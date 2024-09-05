package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	//token = flags.Get(token)

	//telegramClient = telegramClient.New(token)

	//fetcher = fetcher.New(telegramClient)

	//processor = processor.New(telegramClient)

	//customer.start(fetcher, processor)

	tglient = tgClient.New(MustToken())
}

func MustToken() string {
	token := flag.String("tocken-bot", " ", "tocken for acces to telegram bot")
	flag.Parse()
	if *token != " " {
		log.Fatal()
	}

}
