package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sulochan/go-nagios-api/config"
)

var (
	contactList       []map[string]string
	contactStatusList []map[string]string
	serviceStatusList []map[string]string
	serviceList       []map[string]string
	hostStatusList    []map[string]string
	hostList          []map[string]string
	hostgroupList     []map[string]string
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type Info struct {
	Created          string `json:"created"`
	Version          string `json:"version"`
	LastUpdatedCheck string `json:"last_update_check"`
	UpdateAvailable  string `json:"update_available"`
	LastVersion      string `json:"last_version"`
	NewVersion       string `json:"new_version"`
}

type ContactStatus struct {
	ContactName                 string `json:"contact_name"`
	ModifiedAttributes          string `json:"modified_attributes"`
	ModifiedHostAttributes      string `json:"modified_host_attributes"`
	ModifiedServiceAttributes   string `json:"modified_service_attributes"`
	HostNotificationPeriod      string `json:"host_notification_period"`
	ServiceNotification_period  string `json:"service_notification_period"`
	LastHostNotification        string `json:"last_host_notification"`
	LastServiceNotification     string `json:"last_service_notification"`
	HostNotificationsEnabled    string `json:"host_notifications_enabled"`
	ServiceNotificationsEnabled string `json:"service_notifications_enabled"`
}

func init() {
	readObjectCache()
	go spawnRefreshRoutein()
}

func spawnRefreshRoutein() {
	for {
		refreshStatusData()
		time.Sleep(60 * time.Second)
	}
}

func readObjectCache() {
	conf := config.GetConfig()
	log.Printf("Reading object cache from %s", conf.ObjectCacheFile)
	log.Printf("Writing commands to %s", conf.CommandFile)
	dat, err := ioutil.ReadFile(conf.ObjectCacheFile)
	if err != nil {
		panic(err)
	}

	a := strings.SplitAfterN(string(dat), "}", -1)
	for _, i := range a {
		lines := strings.Split(i, "\n")
		if stringInSlice("define command {", lines) {
			// We dont do anything with defined commands for now
		}

		if stringInSlice("define contactgroup {", lines) {
			// We dont do anything with contactgroup for now
		}

		if stringInSlice("define hostgroup {", lines) {
			thisgroup := map[string]string{}
			for _, i := range lines {
				if i == "define hostgroup {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
					thisgroup[strings.TrimSpace(strings.Fields(i)[0])] = strings.TrimSpace(strings.Fields(i)[1])
				}
			}
			hostgroupList = append(hostgroupList, thisgroup)
		}

		if stringInSlice("define contact {", lines) {
			thiscontact := map[string]string{}
			for _, i := range lines {
				if i == "define contact {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
					thiscontact[strings.TrimSpace(strings.Fields(i)[0])] = strings.TrimSpace(strings.Fields(i)[1])
				}
			}
			contactList = append(contactList, thiscontact)
		}

		if stringInSlice("define host {", lines) {
			thishost := map[string]string{}
			for _, i := range lines {
				if i == "define host {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
					thishost[strings.TrimSpace(strings.Fields(i)[0])] = strings.TrimSpace(strings.Fields(i)[1])
				}
			}
			hostList = append(hostList, thishost)
		}

		if stringInSlice("define service {", lines) {
			thisservice := map[string]string{}
			for _, i := range lines {
				if i == "define service {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
					thisservice[strings.TrimSpace(strings.Fields(i)[0])] = strings.TrimSpace(strings.Fields(i)[1])
				}
			}
			serviceList = append(serviceList, thisservice)
		}

	}
}

func refreshStatusData() {
	conf := config.GetConfig()
	log.Printf("Refreshig data from %s", conf.StatusFile)
	dat, err := ioutil.ReadFile(conf.StatusFile)
	if err != nil {
		panic(err)
	}

	a := strings.SplitAfterN(string(dat), "}", -1)
	for _, i := range a {
		lines := strings.Split(i, "\n")
		if stringInSlice("contactstatus {", lines) {
			contactstatus := map[string]string{}
			for _, i := range lines {
				if i == "contactstatus {" || i == "    }" || i == "" {
				} else if strings.TrimSpace(strings.Split(i, " ")[0]) == "}" {
					// Ignore these lines
				} else {
					contactstatus[strings.TrimSpace(strings.Split(i, "=")[0])] = strings.TrimSpace(strings.Split(i, "=")[1])
				}
			}
			contactStatusList = append(contactStatusList, contactstatus)
		}

		if stringInSlice("servicestatus {", lines) {
			service := map[string]string{}
			for _, i := range lines {
				if i == "servicestatus {" || i == "    }" || i == "" {
				} else if strings.TrimSpace(strings.Split(i, " ")[0]) == "}" {
					// Ignore these lines
				} else {
					service[strings.TrimSpace(strings.Split(i, "=")[0])] = strings.TrimSpace(strings.Split(i, "=")[1])
				}
			}
			serviceStatusList = append(serviceStatusList, service)
		}

		if stringInSlice("hoststatus {", lines) {
			host := map[string]string{}
			for _, i := range lines {
				if i == "hoststatus {" || i == "	}" || i == "" {
					// Ignore these lines
				} else {
					host[strings.TrimSpace(strings.Split(i, "=")[0])] = strings.TrimSpace(strings.Split(i, "=")[1])
				}
			}
			hostStatusList = append(hostStatusList, host)
		}
	}
}

// HandleGetContacts returns all configured contactlist
// GET: /contacts
func HandleGetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactList)
}

// HandleGetAllHostStatus returns hoststatus for all hosts
// GET: /hoststatus
func HandleGetAllHostStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hostStatusList)
}

// HandleGetHostStatusForHost returns hoststatus for requested host only
// GET: /hoststatus/<host>
func HandleGetHostStatusForHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	if !ok {
		http.Error(w, "Could not find host to lookup", 400)
		return
	}

	for _, item := range hostStatusList {
		if item["host_name"] == host {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Host not found", 404)
	return
}

// HandleGetServiceStatus return all servicestatus
// GET: /servicestatus
func HandleGetServiceStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceStatusList)
}

// HandleGetServiceStatusForService returns all servicestatus for requested service only
// GET: /servicestatus/<service>
func HandleGetServiceStatusForService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, ok := vars["service"]
	if !ok {
		http.Error(w, "Could not find service to lookup", 400)
		return
	}

	var serviceList []map[string]string
	for _, item := range serviceStatusList {
		if item["service_description"] == service {
			serviceList = append(serviceList, item)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceList)
	return
}

// HandleGetHost retruns host info only on the host requested
// GET: /host/<hostname>
func HandleGetHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	if !ok {
		http.Error(w, "Invalid hostname provided", 400)
		return
	}

	for _, item := range hostList {
		if item["host_name"] == host {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Host Not Found", 404)
	return
}

// HandleGetConfiguredHosts returns a list with configured host names
// GET: /hosts
func HandleGetConfiguredHosts(w http.ResponseWriter, r *http.Request) {
	var thesehosts []string
	for _, item := range hostList {
		h := item["host_name"]
		if !stringInSlice(h, thesehosts) {
			thesehosts = append(thesehosts, h)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thesehosts)
}

// HandleGetConfiguredServices returns a list with configured service names
// GET: /services
func HandleGetConfiguredServices(w http.ResponseWriter, r *http.Request) {
	var services []string
	for _, item := range serviceStatusList {
		if !stringInSlice(item["service_description"], services) {
			services = append(services, item["service_description"])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

type hostGroup struct {
	HostGroupName string   `json:"hostgroup_name"`
	Alias         string   `json:"alias"`
	Members       []string `json:"members"`
}

// HandleGetHostGroups returns all defined hostgroups
// GET: /hostgroups
func HandleGetHostGroups(w http.ResponseWriter, r *http.Request) {
	var hg []hostGroup
	for _, item := range hostgroupList {
		group := hostGroup{HostGroupName: item["hostgroup_name"], Alias: item["alias"], Members: strings.Split(item["members"], ",")}
		hg = append(hg, group)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hg)
}
