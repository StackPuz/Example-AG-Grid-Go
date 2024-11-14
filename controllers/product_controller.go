package controllers

import (
	"app/config"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (con *ProductController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	var offset = (page - 1) * size
	order := c.DefaultQuery("order", "id")
	direction := c.DefaultQuery("direction", "asc")
	var products []models.Product
	query := config.DB.Model(&products)
	search := c.Query("search")
	if search != "" {
		search = "%" + search + "%"
		query.Where("name like ?", search)
	}
	var count int64
	query.Count(&count)
	query.Order(order + " " + direction).
		Offset(offset).
		Limit(size).
		Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products, "count": count})
}
