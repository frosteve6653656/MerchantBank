package usecase

import (
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TransferUseCase interface {
	InsertTransfer(transfer *model.TransferModel, ctx *gin.Context) error
	GetTransferById(id string) (*model.TransferModel, error)
	GetAllTransfer() ([]*model.TransferModel, error)
}

type transferUseCaseImpl struct {
	transferRepo repository.TransferRepo
	userRepo      repository.UserRepo
	bankRepo     repository.BankRepo
}

func (transferUseCase *transferUseCaseImpl) GetTransferById(id string) (*model.TransferModel, error) {

	return transferUseCase.transferRepo.GetTransferById(id)
}

func (transferUseCase *transferUseCaseImpl) GetAllTransfer() ([]*model.TransferModel, error) {
	return transferUseCase.transferRepo.GetAllTransfer()
}

func (transferUseCase *transferUseCaseImpl) InsertTransfer(transfer *model.TransferModel, ctx *gin.Context) error {
	var wallet []float64
	senderExistData, err := transferUseCase.userRepo.GetUserById(transfer.Send_id)
	if err != nil {
		return fmt.Errorf("transferUsecaseImpl.InsertTransfer() 1: %w", err)
	}
	if senderExistData == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with id %v not found", transfer.Send_id),
		}
	}
	bankSenderData, err := transferUseCase.bankRepo.GetBankById(transfer.Send_id)
	if err != nil {
		return fmt.Errorf("transferUsecaseImpl.InsertTransfer() 1: %w", err)
	}
	if bankSenderData == nil {
		return &utils.AppError{
			ErrorCode:    2,
			ErrorMessage: fmt.Sprintf("Bank data with id %v not found", transfer.Send_id),
		}
	}
	if bankSenderData.Wallet < transfer.Nominal {
		return &utils.AppError{
			ErrorCode:    3,
			ErrorMessage: fmt.Sprintf("Bank wallet with id %v balance not enough", transfer.Send_id),
		}
	} else {
		wallet = append(wallet, bankSenderData.Wallet)
	}

	receiveExistData, err := transferUseCase.userRepo.GetUserById(transfer.Receive_id)
	if err != nil {
		return fmt.Errorf("transferUsecaseImpl.InsertTransfer() 1: %w", err)
	}
	if receiveExistData == nil {
		return &utils.AppError{
			ErrorCode:    4,
			ErrorMessage: fmt.Sprintf("User data with id %v not found", transfer.Receive_id),
		}
	}

	bankReceiverData, err := transferUseCase.bankRepo.GetBankById(transfer.Receive_id)
	if err != nil {
		return fmt.Errorf("transferUsecaseImpl.InsertTransfer() 1: %w", err)
	}
	if bankReceiverData == nil {
		return &utils.AppError{
			ErrorCode:    5,
			ErrorMessage: fmt.Sprintf("Bank data with id %v not found", transfer.Receive_id),
		}
	} else {
		wallet = append(wallet, bankReceiverData.Wallet)
	}
	// session := sessions.Default(ctx)
	// username := session.Get("Username")

	return transferUseCase.transferRepo.InsertTransfer(transfer, wallet)
}

func NewTransferUseCase(transferRepo repository.TransferRepo, userRepo repository.UserRepo, bankRepo repository.BankRepo) TransferUseCase {
	return &transferUseCaseImpl{
		transferRepo: transferRepo,
		userRepo:      userRepo,
		bankRepo:     bankRepo,
	}
}
