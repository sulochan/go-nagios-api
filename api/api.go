package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sulochan/go-nagios-api/config"
)

var (
	contactList   []map[string]string
	serviceList   []map[string]string
	hostList      []map[string]string
	hostgroupList []map[string]string
	statusData    *StatusData
	mutex         sync.RWMutex
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

func Initialize() {
	err := readObjectCache()
	if err != nil {
		log.Fatal("Unable to parse object cache file: ", err)
	}
	go spawnRefreshRoutein()
}

func spawnRefreshRoutein() {
	for {
		data, err := refreshStatusData()
		if err != nil {
			log.Println("Unable to refresh status data: ", err)
		} else {
			mutex.Lock()
			statusData = data
			mutex.Unlock()
		}
		time.Sleep(60 * time.Second)
	}
}

func readObjectCache() error {
	conf := config.GetConfig()
	log.Printf("Reading object cache from %s", conf.ObjectCacheFile)
	log.Printf("Writing commands to %s", conf.CommandFile)
	dat, err := ioutil.ReadFile(conf.ObjectCacheFile)
	if err != nil {
		return err
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

	return nil
}

type StatusData struct {
	Contacts     []*ContactStatus
	Services     []*ServiceStatus
	Hosts        []*HostStatus
	HostServices map[string][]*ServiceStatus
}

func NewStatusData() *StatusData {
	return &StatusData{
		HostServices: make(map[string][]*ServiceStatus),
	}
}

type settableType interface {
	setField(key, value string) error
}

func parseBlock(o settableType, objecttype string, lines []string) error {
	start := objecttype + " {"
	for _, i := range lines {
		if i == start || i == "    }" || i == "" || strings.TrimSpace(strings.Split(i, " ")[0]) == "}" {
			// Ignore these lines
		} else {
			o.setField(
				strings.TrimSpace(strings.Split(i, "=")[0]),
				strings.TrimSpace(strings.Split(i, "=")[1]),
			)
		}
	}

	return nil
}

func refreshStatusData() (*StatusData, error) {
	data := NewStatusData()
	conf := config.GetConfig()
	log.Printf("Refreshig data from %s", conf.StatusFile)
	dat, err := ioutil.ReadFile(conf.StatusFile)
	if err != nil {
		return nil, err
	}

	a := strings.SplitAfterN(string(dat), "}", -1)
	for _, i := range a {
		lines := strings.Split(i, "\n")
		if stringInSlice("contactstatus {", lines) {
			obj := &ContactStatus{}
			parseBlock(obj, "contactstatus", lines)
			data.Contacts = append(data.Contacts, obj)
		}

		if stringInSlice("servicestatus {", lines) {
			obj := &ServiceStatus{}
			parseBlock(obj, "servicestatus", lines)
			data.Services = append(data.Services, obj)

			data.HostServices[obj.HostName] = append(data.HostServices[obj.HostName], obj)
		}

		if stringInSlice("hoststatus {", lines) {
			obj := &HostStatus{}
			parseBlock(obj, "hoststatus", lines)
			data.Hosts = append(data.Hosts, obj)
		}
	}

	return data, nil
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
	mutex.RLock()
	defer mutex.RUnlock()
	json.NewEncoder(w).Encode(statusData.Hosts)
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

	mutex.RLock()
	defer mutex.RUnlock()
	for _, item := range statusData.Hosts {
		if item.HostName == host {
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
	mutex.RLock()
	defer mutex.RUnlock()
	json.NewEncoder(w).Encode(statusData.Services)
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

	var serviceList []*ServiceStatus
	mutex.RLock()
	defer mutex.RUnlock()
	for _, item := range statusData.Services {
		if item.ServiceDescription == service {
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

// HandleGetServicesForHost retruns all services defined for the given host
// GET: /host/<hostname>/services
func HandleGetServicesForHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	if !ok {
		http.Error(w, "Invalid hostname provided", 400)
		return
	}

	sList, ok := statusData.HostServices[host]
	if !ok {
		http.Error(w, "Host Not Found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sList)
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
	mutex.RLock()
	defer mutex.RUnlock()
	for _, item := range statusData.Services {
		if !stringInSlice(item.ServiceDescription, services) {
			services = append(services, item.ServiceDescription)
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
