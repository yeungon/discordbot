package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
)

func Handles(dg *discordgo.Session) {
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handle.MessageCreate)
	dg.AddHandler(handle.CheckStudent)
	dg.AddHandler(handle.FindStudent)
	dg.AddHandler(handle.SlashCommandHandler)
}
