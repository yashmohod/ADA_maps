async function initMap() {
    // Request needed libraries.
    const { Map } = await google.maps.importLibrary("maps");
    const myLatlng = { lat: 42.423098646583625, lng: -76.49424475970028 };
    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 15.5,
        center: myLatlng,
    });
    // Create the initial InfoWindow.
    let infoWindow = new google.maps.InfoWindow({
        content: "Click the map to get Lat/Lng!",
        position: myLatlng,
    });

    infoWindow.open(map);
    // Configure the click listener.
    map.addListener("click", (mapsMouseEvent) => {
        // Close the current InfoWindow.
        infoWindow.close();
        // Create a new InfoWindow.
        infoWindow = new google.maps.InfoWindow({
            position: mapsMouseEvent.latLng,
        });
        infoWindow.setContent(
            JSON.stringify(mapsMouseEvent.latLng.toJSON(), null, 2),
        );
        infoWindow.open(map);
    });
}

initMap();