package routes

import (
	controllers "invoice-service/controllers/http"

	"github.com/gin-gonic/gin"
)

type InvoiceRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IInvoiceRoute interface {
	Run()
}

func NewInvoiceRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup) IInvoiceRoute {
	return &InvoiceRoute{controller: controller, group: group}
}

func (r *InvoiceRoute) Run() {
	group := r.group.Group("/invoices")
	group.GET("", r.controller.GetInvoice().FindAllWithoutPagination)
	group.GET("/:id", r.controller.GetInvoice().FindByID)
	group.POST("", r.controller.GetInvoice().Create)
	group.POST("/:id/mark-overdue", r.controller.GetInvoice().MarkOverdue)
}
