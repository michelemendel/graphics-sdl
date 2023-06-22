package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func run1() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("My Window", 0, 0, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	renderer.SetDrawColor(0, 0, 0, 255)
	var x, y int32
	for y = 0; y < 600; y++ {
		for x = 0; x < 800; x++ {
			renderer.DrawPoint(x, y)
		}
	}

	renderer.SetDrawColor(0, 0, 255, 5)
	for y = 100; y < 106; y++ {
		for x = 100; x < 106; x++ {
			renderer.DrawPoint(x, y)
		}
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
