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

type BankHandler struct {
	bankUseCase usecase.BankUseCase
}

func (bankHandler BankHandler) GetAllBank(ctx *gin.Context) {
	bank, err := bankHandler.bankUseCase.GetAllBank()
	if err != nil {
		fmt.Printf("BankHandler.GetAllBank(): %v", err.Error())
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

func (bankHandler BankHandler) GetBankById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID cannot be empty",
		})
		return
	}

	bank, err := bankHandler.bankUseCase.GetBankById(id)
	if err != nil {
		fmt.Printf("BankHandler.GetBankById(): %v", err.Error())
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

func (bankHandler BankHandler) InsertBank(ctx *gin.Context) {
	bank := &model.BankModel{}
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = bankHandler.bankUseCase.InsertBank(bank, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("BankHandler.InsertBank() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("BankHandler.InsertBank() 2: %v", err.Error())
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

func (bankHandler BankHandler) EditBank(ctx *gin.Context) {
	bank := &model.BankModel{}
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = bankHandler.bankUseCase.EditBankById(bank, ctx)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("BankHandler.EditBank() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("BankHandler.EditBank() 2: %v", err.Error())
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

func NewBankHandler(srv *gin.Engine, bankUseCase usecase.BankUseCase) *BankHandler {
	bankHandler := &BankHandler{
		bankUseCase: bankUseCase,
	}

	// route
	srv.POST("/bank", middleware.RequireToken(), middleware.LevelUserAdmin(), bankHandler.InsertBank)
	srv.PUT("/bank", middleware.RequireToken(), middleware.LevelUserAdmin(), bankHandler.EditBank)
	srv.GET("/bank/:id", middleware.RequireToken(), bankHandler.GetBankById)
	srv.GET("/banks", middleware.RequireToken(), middleware.LevelUserAdmin(), bankHandler.GetAllBank)

	return bankHandler
}
