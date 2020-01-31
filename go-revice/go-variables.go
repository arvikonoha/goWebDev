package main

import (
	"fmt"
)

var numberOfTimes int
var num1, num2, num3 int = 1, 2, 3
var anotherNumberOfTimes = 10

func main() {
	shortHand := 10
	isalive := true
	var isdead bool
	var ishappy = false
	myComplex := 10 + 12i
	yes, no, maybe := "yes", "no", "maybe"
	// int8,int16,int32,int64,int128,float,float32,float64,rune,byte,uint8,uint16,uint
	fmt.Println(numberOfTimes, num1, num2, num3, anotherNumberOfTimes, shortHand, isalive, isdead, ishappy, myComplex, yes, no, maybe)

	// to change a string since string is immutable
	str := "Hello"
	by := []byte(str)
	by[0] = 'J'
	str = string(by)
	fmt.Println(str, by)

	str = str + " world" // you can do this of course
	str = str[3:]

	fmt.Println(str)

}
