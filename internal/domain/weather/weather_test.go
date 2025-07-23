package weather

import (
	"testing"
)

func TestNewWeather_Success(t *testing.T) {
	weather, err := NewWeather(25.0, 77.0, 298.15)
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}
	if weather == nil {
		t.Fatal("esperava instância de Weather, mas recebeu nil")
	}
}

func TestNewWeather_MissingFields(t *testing.T) {
	tests := []struct {
		name    string
		celsius float32
		fahr    float32
		kelvin  float32
	}{
		{"Celsius zero", 0, 77, 298.15},
		{"Fahrenheit zero", 25, 0, 298.15},
		{"Kelvin zero", 25, 77, 0},
		{"Todos zero", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewWeather(tc.celsius, tc.fahr, tc.kelvin)
			if err == nil {
				t.Errorf("esperava erro para %s, mas recebeu nil", tc.name)
			}
		})
	}
}
