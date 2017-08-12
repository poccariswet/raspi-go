package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func start() {
	out, err := exec.Command("sh", "-c", "mpg321 -l 0 ./music/morning.mp3 > /dev/null 2>&1 &").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func start2() {
	out, err := exec.Command("sh", "-c", "mpg321 -l 0 ./music/***.mp3").Output()
	if err != nil {
		fmt.Println("******", err)
	}
	fmt.Println(string(out))
}

func stop() {
	out, err := exec.Command("killall", "mpg321").Output()
	if err != nil {
		fmt.Println("-------", err)
	}
	fmt.Println(string(out))
}

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_CHAN_SEACRET"),
		os.Getenv("LINE_APP_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// set up alerm duration
				case *linebot.TextMessage:

					awaketime, err := strconv.Atoi(message.Text)
					if err != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("時間を入力してください\n例) 10\n10時間後にアラームがなります")).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%d時間後に起こしますね！", awaketime))).Do(); err != nil {
							log.Print(err)
						}
						time.Sleep(time.Duration(awaketime) * time.Second)
						start()
					}

				// stop alerm
				case *linebot.StickerMessage:
					stop()
					start2()
					stop()

				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("時間を入力してください\n例) 10\n10時間後にアラームがなります")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
