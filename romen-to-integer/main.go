package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

type RomanNumber struct {
	Sysmbol string
	Value   int
}

var gRoman = []RomanNumber{
	{Sysmbol: "M", Value: 1000},
	{Sysmbol: "D", Value: 500},
	{Sysmbol: "C", Value: 100},
	{Sysmbol: "L", Value: 50},
	{Sysmbol: "X", Value: 10},
	{Sysmbol: "V", Value: 5},
	{Sysmbol: "I", Value: 1},
}

func romanToInt(s string) int {
	var count int
	text := strings.Split(s, "")
	//M = 1000, CM = 900, XC = 90 and IV = 4.
	//L = 50, V= 5, III = 3.
	skip := 0
	for i := 0; i < len(text); i++ {

		if skip > 0 {
			skip -= 1
			fmt.Println("skip")
			continue
		}

		var j = i
		if (i + 1) < len(text) {
			j += 1
		}

		//fmt.Printf("%v  %v \n ", getValue(text[i]), getValue(text[j]))

		if getValue(text[i]) >= getValue(text[j]) {

			count += getValue(text[i])

		} else {
			count += getValue(text[j]) - getValue(text[i])
			skip += 1

		}

	}

	return count
}

func getValue(s string) int {
	for _, roman := range gRoman {
		if s == roman.Sysmbol {
			return roman.Value
		}
	}
	return 0
}
