const loadGeohashLibrary = async () => {
  const ngeohashModule = await import('https://cdn.skypack.dev/ngeohash');
  return ngeohashModule.default;
};

const main = async () => {
  const ngeohash = await loadGeohashLibrary();
  const map = L.map('map').setView([38.963745, 35.243322], 6);
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  }).addTo(map);

  const oldDataRequest = await fetch("http://localhost:6663/getOldKeys");
  const oldData = await oldDataRequest.text();
  oldData.split("\n").forEach(data => {
    if (data) {
      const [geohash, depremSayisi] = data.split(",");
      const decoded = ngeohash.decode(geohash.slice(0, 4));
      L.marker([decoded.latitude, decoded.longitude])
        .addTo(map)
        .bindPopup(`Bolgede olusan 3.0'dan buyuk deprem sayisi: ${depremSayisi}`)
        .openPopup();
    }
  });

  const evtSource = new EventSource("http://localhost:6663/events/message");
  evtSource.addEventListener("message", event => {
    const [geohash, depremSayisi] = event.data.split(",");
    const decoded = ngeohash.decode(geohash.slice(0, 4));
    L.marker([decoded.latitude, decoded.longitude])
      .addTo(map)
      .bindPopup(`Bolgede olusan 3.0'dan buyuk deprem sayisi: ${depremSayisi}`)
      .openPopup();
  });

  evtSource.addEventListener("close", () => {
    setTimeout(() => {
      new EventSource("http://localhost:6663/events/message");
    }, 1000);
  });
};

main();

