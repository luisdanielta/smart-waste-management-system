#include "main.h"

/* types gpio */
typedef enum
{
    input = 0,
    output = 1,
} gpio_mode_t;

typedef struct
{
    float distance;
    uint32_t timeout;
    uint32_t startTimeout;
    uint32_t endTimeout;
    uint32_t dataRaw;
} ultrasonic_data_t;