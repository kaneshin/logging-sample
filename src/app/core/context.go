package core

import (
	"context"

	"github.com/kaneshin/logging-sample/src/track"
)

const (
	tokenKey = "co.kaneshin.token"
	eventKey = "co.kaneshin.event"
)

func ContextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func ContextToken(ctx context.Context) string {
	if v, ok := ctx.Value(tokenKey).(string); ok {
		return v
	}
	return ""
}

func ContextWithEvent(ctx context.Context, d *track.EventData) context.Context {
	list := ContextEvent(ctx)
	list = append(list, d)
	return context.WithValue(ctx, eventKey, list)
}

func ContextEvent(ctx context.Context) []*track.EventData {
	if v, ok := ctx.Value(eventKey).([]*track.EventData); ok {
		return v
	}
	return nil
}
