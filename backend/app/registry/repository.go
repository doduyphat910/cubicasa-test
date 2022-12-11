package registry

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/external/persistence/pgsql"
	"github.com/google/wire"
)

var (
	repositorySet = wire.NewSet(
		pgsql.NewHubRepository,
		pgsql.NewTeamRepository,
		pgsql.NewUserRepository,
		wire.Bind(new(repository.HubRepository), new(*pgsql.HubRepository)),
		wire.Bind(new(repository.TeamRepository), new(*pgsql.TeamRepository)),
		wire.Bind(new(repository.UserRepository), new(*pgsql.UserRepository)),
	)
)
