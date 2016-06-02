package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	. "github.com/hangim/store/models"
)

func ListBooksHandler(c *gin.Context) {
	var books []Book

	if err := DB.Find(&books).Error; err != nil {
		log.Panicln(err)
	}

	for i := range books {
		if err := DB.Model(&books[i]).Association("Authors").Find(&books[i].Authors).Error; err != nil {
			log.Panicln(err)
		}
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"books": books,
	})
}

func ViewBookHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Panicln(err)
	}

	var book = &Book{}
	if err := DB.Find(&book, id).Error; err != nil {
		log.Panicln(err)
	}

	if err := DB.Model(&book).Association("Authors").Find(&book.Authors).Error; err != nil {
		log.Panicln(err)
	}

	c.HTML(http.StatusOK, "book.tmpl", gin.H{
		"book": book,
	})
}
