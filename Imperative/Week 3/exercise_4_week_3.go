package main

import "fmt"

//Estefanía Perez Hidalgo

type Shoes struct {
	brand string
	price int
	size  int
}

var shoes []Shoes

func availableShoes(brand string, price int, size int) []Shoes {
	//First, the array of shoes is created.
	shoe := Shoes{brand, price, size}
	shoes = append(shoes, shoe)
	return shoes
}
func validation(shoes []Shoes, min int, max int) []Shoes {
	//Then, this function goes over the array of shoes
	//and validate the sizes.
	var slice []Shoes
	for i := 0; i < len(shoes); i++ {
		if shoes[i].size >= min && shoes[i].size <= max {
			slice = append(slice, shoes[i])
		}
	}
	fmt.Println("\nThe new list is:")
	for _, r := range slice {
		fmt.Println("The brand:", r.brand, " price:", r.price, " size:", r.size)
	}
	return slice
}
func main() {
	availableShoes("Nike", 73000, 45)
	availableShoes("Flexi", 85000, 35)
	availableShoes("Botero", 30000, 34)
	availableShoes("S&J", 42500, 40)
	availableShoes("Hush Puppies", 10450, 38)
	fmt.Println("\nThe firts list is:")
	for _, r := range shoes {
		fmt.Println("The brand:", r.brand, " price:", r.price, " size:", r.size)
	}
	validation(shoes, 35, 38)
}
