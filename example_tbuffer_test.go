package tbuffer

import (
	"fmt"
	"time"
)

type exampleStruct struct {
	val float64
}

func Example_timedBuffer() {
	b1 := New(10, time.Second, func(i []int) {
		var result int
		for _, v := range i {
			result += v
		}
		fmt.Println(result)
	})
	b1.Put(123)
	b1.Put(234)

	b2 := New(10, time.Second, func(i []exampleStruct) {
		var result float64
		for _, v := range i {
			result += v.val
		}
		fmt.Println(result)
	})
	b2.Put(exampleStruct{123})
	b2.Put(exampleStruct{234})
}
