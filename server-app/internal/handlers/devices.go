package handlers

import (
	"net/http"
	"swms-sa/pkg"
	tp "swms-sa/types"
)

func DevicePost(w http.ResponseWriter, r *http.Request) {
	var db = pkg.ConnDB.GetConn()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	pkg.ConnDB.CloseConn(db)
}

func DeviceGet(w http.ResponseWriter, r *http.Request) {
	data := struct {
		DeviceName string `json:"device_name"`
		DeviceType string `json:"device_type"`
	}{}

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Device added successfully",
		Error:      tp.Error_t{},
		Application: tp.Application_t{
			Type: "application/json",
		},
	}
	res.Send(w)
}
