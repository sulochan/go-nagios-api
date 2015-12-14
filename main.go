package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sulochan/go-nagios-api/api"
	"github.com/sulochan/go-nagios-api/auth"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	chain := alice.New()
	router := mux.NewRouter()
	http.Handle("/", router)
	router.Handle("/contacts", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetContacts)).Methods("GET")

	router.Handle("/hosts", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetConfiguredHosts)).Methods("GET")
	router.Handle("/host/{hostname:[a-z,A-Z,0-9, _.-]+}", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetHost)).Methods("GET")
	router.Handle("/hoststatus", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetAllHostStatus)).Methods("GET")
	router.Handle("/hoststatus/{hostname:[a-z,A-Z,0-9,_.-]+}", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetHostStatusForHost)).Methods("GET")
	router.Handle("/hostgroups", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetHostGroups)).Methods("GET")

	router.Handle("/services", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetConfiguredServices)).Methods("GET")
	router.Handle("/servicestatus", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetServiceStatus)).Methods("GET")
	router.Handle("/servicestatus/{service:[a-z,A-Z,0-9,_.-]+}", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetServiceStatusForService)).Methods("GET")

	// Nagios External Command Handlers
	router.Handle("/disable_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableNotifications)).Methods("POST")
	router.Handle("/enable_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableNotifications)).Methods("POST")
	router.Handle("/disable_host_check", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostCheck)).Methods("POST")
	router.Handle("/enable_host_check", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostCheck)).Methods("POST")
	router.Handle("/disable_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostNotifications)).Methods("POST")
	router.Handle("/enable_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostNotifications)).Methods("POST")
	router.Handle("/acknowledge_host_problem", chain.Append(auth.AuthHandler).ThenFunc(api.HandleAcknowledgeHostProblem)).Methods("POST")
	router.Handle("/acknowledge_service_problem", chain.Append(auth.AuthHandler).ThenFunc(api.HandleAcknowledgeServiceProblem)).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil); err != nil {
		log.Printf("http.ListendAndServer() failed with %s\n", err)
	}
	log.Printf("Exited\n")
}
