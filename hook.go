package content

import (
	"context"
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func createBeforeInsertHookFunc(contextKey uint) func(context.Context, boil.ContextExecutor, *models.Content) error {
	return func(ctx context.Context, exec boil.ContextExecutor, content *models.Content) error {
		creatorID, ok := ctx.Value(contextKey).(int)
		if ok {
			content.CreatorID = null.IntFrom(creatorID)
		}
		return nil
	}
}

/*
	RegisterBeforeInsertHook receives contextKey of the creator id
*/
func RegisterBeforeInsertHook(contextKey uint) {
	log.Print("Content:RegisterBeforeInsertHook")
	models.AddContentHook(boil.BeforeInsertHook, createBeforeInsertHookFunc(contextKey))
}
