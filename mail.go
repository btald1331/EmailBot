package main

import (
	"flag"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/smtp"
	"time"
)

const (
	SERVICE = "emailbot"
	VERSION = "1.0.0"
)

func main() {

	logPath := flag.String("logpath", "/var/log/emailbot.log", "Path to output logs.")
	interval := flag.Int("interval", 48, "Send email every X hours.")
	stop := flag.Int("stop", 168, "Stop sending emails after X hours.")
	rotate := flag.Int("rotate", 1000, "Rotate file after X megabytes.")
	message := flag.String("message", "Hello,\n\nMy name is Brice Aldrich. This is a bot, beep bop boop. This message was delivered by a program I wrote in Golang. I can write code. I can dev. I know 5 spanish words. One being contratame. This will execute every 2 days until contact has been made. In the mean time, please check out the link to my resume.\n\nSource: https://github.com/btald1331/EmailBot.git\nResume: https://drive.google.com/file/d/112VKL35gvsuyPxtuGcfDV7p6Lvap9Ad0/view?usp=sharing\n\nBest Regards,\nBrice Aldrich\n(260) 582-9842\nbrice.aldrich@gmail.com", "Body of email.")
	flag.Parse()

	log.SetOutput(&lumberjack.Logger{
		Filename:   *logPath,
		MaxSize:    *rotate, // megabytes
		MaxBackups: 3,
		Compress:   true, // disabled by default
	})
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) //Remove default logging timestamp.

	send(*message) //Send initial email

	ticker := time.NewTicker(time.Hour * time.Duration(*interval)) //Start interval count
	go func() {
		for range ticker.C {
			send(*message)
		}
	}()
	time.Sleep(time.Hour * time.Duration(*stop)) //Stop sending emails after so many intervals
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func send(body string) {
	from := "<INSERT YOUR EMAIL>"
	pass := "<INSERT YOUR PASS>" //NOTE: This is a plain text password. Not secure.
	to := "<DESTINATION ADDRESS>"
	server := "smtp.gmail.com:587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hire me. I can dev.\n\n" +
		body

	err := smtp.SendMail(server,
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		logError(err)
		return
	}
	logSend(from, to, server, body)
}

func logError(err error) {
	log.Println(SERVICE + " " + VERSION + " " + giveMeZulu() + " " + "EMAIL.ERROR{" + "\"error\":\"" + err.Error() + "\"}")
}

func logSend(from, to, server, body string) {
	log.Println(SERVICE + " " + VERSION + " " + giveMeZulu() + " " + "EMAIL.SENT{ " + "\"from\":\"" + from + "\", \"to\":\"" + to + "\", \"server\":\"" + server + "\", \"body\":\"" + body + "\"}")
}

func giveMeZulu() string {
	t := time.Now()
	return t.Format("RFC3339")
}
