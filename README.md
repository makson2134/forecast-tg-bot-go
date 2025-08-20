# Forecast-tg-bot

Бот для получения прогноза погоды

## Стек

*   **Go 1.24.4**
*   **go-telegram-bot-api/v5**: Библиотека для взаимодействия с Telegram Bot API.
*   **slog**: Стандартная библиотека для структурированных логов.
*   **lib/pq**: Драйвер для подключения к PostgreSQL.
*   **golang-migrate**: Инструмент для управления миграциями схемы базы данных.
*   **cleanenv**: Библиотека для считывания переменных окружения и анмаршаллинга.

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/makson2134/Forecast-tg-bot-go
```

### 2. Настройка окружения

Конфигурацией сервиса управляет файл `.env`. Формат .env файла:

```env
# Токен бота
BOT_TOKEN=

# Ключ от weather api
WEATHER_API_KEY=
WEATHER_BASE_URL=http://api.weatherapi.com/v1
WEATHER_REQUEST_DELAY=1s


LOG_LEVEL=info
LOG_FORMAT=json

DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_NAME=weather_bot
DB_PASSWORD=postgres

TELEGRAM_TIMEOUT=60
HTTP_CLIENT_TIMEOUT=10s

ENVIRONMENT=development
```

Для получения необходимых токенов:
*   `BOT_TOKEN` - получить у [**@BotFather**](https://t.me/BotFather)
*   `WEATHER_API_KEY` - получить на [**WeatherAPI.com**](https://www.weatherapi.com/)

### 3. Запуск

Запуск производится с помощью запуска docker compose в корне проекта

```bash
docker compose up
```

Чтобы остановить приложение:
```bash
docker compose down
```


