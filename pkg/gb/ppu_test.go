package gb

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSpritePriority runs the mooneye sprite_priority.gb test rom and asserts that the
// output frame matches the expected image.
func TestSpritePriority(t *testing.T) {
	// Takes about 10 frames to render the sprite priority image
	const maxPPUIterations = 10

	// Override the palette with the colours in the expected image
	Palettes[CurrentPalette] = []color.RGBA{
		color.RGBA{R: 3, G: 3, B: 3, A: 0xFF},
		color.RGBA{R: 2, G: 2, B: 3, A: 0xFF},
		color.RGBA{R: 1, G: 1, B: 1, A: 0xFF}, // not used in expected image
		color.RGBA{R: 0, G: 0, B: 0, A: 0xFF},
	}

	// Map of colours in the image to color in the palette
	var imageMap = map[color.Color]byte{
		color.Gray{Y: 255}: 3,
		color.Gray{Y: 111}: 2,
		color.Gray{Y: 0}:   0,
	}

	// Load the test ROM and iterate a few frames to load the image
	gb, err := NewGameboy("./../../roms/mooneye/runnable/sprite_priority.gb")
	require.NoError(t, err, "error in init gb %v", err)
	for i := 0; i < maxPPUIterations; i++ {
		gb.Update()
	}

	// Load the expected output image
	img, err := loadImage("../../roms/mooneye/runnable/sprite_priority-expected.png")
	if err != nil {
		t.Fatalf("Could not open expected image: %v", err)
	}

	// Iterate over the image and assert each pixel matches the expected image
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			actual := gb.PreparedData[y*ScreenWidth+x]
			expected, ok := imageMap[img.At(x, y)]
			require.True(t, ok, "unexpected colour in expected image: %v", img.At(x, y))
			require.Equal(t, expected, actual.R, "incorrect pixel at X:%v Y:%x", x, y)
		}
	}
}

// Load a PNG image
func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening image: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %v", err)
	}
	return img, nil
}
