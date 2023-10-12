#include "main.h"

int main()
{
    /* create app */
    Gpio sys(25, GPIO_FUNC_SIO, output);
    App app(&sys);

    /* setup app */
    app.setup();

    /* run app */
    while (true)
        app.loop();

    return 0;
}

/* setup app */
void App::setup()
{
    /* setup system led */
    config.system->setValue(true);
}

/* run app */
void App::loop()
{
    /* toogle system led */
    config.system->setValue(!config.system->getValue());
    sleep_ms(100);
}