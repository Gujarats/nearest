package driver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Gujarats/API-Golang/util"
	"github.com/Gujarats/API-Golang/util/logger"

	"github.com/Gujarats/API-Golang/model/driver/interface"
	"github.com/Gujarats/API-Golang/model/global"
)

// find specific driver with their ID or name.
// if the desired data didn't exist then insert new data
func UpdateDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time for lenght of the process
		startTimer := time.Now()

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

		// convert string to float64
		convertedFloat, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to conver float value")
			return
		}

		latFloat := convertedFloat[0]
		lonFloat := convertedFloat[1]

		driver.Update(name, latFloat, lonFloat, statusBool)

		//return succes response
		elpasedTime := time.Since(startTimer).Seconds()
		w.WriteHeader(http.StatusOK)
		global.SetResponseTime(w, "Succes", "Driver Inserted", elpasedTime)
		return
	})

}

func FindDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time for lenght of the process
		startTimer := time.Now()
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		distance := r.FormValue("distance")

		//checking empty value
		checkValue := util.CheckValue(lat, lon, distance)
		if !checkValue {
			logger.PrintLog("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		floatNumbers, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert float value")
			return
		}
		latFloat := floatNumbers[0]
		lonFloat := floatNumbers[1]

		intNumbers, err := util.ConvertToInt64(distance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert integer value")
			return
		}
		distanceInt := intNumbers[0]

		// get instance
		driverDatas := driver.GetNearLocation(distanceInt, latFloat, lonFloat)

		if len(driverDatas) == 0 {
			//return Bad response
			w.WriteHeader(http.StatusOK)
			elapsedTime := time.Since(startTimer).Seconds()
			response := global.Response{Status: "Success", Message: "Data Found", Latency: elapsedTime, Data: driverDatas}
			json.NewEncoder(w).Encode(response)
			return
		}

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data Found", Latency: elapsedTime, Data: driverDatas}
		json.NewEncoder(w).Encode(response)
		return

	})
}

func InsertDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//start time for lenght of the process
		startTimer := time.Now()
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

		// convert string to float64
		convertedFloat, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert float value")
			return
		}
		latFloat := convertedFloat[0]
		lonFloat := convertedFloat[1]

		// insert driver
		driver.Insert(name, latFloat, lonFloat, statusBool)

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data Inserted", Latency: elapsedTime}
		json.NewEncoder(w).Encode(response)
		return
	})
}
