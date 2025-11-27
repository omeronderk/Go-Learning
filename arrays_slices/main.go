package main

import "fmt"

func main() {
	/*
		var names [3]string
		names[0] = "ömer"
		names[1] = "Önder"
		names[2] = "Kargın"
	var names = [4]string{"Ömer", "Önder", "Kargın","Bilgisayar mühendisi"}//boyut belirli olduğu için boyutu aşan elemanı ekleyemem
	names[1]="Osman"
	fmt.Println(names[0:2])
	//esnek arraylar için slice ları kullanıcaz*/
	var names = []string{"ömer","önder","kargın"}
	names=append(names, "osman")
	fmt.Println(names)
}
