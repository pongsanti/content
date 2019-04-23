package response

import (
	"log"
	"time"

	"github.com/pongsanti/content/db/models"
)

type Content struct {
	ID       int
	Title    string
	Subtitle *string
	Detail   *string
	StartAt  *time.Time
	EndAt    *time.Time
}

func NewContent(c *models.Content) *Content {
	log.Print("Content:NewContent")
	if c == nil {
		return nil
	}

	return &Content{
		ID:       c.ID,
		Title:    c.Title,
		Subtitle: c.Subtitle.Ptr(),
		Detail:   c.Detail.Ptr(),
		StartAt:  c.StartAt.Ptr(),
		EndAt:    c.EndAt.Ptr(),
	}
}

func NewContents(cts models.ContentSlice) []*Content {
	log.Print("Content:NewContents")
	if cts == nil {
		return nil
	}

	res := make([]*Content, len(cts), len(cts))
	for index, ct := range cts {
		res[index] = NewContent(ct)
	}
	return res
}
