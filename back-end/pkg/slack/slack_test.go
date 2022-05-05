package slack

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	slackToken string
	channelId  string
	botId      string
)

func init() {
	godotenv.Load()
	slackToken = os.Getenv("slackToken")
	channelId = os.Getenv("channelId")
	botId = os.Getenv("botId")
}

func TestHistory(t *testing.T) {
	parameters := &MessageParameters{
		Start: "2022-04",
		End:   "2022-05",
	}

	slackClient := NewSlackClient(slackToken, channelId, botId)
	messages := slackClient.GetMessages(parameters)
	messagesFiltered := slackClient.FilterMessages(messages)

	payments := slackClient.ConvertToPayment(messagesFiltered, parameters)

	fmt.Print("message size: ")
	fmt.Println(len(messagesFiltered))
	fmt.Print("payments size: ")
	fmt.Println(len(payments))
	fmt.Println(prettyPrint(payments))

	assert.NotEqual(t, len(payments), 0)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
