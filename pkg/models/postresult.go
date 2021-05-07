package models

type WebhookBody struct {
	Post WebhookPostBody `json:"post"`
}
type WebhookPostBody struct {
	Current WebhookPost `json:"current"`
}

type WebhookPost struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Slug        string `json:"slug"`
	Excerpt     string `json:"excerpt"`
	HtmlMarkup  string `json:"html"`
	PublishedAt string `json:"published_at"`
}
