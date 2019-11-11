package main

import (
	"fmt"

	"github.com/botviet/vibo/utility"
)

func main() {

	fmt.Println(utility.Transform("nghiêng", utility.UNICODE, utility.VIQR))
	fmt.Println(utility.Transform("ngu+o+`i", utility.VIQR, utility.UNICODE))
}
