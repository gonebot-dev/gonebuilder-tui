package menuscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	menuscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/menu_scene"
)

func TestMenuScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(menuscene.MenuScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running menu scene: ", err)
		t.Fail()
	}
}
