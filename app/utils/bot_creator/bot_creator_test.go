package botcreator_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
	botcreator "github.com/gonebot-dev/gonebuilder-tui/app/utils/bot_creator"
)

func TestBotCreator(t *testing.T) {
	go func() {
		err := botcreator.CreateBot("/home/kingcq/WorkSpace", "ABot", "v0.0.1", "A simple test bot.", &[]api.AdapterInfo{
			{
				Name:    "onebotv11",
				Adapter: "onebotv11.OneBotV11",
				Desc:    "OneBotV11 adapter for gonebot. Use it for QQ.",
				DescZH:  "使用 OneBotV11 协议的 QQ 适配器。",
				Package: "github.com/gonebot-dev/goneadapter-onebotv11",
			},
		}, &[]api.PluginInfo{
			{
				Name:    "echo",
				Plugin:  "echo.Echo",
				Desc:    "Replys what you say",
				DescZH:  "回复你说的内容",
				Package: "github.com/gonebot-dev/goneplugin-echo",
			},
		})
		if err != nil {
			t.Error(err)
		}
	}()
	for prog, _ := botcreator.GetInfo(); prog < 6; prog, _ = botcreator.GetInfo() {
		fmt.Printf("\rProgress: %d", prog)
		time.Sleep(time.Millisecond * 10)
	}
}
