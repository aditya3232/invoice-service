package routes

import (
	controllers "invoice-service/controllers/http"
	routes "invoice-service/routes/invoice"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouteRegister interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouteRegister {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) invoiceRoute() routes.IInvoiceRoute {
	return routes.NewInvoiceRoute(r.controller, r.group)
}

func (r *Registry) Serve() {
	r.invoiceRoute().Run()
}
