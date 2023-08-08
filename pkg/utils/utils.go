package utils

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"tpb-tui/internal/models"
)

// OpenURL opens a URL in the default browser
func OpenURL(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		log.Fatalf("Unsupported operating system")
		return
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to open magnet url: %s", err)
	}
}

// Generate a magnet link from a selected torrent
func GenerateMagnetLink(torrent models.Torrent) string {
	base := "magnet:?xt=urn:btih:"
	trackers, err := getTrackersFromFile()
	if err != nil {
		log.Fatalf("Failed to read trackers: %s", err)
	}

	encodedName := url.QueryEscape(torrent.Name)
	magnet := fmt.Sprintf("%s%s&dn=%s", base, torrent.InfoHash, encodedName)
	for _, tracker := range trackers {
		magnet += "&tr=" + url.QueryEscape(tracker)
	}
	return magnet
}
