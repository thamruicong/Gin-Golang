package models

import "testing"

// TestGetAllArticles verifies that the getAllArticles function correctly retrieves
// all articles from the articleList. It checks that the number of articles returned
// matches the expected length and that each article's ID, Title, and Content match
// the expected values in the articleList.
func TestGetAllArticles(t *testing.T) {
	articles := GetAllArticles()

	// Check length of articles is the same
	if len(articles) != len(articleList) {
		t.Errorf("Expected %d articles, got %d", len(articleList), len(articles))
	}

	// Check each article's ID, Title, and Content
	for i, article := range articles {
		expected := articleList[i]
		if article.ID != expected.ID ||
			article.Title != expected.Title ||
			article.Content != expected.Content {
			t.Errorf("In article %d, Expected %+v, got %+v", i, expected, article)
			break
		}
	}
}