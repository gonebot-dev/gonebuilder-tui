package initialscene_test

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	initialscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/initial_scene"
)

func TestInitialScene(t *testing.T) {
	app := tea.NewProgram(initialscene.InitialScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running initial scene: ", err)
		t.Fail()
	}
}
