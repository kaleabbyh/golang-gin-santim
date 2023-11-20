package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/models"
)


func CreateDemo(c *gin.Context) {

    var demoData models.Demo
    if err := c.ShouldBindJSON(&demoData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := db.Create(&demoData)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusCreated, demoData)
}