package main

type Fanout interface {
	Register(output chan<- bool)
	NewListener() <-chan bool
}

type fanout struct {
	outputs []chan<- bool
}

func NewFanout(input <-chan bool) Fanout {
	fo := new(fanout)
	go transmitLoop(fo, input)
	return fo
}

func transmitLoop(fo *fanout, input <-chan bool) {
	for {
		a := <-input
		for _, oc := range fo.outputs {
			go dispatchBool(oc, a)
		}
	}
}

func dispatchBool(c chan<- bool, b bool) {
	c <- b
}

func (fo *fanout) Register(output chan<- bool) {
	fo.outputs = append(fo.outputs, output)
}

func (fo *fanout) NewListener() <-chan bool {
	output := make(chan bool)
	fo.Register(output)
	return output
}

type fanin struct {
	inputs []<-chan gameEvent
	output chan<- gameEvent
}

func NewFanIn(output chan<- gameEvent) *fanin {
	fi := &fanin{output: output}
	return fi
}

func (fo *fanin) Register(input <-chan gameEvent) {
	go dispatchGameEvent(fo.output, input)
}

func dispatchGameEvent(output chan<- gameEvent, input <-chan gameEvent) {
	for {
		output <- <-input
	}
}
