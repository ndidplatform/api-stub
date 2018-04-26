package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RetrieveAllIdpId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	minIal := params["min_ial"]
	minAal := params["min_aal"]
	//TODO Check input
	println(minIal)
	println(minAal)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"idp_name\": \"string\",\"idp_id\": \"string\",\"min_ial\": 0,\"min_aal\": 0}]"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: MIN IAL, AAL value"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: No IDP found"))

}

func RetrieveAllIdpNodeId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	minIal := params["min_ial"]
	minAal := params["min_aal"]
	//TODO Check input
	println(namespace)
	println(identifier)
	println(minIal)
	println(minAal)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"idp_name\": \"string\",\"idp_id\": \"string\",\"min_ial\": 0,\"min_aal\": 0}]"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Identity or MIN IAL, AAL value"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: No IDP found"))

}

func RetrieveListGivenServiceId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	serviceId := params["service_id"]
	//TODO Check input
	println(serviceId)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"as_name\": \"string\",\"as_id\": \"string\",\"min_ial\": 0,\"min_aal\": 0}]"))
	//TODO if return 404 //not define in swaggerhub
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Service Not Found"))
}
