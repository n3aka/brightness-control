package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {

	err = initCam()
	if err != nil {
		log.Println("unable to init the cam.. exiting")
		restore()
		return
	}
	defer cam.Close()

	go keyLogger()

	go bc()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	sig := <-signalChannel
	switch sig {
	case os.Interrupt:
		err = restore()
		if err != nil {
			log.Println(err.Error())
		}
		return
	case syscall.SIGTERM:
		err = restore()
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
}
