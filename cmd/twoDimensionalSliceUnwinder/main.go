package main

import (
	"github.com/endevour-code-writer/twoDimensionalSliceUnwinder/internal/twoDimensionalSliceUnwinder"
	"os"
)

func main() {
	args := os.Args[1:]
	twoDimensionalSliceUnwinder.unwind(args)
}
