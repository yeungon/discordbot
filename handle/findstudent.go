package handle

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func FindStudent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.TrimSpace(m.Content)
	lower := strings.ToLower(content)
	parts := strings.Fields(lower)
	if len(parts) == 0 || len(parts) > 7 {
		return
	}

	command := parts[0]
	arg := ""
	if len(parts) > 1 {
		arg = strings.Join(parts[1:], " ")
	}

	fmt.Printf("command %s, arg: %s\n", command, arg)

	if command == "find" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are running find command, argument is:  %v", arg))
	}
}
