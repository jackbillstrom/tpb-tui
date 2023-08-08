package utils

import (
	"bufio"
	"os"
)

// getTrackersFromFile reads a list of trackers from a trackers.txt file and returns a slice of strings
func getTrackersFromFile() ([]string, error) {
	file, err := os.Open("trackers.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var trackers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trackers = append(trackers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return trackers, nil
}
