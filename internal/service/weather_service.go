package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tg-bot/internal/config"
	"time"
)

type WeatherService struct {
	apiKey       string
	baseURL      string
	httpClient   *http.Client
	requestDelay time.Duration
}
type WeatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		TzId    string `json:"tz_id"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		FeelsLike float64 `json:"feelslike_c"`
		Humidity  int     `json:"humidity"`
		WindKph   float64 `json:"wind_kph"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}
type ForecastData struct {
	Location struct {
		Name string `json:"name"`
		TzId string `json:"tz_id"`
	} `json:"location"`
	Forecast struct {
		ForecastDay []struct {
			Date string `json:"date"`
			Day  struct {
				MaxTempC  float64 `json:"maxtemp_c"`
				MinTempC  float64 `json:"mintemp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
			Hour []struct {
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func NewWeatherService(cfg *config.Config) *WeatherService {
	return &WeatherService{
		apiKey:       cfg.Weather.APIKey,
		baseURL:      cfg.Weather.BaseURL,
		httpClient:   &http.Client{Timeout: cfg.HTTP.ClientTimeout},
		requestDelay: cfg.Weather.RequestDelay,
	}
}
func (w *WeatherService) ValidateCity(city string) (*WeatherData, error) {
	data, err := w.getCurrentWeather(city)
	if err != nil {
		return nil, err
	}
	time.Sleep(w.requestDelay)
	return data, nil
}
func (w *WeatherService) GetCurrentWeather(city string) (*WeatherData, error) {
	data, err := w.getCurrentWeather(city)
	if err != nil {
		return nil, err
	}
	time.Sleep(w.requestDelay)
	return data, nil
}
func (w *WeatherService) GetForecast(city string) (*ForecastData, error) {
	u := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&days=5&lang=ru",
		w.baseURL, w.apiKey, url.QueryEscape(city))
	resp, err := w.httpClient.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		return nil, fmt.Errorf("city not found")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	var forecast ForecastData
	if err := json.Unmarshal(body, &forecast); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	time.Sleep(w.requestDelay)
	return &forecast, nil
}
func (w *WeatherService) getCurrentWeather(city string) (*WeatherData, error) {
	u := fmt.Sprintf("%s/current.json?key=%s&q=%s&lang=ru",
		w.baseURL, w.apiKey, url.QueryEscape(city))
	resp, err := w.httpClient.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		return nil, fmt.Errorf("city not found")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	var weather WeatherData
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &weather, nil
}
