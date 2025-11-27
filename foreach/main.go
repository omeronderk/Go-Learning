package main

import "fmt"

func main() {

	//var numbers = []int{1, 2, 3, 4}

	/*for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/
	/*
	for index , value := range numbers{//arrayin üstünde kendi kendine gezecek index olarak indexi alacak değer olan kısmı ise value ye kendi atayacak 
		fmt.Println(index,value)// index görünmesin istiyorsan for döndgüsünde index yerine alt tire koyabilirsin 
	}*/
	/*var language = "Golang"

	for _ , character := range language{
		fmt.Printf("Character: %c\n" , character)
	}*/
	var names = map[string]int{
		"ömer":10,
		"önder":20,
		"kargın":30,
	}
	for key ,value := range names{
		fmt.Println(key,value)
	}
}