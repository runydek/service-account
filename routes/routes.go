package routes

import (
	"net/http"
	"service-account/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	accountController := controllers.NewAccountController(db)

	e.POST("/daftar", accountController.RegisterAccount)
	e.POST("/tabung", accountController.Deposit)
	e.POST("/tarik", accountController.Withdraw)
	e.GET("/saldo/:no_rekening", accountController.GetBalance)

	e.HTTPErrorHandler = customHTTPErrorHandler
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
}
