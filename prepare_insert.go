package content

import (
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
)

func CreateModel(
	title string,
	subtitle *string,
	detail *string,
) *models.Content {
	log.Print("Content:CreateModel")

	return &models.Content{
		Title:    title,
		Subtitle: null.StringFromPtr(subtitle),
		Detail:   null.StringFromPtr(detail),
	}
}
