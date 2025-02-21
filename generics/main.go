package main

import (
	"cmp"
	"fmt"
	//"golang.org/x/exp/constraints"
)

func main() {
	//v1 := []float64{1, 3, 5, 45, 12, 223, 6, 92, 78, 102}
	//v2 := []float32{9, 23, 1, 23, 8, 98}

	//fmt.Println(Sum01(v1))
	//fmt.Println(Sum01(v2))

	//fmt.Println(Sum02(v1))
	//fmt.Println(Sum02(v2))

	//anyType(1, 8)
	//anyType("1", "3")
	//anyType(2, "as")
	//anyTypeTwo(1, "2")
	//comparableType(2,3)

	/*csInt := CustomSlice[int]{1, 2, 3, 4}
	fmt.Println(csInt)

	csStg := CustomSlice[string]{"a", "b", "4"}
	fmt.Println(csStg)

	x, y := Point(5), Point(2)

	fmt.Println(MinNumber(x, y))*/

	fd := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}}
	fd.Data.PrintOne()

	sd := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	sd.Data.PrintTwo()
}

func Sum01[N int | int32 | int64 | float64 | float32](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

type Number interface {
	int | int32 | int64 | float64 | float32
}

func Sum02[N Number](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
}

func orderValues[N cmp.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}

type CustomSlice[V int | string] []V

func MinNumber[T N1](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point int

type N1 interface {
	int | int32 | int64 | float64 | float32
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}
	MyFirstData struct{}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}
