package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New("7312826275:AAEb2dvmGW7miZ94zYgdls2QqG29y1nr_UU")
	if err != nil {
		log.Print(err)
	}

	b.Start(ctx)
}
