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
	translates["让我们说中文"] = "Let's speak English"
	translates["What can I do for you?"] = "宁需要我做什么？"
	translates["Select an option to continue..."] = "选择一个选项以继续..."
	translates["Create a new gonebot."] = "创建一个新的 gonebot"
	translates["Modify an existing gonebot."] = "修改现有的 gonebot"
	translates["Manage .env configurations."] = "管理 .env 配置"
	translates["Explore plugin repository."] = "探索插件仓库"
	translates["Explore adapter repository."] = "探索适配器仓库"
	translates["Exit the application."] = "退出应用程序"
}
