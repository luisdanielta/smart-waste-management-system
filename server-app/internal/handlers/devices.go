package handlers

import (
	"net/http"
	"swms-sa/internal/models"
	"swms-sa/pkg"
	tp "swms-sa/types"
)

func DeviceGet(w http.ResponseWriter, r *http.Request) {
	var err tp.Error_t

	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	/* database connection */
	var db = pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.ControllerC{})

	/* get data from database */
	var data models.ControllerC
	err.ERR = db.Where("id = ?", id).First(&data).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	if err.ERR != nil {
		res := tp.Response_t{
			StatusCode: http.StatusNotFound,
			Data:       nil,
			Message:    "Not found",
			Error:      err,
			Application: tp.Application_t{
				Type: "application/json",
			},
		}
		res.Send(w)
		return
	}

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Device found successfully",
		Error:      err,
		Application: tp.Application_t{
			Type: "application/json",
		},
	}

	res.Send(w)

}

func DeviceGetAll(w http.ResponseWriter, r *http.Request) {
	var err tp.Error_t

	/* database connection */
	var db = pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.ControllerC{})

	/* get data from database */
	var data []models.ControllerC
	err.ERR = db.Find(&data).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	if err.ERR != nil {
		res := tp.Response_t{
			StatusCode: http.StatusNotFound,
			Data:       nil,
			Message:    "Not found",
			Error:      err,
			Application: tp.Application_t{
				Type: "application/json",
			},
		}
		res.Send(w)
		return
	}

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

func DevicePost(w http.ResponseWriter, r *http.Request) {

	var err tp.Error_t

	err.ERR = r.ParseForm()
	err.Check()

	if !pkg.ValidateRequiredFields(r, []string{"name", "ip"}) {
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

	// db.AutoMigrate(&models.ControllerC{})

	data := models.ControllerC{
		Name:    r.FormValue("name"),
		Ip:      r.FormValue("ip"),
		Battery: 100,
		Signal:  100,
	}

	data.SetStatus()

	/* add data to database */
	err.ERR = db.Create(&data).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Success",
		Error:      err,
		Application: tp.Application_t{
			Type: "application/json",
		},
	}

	res.Send(w)
}

func DeviceDelete(w http.ResponseWriter, r *http.Request) {
	var err tp.Error_t

	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	/* database connection */
	var db = pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.ControllerC{})

	/* delete data from database */
	err.ERR = db.Where("id = ?", id).Delete(&models.ControllerC{}).Error
	err.Check()
	pkg.ConnDB.CloseConn(db)

	res := tp.Response_t{
		StatusCode: http.StatusOK,
		Data:       nil,
		Message:    "Device deleted successfully",
		Error:      err,
		Application: tp.Application_t{
			Type: "application/json",
		},
	}
	res.Send(w)
}
