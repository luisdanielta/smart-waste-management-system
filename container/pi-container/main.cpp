#include "main.h"

int main()
{
    /* setup app */
    stdio_init_all();
    App app;
    app.setup();

    /* run app */
    while (true)
        app.loop();

    return 0;
}

/* setup app */
void App::setup()
{
    /* setup hardware */
    hardware.led = new Gpio(15, GPIO_FUNC_SIO, output);
}

/* run app */
void App::loop()
{
    ultrasonic_data_t ultrasonicData = ultrasonic.getData();
    hardware.led->setValue(true);
    printf("distance: %f\n", ultrasonicData.distance);
}
