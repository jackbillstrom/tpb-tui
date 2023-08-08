package ui

import (
	"tpb-tui/internal/models"

	"github.com/charmbracelet/bubbles/table"
)

// TorrentsToTableModel converts a slice of Torrents to a table.Model
func TorrentsToTableModel(torrents []models.Torrent) table.Model {
	// Define the table columns
	columns := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Seeders", Width: 10},
		{Title: "Leechers", Width: 10},
		{Title: "Size", Width: 20},
		{Title: "Username", Width: 20},
	}

	rows := make([]table.Row, len(torrents))

	for i, t := range torrents {
		row := table.Row{
			t.Name,
			t.Seeders,
			t.Leechers,
			t.Size,
			t.Username,
		}
		rows[i] = row
	}

	// Create the table
	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
}
