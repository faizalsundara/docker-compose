package routes

import (
	"banking/app"

	"github.com/labstack/echo"
)

func New(presenter app.Presenter) *echo.Echo {
	e := echo.New()

	e.POST("/daftar", presenter.UserPresenter.RegisterUser)
	e.POST("/tabung", presenter.UserPresenter.SetorSaldo)
	e.POST("/tarik", presenter.UserPresenter.TarikSaldo)
	e.GET("/saldo/:no_rekening", presenter.UserPresenter.CekSaldo)

	return e
}
