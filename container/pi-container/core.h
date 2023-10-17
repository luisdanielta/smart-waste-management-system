#pragma once
#include "main.h"

class Gpio
{
public:
    /* config gpio - pin, function, pull, drive, etc */
    struct GpioConfig
    {
        uint pin = 0;
        gpio_function function = GPIO_FUNC_SIO;
        bool state = false;
        gpio_mode_t mode = input;
        /* optional */
        char *name = NULL;
        uint value = 0;

    } config;
    Gpio(uint pin, gpio_function fnc, gpio_mode_t mode)
    {
        config.pin = pin ? pin : 0;
        config.mode = mode ? mode : input;
        config.function = fnc ? fnc : GPIO_FUNC_SIO;

        /* set gpio */
        gpio_init(config.pin);

        /* set mode */
        gpio_set_dir(config.pin, config.mode);

        /* set function */
        if (config.function != GPIO_FUNC_SIO)
        {
            gpio_set_function(config.pin, config.function);

            /* switch to alt function
            uint slice_num;
            pwm_config pwmConfig = pwm_get_default_config();
            switch (config.function)
            {
            case GPIO_FUNC_PWM:
                slice_num = pwm_gpio_to_slice_num(config.pin);
                pwm_init(slice_num, &pwmConfig, true);
                pwm_set_wrap(slice_num, 255);
                pwm_set_gpio_level(config.pin, config.value);
                break;
            default:
                printf("gpio function not found...\n");
                break;
            }
            */
        }
    }

    /* reference to gpio */
    void setName(char *name)
    {
        config.name = name;
    }

    /* fns */
    void setValue(bool value)
    {
        if (config.function == GPIO_FUNC_PWM)
        {
            pwm_set_gpio_level(config.pin, value);
            return;
        }
        config.value = value;
        gpio_put(config.pin, config.value);
    }

    bool getValue()
    {
        return gpio_get(config.pin);
    }
};

/* create a function map for ultrasonic */
float map(float x, float in_min, float in_max, float out_min, float out_max)
{
    return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;
}