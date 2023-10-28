/* This type is used to define the status of the device. */
type status_t = {
  statusValue: string;
  statusBattery: string;
  statusSignal: string;
};

interface DeviceProps {
  idNumber: string;
  nameCT: string;
  status: status_t;
  className?: string;
}

export type { DeviceProps, status_t };