package main

// import (
// 	"fmt"
// 	"math"
// 	"math/cmplx"
// 	"unsafe"

// 	"github.com/veandco/go-sdl2/sdl"
// )

// type Color struct {
// 	r, g, b, a byte
// }

// type Pos struct {
// 	x, y float32
// }

// func clear(pixels []byte) {
// 	for i := range pixels {
// 		pixels[i] = 0
// 	}
// }

// const (
// 	winPosX, winPosY, winWidth, winHeight int32 = 0, 0, 1200, 1200
// 	pixelWidth, pixelHeight               int32 = 1, 1
// )

// func run0() {
// 	err := sdl.Init(sdl.INIT_EVERYTHING)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer sdl.Quit()

// 	window, err := sdl.CreateWindow("z", winPosX, winPosY, winWidth, winHeight, sdl.WINDOW_SHOWN)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer window.Destroy()

// 	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer renderer.Destroy()

// 	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, winWidth, winHeight)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer tex.Destroy()

// 	pixels := make([]byte, winWidth*winHeight*4)
// 	pixel := Pixel{Pos{0, 0}, Color{255, 255, 255, 255}}

// 	for {
// 		pixel.updatePixel()
// 		pixel.makePixel(pixels)
// 		tex.Update(nil, unsafe.Pointer(&pixels[0]), int(winWidth*4))
// 		// clear(pixels)
// 		renderer.Copy(tex, nil, nil)
// 		// renderer.Copy(tex, nil, &sdl.Rect{X: int32(pixel.x), Y: int32(pixel.y), W: winWidth, H: winHeight})
// 		renderer.Present()

// 		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 			switch event.(type) {
// 			case *sdl.QuitEvent:
// 				println("Quit")
// 				return
// 			}
// 		}
// 		sdl.Delay(16)
// 	}
// }

// // --------------------------------------------------------------------------------
// func setPixel(p Pos, c Color, pixels []byte) {
// 	index := (int32(p.y)*winWidth + int32(p.x)) * 4
// 	if index > 0 && index < int32(len(pixels)-4) {
// 		pixels[index] = c.r
// 		pixels[index+1] = c.g
// 		pixels[index+2] = c.b
// 		pixels[index+3] = c.a
// 	}
// }

// // --------------------------------------------------------------------------------
// type Pixel struct {
// 	Pos
// 	Color
// }

// func (p *Pixel) makePixel(pixels []byte) {
// 	var x, y int32
// 	for y = 0; y < pixelHeight; y++ {
// 		for x = 0; x < pixelWidth; x++ {
// 			setPixel(Pos{
// 				float32(x) + p.x,
// 				float32(y) + p.y,
// 			}, Color{p.r, p.g, p.b, p.a}, pixels)
// 		}
// 	}
// }

// // --------------------------------------------------------------------------------

// const scale = 308

// const trX, trY float32 = float32(winWidth / 2), float32(winHeight / 2)

// // const trX, trY float32 = 0, 0

// var z complex128 = complex(0, 0)
// var c complex128 = complex(0, 0)
// var cStepR, cStepI float64 = 0.013, 0.001111

// func (p *Pixel) updatePixel() {
// 	p.x = float32(real(z))*scale + trX
// 	p.y = float32(imag(z))*scale + trY

// 	p.r = 100 + byte(math.Abs(real(c))*255)
// 	p.g = 100 + byte(math.Abs(imag(c))*255)
// 	p.b = 100 + byte(cmplx.Abs(c)*255)

// 	prevZ := z
// 	z = z*z + c

// 	if real(z) >= 2 || imag(z) >= 2 ||
// 		close(prevZ, z, 0.001) ||
// 		p.x <= 0 || p.x >= float32(winWidth) ||
// 		p.y <= 0 || p.y >= float32(winHeight) {
// 		// fmt.Println("ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ")
// 		// fmt.Println(z, c, cmplx.Abs(z), cmplx.Abs(c), "(", p.x, p.y, "), (", p.r, p.g, p.b, ")")
// 		z = complex(0, 0)
// 		cStepR += 0.0001
// 		cStepI += 0.0001
// 		c += complex(cStepR, cStepI)
// 	}

// 	fmt.Println(z, c, cmplx.Abs(z), cmplx.Abs(c), "(", p.x, p.y, "), (", p.r, p.g, p.b, ")")
// }

// func close(z1, z2 complex128, e float64) bool {
// 	return math.Abs(real(z1)-real(z2)) < e && math.Abs(imag(z1)-imag(z2)) < e
// }

// func round(z complex128, decs int) complex128 {
// 	return complex(
// 		math.Round(float64(real(z))*math.Pow10(decs))/math.Pow10(decs),
// 		math.Round(float64(imag(z))*math.Pow10(decs))/math.Pow10(decs))
// }
