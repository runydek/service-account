package controllers

import (
	"net/http"
	"service-account/models"
	"service-account/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountController struct {
	DB *gorm.DB
}

func NewAccountController(db *gorm.DB) *AccountController {
	return &AccountController{DB: db}
}

func (ac *AccountController) RegisterAccount(c echo.Context) error {
	req := new(models.RegisterRequest)
	if err := c.Bind(req); err != nil {
		utils.LogError("Failed to bind request", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	existingAccount := new(models.Account)
	if err := ac.DB.Where("nik = ? OR no_hp = ?", req.NIK, req.NoHP).First(existingAccount).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "NIK or no_hp already exists"})
	}

	newAccount := models.Account{
		Nama:       req.Nama,
		NIK:        req.NIK,
		NoHP:       req.NoHP,
		NoRekening: utils.GenerateRekeningNumber(),
		Saldo:      0,
	}

	if err := ac.DB.Create(&newAccount).Error; err != nil {
		utils.LogWarning("Failed to create account", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Failed to create account"})
	}

	utils.LogInfo("Account created", map[string]interface{}{"no_rekening": newAccount.NoRekening})
	return c.JSON(http.StatusOK, map[string]interface{}{"no_rekening": newAccount.NoRekening})
}

func (ac *AccountController) Deposit(c echo.Context) error {
	req := new(models.TransactionRequest)
	if err := c.Bind(req); err != nil {
		utils.LogError("Failed to bind request", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	account := new(models.Account)
	if err := ac.DB.Where("no_rekening = ?", req.NoRekening).First(account).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Account not found"})
	}

	account.Saldo += req.Nominal
	ac.DB.Save(account)

	utils.LogInfo("Deposit successful", map[string]interface{}{"no_rekening": req.NoRekening, "saldo": account.Saldo})
	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": account.Saldo})
}

func (ac *AccountController) Withdraw(c echo.Context) error {
	req := new(models.TransactionRequest)
	if err := c.Bind(req); err != nil {
		utils.LogError("Failed to bind request", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	account := new(models.Account)
	if err := ac.DB.Where("no_rekening = ?", req.NoRekening).First(account).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Account not found"})
	}

	if account.Saldo < req.Nominal {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Insufficient balance"})
	}

	account.Saldo -= req.Nominal
	ac.DB.Save(account)

	utils.LogInfo("Withdrawal successful", map[string]interface{}{"no_rekening": req.NoRekening, "saldo": account.Saldo})
	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": account.Saldo})
}

func (ac *AccountController) GetBalance(c echo.Context) error {
	noRekening := c.Param("no_rekening")
	account := new(models.Account)

	if err := ac.DB.Where("no_rekening = ?", noRekening).First(account).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Account not found"})
	}

	utils.LogInfo("Balance fetched", map[string]interface{}{"no_rekening": noRekening, "saldo": account.Saldo})
	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": account.Saldo})
}
