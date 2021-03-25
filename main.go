// Example program that uses blakjack/webcam library
// for working with V4L2 devices.
// The application reads frames from device and writes them to stdout
// If your device supports motion formats (e.g. H264 or MJPEG) you can
// use it's output as a video stream.
// Example usage: go run stdout_streamer.go | vlc -
package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/blackjack/webcam"
)

type FrameSizes []webcam.FrameSize

func (slice FrameSizes) Len() int {
	return len(slice)
}

//For sorting purposes
func (slice FrameSizes) Less(i, j int) bool {
	ls := slice[i].MaxWidth * slice[i].MaxHeight
	rs := slice[j].MaxWidth * slice[j].MaxHeight
	return ls < rs
}

//For sorting purposes
func (slice FrameSizes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

const (
	V4L2_PIX_FMT_PJPG = 0x47504A50
	V4L2_PIX_FMT_YUYV = 0x56595559
)

var supportedFormats = map[webcam.PixelFormat]bool{
	V4L2_PIX_FMT_PJPG: true,
	V4L2_PIX_FMT_YUYV: false,
}

var flag bool

func main() {

	go keyLogger()

	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		panic(err.Error())
	}
	defer cam.Close()
	// 1448695129
	_, _, _, err = cam.SetImageFormat(1196444237, uint32(160), uint32(120))
	if err != nil {
		panic(err.Error())
	}

	ma := map[string]webcam.ControlID{
		"white_balance":          9963802,
		"backlight_compensation": 9963804,
		"exposure_auto":          10094849,
		"brightness":             9963776,
		"contrast":               9963777,
		"saturation":             9963778,
		"gamma":                  9963792,
		"power_line":             9963800,
		"hue":                    9963779,
		"white_balance_auto":     9963788,
		"gain":                   9963795,
		"sharpness":              9963803,
		"exposure":               10094850,
	}

	err = cam.SetControl(ma["power_line"], int32(0))
	err = cam.SetControl(ma["white_balance"], int32(2800))
	err = cam.SetControl(ma["white_balance_auto"], int32(0))
	err = cam.SetControl(ma["backlight_compensation"], int32(0))
	err = cam.SetControl(ma["contrast"], int32(0))
	err = cam.SetControl(ma["saturation"], int32(0))
	err = cam.SetControl(ma["sharpness"], int32(1))
	err = cam.SetControl(ma["exposure_auto"], int32(1))
	err = cam.SetControl(ma["exposure"], int32(0))
	err = cam.SetControl(ma["brightness"], int32(0))
	err = cam.SetControl(ma["gain"], int32(1))
	err = cam.SetControl(ma["hue"], int32(0))
	err = cam.SetControl(ma["gamma"], int32(100))

	err = cam.StartStreaming()

	if err != nil {
		panic(err.Error())
	}

	for {

		timeout := uint32(5) //5 seconds
		err = cam.WaitForFrame(timeout)

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			fmt.Fprint(os.Stderr, err.Error())
			continue
		default:
			panic(err.Error())
		}
		frame, err := cam.ReadFrame()

		if len(frame) != 0 {
			rdr := bytes.NewReader(frame)

			config, _, err := image.Decode(rdr)
			if err != nil {
				log.Println("error " + err.Error())

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
			br := Res / 407040
			if br < 700 || br > 10000 {
				continue
			}
		END:
			for i := 0; i < 10; i++ {
				d, err := f(br)
				if err == nil {
					err = ioutil.WriteFile("/sys/class/backlight/intel_backlight/brightness", d, 0644)
					if err != nil {
						log.Println(err)
					}
					break END
				}
				br = br + 1
			}
		} else if err != nil {
			panic(err.Error())
		}
	F7:
		for {
			if flag {
				// log.Println("unlocked!")
				break F7
			}
			time.Sleep(3000 * time.Millisecond)
			// log.Println("locked!")
		}

		time.Sleep(700 * time.Millisecond)
	}
}
