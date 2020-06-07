package main

import "github.com/fifoc/fifVideo/common"

var paletteGenerators = make([]common.FifPaletteGenerator, 0)

func registerPaletteGenerator(cb func() common.FifPalette, name string, short string) {
	generator := common.FifPaletteGenerator{
		Cb:    cb,
		Name:  name,
		Short: short,
	}

	paletteGenerators = append(paletteGenerators, generator)
}
