package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/internal/config"
	"github.com/yeungon/discordbot/pkg/helpers"
)

type ResponseData struct {
	Info json.RawMessage `json:"info"`
	Diem []DiemData      `json:"diem"`
}

type StudentInfo struct {
	FMasv  string `json:"f_masv"`
	FHoten string `json:"f_hoten"`
	FLop   string `json:"f_lop"`
	FPhone string `json:"f_phone"`
}

type DiemData struct {
	HK      string `json:"HK"`
	MAMH    string `json:"MAMH"`
	DIEM    string `json:"DIEM"`
	DIEMQT  string `json:"DIEMQT"`
	DIEMTHI string `json:"DIEMTHI"`
	DIEMTL  string `json:"DIEMTL"`
	TenNH   string `json:"tennh"`
	DVHT    string `json:"dvht"`
}

// type ResponseData struct {
// 	Info struct {
// 		FMasv  string `json:"f_masv"`
// 		FHoten string `json:"f_hoten"`
// 		FLop   string `json:"f_lop"`
// 		FPhone string `json:"f_phone"`
// 	} `json:"info"`
// 	Diem []struct {
// 		Hk      string `json:"HK"`
// 		Mamh    string `json:"MAMH"`
// 		Diem    string `json:"DIEM"`
// 		Diemqt  string `json:"DIEMQT"`
// 		Diemthi string `json:"DIEMTHI"`
// 		Diemtl  string `json:"DIEMTL"`
// 		Tennh   string `json:"tennh"`
// 		Dvht    string `json:"dvht"`
// 	} `json:"diem"`
// }

func GetStudentHandler(appConfig *config.AppConfig) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		content := strings.TrimSpace(m.Content)
		parts := strings.Fields(content)
		if len(parts) == 0 || len(parts) != 2 {
			return
		}

		command := strings.ToLower(parts[0])
		student_code := ""
		if len(parts) > 1 {
			student_code = strings.Join(parts[1:], " ")
		}

		student_code = strings.ToUpper(student_code)

		if command == "get" || command == "g" {
			StudentCheckFetch(s, m.ChannelID, student_code)
		}

	}
}

func blankIfEmpty(s string) string {
	if s == "" {
		return "(chưa có)"
	}
	return s
}

func StudentCheckFetch(s *discordgo.Session, channelID string, studentID string) error {
	baseurl := config.Get().SphURLEndpoint
	key := helpers.GenerateHash()
	endpoint := fmt.Sprintf("%s?key=%s&id=%s", baseurl, key, studentID)
	student_info, diem, err := fetchData(endpoint)

	if err != nil {
		log.Printf("❌ Error fetching data: %v", err)
		s.ChannelMessageSend(channelID, fmt.Sprintf("❌ Có lỗi khi fetch dữ liệu: %v", err))
		return nil
	}

	studentInfo := fmt.Sprintf(
		"📄 **Thông tin sinh viên**\n"+
			"• 🆔 Mã số: `%s`\n"+
			"• 👤 Họ và tên: **%s**\n"+
			"• 🏫 Lớp: `%s`\n"+
			"• 📞 SĐT: `%s`\n",
		blankIfEmpty(student_info.FMasv),
		blankIfEmpty(student_info.FHoten),
		blankIfEmpty(student_info.FLop),
		blankIfEmpty(student_info.FPhone),
	)

	if _, err := s.ChannelMessageSend(channelID, studentInfo); err != nil {
		return err
	}

	// 📘 Các môn có điểm tiểu luận
	var hasTL strings.Builder
	for _, subject := range diem {
		if strings.TrimSpace(subject.DIEMTL) != "" {
			hasTL.WriteString(fmt.Sprintf("• `%s` **%s** — Điểm TL: `%s`\n",
				subject.MAMH,
				subject.TenNH,
				subject.DIEMTL,
			))
		}
	}

	if hasTL.Len() > 0 {
		header := "📝 **Các môn có làm tiểu luận:**\n"
		if _, err := s.ChannelMessageSend(channelID, header+hasTL.String()); err != nil {
			return err
		}
	}

	// 📊 Kết quả học tập theo học kỳ
	const maxMessageLength = 2000
	var currentMsg strings.Builder
	currentMsg.WriteString("📊 **Kết quả học tập**\n\n")

	lastSemester := ""
	for _, diem := range diem {
		// Normalize scores
		qt := formatScore(diem.DIEMQT, "_")
		thi := formatScore(diem.DIEMTHI, "_")
		diemFinal := formatScore(diem.DIEM, "_")

		// Add semester header only if new
		var semesterHeader string
		if diem.HK != lastSemester {
			semesterHeader = fmt.Sprintf("📘 **Học kỳ %s**\n", diem.HK)
			lastSemester = diem.HK
		}

		entry := fmt.Sprintf(
			"%s• Môn: **%s** (`%s`)\n"+
				"  • ĐVHT: `%s`, Điểm: `%s`, QT: `%s`, Thi: `%s`\n\n",
			semesterHeader,
			diem.TenNH,
			diem.MAMH,
			diem.DVHT,
			diemFinal,
			qt,
			thi,
		)

		// Send if nearing limit
		if currentMsg.Len()+len(entry) > maxMessageLength {
			if _, err := s.ChannelMessageSend(channelID, currentMsg.String()); err != nil {
				return err
			}
			currentMsg.Reset()
		}

		currentMsg.WriteString(entry)
	}

	// Final flush
	if currentMsg.Len() > 0 {
		if _, err := s.ChannelMessageSend(channelID, currentMsg.String()); err != nil {
			return err
		}
	}

	return nil
}

func formatScore(val string, fallback string) string {
	val = strings.TrimSpace(val)
	if val == "" {
		return fallback
	}
	return val
}

func fetchData(endpoint string) (*StudentInfo, []DiemData, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	var raw ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Handle the "info" field which might be an object or false
	var info StudentInfo
	if string(raw.Info) != "false" {
		if err := json.Unmarshal(raw.Info, &info); err != nil {
			return nil, nil, fmt.Errorf("failed to parse student info: %w", err)
		}
	} else {
		// leave info fields empty
		info = StudentInfo{}
	}

	return &info, raw.Diem, nil
}
