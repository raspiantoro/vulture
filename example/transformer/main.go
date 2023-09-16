package main

import (
	"fmt"

	"github.com/raspiantoro/vulture"
)

type rectangle struct {
	high  int
	width int
}

type area struct {
	result int
}

func main() {
	r := []rectangle{
		{
			high:  10,
			width: 6,
		},
		{
			high:  5,
			width: 7,
		},
		{
			high:  3,
			width: 2,
		},
	}

	list := vulture.FromSlice(r)
	lTrasformer := vulture.Transformer[area](list)
	a := lTrasformer.
		Iter().
		Add(rectangle{2, 2}, func(addition, val rectangle) rectangle {
			return rectangle{
				high:  val.high + addition.high,
				width: val.width + addition.width,
			}
		}).
		Transform().
		Map(func(val rectangle) area {
			return area{
				result: val.high * val.width,
			}
		}).
		Iter().
		FoldLeft(area{0}, func(accumulator, val area) area {
			return area{
				result: accumulator.result + val.result,
			}
		})

	fmt.Println(a)
}
