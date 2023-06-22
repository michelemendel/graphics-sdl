package main

// In golang use sdl2 and texture to drax a blue 10x10 pixel square on a yellow background. The square has an alpha of 55.

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	width  = 640
	height = 480
)

func run4() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		log.Fatalf("Error initializing SDL: %s\n", err)
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		log.Fatalf("Error initializing TTF: %s\n", err)
	}
	defer ttf.Quit()

	window, err := sdl.CreateWindow("SDL2 GOLANG", 0, 0, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Error creating window: %s\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Error creating renderer: %s\n", err)
	}
	defer renderer.Destroy()

	// Create a yellow background
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.Clear()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 10, 10)
	if err != nil {
		log.Fatalf("Error creating texture: %s\n", err)
	}

	renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(0, 0, 255, 55)
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: 10, H: 10})
	renderer.SetRenderTarget(nil)
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(texture, nil, &sdl.Rect{X: 100, Y: 100, W: 10, H: 10})
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
