package base

const (
	MaxWidth = 80

	NewBot = iota
	EditBot
	DotEnv
	Plugins
	Adapters
	ExitApp
)

var (
	// Interactive
	WindowHeight   int    = 50
	WindowWidth    int    = 80
	Lang           string = "en"
	SelectedAction        = -1
)
