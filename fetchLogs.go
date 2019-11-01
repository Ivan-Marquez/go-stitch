package gostitch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LogResponse from Stitch
type LogResponse struct {
	Logs []Logger `json:"logs"`
}

// Logger to represent log rows
type Logger struct {
	Started   string   `json:"started"`
	Completed string   `json:"completed"`
	Function  string   `json:"function_name"`
	Trigger   string   `json:"event_subscription_name"`
	Messages  []string `json:"messages"`
	Error     string   `json:"error"`
}

func setLogRequest(accessToken, projectID, appID string) *http.Request {
	url := fmt.Sprintf("%v/groups/%v/apps/%v/logs", BaseURL, projectID, appID)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer "+accessToken)

	return req
}

// FetchLogs retrieves logs from MongoDB Stitch and returns a LogReponse
func FetchLogs(accessToken, projectID, appID string) (response *LogResponse, err error) {
	client := http.Client{}

	req := setLogRequest(accessToken, projectID, appID)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
