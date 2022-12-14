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

func InjectedTeamUseCase() usecase.TeamUseCaser {
	teamRepository := pgsql.NewTeamRepository()
	hubRepository := pgsql.NewHubRepository()
	teamUseCase := usecase.NewTeamUseCase(teamRepository, hubRepository)
	return teamUseCase
}

func InjectedUserUseCase() usecase.UserUseCaser {
	teamRepository := pgsql.NewTeamRepository()
	userRepository := pgsql.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(teamRepository, userRepository)
	return userUseCase
}
