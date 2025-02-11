package openai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var gpt_url string
var gpt_model string
var gpt_APIkey string
var gpt_ini_promt string

type t_msg struct {
	Role string `json:"role"`
	Text string `json:"content"`
}

func new_t_msg(role string, content string) t_msg {
	return t_msg{
		Role: role,
		Text: content,
	}
}

type t_chat struct {
	Model string  `json:"model"`
	Msgs  []t_msg `json:"messages"`
}

func new_t_chat() t_chat {
	return t_chat{
		Model: gpt_model,
		Msgs: []t_msg{
			{
				Role: "system",
				Text: gpt_ini_promt,
			},
		},
	}
}

type t_reply struct {
	Index int   `json:"index"`
	Msg   t_msg `json:"message"`
}

type t_token_msg struct {
	Prompt_use int `json:"prompt_tokens"`
	Reply_use  int `json:"completion_tokens"`
	Total_use  int `json:"total_tokens"`
}

type t_http_reply struct {
	Choices []t_reply   `json:"choices"`
	Usage   t_token_msg `json:"usage"`
}

var err_msg t_msg = new_t_msg("system", "出错了喵")

func send2gpt(payload *strings.Reader) t_msg {
	client := &http.Client{}

	req, err := http.NewRequest("POST", gpt_url, payload)

	if err != nil {
		fmt.Println(err)
		return err_msg
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", gpt_APIkey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err_msg
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err_msg
	}

	var reply t_http_reply
	json.Unmarshal(body, &reply)

	//fmt.Println(string(body))
	fmt.Println("token use:", reply.Usage.Total_use)
	return reply.Choices[0].Msg
}

var gpt_chat = new_t_chat()

func gpt_reply() t_msg {

	jsonData, err := json.Marshal(gpt_chat)

	if err != nil {
		fmt.Println("json load failed", err)
		return err_msg
	}

	//fmt.Println(string(jsonData))
	payload := strings.NewReader(string(jsonData))

	reply := send2gpt(payload)
	gpt_chat.Msgs = append(gpt_chat.Msgs, reply)
	return reply
}

func Chat(q string) string {
	gpt_chat.Msgs = append(gpt_chat.Msgs, new_t_msg("user", q))
	return gpt_reply().Text
}

func New_chat(q string) string {
	gpt_chat = new_t_chat()
	return Chat(q)
}

func Set_url(url string) {
	gpt_url = url
}

func Set_model(model string) {
	gpt_model = model
}

func Set_APIkey(APIkey string) {
	gpt_APIkey = APIkey
}

func Set_initial_promt(initial_promt string) {
	gpt_ini_promt = initial_promt
}
