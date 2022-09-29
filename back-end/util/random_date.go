package util

import (
	"math/rand"
	"time"
)

func GetRandomDate(year int, month int) time.Time {
	min := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}
