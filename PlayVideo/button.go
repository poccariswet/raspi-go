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

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pushed")

			if ok == true {
				out, err := exec.Command("killall", "omxplayer").Output()
				ok = false
				if err != nil {
					log.Println(err)
					os.Exit(100)
				}
				fmt.Println(string(out))
			}

			//再生
			if ok != true {
				out, err := exec.Command("omxplayer", "-b", "/opt/vc/src/hello_pi/hello_video/test.h264").Output()
				if err != nil {
					log.Print(err)
					os.Exit(200)
				} else {
					ok = true
					fmt.Println(string(out))
				}
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
