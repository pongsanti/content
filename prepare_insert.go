package content

import (
	"log"
	"time"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
)

func CreateModel(
	contentType string,
	title string,
	subtitle *string,
	detail *string,
	startAt *time.Time,
	endAt *time.Time,
	status *string,
	xtime1 *time.Time,
	xtext1 *string,
) *models.Content {
	log.Print("Content:CreateModel")

	return &models.Content{
		ContentType: contentType,
		Title:       title,
		Subtitle:    null.StringFromPtr(subtitle),
		Detail:      null.StringFromPtr(detail),
		StartAt:     null.TimeFromPtr(startAt),
		EndAt:       null.TimeFromPtr(endAt),
		Status:      null.StringFromPtr(status),
		Xtime1:      null.TimeFromPtr(xtime1),
		Xtext1:      null.StringFromPtr(xtext1),
	}
}
