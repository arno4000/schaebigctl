package ai

import (
	"fmt"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
)

func GenAIResponse(promt string, question string) (string, error) {
	chatGPTInput := parseMessage(promt, question)
	pw, err := playwright.Run()
	if err != nil {
		logrus.Errorln(err)
		return "schaebig 🤣", err
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		logrus.Errorln(err)
		return "schaebig 🤣", err
	}
	page, err := browser.NewPage()
	if err != nil {
		logrus.Errorln(err)
		return "schaebig 🤣", err
	}
	if _, err = page.Goto("https://chat.openai.com/"); err != nil {
		logrus.Errorln(err)
		return "schaebig 🤣", err
	}
	page.Click(`//*[@id="__next"]/div[1]/div[2]/div[1]/div/button[1]`)
	page.Type(`//*[@id="username"]`, "chatgpt_mail")
	time.Sleep(time.Second * 1)
	page.Press(`//*[@id="username"]`, "Enter")
	page.Type(`//*[@id="password"]`, "chatgpt_password")
	time.Sleep(time.Second * 1)
	page.Press(`//*[@id="password"]`, "Enter")
	time.Sleep(time.Second * 1)
	page.Click(`//*[@id="radix-:rf:"]/div[2]/div/div[4]/button/div`)
	page.Click(`//*[@id="prompt-textarea"]`)
	time.Sleep(time.Second * 1)
	page.Type(`//*[@id="prompt-textarea"]`, chatGPTInput)
	page.Press(`//*[@id="prompt-textarea"]`, "Enter")

	time.Sleep(time.Second * 15)
	response, err := page.Locator(`//*[@id="__next"]/div[1]/div[2]/div/main/div/div[1]/div/div/div/div[2]/div`).TextContent()
	if err != nil {
		logrus.Errorln(err)
		return "schaebig 🤣", err
	}
	if response == "" {
		logrus.Errorln("Failed to get message. Assuming bot popup came up. Sending message to maintainer")
	}
	response = strings.ReplaceAll(response, "ChatGPT1 / 1", "")
	err = page.Close()
	if err != nil {
		logrus.Errorln(err)
	}
	err = browser.Close()
	if err != nil {
		logrus.Errorln(err)
	}
	return response, err
}

func parseMessage(promt string, question string) string {
	return fmt.Sprintf("%s %s", promt, question)
}
