package tpb

import (
	"encoding/json"
	"net/http"
	"net/url"
	"tpb-tui/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

// SearchTorrents searches from the TPB API
func SearchTorrents(query string) tea.Cmd {
	return func() tea.Msg {
		endpoint := "https://apibay.org/q.php"
		resp, err := http.Get(endpoint + "?q=" + url.QueryEscape(query) + "&cat=")
		if err != nil {
			return models.ErrMsg(err)
		}
		defer resp.Body.Close()

		var torrents []models.Torrent
		if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
			return models.ErrMsg(err)
		}
		return torrents
	}
}
