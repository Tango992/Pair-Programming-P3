package controller

import (
	"fmt"
	"net/http"
	"pair-programming/dto"
	"pair-programming/models"
	"pair-programming/repository"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	Repository repository.Transaction
}

func NewTransactionController(r repository.Transaction) TransactionController {
	return TransactionController{
		Repository: r,
	}
}

func (tc TransactionController) CreateNewTransaction(c echo.Context) error {
	var transactionDataTmp dto.Transaction
	if err := c.Bind(&transactionDataTmp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	
	if err := c.Validate(&transactionDataTmp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	
	transactionData := models.Transaction{
		Description: transactionDataTmp.Description,
		Amount: transactionDataTmp.Amount,
	}
	
	if err := tc.Repository.Post(&transactionData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Transaction created",
	})
}

func (tc TransactionController) GetAllTransactions(c echo.Context) error {
	transactions, err := tc.Repository.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Get all transactions",
		"data": transactions,
	})
}

func (tc TransactionController) DeleteTransactionById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "id cannot be empty"})
	}
	
	if err := tc.Repository.DeleteById(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Deleted at id %v", id),
	})
}


func (tc TransactionController) GetTransactionById(c echo.Context) error{
	id := c.Param("id")

	transactions, err := tc.Repository.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Transaction retrieved successfully",
		"transaction": transactions,
	})
}

func (tc TransactionController) UpdateTransactionById(c echo.Context)error{
	id := c.Param("id") 
	
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	
	if err := c.Validate(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	
	if err := tc.Repository.PutById(id, &transaction); err != nil {
		return c.JSON(http.StatusFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "trasaction update successfully",
		"transaction": transaction,
	})
}
