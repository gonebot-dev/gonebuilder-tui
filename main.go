package gonebuilder

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
}

func (app App) Init() tea.Cmd {
	return nil
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return app, nil
}

func (app App) View() string {
	return "Hello, world!"
}

func main() {
	app := tea.NewProgram(App{}, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
