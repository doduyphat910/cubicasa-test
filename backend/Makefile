mod:
	@go mod tidy
	@go mod vendor

MIG_ENV ?= development
migrate-up:
	cd db;\
	sql-migrate up -env=${MIG_ENV} -config=dbconfig.yml
migrate-down:
	cd db;\
	sql-migrate down -env=${MIG_ENV} -config=dbconfig.yml
migration:
	cd db;\
	sql-migrate new $(fileName)