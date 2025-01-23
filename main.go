package main

import (
	"QQbot/src"
)

func main() {
	src.Set_GPT_url("https://api.deepseek.com/chat/completions")        //GPT网址
	src.Set_model("deepseek-chat")                                      //模型
	src.Set_APIkey("Bearer sk-")                                        //"sk-"后面补上sk-key
	src.Set_initial_promt("你是一个可爱的猫娘助手，你应该积极且可爱地回复别人的问题，并且每句话的结尾要加上喵。") //初始化的提示词
	src.Set_server_ip("127.0.0.1")                                      //服务器ip
	src.Set_server_port("")                                             //服务器端口
	src.Set_client_port("")                                             //客户端端口
	src.Set_botQQ("")                                                   //bot的QQ号
	src.Set_token("")                                                   //[info] [NapCat] [WebUi] WebUi Local Panel Url:
	src.Main()
}
