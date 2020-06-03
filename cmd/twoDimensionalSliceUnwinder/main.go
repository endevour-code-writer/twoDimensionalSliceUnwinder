package main

import (
	"fmt"
	"github.com/endevour-code-writer/twoDimensionalSliceUnwinder/internal/twoDimensionalSliceUnwinder"
	"os"
)

func main() {
	args := os.Args[1:]
	unwinded := twoDimensionalSliceUnwinder.Unwind(args)
	fmt.Printf("It was unwinded to : %v", unwinded)
	fmt.Println()
}
