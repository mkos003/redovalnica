package main

import (
	"fmt"

	"github.com/mkos003/redovalnica/redovalnica"
)

func main() {
	fmt.Print("Hello, World!")

	studenti := make(map[string]redovalnica.Student)
	studenti["1234"] = redovalnica.Student{"Janez", "Novak", []int{10, 6, 5}}
	studenti["4321"] = redovalnica.Student{"Ivan", "Novak", []int{8, 7, 6}}
	studenti["5678"] = redovalnica.Student{"Petra", "Novak", []int{5, 5, 6}}
	studenti["8765"] = redovalnica.Student{"Marko", "Novak", []int{9, 10, 9}}

	r := redovalnica.NewRedovalnica(studenti)
	r.IzpisVsehOcen()
	r.DodajOceno("1234", 10)
	r.IzpisiKoncniUspeh()
}
