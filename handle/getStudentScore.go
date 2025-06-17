package handle

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/yeungon/discordbot/internal/config"
)

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
func StudentCheckFetch(s *discordgo.Session, channelID string, studentID string) error {
	baseurl := config.Get().SPH_URL_ENDPOINT
	key := generateHash()
	endpoint := fmt.Sprintf("%s?key=%s&id=%s", baseurl, key, studentID)

	data, err := fetchData(endpoint)
	if err != nil {
		log.Printf("❌ Error fetching data: %v", err)
		s.ChannelMessageSend(channelID, fmt.Sprintf("❌ Có lỗi khi fetch dữ liệu: %v", err))
		return nil
	}

	// 🧑‍🎓 Thông tin sinh viên
	studentInfo := fmt.Sprintf(
		"📄 **Thông tin sinh viên**\n"+
			"• 🆔 Mã số: `%s`\n"+
			"• 👤 Họ và tên: **%s**\n"+
			"• 🏫 Lớp: `%s`\n"+
			"• 📞 SĐT: `%s`\n",
		data.Info.FMasv,
		data.Info.FHoten,
		data.Info.FLop,
		data.Info.FPhone,
	)

	if _, err := s.ChannelMessageSend(channelID, studentInfo); err != nil {
		return err
	}

	// 📘 Các môn có điểm tiểu luận
	var hasTL strings.Builder
	for _, subject := range data.Diem {
		if strings.TrimSpace(subject.Diemtl) != "" {
			hasTL.WriteString(fmt.Sprintf("• `%s` **%s** — Điểm TL: `%s`\n",
				subject.Mamh,
				subject.Tennh,
				subject.Diemtl,
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
	for _, diem := range data.Diem {
		// Normalize scores
		qt := formatScore(diem.Diemqt, "_")
		thi := formatScore(diem.Diemthi, "_")
		diemFinal := formatScore(diem.Diem, "_")

		// Add semester header only if new
		var semesterHeader string
		if diem.Hk != lastSemester {
			semesterHeader = fmt.Sprintf("📘 **Học kỳ %s**\n", diem.Hk)
			lastSemester = diem.Hk
		}

		entry := fmt.Sprintf(
			"%s• Môn: **%s** (`%s`)\n"+
				"  • ĐVHT: `%s`, Điểm: `%s`, QT: `%s`, Thi: `%s`\n\n",
			semesterHeader,
			diem.Tennh,
			diem.Mamh,
			diem.Dvht,
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

type ResponseData struct {
	Info struct {
		FMasv  string `json:"f_masv"`
		FHoten string `json:"f_hoten"`
		FLop   string `json:"f_lop"`
		FPhone string `json:"f_phone"`
	} `json:"info"`
	Diem []struct {
		Hk      string `json:"HK"`
		Mamh    string `json:"MAMH"`
		Diem    string `json:"DIEM"`
		Diemqt  string `json:"DIEMQT"`
		Diemthi string `json:"DIEMTHI"`
		Diemtl  string `json:"DIEMTL"`
		Tennh   string `json:"tennh"`
		Dvht    string `json:"dvht"`
	} `json:"diem"`
}

func fetchData(url string) (*ResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var data ResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &data, nil
}

func generateHash() string {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc
	prefix := config.Get().SECRET_FIRST
	suffix := config.Get().SECRET_SECOND
	currentTime := time.Now().In(loc)
	dateString := currentTime.Format("212006") // dmyyyy format
	input := prefix + dateString + suffix
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
