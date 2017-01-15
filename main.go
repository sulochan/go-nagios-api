package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sulochan/go-nagios-api/api"
	"github.com/sulochan/go-nagios-api/auth"
	"github.com/sulochan/go-nagios-api/config"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	conf := config.GetConfig()
	chain := alice.New()
	router := mux.NewRouter()
	http.Handle("/", router)
	router.Handle("/contacts", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetContacts)).Methods("GET")

	router.Handle("/hosts", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetConfiguredHosts)).Methods("GET")
	router.Handle("/host/{hostname:[a-z,A-Z,0-9, _.-]+}", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetHost)).Methods("GET")
	router.Handle("/host/{hostname:[a-z,A-Z,0-9, _.-]+}/services", chain.Append(auth.AuthHandler).ThenFunc(api.HandleGetServicesForHost)).Methods("GET")
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
	router.Handle("/add_host_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleAddHostComment)).Methods("POST")
	router.Handle("/add_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleAddServiceComment)).Methods("POST")
	router.Handle("/del_all_host_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDeleteAllHostCommnet)).Methods("POST")
	router.Handle("/del_all_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDeleteAllServiceComment)).Methods("POST")
	router.Handle("/del_host_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDeleteHostComment)).Methods("POST")
	router.Handle("/del_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDeleteServiceComment)).Methods("POST")
	router.Handle("/disable_all_notification_beyond_host", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableAllNotificationBeyondHost)).Methods("POST")
	router.Handle("/enable_all_notification_beyond_host", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableAllNotificationBeyondHost)).Methods("POST")
	router.Handle("/disable_hostgroup_host_checks", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostgroupHostChecks)).Methods("POST")
	router.Handle("/enable_hostgroup_host_checks", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostgroupHostChecks)).Methods("POST")
	router.Handle("/disable_hostgroup_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostgroupHostNotification)).Methods("POST")
	router.Handle("/enable_hostgroup_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostgroupHostNotification)).Methods("POST")
	router.Handle("/disable_hostgroup_svc_checks", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostgroupServiceChecks)).Methods("POST")
	router.Handle("/enable_hostgroup_svc_checks", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostgroupServiceChecks)).Methods("POST")
	router.Handle("/disable_hostgroup_svc_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostgroupServiceNotifications)).Methods("POST")
	router.Handle("/enable_hostgroup_svc_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostgroupServiceNotifications)).Methods("POST")
	router.Handle("/disable_host_and_child_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleDisableHostandChildNotifications)).Methods("POST")
	router.Handle("/enable_host_and_child_notifications", chain.Append(auth.AuthHandler).ThenFunc(api.HandleEnableHostandChildNotifications)).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil); err != nil {
		log.Printf("http.ListendAndServer() failed with %s\n", err)
	}
	log.Printf("Exited\n")
}
