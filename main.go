package main

/*
 * go-tpb-tui is a terminal user interface for The Pirate Bay
 * It uses their public API to search for torrents and display them in a table
 * Pressing enter on a torrent will open the magnet link in your default torrent client (if you have one set up)
 * Pressing Ctrl+C will exit the application
 * Pressing Esc will go back to the search input (if you're in the table view)
 */

import (
	"fmt"
	"log"
	"strings"
	"tpb-tui/internal/models"
	"tpb-tui/internal/ui"
	"tpb-tui/pkg/tpb"
	"tpb-tui/pkg/utils"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// baseStyle is the base style for the application (Snatched from: https://github.com/charmbracelet/bubbletea/blob/master/examples/table/main.go)
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240")).
	Padding(1, 1)

const (
	TextInputView models.ViewState = iota
	TableView
)

// AppState is the application state
type AppState struct {
	ViewState  models.ViewState
	Table      table.Model
	TextInput  textinput.Model
	Torrents   []models.Torrent
	Err        error
	ErrMessage string
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// initialModel returns the initial application state
func initialModel() AppState {
	// Set up the search input field
	ti := textinput.New()
	ti.Placeholder = "The Game 1997"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return AppState{
		TextInput: ti,
		Err:       nil,
	}
}

// Init is called when the application starts
func (m AppState) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles the interaction with the application (key presses, etc.)
func (m AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		switch m.ViewState {
		// ? Textinput visas (default)
		case TextInputView:
			switch msg.Type {
			case tea.KeyEnter:
				query := m.TextInput.Value()
				if strings.TrimSpace(query) == "" {
					m.ErrMessage = "Please enter a valid search query."
					return m, nil
				}
				return m, tpb.SearchTorrents(query)
			case tea.KeyEsc:
				return m, tea.Quit
			}
			m.TextInput, cmd = m.TextInput.Update(msg)
			return m, cmd

		// ? Tabellen visas
		case TableView:
			switch msg.Type {
			case tea.KeyEnter:
				// Get the selected torrent
				selectedIndex := m.Table.Cursor()
				if selectedIndex < len(m.Torrents) {
					selectedTorrent := m.Torrents[selectedIndex]
					magnetLink := utils.GenerateMagnetLink(selectedTorrent)
					utils.OpenURL(magnetLink)
				}
			case tea.KeyEsc:
				m.ViewState = TextInputView
			default:
				m.Table, cmd = m.Table.Update(msg)
			}
			return m, cmd
		}

	case models.ErrMsg:
		m.Err = msg
		return m, nil

	case []models.Torrent:
		m.Torrents = msg
		m.Table = ui.TorrentsToTableModel(m.Torrents)
		m.ViewState = TableView
		return m, nil
	}

	return m, nil
}

// View renders the application
func (m AppState) View() string {
	switch m.ViewState {
	case TextInputView:
		return fmt.Sprintf(
			"Which torrent are you looking for? 🕵\n\n%s\n\n%s\n(esc to quit)",
			m.TextInput.View(),
			m.ErrMessage,
		) + "\n"
	case TableView:
		return baseStyle.Render(m.Table.View()) + "\n\n(Enter to open magnetlink)  (esc to go back)  (ctrl+c to quit)\n"
	default:
		return "Unknown view state!"
	}
}
