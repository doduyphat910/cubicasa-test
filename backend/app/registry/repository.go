package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/external/persistence/pgsql"
	"github.com/google/wire"
)

var (
	repositorySet = wire.NewSet(
		pgsql.NewHubRepository,
		wire.Bind(new(repository.HubRepository), new(pgsql.HubRepository)),
	)
)
