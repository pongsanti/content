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

	var queryMods []qm.QueryMod
	if qmods != nil {
		queryMods = qmods
	}
	return models.Contents(queryMods...).All(ctx, exec)
}

func SelectOne(
	ctx context.Context, exec boil.ContextExecutor, qmods []qm.QueryMod) (
	*models.Content, error) {
	log.Print("Content:SelectOne")

	var queryMods []qm.QueryMod
	if qmods != nil {
		queryMods = qmods
	}
	return models.Contents(queryMods...).One(ctx, exec)
}

func SelectCount(
	ctx context.Context, exec boil.ContextExecutor, qmods []qm.QueryMod) (
	int, error) {
	log.Print("Content:SelectCount")

	var queryMods []qm.QueryMod
	if qmods != nil {
		queryMods = qmods
	}

	count, err := models.Contents(queryMods...).Count(ctx, exec)
	return int(count), err
}
