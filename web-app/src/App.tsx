import "./app.css";
import Layout from "./components/Layout";
import Device from "./components/ui/Device";
import { status_t } from "./components/ui/Device/types";

import { useState, useEffect } from "react";
import MapView from "./components/ui/MapView";
import "maplibre-gl/dist/maplibre-gl.css";

const testD: status_t = {
  statusValue: "Online",
  statusBattery: "100",
  statusSignal: "100",
};

const apiDeviceList: string = "http://localhost:9095/api/device/pi-pico/all";

const getData = async () => {
  const response = await fetch(apiDeviceList);
  const data = await response.json();
  return data.data;
};

export default function App() {
  const [devices, setDevices] = useState([]);

  useEffect(() => {
    getData().then((data) => {
      setDevices(data);
    });
  }, []);

  return (
    <main className="container">
      <Layout />
      <div className="container_main">
        <section className="container_info">
          <MapView />
          <section className="widget_view">
            <div>
              <p className="title">Containers - Total: {devices.length}</p>
              <div>
                {devices.map((device: any) => (
                  <Device
                    key={device.ID}
                    idNumber={device.name}
                    nameCT={device.controller_id}
                    status={testD}
                  />
                ))}
              </div>
            </div>
            <div className="container_notifications"></div>
          </section>
        </section>
        <section className="container_test"></section>
      </div>
    </main>
  );
}
