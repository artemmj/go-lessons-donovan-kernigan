package echo4

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "skip newline char")
var sep = flag.String("s", " ", "separator")

func Echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
