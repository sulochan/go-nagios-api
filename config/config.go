package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Port            string
	ObjectCacheFile string
	StatusFile      string
	CommandFile     string
	Logfile         string
}

var (
	config          *Config
	configLock      = new(sync.RWMutex)
	configfile      *string
	objectCacheFile *string
	statusFile      *string
	commandFile     *string
	port            *string
)

func init() {
	configfile = flag.String("config", "", "path to config file")
	objectCacheFile = flag.String("cachefile", "/usr/local/nagios/var/objects.cache", "Nagios object.cache file location")
	statusFile = flag.String("statusfile", "/usr/local/nagios/var/status.dat", "Nagios status.dat file location")
	commandFile = flag.String("commandfile", "/usr/local/nagios/var/nagios.cmd", "Nagios command file location")
	port = flag.String("port", "9090", "Port to run server on")
	flag.Parse()

	if *configfile != "" {
		loadConfigFile()
	} else {
		loadConfigFlags()
	}
}

func loadConfigFlags() {
	config = &Config{Port: *port, ObjectCacheFile: *objectCacheFile, StatusFile: *statusFile, CommandFile: *commandFile}
}

func loadConfigFile() {
	file, err := ioutil.ReadFile(*configfile)
	if err != nil {
		log.Fatal("open config: ", err)
	}

	temp := new(Config)
	if err = json.Unmarshal(file, temp); err != nil {
		log.Fatal("parse config: ", err)
	}
	configLock.Lock()
	config = temp
	configLock.Unlock()
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}
