import { DeviceProps } from "./types";
import "./styles.css"

export default function Device({ idNumber, nameCT, status }: DeviceProps) {
  return (
    <div className="container_device">
      <span>
        <p>Container Number: {idNumber}</p>
      </span>
      <span className="container_status_device">
        <p className={
          status.statusValue === "Online" ? "status_p" : "status_p_offline"
        }>{status.statusValue}</p>
        <p className={
          status.statusValue === "Online" ? "name_ct_p" : "name_ct_p_offline"
        }>{nameCT}</p>
      </span>
      <span>
        <p>
          Battery: {status.statusBattery}% | Signal: {status.statusSignal}%
        </p>
      </span>
    </div>
  );
}
