package config

import "github.com/urfave/cli/v2"

var (
	EnvFlag = &cli.StringFlag{
		Name:    "env",
		Usage:   "Application environment: development, staging, production",
		EnvVars: []string{"ENV"},
		Value:   "dev",
	}
)

var (
	AppNameFlag = &cli.StringFlag{
		Name:    "app_name",
		Usage:   "Application name",
		EnvVars: []string{"APP_NAME"},
		Value:   "cubicasa_test",
	}
	AppPortFlag = &cli.StringFlag{
		Name:    "app_port",
		Usage:   "Application port",
		EnvVars: []string{"APP_PORT"},
		Value:   "8000",
	}
)

var (
	PGSQLHostFlag = &cli.StringFlag{
		Name:    "db_host",
		Usage:   "specify DB host",
		EnvVars: []string{"PGSQL_HOST"},
		Value:   "localhost",
	}
	PGSQLPortFlag = &cli.StringFlag{
		Name:    "db_port",
		Usage:   "DB port is using by application",
		EnvVars: []string{"PGSQL_PORT"},
		Value:   "5432",
	}
	PGSQLUsernameFlag = &cli.StringFlag{
		Name:    "db_user",
		Usage:   "DB username",
		EnvVars: []string{"PGSQL_USERNAME"},
		Value:   "cubicasa",
	}
	PGSQLPasswordFlag = &cli.StringFlag{
		Name:    "db_password",
		Usage:   "password used for DB user",
		EnvVars: []string{"PGSQL_PASSWORD"},
		Value:   "cubicasa",
	}
	PGSQLNameFlag = &cli.StringFlag{
		Name:    "db_name",
		Usage:   "DB name is using by application",
		EnvVars: []string{"PGSQL_DB"},
		Value:   "cubicasa_test",
	}
	PGSQLMaxOpenConnsFlag = &cli.IntFlag{
		Name:    "db_max_open_conns",
		Usage:   "sets the maximum number of open connections to the database",
		EnvVars: []string{"PGSQL_MAX_OPEN_CONNS"},
		Value:   10,
	}
	PGSQLMaxIdleConnsFlag = &cli.IntFlag{
		Name:    "db_max_idle_conns",
		Usage:   "sets the maximum number of connections in the idle connection pool",
		EnvVars: []string{"PGSQL_MAX_IDLE_CONNS"},
		Value:   5,
	}
	PGSQLConnMaxLifetimeFlag = &cli.IntFlag{
		Name:    "db_conn_max_lifetime",
		Usage:   "sets the maximum amount of time in minutes a connection may be reused",
		EnvVars: []string{"PGSQL_CONN_MAX_LIFETIME"},
		Value:   60,
	}
	PGSQLIsEnabledLogFlag = &cli.BoolFlag{
		Name:    "is_enabled_log",
		Usage:   "Is turn on DB log",
		EnvVars: []string{"PGSQL_IS_ENABLE_LOG"},
		Value:   true,
	}
)
