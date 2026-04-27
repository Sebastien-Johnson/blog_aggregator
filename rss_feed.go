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

//Grabs rss feed/structure from website html
func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	client := &http.Client{}
	//create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}
	//set request header to help program ID the server
	req.Header.Set("User-Agent", "gator")

	//run request and get response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//reads and confirms full file
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//create var to unmarshal xml data into into
	rssData := RSSFeed{}
	err = xml.Unmarshal(dat, &rssData)
	if err != nil {
		return nil, err
	}

	//UnescapeString turns things like "&amp;, &lt;, &ldquo;" into their actual characters
	//Unescape title and description directly
	rssData.Channel.Title = html.UnescapeString(rssData.Channel.Title)
	rssData.Channel.Description = html.UnescapeString(rssData.Channel.Description)

	//Unescape all items *within* tilte and description
	for i := range rssData.Channel.Item {
		rssData.Channel.Item[i].Title = html.UnescapeString(rssData.Channel.Item[i].Title)
		rssData.Channel.Item[i].Description = html.UnescapeString(rssData.Channel.Item[i].Description)
	}
	
	return &rssData, nil
}

