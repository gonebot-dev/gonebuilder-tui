package translaltor

import "github.com/gonebot-dev/gonebuilder-tui/app/base"

var translates = make(map[string]string)

func Translate(text string) string {
	if base.Lang == "en" {
		return text
	}
	val, ok := translates[text]
	if !ok {
		panic("Translation not found!")
	}
	return val
}

func init() {
	translates["简体中文"] = "English"
	translates["What can I do for you?"] = "宁需要我做什么？"
}
