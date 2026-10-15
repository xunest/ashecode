package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// lightTerminal records whether the terminal uses a light background.
// It is detected once at package initialization — before bubbletea takes
// over stdin — because termenv's OSC 11 background probe temporarily
// switches the TTY to raw/no-echo mode and consumes the reply itself.
// On a non-TTY (piped -p output, CI) the probe is skipped and we
// conservatively keep the dark palette.
var lightTerminal = !termenv.HasDarkBackground()

// pickColor selects between a color tuned for dark backgrounds and one
// tuned for light backgrounds.
func pickColor(dark, light string) lipgloss.Color {
	if lightTerminal {
		return lipgloss.Color(light)
	}
	return lipgloss.Color(dark)
}

var (
	// Brand colors — values are (dark background, light background).
	brandPurple    = pickColor("99", "92")
	dimText        = pickColor("242", "244")
	mutedText      = pickColor("245", "240")
	normalText     = pickColor("252", "235")
	brightText     = pickColor("255", "232")
	greenText      = pickColor("78", "28")
	redText        = pickColor("203", "124")
	yellowText     = pickColor("214", "142")
	cyanText       = pickColor("80", "30")
	separatorColor = pickColor("236", "248")

	// Teammate progress tree (truecolor hex).
	progressCyan   = pickColor("#00d7ff", "#008aaf")
	progressGreen  = pickColor("#00ff00", "#008700")
	progressRed    = pickColor("#ff0000", "#c00000")
	progressYellow = pickColor("#ffff00", "#8a7000")

	// Banner
	bannerStyle = lipgloss.NewStyle().
			Foreground(brandPurple).
			Bold(true)

	bannerDimStyle = lipgloss.NewStyle().
			Foreground(dimText)

	// Separator line
	separatorStyle = lipgloss.NewStyle().
			Foreground(separatorColor)

	// User prompt marker
	promptStyle = lipgloss.NewStyle().
			Foreground(cyanText).
			Bold(true)

	// AI response marker
	aiMarkerStyle = lipgloss.NewStyle().
			Foreground(brandPurple).
			Bold(true)

	// AI text
	aiTextStyle = lipgloss.NewStyle().
			Foreground(normalText).
			PaddingLeft(2)

	// Streaming text (slightly dimmer while streaming)
	streamingTextStyle = lipgloss.NewStyle().
				Foreground(normalText).
				PaddingLeft(2)

	// Tool call styles

	toolRunningStyle = lipgloss.NewStyle().
				Foreground(dimText).
				PaddingLeft(2)

	toolDoneStyle = lipgloss.NewStyle().
			Foreground(greenText).
			PaddingLeft(2)

	toolErrorStyle = lipgloss.NewStyle().
			Foreground(redText).
			PaddingLeft(2)

	toolDetailStyle = lipgloss.NewStyle().
			Foreground(dimText).
			PaddingLeft(4)

	// EditFile diff 行样式：新增绿、删除红，上下文行复用 toolDetailStyle
	diffAddStyle = lipgloss.NewStyle().
			Foreground(greenText).
			PaddingLeft(4)

	diffRemoveStyle = lipgloss.NewStyle().
			Foreground(redText).
			PaddingLeft(4)

	// Error message
	errorStyle = lipgloss.NewStyle().
			Foreground(redText).
			PaddingLeft(2)

	// Permission dialog
	permBorderStyle = lipgloss.NewStyle().
			Foreground(yellowText).
			Bold(true)

	permDimStyle = lipgloss.NewStyle().
			Foreground(dimText)

	// Status bar (bottom)
	statusBarStyle = lipgloss.NewStyle().
			Foreground(dimText)

	statusItemStyle = lipgloss.NewStyle().
			Foreground(mutedText)

	// Provider selection
	selectLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(brandPurple).
				Align(lipgloss.Center)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(cyanText).
				Bold(true)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(mutedText)
)

// markdownStyleName returns the glamour standard style that matches the
// terminal background, so markdown code/headings stay readable on both.
func markdownStyleName() string {
	if lightTerminal {
		return "light"
	}
	return "dark"
}
