// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"server/internal/user"
)

// Injectors from wire.go:

func InitializeHandler(db user.DBTX) (*user.Handler, error) {
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	handler := user.NewHandler(service)
	return handler, nil
}
