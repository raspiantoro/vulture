package main

import (
	"fmt"
	"math"

	"github.com/raspiantoro/vulture"
)

func main() {
	b := []int{2, 1, 7, 8, 1, 3, 20}
	a := vulture.FromSlice(b)

	fmt.Println("a: ", a)

	// c := a.Iter().Filter(func(i int) bool { return i < 7 }).Collect()
	// fmt.Println("c: ", c)

	// d := a.Iter().ScanLeft(6, func(acc, v int) int { return int(math.Max(float64(acc), float64(v))) }).Collect()
	// fmt.Println("d: ", d)

	e := a.Iter().OldScanRight(2, func(acc, v int) int { return int(math.Min(float64(acc), float64(v))) }).Collect()
	fmt.Println("e: ", e)

	f := a.Iter().ScanRight(2, func(acc, v int) int { return int(math.Min(float64(acc), float64(v))) }).Collect()
	fmt.Println("f: ", f)

	// g := a.Iter().FoldLeft(20, func(acc, v int) int { return acc - v })
	// fmt.Println("g: ", g)

	// h := a.Iter().FoldRight(20, func(acc, v int) int { return acc - v })
	// fmt.Println("h: ", h)

	fmt.Println("a: ", a)

	// str := "Hello world!!!"
	// strOps := vulture.FromString(str).Iter().Filter(func(b byte) bool { return b == []byte("l")[0] }).Collect()
	// fmt.Println(strOps.IntoString())
}
