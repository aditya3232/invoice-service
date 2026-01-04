package kafka

import (
	kafka "invoice-service/controllers/kafka/payment"
	"invoice-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IKafkaRegistry interface {
	GetPayment() kafka.IPaymentKafka
}

func NewKafkaRegistry(service services.IServiceRegistry) IKafkaRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetPayment() kafka.IPaymentKafka {
	return kafka.NewPaymentKafka(r.service)
}
