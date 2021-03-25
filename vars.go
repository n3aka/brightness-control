package main

import "github.com/blackjack/webcam"

const (
	V4L2_PIX_FMT_PJPG = 0x47504A50
	V4L2_PIX_FMT_YUYV = 0x56595559
)

var (
	supportedFormats = map[webcam.PixelFormat]bool{
		V4L2_PIX_FMT_PJPG: true,
		V4L2_PIX_FMT_YUYV: false,
	}

	flag bool

	cor int

	err error

	quitBC     = make(chan bool)
	quitLogger = make(chan bool)

	camControl = map[string]webcam.ControlID{
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

	cam *webcam.Webcam
)

type FrameSizes []webcam.FrameSize
