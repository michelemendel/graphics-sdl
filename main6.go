package main

// In golang use sdl2 and texture to draw a yellow background. The place a blue 10x10 pixel square on position x=12, y=13 . The square has an alpha of 55.

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth6, winHeight6 int32 = 800, 600

func run6() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 example", 0, 0, winWidth6, winHeight6, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("creating window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	// create yellow background texture
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, int32(winWidth6), int32(winHeight6))
	if err != nil {
		fmt.Println("creating texture:", err)
		return
	}
	defer texture.Destroy()
	renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.Clear()
	renderer.SetRenderTarget(nil)

	// create blue square texture
	squareTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 10, 10)
	if err != nil {
		fmt.Println("creating texture:", err)
		return
	}
	defer squareTex.Destroy()
	renderer.SetRenderTarget(squareTex)
	renderer.SetDrawColor(0, 0, 255, 55)
	renderer.Clear()
	renderer.SetRenderTarget(nil)

	// combine background and square texture
	err = renderer.Copy(texture, nil, nil)
	if err != nil {
		fmt.Println("copying texture:", err)
		return
	}
	err = renderer.Copy(squareTex, &sdl.Rect{X: 12, Y: 13, W: 10, H: 10}, nil)
	if err != nil {
		fmt.Println("copying texture:", err)
		return
	}

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
