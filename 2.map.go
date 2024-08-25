package main

import "fmt"

func map_func() {
	//empty map creation without using make keyword
	var map1 map[int]int
	fmt.Println("Map1: ", map1)
	//without using keyword var and make ->shorthand method
	map2 := map[int]string{} //empty map
	fmt.Println(map2)
	//initialising map with some key-value pairs
	map3 := map[int]string{1: "hello", 2: "world"}
	fmt.Println(map3)
	//using make keyword
	var map4 = make(map[int]int, 5) //initial capacity 5
	//assigning values to map
	map4[1] = 111
	map4[2] = 222
	map4[3] = 333
	fmt.Println(map4)
	//iterating through map

}
