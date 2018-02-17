package podcast

import (
	"bytes"
	"encoding/xml"
)

type Enclosue struct {
	URL string `xml:"url,attr"`
}

type PodcastFeedItem struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
	AudioURL    Enclosue `xml:"enclosure"`
}

type PodcastFeed struct {
	Title       string            `xml:"channel>title"`
	Link        string            `xml:"RssDefault channel>link"`
	Language    string            `xml:"channel>language"`
	Description string            `xml:"channel>description"`
	ImageURL    string            `xml:"channel>image>url"`
	Items       []PodcastFeedItem `xml:"channel>item"`
}

func parsePodcast(content []byte) (error, *PodcastFeed) {
	var feed PodcastFeed
	d := xml.NewDecoder(bytes.NewReader(content))
	d.DefaultSpace = "RssDefault"

	err := d.Decode(&feed)
	if err != nil {
		return err, nil
	}
	return nil, &feed
}
