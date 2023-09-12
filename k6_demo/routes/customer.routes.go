package routes

import (
	"k6_demo/controllers"

	"github.com/gin-gonic/gin"
)

func CustRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/customer", controller.CreateCustomer)
}