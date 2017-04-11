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
func (a *Api) HandleAcknowledgeHostProblem(w http.ResponseWriter, r *http.Request) {
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

	if data.Author == "" {
		http.Error(w, "Error: Author filed is required", 400)
		return
	}

	command := fmt.Sprintf("%s;%s;%d;%d;%d;%s;%s", "ACKNOWLEDGE_HOST_PROBLEM", data.Hostname, data.Sticky, data.Notify, data.Persistent, data.Author, data.Comment)
	a.WriteCommandToFile(w, command)
}

// HandleAcknowledgeServiceProblem ACKNOWLEDGE_SVC_PROBLEM
// POST: /acknowledge_service_problem
//       {sticky:bool, notify:bool, persistent:bool, author:string, comment:string}
func (a *Api) HandleAcknowledgeServiceProblem(w http.ResponseWriter, r *http.Request) {
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
	a.WriteCommandToFile(w, command)
}

// HandleAddHostComment ADD_HOST_COMMENT
// POST: /add_host_comment/<host>
//       {persistent:bool, author:string, comment:string}
func (a *Api) HandleAddHostComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname   string
		Persistent int
		Author     string
		Comment    string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		return
	}

	if data.Hostname == "" {
		http.Error(w, "Missing host", 400)
		return
	}

	if data.Persistent == 0 {
		data.Persistent = 1
	}

	if data.Author == "" {
		http.Error(w, fmt.Sprintf("Error: Author field is required"), 400)
		return
	}

	if data.Comment == "" {
		http.Error(w, fmt.Sprintf("Error: Comment can not be empty"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s;%d;%s;%s", "ADD_HOST_COMMENT", data.Hostname, data.Persistent, data.Author, data.Comment)
	a.WriteCommandToFile(w, command)

}

// HandleAddServiceComment ADD_SVC_COMMENT
// POST: /add_svc_comment/<host>/<service>
//       {persistent:bool, author:string, comment:string}
func (a *Api) HandleAddServiceComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname   string
		Service    string
		Persistent int
		Author     string
		Comment    string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		return
	}

	if data.Hostname == "" {
		http.Error(w, "Missing hostname", 400)
		return
	}

	if data.Persistent == 0 {
		data.Persistent = 1
	}

	if data.Author == "" {
		http.Error(w, fmt.Sprintf("Error: Author field is required"), 400)
		return
	}

	if data.Service == "" {
		http.Error(w, fmt.Sprintf("Error: ServiceDesc can not be empty"), 400)
		return
	}

	if data.Comment == "" {
		http.Error(w, fmt.Sprintf("Error: Comment can not be empty"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s;%s;%d;%s;%s", "ADD_SVC_COMMENT", data.Hostname, data.Service, data.Persistent, data.Author, data.Comment)
	a.WriteCommandToFile(w, command)
}

// HandleDeleteAllHostCommnet DEL_ALL_HOST_COMMENTS
func (a *Api) HandleDeleteAllHostCommnet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DEL_ALL_HOST_COMMENTS", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// HandleDeleteAllServiceComment DEL_ALL_SVC_COMMENTS
func (a *Api) HandleDeleteAllServiceComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DEL_ALL_SVC_COMMENTS", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// HandleDeleteHostComment DEL_HOST_COMMENT
func (a *Api) HandleDeleteHostComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		CommentID string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.CommentID == "" {
		http.Error(w, fmt.Sprintf("Error: CommentID field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DEL_HOST_COMMENT", data.CommentID)
	a.WriteCommandToFile(w, command)
}

// HandleDeleteServiceComment DEL_SVC_COMMENT
func (a *Api) HandleDeleteServiceComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		CommentID string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.CommentID == "" {
		http.Error(w, fmt.Sprintf("Error: CommentID field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DEL_SVC_COMMENT", data.CommentID)
	a.WriteCommandToFile(w, command)
}

// HandleDisableAllNotificationBeyondHost DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST
func (a *Api) HandleDisableAllNotificationBeyondHost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// HandleEnableAllNotificationBeyondHost ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST
func (a *Api) HandleEnableAllNotificationBeyondHost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// HandleDisableHostgroupHostChecks DISABLE_HOSTGROUP_HOST_CHECKS
func (a *Api) HandleDisableHostgroupHostChecks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOSTGROUP_HOST_CHECKS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleEnableHostgroupHostChecks ENABLE_HOSTGROUP_HOST_CHECKS
func (a *Api) HandleEnableHostgroupHostChecks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOSTGROUP_HOST_CHECKS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleDisableHostgroupHostNotification DISABLE_HOSTGROUP_HOST_NOTIFICATIONS
func (a *Api) HandleDisableHostgroupHostNotification(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOSTGROUP_HOST_NOTIFICATIONS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleEnableHostgroupHostNotification ENABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
func (a *Api) HandleEnableHostgroupHostNotification(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOSTGROUP_HOST_NOTIFICATIONS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleDisableHostgroupServiceChecks DISABLE_HOSTGROUP_SVC_CHECKS
func (a *Api) HandleDisableHostgroupServiceChecks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOSTGROUP_SVC_CHECKS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleEnableHostgroupServiceChecks ENABLE_HOSTGROUP_SVC_CHECKS
func (a *Api) HandleEnableHostgroupServiceChecks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOSTGROUP_SVC_CHECKS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleDisableHostgroupServiceNotifications DISABLE_HOSTGROUP_SVC_NOTIFICATIONS
func (a *Api) HandleDisableHostgroupServiceNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOSTGROUP_SVC_NOTIFICATIONS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleEnableHostgroupServiceNotifications ENABLE_HOSTGROUP_SVC_NOTIFICATIONS
func (a *Api) HandleEnableHostgroupServiceNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostgroup string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostgroup == "" {
		http.Error(w, fmt.Sprintf("Error: Hostgroup name field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOSTGROUP_SVC_NOTIFICATIONS", data.Hostgroup)
	a.WriteCommandToFile(w, command)
}

// HandleDisableHostandChildNotifications DISABLE_HOST_AND_CHILD_NOTIFICATIONS
func (a *Api) HandleDisableHostandChildNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOST_AND_CHILD_NOTIFICATIONS", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// ENABLE_HOST_AND_CHILD_NOTIFICATIONS
func (a *Api) HandleEnableHostandChildNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname string
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 400)
		return
	}

	if data.Hostname == "" {
		http.Error(w, fmt.Sprintf("Error: Hostname field is required"), 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOST_AND_CHILD_NOTIFICATIONS", data.Hostname)
	a.WriteCommandToFile(w, command)
}

// DISABLE_HOST_CHECK
// POST: /disable_host_check
func (a *Api) HandleDisableHostCheck(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOST_CHECK", host.Hostname)
	a.WriteCommandToFile(w, command)
}

// ENABLE_HOST_CHECK
// POST: /enable_host_check
func (a *Api) HandleEnableHostCheck(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOST_CHECK", host.Hostname)
	a.WriteCommandToFile(w, command)
}

// DISABLE_HOST_NOTIFICATIONS
// POST: /disable_host_notifications
func (a *Api) HandleDisableHostNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "DISABLE_HOST_NOTIFICATIONS", host.Hostname)
	a.WriteCommandToFile(w, command)
}

// ENABLE_HOST_NOTIFICATIONS
// POST: /enable_host_notifications
func (a *Api) HandleEnableHostNotifications(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var host struct{ Hostname string }
	err := decoder.Decode(&host)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	command := fmt.Sprintf("%s;%s", "ENABLE_HOST_NOTIFICATIONS", host.Hostname)
	a.WriteCommandToFile(w, command)
}

// DISABLE_NOTIFICATIONS
// POST: /disable_notifications
func (a *Api) HandleDisableNotifications(w http.ResponseWriter, r *http.Request) {
	command := "DISABLE_NOTIFICATIONS"
	a.WriteCommandToFile(w, command)
}

// ENABLE_NOTIFICATIONS
// POST: /enable_notifications
func (a *Api) HandleEnableNotifications(w http.ResponseWriter, r *http.Request) {
	command := "ENABLE_NOTIFICATIONS"
	a.WriteCommandToFile(w, command)
}

// SCHEDULE_FORCED_HOST_CHECK
// SCHEDULE_FORCED_HOST_CHECK;<host_name>;<check_time>
func (a *Api) HandleScheduleForcedHostCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_FORCED_HOST_SVC_CHECKS
// SCHEDULE_FORCED_HOST_SVC_CHECKS;<host_name>;<check_time>
func (a *Api) HandleScheduleForcedHostServiceChecks(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_FORCED_SVC_CHECK
// SCHEDULE_FORCED_SVC_CHECK;<host_name>;<service_description>;<check_time>
func (a *Api) HandleScheduleForcedServiceCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_HOST_CHECK
// SCHEDULE_HOST_CHECK;<host_name>;<check_time>
func (a *Api) HandleScheduleHostCheck(w http.ResponseWriter, r *http.Request) {
}

// SCHEDULE_HOST_DOWNTIME
// SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
func (a *Api) HandleScheduleHostDowntime(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Hostname  string `json:"hostname"`
		StartTime int64  `json:"start_time"`
		EndTime   int64  `json:"end_time"`
		Fixed     uint8  `json:"fixed"`
		TriggerId int64  `json:"trigger_id"`
		Duration  int64  `json:"duration"`
		Author    string `json:"author"`
		Comment   string `json:"comment"`
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	if data.Hostname == "" {
		http.Error(w, "Missing host", http.StatusBadRequest)
		return
	}

	if data.Author == "" {
		http.Error(w, fmt.Sprintf("Error: Author field is required"), http.StatusBadRequest)
		return
	}

	if data.Comment == "" {
		http.Error(w, fmt.Sprintf("Error: Comment can not be empty"), http.StatusBadRequest)
		return
	}

	if data.StartTime >= data.EndTime {
		http.Error(w, "start_time must be less than end_time", http.StatusBadRequest)
	}

	if data.Duration == 0 {
		http.Error(w, "duration of maintenance must be greater than 0 seconds", http.StatusBadRequest)
	}

	command := fmt.Sprintf("%s;%s;%d;%d;%d;%d;%d;%s;%s", "SCHEDULE_HOST_DOWNTIME", data.Hostname, data.StartTime, data.EndTime, data.Fixed, data.TriggerId, data.Duration, data.Author, data.Comment)
	a.WriteCommandToFile(w, command)
}

func (a *Api) WriteCommandToFile(w http.ResponseWriter, command string) {
	if err := a.WriteCommand(command); err != nil {
		http.Error(w, "Could not execute command", http.StatusInternalServerError)
		return
	}
}
