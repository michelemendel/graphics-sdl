// Prompt: In golang use sdl2 to animate a yellow sinus curve on a blue background.
// Bing: Here's an example of how you can animate a yellow sinus curve on a blue background using SDL2 in Go:

package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// var err error
// var window *sdl.Window
// var renderer *sdl.Renderer
// var texture *sdl.Texture
var windowTitle string = "Sinus Curve"
var windowWidth, windowHeight int32 = 900, 400
var speed float64 = 2

// var zoom int32 = 1

// var src, dst *sdl.Rect

// --------------------------------------------------------------------------------
func runSin() int {
	window, renderer, err = sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.RENDERER_ACCELERATED)
	window.SetPosition(0, 0)
	window.SetTitle(windowTitle)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, windowWidth, windowHeight)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	var phase float64 = 0
	var points0 []sdl.Point
	var x, y0 float64
	// var y1 float64
	// var points1 []sdl.Point
	// var sumY []sdl.Point
	// var pixels = make([]uint32, windowWidth*windowHeight)

	src = &sdl.Rect{X: 0, Y: 0, W: windowWidth, H: windowHeight}
	var frameWidth int32 = 0
	dst = &sdl.Rect{X: frameWidth, Y: frameWidth, W: windowWidth - 2*frameWidth, H: windowHeight - 2*frameWidth}

	running := true
	for running {
		events(&running, &zoom)
		renderer.SetRenderTarget(texture)

		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		y0 = fn(x/2, phase)
		points0 = append(points0, sdl.Point{X: int32(x), Y: int32(y0) + (windowHeight / (2 * zoom))})
		if len(points0) > 100 {
			points0 = points0[1:]
		}
		// fmt.Println(points0)
		renderer.SetDrawColor(255, 255, 0, 255)
		// renderer.DrawPoints(points0)
		renderer.DrawLines(points0)

		// y1 = fn(x, phase)
		// points1 = append(points1, sdl.Point{X: int32(x), Y: int32(y1) + (windowHeight / 2)})
		// renderer.SetDrawColor(0, 255, 0, 15)
		// renderer.DrawPoints(points1)

		// sumY = append(sumY, sdl.Point{X: int32(x), Y: int32(y0+y1) + (windowHeight / 2)})
		// renderer.SetDrawColor(255, 255, 255, 55)
		// renderer.DrawPoints(sumY)

		x += speed

		renderer.SetRenderTarget(nil)
		renderer.SetDrawColor(255, 1, 255, 255)
		renderer.Clear()
		renderer.Copy(texture, src, dst)
		renderer.Present()

		if int32(x) > windowWidth {
			phase += speed * 1.01
			x = float64(0)
			points0 = nil
			// points1 = nil
			// fmt.Println(phase)
			// clearScreen(renderer, texture)
		}

		sdl.Delay(16)
	}

	return 0
}

func fn(x, phase float64) float64 {
	return math.Sin(x/16+phase) * float64(100/zoom)
}

// func clearScreen(renderer *sdl.Renderer, texture *sdl.Texture) {
func clearScreen(renderer *sdl.Renderer) {
	// renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	// renderer.SetRenderTarget(nil)
}

func events(running *bool, zoom *int32) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			*running = false
		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			keys := ""

			// Modifier keys
			switch t.Keysym.Mod {
			case sdl.KMOD_LALT:
				keys += "Left Alt"
			case sdl.KMOD_LCTRL:
				keys += "Left Control"
			case sdl.KMOD_LSHIFT:
				keys += "Left Shift"
			case sdl.KMOD_LGUI:
				keys += "Left Meta or Windows key"
			case sdl.KMOD_RALT:
				keys += "Right Alt"
			case sdl.KMOD_RCTRL:
				keys += "Right Control"
			case sdl.KMOD_RSHIFT:
				keys += "Right Shift"
			case sdl.KMOD_RGUI:
				keys += "Right Meta or Windows key"
			case sdl.KMOD_NUM:
				keys += "Num Lock"
			case sdl.KMOD_CAPS:
				keys += "Caps Lock"
			case sdl.KMOD_MODE:
				keys += "AltGr Key"
			}

			if keyCode < 10000 {
				if keys != "" {
					keys += " + "
				}

				// If the key is held down, this will fire
				if t.Repeat > 0 {
					// keys += string(keyCode) + " repeating"
				} else {
					if t.State == sdl.PRESSED {
						keys += string(keyCode)
						switch keys {
						case ",":
							if *zoom > 1 {
								*zoom--
							}
							fmt.Println("zoom: ", *zoom)
							src.H = windowHeight / *zoom
						case ".":
							*zoom++
							fmt.Println("zoom: ", *zoom)
							src.H = windowHeight / *zoom
						case "'":
							speed += 0.5
							fmt.Println("speed: ", speed)
						case ";":
							speed -= 0.5
							fmt.Println("speed: ", speed)
						}
					} else if t.State == sdl.RELEASED {
						keys += string(keyCode)
					}
				}

			}

			// if keys != "" {
			// fmt.Println(keys)
			// }
		}
	}
}
