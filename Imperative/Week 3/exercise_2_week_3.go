package main

import "fmt"

//Estefanía Perez Hidalgo

func sustsinPerdidda(array *[6]int, number int, position int, size int) (slice []int) {
	slice = array[:]
	newslice := (size - 1)
	for range slice {
		if newslice == (size - 1) {
			*&slice = append(slice, slice[newslice])
		}
		(slice[newslice+1]) = slice[newslice]
		if newslice == position {
			(slice[newslice]) = number
			break
		}
		newslice--
	}
	fmt.Printf("\nThe slice with the number %d added is: ", number)
	return
}
func sliceChange(array *[6]int, position int) {
	(*array)[position] = 41 //Changes the position 3=67 for 3=10.
	fmt.Printf("The slice with the position %d changed is: %d", position, *array)
}
func main() {
	const size = 6
	array := [size]int{36, 3, 67, 95, 0, 2389}
	fmt.Println("The first slice is:", array)
	sliceChange(&array, 3)
	slice := sustsinPerdidda(&array, 82, 4, 6)
	fmt.Println(slice)
}
