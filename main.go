package main

import (
	"log"

	"github.com/sulochan/go-nagios-api/api"
	"github.com/sulochan/go-nagios-api/config"
)

func main() {
	conf := config.GetConfig()
	api := api.NewApi(conf.Addr)

	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
