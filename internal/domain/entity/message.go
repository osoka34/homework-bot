package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id       uuid.UUID
	Content  string
	SenderId int64
    Username string
	CreateAt time.Time
	GroupId  int64
	Pattern  string
}
