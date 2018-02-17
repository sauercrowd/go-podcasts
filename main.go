package main

import (
	//"github.com/sauercrowd/go-podcasts/pkg/podcast"
	"log"

	"github.com/sauercrowd/go-podcasts/pkg/flags"
	"github.com/sauercrowd/go-podcasts/pkg/podcast"
	"github.com/sauercrowd/go-podcasts/pkg/storage"
)

func main() {
	v := flags.Parse()
	err, db := storage.Setup(v)
	if err != nil {
		log.Println(err)
		return
	}
	err, feed := podcast.Get("http://podcast-ufo.fail/?feed=rss2&cat=2")
	if err != nil {
		log.Println("error: ", err)
		return
	}
	if err := storage.AddOrUpdatePodcast(db, feed); err != nil {
		log.Println(err)
		return
	}

}
