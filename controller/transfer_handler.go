package controller

import (
	"errors"
	"final-project/middleware"
	"final-project/model"
	"final-project/usecase"
	"final-project/utils"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type TransferHandler struct {
	bankUseCase usecase.TransferUseCase
}

func (bankHandler TransferHandler) GetAllTransfer(ctx *gin.Context) {
	bank, err := bankHandler.bankUseCase.GetAllTransfer()
	if err != nil {
		fmt.Printf("TransferHandler.GetAllTransfer(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching bank data",
		})
		return
	}
	if bank == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bank,
	})
}

func (bankHandler TransferHandler) GetTransferById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID cannot be empty",
		})
		return
	}

	bank, err := bankHandler.bankUseCase.GetTransferById(id)
	if err != nil {
		fmt.Printf("TransferHandler.GetTransferById(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching bank data",
		})
		return
	}
	if bank == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bank,
	})
}

func (bankHandler TransferHandler) InsertTransfer(ctx *gin.Context) {
	bank := &model.TransferModel{}
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = bankHandler.bankUseCase.InsertTransfer(bank, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("TransferHandler.InsertTransfer() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("TransferHandler.InsertTransfer() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving bank data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewTransferHandler(srv *gin.Engine, bankUseCase usecase.TransferUseCase) *TransferHandler {
	bankHandler := &TransferHandler{
		bankUseCase: bankUseCase,
	}

	// route
	srv.POST("/bank", middleware.RequireToken(), middleware.LevelUserAdmin(), bankHandler.InsertTransfer)
	srv.GET("/bank/:id", middleware.RequireToken(), bankHandler.GetTransferById)
	srv.GET("/banks", middleware.RequireToken(), middleware.LevelUserAdmin(), bankHandler.GetAllTransfer)

	return bankHandler
}
