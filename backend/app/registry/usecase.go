package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase"
	"github.com/google/wire"
)

var (
	HubUseCaseSet = wire.NewSet(
		repositorySet,
		usecase.NewHubUseCase,
		wire.Bind(new(usecase.HubUseCaser), new(*usecase.HubUseCase)),
	)
	TeamUseCaseSet = wire.NewSet(
		repositorySet,
		usecase.NewTeamUseCase,
		wire.Bind(new(usecase.TeamUseCaser), new(*usecase.TeamUseCase)),
	)
	UserUseCaseSet = wire.NewSet(
		repositorySet,
		usecase.NewUserUseCase,
		wire.Bind(new(usecase.UserUseCaser), new(*usecase.UserUseCase)),
	)
)
