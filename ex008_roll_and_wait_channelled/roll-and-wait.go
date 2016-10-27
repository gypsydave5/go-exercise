package main

import (
	"fmt"
	"math/rand"
	"time"
)

var players = []string{
	"Jo",
	"Alex",
	"Toni",
}

type gameEvent struct {
	player string
	roll   int
}

type die <-chan int

func main() {
	scores := make(map[string]int)
	stopFanout := NewFanout(newTimer(time.Duration(time.Second * 10)))
	stop := stopFanout.NewListener()

	gameEvents := make(chan gameEvent)
	go gameLoop(gameEvents, scores)

	gameEventFanin := NewFanIn(gameEvents)

	for _, name := range players {
		pchan := player(name, stopFanout.NewListener())
		gameEventFanin.Register(pchan)
	}

	<-stop
	fmt.Println(scores)
}

func gameLoop(events <-chan gameEvent, scores map[string]int) {
	for {
		e := <-events
		fmt.Println(e.String())
		scores[e.player] += e.roll
	}
}

func (e *gameEvent) String() string {
	return fmt.Sprintf("%s rolled %d, waiting %d seconds", e.player, e.roll, 6-e.roll)
}

func player(name string, stop <-chan bool) <-chan gameEvent {
	events := make(chan gameEvent)
	die := newDie()

	go playerLoop(name, stop, events, die)
	return events
}

func playerLoop(name string, stop <-chan bool, events chan<- gameEvent, d die) {
	for {
		select {
		case x := <-d:
			events <- gameEvent{name, x}
			time.Sleep(time.Second * time.Duration(6-x))
		case <-stop:
			fmt.Println("%s says Game Over", name)
			return
		}
	}
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

func newDie() die {
	die := make(chan int)
	go roll(die)
	return die
}

func roll(die chan<- int) {
	for {
		die <- rand.Intn(5) + 1
	}
}
