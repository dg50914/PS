package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Meritev struct {
	vrsta    string
	vrednost float32
}

func main() {
	bufferSize := 3
	measurerReportCount := 10

	measurementChan := make(chan Meritev, bufferSize)
	keyboardChan := make(chan bool)

	go measurer(measurementChan, "temperatura", measurerReportCount, -20, 40)
	go measurer(measurementChan, "vlaga", measurerReportCount, 30, 80)
	go measurer(measurementChan, "tlak", measurerReportCount, 900, 1100)

	go readKey(keyboardChan)

	for {
		select {
		case measurement := <-measurementChan:
			switch measurement.vrsta {
			case "temperatura":
				fmt.Println("Meritev temperature: ", measurement.vrednost, "°C")
			case "vlaga":
				fmt.Println("Meritev vlage: ", measurement.vrednost, "%")
			case "tlak":
				fmt.Println("Meritev tlaka: ", measurement.vrednost, "hPa")
			}
		case <-keyboardChan:
			fmt.Println("Prejel pritisk Enter, exiting")
			return

		case <-time.After(time.Second * 5):
			fmt.Println("Sistem neodziven 5 sekund, exiting")
			return
		}
	}
}

func measurer(comm chan<- Meritev, measurementType string, count int, rangeMin, rangeMax int) {
	for range count {
		time.Sleep(time.Millisecond * 100)

		measurement := Meritev{
			vrsta:    measurementType,
			vrednost: float32(rangeMin) + float32(rangeMax-rangeMin)*rand.Float32(),
		}
		comm <- measurement
	}
}

func readKey(input chan bool) {
	_, _ = fmt.Scanln()
	input <- true
}
