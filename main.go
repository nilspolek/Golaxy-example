package main

import (
	"fmt"
	"time"

	"github.com/nilspolek/Golaxy/router"
	"github.com/nilspolek/Golaxy/variable"
)

type Home struct {
	counter *variable.Integer
}

func (h *Home) Increase() {
	for {
		time.Sleep(time.Millisecond * 100)
		h.counter.Set(h.counter.Get() + 1)
	}
}

func (h Home) Render() string {
	return fmt.Sprintf("<h1>counter: %d</h1>", h.counter.Get())
}

func main() {
	r := router.New()
	home := Home{counter: variable.NewInt(r)}
	go home.Increase()
	r.RegisterRoute("home", home)
	r.Render()
	select {}
}
