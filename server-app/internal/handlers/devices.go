package handlers

import (
	"net/http"
	"swms-sa/pkg"
	tp "swms-sa/types"
)

func DevicePost(w http.ResponseWriter, r *http.Request) {

	var err tp.Error_t
	err.Check(r.ParseForm())

	if !pkg.ValidateRequiredFields(r, []string{"name", "type", "location", "status"}) {
		res := tp.Response_t{
			StatusCode: http.StatusBadRequest,
			Data:       nil,
			Message:    "All fields are required",
			Error:      err,
			Application: tp.Application_t{
				Type: "application/json",
			},
		}
		res.Send(w)
		return
	}

	/* database connection */
	var db = pkg.ConnDB.GetConn()

	pkg.ConnDB.CloseConn(db)

}
