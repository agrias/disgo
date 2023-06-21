package handler

import (
	"context"

	"github.com/disgoorg/disgo/events"
)

type Handler interface {
	Handle(ctx context.Context, e *events.InteractionCreate) error
}

type HandlerFunc func(ctx context.Context, e *events.InteractionCreate) error

func (f HandlerFunc) Handle(ctx context.Context, e *events.InteractionCreate) error {
	return f(ctx, e)
}

type (
	Middleware func(next Handler) Handler

	Middlewares []Middleware
)
