package bot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"

	"github.com/osoka34/homework-bot/config"
	"github.com/osoka34/homework-bot/internal/domain/entity"
	"github.com/osoka34/homework-bot/pkg/utils"
)

const (
	pattern = "#ДЗ"
)

type Bot struct {
	teleBot       *telebot.Bot
	mesRepository entity.MessegeRepositoryI
}

func NewBot(cfg *config.Config, mesRepository entity.MessegeRepositoryI) (*Bot, error) {
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
	b.teleBot.Handle(telebot.OnText, b.SaveMessage)
	b.teleBot.Handle("/users_out_of_time", b.GetUsersOutOfTime)

	b.teleBot.Start()
}

func (b *Bot) Stop() {
	b.teleBot.Stop()
}

func (b *Bot) GetUsersOutOfTime(c telebot.Context) error {
	group := c.Chat()
	messagesOut, err := b.mesRepository.GetOutOfTimeByChat(context.TODO(), pattern, group.ID)
	if err != nil {
		return err
	}

	mapUsers := make(map[string]*entity.Message, len(messagesOut))

	for _, message := range messagesOut {
		mapUsers[message.Username] = message
	}

	_, end := utils.GetLastSundayRange()

	builder := strings.Builder{}
	for user := range mapUsers {
		builder.WriteString("@")
		builder.WriteString(mapUsers[user].Username)
		builder.WriteString(" - сдал в: ")
		builder.WriteString(mapUsers[user].CreateAt.Format(time.DateTime))
		builder.WriteString(" - опоздал на: ")
		builder.WriteString(mapUsers[user].CreateAt.Sub(end).String())
		builder.WriteString("\n")
	}

	out := builder.String()

	if out == "" {
		return c.Reply("Опоздавших нет")
	}

	return c.Reply(fmt.Sprintf("Пользователи, которые сделали ДЗ c опозданием:\n%s", out))
}

func (b *Bot) SaveMessage(c telebot.Context) error {
	message := c.Message()
	group := message.Chat
	user := message.Sender

	zap.L().
		Info("SaveMessage",
			zap.String("message", message.Text),
			zap.String("group", group.Title),
			zap.String("user", user.Username),
		)

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
