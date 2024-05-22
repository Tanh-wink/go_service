package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
var (
	Bot = NewWeComBot(conf.RobotURL, conf.MentionWhoList)
)

// 企业微信告警
type WeComBot struct {
	RobotURL       string
	MentionWhoList []string
}

// 构造函数
func NewWeComBot(robotURL string, mentionWhoList []string) *WeComBot {
	return &WeComBot{
		RobotURL:       robotURL,
		MentionWhoList: mentionWhoList,
	}
}

func (w *WeComBot) sendRequest(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", w.RobotURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("WeComBot response: %d, %s\n", resp.StatusCode, body)

	return nil
}

// send text message
func (w *WeComBot) SendText(msg string) error {
	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":             msg,
			"mentioned_mobile_list": w.MentionWhoList,
		},
	}
	return w.sendRequest(data)
}

// send markdown message
func (w *WeComBot) SendMarkdown(md string) error {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": md,
		},
	}
	return w.sendRequest(data)
}

// func main() {
	

// 	err := bot.SendText("This is a text message")
// 	if err != nil {
// 		fmt.Println("Error sending text message:", err)
// 		os.Exit(1)
// 	}

// 	err = bot.SendMarkdown("# This is a markdown message")
// 	if err != nil {
// 		fmt.Println("Error sending markdown message:", err)
// 		os.Exit(1)
// 	}
// }