package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/training_project/controller"
	"github.com/training_project/model"
	"log"
	"net/http"
	"strconv"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.api+json")

	//set response
	var response model.Response

	//get the parameters
	productId := r.FormValue("product_id")
	if productId == "" {
		response.Status = "Failed : no product id provided"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productIdInt, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := controller.GetTalks(productIdInt)
	response.Data = data

	//return response as JSON
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.api+json")

	//declare response
	var response model.Response

	//get the params
	message := r.FormValue("message")
	if message == "" {
		response.Status = "Failed : message empty"
		//return response as JSON
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert data
	err := controller.InsertTalk(message)
	if err != nil {
		log.Fatal(err)
	}

	//return response as JSON
	response.Status = "OK : data inserted"
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
