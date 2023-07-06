package weather

import (
	"backend/internal/domain/entities"
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

func (s *ServiceYandex) GetWeather() (*entities.Weather, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.weather.yandex.ru/v2/forecast?lat=55.75396&lon=37.620393", nil)
	req.Header.Set("X-Yandex-API-Key", "ad8382a6-0730-45fe-a069-7285cedf1dd7")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var smth1 YandexWeather
	err = json.Unmarshal(body, &smth1)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	var weather entities.Weather

	for _, date := range smth1.Forecasts {
		weather.Forecast = append(weather.Forecast, entities.Temperature{
			Value:     date.Parts.DayShort.Temp,
			Date:      date.Date,
			Condition: date.Parts.DayShort.Condition,
		})
	}

	weather.TemperatureNow = entities.Temperature{
		Value: smth1.Fact.Temp,
	}
	weather.FeelsLike = smth1.Fact.FeelsLike
	weather.Condition = smth1.Fact.Condition
	weather.Pressure = smth1.Info.DefPRessureMm

	return &weather, nil
}
