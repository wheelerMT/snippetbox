package main

import "snippetbox.wheelermt.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
