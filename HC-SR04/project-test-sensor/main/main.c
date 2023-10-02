#include <stdio.h>
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
#include "driver/gpio.h"
#include "esp_log.h"
#include <esp_timer.h>

#include "sdkconfig.h"

/* Use project configuration menu (idf.py menuconfig) to choose the GPIO to blink,
   or you can edit the following line and set a number here.
*/

/* delay */
void set_delay(int ms)
{
    vTaskDelay(ms / portTICK_PERIOD_MS);
}

/* hc-sr04 */

typedef struct measurement_
{
    float distance;
    float time;
} measurement_t;

#define ECHO_GPIO GPIO_NUM_18
#define TRIGGER_GPIO GPIO_NUM_5
float distances[] = {};
bool state_recording = false;

measurement_t hc_sr04_measure_cm(void);

/* function prototype */
esp_err_t hc_sr04_init(void);

float map(float value, float in_min, float in_max, float out_min, float out_max)
{
    return (value - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;
}

// with map and ln print line
char *println(float value, float in_min, float in_max, float out_min, float out_max)
{
    float mapped_value = map(value, in_min, in_max, out_min, out_max);
    char *buffer = malloc(100);
    sprintf(buffer, "D:%.2f,T:%.2f", mapped_value, value);
    printf(buffer);
    return buffer;
}

void app_main(void)
{
    /* initialize hc-sr04 */
    ESP_ERROR_CHECK(hc_sr04_init());
    printf("\033[2J\033[1;1H");
    
    while (true)
    {
        measurement_t measurement_hc_sr04 = hc_sr04_measure_cm();

        /* .2f decimal places D: 0.0 T: 0.0 */
        // printf("D: %.2f - T: %.2f\n", measurement_hc_sr04.distance, measurement_hc_sr04.time);
        // set_delay(10);
        //  print with map and ln print line
        //printf("\033[2J\033[1;1H");
        printf("D:%.2f,T:%.2f\n", map(measurement_hc_sr04.distance, 0, 1000, 0, 100), map(measurement_hc_sr04.time, 0, 1000, 0, 100));
    }
}

/* hc-sr04 */
esp_err_t hc_sr04_init(void)
{
    /* configure trigger pin */
    gpio_config_t trigger_gpio_conf;
    trigger_gpio_conf.intr_type = GPIO_INTR_DISABLE;
    trigger_gpio_conf.mode = GPIO_MODE_OUTPUT;
    trigger_gpio_conf.pin_bit_mask = (1ULL << TRIGGER_GPIO);
    trigger_gpio_conf.pull_down_en = GPIO_PULLDOWN_DISABLE;
    trigger_gpio_conf.pull_up_en = GPIO_PULLUP_DISABLE;
    gpio_config(&trigger_gpio_conf);

    /* configure echo pin */
    gpio_config_t echo_gpio_conf;
    echo_gpio_conf.intr_type = GPIO_INTR_ANYEDGE;
    echo_gpio_conf.mode = GPIO_MODE_INPUT;
    echo_gpio_conf.pin_bit_mask = (1ULL << ECHO_GPIO);
    echo_gpio_conf.pull_down_en = GPIO_PULLDOWN_DISABLE;
    echo_gpio_conf.pull_up_en = GPIO_PULLUP_DISABLE;
    gpio_config(&echo_gpio_conf);

    return ESP_OK;
}

/* measure raw used struct measurement */
measurement_t hc_sr04_measure_cm(void)
{
    measurement_t measurement;

    float raw_data = 0;

    /* trigger sensor */
    gpio_set_level(TRIGGER_GPIO, 1);
    set_delay(10);
    gpio_set_level(TRIGGER_GPIO, 0);

    /* wait for echo */
    while (gpio_get_level(ECHO_GPIO) == 0)
    {
    }

    measurement.time = esp_timer_get_time();
    raw_data = measurement.time;

    /* wait for echo end */
    while (gpio_get_level(ECHO_GPIO) == 1)
    {
    }

    measurement.time = esp_timer_get_time() - measurement.time;

    /* calculate distance */
    //measurement.distance = measurement.time * 0.034 / 2;

    measurement.distance = raw_data;
    return measurement;
}