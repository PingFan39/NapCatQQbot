package gemini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var gpt_url string
var gpt_ini_promt string
var use_proxy bool

type t_part struct {
	Text string `json:"text"`
}

type t_system_content struct {
	Parts []t_part `json:"parts"`
}

type t_content struct {
	Role  string   `json:"role"`
	Parts []t_part `json:"parts"`
}

func new_t_content(role string, content string) t_content {
	var parts []t_part
	parts = append(parts, t_part{Text: content})
	return t_content{
		Role:  role,
		Parts: parts,
	}
}

type t_empty struct {
}

type t_tool struct {
	Empty t_empty `json:"google_search"`
}

type t_safety_rating struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type t_config struct {
	MaxToken    int     `json:"maxOutputTokens"`
	Temperature float64 `json:"temperature"`
}

type t_chat struct {
	Instruction    t_system_content  `json:"systemInstruction"`
	Contents       []t_content       `json:"contents"`
	Tools          []t_tool          `json:"tools"`
	SafetySettings []t_safety_rating `json:"safetySettings"`
	Config         t_config          `json:"generationConfig"`
}

var token_limit int = 700
var temp_set float64 = 1

func new_t_chat() t_chat {
	return t_chat{
		Instruction: t_system_content{
			Parts: []t_part{
				{
					Text: gpt_ini_promt,
				},
			},
		},
		Tools: []t_tool{},
		SafetySettings: []t_safety_rating{ // 默认所有安全控制的级别都是不阻拦
			{
				Category:  "HARM_CATEGORY_HARASSMENT",
				Threshold: "BLOCK_NONE",
			},
			{
				Category:  "HARM_CATEGORY_HATE_SPEECH",
				Threshold: "BLOCK_NONE",
			},
			{
				Category:  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
				Threshold: "BLOCK_NONE",
			},
			{
				Category:  "HARM_CATEGORY_DANGEROUS_CONTENT",
				Threshold: "BLOCK_NONE",
			},
			{
				Category:  "HARM_CATEGORY_CIVIC_INTEGRITY",
				Threshold: "BLOCK_NONE",
			},
		},
		Config: t_config{
			MaxToken:    token_limit,
			Temperature: temp_set,
		},
	}
}

var err_msg string = "出错了喵"

func send2gpt(payload *strings.Reader) string {
	client := &http.Client{}

	if use_proxy {
		ProxyURL, err := url.Parse("http://127.0.0.1:7890")
		if err != nil {
			return "代理出错了喵"
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(ProxyURL),
		}

		client = &http.Client{
			Transport: transport,
		}
	}

	req, err := http.NewRequest("POST", gpt_url, payload)

	if err != nil {
		fmt.Println(err)
		return err_msg
	}

	req.Header.Add("Content-Type", "application/json")

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

	var reply map[string]interface{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		fmt.Println(err)
		return err_msg
	}

	reply_parts, ok := reply["candidates"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["parts"].([]interface{})[0].(map[string]interface{})
	if !ok {
		return "我不会喵"
	}

	return reply_parts["text"].(string)
}

var gpt_chat = new_t_chat()

func gpt_reply() string {

	jsonData, err := json.Marshal(gpt_chat)

	if err != nil {
		fmt.Println("json load failed", err)
		return err_msg
	}

	//fmt.Println(string(jsonData))
	payload := strings.NewReader(string(jsonData))

	reply := send2gpt(payload)
	gpt_chat.Contents = append(gpt_chat.Contents, new_t_content("model", reply))
	return reply
}

func Chat(q string) string {
	gpt_chat.Contents = append(gpt_chat.Contents, new_t_content("user", q))
	return gpt_reply()
}

func New_chat(q string) string {
	gpt_chat = new_t_chat()
	return Chat(q)
}

func Set_url(url string) {
	gpt_url = url
}

func Set_initial_promt(initial_promt string) {
	gpt_ini_promt = initial_promt
}
