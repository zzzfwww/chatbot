package main

import (
	"flag"
	"fmt"
	"github.com/kevwan/chatbot/bot"
	"github.com/kevwan/chatbot/bot/adapters/logic"
	"github.com/kevwan/chatbot/bot/adapters/storage"
	"github.com/kevwan/chatbot/cli/ask/internal/config"
	"github.com/kevwan/chatbot/cli/ask/internal/handler"
	"github.com/kevwan/chatbot/cli/ask/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

var (
	configFile = flag.String("f", "etc/bot-api.yaml", "the config file")
	verbose    = flag.Bool("v", false, "verbose mode")
	storeFile  = flag.String("c", "corpus.gob", "the file to store corpora")
	tops       = flag.Int("t", 5, "the number of answers to return")
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	store, err := storage.NewSeparatedMemoryStorage(*storeFile)
	if err != nil {
		log.Fatal(err)
	}

	chatbot := &bot.ChatBot{
		LogicAdapter: logic.NewClosestMatch(store, *tops),
	}
	if *verbose {
		chatbot.LogicAdapter.SetVerbose()
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, chatbot, *tops)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
