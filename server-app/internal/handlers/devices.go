package handlers

import "net/http"

func DevicePost(w http.ResponseWriter, r *http.Request) {

	data := struct {
		DeviceName string `json:"device_name"`
		DeviceType string `json:"device_type"`
	}{}

	res := Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Device added successfully",
		Error:      Error_t{},
		Application: Application_t{
			Type:   "application/json",
			Url:    r.URL.Path,
			Method: r.Method,
		},
	}
	res.Send(w)
}

func DeviceGet(w http.ResponseWriter, r *http.Request) {
	data := struct {
		DeviceName string `json:"device_name"`
		DeviceType string `json:"device_type"`
	}{}

	res := Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Device added successfully",
		Error:      Error_t{},
		Application: Application_t{
			Type:   "application/json",
			Url:    r.URL.Path,
			Method: r.Method,
		},
	}
	res.Send(w)
}
