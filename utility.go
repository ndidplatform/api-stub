package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RetrieveAllIdpId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	minIal := r.URL.Query().Get("min_ial")
	minAal := r.URL.Query().Get("min_aal")
	a, _ := strconv.Atoi(minIal)
	b, _ := strconv.Atoi(minAal)
	if a > 10 && b > 10 { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: No IDP found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"idp_name": "string","idp_id": "string","min_ial": ` + minIal + `,"min_aal": ` + minAal + `}]`))
	}
	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: MIN IAL, AAL value"))

}

func RetrieveAllIdpNodeId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	namespace := params["namespace"]
	identifier := params["identifier"]
	minIal := r.URL.Query().Get("min_ial")
	minAal := r.URL.Query().Get("min_aal")
	a, _ := strconv.Atoi(minIal)
	b, _ := strconv.Atoi(minAal)

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
	} else if a > 10 && b > 10 { //return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: No IDP found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"idp_name": "string","idp_id": "string","min_ial": ` + minIal + `,"min_aal": ` + minAal + `}]`))
	}
	//TODO if return 400
	//w.WriteHeader(StatusBadRequest)
	//w.Write([]byte("Error: Invalid Identity or MIN IAL, AAL value"))

}

func RetrieveListGivenServiceId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//params := mux.Vars(r)
	//serviceId := params["service_id"]

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`[{"as_name": "string","as_id": "string","min_ial": 0,"min_aal": 0}]`))

}
