package main

import (
	"log"

	"github.com/blackjack/webcam"
)

type camControlModel struct {
}

func prepaire() error {
	_, _, _, err = cam.SetImageFormat(1196444237, uint32(160), uint32(120))
	if err != nil {
		return err
	}

	err = cam.SetControl(camControl["power_line"], int32(0))
	err = cam.SetControl(camControl["white_balance"], int32(2800))
	err = cam.SetControl(camControl["white_balance_auto"], int32(0))
	err = cam.SetControl(camControl["backlight_compensation"], int32(0))
	err = cam.SetControl(camControl["contrast"], int32(0))
	err = cam.SetControl(camControl["saturation"], int32(0))
	err = cam.SetControl(camControl["sharpness"], int32(1))
	err = cam.SetControl(camControl["exposure_auto"], int32(1))
	err = cam.SetControl(camControl["exposure"], int32(0))
	err = cam.SetControl(camControl["brightness"], int32(0))
	err = cam.SetControl(camControl["gain"], int32(1))
	err = cam.SetControl(camControl["hue"], int32(0))
	err = cam.SetControl(camControl["gamma"], int32(100))

	err = cam.SetBufferCount(1)
	err = cam.StartStreaming()
	return err
}
func restore() error {
	err := cam.SetControl(camControl["power_line"], int32(0))
	err = cam.SetControl(camControl["white_balance"], int32(4600))
	err = cam.SetControl(camControl["white_balance_auto"], int32(1))
	err = cam.SetControl(camControl["backlight_compensation"], int32(3))
	err = cam.SetControl(camControl["contrast"], int32(0))
	err = cam.SetControl(camControl["saturation"], int32(64))
	err = cam.SetControl(camControl["sharpness"], int32(2))
	err = cam.SetControl(camControl["exposure_auto"], int32(3))
	err = cam.SetControl(camControl["exposure"], int32(156))
	err = cam.SetControl(camControl["brightness"], int32(0))
	err = cam.SetControl(camControl["gain"], int32(1))
	err = cam.SetControl(camControl["hue"], int32(0))
	err = cam.SetControl(camControl["gamma"], int32(100))
	if err != nil {
		log.Println(err)
	}
	return err
}

func initCam() error {
	cam, err = webcam.Open("/dev/video0")
	return err
}
