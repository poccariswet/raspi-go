package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var ok bool

func main() {
	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11")
	ch := make(chan bool)

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pushed")

			if ok == true {
				if <-ch == true {
					out, err := exec.Command("killall", "omxplayer.bin").Output()
					if err != nil {
						log.Println(err)
						os.Exit(100)
					}
					fmt.Println(string(out))
				}
			}

			if ok != true {
				go func(ch chan<- bool) {
					ch <- true
					ok = true
					out, err := exec.Command("omxplayer", "/opt/vc/src/hello_pi/hello_video/test.h264").Output()
					if err != nil {
						log.Println(err)
						os.Exit(100)
					}
					fmt.Println(string(out))
				}(ch)
			}
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
}
