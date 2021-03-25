package main

import (
	"bytes"
	"image"
	"time"

	"github.com/blackjack/webcam"
)


func getImage() ([]byte, error) {
	time.Sleep(time.Millisecond * 700)

	timeout := uint32(5) //5 seconds
	err := cam.WaitForFrame(timeout)

	switch err.(type) {
	case nil:
		frame, index, err := cam.GetFrame()
		err = cam.ReleaseFrame(index)
		return frame, err

	case *webcam.Timeout:
		return nil, err

	default:
	}

	return nil, err
}


func getBvalue(frame []byte) (int, error) {

	if len(frame) == 0 {
		return 0, nil
	}
	rdr := bytes.NewReader(frame)

	config, _, err := image.Decode(rdr)
	if err != nil {
		return 0, err
	}
	var Res uint32
	var T uint
	for i := 0; i < 848; i++ {
		for j := 0; j < 480; j++ {
			x := config.At(i, j)
			r, g, b, _ := x.RGBA()
			Res = Res + ((r*299)+(g*587)+(b*114))/1000
			T = T + 1
		}
	}
	br := int(Res / 407040)
	if br < 400 || br > 10000 {
		return 0, err
	}
	return br, err
}
