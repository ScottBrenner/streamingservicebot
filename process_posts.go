package main

import (
	"fmt"
	"regexp"

	"github.com/turnage/graw/reddit"
)

const replyHeader string = "I found this track on other streaming services!:\n\n"
const replyFooter string = `

----

^(I am a bot, and this action was performed automatically.) [^(Source)](https://github.com/ScottBrenner/streamingservicebot)`

func checkURL(p *reddit.Post) bool {
	var re = regexp.MustCompile(`youtu|youtube|soundcloud|spotify`)
	return re.MatchString(p.URL)
}

func processYouTube(p *reddit.Post) string {
	youtubeSearchResult := youtubeSearch(p.Title)
	return fmt.Sprintf("- [YouTube](%s)\n", youtubeSearchResult)
}

func generateReply(p *reddit.Post) string {
	var replyText string
	streamingServices := make(map[string]string)
	replyText = replyHeader
	streamingServices["YouTube"] = processYouTube(p)
	for _, result := range streamingServices {
		replyText += result
	}
	replyText += replyFooter
	return replyText
}
