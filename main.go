package main

import (
	"github.com/nilspolek/Golaxy/router"
)

type Home struct{}

func (h Home) Render() string {
	return "<h1>hello world</h1>"
}

func main() {
	r := router.New()
	r.RegisterRoute("home", Home{})
	r.Render()
}
