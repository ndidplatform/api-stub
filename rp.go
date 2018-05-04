package main

import (
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
)

type requestForConsent struct {
	ReferenceNumber string          `json:"namespace"`
	IdpList         []string        `json:"idp_list"`
	CallbackUrl     string          `json:"callback_url"`
	AsServiceList   []asServiceList `json:"as_service_list"` // array of object i.e. "as_service_list": [{"service_id": "string", "as_id": ["string"], "count": 0, "params": "string"}]
	RequestMessage  string          `json:"request_message"`
	MinIal          string          `json:"min_ial"`
	MinAal          string          `json:"min_aal"`
	MinIdp          int             `json:"min_idp"`
	RequestTimeout  int             `json:"request_timeout"`
}
type asServiceList struct {
	ServiceId string   `json:"service_id"`
	AsId      []string `json:"as_id"`
	Count     int      `json:"count"`
	Params    string   `json:"params"`
}

func RequestToId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var new requestForConsent
	if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
		panic(err)
	}
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	timeout := r.URL.Query().Get("timeout")

	errs := url.Values{}

	if new.MinAal == "" {
		errs.Add("min_aal", "The min_aal field is required!")
	}
	if new.MinIal == "" {
		errs.Add("min_ial", "The min_ial field is required!")
	}
	if new.MinIdp < 1 {
		errs.Add("min_idp", "The min_idp field cannot be empty or 0")
	}

	if new.RequestMessage == "" {
		errs.Add("request_message", "The request_message field is required!")
	}
	if new.ReferenceNumber != "" {
		if _, err := strconv.Atoi(timeout); err != nil {
			errs.Add("refference_number", "The refference_number field must be the number!")
		}
	}

	if timeout != "" {
		if _, err := strconv.Atoi(timeout); err != nil {
			errs.Add("timeout", "The timeout field must be the number!")
		}
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
	} else if new.RequestTimeout > 0 && new.CallbackUrl != "" { //return 201
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request Accepted – Async processing, please check back or wait for response at Callback URL"))
	} else if timeout == "" { //return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request Successful"))
	} else {
		// return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Invalid Request"))
	}

}

func FetchRequestStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	requestId := params["request_id"]

	if requestId != "xxx" { // return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"request_type": "consent", "request_message": "string", "min_ial": 1, "min_aal": 0, "min_idp": 0, "service_id_list": ["string"], "timeout": 0, "status": "open"}`))
	} else {
		// return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No request Exist"))
	}
}

func GetRefNo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	refno := params["reference_number"]

	if refno != "xxx" { // return 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ex. reference no"))
	} else {
		// return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No reference number in the system"))
	}
}

func GetData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	reqId := params["request_id"]

	if reqId == "xxx" { // return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data found"))
	} else if reqId == "yyy" {
		// return 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request does not exist (Bad/incorrect" + reqId + ")"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"source_node_id": "string", "service_id": "string", "source_signature\": \"string\", \"data\": \"string\"}]`))
	}
}
