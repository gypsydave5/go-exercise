package rollandwait

import (
	"math/rand"
	"time"
)

func rollAndWait(name string, totalTime time.Duration) {
	t := time.Now()
	for {
			result :=
	}
}

func dieRoll() int {
	return rand.Int() * 6
}

type fanOut struct {
		output []chan int
		input chan int
}

func newFanout(input chan int) {
		
}
