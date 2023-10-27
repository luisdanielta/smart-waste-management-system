import "./app.css";
import Layout from "./components/Layout";
import Device from "./components/ui/Device";
import { status_t } from "./components/ui/Device/types";


const testD: status_t = {
  statusValue: "Online",
  statusBattery: "100",
  statusSignal: "100"
}


export default function App() {
  return (
    <main className="container">
      <Layout />

      <Device idNumber="123456789" nameCT="CT-01" status={testD} />

    </main>
  );
}
