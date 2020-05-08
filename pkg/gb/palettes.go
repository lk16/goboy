package gb

import (
	"image/color"

	"github.com/Humpheh/goboy/pkg/bits"
)

const (
	// PaletteGreyscale is the default greyscale gameboy colour palette.
	PaletteGreyscale = byte(iota)
	// PaletteOriginal is more authentic looking green tinted gameboy
	// colour palette  as it would have been on the GameBoy
	PaletteOriginal
	// PaletteBGB used by default in the BGB emulator.
	PaletteBGB
)

// CurrentPalette is the global current DMG palette.
var CurrentPalette = PaletteBGB

// Palettes is an mapping from colour palettes to their colour values
// to be used by the emulator.
var Palettes = [][]color.RGBA{
	// PaletteGreyscale
	{
		color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF},
		color.RGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xFF},
		color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
	},
	// PaletteOriginal
	{
		color.RGBA{R: 0x9B, G: 0xBC, B: 0x0F, A: 0xFF},
		color.RGBA{R: 0x8B, G: 0xAC, B: 0x0F, A: 0xFF},
		color.RGBA{R: 0x30, G: 0x62, B: 0x30, A: 0xFF},
		color.RGBA{R: 0x0F, G: 0x38, B: 0x0F, A: 0xFF},
	},
	// PaletteBGB
	{
		color.RGBA{R: 0xE0, G: 0xF8, B: 0xD0, A: 0xFF},
		color.RGBA{R: 0x88, G: 0xC0, B: 0x70, A: 0xFF},
		color.RGBA{R: 0x34, G: 0x68, B: 0x56, A: 0xFF},
		color.RGBA{R: 0x08, G: 0x18, B: 0x20, A: 0xFF},
	},
}

// GetPaletteColour returns the colour based on the colour index and the currently
// selected palette.
func GetPaletteColour(index uint) color.RGBA {
	return Palettes[CurrentPalette][index]
}

// NewPalette makes a new CGB colour palette.
func NewPalette() *cgbPalette {
	pal := make([]byte, 0x40)
	for i := range pal {
		pal[i] = 0xFF
	}

	return &cgbPalette{Palette: pal}
}

func changePallete() {
	CurrentPalette = (CurrentPalette + 1) % byte(len(Palettes))
}

// Palette for cgb containing information tracking the palette colour info.
type cgbPalette struct {
	// Palette colour information.
	Palette []byte
	// Current index the palette is referencing.
	Index byte
	// If to auto increment on write.
	Inc bool
}

// Update the index the palette is indexing and set
// auto increment if bit 7 is set.
func (pal *cgbPalette) updateIndex(value byte) {
	pal.Index = value & 0x3F
	pal.Inc = bits.Test(value, 7)
}

// Read the palette information stored at the current index.
func (pal *cgbPalette) read() byte {
	return pal.Palette[pal.Index]
}

// Read the current index.
func (pal *cgbPalette) readIndex() byte {
	return pal.Index
}

// Write a value to the palette at the current index.
func (pal *cgbPalette) write(value byte) {
	pal.Palette[pal.Index] = value
	if pal.Inc {
		pal.Index = (pal.Index + 1) & 0x3F
	}
}

// Get the rgba colour for a palette at a colour number.
func (pal *cgbPalette) get(palette, num uint) color.RGBA {
	idx := (palette << 3) | (num << 1)
	colour := uint(pal.Palette[idx]) | uint(pal.Palette[idx+1])<<8

	r := colour & 0x1F
	g := (colour >> 5) & 0x1F
	b := (colour >> 10) & 0x1F

	return color.RGBA{R: colArr[r], G: colArr[g], B: colArr[b], A: 0xFF}
}

// Mapping of the 5 bit colour value to a 8 bit value.
var colArr = []uint8{
	0x0,
	0x8,
	0x10,
	0x18,
	0x20,
	0x29,
	0x31,
	0x39,
	0x41,
	0x4a,
	0x52,
	0x5a,
	0x62,
	0x6a,
	0x73,
	0x7b,
	0x83,
	0x8b,
	0x94,
	0x9c,
	0xa4,
	0xac,
	0xb4,
	0xbd,
	0xc5,
	0xcd,
	0xd5,
	0xde,
	0xe6,
	0xee,
	0xf6,
	0xff,
}
