// Register all scenes here.
package router

import (
	"github.com/gonebot-dev/gonebuilder-tui/app/scene"
	initialscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/initial_scene"
	menuscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/menu_scene"
)

func init() {
	scene.RegisterScene("InitialScene", initialscene.InitialScene)
	scene.RegisterScene("MenuScene", menuscene.MenuScene)
}
