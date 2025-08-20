package commands
import (
	"fmt"
	"log/slog"
	"strings"
	"tg-bot/internal/messages"
	"tg-bot/internal/repository"
	"tg-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func HandleSetCity(bot *tgbotapi.BotAPI, repo repository.Repository, weatherService *service.WeatherService, message *tgbotapi.Message, args []string, msg *messages.Messages) {
	if len(args) == 0 {
		if err := sendMessage(bot, message.Chat.ID, msg.SetCityHelp); err != nil {
			slog.Error("failed to send message", "error", err)
		}
		return
	}
	userID := message.From.ID
	chatID := message.Chat.ID
	user, err := repo.GetUserByTelegramID(userID)
	if err != nil {
		slog.Error("failed to get user", "error", err)
		if err := sendMessage(bot, chatID, msg.DatabaseError); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	if user == nil {
		if err := sendMessage(bot, chatID, msg.NeedRegistration); err != nil {
			slog.Error("failed to send message", "error", err)
		}
		return
	}
	city := strings.Join(args, " ")
	if err := sendMessage(bot, chatID, msg.CheckingCity); err != nil {
		slog.Error("failed to send message", "error", err)
	}
	weatherData, err := weatherService.ValidateCity(city)
	if err != nil {
		var responseText string
		if strings.Contains(err.Error(), "city not found") {
			responseText = msg.CityNotFound
		} else {
			slog.Error("weather API error", "error", err)
			responseText = msg.WeatherAPIError
		}
		if err := sendMessage(bot, chatID, responseText); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	correctCityName := weatherData.Location.Name
	err = repo.UpdateUserCity(userID, correctCityName)
	if err != nil {
		slog.Error("failed to update user city", "error", err)
		if err := sendMessage(bot, chatID, msg.DatabaseError); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	currentWeather := formatCurrentWeather(weatherData)
	responseText := msg.CityUpdated(correctCityName, currentWeather)
	if err := sendMessage(bot, chatID, responseText); err != nil {
		slog.Error("failed to send message", "error", err)
	}
}
func formatCurrentWeather(data *service.WeatherData) string {
	temp := int(data.Current.TempC)
	feelsLike := int(data.Current.FeelsLike)
	description := data.Current.Condition.Text
	emoji := getWeatherEmoji(description)
	return fmt.Sprintf("%s *Сейчас:* %+d°C (ощущается как %+d°C)\n💧 Влажность: %d%%\n🌪 Ветер: %.1f км/ч\n%s",
		emoji, temp, feelsLike, data.Current.Humidity, data.Current.WindKph, description)
}
func getWeatherEmoji(description string) string {
	desc := strings.ToLower(description)
	switch {
	case strings.Contains(desc, "ясно") || strings.Contains(desc, "солнечно"):
		return "☀️"
	case strings.Contains(desc, "облачно") || strings.Contains(desc, "пасмурно"):
		return "☁️"
	case strings.Contains(desc, "дождь") || strings.Contains(desc, "ливень"):
		return "🌧"
	case strings.Contains(desc, "снег"):
		return "❄️"
	case strings.Contains(desc, "гроза") || strings.Contains(desc, "буря"):
		return "⛈"
	case strings.Contains(desc, "туман") || strings.Contains(desc, "дымка"):
		return "🌫"
	case strings.Contains(desc, "переменная"):
		return "🌤"
	default:
		return "🌤"
	}
}
