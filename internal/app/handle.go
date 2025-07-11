package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/handle"
	"github.com/yeungon/discordbot/internal/config"
	"github.com/yeungon/discordbot/menu"
)

func Handles(dg *discordgo.Session, appConfig *config.AppConfig) {
	dg.AddHandler(handle.MessageCreateHandler(appConfig))
	dg.AddHandler(handle.CheckStudentHandler(appConfig))
	dg.AddHandler(handle.SearchStudentHandler(appConfig))
	dg.AddHandler(handle.GetStudentHandler(appConfig))
	dg.AddHandler(menu.InfoMenuCreate)
	dg.AddHandler(menu.InfoInteractionHandler)
	dg.AddHandler(handle.SlashCommandHandler)
}
