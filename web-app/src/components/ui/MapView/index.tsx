import Map, { Marker, NavigationControl } from "react-map-gl/maplibre";
// import maplibregl from "maplibre-gl";

import "./styles.css";
import MarkerLayout from "./Marker";
import React, { useEffect, useState /* , useRef */ } from "react";

const apiDeviceList: string = "http://localhost:9095/api/device/pi-pico/all";

const getData = async () => {
  const response = await fetch(apiDeviceList);
  const data = await response.json();
  return data.data;
};

/* parse location "10.3710218,-75.4647328" */
const parseLocation = (location: string) => {
  const [longitude, latitude] = location.split(",");
  return { latitude: parseFloat(latitude), longitude: parseFloat(longitude) };
};

export default function MapView() {
  const [viewState, setViewState] = React.useState({
    longitude: -75.46552311679915,
    latitude: 10.370463945203852,
    zoom: 19,
    bearing: 178.03357930521588,
    padding: { top: 0, bottom: 0, left: 0, right: 0 },
    pitch: 60,
  });

  const [devices, setDevices] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getData().then((data) => {
      setDevices(data);
    });
    setLoading(false);
  }, []);

  /* marker */

  // const markerRef = useRef<maplibregl.Marker>(null);
  /* 
  useEffect(() => {
    console.log(viewState);
  }, [viewState]);
  */
  return (
    <div className="map_view">
      <Map
        {...viewState}
        style={{ height: 700, borderRadius: "0.5rem" }}
        onMove={(evt) => setViewState(evt.viewState)}
        mapStyle="https://api.maptiler.com/maps/streets/style.json?key=hpIsP2oOBBR32z1kce5w" // I Know, I Know, I Know
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
        {/* 
        <Marker
          longitude={-75.4658906868683}
          latitude={10.37037350164151}
          draggable
          onDragStart={() => null}
          onDrag={() => null}
          onDragEnd={() => console.log(markerRef.current?.getLngLat())}
          ref={markerRef}
        >
          <div
            style={{
              width: "1rem",
              height: "1rem",
              backgroundColor: "red",
              borderRadius: "50%",
            }}
          />
        </Marker>
          */}
      </Map>
    </div>
  );
}