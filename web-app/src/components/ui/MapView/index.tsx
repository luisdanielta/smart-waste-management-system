import Map, { Marker, NavigationControl } from "react-map-gl/maplibre";
import "./styles.css";
import MarkerLayout from "./Marker";
import React, { useEffect, useState } from "react";

const apiDeviceList: string = "http://localhost:9095/api/device/pi-pico/all";

const getData = async () => {
  const response = await fetch(apiDeviceList);
  const data = await response.json();
  return data.data;
};

/* parse location "10.420877, -75.551078" */
const parseLocation = (location: string) => {
  const [latitude, longitude] = location.split(",");
  return { latitude: parseFloat(latitude), longitude: parseFloat(longitude) };
};

export default function MapView() {
  const [viewState, setViewState] = React.useState({
    longitude: -75.548598,
    latitude: 10.422366,
    zoom: 17,
  });

  const [devices, setDevices] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getData().then((data) => {
      setDevices(data);
    });
    setLoading(false);
  }, []);

  return (
    <div className="map_view">
      <Map
        {...viewState}
        style={{ height: 730 }}
        onMove={(evt) => setViewState(evt.viewState)}
        mapStyle="https://api.maptiler.com/maps/streets/style.json?key=hpIsP2oOBBR32z1kce5w"
      >
        <NavigationControl position="bottom-right" />
        {!loading && (
          <>
            {devices.map((device: any) => (
              <Marker
                key={device.ID}
                longitude={parseLocation(device.location).longitude}
                latitude={parseLocation(device.location).latitude}
              >
                <MarkerLayout name={device.name} />
              </Marker>
            ))}
          </>
        )}
      </Map>
    </div>
  );
}
