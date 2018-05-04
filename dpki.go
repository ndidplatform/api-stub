package main

import (
	"net/http"
	"net/url"

	"encoding/json"
)

type UrlDpki struct {
	Url string `json:"url"`
}

/*
type NodeDpki struct {
	NodeId              string `json:"node_id"`
	NodeName            string `json:"node_name"`
	NodeKey             string `json:"node_key"`
	NodeKeyType         string `json:"node_key_type"`
	NodeKeyMethod       string `json:"node_key_method"`
	NodeMasterKey       string `json:"node_master_key"`
	NodeMasterKeyType   string `json:"node_master_key_type"`
	NodeMasterKeyMethod string `json:"node_master_key_method"`
}
*/
/*
func CreatedNodeDpki(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var nodePki NodeDpki
	_ = decoder.Decode(&nodePki)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successful"))
	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("error, such as duplicate node_id"))

}

func UpdateNodeName(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var nodePki NodeDpki
	_ = decoder.Decode(&nodePki)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successful"))
	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("error: unauthorized"))
}
*/
func RegisterCallback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var urlDpki UrlDpki
	if err := json.NewDecoder(r.Body).Decode(&urlDpki); err != nil {
		panic(err)
	}

	errs := url.Values{}

	if urlDpki.Url == "" {
		errs.Add("url", "The url field is required!")
	}

	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration Successful"))

	//TODO if return 400 //not have in swaggerhub
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Response"))

	//TODO if return 404 // not have in swaggerhub
	//w.WriteHeader(http.StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))

}

func RegisterCallbackMasterKey(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var urlDpki UrlDpki
	if err := json.NewDecoder(r.Body).Decode(&urlDpki); err != nil {
		panic(err)
	}

	errs := url.Values{}

	if urlDpki.Url == "" {
		errs.Add("url", "The url field is required!")
	}

	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//TODO Check input
	//TODO if return 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration Successful"))

	//TODO if return 400 //not have in swaggerhub
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Response"))

	//TODO if return 404 // not have in swaggerhub
	//w.WriteHeader(http.StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))

}
