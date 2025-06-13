package handle

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Define responses to specific messages
var textTriggers = map[string]string{
	"ping": "Xin chào, tiếng Việt tốt chứ?!",
	"pong": "Có gì hay Ping!",
	"hi":   `Phát biểu tại họp báo, nhà báo Phùng Công Sưởng - Tổng biên tập báo Tiền Phong, Trưởng BTC Hoa hậu Việt Nam 2024 - khẳng định 6 tháng là chặng đường dài của sự nỗ lực không ngừng, của tinh thần bền bỉ, cầu tiến. Đây cũng là hành trình trưởng thành đầy xúc cảm của các thí sinh Hoa hậu Việt Nam 2024.`,
}

// Main message handler
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Normalize message
	content := strings.ToLower(strings.TrimSpace(m.Content))

	// Check if there's a matching trigger
	if response, ok := textTriggers[content]; ok {
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
