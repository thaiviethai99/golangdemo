package main

import "fmt"

func map1() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	delete(x, "key")
}

func map2() {
	x := make(map[string]int)
	x["key"] = 10

	value, ok := x["key"]
	fmt.Println(value, ok)

	value2, ok2 := x["key2"]
	fmt.Println(value2, ok2)
}
func main() {
	//var x map[string]int
	//x := make(map[string]int)
	/* elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be" : "Beryllium",
		"B" : "Boron",
		"C" : "Carbon",
		"N" : "Nitrogen",
		"O" : "Oxygen",
		"F" : "Fluorine",
		"Ne" : "Neon",
	} */
	//map1()
	map2()
}
