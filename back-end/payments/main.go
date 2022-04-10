package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	// 슬랙에서 채팅 내역 가져오기
	slackClient := NewSlackClient()
	messages := slackClient.GetMessages()

	// 채팅 내역을 []Payment로 데이터 가공
	var payments []Payment
	for _, message := range messages {
		if len(message.ClientMsgID) > 0 {
			payments = append(payments, Payment{
				Date: message.Text,
			})
		}
	}

	// 가공한 []Payment 리턴
	body, err := json.Marshal(payments)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "payments-handler",
		},
	}

	return resp, nil
}

type SlackHistory struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ClientMsgID string `json:"client_msg_id,omitempty"`
	Text        string `json:"text"`
	Ts          string `json:"ts"`
}

type Payment struct {
	Date               string `json:"date"`
	Name               string `json:"name"`
	Category           string `json:"category"`
	Method             string `json:"method"`
	Price              int    `json:"price"`
	MonthlyInstallment int    `json:"monthlyInstallment"`
}

type SlackClient struct {
	token      string
	channelId  string
	httpClient *http.Client
}

func NewSlackClient() *SlackClient {
	return &SlackClient{
		token:     os.Getenv("slackToken"),
		channelId: os.Getenv("channelId"),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *SlackClient) GetMessages() []Message {
	req, err := http.NewRequest("GET", "https://slack.com/api/conversations.history", nil)
	errorHandler(err)

	// 헤더
	req.Header.Add("Authorization", "Bearer "+s.token)

	// urlQeury
	q := req.URL.Query()
	q.Add("channel", s.channelId)
	req.URL.RawQuery = q.Encode()

	resp, err := s.httpClient.Do(req)
	errorHandler(err)

	defer resp.Body.Close()

	slackHistory := SlackHistory{}
	err = json.NewDecoder(resp.Body).Decode(&slackHistory)
	errorHandler(err)

	return slackHistory.Messages
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(Handler)
}
