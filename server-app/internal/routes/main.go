package routes

import (
	"log"
	"net/http"
	"swms-sa/internal/handlers"
	"swms-sa/internal/handlers/devices"
)

type Route struct {
	Url    string
	Func   []http.HandlerFunc
	Method []string
}

func Handler(funcs []http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* check method */
		for i := 0; i < len(methods); i++ {
			if r.Method == methods[i] {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
				w.Header().Set("Access-Control-Allow-Methods", methods[i])
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				log.Printf("url: %s - method: %s", r.URL.Path, r.Method)
				funcs[i](w, r)
				return
			}
		}
		/* method not allowed */
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
	}
}

var MakeRoute = []Route{ // List of the routes on server
	{ /* / - GET */
		Url:    "/",
		Func:   []http.HandlerFunc{handlers.HandleIndex},
		Method: []string{"GET"},
	},

	/* /api/esp32/device - GET, POST, DELETE */
	{
		Url:    "/api/device/esp32",
		Func:   []http.HandlerFunc{devices.DeviceEspGet},
		Method: []string{"GET"},
	},
	{
		Url:    "/api/device/esp32/all",
		Func:   []http.HandlerFunc{devices.DeviceEspGetAll},
		Method: []string{"GET"},
	},
	{
		Url:    "/api/device/esp32/add",
		Func:   []http.HandlerFunc{devices.DeviceEspPost},
		Method: []string{"POST"},
	},
	{
		Url:    "/api/device/esp32/delete",
		Func:   []http.HandlerFunc{devices.DeviceEspDelete},
		Method: []string{"DELETE"},
	},
	/* /api/pi-pico/device - GET, POST, DELETE */
	{
		Url:    "/api/device/pi-pico/all",
		Func:   []http.HandlerFunc{devices.DevicePicoGetAll},
		Method: []string{"GET"},
	},
	{
		Url:    "/api/device/pi-pico/add",
		Func:   []http.HandlerFunc{devices.DevicePicoPost},
		Method: []string{"POST"},
	},
}
