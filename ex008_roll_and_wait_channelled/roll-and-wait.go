package main

import (
	"math/rand"
	"strconv"
	"time"
)

func main() {
	scores := make([]int, 10)
	stop := newTimer(time.Duration(time.Second * 60))
	stopFanout := NewFanout(stop)

	gameEvents := make(chan gameEvent)
	gameEventFanin := NewFanIn(gameEvents)

	for i := range scores {
		pchan := player(strconv.Itoa(i), stopFanout.NewListener())
		gameEventFanin.Register(pchan)
	}

	gameOver := <-stopFanout
}

type gameEvent struct {
	player string
	roll   int
}

func player(name string, stop <-chan bool) (roll <-chan gameEvent) {
	roll = make(chan gameEvent)
	return roll
}

func rollDie() int {
	return rand.Intn(5) + 1
}

func newTimer(d time.Duration) <-chan bool {
	timer := make(chan bool)

	go func() {
		time.Sleep(d)
		timer <- true
	}()

	return timer
}
