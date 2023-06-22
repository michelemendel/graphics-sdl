package main

import (
	"fmt"
	"math/cmplx"

	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	r, g, b, a byte
}

// type Pos struct {
// x, y float32
// }

type Pixel struct {
	Point sdl.Point
	Color
}

var err error
var window *sdl.Window
var renderer *sdl.Renderer
var texture *sdl.Texture
var winTitle string = "Wrong fractal"
var winPosX, winPosY, winWidth, winHeight int32 = 0, 0, 1000, 1000

var z complex128
var c complex128

// var speed float64 = 2
var zoom int32 = 1
var src, dst *sdl.Rect

// var fromX, toX, fromY, toY float64 = -2.0, 0.47, -1.12, 1.12
var fromX, toX float64 = -2.0, 0.6
var fromY, toY float64 = -1.22, 1.22
var steps float64 = 0.001
var iterations int32 = 100
var zfBase float64 = (fromX - toX) / (fromY - toY)
var zf float64

var runAgain bool = true

// var points []sdl.Point

// var pixelWidth, pixelHeight int32 = 1, 1

// --------------------------------------------------------------------------------
// Init

func initSDL(winTitle string, winPosX, winPosY, winWidth, winHeight int32) (*sdl.Renderer, *sdl.Texture) {
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, renderer, err = sdl.CreateWindowAndRenderer(winWidth, winHeight, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	window.SetPosition(winPosX, winPosY)
	window.SetTitle(winTitle)

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_TARGET, winWidth, winHeight)
	if err != nil {
		panic(err)
	}

	return renderer, texture
}

func cleanUp() {
	sdl.Quit()
	window.Destroy()
	texture.Destroy()
}

// --------------------------------------------------------------------------------
// Run function
func runFract() {
	renderer, texture = initSDL(winTitle, winPosX, winPosY, winWidth, winHeight)
	defer cleanUp()

	var frameWidth int32 = 0
	src = &sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight}
	dst = &sdl.Rect{X: frameWidth, Y: frameWidth, W: winWidth - 2*frameWidth, H: winHeight - 2*frameWidth}

	// Clear screen
	renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	running := true
	for running {
		eventHandling(&running, &zoom)
		renderer.SetRenderTarget(texture)

		var iters, i int32
		var cx, cy, x, y float64
		var px, py float32
		var r, g, b, a byte

		// scaleX := float64(winWidth) / math.Abs(fromX-toX)
		// scaleY := float64(winHeight) / math.Abs(fromY-toY)

		if runAgain {
			zf = (fromX - toX) / (fromY - toY)
			zfRatio := zf / zfBase
			// steps /= math.Sqrt(zfRatio)
			iterations = int32(float64(iterations) * zfRatio)
			fmt.Printf("X:(%.4f,%.4f), Y:(%.4f,%.4f), Z:(%.4f,%.4f,%.4f), S:%.4f, I:%d\n", fromX, toX, fromY, toY, zfBase, zf, zfRatio, steps, iterations)
			for x = 0.0; x < 1.0; x += steps {
				for y = 0.0; y < 1.0; y += steps {
					cx = lerp(fromX, toX, x)
					cy = lerp(fromY, toY, y)
					c = complex(cx, cy)
					z = complex(0, 0)
					iters = 0
					for i = 0; i < iterations; i++ {
						z = z*z + c
						if cmplx.Abs(z) > 2 {
							iters = i
							break
						}
					}

					px = float32(x) * float32(winWidth)  //float32(x*scaleX) + float32(winWidth/2)
					py = float32(y) * float32(winHeight) //float32(y*scaleY) + float32(winHeight/2)
					// fmt.Printf("(%v,%v) -> (%v,%v)\n", cx, cy, x, y)
					if iters == 0 {
						renderer.SetDrawColor(0, 0, 0, 255)
						renderer.DrawPointF(px, py)
					} else {
						r = byte(iters * 12 % 255)
						g = byte(iters * 13 % 255)
						b = byte(iters * 14 % 255)
						a = 255
						renderer.SetDrawColor(r, g, b, a)
						renderer.DrawPointF(px, py)
					}
				}
			}
			runAgain = false
			fmt.Println("Done!")
		}

		renderer.SetRenderTarget(nil)
		// Frame, if shown
		// renderer.SetDrawColor(255, 100, 0, 255)
		// renderer.Clear()

		renderer.Copy(texture, src, dst)
		renderer.Present()

		sdl.Delay(16)
	}
}

func lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

// --------------------------------------------------------------------------------
// Events

func eventHandling(running *bool, zoom *int32) {
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
					if t.State == sdl.RELEASED {
						keys += string(keyCode)
						switch keys {
						case "a":
							fromX -= 0.1
							toX -= 0.1
							runAgain = true
						case "d":
							fromX += 0.1
							toX += 0.1
							runAgain = true
						case "w":
							fromY -= 0.1
							toY -= 0.1
							runAgain = true
						case "s":
							fromY += 0.1
							toY += 0.1
							runAgain = true
							// Zoom in
						case "z":
							fromX += 0.1
							toX -= 0.1
							fromY += 0.1
							toY -= 0.1
							runAgain = true
							// Zoom out
						case "x":
							fromX -= 0.1
							toX += 0.1
							fromY -= 0.1
							toY += 0.1
							runAgain = true
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
