package abs

import (
	"encoding/json"
	"strconv"
	"strings"
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
	Latest string
	Oldest string
	Limit  int
}

type Message struct {
	ClientMsgID string `json:"client_msg_id,omitempty"`
	Text        string `json:"text"`
	Ts          string `json:"ts"`
}

func NewHistoryParameters() HistoryParameters {
	return HistoryParameters{
		Latest: DEFAULT_HISTORY_LATEST,
		Oldest: DEFAULT_HISTORY_OLDEST,
		Limit:  DEFAULT_HISTORY_Limit,
	}
}

func (s *SlackClient) GetMessages() []Message {
	// url 파라미터 설정
	historyParameters := NewHistoryParameters()
	var historyParametersMap map[string]string
	data, _ := json.Marshal(historyParameters)
	json.Unmarshal(data, &historyParametersMap)

	// Slack API 통신
	var history History
	s.NewAPI("https://slack.com/api/conversations.history", historyParametersMap, &history)
	return history.Messages
}

func (s *SlackClient) FilterMessages(messages []Message) []Message {
	mentionFilter := "<@" + s.botId + "> "
	var messagesFiltered []Message
	for _, message := range messages {
		// 멘션이 되었고 사용자가 친 채팅만 필터링
		if strings.Index(message.Text, mentionFilter) == 0 && len(message.ClientMsgID) > 0 {
			// 멘션 텍스트 제거
			message.Text = strings.Replace(message.Text, mentionFilter, "", 1)
			messagesFiltered = append(messagesFiltered, message)
		}
	}

	return messagesFiltered
}

func (s *SlackClient) ConvertToPayment(messagesFiltered []Message) []Payment {
	var payments []Payment
	for _, message := range messagesFiltered {
		txtSlice := strings.Split(message.Text, ";")
		if len(txtSlice) >= 6 {
			price, _ := strconv.Atoi(strings.Trim(txtSlice[4], " "))
			monthlyInstallment, _ := strconv.Atoi(strings.Trim(txtSlice[5], " "))
			payments = append(payments, Payment{
				Date:               strings.Trim(txtSlice[0], " "),
				Method:             strings.Trim(txtSlice[1], " "),
				Category:           strings.Trim(txtSlice[2], " "),
				Name:               strings.Trim(txtSlice[3], " "),
				Price:              price,
				MonthlyInstallment: monthlyInstallment,
			})
		}
	}
	return payments
}
