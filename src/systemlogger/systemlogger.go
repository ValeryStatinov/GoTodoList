package systemlogger

import (
	"log"
	"os"
)

func Log(msgs ...string) {
	f, err := os.OpenFile("system.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	result := ""
	for _, msg := range msgs {
		result = result + " " + msg
	}

	logger := log.New(f, "", log.LstdFlags|log.LUTC)
	logger.Println(result)
}
