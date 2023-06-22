package main

func main() {
	// run0()
	// run1()
	// run3()
	// run4()
	// run5()
	// run6()
	// runSin()
	runFract()
	// lerping()

	// var z complex128 = 1 + 5i
	// fmt.Println(cmplx.Abs(z) * cmplx.Abs(z))
}

// func lerping() {
// 	a, b := -2.0, 2.0

// 	t := 0.5
// 	fmt.Println(t, lerp(a, b, t))

// 	fns := []func(a, b, t float64) float64{lerp, lerpClamp}
// 	for _, fn := range fns {
// 		testLerp(a, b, fn)
// 	}
// }

// func testLerp(a, b float64, lerpFn func(a, b, t float64) float64) {
// 	for t := 0.0; t <= 1.0; t += 0.1 {
// 		fmt.Printf("%v:\t%v\n", t, lerpFn(a, b, t))
// 	}
// 	fmt.Println("--------")
// }

// func lerp(a, b, t float64) float64 {
// 	return a + t*(b-a)
// }

// func lerpClamp(a, b, t float64) float64 {
// 	if t < 0 {
// 		t = 0
// 	}
// 	if t > 1 {
// 		t = 1
// 	}
// 	return a + t*(b-a)
// }
