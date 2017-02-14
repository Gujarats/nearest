package review

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("shop_id = ", shopIDString)

	if shopIDString == "" {
		fmt.Println("empty params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//conver parameters to specific type data
	shopID, err := strconv.ParseInt(shopIDString, 10, 64)
	if err != nil {
		fmt.Println("Error convert params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reviewStruct.PassParams(shopID)

	// passing parameters to struct Data
	reviewInterface := reviewStruct.GetStruct()

	if !reviewInterface.Exist() {

		// create failed response
		resp := global.Response{}
		resp.Status = "Failed"
		resp.Message = "Data is Not Exist"

		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create succes response
	resp := global.Response{}
	resp.Status = "Success"
	resp.Message = "Data Exist"

	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
	return

}
