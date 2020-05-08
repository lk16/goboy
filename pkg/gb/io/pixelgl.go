package io

import (
	"image/color"
	"log"
	"time"

	"math"

	"github.com/Humpheh/goboy/pkg/gb"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// minRenderInterval limits how many frames we render per second
const minRenderInterval = time.Second / 60

// PixelScale is the multiplier on the pixels on display
var PixelScale float64 = 3

// PixelsIOBinding binds screen output and input using the pixels library.
type PixelsIOBinding struct {
	window    *pixelgl.Window
	picture   *pixel.PictureData
	lastFrame time.Time
}

// NewPixelsIOBinding returns a new Pixelsgl IOBinding
func NewPixelsIOBinding(enableVSync bool) *PixelsIOBinding {
	windowConfig := pixelgl.WindowConfig{
		Title: "GoBoy",
		Bounds: pixel.R(
			0, 0,
			float64(gb.ScreenWidth*PixelScale), float64(gb.ScreenHeight*PixelScale),
		),
		VSync:     enableVSync,
		Resizable: true,
	}

	window, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}

	// Hack so that pixelgl renders on Darwin
	window.SetPos(window.GetPos().Add(pixel.V(0, 1)))

	picture := &pixel.PictureData{
		Pix:    make([]color.RGBA, gb.ScreenWidth*gb.ScreenHeight),
		Stride: gb.ScreenWidth,
		Rect:   pixel.R(0, 0, gb.ScreenWidth, gb.ScreenHeight),
	}

	monitor := PixelsIOBinding{
		window:  window,
		picture: picture,
	}

	monitor.updateCamera()

	return &monitor
}

// updateCamera updates the window camera to center the output.
func (mon *PixelsIOBinding) updateCamera() {
	xScale := mon.window.Bounds().W() / 160
	yScale := mon.window.Bounds().H() / 144
	scale := math.Min(yScale, xScale)

	shift := mon.window.Bounds().Size().Scaled(0.5)
	cam := pixel.IM.Scaled(pixel.ZV, scale).ScaledXY(pixel.ZV, pixel.Vec{X: 1, Y: -1}).Moved(shift)
	mon.window.SetMatrix(cam)
}

// IsRunning returns if the game should still be running. When
// the window is closed this will be false so the game stops.
func (mon *PixelsIOBinding) IsRunning() bool {
	return !mon.window.Closed()
}

// Render renders the pixels on the screen.
func (mon *PixelsIOBinding) Render(frame *gb.Frame) {

	if time.Since(mon.lastFrame) < minRenderInterval {
		return
	}

	mon.lastFrame = time.Now()

	go func() {

		mon.picture.Pix = frame[:]

		rgba := gb.GetPaletteColour(3)
		mon.window.Clear(rgba)

		spr := pixel.NewSprite(mon.picture, pixel.R(0, 0, gb.ScreenWidth, gb.ScreenHeight))
		spr.Draw(mon.window, pixel.IM)

		mon.updateCamera()
		mon.window.Update()
	}()
}

// SetTitle sets the title of the game window.
func (mon *PixelsIOBinding) SetTitle(title string) {
	mon.window.SetTitle(title)
}

// Toggle the fullscreen window on the main monitor.
func (mon *PixelsIOBinding) toggleFullscreen() {
	if mon.window.Monitor() == nil {
		monitor := pixelgl.PrimaryMonitor()
		_, height := monitor.Size()
		mon.window.SetMonitor(monitor)
		PixelScale = height / 144
	} else {
		mon.window.SetMonitor(nil)
		PixelScale = 3
	}
}

var keyMap = map[pixelgl.Button]gb.Button{
	pixelgl.KeyZ:         gb.ButtonA,
	pixelgl.KeyX:         gb.ButtonB,
	pixelgl.KeyBackspace: gb.ButtonSelect,
	pixelgl.KeyEnter:     gb.ButtonStart,
	pixelgl.KeyRight:     gb.ButtonRight,
	pixelgl.KeyLeft:      gb.ButtonLeft,
	pixelgl.KeyUp:        gb.ButtonUp,
	pixelgl.KeyDown:      gb.ButtonDown,

	pixelgl.KeyEscape: gb.ButtonPause,
	pixelgl.KeyEqual:  gb.ButtonChangePallete,
	pixelgl.KeyQ:      gb.ButtonToggleBackground,
	pixelgl.KeyW:      gb.ButtonToggleSprites,
	pixelgl.KeyE:      gb.ButttonToggleOutputOpCode,
	pixelgl.KeyD:      gb.ButtonPrintBGMap,
	pixelgl.Key7:      gb.ButtonToggleSoundChannel1,
	pixelgl.Key8:      gb.ButtonToggleSoundChannel2,
	pixelgl.Key9:      gb.ButtonToggleSoundChannel3,
	pixelgl.Key0:      gb.ButtonToggleSoundChannel4,
}

// ProcessInput checks the input and process it.
func (mon *PixelsIOBinding) ButtonInput() gb.ButtonInput {

	if mon.window.JustPressed(pixelgl.KeyF) {
		mon.toggleFullscreen()
	}

	var buttonInput gb.ButtonInput

	for handledKey, button := range keyMap {
		if mon.window.JustPressed(handledKey) {
			buttonInput.Pressed = append(buttonInput.Pressed, button)
		}
		if mon.window.JustReleased(handledKey) {
			buttonInput.Released = append(buttonInput.Released, button)
		}
	}

	return buttonInput
}
