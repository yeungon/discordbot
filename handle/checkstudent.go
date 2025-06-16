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

func CheckStudentHandler(appConfig *config.AppConfig) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
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

		if len(parts[1]) != 10 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ Mã sinh viên không chính xác:  %v \n.✅ Mã sinh viên có 10 kí tự !", arg))
			return
		}

		if command == "check" {
			student_id_uppercase := strings.ToUpper(arg)
			err, studentData := CheckStudentModel(appConfig.Query, student_id_uppercase)

			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					fmt.Println("Student not found")
					s.ChannelMessageSend(m.ChannelID, fmt.Sprint("Không có mã sinh viên này trong hệ thống!"))
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ Có lỗi xảy ra, không thể truy xuất dữ liệu với mã sinh viên %v. Lỗi:  %v", arg, err))
					log.Println("Query error:", err)
				}
				return
			}
			s.ChannelMessageSend(m.ChannelID, FormatStudentInfo(*studentData))
		}
	}
}

func CheckStudentModel(query *db.Queries, student_id string) (error, *db.Student) {
	ctx := context.Background()
	student_code := sql.NullString{
		String: student_id,
		Valid:  true,
	}
	ListUser, err := query.GetStudentByStudentCode(ctx, student_code)
	if err != nil {
		log.Print("Error when querying user list", err)
		return err, nil
	}
	fmt.Println(&ListUser)
	return nil, &ListUser
}

func SafeString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "Không có dữ liệu"
}

func FormatStudentInfo(s db.Student) string {
	return fmt.Sprintf(
		`📄 **Thông tin sinh viên**
**👤 Họ và tên**: %s
**🎓 Mã sinh viên**: %s
**🧑 Giới tính**: %s
**🎂 Ngày sinh**: %s (định dạng: %s)
**🏫 Lớp**: %s (%s)
**🌍 Dân tộc**: %s
**🆔 Số CMND/CCCD**: %s
**📞 Số điện thoại**: %s
**✉️ Email**: %s
**🏠 Tỉnh/Thành phố**: %s
**📍 Địa chỉ**: %s
**📝 Ghi chú**: %s`,
		SafeString(s.Name),
		SafeString(s.StudentCode),
		SafeString(s.Gender),
		SafeString(s.Dob),
		SafeString(s.DobFormat),
		SafeString(s.Class),
		SafeString(s.ClassCode),
		SafeString(s.Ethnic),
		SafeString(s.NationalID),
		SafeString(s.Phone),
		SafeString(s.Email),
		SafeString(s.Province),
		SafeString(s.Address),
		SafeString(s.Notes),
	)
}
