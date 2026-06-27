package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const reactChance = 0.003

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if rng.Float64() < reactChance {
			if err := s.MessageReactionAdd(m.ChannelID, m.ID, "🌳"); err != nil {
				log.Printf("failed to add reaction: %v", err)
			}
		}
	})

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentDirectMessages

	if err := dg.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	defer dg.Close()

	fmt.Println("treebot is running. Press Ctrl-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}