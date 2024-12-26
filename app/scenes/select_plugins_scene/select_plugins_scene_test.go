package selectpluginsscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	selectpluginsscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/select_plugins_scene"
)

func TestSelectPluginsScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(selectpluginsscene.SelectPluginsScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running select plugins scene: ", err)
		t.Fail()
	}
}
