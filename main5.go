package main

// In golang use sdl2 and texture to drax a blue 10x10 pixel square on a yellow background. The square has an alpha of 55.

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

func run5() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println(err)
		return
	}

	window, err := sdl.CreateWindow("SDL2 Pixel Drawing", 0, 0, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the background color
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.Clear()

	// Create a texture
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, 10, 10)
	if err != nil {
		panic(err)
	}
	tex.Lock(&sdl.Rect{})
	pixels := make([]byte, 10*10*4)
	for i := 0; i < 10*10; i++ {
		pixels[i*4] = 0     // Blue
		pixels[i*4+1] = 0   // Green
		pixels[i*4+2] = 255 // Red
		pixels[i*4+3] = 55  // Alpha
	}
	tex.Update(nil, unsafe.Pointer(&pixels[0]), 10*4)
	tex.Unlock()
	renderer.Copy(tex, nil, &sdl.Rect{X: 100, Y: 100, W: 10, H: 10})
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
