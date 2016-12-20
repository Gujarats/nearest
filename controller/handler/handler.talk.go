package handler

import (
	"encoding/json"
	"fmt"
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
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Content-Type", "application/vnd.api+json")
	fmt.Println("hellloooo")
	var m model.Message
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	fmt.Println("hellloooo2")
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 400)

		fmt.Println("Masuk Error")
		return
	}

	fmt.Println("hellloooo3")

	fmt.Println("ShopId = ", m.ShopId)
	return
}
