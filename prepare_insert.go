package content

import (
	"log"
	"time"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
)

func CreateModel(
	title string,
	subtitle *string,
	detail *string,
	startAt *time.Time,
	endAt *time.Time,
) *models.Content {
	log.Print("Content:CreateModel")

	return &models.Content{
		Title:    title,
		Subtitle: null.StringFromPtr(subtitle),
		Detail:   null.StringFromPtr(detail),
		StartAt:  null.TimeFromPtr(startAt),
		EndAt:    null.TimeFromPtr(endAt),
	}
}
