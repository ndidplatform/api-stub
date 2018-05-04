package main

import (
	"net/http"
	"net/url"

	"encoding/json"

	"github.com/gorilla/mux"
)

type AsService struct {
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	MinIal      int    `json:"min_ial"`
	MinAal      int    `json:"min_aal"`
	Url         string `json:"url"`
}

func AddUpdateAsService(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var asService AsService
	if err := json.NewDecoder(r.Body).Decode(&asService); err != nil {
		panic(err)
	}
	errs := url.Values{}

	if asService.ServiceId == "" {
		errs.Add("service_id", "The service_id field is required!")
	}
	if asService.ServiceName == "" {
		errs.Add("service_name", "The service_name field is required!")
	}
	if asService.MinIal == 0 {
		errs.Add("min_ial", "The min_ial field is required or greter than 0")
	}
	if asService.MinAal == 0 {
		errs.Add("min_aal", "The min_aal field is required greter than 0")
	}
	if asService.Url == "" {
		errs.Add("url", "The url field is required!")
	}
	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

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

	if serviceId == "xxx" { // return 404
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Service Not Found"))
	} else { // return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"service_id":` + serviceId + ` ,"service_name": "string","min_ial": 0,"min_aal": 0,"url": "string"}`))
	}
}
