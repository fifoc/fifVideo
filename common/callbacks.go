package common

type FifPaletteGenerator struct {
	Cb    func() FifPalette
	Name  string
	Short string
}
