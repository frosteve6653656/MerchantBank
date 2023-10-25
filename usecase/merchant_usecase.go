package usecase

import (
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type MerchantUseCase interface {
	InsertMerchant(merc *model.MerchantRequestModel, ctx *gin.Context) error
	GetMerchantById(id string) (*model.MerchantModel, error)
	GetMerchantByName(name string) (*model.MerchantModel, error)
	GetAllMerchant() ([]*model.MerchantModel, error)
	EditMerchantById(merc *model.MerchantModel, ctx *gin.Context) error
}

type merchantUseCaseImpl struct {
	merchantRepo repository.MerchantRepo
	userRepo  repository.UserRepo
}

func (merchantRepo *merchantUseCaseImpl) GetMerchantById(id string) (*model.MerchantModel, error) {

	return merchantRepo.merchantRepo.GetMerchantById(id)
}

func (merchantRepo *merchantUseCaseImpl) GetAllMerchant() ([]*model.MerchantModel, error) {
	return merchantRepo.merchantRepo.GetAllMerchant()
}

func (merchantRepo *merchantUseCaseImpl) GetMerchantByName(name string) (*model.MerchantModel, error) {
	existData, err := merchantRepo.merchantRepo.GetMerchantByName(name)
	if err != nil {
		return nil, fmt.Errorf("merchantUsecaseImpl.GetMerchantByName(): %w", err)
	}

	return existData, nil
}

func (merchantRepo *merchantUseCaseImpl) InsertMerchant(merc *model.MerchantRequestModel, ctx *gin.Context) error {

	if merc.FullName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if merc.Username == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Username cannot be empty",
		}
	}
	if merc.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}
	if len(merc.Password) < 8 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password must be at least 8 characters",
		}
	}
	if merc.NoPhone == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Phone number cannot be empty",
		}
	}
	if len(merc.NoPhone) < 11 || len(merc.NoPhone) > 15 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid phone number",
		}
	}
	if merc.Email == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Email cannot be empty",
		}
	}
	if !utils.ValidateEmail(merc.Email) {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid email",
		}
	}
	if merc.Address == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Address cannot be empty",
		}
	}

	existDataUsr, err := merchantRepo.userRepo.GetUserByName(merc.Username)
	if err != nil {
		return fmt.Errorf("merchantUsecaseImpl.InsertMerchant() 1: %w", err)
	}
	if existDataUsr != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with username %v already exists", merc.Username),
		}
	}

	existDataCust, err := merchantRepo.merchantRepo.GetMerchantByName(merc.FullName)
	if err != nil {
		return fmt.Errorf("merchantUsecaseImpl.InsertMerchant() 2: %w", err)
	}
	if existDataCust != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Merchant data with name %v already exists", merc.FullName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(merc.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("merchantUsecaseImpl.GenerateFromPassword(): %w", err)
	}

	// session := sessions.Default(ctx)
	// username := session.Get("Username")
	merc.Password = string(passHash)
	merc.Role = "Merchant"
	merc.Active = true

	return merchantRepo.merchantRepo.InsertMerchant(merc)
}

func (merchantRepo *merchantUseCaseImpl) EditMerchantById(merc *model.MerchantModel, ctx *gin.Context) error {
	if merc.FullName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if merc.NoPhone == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Phone number cannot be empty",
		}
	}
	if len(merc.NoPhone) < 11 || len(merc.NoPhone) > 15 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid phone number",
		}
	}
	if merc.Email == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Email cannot be empty",
		}
	}
	if !utils.ValidateEmail(merc.Email) {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid email",
		}
	}
	if merc.Address == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Address cannot be empty",
		}
	}

	existDataCustId, err := merchantRepo.merchantRepo.GetMerchantById(merc.ID)
	if err != nil {
		return fmt.Errorf("merchantUsecaseImpl.UpdateMerchant() 2: %w", err)
	}
	if existDataCustId == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Merchant data with ID %v does not exist", merc.ID),
		}
	}

	existDataCust, err := merchantRepo.merchantRepo.GetMerchantByName(merc.FullName)
	if err != nil {
		return fmt.Errorf("merchantUsecaseImpl.UpdateMerchant() 3: %w", err)
	}
	if existDataCust != nil && merc.ID != existDataCust.ID {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Merchant data with name %v already exists", merc.FullName),
		}
	}
	merc.User_id = existDataCustId.User_id

	// session := sessions.Default(ctx)
	// username := session.Get("Username")

	return merchantRepo.merchantRepo.EditMerchantById(*merc)
}

func NewMerchantUseCase(merchantRepo repository.MerchantRepo, userRepo repository.UserRepo) MerchantUseCase {
	return &merchantUseCaseImpl{
		merchantRepo: merchantRepo,
		userRepo:  userRepo,
	}
}
