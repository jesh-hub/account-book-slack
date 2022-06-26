package abs

import (
	"github.com/goccy/go-json"
	"testing"
)

func init() {
	ConnectDB()
}

func TestConnect(t *testing.T) {
	ConnectDB()
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
