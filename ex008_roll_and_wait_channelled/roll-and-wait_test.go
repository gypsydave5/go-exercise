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
