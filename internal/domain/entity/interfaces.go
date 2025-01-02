package entity

import (
	"context"
)

type MessegeRepositoryI interface {
	CreateMessage(ctx context.Context, message *Message) error
	GetOnTimeByChat(
		ctx context.Context,
		pattern string,
		chat int64,
	) (messages []*Message, err error)
	GetOutOfTimeByChat(
		ctx context.Context,
		pattern string,
		chat int64,
	) (messages []*Message, err error)
}
