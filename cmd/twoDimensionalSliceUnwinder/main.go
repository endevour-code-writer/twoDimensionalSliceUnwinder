package main

import (
	"fmt"
	"github.com/endevour-code-writer/twoDimensionalSliceUnwinder/internal/twoDimensionalSliceUnwinder"
	"os"
)

func main() {
	args := os.Args[1:]
	unwound := twoDimensionalSliceUnwinder.Unwind(args)
	fmt.Printf("It was unwinded to : %v", unwound)
	fmt.Println()
}
