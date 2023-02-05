package main

type WebhookBody struct {
	Content    string   `json:"content"`
	Embeds     []Embed  `json:"embeds"`
}

type Embed struct {
	Image Image `json:"image"`
}

type Image struct {
	URL string `json:"url"`
}

