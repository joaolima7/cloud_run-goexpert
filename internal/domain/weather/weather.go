package weather

import "errors"

type Weather struct {
	CelsiusWeather    float32
	FahrenheitWeather float32
	KelvinWeather     float32
}

func NewWeather(celsiusWeather float32, fahrenheitWeather float32, kelvinWeather float32) (*Weather, error) {
	weather := &Weather{
		CelsiusWeather:    celsiusWeather,
		FahrenheitWeather: fahrenheitWeather,
		KelvinWeather:     kelvinWeather,
	}

	err := weather.Validate()
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (w *Weather) Validate() error {
	if w.CelsiusWeather == 0 || w.FahrenheitWeather == 0 || w.KelvinWeather == 0 {
		return errors.New("o clima deve conter todos os campos")
	}

	return nil
}
