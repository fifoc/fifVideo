package main

import "flag"

var cliMode string
var cliMedia string
var cliInput string
var cliOutput string
var cliPalette string

func registerFlags() {
	flag.StringVar(&cliMode, "mode", "normal", "Set application mode")
	flag.StringVar(&cliMedia, "media", "image", "What media to parse? (image or video)")
	flag.StringVar(&cliInput, "input", "input.png", "Input file")
	flag.StringVar(&cliOutput, "output", "output.fif", "Output file")
	flag.StringVar(&cliPalette, "palette", "ocgpu", "Palette mode to use")
}
