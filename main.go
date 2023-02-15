package main

import (
	"fmt"

	"github.com/xiusin/phpingo/array"
)

type S struct {
	Name string
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr1 := []int{2, 12, 32, 9}
	arr2 := []int{12, 0}

	// fmt.Println(array.Reverse[int](arr))
	// fmt.Println(array.Slice[int](arr, 2, 200))
	// fmt.Println(array.Splice[int](&arr, 2, 4, []int{21, 22, 23, 24, 25}))
	// fmt.Println(arr)
	fmt.Println(array.Intersect[int](arr1, arr, arr2))

	// fmt.Println(array.Chunk[int](arr, 2))
	// fmt.Println("sum", array.Sum[int](arr))
	// array.Push[int](&arr, 10)
	// fmt.Println("sum", array.Sum[int](arr))
	// val, _ := array.Pop[int](&arr)
	// fmt.Println(val, arr)
	// var arr1 = []int{1}
	// fmt.Println(array.Shift[int](&arr1))
	// fmt.Println(arr1)

	var mapArr = []map[string]string{{"name": "张三"}}

	var structArr = []S{{Name: "张三"}, {Name: "李四"}}

	fmt.Println(array.Column[string](&mapArr, "name"))
	fmt.Println(array.Column[string](&structArr, "Name"))

	// fmt.Println(array.Values[int](arr))
}
