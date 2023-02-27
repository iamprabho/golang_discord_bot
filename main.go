package main

import (
	"fmt"

	"github.com/iamprabho/discord-ping/bot"
	"github.com/iamprabho/discord-ping/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}