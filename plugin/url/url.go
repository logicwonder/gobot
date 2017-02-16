package url

import (
	"fmt"
	"net/url"

	"github.com/gabeguz/gobot/bot"
	"github.com/thatguystone/swan"
)

type Url struct{}

func (p Url) Name() string {
	return "Url v1.0"
}

func (p Url) Execute(msg bot.Message, bot bot.Bot) error {

	if msg.From() == bot.FullName() {
		return nil
	}

	u, err := url.ParseRequestURI(msg.Body())
	if err != nil {
		return err
	}

	a, err := swan.FromURL(u.String())
	if err != nil {
		return (err)
	}

	// Respond with the article title
	fmt.Printf("Title: %v\n", a)
	bot.Send(a.Meta.Title)

	return nil
}