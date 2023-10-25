package manager

import (
	"final-project/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetCustomerUsecase() usecase.CustomerUseCase
	GetMerchantUsecase() usecase.MerchantUseCase
	GetBankUsecase() usecase.BankUseCase
	GetTransferUsecase() usecase.TransferUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase  usecase.UserUseCase
	lgUsecase   usecase.LoginUseCase
	custUsecase usecase.CustomerUseCase
	mercUsecase usecase.MerchantUseCase
	bankUsecase usecase.BankUseCase
	trxusecase  usecase.TransferUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadMerchantUsecase sync.Once
var onceLoadBankUsecase sync.Once
var onceLoadTransferUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.lgUsecase = usecase.NewLoginUseCase(um.repoManager.GetLoginRepo(), um.repoManager.GetCustomerRepo())
	})
	return um.lgUsecase
}

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUseCase {
	onceLoadCustomerUsecase.Do(func() {
		um.custUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo(), um.repoManager.GetUserRepo())
	})
	return um.custUsecase
}

func (um *usecaseManager) GetMerchantUsecase() usecase.MerchantUseCase {
	onceLoadMerchantUsecase.Do(func() {
		um.mercUsecase = usecase.NewMerchantUseCase(um.repoManager.GetMerchantRepo(), um.repoManager.GetUserRepo())
	})
	return um.mercUsecase
}

func (um *usecaseManager) GetBankUsecase() usecase.BankUseCase {
	onceLoadBankUsecase.Do(func() {
		um.bankUsecase = usecase.NewBankUseCase(um.repoManager.GetBankRepo(), um.repoManager.GetUserRepo())
	})
	return um.bankUsecase
}

func (um *usecaseManager) GetTransferUsecase() usecase.TransferUseCase {
	onceLoadTransferUsecase.Do(func() {
		um.trxusecase = usecase.NewTransferUseCase(um.repoManager.GetTransferRepo(), um.repoManager.GetUserRepo(), um.repoManager.GetBankRepo())
	})
	return um.trxusecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
