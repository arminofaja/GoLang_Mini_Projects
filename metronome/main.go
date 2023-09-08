package main

import (
	"fmt"
	"time"
)

func metronome(bpm int) {
	interval := time.Minute / time.Duration(bpm)

	for i := 1; ; i++ {
		fmt.Print(i, " ")
		if i == 4 {
			i = 0
		}
		time.Sleep(interval)

		fmt.Print("\x07") // This produces a beep sound in some terminals

	}
}

func main() {
	bpm := 120 // Adjust the BPM as needed

	go metronome(bpm)

	fmt.Printf("Metronome started at %d BPM. Press Enter to stop.\n", bpm)
	fmt.Scanln() // Wait for Enter key to stop the metronome

	// On some platforms, the beep sound may not work as expected.
	fmt.Println("Metronome stopped.")
}
