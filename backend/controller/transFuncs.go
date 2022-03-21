package controller

import (
	"dhlabs/backend/models"
	"dhlabs/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransController struct {
	TransactService services.TransactService
}

func TransControllerConstruct(transService services.TransactService) TransController {
	return TransController{
		TransactService: transService,
	}
}

func (tc *TransController) NewTrans(c *gin.Context) {
	var transaction models.Transact
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := tc.TransactService.NewTrans(&transaction); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TransController) GetTrans(c *gin.Context) {
	transacs, err := tc.TransactService.GetTrans()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transacs)
}

func (tc *TransController) Routes(rg *gin.RouterGroup) {
	routes := rg.Group("/transactions")
	routes.GET("/getall", tc.GetTrans)
	routes.POST("/add", tc.NewTrans)
}
