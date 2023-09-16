package main

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/raspiantoro/vulture"
)

func main() {
	b := []int{2, 1, 7, 8, 1, 3, 20}
	a := vulture.FromSlice(b)

	fmt.Println("a: ", a)

	// c := a.Iter().Filter(func(i int) bool { return i < 7 }).Collect()
	// fmt.Println("c: ", c)

	e := a.Iter().ScanLeft(6, func(acc, v int) int { return int(math.Max(float64(acc), float64(v))) }).Collect()
	fmt.Println("e: ", e)

	fmt.Println("a: ", a)

	f := a.IterRef().ScanLeft(6, func(acc, v int) int { return int(math.Max(float64(acc), float64(v))) }).Collect()
	fmt.Println("f: ", *f)
	fmt.Println("f address: ", unsafe.Pointer(f))

	fmt.Println("a: ", a)
	fmt.Println("a address: ", unsafe.Pointer(&a))

	h := a.Iter().ScanRight(2, func(acc, v int) int { return int(math.Min(float64(acc), float64(v))) }).Collect()
	fmt.Println("h: ", h)

	fmt.Println("a: ", a)

	i := a.IterRef().ScanRight(2, func(acc, v int) int { return int(math.Min(float64(acc), float64(v))) }).Collect()
	fmt.Println("i: ", *i)
	fmt.Println("i address: ", unsafe.Pointer(i))

	// j := a.Iter().FoldLeft(20, func(acc, v int) int { return acc - v })
	// fmt.Println("j: ", j)

	// k := a.Iter().FoldRight(20, func(acc, v int) int { return acc - v })
	// fmt.Println("k: ", k)

	fmt.Println("a: ", a)
	fmt.Println("a address: ", unsafe.Pointer(&a))

	l := vulture.Ranges(int8(-8), 8).Collect()
	fmt.Println("l: ", l)

	m := vulture.Ranges[int8](8, -8).Collect()
	fmt.Println("m: ", m)

	n := vulture.Ranges[int8](-8, 8).Exclusive().Collect()
	fmt.Println("n: ", n)

	o := vulture.Ranges[int8](8, -8).Exclusive().Collect()
	fmt.Println("o: ", o)

	// str := "Hello world!!!"
	// strOps := vulture.FromString(str).Iter().Filter(func(b byte) bool { return b == []byte("l")[0] }).Collect()
	// fmt.Println(strOps.IntoString())
}
