package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/sqlite"

// type Item struct {
// 	Name  string
// 	Value int
// }

// type UserSession struct {
// 	Items []Item
// }

// var db *gorm.DB

// type Page struct {
// 	gorm.Model

// 	Title string `gorm:"column:title"`
// 	Body  []byte `gorm:"column:body"`
// }

func main() {
	// tempDB, _ := gorm.Open("sqlite3", "book.db")
	// db = tempDB
	// defer db.Close()

	// db.AutoMigrate(&Page{}) // if there's no table in the db, it will create one

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// r.GET("/:title", func(c *gin.Context) {
	// 	title := c.Param("title")

	// 	var p Page

	// 	if err := db.Where("title = ?", title).First(&p).Error; err != nil {
	// 		c.AbortWithStatus(404)
	// 		return
	// 	}

	// 	c.HTML(200, "view.html", gin.H{
	// 		"page": p,
	// 	})
	// })

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
		price := c.PostForm("price")

		c.HTML(200, "item.tmpl", gin.H{
			"name":  name,
			"price": price,
		})
	})

	// //r.StaticFS("/static", http.Dir("static"))

	r.Run(":8080") //r.Run(":" + os.Getenv("PORT"))
}
