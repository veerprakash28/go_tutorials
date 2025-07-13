package main

import (
	"fmt"
)

func main() {
	// Arrays
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr2)

	intArr3 := [...]int32{1,2,3}
	fmt.Println(intArr3)

	// Slice

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))
	fmt.Println(intSlice3)

	// Map

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45, "John": 23}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])
	delete(myMap2, "Adam")
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Loops
	for name, age:= range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v:= range(intArr3) {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
}
