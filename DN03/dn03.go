package main

import (
	"fmt"
	"sync"
)

type narocilo interface {
	obdelaj()
}

type izdelek struct {
	imeIzdelka string
	cena       float64
	teza       float64
}

func (i izdelek) obdelaj() {
	defer wg.Done()

	lock.Lock()
	stNarocil++
	fmt.Printf("Št. naročila: %d\nIme izdelka: %s\nCena: %.2f €\n Teža: %.3f kg\n\n", stNarocil, i.imeIzdelka, i.cena, i.teza)
	promet += i.cena
	lock.Unlock()
}

type eknjiga struct {
	naslovKnjige string
	cena         float64
}

func (e eknjiga) obdelaj() {
	defer wg.Done()

	lock.Lock()
	stNarocil++
	fmt.Printf("Št. naročila: %d\nNaslov knjige: %s\nCena: %.2f €\n\n", stNarocil, e.naslovKnjige, e.cena)

	promet += e.cena
	lock.Unlock()
}

type spletniTecaj struct {
	imeTecaja   string
	trajanjeUre int
	cenaUre     float64
}

func (st spletniTecaj) obdelaj() {
	defer wg.Done()

	lock.Lock()
	stNarocil++
	fmt.Printf("Št. naročila: %d\nIme tečaja: %s\nCena ure: %.2f €\n Trajanje: %d ur\n\n", stNarocil, st.imeTecaja, st.cenaUre, st.trajanjeUre)
	promet += st.cenaUre * float64(st.trajanjeUre)
	lock.Unlock()
}

var promet float64
var stNarocil int
var lock sync.Mutex
var wg sync.WaitGroup

func main() {

	narocila := []narocilo{
		izdelek{
			imeIzdelka: "Jabolko",
			cena:       0.4,
			teza:       0.3,
		},
		izdelek{
			imeIzdelka: "Palični mešalnik",
			cena:       150,
			teza:       2.4,
		},
		izdelek{
			imeIzdelka: "Šotor",
			cena:       500,
			teza:       45,
		},
		eknjiga{
			naslovKnjige: "Harry Potter in primanjkanje cvetličnih lončkov",
			cena:         15,
		},
		eknjiga{
			naslovKnjige: "27 odtenkov rdeče",
			cena:         21,
		},
		eknjiga{
			naslovKnjige: "SSKJ",
			cena:         30,
		},
		spletniTecaj{
			imeTecaja:   "HTML v 10 sekundah",
			trajanjeUre: 10,
			cenaUre:     10,
		},
		spletniTecaj{
			imeTecaja:   "Primanjkanje idej",
			trajanjeUre: 16,
			cenaUre:     3,
		},
		eknjiga{
			naslovKnjige: "Matematika 3: delovni zvezek",
			cena:         90,
		},
		izdelek{
			imeIzdelka: "Matematika 3: delovni zvesek fizična različica",
			cena:       100,
			teza:       1,
		},
	}

	for _, item := range narocila {
		wg.Add(1)
		go item.obdelaj()
	}

	wg.Wait()

	fmt.Printf("Skupno naročil: %d, skupno prometa: %.2f €\n", stNarocil, promet)
}
