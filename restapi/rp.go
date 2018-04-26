package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type requestForConsent struct {
	Namespace  string `json:"namespace"`
	Identifier string `json:"identifier"`
	Timeout    string `json:"timeout"`

	ReferenceNumber string `json:"namespace"`
	IdpList         string `json:"idp_list"` //array of string i.e. "idp_list": ["string"]
	CallbackUrl     string `json:"callback_url"`
	AsServiceList   string `json:"as_service_list"` // array of object i.e. "as_service_list": [{"service_id": "string", "as_id": ["string"], "count": 0, "params": "string"}]
	RequestMessage  string `json:"request_message"`
	MinIal          string `json:"min_ial"`
	MinAal          string `json:"min_aal"`
	MinIdp          string `json:"min_idp"`
	RequestTimeout  string `json:"namespace"`
}

func RequestToId(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new identifiers
	_ = decoder.Decode(&new)
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]

	//TODO Check input
	println(namespace)
	println(identifier)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request Successful"))

	//TODO if return 202
	//w.WriteHeader(http.StatusAccepted)
	//w.Write([]byte("Request Accepted – Async processing, please check back or wait for response at Callback URL"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Request"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))
}

func FetchRequestStatus(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new identifiers
	_ = decoder.Decode(&new)
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]

	//TODO Check input
	println(namespace)
	println(identifier)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"request_type\": \"consent\", \"request_message\": \"string\", \"min_ial\": 0, \"min_aal\": 0, \"min_idp\": 0, \"service_id_list\": [\"string\"], \"timeout\": 0, \"status\": \"open\"}"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("No request Exist"))
}

func GetRefNo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new identifiers
	_ = decoder.Decode(&new)
	defer r.Body.Close()
	params := mux.Vars(r)
	refno := params["reference_number"]

	//TODO Check input
	println(refno)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ex. reference no"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("No reference number in the system"))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new identifiers
	_ = decoder.Decode(&new)
	defer r.Body.Close()
	params := mux.Vars(r)
	reqId := params["request_id"]

	//TODO Check input
	println(reqId)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"source_node_id\": \"string\", \"service_id\": \"string\", \"source_signature\": \"string\", \"data\": \"string\"}]"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	w.Write([]byte("Request does not exist (Bad/incorrect" + reqId + ")"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("No data found"))
}
