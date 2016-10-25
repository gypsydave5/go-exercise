package main

import (
	"fmt"
	"strconv"
	"time"
)

const playerCount = 4

func main() {
	scores := make([]int, playerCount)
	stop := newTimer(time.Duration(time.Second * 60))
	stopFanout := NewFanout(stop)

	for i := range scores {
		player(strconv.Itoa(i), stopFanout.NewListener())
	}
}

type Fanout interface {
	Register(output chan<- bool)
	NewListener() <-chan bool
}

type fanout struct {
	outputs []chan<- bool
}

type gameEvent struct {
	player string
	roll   int
}

func NewFanout(input <-chan bool) Fanout {
	fo := new(fanout)

	go func() {
		for {
			a := <-input
			fmt.Println(a, fo.outputs)
			for _, oc := range fo.outputs {
				go dispatch(oc, a)
			}
		}
	}()

	return fo
}

func dispatch(c chan<- bool, b bool) {
	c <- b
}

func (fo *fanout) Register(output chan<- bool) {
	fo.outputs = append(fo.outputs, output)
}

func (fo *fanout) NewListener() <-chan bool {
	o := make(chan bool)
	fo.Register(o)
	return o
}

func player(name string, stop <-chan bool) (roll <-chan gameEvent) {
	roll = make(chan gameEvent)
	return roll
}

func newTimer(d time.Duration) <-chan bool {
	timer := make(chan bool)

	go func() {
		time.Sleep(d)
		timer <- true
	}()

	return timer
}
