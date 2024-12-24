package selectadaptersscene_test

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	selectadaptersscene "github.com/gonebot-dev/gonebuilder-tui/app/scenes/select_adapters_scene"
)

func TestSelectAdapterScene(t *testing.T) {
	os.Setenv("DEBUG", "true")
	app := tea.NewProgram(selectadaptersscene.SelectAdaptersScene, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running select adapters scene: ", err)
		t.Fail()
	}
}
