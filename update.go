package content

import (
	"context"
	"errors"
	"log"

	"github.com/pongsanti/content/db/models"

	"github.com/volatiletech/sqlboiler/boil"
)

func Update(
	ctx context.Context,
	exec boil.ContextExecutor,
	content *models.Content,
	columns boil.Columns,
) (int, error) {
	log.Print("Content:Update")
	if content == nil {
		return 0, errors.New(contentIsNil)
	}

	count, err := content.Update(ctx, exec, columns)
	return int(count), err
}
