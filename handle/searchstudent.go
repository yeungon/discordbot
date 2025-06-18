package handle

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/internal/config"
	db "github.com/yeungon/discordbot/internal/pg"
)

func SearchStudentHandler(appConfig *config.AppConfig) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		content := strings.TrimSpace(m.Content)
		lower := strings.ToLower(content)
		parts := strings.Fields(lower)
		if len(parts) == 0 || len(parts) > 10 {
			return
		}

		command := parts[0]
		arg := ""
		if len(parts) > 1 {
			arg = strings.Join(parts[1:], " ")
		}

		if command == "search" || command == "s" {
			//s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are running find command, argument is:  %v", arg))
			students, err := searchStudentsByKeyword(appConfig.Query, arg)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					fmt.Println("Student not found")
					s.ChannelMessageSend(m.ChannelID, fmt.Sprint("Không thông tin sinh viên liên quan trong hệ thống!"))
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ Có lỗi xảy ra, không thể truy xuất dữ liệu với mã sinh viên %v. Lỗi:  %v", arg, err))
					log.Println("Query error:", err)
				}
				return
			}

			err = sendStudentSearchResults(s, m.ChannelID, students)
			if err != nil {
				fmt.Println("err", err)
				s.ChannelMessageSend(m.ChannelID, "❌ Failed to send results.")
			}

		}

	}
}

func searchStudentsByKeyword(query *db.Queries, keyword_input string) ([]db.Student, error) {
	ctx := context.Background()
	students, err := query.SearchStudentsByPhrase(ctx, keyword_input)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return students, nil
}

func sendStudentSearchResults(s *discordgo.Session, channelID string, students []db.Student) error {
	if len(students) == 0 {
		_, err := s.ChannelMessageSend(channelID, "❌ No students found.")
		return err
	}

	const maxStudents = 60
	const chunkSize = 5

	if len(students) > maxStudents {
		students = students[:maxStudents]
	}

	for i := 0; i < len(students); i += chunkSize {
		end := i + chunkSize
		if end > len(students) {
			end = len(students)
		}

		var message strings.Builder
		message.WriteString(fmt.Sprintf("**🔍 Kết quả tìm kiếm (từ %d đến %d):**\n\n", i+1, end))

		for j, student := range students[i:end] {
			message.WriteString(fmt.Sprintf(
				"**%d. %s** (%s / %s)\n"+
					"👤 Code: `%s`, Gender: %s, Ethnic: %s\n"+
					"🎂 DOB: %s\n"+
					"🆔 ID: %s\n"+
					"📞 Phone: %s\n"+
					"📧 Email: %s\n"+
					"📍 Province: %s\n"+
					"🏠 Address: %s\n"+
					"📝 Notes: %s\n\n",
				i+j+1,
				safeString(student.Name),
				safeString(student.Class),
				safeString(student.ClassCode),
				safeString(student.StudentCode),
				safeString(student.Gender),
				safeString(student.Ethnic),
				safeString(student.DobFormat),
				safeString(student.NationalID),
				safeString(student.Phone),
				safeString(student.Email),
				safeString(student.Province),
				safeString(student.Address),
				safeString(student.Notes),
			))
		}

		content := message.String()
		if len(content) > 2000 {
			content = content[:1990] + "\n...🔻 Message truncated due to Discord limit."
		}

		_, err := s.ChannelMessageSend(channelID, content)
		if err != nil {
			return err
		}
	}

	return nil
}

func safeString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
