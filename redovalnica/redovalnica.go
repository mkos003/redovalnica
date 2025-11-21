// Package redovalnica implements a simple gradebook system for managing students
// and their grades.
//
// The package allows storing students, assigning grades, calculating averages,
// and determining their final academic success. Each student is identified by
// a unique enrolment number (vpisna številka).
//
// Example usage:
//
//	r := redovalnica.NewRedovalnica(map[string]redovalnica.Student{
//	    "63120001": {ime: "Ana", priimek: "Novak", ocene: []int{8, 9}},
//	})
//
//	r.DodajOceno("63120001", 10)
//	r.IzpisVsehOcen()
//	r.IzpisiKoncniUspeh()
//
// The gradebook provides the following functionalities:
//   - Adding grades for a student
//   - Printing all grades for all students
//   - Calculating grade averages
//   - Evaluating final academic success based on average grade
package redovalnica

import (
	"fmt"
)

// Student represents a single student and their grades.
type Student struct {
	ime     string
	priimek string
	ocene   []int
}

// NewStudent creates a new student with the given name, surname, and grades.
func NewStudent(ime, priimek string, ocene []int) Student {
	return Student{
		ime:     ime,
		priimek: priimek,
		ocene:   ocene,
	}
}

// Redovalnica represents a gradebook containing multiple students.
type Redovalnica struct {
	studenti map[string]Student
}

// NewRedovalnica creates a new gradebook with an initial map of students.
func NewRedovalnica(s map[string]Student) *Redovalnica {
	r := &Redovalnica{
		studenti: s,
	}
	return r
}

// DodajOceno adds a new grade for a student with the given enrolment number.
// The method checks whether the student exists and whether the grade is valid.
func (r *Redovalnica) DodajOceno(vpisnaStevilka string, ocena int) {
	_, ok := r.studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Napaka! Študent z vpisno številko ", vpisnaStevilka, " ne obstaja.")
		return
	}

	if ocena < 1 || ocena > 10 {
		fmt.Println("Napaka! Ocena mora biti med 1 in 10.")
		return
	}

	student := r.studenti[vpisnaStevilka]
	student.ocene = append(student.ocene, ocena)

	// Kopijo studenta zapišemo nazaj na originalno mesto
	r.studenti[vpisnaStevilka] = student
}

// IzpisVsehOcen prints all students and their corresponding grades.
func (r *Redovalnica) IzpisVsehOcen() {
	fmt.Printf("Redovalnica:\n")
	for vpisna, student := range r.studenti {
		fmt.Printf("%s - %s %s: [", vpisna, student.ime, student.priimek)
		for _, ocena := range student.ocene {
			fmt.Printf(" %d", ocena)
		}
		fmt.Printf(" ]\n")
	}
}

// IzpisiKoncniUspeh prints the final academic success for each student,
// determined by their average grade.
func (r *Redovalnica) IzpisiKoncniUspeh() {
	for vpisna, student := range r.studenti {
		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", student.ime, student.priimek, povprecje(r, vpisna), koncniUspeh(povprecje(r, vpisna)))
	}
	fmt.Printf("\n")
}

// koncniUspeh returns a string describing final academic success
// based on the student's average grade.
func koncniUspeh(povprecje float64) string {
	switch {
	case povprecje >= 9.0:
		return "Odličen študent!"
	case povprecje >= 6.0 && povprecje < 9.0:
		return "Povprečen študent!"
	case povprecje < 6.0:
		return "Neuspešen študent!"
	default:
		return "Napaka!"
	}
}

// povprecje calculates the average grade for a student. If the student does not
// exist, the function prints an error and returns -1.
func povprecje(r *Redovalnica, vpisnaStevilka string) float64 {
	_, ok := r.studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Napaka! Študent z vpisno številko ", vpisnaStevilka, " ne obstaja.")
		return -1.0
	}

	student := r.studenti[vpisnaStevilka]
	var sestevek int
	for _, ocena := range student.ocene {
		sestevek += ocena
	}

	return float64(sestevek) / float64(len(student.ocene))
}
