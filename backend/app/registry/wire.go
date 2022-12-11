//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase"
	"github.com/google/wire"
)

func InjectedHubUseCase() usecase.HubUseCaser {
	wire.Build(HubUseCaseSet)
	return nil
}

func InjectedTeamUseCase() usecase.TeamUseCaser {
	wire.Build(TeamUseCaseSet)
	return nil
}

func InjectedUserUseCase() usecase.UserUseCaser {
	wire.Build(UserUseCaseSet)
	return nil
}
