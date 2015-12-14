package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// All external commands https://old.nagios.org/developerinfo/externalcommands/commandlist.php

// HandleAcknowledgeHostProblem ACKNOWLEDGE_HOST_PROBLEM
// POST: /acknowledge_host_problem/<host>
//       {sticky:bool, notify:bool, persistent:bool, author:string, comment:string}
func HandleAcknowledgeHostProblem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname   string
		Sticky     int
		Notify     int
		Persistent int
		Author     string
		Comment    string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, "Missing host", 400)
		return
	}

	if data.Sticky == 0 {
		data.Sticky = 2
	}

	if data.Notify == 0 {
		data.Notify = 1
	}

	if data.Persistent == 0 {
		data.Persistent = 1
	}

	command := fmt.Sprintf("%s;%s;%d;%d;%d;%s;%s", "ACKNOWLEDGE_HOST_PROBLEM", data.Hostname, data.Sticky, data.Notify, data.Persistent, data.Author, data.Comment)
	WriteCommandToFile(w, command)
}

// HandleAcknowledgeServiceProblem ACKNOWLEDGE_SVC_PROBLEM
// POST: /acknowledge_service_problem
//       {sticky:bool, notify:bool, persistent:bool, author:string, comment:string}
func HandleAcknowledgeServiceProblem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname           string
		ServiceDescription string
		Sticky             int
		Notify             int
		Persistent         int
		Author             string
		Comment            string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, "Missing hostname in data.", 400)
		return
	}

	if data.ServiceDescription == "" {
		http.Error(w, "Missing servicedescription in data.", 400)
		return
	}

	if data.Sticky == 0 {
		data.Sticky = 2
	}

	if data.Notify == 0 {
		data.Notify = 1
	}

	if data.Persistent == 0 {
		data.Persistent = 1
	}

	command := fmt.Sprintf("%s;%s;%s;%d;%d;%d;%s;%s", "ACKNOWLEDGE_SVC_PROBLEM", data.Hostname, data.ServiceDescription, data.Sticky, data.Notify, data.Persistent, data.Author, data.Comment)
	WriteCommandToFile(w, command)
}

// HandleAddHostComment ADD_HOST_COMMENT
// POST: /add_host_comment/<host>
//       {persistent:bool, author:string, comment:string}
func HandleAddHostComment(w http.ResponseWriter, r *http.Request) {
}

// HandleAddServiceComment ADD_SVC_COMMENT
// POST: /add_svc_comment/<host>/<service>
//       {persistent:bool, author:string, comment:string}
func HandleAddServiceComment(w http.ResponseWriter, r *http.Request) {
}

// HandleDeleteAllHostCommnet DEL_ALL_HOST_COMMENTS
func HandleDeleteAllHostCommnet(w http.ResponseWriter, r *http.Request) {
}

// HandleDeleteAllServiceComment DEL_ALL_SVC_COMMENTS
func HandleDeleteAllServiceComment(w http.ResponseWriter, r *http.Request) {
}

// HandleDeleteHostComment DEL_HOST_COMMENT
func HandleDeleteHostComment(w http.ResponseWriter, r *http.Request) {
}

// HandleDeleteServiceComment DEL_SVC_COMMENT
func HandleDeleteServiceComment(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableAllNotificationBeyondHost DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST
func HandleDisableAllNotificationBeyondHost(w http.ResponseWriter, r *http.Request) {
}

// HandleEnableAllNotificationBeyondHost ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST
func HandleEnableAllNotificationBeyondHost(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableHostgroupHostChecks DISABLE_HOSTGROUP_HOST_CHECKS
func HandleDisableHostgroupHostChecks(w http.ResponseWriter, r *http.Request) {
}

// HandleEnableHostgroupHostChecks ENABLE_HOSTGROUP_HOST_CHECKS
func HandleEnableHostgroupHostChecks(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableHostgroupHostNotification DISABLE_HOSTGROUP_HOST_NOTIFICATIONS
func HandleDisableHostgroupHostNotification(w http.ResponseWriter, r *http.Request) {
}

// HandleEnableHostgroupHostNotification ENABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
func HandleEnableHostgroupHostNotification(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableHostgroupServiceChecks DISABLE_HOSTGROUP_SVC_CHECKS
func HandleDisableHostgroupServiceChecks(w http.ResponseWriter, r *http.Request) {
}

// HandleEnableHostgroupServiceChecks ENABLE_HOSTGROUP_SVC_CHECKS
func HandleEnableHostgroupServiceChecks(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableHostgroupServiceNotifications DISABLE_HOSTGROUP_SVC_NOTIFICATIONS
func HandleDisableHostgroupServiceNotifications(w http.ResponseWriter, r *http.Request) {
}

// HandleEnableHostgroupServiceNotifications ENABLE_HOSTGROUP_SVC_NOTIFICATIONS
func HandleEnableHostgroupServiceNotifications(w http.ResponseWriter, r *http.Request) {
}

// HandleDisableHostandChildNotifications DISABLE_HOST_AND_CHILD_NOTIFICATIONS
func HandleDisableHostandChildNotifications(w http.ResponseWriter, r *http.Request) {
}

// ENABLE_HOST_AND_CHILD_NOTIFICATIONS
func HandleEnableHostandChildNotifications(w http.ResponseWriter, r *http.Request) {
}

// DISABLE_HOST_CHECK
// POST: /disable_host_check
func HandleDisableHostCheck(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOST_CHECK", host.Hostname)
	WriteCommandToFile(w, command)
}

// ENABLE_HOST_CHECK
// POST: /enable_host_check
func HandleEnableHostCheck(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOST_CHECK", host.Hostname)
	WriteCommandToFile(w, command)
}

// DISABLE_HOST_NOTIFICATIONS
// POST: /disable_host_notifications
func HandleDisableHostNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOST_NOTIFICATIONS", host.Hostname)
	WriteCommandToFile(w, command)
}

// ENABLE_HOST_NOTIFICATIONS
// POST: /enable_host_notifications
func HandleEnableHostNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOST_NOTIFICATIONS", host.Hostname)
	WriteCommandToFile(w, command)
}

// DISABLE_NOTIFICATIONS
// POST: /disable_notifications
func HandleDisableNotifications(w http.ResponseWriter, r *http.Request) {
	command := "DISABLE_NOTIFICATIONS"
	WriteCommandToFile(w, command)
}

// ENABLE_NOTIFICATIONS
// POST: /enable_notifications
func HandleEnableNotifications(w http.ResponseWriter, r *http.Request) {
	command := "ENABLE_NOTIFICATIONS"
	WriteCommandToFile(w, command)
}

// SCHEDULE_FORCED_HOST_CHECK
// SCHEDULE_FORCED_HOST_CHECK;<host_name>;<check_time>
func HandleScheduleForcedHostCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_FORCED_HOST_SVC_CHECKS
// SCHEDULE_FORCED_HOST_SVC_CHECKS;<host_name>;<check_time>
func HandleScheduleForcedHostServiceChecks(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_FORCED_SVC_CHECK
// SCHEDULE_FORCED_SVC_CHECK;<host_name>;<service_description>;<check_time>
func HandleScheduleForcedServiceCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_HOST_CHECK
// SCHEDULE_HOST_CHECK;<host_name>;<check_time>
func HandleScheduleHostCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_HOST_DOWNTIME
// SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
func HandleScheduleHostDowntime(w http.ResponseWriter, r *http.Request) {
}

func WriteCommandToFile(w http.ResponseWriter, command string) {
	if err := WriteCommand(command); err != nil {
		http.Error(w, "Could not execute command", 500)
		return
	}
}
