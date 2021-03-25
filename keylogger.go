package main

import (
	"log"
	"time"

	"github.com/MarinX/keylogger"
)

func keyLogger() error {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	for {
		// find keyboard device, does not require a root permission
		keyboard := keylogger.FindKeyboardDevice()

		// check if we found a path to keyboard
		if len(keyboard) <= 0 {
			log.Println("No keyboard found...you will need to provide manual input path")
			return nil
		}

		log.Println("Found a keyboard at", keyboard)
		// init keylogger with keyboard
		k, err := keylogger.New(keyboard)
		if err != nil {
			log.Println(err)
		}
		defer k.Close()

		events := k.Read()

		// range of events
		for e := range events {
			switch e.Type {
			// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
			// check the input_event.go for more events
			case keylogger.EvKey:

				// if the state of key is pressed
				if e.KeyPress() {
					if e.KeyString() == "F7" {
						// log.Println("[event] press key ", e.KeyString())
						if flag {
							flag = false
						} else {
							flag = true
						}
					}
				}

				// // if the state of key is released
				// if e.KeyRelease() {
				// 	log.Println("[event] release key ", e.KeyString())
				// }

				break
			}
		}
	}
	time.Sleep(300 * time.Millisecond)
	return nil
}
