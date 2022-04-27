package slack

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cloudflare/ahocorasick"
)

const (
	DEFAULT_HISTORY_LATEST = ""
	DEFAULT_HISTORY_OLDEST = "0"
	DEFAULT_HISTORY_Limit  = 100
)

type History struct {
	Messages []Message `json:"messages"`
}

type HistoryParameters struct {
	Latest string `json:"latest"`
	Oldest string `json:"oldest"`
	Limit  int    `json:"limit"`
}

type Message struct {
	ClientMsgID string `json:"client_msg_id,omitempty"`
	Text        string `json:"text"`
	Ts          string `json:"ts"`
}

type MessageParameters struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (mp *MessageParameters) StartAsTime() time.Time {
	startTime, err := time.Parse("2006-01", mp.Start)
	errorHandler(err)
	return startTime
}

func (mp *MessageParameters) EndAsTime() time.Time {
	endTime, err := time.Parse("2006-01", mp.End)
	errorHandler(err)
	return endTime
}

func NewHistoryParameters() HistoryParameters {
	return HistoryParameters{
		Latest: DEFAULT_HISTORY_LATEST,
		Oldest: DEFAULT_HISTORY_OLDEST,
		Limit:  DEFAULT_HISTORY_Limit,
	}
}

func (s *SlackClient) GetMessages(messageParameters MessageParameters) []Message {
	// url 파라미터 설정
	historyParameters := NewHistoryParameters()
	if len(messageParameters.Start) > 0 && len(messageParameters.End) > 0 {
		historyParameters.Oldest = fmt.Sprintf("%v", messageParameters.StartAsTime().Unix())
		// 늦게 입력된 채팅 크롤링을 위해서 end + 1달 처리
		historyParameters.Latest = fmt.Sprintf("%v", messageParameters.EndAsTime().AddDate(0, 1, 0).Unix())
	}

	// Slack API 통신
	var history History
	s.NewAPI("https://slack.com/api/conversations.history", historyParameters, &history)
	return history.Messages
}

func (s *SlackClient) FilterMessages(messages []Message) []Message {
	mentionFilter := "<@" + s.botId + "> "
	m := ahocorasick.NewStringMatcher([]string{mentionFilter})
	var messagesFiltered []Message
	for _, message := range messages {
		// bot이 멘션된 채팅만 필터링
		hits := m.Match([]byte(message.Text))
		if len(hits) == 1 {
			// 멘션 텍스트 제거
			message.Text = strings.Replace(message.Text, mentionFilter, "", 1)
			messagesFiltered = append(messagesFiltered, message)
		}
	}

	return messagesFiltered
}

func (s *SlackClient) ConvertToPayment(messagesFiltered []Message, messageParameters MessageParameters) []Payment {
	trim := func(s string) string {
		return strings.Trim(s, " ")
	}

	dateFilter := func(date string) bool {
		dateTime, err := time.Parse("2006-01-02", date)
		errorHandler(err)

		if dateTime.Unix() >= messageParameters.StartAsTime().Unix() &&
			dateTime.Unix() < messageParameters.EndAsTime().AddDate(0, 1, 0).Unix() {
			return true
		} else {
			return false
		}
	}

	var payments []Payment
	for _, message := range messagesFiltered {
		txtSlice := strings.Split(message.Text, ";")
		if len(txtSlice) >= 6 {
			date := trim(txtSlice[0])
			if dateFilter(date) {
				price, _ := strconv.Atoi(trim(txtSlice[4]))
				monthlyInstallment, _ := strconv.Atoi(trim(txtSlice[5]))
				payments = append(payments, Payment{
					Date:               date,
					Method:             trim(txtSlice[1]),
					Category:           trim(txtSlice[2]),
					Name:               trim(txtSlice[3]),
					Price:              price,
					MonthlyInstallment: monthlyInstallment,
				})
			}
		}
	}
	return payments
}
