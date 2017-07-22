package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"os"
	"os/exec"
)

func stop() {
	out, err := exec.Command("killall", "mpg321").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func start() {
	out, err := exec.Command("sh", "-c", "mpg321 ./music/morning.mp3 > /dev/null 2>&1 &").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(string(out))
}


func main() {
	start()
	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pushed")
			stop()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
	// stop()
}
