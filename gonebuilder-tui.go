package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gonebot-dev/gonebuilder-tui/app"
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	if os.Getenv("GONEREPO") == "" {
		os.Setenv("GONEREPO", "gonebot-dev/gonerepo")
	}

	app := tea.NewProgram(app.CreateApp(), tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Println("Error running gonebuilder: ", err)
		os.Exit(1)
	}
	if base.PostFunc != nil {
		base.PostFunc()
	}
}
