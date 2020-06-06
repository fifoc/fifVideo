package common

type FifColor uint32

func (a FifColor) split() (int, int, int) {
	r := int(a >> 16)
	r = r & 0xFF
	g := int(a >> 8)
	g = g & 0xFF
	b := int(a)
	b = b & 0xFF

	return r, g, b
}
