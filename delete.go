package content

import (
	"context"
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func Delete(ctx context.Context,
	exec boil.ContextExecutor,
	content *models.Content) (int, error) {
	log.Print("Content:Delete")
	if content == nil {
		return 0, errContentIsNil
	}

	count, err := content.Delete(ctx, exec)
	return int(count), err
}
