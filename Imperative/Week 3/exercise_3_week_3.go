package main

import (
	"fmt"
	"strings"
)

//Estefanía Pérez Hidalgo

func main() {
	string := "This is an exercise in Golang"
	substring := "exercise"
	fmt.Println("\nUsing the library:")
	libraryComparison(string, substring)
	fmt.Println("\nNot using the library:")
	withoutLibraryComparison(string, substring)
}
func libraryComparison(string string, substring string) {
	if strings.Contains(strings.ToLower(string), strings.ToLower(substring)) {
		fmt.Println("The substring is inserted the string:", true)
		return
	}
	fmt.Println("The substring is not inserted the string:", false)
}
func withoutLibraryComparison(string, substring string) {
	index := 0
	flag := false
	for i := 0; i < len(string); i++ {
		if index == len(substring) {
			flag = true
			break
		}
		if string[i] != substring[index] {
			if index >= 1 {
				index = 0
			}
			continue
		} else {
			index++
		}
	}
	if flag {
		fmt.Println("The substring is inserted the string:", flag)
		return
	}
	fmt.Println("The substring is not inserted the string:", flag)
}
