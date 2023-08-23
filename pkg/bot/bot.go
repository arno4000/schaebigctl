package bot

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arno4000/schaebigctl/pkg/ai"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartBot(token string) {
	logrus.Infoln("Started schaebigctl bot")
	dg, err := discordgo.New(token)

	if err != nil {
		logrus.Errorln(err)
	}

	dg.AddHandler(func(s *discordgo.Session, t *discordgo.ThreadCreate) {

		if t.NewlyCreated {
			prompts := []string{
				`Your persona:You are a frustrated tech support worker. Explain to the caller why the topic they are asking about is shabby, but replace shabby with schaebig in your sentences. Your sentence is not allowed to include the words "language model", "OpenAI", "AI" or anything similar. Pretend you are a human. In addition do not start your answer with "it is schaebig",  "because" or anything similar. Do not offer to continue to talk with the person.  Include a sentence where you make fun of the caller and insert :joy: at the end. here's the question:`,
				`Your persona:You are a professional tech support worker. Explain to the caller how they can solve their question. Your sentence is not allowed to include the words "language model", "OpenAI", "AI" or anything similar. Pretend you are a human. Do not offer to continue to talk with the person. here's the question:`,
			}
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randomNumber := r.Intn(2)
			prompt := prompts[randomNumber]
			timeout := viper.GetInt("maxTimeoutMinutes")
			var waitDuration time.Duration
			if timeout > 0 {
				waitDuration = time.Duration(rand.Intn(viper.GetInt("maxTimeoutMinutes")))
				logrus.Infof("Posting message to thread %s in %d minutes", t.Name, waitDuration)
				time.Sleep(waitDuration * time.Minute)
			} else {
				logrus.Infof("Posting message to thread %s now", t.Name)
			}

			threadMessage, err := GetThreadMessage(t.ID, token)
			if err != nil {
				logrus.Errorln(err)
			}

			answer, err := ai.GenAIResponse(prompt, threadMessage)
			if err != nil {
				logrus.Errorln(err)
			}
			logrus.Infof("Posted message to thread %s, ID: %s, took %d minutes", t.Name, t.ID, waitDuration)
			_, err = s.ChannelMessageSend(t.ID, answer)
			if err != nil {
				logrus.Errorln(err)
			}
		}

	})

	err = dg.Open()
	if err != nil {
		logrus.Errorln(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGILL, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
	logrus.Infoln("Exiting schaebigctl bot. Bye. Stay schaebig!")

}

func GetThreadMessage(threadID string, token string) (string, error) {
	var messages ThreadMessage
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/channels/"+threadID+"/messages", nil)
	if err != nil {
		logrus.Errorln(err)
	}
	req.Header.Set("authority", "discord.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "de-DE,de;q=0.8")
	req.Header.Set("authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln(err)
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		logrus.Errorln(err)
	}
	firstMessage := messages[len(messages)-1].Content
	return firstMessage, err
}
