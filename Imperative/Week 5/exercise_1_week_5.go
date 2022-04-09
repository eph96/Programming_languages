package main

import "fmt"

//Estefanía Pérez Hidalgo

type Person struct {
	name  string
	id    [8]int
	score int
}

func sampling(group []Person, sampling [8]int) {
	cont := 0
	index := 0
	amount := 0
	for i := 0; i < len(group); i++ {
		cont = 0
		index = 0
		for y := 0; y < len(sampling); y++ {
			if sampling[y] == group[i].id[index] {
				cont++
			}
			if y == (len(sampling) - 1) {
				group[i].score = cont
			}
			index++
		}
	}
	for j := 0; j < len(group); j++ {
		index = 0
		amount = 0
		for index < len(group) {
			if group[j].score < group[index].score {
				amount++
			}
			if amount >= 3 {
				group = append(group[:j], group[j+1:]...)
				j--
			}
			index++
		}
	}
	fmt.Println("\nThe top 3 people with the highest scores are:")
	for _, r := range group {
		fmt.Println(r.name, "with an score of:", r.score)
	}
}
func main() {
	basesampling := [8]int{0, 1, 0, 0, 1, 0, 0, 1}
	personas := []Person{
		{"María", [8]int{0, 1, 0, 1, 1, 0, 0, 1}, 0},
		{"Juan", [8]int{0, 1, 0, 1, 0, 1, 0, 1}, 0},
		{"Andrea", [8]int{0, 0, 0, 0, 0, 1, 1, 1}, 0},
		{"Luis", [8]int{0, 0, 1, 1, 1, 0, 0, 1}, 0},
		{"Sonia", [8]int{1, 0, 0, 1, 0, 1, 0, 0}, 0},
		{"Carlos", [8]int{1, 1, 0, 1, 0, 0, 1, 0}, 0}}
	sampling(personas, basesampling)
}
