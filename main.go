package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Item struct {
	Name  string
	Value int
}

const dbURL = "postgres://Nore@localhost:5432/testing?sslmode=disable"

var db *sqlx.DB
var err error

func saveItemToDB(name, value string) (*Item, error) {
	query := "INSERT INTO items (name, value) VALUES ($1, $2) RETURNING *"

	item := &Item{}
	err := db.Get(item, query, name, value)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func fetchAllItems() []*Item {
	query := "SELECT * FROM items"
	items := []*Item{}
	err := db.Select(&items, query)
	if err != nil {
		panic("failed to fetch all items from the database")
	}
	return items
}

func calcTotalItemWorth(items []*Item) int {
	totalWorth := 0
	for _, item := range items {
		totalWorth = totalWorth + item.Value
	}
	return totalWorth
}

func main() {
	db, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.Exec("CREATE TABLE items (name varchar(100), value int)")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		items := fetchAllItems()
		totalWorth := calcTotalItemWorth(items)

		c.HTML(200, "index.tmpl", gin.H{
			"items":       items,
			"total_worth": totalWorth,
		})
	})

	r.GET("/new", func(c *gin.Context) {
		c.HTML(200, "new.tmpl", gin.H{})
	})

	r.POST("/new", func(c *gin.Context) {
		item, err := saveItemToDB(c.PostForm("item_name"), c.PostForm("value"))
		if err != nil {
			panic(err.Error())
		}

		c.HTML(200, "item.tmpl", gin.H{
			"name":  item.Name,
			"value": item.Value,
		})
	})

	r.Run(":8080")
}
