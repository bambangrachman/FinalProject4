package handler

import (
	"finalproject4/config"
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
	CreateTest(ctx *gin.Context)
	DeleteTransactionHistory(ctx *gin.Context)
}

type transactionHistoryHandler struct {
	transactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionHistoryService service.TransactionHistoryService) *transactionHistoryHandler {
	return &transactionHistoryHandler{transactionHistoryService}
}

func (th *transactionHistoryHandler) GetAllTransactionHistory(ctx *gin.Context) {

	role := ctx.MustGet("currentUserRole").(string)
	currentUser := ctx.MustGet("currentUser").(model.User)
	if role == "customer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "You are not authorized to access this endpoint",
		})
		return
	}

	id := int(currentUser.ID)
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

	role := ctx.MustGet("currentUserRole").(string)
	currentUser := ctx.MustGet("currentUser").(model.User)
	id := int(currentUser.ID)
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

func (th *transactionHistoryHandler) CreateTest(ctx *gin.Context) {
	var transaction model.TransactionHistoryInput

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid type of product and quantity"})
		return
	}

	currentUser := ctx.MustGet("currentUser").(model.User)
	userID := int(currentUser.ID)

	transactionData, err := th.transactionHistoryService.CreateTransactionTest(transaction, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	formatter := model.FormatTransaction(transactionData)
	ctx.JSON(http.StatusOK, formatter)

}

func (th *transactionHistoryHandler) CreateTransactionHistory(ctx *gin.Context) {
	db := config.GetDB()
	println("tes0")
	transaction := model.TransactionHistory{}

	println("tes1")
	currentUser := ctx.MustGet("currentUser").(model.User)

	id := int(currentUser.ID)
	println("tes2")
	// tes nya berhenti disini apakah shouldbindjsonnya ada yg salah?
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	//Cek Apakah product ada dan stok ada
	Product := model.Product{}
	err = db.First(&Product, transaction.ProductID).Error
	println("tes3")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	if transaction.Quantity > Product.Stock {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Product stock is not enough",
		})
		return
	}

	transaction.TotalPrice = transaction.Quantity * Product.Price
	println("tes4")
	User := model.User{}
	err = db.First(&User, id).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	if transaction.TotalPrice > int(User.Balance) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Saldo is not enough",
		})
		return
	}
	println("tes5")
	//Pengurangan stok product
	db.Model(&Product).Where("id = ?", Product.ID).Update("stock", Product.Stock-transaction.Quantity)

	//Pengurangan saldo User
	db.Model(&User).Where("id = ?", User.ID).Update("balance", User.Balance-int(transaction.TotalPrice))
	println("tes6")
	//Create Transaction
	transaction.UserID = id

	err = db.Debug().Create(&transaction).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	println("tes7")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Your have succesfully purchased the product",
		"transaction_bill": `{
			"total_price":Transaction.TotalPrice,
			"quantity":Transaction.Quantity,
			"Product_title":Product.Title,
		}`,
	})
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
	currentUser := ctx.MustGet("currentUser").(model.User)
	currentUserID := int(currentUser.ID)
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
