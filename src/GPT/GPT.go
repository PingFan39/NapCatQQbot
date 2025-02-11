package GPT

import (
	"QQbot/src/GPT/gemini"
	"QQbot/src/GPT/openai"
)

var Is_OpenAI_format bool
var Is_gemini_format bool

func preprocess(p_text *string) {
	text := *p_text
	if text == "帮助" || text == "help" || len(text) == 0 {
		*p_text = "不需要多余文字，请介绍自己并用自己的话附上以下内容：如需要任何帮助，请跳转到https://github.com/PingFan39/NapCatQQbot"
	}
}

func Chat(text string) string {
	preprocess(&text)
	if Is_OpenAI_format {
		return openai.Chat(text)
	}
	return gemini.Chat(text)
}

func New_chat(text string) string {
	preprocess(&text)
	if Is_OpenAI_format {
		return openai.New_chat(text)
	}
	return gemini.New_chat(text)
}
