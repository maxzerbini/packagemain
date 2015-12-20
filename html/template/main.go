package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	Id         int
	Name       string
	Note       string
	Category   int
	UpdateDate time.Time
	Style      Style
}

type Style struct {
	Color  string
	Font   string
	Bold   bool
	Border bool
}

func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("mytemplate.html"))
	router.SetHTMLTemplate(html)
	router.GET("/index", RenderTemplate)
	router.Run(":8080")
}

func RenderTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "mytemplate.html", gin.H{
		"Title":    "Product Catalog",
		"Products": getProductList(),
	})
}

func getProductList() *[]Product {
	list := make([]Product, 8, 8)
	for i := 0; i < 8; i++ {
		list[i] = Product{Id: i, Name: "Product N." + strconv.Itoa(i), Note: "new product", Category: i*10 + 100, UpdateDate: time.Now()}
		if i%2 == 0 {
			list[i].Style.Color = "green"
			list[i].Style.Font = "Arial"
			list[i].Style.Bold = true
			list[i].Style.Border = true
		} else {
			list[i].Style.Color = "red"
			list[i].Style.Font = "Arial"
			list[i].Style.Bold = true
			list[i].Style.Border = true
		}
	}
	return &list
}
