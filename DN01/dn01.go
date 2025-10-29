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
		"123": {
			ime:     "Franc",
			priimek: "Horvat",
			ocene:   []int{10, 9, 10, 9, 10, 10, 9},
		},
		"456": {
			ime:     "Ana",
			priimek: "Kovačič",
			ocene:   []int{6, 7, 6, 7, 6, 7, 6, 7},
		},
		"789": {
			ime:     "Janez",
			priimek: "Novak",
			ocene:   []int{4, 5, 6, 5, 4, 5, 4, 5, 4},
		},
		"234": {
			ime:     "Zdravo",
			priimek: "Neocenjen",
			ocene:   []int{7, 8, 9},
		},
	}

	fmt.Println("Začetno stanje:")
	izpisRedovalnice(studenti)
	fmt.Println()

	fmt.Println("Dodajanje ocen:")
	fmt.Println()
	dodajOceno(studenti, "123", 8)
	izpisRedovalnice(studenti)
	fmt.Println()

	dodajOceno(studenti, "123", 15)
	izpisRedovalnice(studenti)
	fmt.Println()

	dodajOceno(studenti, "abc", 4)
	izpisRedovalnice(studenti)
	fmt.Println()

	fmt.Println("Povprečje ocen:")
	fmt.Println()
	result := povprecje(studenti, "123")
	fmt.Printf("Povprečje študenta 123: %f\n", result)

	result2 := povprecje(studenti, "456")
	fmt.Printf("Povprečje študenta 456: %f\n", result2)

	result3 := povprecje(studenti, "789")
	fmt.Printf("Povprečje študenta 789: %f\n", result3)

	result4 := povprecje(studenti, "234")
	fmt.Printf("Povprečje študenta 234: %f\n", result4)

	result5 := povprecje(studenti, "abc")
	fmt.Printf("Povprečje študenta abc: %f\n", result5)

	fmt.Println()
	izpisiKoncniUspeh(studenti)
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

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1
	}

	if len(student.ocene) < 6 {
		return 0
	}

	avg := 0.0
	for _, ocena := range student.ocene {
		avg += float64(ocena)
	}

	avg /= float64(len(student.ocene))
	return avg
}

func izpisiKoncniUspeh(studenti map[string]Student) {
	fmt.Println("KONČNI USPEH:")
	for vpisna, student := range studenti {
		studentAvg := povprecje(studenti, vpisna)

		fmt.Printf("%s: povprečna ocena %.1f -> ", student, studentAvg)

		switch {
		case studentAvg >= 9:
			fmt.Printf("Odličen študent!\n")
		case studentAvg < 9 && studentAvg >= 6:
			fmt.Printf("Povprečen študent\n")
		case studentAvg < 6:
			fmt.Printf("Neuspešen študent\n")
		}
	}
}
