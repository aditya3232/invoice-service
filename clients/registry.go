package clients

import (
	"invoice-service/clients/config"
	clients "invoice-service/clients/customer"
	config2 "invoice-service/config"
)

type ClientRegistry struct{}

type IClientRegistry interface {
	GetCustomer() clients.ICustomerClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientRegistry{}
}

func (c *ClientRegistry) GetCustomer() clients.ICustomerClient {
	return clients.NewCustomerClient(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.Customer.Host),
		))
}
