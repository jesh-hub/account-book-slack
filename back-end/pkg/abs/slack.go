package abs

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SlackClient struct {
	token      string
	channelId  string
	botId      string
	httpClient *http.Client
}

func NewSlackClient(token string, channelId string, botId string) *SlackClient {
	return &SlackClient{
		token:     token,
		channelId: channelId,
		botId:     botId,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *SlackClient) NewAPI(uri string, urlQuery interface{}, target interface{}) {
	req, err := http.NewRequest("GET", uri, nil)
	errorHandler(err)

	// 헤더에 토큰 설정
	req.Header.Add("Authorization", "Bearer "+s.token)

	// url 파라미터 설정
	q := req.URL.Query()
	q.Add("channel", s.channelId)
	if urlQuery != nil {
		var urlQueryMap map[string]string
		urlQueryJson, _ := json.Marshal(urlQuery)
		json.Unmarshal(urlQueryJson, &urlQueryMap)
		for k, v := range urlQueryMap {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err := s.httpClient.Do(req)
	errorHandler(err)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
