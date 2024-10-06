package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "gator")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var feed RSSFeed
	err = xml.Unmarshal(content, &feed)
	if err != nil {
		return nil, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, value := range feed.Channel.Item {
		feed.Channel.Item[i].Description = html.UnescapeString(value.Description)
		feed.Channel.Item[i].Title = html.UnescapeString(value.Title)
	}
	return &feed, nil
}
