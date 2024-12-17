package initialscene

import "github.com/charmbracelet/lipgloss"

const Banner = `
=========================================
   ______                 __          __
  / ____/___  ____  ___  / /_  ____  / /_
 / / __/ __ \/ __ \/ _ \/ __ \/ __ \/ __/
/ /_/ / /_/ / / / /  __/ /_/ / /_/ / /_
\____/\____/_/ /_/\___/_.___/\____/\__/
=========================================
`

var MainFrame = struct {
	Style   lipgloss.Style
	Padding int
}{
	Style: lipgloss.NewStyle().
		Bold(true).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center),
	Padding: 8,
}
