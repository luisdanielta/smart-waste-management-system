#pragma once
#include <stdint.h>
#include "pico/stdlib.h"
#include "hardware/pwm.h"
#include "hardware/gpio.h"
#include <string>
#include "hardware/clocks.h"
#include "types.h"
#include <cstdlib>
#include "core.h"
#include "ultra_sonic.h"

class App
{
protected:
    struct HardwareConfig_
    {
        Gpio *led = NULL;
        void reset() = delete;
    } hardware;

    struct AppConfig_
    {
        bool dev = false;

        /* hardware */
        HardwareConfig_ hardware;
    } config;

    /* gpio ultrasonic */
    Gpio triggerUltrasonic = Gpio(10, GPIO_FUNC_SIO, output);
    Gpio echoUltrasonic = Gpio(11, GPIO_FUNC_SIO, input);

    void delay(uint32_t ms)
    {
        sleep_ms(ms);
    }

public:
    App()
    {
        /* setup pico sdk */
        stdio_init_all();
    };

    /* loop and setup instance */
    void setup();
    void loop();

    /* hardware */
    Ultrasonic ultrasonic = Ultrasonic(&triggerUltrasonic, &echoUltrasonic);
};