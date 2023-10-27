package helpers

import (
	"io"
	"log"
	"net/http"
	"os"
)

// helper function to handleError for errors
func LogError(msg string, err error) {
	if err != nil {
		log.Printf("%s: %v", msg, err)
	}
}

func DownloadCoverArt(url string, tmp string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
