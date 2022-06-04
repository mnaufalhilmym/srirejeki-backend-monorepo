package filehandler

import (
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"log"
	"os"
	"os/signal"
)

func CloseFileOnInterupt(file *os.File) {
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		// do last actions and wait for all write operations to end
		log.Printf("Closing file log...\n")
		if err := file.Close(); err != nil {
			errorHandler.LogErrorThenContinue("fileHandler/Close1", err)
		}
		os.Exit(0)
	}()
}
