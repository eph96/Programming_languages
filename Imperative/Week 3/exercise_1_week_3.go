package main

import "fmt"

//Estefanía Perez Hidalgo

func seekWord(chain string, word string) string {
	var tempWord string
	var tempLetter string
	num := 0
	for _, v := range chain {
		if v == int32(word[num]) {
			tempWord = tempLetter + string(word[num])
			tempLetter = tempWord
			if word == tempWord {
				return tempWord
			}
			num++
			continue
		}
		tempWord = ""
	}
	return tempWord
}
func main() {
	cadena := "Escriba un programa que a partir de una cadena."
	wordSought := "partir"
	fmt.Printf("\nThe word: \"%v\" was found.\n", seekWord(cadena, wordSought))
}
