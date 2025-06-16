package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
	"github.com/yeungon/discordbot/internal/config"
	db "github.com/yeungon/discordbot/internal/pg"
)

func Boot() {
	dg, err := discordgo.New("Bot " + config.Token())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dbConn := DatabaseConnect()
	defer dbConn.Close()

	// ---------------setting config for the App------------
	appConfig := config.NewApp(true, true)
	appConfig.Query = db.New(dbConn)
	appConfig.Debug = true

	// ---------------setting config for the App------------

	Handles(dg, appConfig)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Register commands
	guildID := os.Getenv("SERVER_ID")
	for _, cmd := range handle.Commands {
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
