package utils

func ConvertCelsiusToFahrenheit(celsius float32) float32 {
	return (celsius * 1.8) + 32
}

func ConvertCelsiusToKelvin(celsius float32) float32 {
	return celsius + 273
}
