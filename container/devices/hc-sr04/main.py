import serial
from matplotlib import pyplot as plt
import time
import numpy as np


class Device:
    port: str  # serial port COM4 - \dev\ttyACM0
    baudrate: int  # 9600 - 115200
    device: serial.Serial

    def __init__(self, port, baudrate) -> None:
        self.port = port
        self.baudrate = baudrate

    def connect(self) -> None:
        try:
            self.device = serial.Serial(self.port, self.baudrate)
            print("Connected to " + self.device.name)
        except Exception as e:
            print(e)

    def disconnect(self) -> None:
        self.device.close()

    def read(self) -> str:
        if self.device.is_open:
            data = self.device.readline()  # read a '\n' terminated line
            return data.strip().decode("utf-8", "ignore")
        return None

    def write(self, data: str) -> None:
        self.device.write(data.encode("utf-8"))


esp32 = Device("/dev/cu.usbserial-0001", 115200)
esp32.connect()
time.sleep(1)  # wait for the device to be ready


def splitData(value: str) -> list:
    # validate initial character is "D"
    if value[0] != "D":
        return

    value = value.split(",")  # Value #1: Data raw, Value #2: Time
    value = [
        float(value[0].split(":")[1]),
        float(value[1].split(":")[1]),
    ]  # Value #1: Data raw, Value #2: Time
    return value


d = []
t = []

start_time = time.time()
while time.time() - start_time < 5:
    value = esp32.read()
    value = splitData(value)
    if value is not None:
        print(value)
        d.append(value[0])
        t.append(value[1])

esp32.disconnect()

# Plot
d = np.array(d)
t = np.array(t)

plt.plot(d, t)
plt.show()
