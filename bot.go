package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type streamingServiceBot struct {
	bot reddit.Bot
}

func (r *streamingServiceBot) Post(p *reddit.Post) error {
	if checkURL(p) {
		return r.bot.Reply(
			p.Name,
			generateReply(p),
		)
	}
	return nil
}

func main() {
	cfg := reddit.BotConfig{
		Agent: "graw:streamingservicebot:0.0.1 by /u/scottstimo",
		App: reddit.App{
			ID:       os.Getenv("SSB_REDDIT_ID"),
			Secret:   os.Getenv("SSB_REDDIT_SECRET"),
			Username: os.Getenv("SSB_REDDIT_USERNAME"),
			Password: os.Getenv("SSB_REDDIT_PASSWORD"),
		},
		Rate:   0,
		Client: &http.Client{},
	}
	if bot, err := reddit.NewBot(cfg); err != nil {
		fmt.Println("Failed to create bot handle: ", err)
	} else {
		cfg := graw.Config{Subreddits: strings.Split(os.Getenv("SSB_REDDIT_SUBREDDITS"), "+")}
		handler := &streamingServiceBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
	}
}
