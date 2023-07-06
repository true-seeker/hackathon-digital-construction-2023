package weather

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/widget/weather"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServiceYandex struct {
}

func NewWeatherYandexService() *ServiceYandex {
	return &ServiceYandex{}
}

type YandexWeather struct {
	Forecasts []Forecast `json:"forecasts,omitempty"`
	Fact      Fact       `json:"fact"`
	Info      Info       `json:"info"`
}
type Forecast struct {
	Parts Part   `json:"parts,omitempty"`
	Date  string `json:"date,omitempty"`
}

type Info struct {
	DefPRessureMm int `json:"def_pressure_mm"`
}

type Fact struct {
	Temp      int    `json:"temp"`
	FeelsLike int    `json:"feels_like"`
	Condition string `json:"condition"`
}

type Part struct {
	DayShort DayShort `json:"day_short"`
}

type DayShort struct {
	Temp      int    `json:"temp,omitempty"`
	Condition string `json:"condition"`
}

func (s *ServiceYandex) GetWeather(locationData weather.LocationData) (*entities.Weather, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%f&lon=%f", locationData.Latitude, locationData.Longitude), nil)
	req.Header.Set("X-Yandex-API-Key", "ad8382a6-0730-45fe-a069-7285cedf1dd7")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var yandexWeather YandexWeather
	err = json.Unmarshal(body, &yandexWeather)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	var w entities.Weather

	for _, date := range yandexWeather.Forecasts {
		w.Forecast = append(w.Forecast, entities.Temperature{
			Value:     date.Parts.DayShort.Temp,
			Date:      date.Date,
			Condition: date.Parts.DayShort.Condition,
		})
	}

	w.TemperatureNow = entities.Temperature{
		Value: yandexWeather.Fact.Temp,
	}
	w.FeelsLike = yandexWeather.Fact.FeelsLike
	w.Condition = yandexWeather.Fact.Condition
	w.Pressure = yandexWeather.Info.DefPRessureMm

	return &w, nil
}
