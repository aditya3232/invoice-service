package controllers

import (
	"invoice-service/common/response"
	"invoice-service/domain/dto"
	"invoice-service/services"
	"net/http"
	"strconv"

	errWrap "invoice-service/common/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InvoiceController struct {
	service services.IServiceRegistry
}

type IInvoiceController interface {
	FindByID(*gin.Context)
	Create(*gin.Context)
	FindAllWithoutPagination(*gin.Context)
}

func NewInvoiceController(service services.IServiceRegistry) IInvoiceController {
	return &InvoiceController{service: service}
}

func (c *InvoiceController) FindByID(ctx *gin.Context) {
	reqID, _ := strconv.Atoi(ctx.Param("id"))
	invoice, err := c.service.GetInvoice().FindByID(ctx.Request.Context(), reqID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: invoice,
		Gin:  ctx,
	})
}

func (c *InvoiceController) Create(ctx *gin.Context) {
	request := &dto.InvoiceRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	invoice, err := c.service.GetInvoice().Create(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: invoice,
		Gin:  ctx,
	})
}

func (c *InvoiceController) FindAllWithoutPagination(ctx *gin.Context) {
	var params dto.InvoiceRequestParam
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(params); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Err:     err,
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	invoices, err := c.service.GetInvoice().FindAllWithoutPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: invoices,
		Gin:  ctx,
	})
}
