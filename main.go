package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/hangim/store/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.StaticFS("/static/", http.Dir("static"))

	router.GET("/", ListBooksHandler)
	router.GET("/book/:id", ViewBookHandler)

	router.Run(":9000")
}
