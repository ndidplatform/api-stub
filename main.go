package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// Identity Management API
	// https://app.swaggerhub.com/apis/NDID/identity/1.0
	router.HandleFunc("/identity", CreateIdentity).Methods("POST")
	router.HandleFunc("/identity/{namespace}/{identifier}", FetchIdentity).Methods("GET")
	router.HandleFunc("/identity/{namespace}/{identifier}", NewIdentifiers).Methods("POST")
	router.HandleFunc("/identity/{namespace}/{identifier}/endorsement", FetchEndorsement).Methods("GET")
	router.HandleFunc("/identity/{namespace}/{identifier}/endorsement", SubmitEndorsement).Methods("POST")
	router.HandleFunc("/identity/{namespace}/{identifier}/accessors", AddAccessorMethod).Methods("POST")
	router.HandleFunc("/identity/{namespace}/{identifier}/requests/history", FetchRequestHistory).Methods("GET")

	// IDP
	// https://app.swaggerhub.com/apis/NDID/identity_provider/1.0
	router.HandleFunc("/idp/callback", RetrieveCallbackUrl).Methods("GET")
	router.HandleFunc("/idp/callback", SetUrlCallback).Methods("POST")
	router.HandleFunc("/idp/response", IdpResponse).Methods("POST")

	// RP
	// https://app.swaggerhub.com/apis/NDID/relying_party_api/1.0
	router.HandleFunc("/rp/requests/{namespace}/{identifier}", RequestToId).Methods("POST")
	router.HandleFunc("/rp/requests/{namespace}/{identifier}/{timeout}", RequestToId).Methods("POST")
	router.HandleFunc("/rp/requests/{request_id}", FetchRequestStatus).Methods("GET")
	router.HandleFunc("/rp/requests/reference/{reference_number}", GetRefNo).Methods("GET")
	router.HandleFunc("/rp/requests/data/{request_id}", GetData).Methods("GET")

	// AS
	// https://app.swaggerhub.com/apis/NDID/authoritative_source_api/1.0
	router.HandleFunc("/as/service/{service_id}", AddUpdateAsService).Methods("POST")
	router.HandleFunc("/as/service/{service_id}", GetAsServiceInfo).Methods("GET")

	// Utility API
	// https://app.swaggerhub.com/apis/NDID/utility/1.0
	router.HandleFunc("/utility/idp", RetrieveAllIdpId).Methods("GET")
	router.HandleFunc("/utility/idp/{namespace}/{identifier}", RetrieveAllIdpNodeId).Methods("GET")
	router.HandleFunc("/utility/as/{service_id}", RetrieveListGivenServiceId).Methods("GET")

	// DPKI API
	// https://app.swaggerhub.com/apis/NDID/dpki/1.0
	// Only NDID can use the first two API
	//router.HandleFunc("/dpki/node/create", CreatedNodeDpki).Methods("POST")
	//router.HandleFunc("/dpki/node/update", UpdateNodeName).Methods("POST")
	router.HandleFunc("/dpki/node/register_callback", RegisterCallback).Methods("POST")
	router.HandleFunc("/dpki/node/register_callback_master", RegisterCallbackMasterKey).Methods("POST")

	// Callback API (not provided by NDID, rather, expected by NDID, implemented by IDP, RP and AS respectively
	// IDP Callback
	// https://app.swaggerhub.com/apis/NDID/idp_callback/1.0
	// RP Callback
	// https://app.swaggerhub.com/apis/NDID/rp_callback/1.0
	// AS Callback
	// https://app.swaggerhub.com/apis/NDID/as_callback/1.0
	// DPKI callback
	// https://app.swaggerhub.com/apis/NDID/dpki_callback/1.0

	log.Fatal(http.ListenAndServe(":8000", router))

}
