package usecase

import (
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BankUseCase interface {
	InsertBank(bank *model.BankModel, ctx *gin.Context) error
	GetBankById(id string) (*model.BankModel, error)
	GetAllBank() ([]*model.BankModel, error)
	EditBankById(bank *model.BankModel, ctx *gin.Context) error
}

type bankUseCaseImpl struct {
	bankRepo repository.BankRepo
	userRepo  repository.UserRepo
}

func (bankUseCase *bankUseCaseImpl) GetBankById(id string) (*model.BankModel, error) {

	return bankUseCase.bankRepo.GetBankById(id)
}

func (bankUseCase *bankUseCaseImpl) GetAllBank() ([]*model.BankModel, error) {
	return bankUseCase.bankRepo.GetAllBank()
}

func (bankUseCase *bankUseCaseImpl) InsertBank(bank *model.BankModel, ctx *gin.Context) error {

	existDataUsr, err := bankUseCase.userRepo.GetUserById(bank.User_id)
	if err != nil {
		return fmt.Errorf("bankUsecaseImpl.InsertBank() 1: %w", err)
	}
	if existDataUsr != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with id %v already exists", bank.User_id),
		}
	}

	// session := sessions.Default(ctx)
	// username := session.Get("Username")

	return bankUseCase.bankRepo.InsertBank(bank)
}

func (bankUseCase *bankUseCaseImpl) EditBankById(bank *model.BankModel, ctx *gin.Context) error {
	existDataCustId, err := bankUseCase.bankRepo.GetBankByUserId(bank.ID)
	if err != nil {
		return fmt.Errorf("bankUsecaseImpl.UpdateBank() 2: %w", err)
	}
	if existDataCustId == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Bank data with ID %v does not exist", bank.ID),
		}
	}
	bank.User_id = existDataCustId.User_id

	// session := sessions.Default(ctx)
	// username := session.Get("Username")

	return bankUseCase.bankRepo.EditBankById(*bank)
}

func NewBankUseCase(bankRepo repository.BankRepo, userRepo repository.UserRepo) BankUseCase {
	return &bankUseCaseImpl{
		bankRepo: bankRepo,
		userRepo:  userRepo,
	}
}
