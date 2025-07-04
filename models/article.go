package models

import (
	"errors"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Content of Article 1"},
	{ID: 2, Title: "Article 2", Content: "Content of Article 2"},
}

func GetAllArticles() []Article {
	return articleList
}

func SetAllArticles(articles []Article) {
	articleList = articles
}

func GetArticleByID(id int) (*Article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found")
}