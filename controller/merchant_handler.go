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

type MerchantHandler struct {
	merchantUseCase usecase.MerchantUseCase
}

func (mercHandler MerchantHandler) GetAllMerchant(ctx *gin.Context) {
	merc, err := mercHandler.merchantUseCase.GetAllMerchant()
	if err != nil {
		fmt.Printf("MerchantHandler.GetAllMerchant(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching merchant data",
		})
		return
	}
	if merc == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    merc,
	})
}

func (mercHandler MerchantHandler) GetMerchantById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID cannot be empty",
		})
		return
	}

	merc, err := mercHandler.merchantUseCase.GetMerchantById(id)
	if err != nil {
		fmt.Printf("MerchantHandler.GetMerchantById(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching merchant data",
		})
		return
	}
	if merc == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    merc,
	})
}

func (mercHandler MerchantHandler) GetMerchantByName(ctx *gin.Context) {
	merchant := &model.MerchantModel{}
	err := ctx.ShouldBindJSON(&merchant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if merchant.FullName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}

	merc, err := mercHandler.merchantUseCase.GetMerchantByName(merchant.FullName)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.GetMerchantByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.GetMerchantByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching merchant data",
			})
			return
		}
		return
	}
	if merc == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    merc,
	})
}

func (mercHandler MerchantHandler) InsertMerchant(ctx *gin.Context) {
	merc := &model.MerchantRequestModel{}
	err := ctx.ShouldBindJSON(&merc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = mercHandler.merchantUseCase.InsertMerchant(merc, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.InsertMerchant() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.InsertMerchant() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving merchant data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (mercHandler MerchantHandler) EditMerchant(ctx *gin.Context) {
	merc := &model.MerchantModel{}
	err := ctx.ShouldBindJSON(&merc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = mercHandler.merchantUseCase.EditMerchantById(merc, ctx)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.EditMerchant() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.EditMerchant() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving merchant data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewMerchantHandler(srv *gin.Engine, merchantUseCase usecase.MerchantUseCase) *MerchantHandler {
	mercHandler := &MerchantHandler{
		merchantUseCase: merchantUseCase,
	}

	// route
	srv.POST("/merchant", middleware.RequireToken(), middleware.LevelUserAdmin(), mercHandler.InsertMerchant)
	srv.PUT("/merchant", middleware.RequireToken(), middleware.LevelUserAdmin(), mercHandler.EditMerchant)
	srv.GET("/merchant/:id", middleware.RequireToken(), mercHandler.GetMerchantById)
	srv.GET("/merchant", middleware.RequireToken(), middleware.LevelUserAdmin(), mercHandler.GetMerchantByName)
	srv.GET("/merchants", middleware.RequireToken(), middleware.LevelUserAdmin(), mercHandler.GetAllMerchant)

	return mercHandler
}
