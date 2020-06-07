package main

import (
	"flag"
	"fmt"
	"github.com/fifoc/fifVideo/palettes"
	"os"
)

func init() {
	registerFlags()

	registerPaletteGenerator(palettes.FifOcPalette, "Open Computers T3 GPU", "ocgpu")
}

func main() {
	flag.Parse()

	if cliMode == "debug" {
		debugMode()
	} else if cliMode == "normal" {

	} else {
		fmt.Println("Unknown mode: " + cliMode)
		os.Exit(1)
	}
}
