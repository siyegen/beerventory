package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

type BeerList []Beer

func (b *BeerList) JSON() ([]byte, error) {
	return json.Marshal(b)
}

type Beer struct {
	Upc  int
	Type string
	Name string
	Qty  int
}

func (b *Beer) JSON() ([]byte, error) {
	return json.Marshal(b)
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

	m.Post("/beer", func(req *http.Request) (int, string) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Print("Couldn't read post body", err)
			return 400, "fish"
		}
		log.Print("Post beer")
		var beer Beer
		err = json.Unmarshal(body, &beer)
		if err != nil {
			log.Print("Couldn't unmarshal json", err)
			return 400, "fish"
		}

		_, err = db.Exec(
			"insert into beer set upc=?, type=?, name=?, qty=?",
			beer.Upc, beer.Type, beer.Name, beer.Qty,
		)
		if err != nil {
			log.Print("Couldn't save beer", err)
			return 400, "fish"
		}

		beerJson, _ := beer.JSON()
		return 200, string(beerJson)
	})

	m.Put("/beer/:upc", func(req *http.Request, params martini.Params) (int, string) {
		upc := params["upc"]
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Print("Couldn't read post body", err)
			return 400, "fish"
		}
		log.Print("Post beer")
		var beer Beer
		err = json.Unmarshal(body, &beer)
		if err != nil {
			log.Print("Couldn't unmarshal json", err)
			return 400, "fish"
		}

		_, err = db.Exec(
			"update beer set type=?, name=?, qty=? where upc=?",
			beer.Type, beer.Name, beer.Qty, upc,
		)
		if err != nil {
			log.Print("Couldn't save beer", err)
			return 400, "fish"
		}

		beerJson, _ := beer.JSON()
		return 200, string(beerJson)
	})

	m.Delete("/beer/:upc", func(params martini.Params) (int, string) {
		_, err = db.Exec("delete from beer where upc=? limit 1", params["upc"])
		if err != nil {
			log.Print("Couldn't delete beer", err)
			return 400, "fish"
		}
		return 200, `{"deleted":true}`
	})

	m.Post("/checkout", func() string {
		return "added beer"
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
		err := res.Scan(&curBeer.Upc, &curBeer.Name, &curBeer.Type, &curBeer.Qty)
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
