package routes

import (
	"pair-programming/controller"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Echo, controller controller.TransactionController) {
	e.POST("/transactions", controller.CreateNewTransaction)
	e.GET("/transactions", controller.GetAllTransactions)
	e.GET("/transactions/:id", controller.GetTransactionById)
	e.PUT("/transactions/:id", controller.UpdateTransactionById)
	e.DELETE("/transactions/:id", controller.DeleteTransactionById)
}