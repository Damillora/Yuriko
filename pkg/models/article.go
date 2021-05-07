package models

type Article struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
	Slug        string  `json:"slug"`
	Excerpt     string  `json:"excerpt"`
	HtmlMarkup  string  `json:"html"`
	PublishedAt float64 `json:"published_at"`
}

type ArticleCreateModel struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
	Slug        string  `json:"slug"`
	Excerpt     string  `json:"excerpt"`
	HtmlMarkup  string  `json:"html"`
	PublishedAt float64 `json:"published_at"`
}

type ArticleSearchResult struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Slug        string `json:"slug"`
	Excerpt     string `json:"excerpt"`
	PublishedAt string `json:"published_at"`
}
