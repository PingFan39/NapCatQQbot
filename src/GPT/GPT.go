package GPT

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var gpt_url string
var gpt_model string
var gpt_api_key string
var gpt_initial_promt string

type gpt_msg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chat_session struct {
	Model    string    `json:"model"`
	Messages []gpt_msg `json:"messages"`
}

type gpt_reply struct {
	Index   int     `json:"index"`
	Message gpt_msg `json:"message"`
}

type token_use struct {
	Prompt_use int `json:"prompt_tokens"`
	Reply_use  int `json:"completion_tokens"`
	Total_use  int `json:"total_tokens"`
}

type http_reply struct {
	Choices []gpt_reply `json:"choices"`
	Usage   token_use   `json:"usage"`
}

func New_GPTmsg(mrole string, mcontent string) gpt_msg {
	return gpt_msg{
		Role:    mrole,
		Content: mcontent,
	}
}

func NewGPTChat() chat_session {
	return chat_session{
		Model: gpt_model, // 如果是openai则改为想用的模型如gpt-4o gpt-4o-mini
		Messages: []gpt_msg{
			{
				Role:    "system",
				Content: gpt_initial_promt,
			},
		},
	}
}

func send2gpt(payload *strings.Reader) string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", gpt_url, payload)

	if err != nil {
		fmt.Println(err)
		return "出错了喵"
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", gpt_api_key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "出错了喵"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "出错了喵"
	}

	var reply http_reply
	json.Unmarshal(body, &reply)

	//fmt.Println(string(body))
	fmt.Println("token use:", reply.Usage.Total_use)
	return reply.Choices[0].Message.Content // 默认采用第一个回复
}

func AIreply(question string) string { // 创建一个新对话并返回gpt的回答
	gpt_session := NewGPTChat()
	msg := New_GPTmsg("user", question)
	gpt_session.Messages = append(gpt_session.Messages, msg)

	jsonData, err := json.Marshal(gpt_session)

	if err != nil {
		fmt.Println("json load failed", err)
		return "出错了喵"
	}

	//fmt.Println(string(jsonData))
	payload := strings.NewReader(string(jsonData))

	return send2gpt(payload)
}

func Set_url(url string) {
	gpt_url = url
}

func Set_model(model string) {
	gpt_model = model
}

func Set_APIkey(APIkey string) {
	gpt_api_key = APIkey
}

func Set_initial_promt(initial_promt string) {
	gpt_initial_promt = initial_promt
}
