package review

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/training_project/controller/review/struct"
	"github.com/training_project/model/global"
)

func CheckDataExist(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-ID, X-Device, X-Method, Date, Req-Date, Authorization, X-TKPD-DEBUG, Cookie")

	// check required parameters
	shopIDString := r.FormValue("shop_id")

	if shopIDString == "" {
		// create failed response
		setResponse(w, "Failed", "Params Empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//conver parameters to specific type data
	shopID, err := strconv.ParseInt(shopIDString, 10, 64)
	if err != nil {
		// create failed response
		setResponse(w, "Failed", "Error parsing params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reviewStruct.PassParams(shopID)

	// passing parameters to struct Data
	reviewInterface := reviewStruct.GetStruct()

	if !reviewInterface.Exist() {
		// create failed response
		setResponse(w, "Failed", "Data is Not Exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create succes response
	setResponse(w, "Succes", "Data Exist")
	w.WriteHeader(http.StatusOK)
	return

}

func setResponse(w http.ResponseWriter, Status string, Message string) {
	resp := global.Response{}
	resp.Status = Status
	resp.Message = Message
	json.NewEncoder(w).Encode(resp)
}
