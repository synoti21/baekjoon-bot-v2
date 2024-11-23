package discordbot

import (
	"log"

	"github.com/synoti21/baekjoon-slack-bot/internal/bots/discord"
)

func main() {
	h, err := discord.NewRequestHandler()
	if err != nil {
		log.Fatalf("Fail")
	}
	h.HTTPServe()
}
