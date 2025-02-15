package controller

import (
	// "banking/client"
	"banking/dto"
	"context"

	"banking/model"
	"fmt"
	"net/http"

	"regexp"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
)

type Controller struct {
	bankingController model.BusinessBankingInterface
}

func NewController(bankingController model.BusinessBankingInterface) *Controller {
	return &Controller{bankingController: bankingController}
}

func (ctl *Controller) RegisterUser(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(context.Background(), "RegisterUser-Controller")
	defer trace.Finish()
	var reqRegist dto.UserRequest
	errBind := c.Bind(&reqRegist)
	fmt.Println("reqRegist-2", reqRegist)
	if errBind != nil {
		trace.LogKV("Error Bind", errBind)
		return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", errBind.Error()))
	}

	// model.BankingInterface.RegisterUser()
	reqRegistCl := dto.UserRequest{
		Nama: reqRegist.Nama,
		NIK:  reqRegist.NIK,
		NoHp: reqRegist.NoHp,
	}
	trace.LogKV("Request-RegisterUser-Controller", reqRegistCl)
	resRegist, err := ctl.bankingController.RegisterUser(ctx, reqRegistCl)
	if err != nil {
		trace.LogKV("Error", err)
		match, _ := regexp.MatchString("duplicate", err.Error())
		if match {
			return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", "NIK / No HP sudah pernah digunakan"))
		}
		return c.JSON(http.StatusInternalServerError, dto.ResponseFailed("Transaksi gagal", err.Error()))
	}
	trace.LogKV("Response-RegisterUser-Controller", resRegist)

	var userBank model.Saldo
	userBank.NoRekening = resRegist.NoRekening
	return c.JSON(http.StatusOK, dto.ResponseSuccesWithData("Rekening Berhasil Terbentuk", userBank))
}

func (ctl *Controller) SetorSaldo(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(context.Background(), "SetorSaldo-Controller")
	defer trace.Finish()
	var reqSaldo dto.SetorSaldo
	errBind := c.Bind(&reqSaldo)
	fmt.Println("reqSaldo-2", reqSaldo)
	if errBind != nil {
		trace.LogKV("Error Bind", errBind)
		return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", errBind.Error()))
	}

	// model.BankingInterface.RegisterUser()
	reqSaldoTabung := dto.SetorSaldo{
		NoRekening: reqSaldo.NoRekening,
		Nominal:    reqSaldo.Nominal,
	}
	trace.LogKV("Request-SetorSaldo-Controller", reqSaldoTabung)
	resSaldo, err := ctl.bankingController.SetorSaldo(ctx, reqSaldoTabung)
	if err != nil {
		trace.LogKV("Error", err)
		match, _ := regexp.MatchString("no rows", err.Error())
		if match {
			return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", "Norek tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, dto.ResponseFailed("Transaksi gagal", err.Error()))
	}
	trace.LogKV("Response-SetorSaldo-Controller", resSaldo)
	var userBank model.Saldo
	userBank.Saldo = resSaldo.Saldo
	userBank.NoRekening = resSaldo.NoRekening
	return c.JSON(http.StatusOK, dto.ResponseSuccesWithData("Transaksi berhasil", userBank))
}

func (ctl *Controller) TarikSaldo(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(context.Background(), "	TarikSaldo-Controller")
	defer trace.Finish()
	var reqSaldo dto.TarikSaldo
	errBind := c.Bind(&reqSaldo)
	fmt.Println("reqSaldo-2", reqSaldo)
	if errBind != nil {
		trace.LogKV("Error Bind", errBind)
		return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", errBind.Error()))
	}

	// model.BankingInterface.RegisterUser()
	reqTarik := dto.TarikSaldo{
		NoRekening: reqSaldo.NoRekening,
		Nominal:    reqSaldo.Nominal,
	}
	trace.LogKV("Request-TarikSaldo-Controller", reqTarik)
	resSaldo, err := ctl.bankingController.TarikSaldo(ctx, reqTarik)
	if err != nil {
		trace.LogKV("Error", err)
		match, _ := regexp.MatchString("no rows", err.Error())
		if match {
			return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", "Norek tidak ditemukan"))
		}
		return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", err.Error()))
	}
	trace.LogKV("Response-TarikSaldo-Controller", resSaldo)
	var userBank model.Saldo
	userBank.Saldo = resSaldo.Saldo
	userBank.NoRekening = resSaldo.NoRekening
	return c.JSON(http.StatusOK, dto.ResponseSuccesWithData("Transaksi berhasil", userBank))
}

func (ctl *Controller) CekSaldo(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(context.Background(), "	CekSaldo-Controller")
	fmt.Println("Masuk CEKSALDOO")
	defer trace.Finish()
	NoRekening := c.Param("no_rekening")
	fmt.Println("Masuk CEKSALDOO", NoRekening)

	// model.BankingInterface.RegisterUser()
	trace.LogKV("Request-CekSaldo-Controller", NoRekening)
	resSaldo, err := ctl.bankingController.CekSaldo(ctx, NoRekening)
	if err != nil {
		trace.LogKV("Error", err)
		match, _ := regexp.MatchString("no rows", err.Error())
		if match {
			return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", "Norek tidak ditemukan"))
		}
		return c.JSON(http.StatusBadRequest, dto.ResponseFailed("Transaksi gagal", err.Error()))
	}
	trace.LogKV("Response-CekSaldo-Controller", resSaldo)
	var userBank model.Saldo
	userBank.Saldo = resSaldo.Saldo
	userBank.NoRekening = resSaldo.NoRekening
	return c.JSON(http.StatusOK, dto.ResponseSuccesWithData("Transaksi berhasil", userBank))
}
