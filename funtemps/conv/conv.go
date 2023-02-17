package conv


/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius     Celsius = (Fahrenheit - 32)*(5/9)
    CelsiusToFahrenheit     Fahrenheit = Celsius*(9/5) + 32
    KelvinToFahrenheit      Fahrenheit = (Kelvin - 273.15)*(9/5) + 32
    CelsiusToKelvin         Kelvin = Celsius + 273.15
    KelvinToCelsius         Celsius = Kelvin - 273.15
    FahrenheitToKelvin     Kelvin = (Fahrenheit - 32) * (5/9) + 273.15
*/

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	
	return (value - 32) * (5.0 / 9.0)
}

// De andre konverteringsfunksjonene implementere her
// ...

func CelsiusToFahrenheit(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	
	return value*(9.0/5.0) + 32
}

func KelvinToFahrenheit(value float64) float64 {

	return (value-273.15)*(9.0 / 5.0) + 32
}

func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

func KelvinToCelsius(value float64) float64 {

	return  value - 273.15
}

func FahrenheitToKelvin(value float64) float64 {
	
	return (value-32)*(5.0/9.0) + 273.15
}
