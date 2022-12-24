package svc

import (
	"github.com/kevwan/chatbot/bot"
	"github.com/kevwan/chatbot/cli/ask/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	ChatBot *bot.ChatBot
	Tops    int
}

func NewServiceContext(c config.Config, bot *bot.ChatBot, tops int) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ChatBot: bot,
		Tops:    tops,
	}
}
