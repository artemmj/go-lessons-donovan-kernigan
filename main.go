package main

import (
	"fmt"
	conv "go-lessons-donovan-kernigan/ch2/converter"
)

func main() {
	fmt.Printf("100 meters = %.2f foots\n", conv.MetersToFoots(100))
	fmt.Printf("500 meters = %.2f foots\n", conv.MetersToFoots(500))
	fmt.Printf("999 meters = %.2f foots\n", conv.MetersToFoots(999))

	fmt.Printf("100 foots = %.2f meters\n", conv.FootsToMeters(100))
	fmt.Printf("500 foots = %.2f meters\n", conv.FootsToMeters(500))
	fmt.Printf("999 foots = %.2f meters\n", conv.FootsToMeters(999))
}
