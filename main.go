package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Item struct {
	Name  string
	Value int
}

// type UserSession struct {
// 	Items []Item
// }

// type Page struct {
// 	gorm.Model

// 	Title string `gorm:"column:title"`
// 	Body  []byte `gorm:"column:body"`
// }

const dbURL = "postgres://Nore@localhost:5432/testing?sslmode=disable"

func main() {
	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.Exec("CREATE TABLE items (name varchar(100), price int)")

	// db.AutoMigrate(&Item{}) // if there's no table in the db, it will create one

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/edit/:title/:position", func(c *gin.Context) {
		// title := c.Param("title")
		fmt.Println(c.Accepted)
		c.HTML(200, "edit.tmpl", gin.H{})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("item_name")
		value := c.PostForm("value")

		c.HTML(200, "item.tmpl", gin.H{
			"name":  name,
			"value": value,
		})
	})

	// //r.StaticFS("/static", http.Dir("static"))

	r.Run(":8080") //r.Run(":" + os.Getenv("PORT"))
}
