package initialscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	initialscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/initial_scene"
)

func TestInitialScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(initialscene.InitialScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running initial scene: ", err)
		t.Fail()
	}
}
