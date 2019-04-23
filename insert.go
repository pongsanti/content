package content

import (
	"context"
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func Insert(
	ctx context.Context,
	exec boil.ContextExecutor,
	content *models.Content,
) error {
	log.Print("Content:Insert")
	if content == nil {
		return errContentIsNil
	}

	return content.Insert(ctx, exec, boil.Infer())
}
