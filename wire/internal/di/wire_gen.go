// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"wire-demo/internal/application/service"
	"wire-demo/internal/infrastructure/db"
	"wire-demo/internal/infrastructure/repository"
	"wire-demo/internal/interfaces/handler"
)

// Injectors from wire.go:

func InitializeOrderHandler(dsn string) (*handler.OrderHandler, error) {
	sqlDB, err := db.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	orderRepository := repository.NewOrderRepository(sqlDB)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	return orderHandler, nil
}
