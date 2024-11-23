package slackbot

import (
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/slack"
)

func main() {
	s := slack.NewRequestHandler()
	s.HTTPServe()
}
