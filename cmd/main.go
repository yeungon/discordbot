package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
)

func main() {
	Token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handle.MessageCreate)
	dg.AddHandler(handle.SlashCommandHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Register commands
	guildID := os.Getenv("SERVER_ID")
	for _, cmd := range commands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, cmd)
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v\n", cmd.Name, err)
		}
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
