package test

import (
	"abs/model"
	"abs/service"
	"abs/util"
	"encoding/csv"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var paymentDocs []interface{}
	var err error

	foods, err = getFoodData()
	assert.Equal(t, err, nil)

	shoppings, err = getShoppingData()
	assert.Equal(t, err, nil)

	var name string
	var category string

	groups, err := service.FindGroupByEmail(groupFind)
	paymentMethods, err := service.FindPaymentMethodByGroupId(groups[0].ID.Hex())
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

			paymentDocs = append(paymentDocs, &model.Payment{
				Date:               primitive.NewDateTimeFromTime(util.GetRandomDate(year, month)),
				Name:               name,
				Category:           category,
				Price:              getRandomPrice(5000, 100000),
				MonthlyInstallment: 0,
				PaymentMethodId:    paymentMethods[0].ID,
				GroupId:            groups[0].ID,
				RegUserId:          "guest@wejesh.com",
				ModUserId:          "",
			})
		}
	}
	assert.Greater(t, len(paymentDocs), 0)

	fmt.Println("결제내역 개수 : ", len(paymentDocs))

	// delete data existed
	err = service.DeletePaymentMany(groups[0].ID.Hex())
	assert.Equal(t, err, nil)

	// bulk insert
	paymentColl := mgm.Coll(&model.Payment{})
	_, err = paymentColl.InsertMany(mgm.Ctx(), paymentDocs)
	assert.Equal(t, err, nil)
}

func getRandomPrice(min int, max int) int {
	return (rand.Intn(max-min+1) + min) / 100 * 100
}

func getRadomDataShopping() (string, string) {
	min := 1
	max := len(shoppings)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(max-min+1) + min
	return shoppings[randomIndex][0], "쇼핑"
}

func getRadomDataFood() (string, string) {
	min := 1
	max := len(foods)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(max-min+1) + min
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

	fmt.Println("음식점 데이터 개수 : ", len(foods))
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

	fmt.Println("쇼핑 데이터 개수 : ", len(shoppings))
	return shoppings, nil
}
