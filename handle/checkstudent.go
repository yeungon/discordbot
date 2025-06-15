package handle

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CheckStudent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.TrimSpace(m.Content)
	lower := strings.ToLower(content)
	parts := strings.Fields(lower)
	if len(parts) == 0 || len(parts) != 2 {
		return
	}

	command := parts[0]
	arg := ""
	if len(parts) > 1 {
		arg = parts[1]
	}

	if command == "check" {
		fmt.Printf("command %s, arg: %s\n", command, arg)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are running check command, argument is:  %v", arg))
	}
}
