package review

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/training_project/model/global"
	"github.com/training_project/model/review/instance"
)

func CheckDataExist(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-ID, X-Device, X-Method, Date, Req-Date, Authorization, X-TKPD-DEBUG, Cookie")

	// check required parameters
	shopIDString := r.FormValue("shop_id")

	if shopIDString == "" {
		// create failed response
		w.WriteHeader(http.StatusBadRequest)
		setResponse(w, "Failed", "Params Empty")
		return
	}

	//conver parameters to specific type data
	shopID, err := strconv.ParseInt(shopIDString, 10, 64)
	if err != nil {
		// create failed response
		w.WriteHeader(http.StatusBadRequest)
		setResponse(w, "Failed", "Error parsing params")
		return
	}

	reviewInstance.PassParams(shopID)

	// passing parameters to struct Data
	review := reviewInstance.GetReviewInstance()

	if !review.Exist() {
		// create failed response
		w.WriteHeader(http.StatusBadRequest)
		setResponse(w, "Failed", "Data is Not Exist")
		return
	}

	// create succes response
	w.WriteHeader(http.StatusOK)
	setResponse(w, "Succes", "Data Exist")
	return

}

func setResponse(w http.ResponseWriter, Status string, Message string) {
	resp := global.Response{}
	resp.Status = Status
	resp.Message = Message
	json.NewEncoder(w).Encode(resp)
}
