package main

import (
	"github.com/arno4000/schaebigctl/pkg/bot"
	"github.com/arno4000/schaebigctl/pkg/config"
	"github.com/spf13/viper"
)

func main() {

	config.InitConfig()
	bot.StartBot(viper.GetString("discordToken"))
}
