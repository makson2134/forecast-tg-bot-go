package main
import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tg-bot/internal/config"
	"tg-bot/internal/database"
	"tg-bot/internal/handler"
	"tg-bot/internal/messages"
	"tg-bot/internal/repository"
	"tg-bot/internal/service"
	"tg-bot/pkg/logger"
	"tg-bot/pkg/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	logger.Initialize(cfg)
	slog.Info("config loaded successfully")
	db, err := sql.Open("postgres", utils.GetDSN(cfg))
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}
	slog.Info("database connected successfully")
	if err := database.RunMigrations(db); err != nil {
		panic("failed to run migrations: " + err.Error())
	}
	slog.Info("migrations applied successfully")
	repo := repository.New(db)
	weatherService := service.NewWeatherService(cfg)
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		panic("failed to create telegram bot: " + err.Error())
	}
	if cfg.Environment == "development" {
		bot.Debug = true
	}
	slog.Info("telegram bot authorized", "username", bot.Self.UserName)
	telegramHandler := handler.NewTelegramHandler(bot, repo, weatherService, cfg, &messages.BotMessages)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		slog.Info("starting telegram bot...")
		if err := telegramHandler.Start(); err != nil {
			slog.Error("telegram bot error", "error", err)
			cancel()
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		slog.Info("shutting down gracefully...")
		telegramHandler.Stop()
		cancel()
	}()
	slog.Info("application started")
	<-ctx.Done()
	slog.Info("application stopped")
}
