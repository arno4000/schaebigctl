package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var (
	Token string
)

func main() {

	//todo clean config
	Token = "token_place_holder"

	dg, err := discordgo.New(Token)

	if err != nil {
		logrus.Errorln(err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		fmt.Printf("User: %s, Message: %s, Channel ID: %s\n", m.Author.Username, m.Content, m.ChannelID)
		if m.Author.Username != "igdukekl" {
			//todo meaningful messages
			_, err := s.ChannelMessageSend(m.ChannelID, "kekw")
			if err != nil {
				logrus.Errorln(err)
			}
		}
	})

	err = dg.Open()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGILL, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
