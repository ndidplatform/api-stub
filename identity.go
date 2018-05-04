package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

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
	Identifiers []string `json:"identifiers"` //array of string
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
type history struct {
	RequestId   string `json:"request_id"`
	RequestBody string `json:"request_body"`
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

func (id *identity) validateIdentiy() url.Values {
	errs := url.Values{}

	// check if the field empty
	if id.Namespace == "" {
		errs.Add("namespace", "The namespace field is required!")
	}
	if id.Identifier == "" {
		errs.Add("identifier", "The identifier field is required!")
	}
	if id.Secret == "" {
		errs.Add("secret", "The secret field is required!")
	}
	if id.AccessorType == "" {
		errs.Add("accessor_type", "The accessor_type field is required!")
	}
	if id.AccessorKey == "" {
		errs.Add("accessor_key", "The accessor_key field is required!")
	}
	if id.AccessorId == "" {
		errs.Add("accessor_id", "The accessor_id field is required!")
	}

	// check the title field is between 3 to 120 chars
	//if len(id.Title) < 3 || len(id.Title) > 120 {
	//	errs.Add("title", "The title field must be between 3-120 chars!")
	//}

	//if len(id.Body) < 50 || len(id.Body) > 500 {
	//	errs.Add("body", "The title field must be between 50-500 chars!")
	//}

	// Note: for checking the complex or regex related validation I use govalidator here

	return errs
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
	////////////////

	defer r.Body.Close()
	iden := &identity{}

	if err := json.NewDecoder(r.Body).Decode(iden); err != nil {
		panic(err)
	}
	//Validate
	if validErrs := iden.validateIdentiy(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	ids := getIdentity()

	idExist := false
	for _, p := range ids {
		if p.Namespace == iden.Namespace && p.Identifier == iden.Identifier {
			idExist = true
		}
	}

	if idExist { //403
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error: Identity already exist"))
	} else if iden.AccessorKey == "" || iden.AccessorType == "" { //400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid accessor method or accessor value"))
	} else {
		ids = append(ids, identity{Namespace: iden.Namespace, Identifier: iden.Identifier, Secret: iden.Secret, AccessorType: iden.AccessorType, AccessorKey: iden.AccessorKey, AccessorId: iden.AccessorId})
		Json, _ := json.Marshal(ids)
		_ = ioutil.WriteFile("identity.json", Json, 0644)
		// now result has your targeted JSON structure
		fmt.Println(iden.toString())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Identity Created"))
	}
}

func FetchIdentity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]

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
	defer r.Body.Close()
	var new identifiers
	if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
		panic(err)
	}
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	//TODO Check input
	errs := url.Values{}
	// check if the field empty
	if len(new.Identifiers) == 0 {
		errs.Add("identifiers", "The identifiers field cannot empty!")
	}
	//Validate
	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//////////////////
	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		_ = json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Identity does not exist"))
	} else if new.Identifiers == nil { //return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid endorsement type"))
	} else { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Endorsement Successful"))
	}
}

func FetchEndorsement(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]

	//Check input
	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Identity does not exist"))
	} else if namespace == "cid" && identifier == "2345678901234" { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No endorsement Exist"))
	} else { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"endorsement_type": "string", "endrosement_value": "string"}`))
	}
}

func SubmitEndorsement(w http.ResponseWriter, r *http.Request) {
	////////////////////
	defer r.Body.Close()
	var new endorsement
	if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
		panic(err)
	}

	//TODO Check input
	errs := url.Values{}
	// check if the field empty
	if new.Secret == "" {
		errs.Add("secret", "The secret field is required!")
	}
	if new.AccessorType == "" {
		errs.Add("accessor_type", "The accessor_type field is required!")
	}
	if new.AccessorKey == "" {
		errs.Add("accessor_key", "The accessor_key field is required!")
	}
	if new.AccessorId == "" {
		errs.Add("accessor_id", "The accessor_id field is required!")
	}
	//Validate
	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	////////////////////

	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	//////////////////////////
	//Check input
	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Identity does not exist"))
	} else if new.AccessorType == "" { //return 404
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid endorsement type"))
	} else { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Endorsement Successful"))
	}

}
func AddAccessorMethod(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var new accessor
	if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]

	//TODO Check input
	errs := url.Values{}
	// check if the field empty
	if new.AccessorType == "" {
		errs.Add("accessor_type", "The accessor_type field is required!")
	}
	if new.AccessorKey == "" {
		errs.Add("accessor_key", "The accessor_key field is required!")
	}
	if new.AccessorId == "" {
		errs.Add("accessor_id", "The accessor_id field is required!")
	}
	//Validate
	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//Check input
	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: Identity does not exist"))
	} else if new.AccessorType == "" { //return 404
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid endorsement type"))
	} else { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Endorsement Successful"))
	}

}

func FetchRequestHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	count := r.URL.Query().Get("count")
	//////////////////////////////////////
	errs := url.Values{}
	// check if the field empty

	if count == "" {
		errs.Add("count", "The count field is required!")
	}
	if _, err := strconv.Atoi(count); err != nil {
		errs.Add("count", "The count field must be the number!")
	}
	//Validate
	if len(errs) > 0 {
		err := map[string]interface{}{"validationError": errs}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	//Check input
	var id identity
	ids := getIdentity()
	exist := false
	for _, p := range ids {
		json.Unmarshal([]byte(p.toString()), &id)
		if id.Namespace == namespace && id.Identifier == identifier {
			exist = true
		}
	}
	if exist == false { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No request Exist"))
	} else { //return 200
		raw := []byte(`{"request_id": "string", "request_body": "string"}`)
		var c []history
		json.Unmarshal(raw, &c)
		k, _ := strconv.Atoi(count)
		for i := 0; i < k; i++ {
			c = append(c, history{RequestId: "string", RequestBody: "string"})

		}
		Json, _ := json.Marshal(c)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Json))
	}
}
