package service

import (
	"encoding/json"
	"testing"
)

func TestWeatherDataJSONParsing(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		wantErr  bool
		wantTemp float64
		wantDesc string
	}{
		{
			name:     "valid response",
			json:     `{"current":{"temp_c":25.5,"feelslike_c":28.2,"humidity":65,"wind_kph":12.5,"condition":{"text":"Ясно"}}}`,
			wantErr:  false,
			wantTemp: 25.5,
			wantDesc: "Ясно",
		},
		{
			name:     "negative temp",
			json:     `{"current":{"temp_c":-15.7,"condition":{"text":"Снег"}}}`,
			wantErr:  false,
			wantTemp: -15.7,
			wantDesc: "Снег",
		},
		{
			name:    "invalid json",
			json:    `{"current": invalid}`,
			wantErr: true,
		},
		{
			name:     "empty",
			json:     `{}`,
			wantErr:  false,
			wantTemp: 0.0,
			wantDesc: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var weather WeatherData
			err := json.Unmarshal([]byte(tt.json), &weather)

			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if weather.Current.TempC != tt.wantTemp {
					t.Errorf("temp = %v, want %v", weather.Current.TempC, tt.wantTemp)
				}
				if weather.Current.Condition.Text != tt.wantDesc {
					t.Errorf("condition = %q, want %q", weather.Current.Condition.Text, tt.wantDesc)
				}
			}
		})
	}
}

func TestForecastDataJSONParsing(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		wantErr  bool
		wantDays int
	}{
		{
			name:     "valid forecast",
			json:     `{"forecast":{"forecastday":[{"day":{"maxtemp_c":25.0,"mintemp_c":15.0,"condition":{"text":"Ясно"}}},{"day":{"maxtemp_c":22.0,"mintemp_c":12.0,"condition":{"text":"Дождь"}}}]}}`,
			wantErr:  false,
			wantDays: 2,
		},
		{
			name:     "empty forecast",
			json:     `{"forecast":{"forecastday":[]}}`,
			wantErr:  false,
			wantDays: 0,
		},
		{
			name:    "invalid json",
			json:    `{"forecast": invalid}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var forecast ForecastData
			err := json.Unmarshal([]byte(tt.json), &forecast)

			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				days := len(forecast.Forecast.ForecastDay)
				if days != tt.wantDays {
					t.Errorf("days = %d, want %d", days, tt.wantDays)
				}
			}
		})
	}
}
