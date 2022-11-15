package main

import (
	"fmt"

	"github.com/raspiantoro/vulture"
)

func main() {
	b := []int{2, 5, 7, 8, 1, 3, 20}
	a := vulture.IntoSlice(b)

	fmt.Println(a.IntoString())

	c := a.Filter(func(i int) bool { return i < 7 })
	fmt.Println(c)

	d := a.ScanLeft(vulture.ScanTypeMin)
	fmt.Println(d)

	e := a.ScanLeft(vulture.ScanTypeMax)
	fmt.Println(e)

	str := "Hello world!!!"
	strOps := vulture.FromString(str).Filter(func(b byte) bool { return b == []byte("l")[0] })
	fmt.Println(strOps.IntoString())
}
