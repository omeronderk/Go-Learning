package main

import (
	"fmt"
	
)

func main() {
/*
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	isInStock=true
	discount=0.37
	quantity = 10
	productName = "GoLang Dersleri"
	fmt.Println(isInStock,reflect.TypeOf(isInStock))
	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount,reflect.TypeOf(discount))
	*/

	/*var productName string="GoLang Dersleri"
	var name ="Ömer"
	fmt.Println(productName,reflect.TypeOf(productName))
	fmt.Println(name,reflect.TypeOf(name))*/
	/*productName  :="Golang Dersleri"
	fmt.Println(productName,reflect.TypeOf(productName))*/
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	isInStock=true
	discount=0.37
	quantity = 10
	productName = "GoLang Dersleri"

	//fmt.Println("productName",productName,"quantity",quantity)
	fmt.Printf("Product name: %s, Quantity:%d, Discount:%f, IsInStock %t",productName,quantity,discount,isInStock)
}

