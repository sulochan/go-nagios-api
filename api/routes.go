package api

import (
	"github.com/sulochan/go-nagios-api/auth"

	"github.com/justinas/alice"
)

func (s *Api) buildRoutes() {
	chain := alice.New()

	s.router.Handle("/contacts", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetContacts)).Methods("GET")

	s.router.Handle("/hosts", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetConfiguredHosts)).Methods("GET")
	s.router.Handle("/host/{hostname:[a-z,A-Z,0-9, _.-]+}", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetHost)).Methods("GET")
	s.router.Handle("/host/{hostname:[a-z,A-Z,0-9, _.-]+}/services", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetServicesForHost)).Methods("GET")
	s.router.Handle("/hoststatus", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetAllHostStatus)).Methods("GET")
	s.router.Handle("/hoststatus/{hostname:[a-z,A-Z,0-9,_.-]+}", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetHostStatusForHost)).Methods("GET")
	s.router.Handle("/hostgroups", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetHostGroups)).Methods("GET")

	s.router.Handle("/services", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetConfiguredServices)).Methods("GET")
	s.router.Handle("/servicestatus", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetServiceStatus)).Methods("GET")
	s.router.Handle("/servicestatus/{service:[a-z,A-Z,0-9,_.-]+}", chain.Append(auth.AuthHandler).ThenFunc(s.HandleGetServiceStatusForService)).Methods("GET")

	// Nagios External Command Handlers
	s.router.Handle("/disable_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableNotifications)).Methods("POST")
	s.router.Handle("/enable_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableNotifications)).Methods("POST")
	s.router.Handle("/disable_host_check", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostCheck)).Methods("POST")
	s.router.Handle("/enable_host_check", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostCheck)).Methods("POST")
	s.router.Handle("/disable_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostNotifications)).Methods("POST")
	s.router.Handle("/enable_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostNotifications)).Methods("POST")
	s.router.Handle("/acknowledge_host_problem", chain.Append(auth.AuthHandler).ThenFunc(s.HandleAcknowledgeHostProblem)).Methods("POST")
	s.router.Handle("/acknowledge_service_problem", chain.Append(auth.AuthHandler).ThenFunc(s.HandleAcknowledgeServiceProblem)).Methods("POST")
	s.router.Handle("/add_host_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleAddHostComment)).Methods("POST")
	s.router.Handle("/add_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleAddServiceComment)).Methods("POST")
	s.router.Handle("/del_all_host_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDeleteAllHostCommnet)).Methods("POST")
	s.router.Handle("/del_all_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDeleteAllServiceComment)).Methods("POST")
	s.router.Handle("/del_host_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDeleteHostComment)).Methods("POST")
	s.router.Handle("/del_svc_comment", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDeleteServiceComment)).Methods("POST")
	s.router.Handle("/disable_all_notification_beyond_host", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableAllNotificationBeyondHost)).Methods("POST")
	s.router.Handle("/enable_all_notification_beyond_host", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableAllNotificationBeyondHost)).Methods("POST")
	s.router.Handle("/disable_hostgroup_host_checks", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostgroupHostChecks)).Methods("POST")
	s.router.Handle("/enable_hostgroup_host_checks", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostgroupHostChecks)).Methods("POST")
	s.router.Handle("/disable_hostgroup_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostgroupHostNotification)).Methods("POST")
	s.router.Handle("/enable_hostgroup_host_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostgroupHostNotification)).Methods("POST")
	s.router.Handle("/disable_hostgroup_svc_checks", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostgroupServiceChecks)).Methods("POST")
	s.router.Handle("/enable_hostgroup_svc_checks", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostgroupServiceChecks)).Methods("POST")
	s.router.Handle("/disable_hostgroup_svc_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostgroupServiceNotifications)).Methods("POST")
	s.router.Handle("/enable_hostgroup_svc_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostgroupServiceNotifications)).Methods("POST")
	s.router.Handle("/disable_host_and_child_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleDisableHostandChildNotifications)).Methods("POST")
	s.router.Handle("/enable_host_and_child_notifications", chain.Append(auth.AuthHandler).ThenFunc(s.HandleEnableHostandChildNotifications)).Methods("POST")
	s.router.Handle("/schedule_host_downtime", chain.Append(auth.AuthHandler).ThenFunc(s.HandleScheduleHostDowntime)).Methods("POST")
}
