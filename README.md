<h1>Servisler:</h1>

<h2>Kafka servisi:</h2> 
Kafka servisinin temel amacı diğer servislerden http requesti ile mesaj alıp kafka broker'ına ileterek
mesajı kuyruklamaktır. Diğer servisler de bu broker'a bağlanıp mesajları consume edebilir. 

<h3>Endpoint:</h3>
POST localhost:8081/sendMessageToKafka

<h3>Örnek request body:</h3>
{
    "message": "lat:-80.084128,lon:32.590355,mag:8.544139", //mesajımız (bu formatta olmalı) <br>
    "brokerAddr": "kafka:9092",  							//istek atılacak broker <br>
    "topic": "earthquake",									//mesaj gönderilecek topic <br>
    "partition": 0											//mesaj gönderilecek partition <br>
}
<h3>Response listesi:</h3>
-StatusBadRequest -> gönderilen request yanlış formatlanırsa
-StatusUnprocessableEntity -> belirtilen broker, topic veya partition yoksa
-InternalServerError -> request kaynaklı  değil, server kaynaklı bir sorun oluştuysa
-StatusOk -> mesaj başarıyla eklendiyse

--------------------------------------------------------------------------------------------

<h2>Earthquake servisi:</h2>
Earthquake servisi kafka servisine mesaj produce eder. 3 adet endpoint bulunmakta. "localhost:8080/inputEarthquake" json olarak kullanıcıdan deprem verileri girmesini  ister, "localhost:8080/startRandomEarthquake" 2 saniyede bir sürekli rastgele deprem oluşturur ve bize deprem oluşturmayı durdurmak için bir id döner, 
"localhost:8080/stopRandomEarthquake/id" ise bizden id'yi alır ve deprem üretmeyi durdurur.

<h3>Endpoint:</h3>
POST localhost:8080/inputEarthquake

<h3>Örnek request body:</h3>
{
    "lat": 24.771959, //enlem verisi
    "lon": 46.217018, //boylam verisi
    "mag": 8		  //deprem şiddeti
}

<h3>Response listesi</h3>
-StatusBadRequest -> gönderilen request yanlış formatlanırsa
-StatusUnprocessableEntity -> enlem, boylam, şiddet verileri yanlış aralıktaysa (min magnitude = 1, max magnitude = 10)
-StatusInternalServerError -> kafka servisine http requesti yollayamıyorsa
-StatusOk -> Deprem başarıyla oluştuysa

<h3>Endpoint:</h3>
POST localhost:8080/startRandomEarthquake

<h3>Örnek request body</h3>
Body gönderilmez

<h3>Response listesi:</h3>
id -> deprem oluşturucunun id'si, durdurmak için bu id'yi kullanırız.

<h3>Endpoint</h3>
DEL localhost:8080/stopRandomEarthquake/id

</h3>Örnek request body</h3>
Body gönderilmez

</h3>Response listesi:</h3>
StatusBadRequest -> id gönderilmediyse
StatusUnprocessableEntity -> olmayan bir id gönderildiyse
StatusOK -> deprem oluşturucu başarıyla durdurulduysa

----------------------------------------------------------------------------------------------


<h2>Record_earthquake servisi</h2>
Bu servis kafka broker'da kuyruklanmış tüm verileri consume eder ve postgresql database'inde depolar. 

<h3>Endpoint:</h3>
GET localhost:8082/getEarthquakes

<h3>Örnek request body</h3>
Body gönderilmez

<h3>Response listesi:</h3>
StatusInternalServerError -> database'den veri çekerken hata oluşursa
StatusOK -> veriler başarıyla getirilirse

---------------------------------------------------------------------------------------------

<h2>Filtering Earthquake servisi</h2>
Bu servis de kafkada kuyruklanmış verileri consume eder, ancak verileri consume ederken yalnızca 3.0'dan büyük depremleri önemser ve SSE (server sent event) bağlantısı açarak front-end'e oluşturulan depremlerin geohash değerini ve o geohash'da kaç deprem olduğunu döndürür. Geohash doğruluğu 4 karakter olarak ayarlanmıştır, 40 km ve çevresindeki depremleri aynı bölge olarak sayar. Front-end'de de istek atılırken geohash karakterinin 4 olarak ayarlanması gerekmektedir.

<h3>Endpoint:</h3>
localhost:6663/events/message -> SSE'e bağlanmak için

<h3>Response listesi</h3>
"geohash,deprem_sayisi" döner örnek olarak "3daw1,3" şeklinde veriler döner

<h3>Endpoint:</h3>
localhost:6663/getOldKeys -> client SSE'e bağlanmamışken oluşturulan geohash'leri döner

<h3>Response listesi:</h3>
-
<br><br>
<h1>Kullanım:</h1>
git clone git@github.com:pragmataW/kartaca-earthquake.git <br>
cd kartaca-earthquake <br>
docker-compose up <br>
