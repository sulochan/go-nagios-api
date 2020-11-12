package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr            string
	ObjectCacheFile string
	StatusFile      string
	CommandFile     string
}

var (
	config          *Config
	configfile      *string
	objectCacheFile *string
	statusFile      *string
	commandFile     *string
	addr            *string
)

func init() {
	configfile = flag.String("config", "", "path to config file")
	objectCacheFile = flag.String("cachefile", "/usr/local/nagios/var/objects.cache", "Nagios object.cache file location")
	statusFile = flag.String("statusfile", "/usr/local/nagios/var/status.dat", "Nagios status.dat file location")
	commandFile = flag.String("commandfile", "/usr/local/nagios/var/nagios.cmd", "Nagios command file location")
	addr = flag.String("addr", ":9090", "The interface and port to run server on")
	flag.Parse()

	if *configfile != "" {
		loadConfigFile()
	} else {
		loadConfigFlags()
	}
}

func loadConfigFlags() {
	config = &Config{Addr: *addr, ObjectCacheFile: *objectCacheFile, StatusFile: *statusFile, CommandFile: *commandFile}
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
	config = temp
}

func GetConfig() *Config {
	return config
}
