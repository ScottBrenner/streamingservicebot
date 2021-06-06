package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

const replyHeader string = "I found this track on other streaming services!:\n\n"
const replyFooter string = `

----

^(I am a bot, and this action was performed automatically.) [^(Source)](https://github.com/ScottBrenner/streamingservicebot)`

type streamingServiceBot struct {
	bot reddit.Bot
}

func (r *streamingServiceBot) Post(p *reddit.Post) error {
	var re = regexp.MustCompile(`youtu|youtube|soundcloud|spotify`)
	if re.MatchString(p.URL) {
		replyText := replyHeader
		streamingServices := make(map[string]string)
		fmt.Printf("%s posted \"%s\"!", p.Author, p.Title)
		youtubeSearchResult := youtubeSearch(p.Title)
		streamingServices["YouTube"] = fmt.Sprintf("- [YouTube](%s)\n", youtubeSearchResult)
		for _, result := range streamingServices {
			replyText += result
		}
		replyText += replyFooter
		return r.bot.Reply(
			p.Name,
			replyText,
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
		cfg := graw.Config{Subreddits: []string{"streamingservicebot"}}
		handler := &streamingServiceBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
	}
}
