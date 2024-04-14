Servisler:

Kafka servisi: 
Kafka servisinin temel amacı diğer servislerden http requesti ile mesaj alıp localhost:9092 portunda çalışan kafka broker'ına ileterek
mesajı kuyruklamaktır. Diğer servisler de bu broker'a bağlanıp mesajları consume edebilir. 

Endpoint: 
POST localhost:8081/sendMessageToKafka

Örnek request body:
{
    "message": "lat:-80.084128,lon:32.590355,mag:8.544139", //mesajımız (consume eden servise göre formatlanır,consume edecek servisler bu formatta işleyecek)
    "brokerAddr": "kafka:9092",  							//istek atılacak broker
    "topic": "earthquake",									//mesaj gönderilecek topic
    "partition": 0											//mesaj gönderilecek partition
}
Response listesi:
-StatusBadRequest -> gönderilen request yanlış formatlanırsa
-StatusUnprocessableEntity -> belirtilen broker, topic veya partition yoksa
-InternalServerError -> request kaynaklı  değil, server kaynaklı bir sorun oluştuysa
-StatusOk -> mesaj başarıyla eklendiyse

--------------------------------------------------------------------------------------------

Earthquake servisi:
Earthquake servisi kafka servisine mesaj produce eder. 3 adet endpoint bulunmakta. "localhost:8080/inputEarthquake" json olarak kullanıcıdan deprem verileri girmesini  ister, "localhost:8080/startRandomEarthquake" 2 saniyede bir sürekli rastgele deprem oluşturur ve bize deprem oluşturmayı durdurmak için bir id döner, 
"localhost:8080/stopRandomEarthquake/id" ise bizden id'yi alır ve deprem üretmeyi durdurur.

Endpoint:
POST localhost:8080/inputEarthquake

Örnek request body:
{
    "lat": 24.771959, //enlem verisi
    "lon": 46.217018, //boylam verisi
    "mag": 8		  //deprem şiddeti
}

Response listesi
-StatusBadRequest -> gönderilen request yanlış formatlanırsa
-StatusUnprocessableEntity -> enlem, boylam, şiddet verileri yanlış aralıktaysa (min magnitude = 1, max magnitude = 10)
-StatusInternalServerError -> kafka servisine http requesti yollayamıyorsa
-StatusOk -> Deprem başarıyla oluştuysa

**********************************

Endpoint:
POST localhost:8080/startRandomEarthquake

Örnek request body
Body gönderilmez

Response listesi:
id -> deprem oluşturucunun id'si, durdurmak için bu id'yi kullanırız.

**********************************

Endpoint
DEL localhost:8080/stopRandomEarthquake/id

Örnek request body
Body gönderilmez

Response listesi:
StatusBadRequest -> id gönderilmediyse
StatusUnprocessableEntity -> olmayan bir id gönderildiyse
StatusOK -> deprem oluşturucu başarıyla durdurulduysa

----------------------------------------------------------------------------------------------


Record_earthquake servisi
Bu servis kafka broker'da kuyruklanmış tüm verileri consume eder ve postgresql database'inde depolar. 

Endpoint:
GET localhost:8082/getEarthquakes

Örnek request body
Body gönderilmez

Response listesi:
StatusInternalServerError -> database'den veri çekerken hata oluşursa
StatusOK -> veriler başarıyla getirilirse

---------------------------------------------------------------------------------------------

Filtering Earthquake servisi
Bu servis de kafkada kuyruklanmış verileri consume eder, ancak verileri consume ederken yalnızca 3.0'dan büyük depremleri önemser ve SSE (server sent event) bağlantısı açarak front-end'e oluşturulan depremlerin geohash değerini ve o geohash'da kaç deprem olduğunu döndürür. Geohash doğruluğu 4 karakter olarak ayarlanmıştır, 40 km ve çevresindeki depremleri aynı bölge olarak sayar. Front-end'de de istek atılırken geohash karakterinin 4 olarak ayarlanması gerekmektedir.

Endpoint:
localhost:6663/events/message -> SSE'e bağlanmak için

Response listesi
"geohash,deprem_sayisi" döner örnek olarak "3daw1,3" şeklinde veriler döner


******************************

Endpoint:
localhost:6663/getOldKeys -> client SSE'e bağlanmamışken oluşturulan geohash'leri döner

Response listesi:
-