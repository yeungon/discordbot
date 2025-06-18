package menu

import "github.com/bwmarrin/discordgo"

func InfoMenuCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "info" || m.Content == "Info" {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "Các tác vụ hiện có của bot:\n",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Quy định làm TL",
							Style:    discordgo.PrimaryButton,
							CustomID: "info_quydinh_lam_tl",
						},
						discordgo.Button{
							Label:    "Quy định làm KL",
							Style:    discordgo.PrimaryButton,
							CustomID: "info_quydinh_lam_kt",
						},
						discordgo.Button{
							Label:    "Trích dẫn TLTK",
							Style:    discordgo.PrimaryButton,
							CustomID: "info_trichdan_apa",
						},
						discordgo.Button{
							Label:    "Sổ tay sinh viên",
							Style:    discordgo.PrimaryButton,
							CustomID: "info_sotay_sinhvien",
						},
					},
				},
				// The message may have multiple actions rows.
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Sơ đồ trường",
							Style:    discordgo.SuccessButton,
							CustomID: "info_sodo_truong",
						},
						discordgo.Button{
							Label:    "List Khóa luận",
							Style:    discordgo.SuccessButton,
							CustomID: "xem_diem_server_abcd",
						},
						discordgo.Button{
							Label:    "Mật khẩu Wifi",
							Style:    discordgo.SuccessButton,
							CustomID: "xem_diem_server_xyz",
						},
						discordgo.Button{
							Label:    "Kế hoạch năm học",
							Style:    discordgo.SuccessButton,
							CustomID: "xem_diem_server_abc",
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Quy định Ngoại ngữ",
							Style:    discordgo.SecondaryButton,
							CustomID: "info_quydinh_ngoaingu",
						},
						discordgo.Button{
							Label:    "Kế hoạch năm học",
							Style:    discordgo.SecondaryButton,
							CustomID: "kehoach_namhoc",
						},
						discordgo.Button{
							Label:    "Xét tốt nghiệp",
							Style:    discordgo.SecondaryButton,
							CustomID: "dieukien_xet_tn",
						},
						discordgo.Button{
							Label:    "Mật khẩu Wifi",
							Style:    discordgo.SecondaryButton,
							CustomID: "info_matkhau_wifi",
						},
					},
				},
			},
		})
	}
}

func InfoInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		case "info_quydinh_lam_tl":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "📘 **Quy định làm tiểu luận: **.\n" + dieukien_tl,
				},
			})
		case "info_matkhau_wifi":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Trường: dhsph19572010\n Khoa: TU16051996",
				},
			})

		}

	}
}
