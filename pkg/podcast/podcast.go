package podcast

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) (error, *PodcastFeed) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not load RSS")
		return err, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not read Response Body: %v", err)
		return err, nil
	}
	return parsePodcast(body)
}
