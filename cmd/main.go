package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/osoka34/homework-bot/config"
	"github.com/osoka34/homework-bot/internal/infrastructure/bot"
	"github.com/osoka34/homework-bot/internal/infrastructure/postgres/message"
	"github.com/osoka34/homework-bot/pkg/storage"
	"github.com/osoka34/homework-bot/pkg/utils"
)

func main() {
	l, err := utils.InitJSONLogger()
	if err != nil {
		panic(err)
	}

	defer l.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("error reading config file", zap.Error(err))
	}

	l.Info("Config loaded", zap.Any("config", cfg))

	db, err := storage.InitPostgres(&cfg.Postgres)
	if err != nil {
		l.Fatal("failed to connect to PostgreSQL", zap.Error(err))
	}

	mesRepository := message.NewMessageRepository(db)
	bot, err := bot.NewBot(cfg, mesRepository)
	if err != nil {
		l.Fatal("failed to create bot", zap.Error(err))
	}

	go bot.Start()

	l.Info("Bot started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	l.Info("Shutting down...")

	bot.Stop()

	sqlDB, err := db.DB()
	if err == nil {
		if closeErr := sqlDB.Close(); closeErr != nil {
			l.Error("Error closing database connection", zap.Error(closeErr))
		}
	}
}
