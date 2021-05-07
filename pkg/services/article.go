package services

import (
	"log"
	"time"

	"github.com/Damillora/Yuriko/pkg/database"
	"github.com/Damillora/Yuriko/pkg/models"
	"github.com/typesense/typesense-go/typesense/api"
)

type SearchResult struct {
	Page    int                          `json:"page"`
	PerPage int                          `json:"per_page"`
	Result  []models.ArticleSearchResult `json:"result"`
}

func SearchArticles(q string, page int) (SearchResult, error) {
	perPage := 10
	searchParameters := &api.SearchCollectionParams{
		Q:       q,
		QueryBy: []string{"title", "slug", "excerpt", "html"},
		Page:    &page,
		PerPage: &perPage,
	}
	searchResult, err := database.Client.Collection("articles").Documents().Search(searchParameters)
	if err != nil {
		log.Println("Search error: " + err.Error())
		return SearchResult{}, err
	}
	var articles []models.ArticleSearchResult

	for _, hit := range searchResult.Hits {
		article := models.ArticleSearchResult{
			Title:       hit.Document["title"].(string),
			Url:         hit.Document["url"].(string),
			Slug:        hit.Document["slug"].(string),
			Excerpt:     hit.Document["excerpt"].(string),
			PublishedAt: time.Unix(int64(hit.Document["published_at"].(float64)), 0).String(),
		}
		articles = append(articles, article)
	}

	return SearchResult{
		Page:    page,
		PerPage: perPage,
		Result:  articles,
	}, nil
}

func ImportArticle(article models.ArticleCreateModel) {
	document := article

	database.Client.Collection("articles").Documents().Upsert(document)

}
