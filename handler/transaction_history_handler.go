package handler

import (
	"finalproject4/helper"
	"finalproject4/model"
	"finalproject4/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHistoryHandler interface {
	GetAllTransactionHistory(ctx *gin.Context)
	GetTransactionHistoryByUserId(ctx *gin.Context)
	CreateTransactionHistory(ctx *gin.Context)
	DeleteTransactionHistory(ctx *gin.Context)
}

type transactionHistoryHandler struct {
	transactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionHistoryService service.TransactionHistoryService) *transactionHistoryHandler {
	return &transactionHistoryHandler{transactionHistoryService}
}

func (th *transactionHistoryHandler) GetAllTransactionHistory(ctx *gin.Context) {

	role := ctx.MustGet("role").(string)
	id := ctx.MustGet("userID").(int)
	transactionHistory, err := th.transactionHistoryService.GetAllTransactionHistory(role, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatGetAdminTransaction(transactionHistory)
	ctx.JSON(http.StatusOK, formatter)
}

func (th *transactionHistoryHandler) GetTransactionHistoryByUserId(ctx *gin.Context) {

	role := ctx.MustGet("role").(string)
	id := ctx.MustGet("userID").(int)
	transactionHistory, err := th.transactionHistoryService.GetAllTransactionHistory(role, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatGetUserTransaction(transactionHistory)
	ctx.JSON(http.StatusOK, formatter)

}

func (th *transactionHistoryHandler) CreateTransactionHistory(ctx *gin.Context) {
	var transactions model.TransactionHistoryInput
	id, _ := helper.GetUserID(ctx)
	err := ctx.ShouldBindJSON(&transactions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	transactionData, err := th.transactionHistoryService.CreateTransactionHistory(transactions, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatTransaction(transactionData)
	ctx.JSON(http.StatusCreated, formatter)
}

func (h *transactionHistoryHandler) DeleteTransactionHistory(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	currentUserID, _ := helper.GetUserID(ctx)
	err = h.transactionHistoryService.DeleteTransactionHistory(id, currentUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	format := model.FormatDeleteTransaction()
	ctx.JSON(http.StatusOK, format)
}
