package main

import (
	"fmt"

	"github.com/raspiantoro/vulture"
)

func main() {
	b := []int{2, 5, 7, 8, 1, 3, 20}
	a := vulture.IntoSlice(b)

	// fmt.Println(a.IntoString())

	// c := a.Filter(func(i int) bool { return i < 7 })
	// fmt.Println(c)

	// d := a.ScanLeft(6, func(acc, v int) int { return int(math.Max(float64(acc), float64(v))) })
	// fmt.Println(d)

	// e := a.ScanRight(2, func(acc, v int) int { return int(math.Min(float64(acc), float64(v))) })
	// fmt.Println(e)

	// f := a.FoldLeft(3, func(acc, v int) int { return acc - v })
	// fmt.Println(f)

	g := a.FoldRight(20, func(acc, v int) int { return acc - v })
	fmt.Println(g)

	// str := "Hello world!!!"
	// strOps := vulture.FromString(str).Filter(func(b byte) bool { return b == []byte("l")[0] })
	// fmt.Println(strOps.IntoString())
}
