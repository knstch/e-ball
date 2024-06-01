package router

import (
	"github.com/go-telegram/bot"
	"github.com/knstch/e-ball/internal/app/handlers"
)

func NewRouter(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Greetings)
}
