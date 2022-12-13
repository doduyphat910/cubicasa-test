package aggregate

import "github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"

type UserAggregate struct {
	entity.User
	entity.Team
	entity.Hub
}
