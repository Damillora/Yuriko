package app

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Damillora/Yuriko/pkg/config"
	"github.com/Damillora/Yuriko/pkg/models"
	"github.com/Damillora/Yuriko/pkg/services"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(g *gin.Engine) {
	api := g.Group("/api")
	api.GET("/article/search", SearchArticles)
	api.POST("/webhook", Webhook)
}

func SearchArticles(c *gin.Context) {
	q := c.Query("q")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println("Page not string")
	}
	result, err2 := services.SearchArticles(q, page)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func Webhook(c *gin.Context) {
	webhookKey := c.Query("key")
	if webhookKey != config.CurrentConfig.WebhookApiKey {
		c.String(http.StatusUnauthorized, "no api key")
		return
	}
	var body models.WebhookBody
	err := c.BindJSON(&body)
	if err == nil {
		time, err := time.Parse(time.RFC3339, body.Post.Current.PublishedAt)
		if err != nil {
			log.Println(err.Error())
			return
		}
		newArticle := models.ArticleCreateModel{
			ID:          body.Post.Current.ID,
			Title:       body.Post.Current.Title,
			Url:         body.Post.Current.Url,
			Slug:        body.Post.Current.Slug,
			Excerpt:     body.Post.Current.Excerpt,
			HtmlMarkup:  body.Post.Current.HtmlMarkup,
			PrimaryTag:  body.Post.Current.PrimaryTag,
			PublishedAt: float64(time.Unix()),
		}
		services.ImportArticle(newArticle)
		c.String(http.StatusOK, "ok")
	} else {
		log.Println(err.Error())
	}
}
