package main

import (
	"fmt"
	"time"

	"github.com/YSZhuoyang/go-dispatcher/dispatcher"
)

type Dormir interface {
	Do()
}

type dormir struct {
	time       int
	dispatcher dispatcher.Dispatcher
	end        chan any
}

func NewDormir(time int, disp dispatcher.Dispatcher) Dormir {
	return &dormir{
		time:       time,
		dispatcher: disp,
		end:        make(chan any),
	}
}

func (p *dormir) Do() {
	fmt.Printf("Comienza job de %d segundos\n", p.time)
	time.Sleep(time.Duration(p.time) * time.Second)
	fmt.Printf("Termina job de %d segundos\n", p.time)
	fmt.Printf("Workers disponibles: %d \n", p.dispatcher.GetNumWorkersAvail()+1)
}

var GlobalInt int = 10

func main() {
	times := [10]int{10, 3, 5, 6, 7, 9, 2, 6, 4, 12}
	disp, _ := dispatcher.NewDispatcher(3)

	fmt.Println("Workers totales", disp.GetTotalNumWorkers())
	fmt.Println("Workers disponibles", disp.GetNumWorkersAvail())

	for _, t := range times {
		dormir := NewDormir(t, disp)
		disp.Dispatch(dormir)
	}

	fmt.Println("Workers dispatcheado")
	fmt.Println("Workers disponibles", disp.GetNumWorkersAvail())
	disp.Finalize()

}
