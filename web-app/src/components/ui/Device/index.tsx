import { DeviceProps } from "./types";

export default function Device({ idNumber, nameCT, status }: DeviceProps) {
  return (
    <div className="container_device">
      <span>
        <p>Container Number: {idNumber}</p>
      </span>
      <span>
        <p>{status.statusValue}</p>
        <p>{nameCT}</p>
      </span>
      <span>
        <p>
          Battery: {status.statusBattery} | Signal: {status.statusSignal}
        </p>
      </span>
    </div>
  );
}
