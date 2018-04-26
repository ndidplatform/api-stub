package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
)

type identity struct {
	Namespace    string `json:"namespace"`
	Identifier   string `json:"identifier"`
	Secret       string `json:"secret"`
	AccessorType string `json:"accessor_type"`
	AccessorKey  string `json:"accessor_key"`
	AccessorId   string `json:"accessor_id"`
}

type identifiers struct {
	Identifiers string `json:"identifiers"` //array of string
}

type endorsement struct { //same as identity
	Namespace    string `json:"namespace"`
	Identifier   string `json:"identifier"`
	Secret       string `json:"secret"`
	AccessorType string `json:"accessor_type"`
	AccessorKey  string `json:"accessor_key"`
	AccessorId   string `json:"accessor_id"`
}

type accessor struct {
	AccessorType string `json:"accessor_type"`
	AccessorKey  string `json:"accessor_key"`
	AccessorId   string `json:"accessor_id"`
}

func (p identity) toString() string {
	return toJson(p)
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
func getIdentity() []identity {
	raw, err := ioutil.ReadFile("./identity.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []identity
	json.Unmarshal(raw, &c)
	return c
}

func CreateIdentity(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var id identity
	_ = decoder.Decode(&id)
	defer r.Body.Close()
	//TODO Check input
	//TODO if return 201
	if true {
		ids := getIdentity()
		_ = json.Unmarshal([]byte(id.toString()), &ids)
		_ = json.Unmarshal([]byte(id.toString()), &id)
		//fmt.Println(id.Namespace)
		ids = append(ids, identity{Namespace: id.Namespace, Identifier: id.Identifier, Secret: id.Secret, AccessorType: id.AccessorType, AccessorKey: id.AccessorKey, AccessorId: id.AccessorId})
		/*
			for _, p := range ids {
				fmt.Println(p.toString())
			}
		*/
		Json, _ := json.Marshal(ids)
		_ = ioutil.WriteFile("identity.json", Json, 0644)
		// now result has your targeted JSON structure
	}
	fmt.Println(id.toString())
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Identity Created"))
	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid accessor method or accessor value"))
	//TODO if return 403
	//w.WriteHeader(http.StatusForbidden)
	//w.Write([]byte("Error: Identity already exist"))

	// define slice of Identification

}

func FetchIdentity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	//TODO Check input
	//Test
	var id identity
	ids := getIdentity()
	exist := 0
	for _, p := range ids {
		_ = json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = 1
		}
	}
	if exist == 1 {
		//return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ID exist"))
		//fmt.Println(p.toString())
	} else {
		//return 404
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID not found"))
	}
}

func NewIdentifiers(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("Endorsement Successful"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid endorsement type"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))
}

func FetchEndorsement(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new endorsement
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
	w.Write([]byte("{\"endorsement_type\": \"string\", \"endrosement_value\": \"string\"}"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("No endorsement Exist"))
}

func SubmitEndorsement(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new endorsement
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
	w.Write([]byte("Endorsement Successful"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid endorsement type"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))
}
func AddAccessorMethod(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new accessor
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
	w.Write([]byte("Endorsement Successful"))

	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid accessor type"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("Error: Identity does not exist"))
}

func FetchRequestHistory(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var new endorsement
	_ = decoder.Decode(&new)
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	count := params["count"]

	//TODO Check input
	println(namespace)
	println(identifier)
	println(count)

	//TODO if return 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"request_id\": \"string\", \"request_body\": \"string\"}]"))

	//TODO if return 404
	//w.WriteHeader(StatusNotFound)
	//w.Write([]byte("No request Exist"))
}
