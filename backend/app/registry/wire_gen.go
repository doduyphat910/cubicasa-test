// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/external/persistence/pgsql"
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase"
)

// Injectors from wire.go:

func InjectedHubUseCase() usecase.HubUseCaser {
	hubRepository := pgsql.NewHubRepository()
	hubUseCase := usecase.NewHubUseCase(hubRepository)
	return hubUseCase
}