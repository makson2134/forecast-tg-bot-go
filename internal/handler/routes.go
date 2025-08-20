package handler
import (
	"log/slog"
	"strings"
	"tg-bot/internal/handler/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func (h *TelegramHandler) handleMessage(message *tgbotapi.Message) {
	userID := message.From.ID
	chatID := message.Chat.ID
	text := message.Text
	slog.Info("received message",
		"user_id", userID,
		"chat_id", chatID,
		"text", text)
	if message.IsCommand() {
		h.handleCommand(message)
		return
	}
	if err := h.sendMessage(chatID, h.msg.NotACommand); err != nil {
		slog.Error("failed to send message", "error", err, "chat_id", chatID)
	}
}
func (h *TelegramHandler) handleCommand(message *tgbotapi.Message) {
	command := message.Command()
	args := strings.Fields(message.CommandArguments())
	switch command {
	case "start":
		commands.HandleStart(h.bot, h.repo, message, h.msg)
	case "setcity":
		commands.HandleSetCity(h.bot, h.repo, h.weatherService, message, args, h.msg)
	case "weather":
		commands.HandleWeather(h.bot, h.repo, h.weatherService, message, h.msg)
	case "info":
		h.handleInfo(message)
	default:
		if err := h.sendMessage(message.Chat.ID, h.msg.UnknownCommand); err != nil {
			slog.Error("failed to send message", "error", err, "chat_id", message.Chat.ID)
		}
	}
}
func (h *TelegramHandler) handleInfo(message *tgbotapi.Message) {
	if err := h.sendMessage(message.Chat.ID, h.msg.CommandsList); err != nil {
		slog.Error("failed to send info message", "error", err, "chat_id", message.Chat.ID)
	}
}
