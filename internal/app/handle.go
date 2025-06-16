package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
	"github.com/yeungon/discordbot/internal/config"
)

func Handles(dg *discordgo.Session, appConfig *config.AppConfig) {
	dg.AddHandler(handle.MessageCreateHandler(appConfig))
	dg.AddHandler(handle.CheckStudentHandler(appConfig))
	dg.AddHandler(handle.SearchStudentHandler(appConfig))
	dg.AddHandler(handle.MessageButtonCreate)
	dg.AddHandler(handle.InteractionHandler)
	dg.AddHandler(handle.SlashCommandHandler)
}
