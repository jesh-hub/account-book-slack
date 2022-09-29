package test

import (
	"abs/model"
	"encoding/csv"
	"fmt"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	foods     [][]string
	shoppings [][]string
)

func TestRandomPayment(t *testing.T) {
	var paymentAdds []*model.PaymentAdd
	var name string
	var category string
	var err error

	foods, err = getFoodData()
	assert.Equal(t, err, nil)

	shoppings, err = getShoppingData()
	assert.Equal(t, err, nil)

	year := 2022
	months := []int{6, 7, 8, 9, 10, 11, 12}
	sampleCount := 200
	for _, month := range months {
		for i := 0; i < sampleCount; i++ {
			if i%2 == 0 {
				name, category = getRadomDataShopping()
			} else {
				name, category = getRadomDataFood()
			}
			if err != nil {
				log.Println(err)
				break
			}

			paymentAdds = append(paymentAdds, &model.PaymentAdd{
				Date:               primitive.NewDateTimeFromTime(getRandomDate(year, month)),
				Name:               name,
				Category:           category,
				Price:              getRandomPrice(5000, 100000),
				MonthlyInstallment: 0,
				PaymentMethodId:    primitive.ObjectID{},
				RegUserId:          "",
			})
		}
	}

	fmt.Printf("%# v\n", pretty.Formatter(paymentAdds))
}

func getRandomDate(year int, month int) time.Time {
	min := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func getRandomPrice(min int, max int) int {
	return (rand.Intn(max-min+1) + min) / 100 * 100
}

func getRadomDataShopping() (string, string) {
	min := 1
	max := len(shoppings)
	randomIndex := rand.Intn(max-min+1) + min
	return shoppings[randomIndex][0], "쇼핑"
}

func getRadomDataFood() (string, string) {
	min := 1
	max := len(foods)
	randomIndex := rand.Intn(max-min+1) + min
	return foods[randomIndex][0], foods[randomIndex][1]
}

func getFoodData() ([][]string, error) {
	file, err := os.Open("/Users/choshsh/Downloads/food_data.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	foods, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func getShoppingData() ([][]string, error) {
	file, err := os.Open("/Users/choshsh/Downloads/shopping_data.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 1
	shoppings, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return shoppings, nil
}
