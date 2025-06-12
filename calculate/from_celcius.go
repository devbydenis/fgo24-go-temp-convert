package calculate

func FromCelciusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func FromCelciusToKelvin(c float64) float64 {
	return c + 273.15
}

func FromCelciusToReamur(c float64) float64 {
	return c * 4 / 5
}
