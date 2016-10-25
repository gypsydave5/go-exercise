package main

import (
	"fmt"
	"testing"
)

func TestFanout(t *testing.T) {
	inputChan := make(chan bool)
	fo := NewFanout(inputChan)

	outputChanOne := make(chan bool)
	fo.Register(outputChanOne)

	outputChanTwo := fo.NewListener()

	fmt.Printf("%#v\n", fo)
	inputChan <- true

	if x := <-outputChanTwo; x != true {
		t.Error("Did not recieve message on channel two")
	}

	if x := <-outputChanOne; x != true {
		t.Error("Did not recieve message on channel one")
	}
}

func TestFanIn(t *testing.T) {
	outputChan := make(chan gameEvent)
	fi := NewFanIn(outputChan)

	inputChanOne := make(chan gameEvent)
	inputChanTwo := make(chan gameEvent)

	fi.Register(inputChanOne)
	fi.Register(inputChanTwo)

	inputChanOne <- gameEvent{"bob", 5}
	inputChanTwo <- gameEvent{"joe", 6}

	if x := <-outputChan; x.player != "bob" {
		t.Error("Did not recieve message")
	}

	if x := <-outputChan; x.player != "joe" {
		t.Error("Did not recieve message")
	}

}
