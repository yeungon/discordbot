package handle

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	name1  = "Alice"
	age1   = 24
	city1  = "New York"
	job1   = "Engineer"
	score1 = 89.5

	name2  = "Bob"
	age2   = 30
	city2  = "London"
	job2   = "Designer"
	score2 = 92.0
)

var table string = fmt.Sprintf("```"+
	"| %-8s | %-3s | %-9s | %-10s | %-5s |\n"+
	"|----------|-----|-----------|------------|-------|\n"+
	"| %-8s | %-3d | %-9s | %-10s | %-5.1f |\n"+
	"| %-8s | %-3d | %-9s | %-10s | %-5.1f |\n"+

	"```",
	"Name", "Age", "City", "Job", "Score",
	name1, age1, city1, job1, score1,
	name2, age2, city2, job2, score2,
)

// Define responses to specific messages
var textTriggers = map[string]string{
	"status": "**Production** Production or Dev",
	"ping":   table,
	"pong":   "Có gì hay Ping!",
	"hi":     `Phát biểu tại họp báo, nhà báo Phùng Công Sưởng - Tổng biên tập báo Tiền Phong, Trưởng BTC Hoa hậu Việt Nam 2024 - khẳng định 6 tháng là chặng đường dài của sự nỗ lực không ngừng, của tinh thần bền bỉ, cầu tiến. Đây cũng là hành trình trưởng thành đầy xúc cảm của các thí sinh Hoa hậu Việt Nam 2024.`,
}

// Main message handler
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.TrimSpace(m.Content)
	lower := strings.ToLower(content)
	parts := strings.Fields(lower)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	arg := ""
	if len(parts) > 1 {
		arg = parts[1]
	}

	fmt.Printf("command %s, arg: %s", command, arg)

	// Check if there's a matching trigger
	if response, ok := textTriggers[content]; ok {
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
