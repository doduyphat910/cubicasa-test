package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase"
	"github.com/google/wire"
)

var (
	HubUseCaseSet = wire.NewSet(
		repositorySet,
		usecase.NewHubUseCase,
		wire.Bind(new(usecase.HubUseCaser), new(usecase.HubUseCase)),
	)
)
