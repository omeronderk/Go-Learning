package main

import "fmt"

func main() {
	/*

	var names map[string]int //key value ilişkisi önce keyin tipi sonra value nin tipini belirtmem gerekiyor

	names = make(map[string]int, 0)

	names["ömer"] = 1
	names["önder"] = 2
	names["kargın"] = 3

	fmt.Println(names)*/
	names := map[string]int{
		"ömer":1,
		"önder":2,
		"Kargın":3,
	}
	delete(names,"önder")
	fmt.Println(names)
}