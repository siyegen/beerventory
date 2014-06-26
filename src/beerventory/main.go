package main

import (
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	fmt.Println("Drink beer")
	m := martini.Classic()

	m.Get("/beer", func() string {
		return "beer"
	})

	m.Post("/beer", func() string {
		return "added beer"
	})

	m.Put("/beer/:id", func(params martini.Params) string {
		return fmt.Sprintf("edited beer %d", params["id"])
	})

	m.Delete("/beer/:id", func(params martini.Params) string {
		return fmt.Sprintf("deleted beer %d", params["id"])
	})

	m.Post("/checkout", func() string {
		return "added beer"
	})

	m.Run()
}
