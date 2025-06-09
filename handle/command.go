package handle

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "hello":
		options := i.ApplicationCommandData().Options
		name := "bạn"
		if len(options) > 0 {
			name = options[0].StringValue()
		}

		fmt.Println(options)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Xin chào, %s! 👋", name),
			},
		})
	case "ping":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong slash!",
			},
		})
	case "pong":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Ping slash!",
			},
		})
	default:
		// Optional: handle unknown commands or ignore
	}
}
