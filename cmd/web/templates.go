package main
import "snippetbox.xuche.net/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}