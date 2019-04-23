package content

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/pongsanti/content/db/models"
)

func Select(
	ctx context.Context, exec boil.ContextExecutor, qmods []qm.QueryMod) (
	models.ContentSlice, error) {
	log.Print("Content:Select")
	if qmods == nil {
		return nil, errNoQueries
	}
	return models.Contents(qmods...).All(ctx, exec)
}

func SelectCount(
	ctx context.Context, exec boil.ContextExecutor, qmods []qm.QueryMod) (
	int, error) {
	log.Print("Content:SelectCount")
	if qmods == nil {
		return 0, errNoQueries
	}
	count, err := models.Contents(qmods...).Count(ctx, exec)
	return int(count), err
}
