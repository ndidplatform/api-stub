package main

import (
	"encoding/json"
	"net/http"
	"net/url"
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
	//params := mux.Vars(r)
	//TODO if return 201
	//w.WriteHeader(http.StatusOK) // Shoud be reviewed
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`"{"url": "string"}`))
	//TODO if return 404 // Shoud be reviewed
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("ID not found"))
}

func SetUrlCallback(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var callback IdpCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		panic(err)
	}

	errs := url.Values{}

	if callback.Url == "" {
		errs.Add("url", "The url field is required!")
	}

	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//TODO Check input
	//TODO if return 200 // All response should be reviewed
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Callback url set"))

	//TODO if return 400 // All response should be reviewed
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid accessor method or accessor value"))

	//TODO if return 403 // All response should be reviewed
	//w.WriteHeader(http.StatusForbidden)
	//w.Write([]byte("Error: Identity already exist"))

}

func IdpResponse(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var response IdpResponseData
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		panic(err)
	}
	errs := url.Values{}

	if response.Status == "" {
		errs.Add("status", "The status field is required!")
	}
	if response.RequestId == "" {
		errs.Add("request_id", "The request_id field is required!")
	}
	if response.Namespace == "" {
		errs.Add("namespace", "The namespace field is required!")
	}
	if response.Identifier == "" {
		errs.Add("identifier", "The identifier field is required!")
	}
	if response.Secret == "" {
		errs.Add("secret", "The secret field is required!")
	}
	if response.Aal == "" {
		errs.Add("aal", "The aal field is required!")
	}
	if response.Signature == "" {
		errs.Add("signature", "The signature field is required!")
	}
	if response.AccessorId == "" {
		errs.Add("accessorId", "The accessorId field is required!")
	}

	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == response.Namespace && id.Identifier == response.Identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Identity does not exist"))
	} else if response.RequestId == "xxx" { // return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid Response"))
	} else { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request Successful"))
	}

}
