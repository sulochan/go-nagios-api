package api

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ContactStatus struct {
	ContactName                 string            `json:"contact_name"`
	HostNotificationPeriod      string            `json:"host_notification_period"`
	HostNotificationsEnabled    string            `json:"host_notifications_enabled"`
	LastHostNotification        string            `json:"last_host_notification"`
	LastServiceNotification     string            `json:"last_service_notification"`
	ModifiedAttributes          string            `json:"modified_attributes"`
	ModifiedHostAttributes      string            `json:"modified_host_attributes"`
	ModifiedServiceAttributes   string            `json:"modified_service_attributes"`
	ServiceNotificationPeriod   string            `json:"service_notification_period"`
	ServiceNotificationsEnabled string            `json:"service_notifications_enabled"`
	CustomVariables             map[string]string `json:"custom_variables,omitempty"`
}
type HostStatus struct {
	AcknowledgementType        string            `json:"acknowledgement_type"`
	ActiveChecksEnabled        string            `json:"active_checks_enabled"`
	CheckCommand               string            `json:"check_command"`
	CheckExecutionTime         string            `json:"check_execution_time"`
	CheckInterval              string            `json:"check_interval"`
	CheckLatency               string            `json:"check_latency"`
	CheckOptions               string            `json:"check_options"`
	CheckPeriod                string            `json:"check_period"`
	CheckType                  string            `json:"check_type"`
	CurrentAttempt             string            `json:"current_attempt"`
	CurrentEventId             string            `json:"current_event_id"`
	CurrentNotificationId      string            `json:"current_notification_id"`
	CurrentNotificationNumber  string            `json:"current_notification_number"`
	CurrentProblemId           string            `json:"current_problem_id"`
	CurrentState               string            `json:"current_state"`
	EventHandler               string            `json:"event_handler"`
	EventHandlerEnabled        string            `json:"event_handler_enabled"`
	FlapDetectionEnabled       string            `json:"flap_detection_enabled"`
	HasBeenChecked             string            `json:"has_been_checked"`
	HostName                   string            `json:"host_name"`
	IsFlapping                 string            `json:"is_flapping"`
	LastCheck                  string            `json:"last_check"`
	LastEventId                string            `json:"last_event_id"`
	LastHardState              string            `json:"last_hard_state"`
	LastHardStateChange        string            `json:"last_hard_state_change"`
	LastNotification           string            `json:"last_notification"`
	LastProblemId              string            `json:"last_problem_id"`
	LastStateChange            string            `json:"last_state_change"`
	LastTimeDown               string            `json:"last_time_down"`
	LastTimeUnreachable        string            `json:"last_time_unreachable"`
	LastTimeUp                 string            `json:"last_time_up"`
	LastUpdate                 string            `json:"last_update"`
	LongPluginOutput           string            `json:"long_plugin_output"`
	MaxAttempts                string            `json:"max_attempts"`
	ModifiedAttributes         string            `json:"modified_attributes"`
	NextCheck                  string            `json:"next_check"`
	NextNotification           string            `json:"next_notification"`
	NoMoreNotifications        string            `json:"no_more_notifications"`
	NotificationPeriod         string            `json:"notification_period"`
	NotificationsEnabled       string            `json:"notifications_enabled"`
	Obsess                     string            `json:"obsess"`
	PassiveChecksEnabled       string            `json:"passive_checks_enabled"`
	PercentStateChange         string            `json:"percent_state_change"`
	PerformanceData            string            `json:"performance_data"`
	PluginOutput               string            `json:"plugin_output"`
	ProblemHasBeenAcknowledged string            `json:"problem_has_been_acknowledged"`
	ProcessPerformanceData     string            `json:"process_performance_data"`
	RetryInterval              string            `json:"retry_interval"`
	ScheduledDowntimeDepth     string            `json:"scheduled_downtime_depth"`
	ShouldBeScheduled          string            `json:"should_be_scheduled"`
	StateType                  string            `json:"state_type"`
	CustomVariables            map[string]string `json:"custom_variables,omitempty"`
}
type ProgramStatus struct {
	ActiveHostChecksEnabled          string            `json:"active_host_checks_enabled"`
	ActiveOndemandHostCheckStats     string            `json:"active_ondemand_host_check_stats"`
	ActiveOndemandServiceCheckStats  string            `json:"active_ondemand_service_check_stats"`
	ActiveScheduledHostCheckStats    string            `json:"active_scheduled_host_check_stats"`
	ActiveScheduledServiceCheckStats string            `json:"active_scheduled_service_check_stats"`
	ActiveServiceChecksEnabled       string            `json:"active_service_checks_enabled"`
	CachedHostCheckStats             string            `json:"cached_host_check_stats"`
	CachedServiceCheckStats          string            `json:"cached_service_check_stats"`
	CheckHostFreshness               string            `json:"check_host_freshness"`
	CheckServiceFreshness            string            `json:"check_service_freshness"`
	DaemonMode                       string            `json:"daemon_mode"`
	EnableEventHandlers              string            `json:"enable_event_handlers"`
	EnableFlapDetection              string            `json:"enable_flap_detection"`
	EnableNotifications              string            `json:"enable_notifications"`
	ExternalCommandStats             string            `json:"external_command_stats"`
	GlobalHostEventHandler           string            `json:"global_host_event_handler"`
	GlobalServiceEventHandler        string            `json:"global_service_event_handler"`
	LastLogRotation                  string            `json:"last_log_rotation"`
	ModifiedHostAttributes           string            `json:"modified_host_attributes"`
	ModifiedServiceAttributes        string            `json:"modified_service_attributes"`
	NagiosPid                        string            `json:"nagios_pid"`
	NextCommentId                    string            `json:"next_comment_id"`
	NextDowntimeId                   string            `json:"next_downtime_id"`
	NextEventId                      string            `json:"next_event_id"`
	NextNotificationId               string            `json:"next_notification_id"`
	NextProblemId                    string            `json:"next_problem_id"`
	ObsessOverHosts                  string            `json:"obsess_over_hosts"`
	ObsessOverServices               string            `json:"obsess_over_services"`
	ParallelHostCheckStats           string            `json:"parallel_host_check_stats"`
	PassiveHostCheckStats            string            `json:"passive_host_check_stats"`
	PassiveHostChecksEnabled         string            `json:"passive_host_checks_enabled"`
	PassiveServiceCheckStats         string            `json:"passive_service_check_stats"`
	PassiveServiceChecksEnabled      string            `json:"passive_service_checks_enabled"`
	ProcessPerformanceData           string            `json:"process_performance_data"`
	ProgramStart                     string            `json:"program_start"`
	SerialHostCheckStats             string            `json:"serial_host_check_stats"`
	CustomVariables                  map[string]string `json:"custom_variables,omitempty"`
}
type ServiceStatus struct {
	AcknowledgementType        string            `json:"acknowledgement_type"`
	ActiveChecksEnabled        string            `json:"active_checks_enabled"`
	CheckCommand               string            `json:"check_command"`
	CheckExecutionTime         string            `json:"check_execution_time"`
	CheckInterval              string            `json:"check_interval"`
	CheckLatency               string            `json:"check_latency"`
	CheckOptions               string            `json:"check_options"`
	CheckPeriod                string            `json:"check_period"`
	CheckType                  string            `json:"check_type"`
	CurrentAttempt             string            `json:"current_attempt"`
	CurrentEventId             string            `json:"current_event_id"`
	CurrentNotificationId      string            `json:"current_notification_id"`
	CurrentNotificationNumber  string            `json:"current_notification_number"`
	CurrentProblemId           string            `json:"current_problem_id"`
	CurrentState               string            `json:"current_state"`
	EventHandler               string            `json:"event_handler"`
	EventHandlerEnabled        string            `json:"event_handler_enabled"`
	FlapDetectionEnabled       string            `json:"flap_detection_enabled"`
	HasBeenChecked             string            `json:"has_been_checked"`
	HostName                   string            `json:"host_name"`
	IsFlapping                 string            `json:"is_flapping"`
	LastCheck                  string            `json:"last_check"`
	LastEventId                string            `json:"last_event_id"`
	LastHardState              string            `json:"last_hard_state"`
	LastHardStateChange        string            `json:"last_hard_state_change"`
	LastNotification           string            `json:"last_notification"`
	LastProblemId              string            `json:"last_problem_id"`
	LastStateChange            string            `json:"last_state_change"`
	LastTimeCritical           string            `json:"last_time_critical"`
	LastTimeOk                 string            `json:"last_time_ok"`
	LastTimeUnknown            string            `json:"last_time_unknown"`
	LastTimeWarning            string            `json:"last_time_warning"`
	LastUpdate                 string            `json:"last_update"`
	LongPluginOutput           string            `json:"long_plugin_output"`
	MaxAttempts                string            `json:"max_attempts"`
	ModifiedAttributes         string            `json:"modified_attributes"`
	NextCheck                  string            `json:"next_check"`
	NextNotification           string            `json:"next_notification"`
	NoMoreNotifications        string            `json:"no_more_notifications"`
	NotificationPeriod         string            `json:"notification_period"`
	NotificationsEnabled       string            `json:"notifications_enabled"`
	Obsess                     string            `json:"obsess"`
	PassiveChecksEnabled       string            `json:"passive_checks_enabled"`
	PercentStateChange         string            `json:"percent_state_change"`
	PerformanceData            string            `json:"performance_data"`
	PluginOutput               string            `json:"plugin_output"`
	ProblemHasBeenAcknowledged string            `json:"problem_has_been_acknowledged"`
	ProcessPerformanceData     string            `json:"process_performance_data"`
	RetryInterval              string            `json:"retry_interval"`
	ScheduledDowntimeDepth     string            `json:"scheduled_downtime_depth"`
	ServiceDescription         string            `json:"service_description"`
	ShouldBeScheduled          string            `json:"should_be_scheduled"`
	StateType                  string            `json:"state_type"`
	CustomVariables            map[string]string `json:"custom_variables,omitempty"`
}

func (o *ContactStatus) setField(key, value string) error {
	return setField(o, key, value)
}

func (o *HostStatus) setField(key, value string) error {
	return setField(o, key, value)
}

func (o *ProgramStatus) setField(key, value string) error {
	return setField(o, key, value)
}

func (o *ServiceStatus) setField(key, value string) error {
	return setField(o, key, value)
}

func (o *ContactStatus) setCustomVariable(key, value string) {
	if o.CustomVariables == nil {
		o.CustomVariables = make(map[string]string)
	}
	o.CustomVariables[key] = value
}

func (o *HostStatus) setCustomVariable(key, value string) {
	if o.CustomVariables == nil {
		o.CustomVariables = make(map[string]string)
	}
	o.CustomVariables[key] = value
}

func (o *ProgramStatus) setCustomVariable(key, value string) {
	if o.CustomVariables == nil {
		o.CustomVariables = make(map[string]string)
	}
	o.CustomVariables[key] = value
}

func (o *ServiceStatus) setCustomVariable(key, value string) {
	if o.CustomVariables == nil {
		o.CustomVariables = make(map[string]string)
	}
	o.CustomVariables[key] = value
}

// setField sets a field in a struct based on the JSON tag associated with the struct
func setField(obj interface{}, name string, value interface{}) error {
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		elem := val.Type().Field(i)
		tag := elem.Tag
		field := val.Field(i)

		js := tag.Get("json")
		if js == "" {
			continue
		}

		comma := strings.Index(js, ",")
		if comma != -1 {
			js = js[0:comma]
		}

		if js == name {
			if !field.CanSet() {
				return fmt.Errorf("Cannot set %s field value", name)
			}

			valValue := reflect.ValueOf(value)
			if field.Type() != valValue.Type() {
				return errors.New("Provided value type didn't match obj field type")
			}

			field.Set(valValue)
			return nil
		}
	}
	return fmt.Errorf("No such field: %s in obj", name)
}
