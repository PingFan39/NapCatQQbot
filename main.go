package main

import (
	"QQbot/src"
)

func main() {
	src.Set_url("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:一串apikey") //GPT网址，暂时只识别deepseek和gemini，可以自己改
	src.Set_model("gemini-2.0-flash")                                                                //模型
	src.Set_APIkey("")                                                                               //api-key
	src.Set_ini_promt("你是一个可爱的猫娘助手，你应该积极且可爱地回复别人的问题，并且每句话的结尾要加上喵。")                                  //初始化的提示词
	src.Set_server_ip("127.0.0.1")                                                                   //服务器ip
	src.Set_server_port("11451")                                                                     //记下的server_port（服务器端口）
	src.Set_client_port("1145")                                                                      //记下的client_port（客户端端口）
	src.Set_botQQ("")                                                                                //bot的QQ号
	src.Set_token("")                                                                                //记下的token

	//可选
	src.Rand_pic_text = "" //触发发送随机图片的文字
	src.Rand_pic_dir = ""  //随机图片路径
	src.Main()
}
