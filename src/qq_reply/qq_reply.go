package qq_reply

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url string
var bot_token string

func send_payload(payload *strings.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bot_token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)[0:0])
}

func Reply_text(group_id string, text string) {
	text = strings.ReplaceAll(text, "\n", `\n`)
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "text",
			"data": {
				"text": "` + text + `"
			}
		}
	]
}`)

	send_payload(payload)
}

func Reply_picture(group_id string, dir string) {
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "image",
			"data": {
				"file": "` + dir + `"
			}
		}
	]
}`)

	send_payload(payload)
}

func Group_record(group_id string, record_path string) {
	payload := strings.NewReader(`{
	"group_id": "` + group_id + `",
	"message": [
		{
			"type": "record",
			"data": {
				"file": "` + record_path + `"
			}
		}
	]
}`)

	send_payload(payload)
}

func Set_url(ip string, port string) {
	url = "http://" + ip + ":" + port + "/send_group_msg"
}

func Set_token(token string) {
	bot_token = token
}

/*		{
		"type": "at",
		"data": {
			"qq": "` + user_id + `"
		}
	},*/
