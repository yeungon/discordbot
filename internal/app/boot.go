package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
	"github.com/yeungon/discordbot/internal/config"
	db "github.com/yeungon/discordbot/internal/pg"
	"github.com/yeungon/discordbot/pkg/logging"
)

func Boot() {

	logging.Log()
	// When shutting down the application, ensure the log is closed properly
	defer logging.CloseLog()

	slog.Info("🟢 Bot is starting...")

	//---------------Run config first to get the environment variables------------
	config.New()
	cfg := config.Get()
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//---------------Connect to database------------
	dbConn := DatabaseConnect(cfg.PostgresURL)
	defer dbConn.Close()

	// ---------------setting config for the App------------
	appConfig := config.NewApp(true, true)
	appConfig.Query = db.New(dbConn)
	appConfig.Debug = false

	// ---------------setting config for the App------------

	Handles(dg, appConfig)

	// Enable necessary intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

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
