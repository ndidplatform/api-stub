package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type IdpCallback struct {
	Url string `json:"url"` //API should be revised
}

type IdpResponseData struct {
	Status     string `json:"status"`
	RequestId  string `json:"request_id"`
	Namespace  string `json:"namespace"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	Aal        string `json:"aal"`
	Signature  string `json:"signature"`
	AccessorId string `json:"accessor_id"`
}

func RetrieveCallbackUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	//TODO Check input
	println(namespace)
	println(identifier)
	//TODO if return 201
	//w.WriteHeader(http.StatusOK) // Shoud be reviewed
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"url\": \"string\"}"))
	//TODO if return 404 // Shoud be reviewed
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("ID not found"))
}

func SetUrlCallback(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var url IdpCallback
	_ = decoder.Decode(&url)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 200 // All response should be reviewed
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Identity Created"))

	//TODO if return 400 // All response should be reviewed
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid accessor method or accessor value"))

	//TODO if return 403 // All response should be reviewed
	//w.WriteHeader(http.StatusForbidden)
	//w.Write([]byte("Error: Identity already exist"))

}

func IdpResponse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var response IdpResponseData
	_ = decoder.Decode(&response)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Response Successful"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Response"))

	//TODO if return 404 // All
	//w.WriteHeader(http.StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))

}
