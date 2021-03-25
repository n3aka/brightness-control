package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TheCreeper/go-notify"
)

// crw-rw---- 1 root input 13, 68 Feb 15 13:35 event4
func notifier(arg interface{}) {
	n := ""
	var t int32 = 1000
	switch arg.(type) {
	case int:
		n = strconv.Itoa(arg.(int))
	case string:
		n = arg.(string)
	case error:
		t = 10000
		n = fmt.Sprintf("%v", arg)
		log.Println(fmt.Sprintf("%v", arg))
	default:
	}
	ntf := notify.NewNotification("Brightnes Control", n)
	ntf.Timeout = t
	ntf.AppIcon = "presence"
	if _, err := ntf.Show(); err != nil {
		return
	}
}
