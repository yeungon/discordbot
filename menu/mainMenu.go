package menu

import "github.com/bwmarrin/discordgo"

var (
	longText = `At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!`
)

func MessageButtonCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "menu" || m.Content == "help" {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "Các tác vụ hiện có của bot:",
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
							Style:    discordgo.SuccessButton,
							CustomID: "xem_tkb",
						},
						discordgo.Button{
							Label:    "Thông tin hữu ích",
							Style:    discordgo.PrimaryButton,
							CustomID: "xem_diem_sv",
						},
						discordgo.Button{
							Label:    "PrimaryButton button ",
							Style:    discordgo.PrimaryButton,
							CustomID: "xem_diem_server_Abc",
						},
					},
				},
				// The message may have multiple actions rows.
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Discord Developers server",
							Style:    discordgo.DangerButton,
							CustomID: "xem_diem_server",
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							// Select menu, as other components, must have a customID, so we set it to this value.
							CustomID:    "select",
							Placeholder: "Choose your favorite programming language 👇",
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Go",
									// As with components, this things must have their own unique "id" to identify which is which.
									// In this case such id is Value field.
									Value: "go",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🦦",
									},
									// You can also make it a default option, but in this case we won't.
									Default:     false,
									Description: "Go programming language",
								},
								{
									Label: "JS",
									Value: "js",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🟨",
									},
									Description: "JavaScript programming language",
								},
								{
									Label: "Python",
									Value: "py",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🐍",
									},
									Description: "Python programming language",
								},
							},
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
			// Respond with a message + a new menu
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    "📘 Vui lòng chọn học kỳ để xem điểm:",
					Components: []discordgo.MessageComponent{semesterMenu},
				},
			})
		case "xem_tkb":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "📅 Bạn đã chọn **Xem thời khóa biểu**! \n " + longText,
				},
			})
		}
	}
}

var semesterMenu = discordgo.ActionsRow{
	Components: []discordgo.MessageComponent{
		discordgo.SelectMenu{
			CustomID:    "select_semester",
			Placeholder: "📘 Chọn học kỳ để xem điểm",
			Options: []discordgo.SelectMenuOption{
				{
					Label:       "Học kỳ 241",
					Value:       "241",
					Description: "Xem điểm học kỳ 241",
				},
				{
					Label:       "Học kỳ 242",
					Value:       "242",
					Description: "Xem điểm học kỳ 242",
				},
			},
		},
	},
}
