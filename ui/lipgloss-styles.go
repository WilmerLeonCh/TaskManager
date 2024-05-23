package ui

import "github.com/charmbracelet/lipgloss"

const (
	white   = lipgloss.Color("#fefffe")
	skyBlue = lipgloss.Color("#8ACDEA")
	gray    = lipgloss.Color("#939ab4")
)

var PromptStyle = lipgloss.
	NewStyle().
	Foreground(skyBlue).
	Bold(true)
var PlaceholderStyle = lipgloss.NewStyle().
	Foreground(gray).
	Padding(0, 1)
var TitleListStyle = lipgloss.
	NewStyle().
	Foreground(skyBlue).
	Width(50).
	Bold(true)
var ItemListStyle = lipgloss.
	NewStyle().
	Foreground(skyBlue)
var TableStyle = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(skyBlue).
	Padding(0, 1)
var HeaderTableStyle = lipgloss.
	NewStyle().
	Foreground(skyBlue).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(skyBlue).
	BorderBottom(true).
	Bold(true)
var SelectedTableStyle = lipgloss.
	NewStyle().
	Foreground(white).
	Background(skyBlue)
var MessageStyle = lipgloss.NewStyle().
	Foreground(gray).
	Italic(true)
