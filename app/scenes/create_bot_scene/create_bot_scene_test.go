package createbotscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	createbotscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/create_bot_scene"
)

func TestCreateBotScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(createbotscene.CreateBotScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running new bot scene: ", err)
		t.Fail()
	}
}
