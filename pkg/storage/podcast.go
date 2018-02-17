package storage

import (
	"database/sql"
	"log"

	"github.com/sauercrowd/go-podcasts/pkg/podcast"
)

//AddOrUpdatePodcast inserts every podcast episode into the database if possible
func AddOrUpdatePodcast(db *sql.DB, podcast *podcast.PodcastFeed) error {
	if err := addOrUpdatePodcast(db, podcast); err != nil {
		return err
	}
	for _, episode := range podcast.Items {
		if err := addOrUpdateEpisode(db, podcast.Link, &episode); err != nil {
			return err
		}
	}
	return nil
}

func addOrUpdatePodcast(db *sql.DB, podcast *podcast.PodcastFeed) error {
	var podcastid int64
	dbStr := "INSERT INTO podcasts(podcastlink, podcastname, lang, description, imageurl, updated) VALUES($1, $2, $3, $4, $5, NOW()) ON CONFLICT (podcastlink) DO UPDATE SET updated=NOW(), description=$4, imageurl=$5 RETURNING podcastid"
	err := db.QueryRow(dbStr, podcast.Link, podcast.Title, podcast.Language, podcast.Description, podcast.ImageURL).Scan(&podcastid)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func addOrUpdateEpisode(db *sql.DB, podcastlink string, episode *podcast.PodcastFeedItem) error {

	dbStr := "INSERT INTO episodes(episodelink, podcastlink, episodename, episodedescription, audiourl, pubdate) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (audiourl) DO UPDATE SET episodedescription=$4"
	err := db.QueryRow(dbStr, episode.Link, podcastlink, episode.Title, episode.Description, episode.AudioURL.URL, episode.PubDate).Scan()
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return err
	}
	return nil
}
