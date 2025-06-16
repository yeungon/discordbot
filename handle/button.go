package handle

import (
	"github.com/bwmarrin/discordgo"
)

// func MessageButtonCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == s.State.User.ID {
// 		return
// 	}

// 	if m.Content == "!button" {
// 		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
// 			Content: "Click the button below:",
// 			Components: []discordgo.MessageComponent{
// 				discordgo.ActionsRow{
// 					Components: []discordgo.MessageComponent{
// 						discordgo.Button{
// 							Label:    "Xem điểm sinh viên!",
// 							Style:    discordgo.PrimaryButton,
// 							CustomID: "click_me_button",
// 						},
// 					},
// 				},
// 			},
// 		})
// 	}
// }

// func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if i.Type == discordgo.InteractionMessageComponent {
// 		switch i.MessageComponentData().CustomID {
// 		case "click_me_button":
// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Content: "🎉 You clicked the button!",
// 				},
// 			})
// 		}
// 	}
// }

func MessageButtonCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!button" {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "Click một trong các nút bên dưới:",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Xem điểm sinh viên",
							Style:    discordgo.PrimaryButton,
							CustomID: "xem_diem",
						},
						discordgo.Button{
							Label:    "Xem thời khóa biểu",
							Style:    discordgo.SecondaryButton,
							CustomID: "xem_tkb",
						},
					},
				},
			},
		})
	}
}

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent {
		switch i.MessageComponentData().CustomID {
		case "xem_diem":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "📊 Bạn đã chọn **Xem điểm sinh viên**!",
				},
			})
		case "xem_tkb":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "📅 Bạn đã chọn **Xem thời khóa biểu**!",
				},
			})
		}
	}
}
