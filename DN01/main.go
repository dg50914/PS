package main

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func (s Student) String() string {
	return fmt.Sprintf("%s %s", s.ime, s.priimek)
}

func main() {

	studenti := map[string]Student{
		"123": Student{
			ime:     "Franc",
			priimek: "Horvat",
			ocene:   []int{10, 9, 8, 7, 6},
		},
		"456": Student{
			ime:     "Ana",
			priimek: "Kovačič",
			ocene:   []int{10, 9, 8, 7, 6},
		},
		"789": Student{
			ime:     "Janez",
			priimek: "Novak",
			ocene:   []int{10, 9, 8, 7, 6},
		},
	}

	fmt.Println("Začetno stanje:")
	izpisRedovalnice(studenti)
	fmt.Println()

	fmt.Println("Dodajanje ocen:")
	fmt.Println()
	dodajOceno(studenti, "123", 5)
	izpisRedovalnice(studenti)
	fmt.Println()

	dodajOceno(studenti, "123", 15)
	izpisRedovalnice(studenti)
	fmt.Println()

	dodajOceno(studenti, "abc", 4)
	izpisRedovalnice(studenti)
	fmt.Println()
}

func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if !(ocena >= 0 && ocena <= 10) {
		fmt.Printf("Ocena %d ni v intervalu [0, 10]\n", ocena)
		return
	}

	student, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Printf("Študent z vpisno številko %s ne obstaja\n", vpisnaStevilka)
		return
	}

	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
	fmt.Printf("Ocena %d dodana študentu/ki %s\n", ocena, student)
}

func izpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s: ", vpisna, student)
		fmt.Println(student.ocene)
	}
}
