package message

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/osoka34/homework-bot/internal/domain/entity"
	"github.com/osoka34/homework-bot/pkg/utils"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (m *MessageRepository) CreateMessage(ctx context.Context, message *entity.Message) error {
	dao := FromEntity(message)
	return m.db.WithContext(ctx).Create(dao).Error
}

func (m *MessageRepository) GetOnTimeByChat(
	ctx context.Context,
	pattern string,
	chat int64,
) (messages []*entity.Message, err error) {
	start, end := utils.GetLastSundayRange()
	if err = m.db.WithContext(ctx).
		Where(
			"create_at BETWEEN ? AND ? AND pattern like ? AND group_id = ?",
			start,
			end,
			strings.Join([]string{"%", pattern, "%"}, ""),
			chat).
		Find(&messages).
		Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *MessageRepository) GetOutOfTimeByChat(
	ctx context.Context,
	pattern string,
	chat int64,
) (messages []*entity.Message, err error) {
	start, end := utils.GetLastSundayRange()
	if err = m.db.WithContext(ctx).
		Where(
			"create_at BETWEEN ? AND ? AND pattern = ? AND group_id = ?",
			end,
			start.Add(7*24*time.Hour),
			pattern,
			chat).
		Find(&messages).
		Error; err != nil {
		return nil, err
	}
	return messages, nil
}
