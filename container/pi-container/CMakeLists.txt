cmake_minimum_required(VERSION 3.13)

# initialize the SDK based on PICO_SDK_PATH
# note: this must happen before project()
include(pico_sdk_import.cmake)
set(PICO_CXX_ENABLE_EXCEPTIONS 1)
pico_sdk_init()

# Set the path to the Pico SDK

project(pi_container C CXX ASM)

add_executable(main
    main.cpp
)


# add files motor_control.c and motor_control.h
target_sources(main PRIVATE types.h core.h ultra_sonic.h)

# Add pico_stdlib library which aggregates commonly used features
target_link_libraries(main pico_stdlib pico_multicore)

# Add hardware pwm library
target_link_libraries(main hardware_gpio hardware_i2c hardware_timer hardware_pwm hardware_clocks)

# config uart output USB
pico_enable_stdio_usb(main 1)
pico_enable_stdio_uart(main 0)

# create map/bin/hex/uf2 file in addition to ELF.
pico_add_extra_outputs(main)