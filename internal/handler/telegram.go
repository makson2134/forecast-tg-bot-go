package handler

import (
	"log/slog"
	"tg-bot/internal/config"
	"tg-bot/internal/messages"
	"tg-bot/internal/repository"
	"tg-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandler struct {
	bot            *tgbotapi.BotAPI
	repo           repository.Repository
	weatherService *service.WeatherService
	cfg            *config.Config
	msg            *messages.Messages
}

func NewTelegramHandler(bot *tgbotapi.BotAPI, repo repository.Repository, weatherService *service.WeatherService, cfg *config.Config, msg *messages.Messages) *TelegramHandler {
	return &TelegramHandler{
		bot:            bot,
		repo:           repo,
		weatherService: weatherService,
		cfg:            cfg,
		msg:            msg,
	}
}
func (h *TelegramHandler) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = h.cfg.Telegram.Timeout
	updates := h.bot.GetUpdatesChan(u)
	slog.Info("telegram bot started, listening for updates...")
	for update := range updates {
		if update.Message != nil {
			go h.handleMessage(update.Message)
		}
	}
	return nil
}
func (h *TelegramHandler) Stop() {
	h.bot.StopReceivingUpdates()
}
func (h *TelegramHandler) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := h.bot.Send(msg)
	return err
}
