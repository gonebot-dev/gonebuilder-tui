package base

import "github.com/charmbracelet/lipgloss"

var Colors = struct {
	Red      lipgloss.CompleteAdaptiveColor
	Green    lipgloss.CompleteAdaptiveColor
	Blue     lipgloss.CompleteAdaptiveColor
	Yellow   lipgloss.CompleteAdaptiveColor
	Gray     lipgloss.CompleteAdaptiveColor
	Lavender lipgloss.CompleteAdaptiveColor
	GoBlue   lipgloss.CompleteAdaptiveColor
	Primary  lipgloss.CompleteAdaptiveColor
	Weak     lipgloss.CompleteAdaptiveColor
}{
	Red: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "1",
			ANSI256:   "124",
			TrueColor: "#f56c6c",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "9",
			ANSI256:   "196",
			TrueColor: "#f56c6c",
		},
	},
	Green: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "2",
			ANSI256:   "46",
			TrueColor: "#67c23a",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "10",
			ANSI256:   "82",
			TrueColor: "#67c23a",
		},
	},
	Blue: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "4",
			ANSI256:   "33",
			TrueColor: "#409eff",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "12",
			ANSI256:   "39",
			TrueColor: "#409eff",
		},
	},
	Yellow: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "3",
			ANSI256:   "226",
			TrueColor: "#e6a23c",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "11",
			ANSI256:   "220",
			TrueColor: "#e6a23c",
		},
	},
	Gray: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "8",
			ANSI256:   "239",
			TrueColor: "#6c6e72",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "7",
			ANSI256:   "248",
			TrueColor: "#cfd3dc",
		},
	},
	Lavender: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "5",
			ANSI256:   "93",
			TrueColor: "#c6a0f6",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "13",
			ANSI256:   "135",
			TrueColor: "#c6a0f6",
		},
	},
	GoBlue: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "6",
			ANSI256:   "31",
			TrueColor: "#367b99",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "14",
			ANSI256:   "45",
			TrueColor: "#367b99",
		},
	},
	Primary: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "0",
			ANSI256:   "235",
			TrueColor: "#303133",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "15",
			ANSI256:   "255",
			TrueColor: "#e5eaf3",
		},
	},
	Weak: lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			ANSI:      "8",
			ANSI256:   "239",
			TrueColor: "#6c6e72",
		},
		Dark: lipgloss.CompleteColor{
			ANSI:      "7",
			ANSI256:   "248",
			TrueColor: "#6c6e72",
		},
	},
}
