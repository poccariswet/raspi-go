package main

import (
	"log"
	"time"

	"github.com/jleight/omxplayer"
)

func main() {
	omxplayer.SetUser("pi", "/home/pi")
	player, err := omxplayer.New("/opt/vc/src/hello_pi/hello_video/test.h264")
	if err != nil {
		log.Fatalln(err, "00000000000000000")
	}
	//_, err = player.CanPlay()
	//if err != nil {
	//	log.Fatalln(err, "$$$$$$$$$$$$$$$$$")
	//}

	player.WaitForReady()
	time.Sleep(5 * time.Second)
	err = player.PlayPause()
	if err != nil {
		log.Fatalln(err, "11111111111111111")
	}

	time.Sleep(5 * time.Second)
	err = player.ShowSubtitles()
	if err != nil {
		log.Fatalln(err, "22222222222222222")
	}

	time.Sleep(5 * time.Second)
	err = player.Quit()
	if err != nil {
		log.Fatalln(err, "33333333333333333")
	}
}
