package usecase

import (
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	InsertCustomer(cust *model.CustomerRequestModel, ctx *gin.Context) error
	GetCustomerById(id string) (*model.CustomerModel, error)
	GetCustomerByName(name string) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	EditCustomerById(cust *model.CustomerModel, ctx *gin.Context) error
}

type customerUseCaseImpl struct {
	custRepo repository.CustomerRepo
	userRepo  repository.UserRepo
}

func (custUseCase *customerUseCaseImpl) GetCustomerById(id string) (*model.CustomerModel, error) {

	return custUseCase.custRepo.GetCustomerById(id)
}

func (custUseCase *customerUseCaseImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	return custUseCase.custRepo.GetAllCustomer()
}

func (custUseCase *customerUseCaseImpl) GetCustomerByName(name string) (*model.CustomerModel, error) {
	existData, err := custUseCase.custRepo.GetCustomerByName(name)
	if err != nil {
		return nil, fmt.Errorf("customerUsecaseImpl.GetCustomerByName(): %w", err)
	}

	return existData, nil
}

func (custUseCase *customerUseCaseImpl) InsertCustomer(cust *model.CustomerRequestModel, ctx *gin.Context) error {

	if cust.FullName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if cust.Username == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Username cannot be empty",
		}
	}
	if cust.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}
	if len(cust.Password) < 8 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password must be at least 8 characters",
		}
	}
	if cust.NoPhone == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Phone number cannot be empty",
		}
	}
	if len(cust.NoPhone) < 11 || len(cust.NoPhone) > 15 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid phone number",
		}
	}
	if cust.Email == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Email cannot be empty",
		}
	}
	if !utils.ValidateEmail(cust.Email) {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid email",
		}
	}
	if cust.Address == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Address cannot be empty",
		}
	}

	existDataUsr, err := custUseCase.userRepo.GetUserByName(cust.Username)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.InsertCustomer() 1: %w", err)
	}
	if existDataUsr != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with username %v already exists", cust.Username),
		}
	}

	existDataCust, err := custUseCase.custRepo.GetCustomerByName(cust.FullName)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.InsertCustomer() 2: %w", err)
	}
	if existDataCust != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Customer data with name %v already exists", cust.FullName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(cust.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.GenerateFromPassword(): %w", err)
	}

	// session := sessions.Default(ctx)
	// username := session.Get("Username")
	cust.Password = string(passHash)
	cust.Role = "Customer"
	cust.Active = true

	return custUseCase.custRepo.InsertCustomer(cust)
}

func (custUseCase *customerUseCaseImpl) EditCustomerById(cust *model.CustomerModel, ctx *gin.Context) error {
	if cust.FullName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if cust.NoPhone == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Phone number cannot be empty",
		}
	}
	if len(cust.NoPhone) < 11 || len(cust.NoPhone) > 15 {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid phone number",
		}
	}
	if cust.Email == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Email cannot be empty",
		}
	}
	if !utils.ValidateEmail(cust.Email) {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Invalid email",
		}
	}
	if cust.Address == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Address cannot be empty",
		}
	}

	existDataCustId, err := custUseCase.custRepo.GetCustomerById(cust.ID)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.UpdateCustomer() 2: %w", err)
	}
	if existDataCustId == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Customer data with ID %v does not exist", cust.ID),
		}
	}

	existDataCust, err := custUseCase.custRepo.GetCustomerByName(cust.FullName)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.UpdateCustomer() 3: %w", err)
	}
	if existDataCust != nil && cust.ID != existDataCust.ID {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Customer data with name %v already exists", cust.FullName),
		}
	}
	cust.User_id = existDataCustId.User_id

	// session := sessions.Default(ctx)
	// username := session.Get("Username")

	return custUseCase.custRepo.EditCustomerById(*cust)
}

func NewCustomerUseCase(custRepo repository.CustomerRepo, userRepo repository.UserRepo) CustomerUseCase {
	return &customerUseCaseImpl{
		custRepo: custRepo,
		userRepo:  userRepo,
	}
}
