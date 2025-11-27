package main

import "fmt"

func main() {

	/*total:=add(10,20)
	fmt.Println(total)*/
	/*total , difference , multiply := calculation(10,20)
	fmt.Println(total,difference,multiply)*/
	/*var numbers = []int{10,20,2,3,5}
	fmt.Println(sum_1(numbers))*/
	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(3,5,6,6,6,4))
}

/*sum1 den farklı olarak parametre olarak 3 int girdik ama 4 . bir değişken girmek istediğimde fonksiyonu tekrardan düzeltmem gerek bunun yerine sum fonksiyonunu daha farklı yaz*/
func sum(numbers ...int) int {
	sum :=0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
func sum_1(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value

	}
	return sum

}

func calculation(x int, y int) (int, int, int) {
	return x + y, x - y, x * y

}
func add(x int, y int) int {

	//fmt.Println(x+y)
	return (x + y)
}
