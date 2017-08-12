package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan bool)

	go func(q chan<- bool) {
		time.Sleep(1 * time.Second)
		q <- true
	}(q)

	fmt.Println(<-q) // q に何か入るまで待つ

}
