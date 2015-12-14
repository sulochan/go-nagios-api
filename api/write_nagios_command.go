package api

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

func WriteCommand(command string) error {
	timeNow := time.Now()
	s := fmt.Sprintf("[%d]", timeNow.Unix())
	commandToWrite := fmt.Sprintf("%s %s%s", s, command, "\n")

	f, err := os.OpenFile(*commandFile, os.O_APPEND|os.O_WRONLY, 0600)
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
