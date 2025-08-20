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
func HandleWeather(bot *tgbotapi.BotAPI, repo repository.Repository, weatherService *service.WeatherService, message *tgbotapi.Message, msg *messages.Messages) {
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
	if user.City == "" {
		if err := sendMessage(bot, chatID, msg.NeedCity); err != nil {
			slog.Error("failed to send message", "error", err)
		}
		return
	}
	if err := sendMessage(bot, chatID, msg.WeatherLoading); err != nil {
		slog.Error("failed to send message", "error", err)
	}
	currentWeather, err := weatherService.GetCurrentWeather(user.City)
	if err != nil {
		slog.Error("failed to get current weather", "error", err)
		if err := sendMessage(bot, chatID, msg.WeatherAPIError); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	forecast, err := weatherService.GetForecast(user.City)
	if err != nil {
		slog.Error("failed to get forecast", "error", err)
		if err := sendMessage(bot, chatID, msg.WeatherAPIError); err != nil {
			slog.Error("failed to send error message", "error", err)
		}
		return
	}
	response := formatWeatherResponse(user.City, currentWeather, forecast)
	if err := sendMessage(bot, chatID, response); err != nil {
		slog.Error("failed to send weather response", "error", err)
	}
}
func formatWeatherResponse(city string, current *service.WeatherData, forecast *service.ForecastData) string {
	var response strings.Builder
	response.WriteString(fmt.Sprintf("🌤 *Прогноз погоды для города %s*\n\n", city))
	currentTemp := int(current.Current.TempC)
	currentFeels := int(current.Current.FeelsLike)
	currentDesc := current.Current.Condition.Text
	currentEmoji := getWeatherEmoji(currentDesc)
	response.WriteString(fmt.Sprintf("%s *Сейчас:* %+d°C (ощущается как %+d°C)\n", currentEmoji, currentTemp, currentFeels))
	response.WriteString(fmt.Sprintf("💧 Влажность: %d%%\n🌪 Ветер: %.1f км/ч\n", current.Current.Humidity, current.Current.WindKph))
	response.WriteString(fmt.Sprintf("_%s_\n\n", currentDesc))
	response.WriteString("📅 *Прогноз на 5 дней:*\n")
	for i, day := range forecast.Forecast.ForecastDay {
		if i >= 5 {
			break
		}
		dayName := getDayName(i)
		maxTemp := int(day.Day.MaxTempC)
		minTemp := int(day.Day.MinTempC)
		dayDesc := day.Day.Condition.Text
		dayEmoji := getWeatherEmoji(dayDesc)
		response.WriteString(fmt.Sprintf("%s *%s:* %+d°C...%+d°C %s\n",
			dayEmoji, dayName, minTemp, maxTemp, dayDesc))
	}
	return response.String()
}
func getDayName(dayIndex int) string {
	switch dayIndex {
	case 0:
		return "Сегодня"
	case 1:
		return "Завтра"
	case 2:
		return "Послезавтра"
	default:
		return fmt.Sprintf("День %d", dayIndex+1)
	}
}
