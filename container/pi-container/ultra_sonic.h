#pragma once
#include "main.h"

class Ultrasonic // HC-SR04
{
protected:
    ultrasonic_data_t data = {};

    struct UltrasonicConfig
    {
        Gpio *triggerPin;
        Gpio *echoPin;
        /* optional */
        char *name = NULL;
    } config;

    void trigger()
    {
        config.triggerPin->setValue(true);
        sleep_us(10);
        config.triggerPin->setValue(false);
    }

    void waitEcho()
    {
        while (!config.echoPin->getValue())
            ;
        data.startTimeout = time_us_32();

        while (config.echoPin->getValue())
            ;
        data.endTimeout = time_us_32();
    }

    void clearData()
    {
        data = {};
    }

public:
    Ultrasonic(Gpio *triggerPin, Gpio *echoPin)
    {
        config.triggerPin = triggerPin;
        config.echoPin = echoPin;
    }

    ultrasonic_data_t getData()
    {
        clearData();
        trigger();
        waitEcho();
        data.timeout = data.endTimeout - data.startTimeout;
        return data;
    }

    ultrasonic_data_t getDistance()
    {
        ultrasonic_data_t valuePreviews = getData();
        float distanceVP = (valuePreviews.timeout * 0.0343) / 2;

        /* 50 us */
        sleep_us(50);

        ultrasonic_data_t data = getData();
        float distance = (data.timeout * 0.0343) / 2;

        float p = (distanceVP + distance) / 2;
        data.distance = p;
        delete &p, &distanceVP, &distance, &valuePreviews;

        return data;
    }
};