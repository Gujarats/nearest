package driver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/training_project/util"
	"github.com/training_project/util/logger"

	"github.com/training_project/model/driver/instance"
	"github.com/training_project/model/global"
)

// find specific driver with their ID or name.
// if the desired data didn't exist then insert new data
func UpdateDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	name := r.FormValue("name")
	lat := r.FormValue("latitude")
	lon := r.FormValue("longitude")
	status := r.FormValue("status")

	isAllExist := util.CheckValue(name, lat, lon, status)
	if !isAllExist {
		logger.PrintLog("Required Params Empty")

		//return Bad response
		w.WriteHeader(http.StatusBadRequest)
		global.SetResponse(w, "Failed", "Required Params Empty")
		return
	}

	// convert string to bool
	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		//return Bad response
		logger.PrintLog("Failed to Parse Boolean")
		w.WriteHeader(http.StatusBadRequest)
		global.SetResponse(w, "Failed", "Parse Boolean Erro")
		return
	}

	// get instance
	driver := driverInstance.GetInstance()

	driver.Update(name, lat, lon, statusBool)

	//return succes response
	w.WriteHeader(http.StatusOK)
	global.SetResponse(w, "Succes", "Driver Inserted")
	return

}

func FindDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	name := r.FormValue("name")

	checkValue := util.CheckValue(name)
	if !checkValue {
		logger.PrintLog("Required Params Empty")

		//return Bad response
		w.WriteHeader(http.StatusBadRequest)
		global.SetResponse(w, "Failed", "Required Params Empty")
		return
	}

	// get instance
	driver := driverInstance.GetInstance()

	driverData := driver.Find(name)

	if driverData.Name == "" {
		//return Bad response
		w.WriteHeader(http.StatusOK)
		global.SetResponse(w, "Success", "Data Not Found")
		return
	}

	//return succes response
	w.WriteHeader(http.StatusOK)
	response := global.Response{Status: "Success", Message: "Data Found", Data: driverData}
	json.NewEncoder(w).Encode(response)
	return
}

func InsertDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// getting the parameters
	name := r.FormValue("name")
	lat := r.FormValue("latitude")
	lon := r.FormValue("longitude")
	status := r.FormValue("status")

	isAllExist := util.CheckValue(name, lat, lon, status)
	if !isAllExist {
		logger.PrintLog("Required Params Empty")

		//return Bad response
		w.WriteHeader(http.StatusBadRequest)
		global.SetResponse(w, "Failed", "Required Params Empty")
		return
	}

	// convert string to bool
	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		//return Bad response
		logger.PrintLog("Failed to Parse Boolean")
		w.WriteHeader(http.StatusBadRequest)
		global.SetResponse(w, "Failed", "Parse Boolean Erro")
		return
	}

	// get Driver instance
	driver := driverInstance.GetInstance()

	driver.Insert(name, lat, lon, statusBool)

	//return succes response
	w.WriteHeader(http.StatusOK)
	global.SetResponse(w, "Succes", "Driver Inserted")
	return
}
