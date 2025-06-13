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
		handleHelloCommand(s, i)
	case "ping":
		handlePingCommand(s, i)
	case "pong":
		handlePongCommand(s, i)
	default:
		// Optional: handle unknown commands
	}
}

func handleHelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
}

func handlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "test Pong slash!",
		},
	})
}

func handlePongCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "hello là Ping slash!",
		},
	})
}
