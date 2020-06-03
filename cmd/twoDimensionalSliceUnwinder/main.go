package main

import (
	"fmt"
	"github.com/endevour-code-writer/twoDimensionalSliceUnwinder/internal/twoDimensionalSliceUnwinder"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(twoDimensionalSliceUnwinder.Unwind(args))
}
