package main

import (
	"log"
	"time"

	"github.com/jleight/omxplayer"
)

func main() {
	log.SetFlags(log.Lshortfile)
	omxplayer.SetUser("pi", "/")
	player, err := omxplayer.New("/opt/vc/src/hello_pi/hello_video/test.h264")
	if err != nil {
		log.Println(err)
	}

	player.WaitForReady()
	time.Sleep(5 * time.Second)
	err = player.PlayPause()
	if err != nil {
		log.Println(err)
	}

	time.Sleep(5 * time.Second)
	err = player.ShowSubtitles()
	if err != nil {
		log.Println(err)
	}

	time.Sleep(5 * time.Second)
	err = player.Quit()
	if err != nil {
		log.Println(err)
	}
}
