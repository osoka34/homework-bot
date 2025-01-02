package message

import (
	"time"

	"github.com/google/uuid"

	"github.com/osoka34/homework-bot/internal/domain/entity"
)

const tableName = "messages"

func (m *MessageDAO) TableName() string {
	return tableName
}

type MessageDAO struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Content  string    `gorm:"type:text"`
	SenderId int64     `gorm:"type:bigint"`
	CreateAt time.Time `gorm:"type:timestamp"`
    Username string    `gorm:"type:varchar(70)"`
	GroupId  int64     `gorm:"type:bigint"`
	Pattern  string    `gorm:"type:varchar(255)"`
}

func FromEntity(message *entity.Message) *MessageDAO {
	return &MessageDAO{
		Id:       message.Id,
		Content:  message.Content,
		SenderId: message.SenderId,
        Username: message.Username,
		CreateAt: message.CreateAt,
		GroupId:  message.GroupId,
		Pattern:  message.Pattern,
	}
}


func (m *MessageDAO) ToEntity() *entity.Message {
    return &entity.Message{
        Id:       m.Id,
        Content:  m.Content,
        SenderId: m.SenderId,
        Username: m.Username,
        CreateAt: m.CreateAt,
        GroupId:  m.GroupId,
        Pattern:  m.Pattern,
    }
}

