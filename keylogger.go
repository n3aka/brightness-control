package main

import (
	"log"

	keylogger "git.olinux.dev/go/brightness-control/lib/keylogger"
)

func keyLogger() {

	k, err := keylogger.New("/dev/input/event4")

	if err != nil {
		log.Println(err)
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			if e.KeyPress() {
				if e.KeyString() == "F7" {
					log.Println("[event] press key ", e.KeyString())
					if !flag {
						prepaire()
						log.Println("running...")
						flag = true
					} else {
						restore()
						log.Println("stopped...")
						flag = false
					}
				}
				if e.KeyString() == "F8" {
					// log.Println("[event] press key ", e.KeyString())
					if cor >= -5000 {
						cor = cor - 500
					}
					notifier(cor)
				}
				if e.KeyString() == "F9" {
					// log.Println("[event] press key ", e.KeyString())
					if cor <= 10000 {
						cor = cor + 500
					}
					notifier(cor)
				}
			}
			break
		}
	}
}
