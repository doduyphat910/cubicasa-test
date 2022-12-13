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

gen:
	@#go generate ./...
	@swag init -g app/external/framework/routes/route.go -o app/interface/api/docs --exclude db,scripts,vendor

