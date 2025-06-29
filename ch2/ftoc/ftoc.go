package ftoc

import "fmt"

func WriteTempInfo2() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, FahrToCels(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, FahrToCels(boilingF))

}

func FahrToCels(f float64) float64 {
	return (f - 32) * 5 / 9
}
