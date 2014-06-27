package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

type Beer struct {
	Upc  string
	Type string
	Name string
	Qty  int
}

func main() {
	fmt.Println("Drink beer")
	host := "localhost"
	port := 3306

	mysqlAddr := fmt.Sprintf("root@tcp(%s:%d)/beerventory", host, port)
	db, err := sql.Open("mysql", mysqlAddr)
	if err != nil {
		log.Fatalf("Could not connect to mysql at %s:%s", host, port)
	}

	m := martini.Classic()
	m.Use(SetJsonContentType)

	m.Get("/beer", func() (int, string) {
		fmt.Print("fuck")
		res, err := db.Query("Select * from beer")
		if err != nil {
			log.Printf("Couldn't query for beer")
			return 500, "No beer here"
		}
		statusCode, beersJson := QueryMakerZero(res)
		return statusCode, string(beersJson)
	})

	m.Get("/beer/:id", func(params martini.Params) (int, string) {
		log.Print("beer beer", params["id"])
		res, err := db.Query("Select * from beer where upc = ? limit 1", params["id"])
		if err != nil {
			log.Printf("Couldn't query for beer", err)
			return 500, "No beer here"
		}
		statusCode, beersJson := QueryMakerZero(res)
		return statusCode, string(beersJson)
	})

	m.Post("/beer", func(req http.Request) string {

		log.Print("Post beer")
		// err := db.Exec("", ...)
		return "added beer"
	})

	m.Put("/beer/:id", func(params martini.Params) string {
		return fmt.Sprintf("edited beer %s", params["id"])
	})

	m.Delete("/beer/:id", func(params martini.Params) string {
		return fmt.Sprintf("deleted beer %s", params["id"])
	})

	m.Post("/checkout", func() string {
		return "added beer"
	})

	m.Get("/fish", func() string {
		return "fish"
	})

	m.Run()
}

func SetJsonContentType(res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
}

func QueryMakerZero(res *sql.Rows) (int, string) {
	beers := make([]Beer, 0)
	for res.Next() {
		var curBeer Beer
		err := res.Scan(&curBeer.Upc, &curBeer.Type, &curBeer.Name, &curBeer.Qty)
		if err != nil {
			log.Print("No scan", err)
		}
		beers = append(beers, curBeer)
	}

	var beersJson []byte
	var err error
	if len(beers) == 1 {
		beersJson, err = json.Marshal(beers[0])
	} else {
		beersJson, err = json.Marshal(beers)
	}
	if err != nil {
		return 500, "Json marshalling error"
	}
	return 200, string(beersJson)
}
