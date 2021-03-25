package main

import (
	_ "image/jpeg"
	"log"
	"time"
)

func bc() {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			bc()
		}
	}()

	for {
		if !flag {
			time.Sleep(3000 * time.Millisecond)
			continue
		}
		frame, err := getImage()
		if err != nil {
			log.Println(err)
		}

		br, err := getBvalue(frame)
		if err != nil {
			log.Println(err)
		}

		val := getAbsB(br)

		err = setScreenBRightnes(val)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(900 * time.Millisecond)
	}
}
