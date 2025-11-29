package main

import "fmt"

func main() {

	// var a int
	// var p *int
	// p = &a
	// a = 10
	// /*
	// fmt.Println(a)
	// fmt.Println(p)// adresi almak
	// fmt.Println(*p)// o adresteki değeri yazdırmak*/
	// *p = 20 // pointer üzerinden değişken değeri değiştirmek
	// fmt.Println(a),
	// var a = 10
	// var b int
	// var p *int

	// b=a // eşitlik değil sadece a daki değeri b ye kopyalamak
	// p = &a

	// *p = 20
	// fmt.Println(a,b)
	// var a int = 10
	// fmt.Println(a)
	// add12(a)
	// fmt.Println(a) //fonksiyona girdikten sonra a değeri değişmiyor çünkü a nın değeri sadece kopyalanıyor eğer a nın değişmesini istersen pointer üzerinden yapabilirsin
	// add12Pointer(&a)
	// fmt.Println(a)
	var numbers = []int {1,2,3} // slice- aray olduğunda ekstra pointer kullanmama gerek yok zaten referans olarak gönderiyor pointerları ilkel veri tiplerindeki değişikliklerde kullanıyoruz
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
	
}
func changeValue(numbers []int){
	numbers[0] = 1000
}
func add12(x int) { //pas by value
	x = x + 12
}
func add12Pointer(x *int){ // pass by reference
	*x = *x + 12
}
