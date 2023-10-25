package handlers

import (
	"net/http"
	tp "swms-sa/types"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       nil,
		Message:    "Welcome to SWMS-SA",
		Error:      tp.Error_t{},
		Application: tp.Application_t{
			Type: "application/json",
		},
	}

	res.Send(w)
}
