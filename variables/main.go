package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	//Variables
	var myIntVar int
	myIntVar = 12
	fmt.Println("Mi variable es:", myIntVar)

	//no acepta valores negativos uint
	var myUintVar uint
	myUintVar = 12
	fmt.Println("Mi variable es:", myUintVar)

	var myStringVar string
	myStringVar = "Mi variable string"
	fmt.Println("Mi variable es:", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("Mi variable es:", myBoolVar)

	//Constantes
	const myIntConst int = 12
	fmt.Println("Mi constante es: ", myIntConst)

	const myFirstStringConst = "a12"
	fmt.Println("Mi constante es:", myFirstStringConst)

	const myStringConst string = "test"
	fmt.Println("Mi constante es: ", myStringConst)

	//Bloques

	//Definir variables dentro de un scope {}, no usa la memoria en la ejecución
	//vive solo dentro del scope {}
	{
		myScopeVariable := 40
		{
			fmt.Println("Mi variable: ", myScopeVariable)
		}
		fmt.Println("Mi variable: ", myScopeVariable)
	}

	//Convertir a String
	{
		floatVar := 33.11
		fmt.Println("type: %T, value: %f\n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Println("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Println("type: %T, value: %d\n", intVar, intVar)

		intStrVar := fmt.Sprintf("%d - %d", intVar, 23)
		fmt.Println("type: %T, value: %s\n", intStrVar, intStrVar)

		intVar2 := 342
		intStrVar2 := strconv.Itoa(intVar2)
		fmt.Println("type: %T, value: %s\n", intStrVar2, intStrVar2)
	}

	//Convertir a numero
	strIntVar2, err := strconv.Atoi("15")
	fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar2, strIntVar2, err)

	strIntVar3, err := strconv.ParseInt("10", 10, 64)
	fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar3, strIntVar3, err)

	strFloatVar, err := strconv.ParseFloat("-11-2a", 64)
	fmt.Printf("type: %T, value: %f, err: %v\n", strFloatVar, strFloatVar, err)

	//Operadores
	yearsOld := 33

	fmt.Println(yearsOld > 30)  //true
	fmt.Println(yearsOld < 33)  //false
	fmt.Println(yearsOld <= 33) //true
	fmt.Println(yearsOld >= 40) //false
	fmt.Println(yearsOld == 33) //true

	fmt.Println(yearsOld < 33 || yearsOld == 33) //(false o true) = false
	fmt.Println(yearsOld < 33 || yearsOld == 34) //(false o false) = false
	fmt.Println(yearsOld < 40 || yearsOld == 33) //(true o true) = true

	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(yearsOld < 40)    // true
	fmt.Println(!(yearsOld < 40)) // false

	fmt.Println(yearsOld < 25 && yearsOld == 33 || yearsOld == 40)   //(false and true or true) = true
	fmt.Println(yearsOld < 25 && (yearsOld == 33 || yearsOld == 40)) //(false and (true or true)) = false

	//Condicional if
	yearsOld1 := 16

	if yearsOld1 > 18 {
		fmt.Println("%d es mayor a 18\n", yearsOld1)
	}
	boolVar := true

	if boolVar {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	if value := true; value {
		fmt.Println("es verdadero")
	}

	number := 3

	if number == 1 {
		fmt.Println("uno")
	} else if number == 2 {
		fmt.Println("dos")
	} else if number == 3 {
		fmt.Println("tres")
	} else {
		fmt.Println("ninguna es valida")
	}

	//Condicional Switch
	switch number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("ninguna es valida")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("ninguna es valida")
	}

	switch {
	case number > 0:
		fmt.Println("positivo")
	case number < 0:
		fmt.Println("negativo")
	case number == 0:
		fmt.Println("es cero")
	}

	//Bucle for
	sum := 0

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
		sum++
	}

	fmt.Println(sum)

	sum = 0

	for sum < 1000 {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	//Arrays
	var myIntVar1 int = 30
	fmt.Println("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVar1, myIntVar1, unsafe.Sizeof(myIntVar1), unsafe.Sizeof(myIntVar1)*8)

	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, " - size: ", len(myArrayVar2))

	myArrayVar1[0] = 2
	myArrayVar1[1] = 5
	myArrayVar1[2] = 9
	//fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1), " - get index 1: ", myArrayVar1[1] )

	fmt.Println("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	/*
		for i := range myArrayVar1{
			fmt.Println("index: ", i, " - value: ", myArrayVar1[i])
		}
	*/

	for _, v := range myArrayVar1 {
		fmt.Println("value: ", v)
	}

	//Slices
	myArrayVar := [5]int{3, 6, 6, 10, 16}
	fmt.Println("array: ", myArrayVar, " - len: ", len(myArrayVar))

	mySliceVar := []int{}

	mySliceVar = append(mySliceVar, 12, 34, 54, 31, 12)
	fmt.Println("slice: ", mySliceVar, " - len: ", len(mySliceVar))

	mySliceVar2 := myArrayVar[2:4]

	mySliceVar2[0] = 19
	fmt.Println("slice: ", mySliceVar2, " - len: ", len(mySliceVar2), " - address: ", &mySliceVar2[0])

	fmt.Println("array: ", mySliceVar, " - len: ", len(mySliceVar), " - address: ", &mySliceVar2[2])

	mySliceVar3 := mySliceVar[2:]
	fmt.Println(mySliceVar3)

	//mySliceVar4 := make([]int, 80)
	//fmt.Println("slice: ", mySliceVar4, " - len: ", len(mySliceVar4))

	mySliceVar5 := []int{1, 2, 6, 11, 20, 5, 1, 0}
	fmt.Println(mySliceVar5)

	//Maps
	map1 := make(map[int]string)

	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[99] = "Z"
	map1[-5] = "a"

	fmt.Println(map1)
	fmt.Println(map1[-5])

	map2 := make(map[int][]string)
	map2[1] = []string{"A", "B"}
	map2[3] = []string{"Z", "A", "9"}

	fmt.Println(map2)
	fmt.Println(map2[3])

	map1[1] = "LA"

	delete(map1, 2)
	delete(map1, 3)
	map1[9] = ""
	fmt.Println(map1)
	fmt.Println(map2[8])

	v, ok := map1[9]
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	v, ok = map1[9]
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	map3 := map[int]string{
		4: "A",
		1: "N",
		8: "90",
	}
	fmt.Println(map3)

	for k := range map1 {
		fmt.Println("key: ", " - value: ", map1[k])
	}

	for key, value := range map2 {
		fmt.Println("key: ", key)

		for _, v := range value {
			fmt.Println("Value: ", v)
		}

		fmt.Println()
	}
}
