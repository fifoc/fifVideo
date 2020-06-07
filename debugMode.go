package main

import (
	"fmt"
	"os"
)

func debugMode() {
	fmt.Println("Available palette modes:")
	for i := 0; i < len(paletteGenerators); i++ {
		fmt.Println("-", paletteGenerators[i].Name, "(", paletteGenerators[i].Short, ")")
	}
	os.Exit(0)
}
