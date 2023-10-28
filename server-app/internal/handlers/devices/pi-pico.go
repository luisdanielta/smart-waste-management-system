package devices

import (
	"net/http"
	"swms-sa/internal/models"
	"swms-sa/pkg"
	tp "swms-sa/types"
)

func DevicePicoGetAll(w http.ResponseWriter, r *http.Request) {
	var err tp.Error_t

	/* database connection */
	var db = pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.Container{})

	/* get data from database */
	var data []models.Container
	err.ERR = db.Find(&data).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Devices found successfully",
		Error:      err,
		Application: tp.Application_t{
			Type: "application/json",
		},
	}

	res.Send(w)

}

func DevicePicoPost(w http.ResponseWriter, r *http.Request) {
	var err tp.Error_t

	err.ERR = r.ParseForm()
	err.Check()

	if !pkg.ValidateRequiredFields(r, []string{"controller_id", "name", "location", "usecase", "size"}) {
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

	var db = pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.Container{})

	data := models.Container{
		ControllerId: r.FormValue("controller_id"),
		Name:         r.FormValue("name"),
		Location:     r.FormValue("location"),
		Usecase:      r.FormValue("usecase"),
		Size:         r.FormValue("size"),
	}

	err.ERR = db.Create(&data).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Container created successfully",
		Error:      err,
		Application: tp.Application_t{
			Type: "application/json",
		},
	}

	res.Send(w)

}
