package boiling

import "fmt"

const boilingF = 212.0

func WriteTempInfo() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Temp of boiling = %g°F or %g°C\n", f, c)
}
