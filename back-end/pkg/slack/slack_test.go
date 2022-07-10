package slack

import (
	"abs/database"
	"abs/service"
	"abs/util"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
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
		Start: "2022-06",
		End:   "2022-08",
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

func TestConvertSlackToMongo(t *testing.T) {
	parameters := &MessageParameters{
		Start: "2022-03",
		End:   "2022-08",
	}

	slackClient := NewSlackClient(slackToken, channelId, botId)
	messages := slackClient.GetMessages(parameters)
	messagesFiltered := slackClient.FilterMessages(messages)

	payments := slackClient.ConvertToPayment(messagesFiltered, parameters)

	database.Init()
	for _, payment := range payments {
		date, _ := time.Parse("2006-01-02", payment.Date)
		date = date.UTC()
		fmt.Println(primitive.DateTime(date.UnixMilli()).Time())
		data := &service.Payment{
			Date:               primitive.DateTime(date.UnixMilli()),
			Name:               payment.Name,
			Category:           payment.Category,
			Price:              payment.Price,
			MonthlyInstallment: payment.MonthlyInstallment,
			PaymentMethodId:    util.ConvertStringToObjectId("62cae104f3ff23947354f590"),
			GroupId:            util.ConvertStringToObjectId("62cae0f5f3ff23947354f58f"),
			RegUserId:          "jee824k@gmail.com",
		}
		service.AddPayment(data)
	}
	assert.NotEqual(t, len(payments), 0)
}
