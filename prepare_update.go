package content

import (
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
)

func UpdateModel(
	m *models.Content,
	title *string,
	subtitle *string,
	detail *string) {
	log.Print("Content:UpdateModel")

	if title != nil {
		m.Title = *title
	}
	if subtitle != nil {
		m.Subtitle = null.StringFromPtr(subtitle)
	}
	if detail != nil {
		m.Detail = null.StringFromPtr(detail)
	}
}
