package main

import "time"

// SearchResult Json object for api search results
type SearchResult struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID     string `json:"id"`
		Parent struct {
			ID       string      `json:"id"`
			Name     string      `json:"name"`
			Color    string      `json:"color"`
			Icon     interface{} `json:"icon"`
			Created  time.Time   `json:"created"`
			Modified time.Time   `json:"modified"`
			Parent   string      `json:"parent"`
		} `json:"parent"`
		Tags []struct {
			ID       string    `json:"id"`
			Name     string    `json:"name"`
			Color    string    `json:"color"`
			Created  time.Time `json:"created"`
			Modified time.Time `json:"modified"`
		} `json:"tags"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		Domain      string      `json:"domain"`
		Created     time.Time   `json:"created"`
		Modified    time.Time   `json:"modified"`
		Meta        interface{} `json:"meta"`
	} `json:"results"`
}
