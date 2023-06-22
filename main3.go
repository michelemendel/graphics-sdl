package main

// In golang use sdl2 and texture to drax a blue pixel on a yellow background.

import (
	"github.com/veandco/go-sdl2/sdl"
)

func run3() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", 0, 0, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Draw the yellow background color
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	// Draw the square
	// texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 100, 100)

	// renderer.SetRenderTarget(texture)
	// texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	// texture.SetAlphaMod(255)
	// renderer.SetDrawColor(255, 0, 0, 55)
	// texture.SetColorMod(255, 255, 255)
	// renderer.Copy(texture, nil, &sdl.Rect{X: 400, Y: 300, W: 100, H: 100})

	// renderer.SetDrawBlendMode(sdl.BLENDMODE_MOD)

	var x, y int32
	for y = 10; y < 110; y++ {
		for x = 10; x < 110; x++ {
			renderer.SetDrawColor(255, 0, 255, 255)
			renderer.DrawPoint(x, y)
		}
	}

	// renderer.FillRect(&sdl.Rect{X: 20, Y: 20, W: 100, H: 100})

	// renderer.SetDrawColor(0, 255, 0, 10)
	// renderer.FillRect(&sdl.Rect{X: 100, Y: 100, W: 100, H: 100})

	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				return
			}
		}
		sdl.Delay(16)
	}
}
