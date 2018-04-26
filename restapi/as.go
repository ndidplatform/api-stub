package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type AsService struct {
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	MinIal      string `json:"min_ial"`
	MinAal      string `json:"min_aal"`
	Url         string `json:"url"`
}

func AddUpdateAsService(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var url IdpCallback
	_ = decoder.Decode(&url)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update Successful"))

	//TODO if return 201
	//w.WriteHeader(StatusCreated)
	//w.Write([]byte("Create Successful"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error something is wrong"))

}

func GetAsServiceInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	serviceId := params["service_id"]
	//TODO Check input
	println(serviceId)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"service_id\": \"string\",\"service_name\": \"string\",\"min_ial\": 0,\"min_aal\": 0,\"url\": \"string\"}"))

	//TODO if return 404
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Service Not Found"))
}
