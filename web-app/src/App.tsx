import "./app.css";
import Layout from "./components/Layout";
import Device from "./components/ui/Device";
import { status_t } from "./components/ui/Device/types";

const testD: status_t = {
  statusValue: "Online",
  statusBattery: "100",
  statusSignal: "100",
};

const deviceList = [
  {
    idNumber: "123456789",
    nameCT: "CT - Marina Park",
    status: testD,
  },
  {
    idNumber: "123456789",
    nameCT: "CT - UTB School",
    status: testD,
  },
  {
    idNumber: "123456789",
    nameCT: "CT - Clock Tower",
    status: testD,
  },
  {
    idNumber: "123456789",
    nameCT: "CT - Bocagrande",
    status: testD,
  },
];

export default function App() {
  return (
    <main className="container">
      <Layout />
      <div className="container_main">
        <section className="container_info">
          <section className="map_view"></section>
          <section className="widget_view">
            <div>
              <p className="title">Containers - Total: 57</p>
              <div>
                {deviceList.map((device, index) => (
                  <Device
                    className="container_devices"
                    key={index}
                    idNumber={device.idNumber}
                    nameCT={device.nameCT}
                    status={device.status}
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
