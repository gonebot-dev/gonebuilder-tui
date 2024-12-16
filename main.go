package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app"
)

func main() {
	app := tea.NewProgram(app.CreateApp(), tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running gonebuilder: ", err)
		os.Exit(1)
	}
}
