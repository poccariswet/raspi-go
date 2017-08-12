package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	q := make(chan bool)
  music := "Music/iTunes/iTunes\ Media/Music/10-FEET/thread/03\ その向こうへ.m4a"

	go func(q chan<- bool) { // send
		time.Sleep(3 * time.Second)
		out, err := exec.Command("afplay", music ,"&").Output()
		if err != nil {
			log.Println(err)
			os.Exit(100)
		}
		fmt.Println(string(out))
		q <- true
	}(q)

	//fmt.Println(<-q) // q に何か入るまで待つ
	if ok := <-q; ok == true {
		out, err := exec.Command("killall", "afplay").Output()
		if err != nil {
			log.Println(err)
			os.Exit(200)
		}
		fmt.Println(string(out))
	}

}
