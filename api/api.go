package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Api struct {
	router          *mux.Router
	addr            string
	fileObjectCache string
	fileCommand     string
	fileStatus      string
	statusData      *StatusData
	staticData      *StaticData
	mutex           sync.RWMutex
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func NewApi(addr, fileObjectCache, fileCommand, fileStatus string) *Api {
	api := &Api{
		addr:            addr,
		router:          mux.NewRouter(),
		fileObjectCache: fileObjectCache,
		fileCommand:     fileCommand,
		fileStatus:      fileStatus,
	}

	api.buildRoutes()
	return api
}

func (s *Api) Run() error {
	log.Println("Reading object cache from ", s.fileObjectCache)
	log.Println("Writing commands to ", s.fileCommand)

	oc, err := os.Open(s.fileObjectCache)
	if err != nil {
		return err
	}
	defer oc.Close()

	s.staticData, err = readObjectCache(oc)
	if err != nil {
		return fmt.Errorf("Unable to parse object cache file: %s", err)
	}
	go s.spawnRefreshRoutein()

	http.Handle("/", s.router)

	err = http.ListenAndServe(s.addr, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Api) spawnRefreshRoutein() {
	for {
		data, err := s.refreshStatusDataFile()
		if err != nil {
			log.Println("Unable to refresh status data: ", err)
		} else {
			s.mutex.Lock()
			s.statusData = data
			s.mutex.Unlock()
		}
		time.Sleep(60 * time.Second)
	}
}

func readObjectCache(in io.Reader) (*StaticData, error) {
	data := NewStaticData()
	dat, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
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
          if(len(strings.Fields(i)) != 0){
            thisgroup[strings.TrimSpace(strings.Fields(i)[0])] = strings.Join(strings.Fields(i)[1:], " ")
          }
				}
			}
			data.hostgroupList = append(data.hostgroupList, thisgroup)
		}

		if stringInSlice("define contact {", lines) {
			thiscontact := map[string]string{}
			for _, i := range lines {
				if i == "define contact {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
          if(len(strings.Fields(i)) != 0) {
					  thiscontact[strings.TrimSpace(strings.Fields(i)[0])] = strings.Join(strings.Fields(i)[1:], " ")
				  }
				}
			}
			data.contactList = append(data.contactList, thiscontact)
		}

		if stringInSlice("define host {", lines) {
			thishost := map[string]string{}
			for _, i := range lines {
				if i == "define host {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
          if(len(strings.Fields(i)) != 0) {
					  thishost[strings.TrimSpace(strings.Fields(i)[0])] = strings.Join(strings.Fields(i)[1:], " ")
			    }
				}
			}
			data.hostList = append(data.hostList, thishost)
		}

		if stringInSlice("define service {", lines) {
			thisservice := map[string]string{}
			for _, i := range lines {
				if i == "define service {" || strings.TrimSpace(i) == "}" || i == "" {
					// Ignore these lines
				} else {
          if(len(strings.Fields(i)) != 0) {
					  thisservice[strings.TrimSpace(strings.Fields(i)[0])] = strings.Join(strings.Fields(i)[1:], " ")
			    }
				}
			}
			data.serviceList = append(data.serviceList, thisservice)
		}

	}

	return data, nil
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

type StaticData struct {
	contactList   []map[string]string
	serviceList   []map[string]string
	hostList      []map[string]string
	hostgroupList []map[string]string
}

func NewStaticData() *StaticData {
	return &StaticData{}
}

type settableType interface {
	setField(key, value string) error
	setCustomVariable(key, value string)
}

func parseBlock(o settableType, objecttype string, lines []string) error {
	start := objecttype + " {"
	for _, i := range lines {
		if i == start || i == "    }" || i == "" || strings.TrimSpace(strings.Split(i, " ")[0]) == "}" {
			// Ignore these lines
		} else {
			pieces := strings.SplitN(strings.TrimSpace(i), "=", 2)
			if strings.HasPrefix(pieces[0], "_") {
				o.setCustomVariable(strings.TrimPrefix(pieces[0], "_"), strings.SplitN(pieces[1], ";", 2)[1])
			} else {
				o.setField(pieces[0], pieces[1])
			}
		}
	}

	return nil
}

func (s *Api) refreshStatusDataFile() (*StatusData, error) {
	log.Println("Refreshig data from ", s.fileStatus)

	fh, err := os.Open(s.fileStatus)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	return refreshStatusData(fh)
}

func refreshStatusData(fh io.Reader) (*StatusData, error) {
	data := NewStatusData()
	dat, err := ioutil.ReadAll(fh)
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
func (a *Api) HandleGetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.staticData.contactList)
}

// HandleGetAllHostStatus returns hoststatus for all hosts
// GET: /hoststatus
func (a *Api) HandleGetAllHostStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	json.NewEncoder(w).Encode(a.statusData.Hosts)
}

// HandleGetHostStatusForHost returns hoststatus for requested host only
// GET: /hoststatus/<host>
func (a *Api) HandleGetHostStatusForHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	if !ok {
		http.Error(w, "Could not find host to lookup", 400)
		return
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()
	for _, item := range a.statusData.Hosts {
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
func (a *Api) HandleGetServiceStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	json.NewEncoder(w).Encode(a.statusData.Services)
}

// HandleGetServiceStatusForService returns all servicestatus for requested service only
// GET: /servicestatus/<service>
func (a *Api) HandleGetServiceStatusForService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, ok := vars["service"]
	if !ok {
		http.Error(w, "Could not find service to lookup", 400)
		return
	}

	var serviceList []*ServiceStatus
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	for _, item := range a.statusData.Services {
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
func (a *Api) HandleGetHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	log.Println(host)
	if !ok {
		http.Error(w, "Invalid hostname provided", 400)
		return
	}

	for _, item := range a.staticData.hostList {
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
func (a *Api) HandleGetServicesForHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host, ok := vars["hostname"]
	if !ok {
		http.Error(w, "Invalid hostname provided", 400)
		return
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	sList, ok := a.statusData.HostServices[host]
	if !ok {
		http.Error(w, "Host Not Found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sList)
}

// HandleGetConfiguredHosts returns a list with configured host names
// GET: /hosts
func (a *Api) HandleGetConfiguredHosts(w http.ResponseWriter, r *http.Request) {
	var thesehosts []string
	for _, item := range a.staticData.hostList {
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
func (a *Api) HandleGetConfiguredServices(w http.ResponseWriter, r *http.Request) {
	var services []string
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	for _, item := range a.statusData.Services {
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
func (a *Api) HandleGetHostGroups(w http.ResponseWriter, r *http.Request) {
	var hg []hostGroup
	for _, item := range a.staticData.hostgroupList {
		group := hostGroup{HostGroupName: item["hostgroup_name"], Alias: item["alias"], Members: strings.Split(item["members"], ",")}
		hg = append(hg, group)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hg)
}
