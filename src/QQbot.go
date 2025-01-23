package src

import (
	"QQbot/src/GPT"
	"QQbot/src/qq_reply"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func reduce_pre_sp(msg string, pre string) string {
	msg = msg[len(pre):]
	msg = strings.TrimLeft(msg, " ")
	return msg
}

var at_botQQ string

func group_handler(msg *group_msg) {
	if strings.HasPrefix(msg.CQ_text, at_botQQ) { //@自己的
		group_id := strconv.Itoa(msg.Group_id)
		text := reduce_pre_sp(msg.CQ_text, at_botQQ)

		if len(text) == 0 {
			qq_reply.Reply_text(group_id, "什么事喵？")
			return
		}

		if text == "miku" {
			//之后再写
		}

		// if rest_msg == "今日运势" || rest_msg == "y" {
		// 	fmt.Println(work_dir+"/luck_files/back/", work_dir+"/luck_files/front/", work_dir+"/luck_files/cute.ttf", work_dir+"/luck_files/luck_pics/")
		// 	uid := msg.Sender.QQ
		// 	luck.Gen_Pic(int64(uid), work_dir+"/luck_files/back/", work_dir+"/luck_files/front/", work_dir+"/luck_files/cute.ttf", work_dir+"/luck_files/luck_pics/")
		// 	qq_reply.Luck_reply(strconv.Itoa(msg.Group_id), strconv.Itoa(uid), work_dir+"/luck_files/luck_pics/"+strconv.Itoa(uid)+".png")
		// 	return
		// }

		//text = reduce_pre_sp(text, "那我问你 ")
		fmt.Println("问题：" + text)
		reply := GPT.AIreply(text)
		fmt.Println("回答：" + reply)
		reply_len := len(reply)
		max_len := 3000
		if reply_len > max_len {
			end := "。。可惜这里地方太小，写不下喵。"
			reply = reply[0:max_len-len(end)] + end
		}

		qq_reply.Reply_text(group_id, reply)
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

func Main() {
	qq_reply.Set_url(server_ip, server_port)

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

func Set_GPT_url(url string) {
	GPT.Set_url(url)
}

func Set_model(model string) {
	GPT.Set_model(model)
}

func Set_APIkey(APIkey string) {
	GPT.Set_APIkey(APIkey)
}

func Set_initial_promt(initial_promt string) {
	GPT.Set_initial_promt(initial_promt)
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
