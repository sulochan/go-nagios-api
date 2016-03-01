package api

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/sulochan/go-nagios-api/config"
)

func WriteCommand(command string) error {
	conf := config.GetConfig()
	timeNow := time.Now()
	s := fmt.Sprintf("[%d]", timeNow.Unix())
	commandToWrite := fmt.Sprintf("%s %s%s", s, command, "\n")

	f, err := os.OpenFile(conf.CommandFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Error(err)
		return err
	}

	if _, err = f.WriteString(commandToWrite); err != nil {
		log.Error(err)
		return err
	}

	defer f.Close()
	return nil
}
