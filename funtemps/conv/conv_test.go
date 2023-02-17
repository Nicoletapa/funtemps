package conv

import (
	//"reflect"
	"math"
	"testing"
)

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134.01, want: 56.67},
	}
	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

/*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/


// De andre testfunksjonene implementeres her
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want float64
	}

	tests := []test {
		{input:100 , want: 212.00},
	}
	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}
	
func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want float64
	}

	tests := []test {
		{input:100 , want: -279.67},
	}
	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

	

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want float64
	}

	tests := []test {
		{input:100 , want: 373.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestKelvinToCelsiust(t*testing.T) {
	type test struct {
		input float64
		want float64
	}

	tests := []test {
		{input:100 , want: -173.15},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}


func TestFahrenheitToKelvint(t*testing.T) {
	type test struct {
		input float64
		want float64
	}

	tests := []test {
		{input:100.01 , want: 310.93},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

