package commands
import (
	"log/slog"
	"tg-bot/internal/messages"
	"tg-bot/internal/models"
	"tg-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func HandleStart(bot *tgbotapi.BotAPI, repo repository.Repository, message *tgbotapi.Message, msg *messages.Messages) {
	userID := message.From.ID
	username := message.From.UserName
	chatID := message.Chat.ID
	user, err := repo.GetUserByTelegramID(userID)
	if err != nil {
		slog.Error("failed to get user", "error", err)
		if err := sendMessage(bot, chatID, msg.DatabaseError); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	if user != nil {
		responseText := msg.UserAlreadyExists(username, user.City)
		if err := sendMessage(bot, chatID, responseText); err != nil {
			slog.Error("failed to send already exists message", "error", err)
		}
		return
	}
	newUser := &models.User{
		TelegramID: userID,
		Username:   username,
		City:       "", 
	}
	err = repo.CreateUser(newUser)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		if err := sendMessage(bot, chatID, msg.DatabaseError); err != nil {
			slog.Error("failed to send registration error message", "error", err)
		}
		return
	}
	if err := sendMessage(bot, chatID, msg.FirstTimeWelcome); err != nil {
		slog.Error("failed to send welcome message", "error", err)
	}
}
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(msg)
	return err
}
