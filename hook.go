package content

import (
	"context"
	"log"

	"github.com/pongsanti/content/db/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

type contentCreatorContextKey uint

const ContentCreatorCtxKey contentCreatorContextKey = 0

func beforeInsertHookFunc(ctx context.Context, exec boil.ContextExecutor, content *models.Content) error {
	creatorID, ok := ctx.Value(ContentCreatorCtxKey).(int)
	if ok {
		content.CreatorID = null.IntFrom(creatorID)
	}
	return nil
}

/*
	RegisterBeforeInsertHook receives contextKey of the creator id
*/
func RegisterBeforeInsertHook() {
	log.Print("Content:RegisterBeforeInsertHook")
	models.AddContentHook(boil.BeforeInsertHook, beforeInsertHookFunc)
}
