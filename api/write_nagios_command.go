package api

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func (a *Api) WriteCommand(command string) error {
	commandToWrite := fmt.Sprintf("[%d] %s\n", time.Now().Unix(), command)

	f, err := os.OpenFile(a.fileCommand, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(commandToWrite); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
