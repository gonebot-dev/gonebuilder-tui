package newbotscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	newbotscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/new_bot_scene"
)

func TestNewBotScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(newbotscene.NewBotScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running menu sccene: ", err)
		t.Fail()
	}
}
