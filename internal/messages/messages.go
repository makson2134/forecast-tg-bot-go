package messages
import "fmt"
type Messages struct {
	DatabaseError   string
	WeatherAPIError string
	CheckingCity   string
	WeatherLoading string
	CityUpdatedTmpl string
	SetCityHelp      string
	NeedCity         string
	NeedRegistration string
	UnknownCommand string
	NotACommand    string
	CityNotFound   string
	FirstTimeWelcome      string
	CommandsList          string
	UserAlreadyExistsTmpl string
}
var BotMessages = Messages{
	DatabaseError:   "Произошла внутренняя ошибка. Пожалуйста, попробуйте повторить Ваш запрос позже.",
	WeatherAPIError: "Не удалось получить данные от сервиса погоды. Пожалуйста, попробуйте позже.",
	CheckingCity:   "Проверяю информацию о городе...",
	WeatherLoading: "Получаю прогноз погоды для Вас...",
	CityUpdatedTmpl: "Ваш город успешно изменен на *%s*.\n\n%s\n\nЧтобы получить прогноз на неделю, используйте команду /weather.",
	SetCityHelp:      "Пожалуйста, укажите название города после команды.\nНапример: /setcity Москва",
	NeedCity:         "Пожалуйста, сначала укажите Ваш город с помощью команды /setcity.\nНапример: /setcity Москва",
	NeedRegistration: "Похоже, Вы еще не начали работу с ботом. Пожалуйста, используйте команду /start.",
	UnknownCommand: "Неизвестная команда. Пожалуйста, используйте /info, чтобы посмотреть список доступных команд.",
	NotACommand:    "Я понимаю только команды, которые начинаются с символа '/'. Пожалуйста, используйте /info для просмотра списка команд.",
	CityNotFound:   "К сожалению, город с таким названием не найден. Пожалуйста, проверьте правильность написания и попробуйте снова.",
	FirstTimeWelcome: `*Добро пожаловать!*
Я погодный бот. Чтобы начать работу, пожалуйста, укажите Ваш город.
Например: /setcity Санкт-Петербург
Для получения полного списка команд используйте /info.`,
	CommandsList: `*Список доступных команд:*
/start - Начало работы с ботом.
/setcity [название города] - Установить или изменить Ваш город.
/weather - Получить прогноз погоды для Вашего города.
/info - Показать это справочное сообщение.`,
	UserAlreadyExistsTmpl: "Рад Вас снова видеть, %s! Вы уже зарегистрированы. Ваш город: *%s*.",
}
func (m *Messages) CityUpdated(cityName, weatherInfo string) string {
	return fmt.Sprintf(m.CityUpdatedTmpl, cityName, weatherInfo)
}
func (m *Messages) UserAlreadyExists(userName, cityName string) string {
	if cityName == "" {
		cityName = "не указан. Пожалуйста, укажите его с помощью команды /setcity"
	}
	return fmt.Sprintf(m.UserAlreadyExistsTmpl, userName, cityName)
}
