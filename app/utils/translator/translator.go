package translaltor

import (
	"fmt"

	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

var translates = make(map[string]string)

func Translate(text string) string {
	if base.Lang == "en" {
		return text
	}
	val, ok := translates[text]
	if !ok {
		panic(fmt.Sprintf("Translation not found: %s!", text))
	}
	return val
}

func init() {
	// ! Footer
	translates["Exit"] = "退出"
	translates["让我们说中文"] = "Let's speak English"
	translates["Refresh"] = "刷新"
	translates["Prev"] = "上一步"
	translates["Next"] = "下一步"

	// ! Menu Scene
	translates["What can I do for you?"] = "宁需要我做什么？"
	translates["Select an option to continue..."] = "选择一个选项以继续..."
	translates["Create a new gonebot."] = "创建一个新的 gonebot"
	translates["Modify an existing gonebot."] = "修改现有的 gonebot"
	translates["Manage .env configurations."] = "管理 .env 配置"
	translates["Explore plugin repository."] = "探索插件仓库"
	translates["Explore adapter repository."] = "探索适配器仓库"
	translates["Exit the application."] = "退出应用程序"

	// ! New Bot Scene
	translates["Enter bot name:"] = "输入 bot 名称："
	translates["Name of your bot."] = "你的 bot 名称"
	translates["Enter bot version:"] = "输入 bot 版本号："
	translates["Version of your bot."] = "你的 bot 初始版本号"
	translates["Enter bot description:"] = "输入 bot 描述："
	translates["A short description of your bot."] = "简短描述你的 bot"
	translates["Select a folder..."] = "选择一个文件夹..."
	translates["We will create your bot folder here."] = "我们将在这里创建你的 bot 文件夹"

	// ! Select Adapters Scene
	translates["Syncing Repository..."] = "正在同步仓库..."
	translates["Select Adapters..."] = "选择适配器..."

	// ! [Component]Selected List
	translates["Selected Adapters"] = "已选择的适配器"
	translates["Selected Plugins"] = "已选择的插件"

	// ! Select Plugins Scene
	translates["Select Plugins..."] = "选择插件..."
}
