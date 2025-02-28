package src

import (
	"QQbot/src/GPT"
	"QQbot/src/GPT/gemini"
	"QQbot/src/GPT/openai"
	"QQbot/src/qq_reply"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var work_dir string

type post_head struct {
	Mtype  string `json:"message_type"`
	Mid    int    `json:"message_id"`
	Mstype string `json:"sub_type"`
}

type sender struct {
	QQ   int    `json:"user_id"`
	Name string `json:"nickname"`
	Role string `json:"role"`
}

type group_msg struct {
	Sender   sender `json:"sender"`
	Group_id int    `json:"group_id"`
	CQ_text  string `json:"raw_message"`
}

func has_pres(p_text *string, pres string) bool {
	text := *p_text
	if strings.HasPrefix(text, pres) {
		text = text[len(pres):]
		text = strings.TrimLeft(text, " ")
		*p_text = text
		return true
	}
	return false
}

var at_botQQ string

var Rand_pic_text string
var Rand_pic_dir string

func group_handler(msg *group_msg) {
	id := strconv.Itoa(msg.Group_id)
	text := msg.CQ_text
	if has_pres(&text, at_botQQ) { //@自己的
		if text == Rand_pic_text {
			file_names, err := os.ReadDir(Rand_pic_dir)
			if err != nil {
				fmt.Println("路径有问题喵:", err)
				return
			}
			qq_reply.Reply_picture(
				id,
				Rand_pic_dir+file_names[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(file_names))].Name())
			return
		}
		var reply string
		if has_pres(&text, "，") || has_pres(&text, ",") {
			fmt.Println("追问：" + text)
			reply = GPT.Chat(text)
		} else {
			fmt.Println("首问：" + text)
			reply = GPT.New_chat(text)
		}

		fmt.Println("回答：" + reply)
		reply_len := len(reply)
		max_len := 3000
		if reply_len > max_len {
			end := "。。可惜这里地方太小，写不下喵。"
			reply = reply[0:max_len-len(end)] + end
		}

		qq_reply.Reply_text(id, reply)
	}
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not a POST message", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body read failed", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var head post_head
	json.Unmarshal(body, &head)
	//fmt.Printf("Head_info: %+v\n", head)

	if head.Mtype == "group" {
		var msg group_msg
		json.Unmarshal(body, &msg)
		//fmt.Printf("group_info: %+v\n", msg)
		group_handler(&msg)
	}

	//fmt.Println(string(body))
}

var server_ip string
var server_port string
var client_port string

var gpt_url string
var gpt_model string
var gpt_APIkey string
var gpt_ini_promt string

func Main() {

	qq_reply.Set_url(server_ip, server_port)

	if strings.Contains(gpt_url, "openai") {
		openai.Set_url(gpt_url)
		openai.Set_model(gpt_model)
		openai.Set_APIkey(gpt_APIkey)
		openai.Set_initial_promt(gpt_ini_promt)
		GPT.Is_OpenAI_format = true
	}
	if strings.Contains(gpt_url, "googleapis") {
		gemini.Set_url(gpt_url)
		gemini.Set_initial_promt(gpt_ini_promt)
		GPT.Is_gemini_format = true
	}

	var err error
	work_dir, err = os.Getwd()
	if err != nil {
		fmt.Println("无法获取工作目录:", err)
		return
	}

	http.HandleFunc("/", http_handler)
	fmt.Println("服务器启动成功")
	err = http.ListenAndServe(":"+client_port, nil)

	if err != nil {
		fmt.Println("不了一点\n错误信息:", err)
	}
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

func Set_ini_promt(initial_promt string) {
	gpt_ini_promt = initial_promt
}

func Set_botQQ(qq string) {
	at_botQQ = "[CQ:at,qq=" + qq + "]"
}

func Set_server_ip(ip string) {
	server_ip = ip
}

func Set_server_port(port string) {
	server_port = port
}

func Set_client_port(port string) {
	client_port = port
}

func Set_token(token string) {
	qq_reply.Set_token(token)
}
