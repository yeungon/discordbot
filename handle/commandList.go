package handle

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Replies with a personalized greeting.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Your name",
					Required:    true,
				},
			},
		},
		{
			Name:        "ping",
			Description: "Ping với command thay vì message.",
		},

		{
			Name:        "pong",
			Description: "Pong với command thay vì message.",
		},
	}
)
