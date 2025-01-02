package bot

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"gopkg.in/telebot.v4"

	"github.com/osoka34/homework-bot/config"
	"github.com/osoka34/homework-bot/internal/domain/entity"
)

const (
	pattern = "#ДЗ"
)

type Bot struct {
	teleBot       *telebot.Bot
	mesRepository entity.MessegeRepositoryI
}

func NewBot(cfg config.Config, mesRepository entity.MessegeRepositoryI) (*Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &Bot{
		teleBot:       bot,
		mesRepository: mesRepository,
	}, nil
}

func (b *Bot) Start() {
	b.teleBot.Start()

	b.teleBot.Handle(telebot.OnText, b.SaveMessage)
	b.teleBot.Handle("/users_on_time", b.GetUsersOnTime)
}

func (b *Bot) Stop() {
	b.teleBot.Stop()
}

func (b *Bot) GetUsersOnTime(c telebot.Context) error {
	group := c.Chat()
	messages, err := b.mesRepository.GetOnTimeByChat(context.TODO(), pattern, group.ID)
	if err != nil {
		return err
	}

	mapUsers := make(map[string]struct{}, len(messages))

	for _, message := range messages {
		mapUsers[message.Username] = struct{}{}
	}

	users := make([]string, 0, len(mapUsers))
	for user := range mapUsers {
		users = append(users, user)
	}

	return c.Send("Пользователи, которые сделали дз: %v", users)
}

func (b *Bot) SaveMessage(c telebot.Context) error {
	message := c.Message()
	group := message.Chat
	user := message.Sender

	if strings.Contains(message.Text, pattern) {
		entry := entity.Message{
			Id:       uuid.New(),
			Content:  message.Text,
			SenderId: user.ID,
			Username: user.Username,
			GroupId:  group.ID,
			Pattern:  pattern,
			CreateAt: message.Time(),
		}

		if err := b.mesRepository.CreateMessage(context.TODO(), &entry); err != nil {
			return c.Reply("Ошибка при сохранении сообщения.")
		}

		return c.Reply("Сообщение сохранено!")
	}
	return nil
}
