package config

import (
	"sync"

	"github.com/urfave/cli/v2"
)

type App struct {
	Name string
	Port string
}
type PGSQL struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	IsEnabledLog    bool
}

type Config struct {
	Env       string
	App       App
	PGSQL     PGSQL
	BasicAuth BasicAuth
}
type BasicAuth struct {
	Username string
	Password string
}

var (
	configOnce      sync.Once
	singletonConfig *Config
)

func Init(ctx *cli.Context) *Config {
	conf := &Config{
		Env: ctx.String(EnvFlag.Name),
		App: App{
			Port: ctx.String(AppPortFlag.Name),
			Name: ctx.String(AppNameFlag.Name),
		},
		PGSQL: PGSQL{
			Host:            ctx.String(PGSQLHostFlag.Name),
			Port:            ctx.String(PGSQLPortFlag.Name),
			Username:        ctx.String(PGSQLUsernameFlag.Name),
			Password:        ctx.String(PGSQLPasswordFlag.Name),
			Name:            ctx.String(PGSQLNameFlag.Name),
			MaxOpenConns:    ctx.Int(PGSQLMaxOpenConnsFlag.Name),
			MaxIdleConns:    ctx.Int(PGSQLMaxIdleConnsFlag.Name),
			ConnMaxLifetime: ctx.Int(PGSQLConnMaxLifetimeFlag.Name),
			IsEnabledLog:    ctx.Bool(PGSQLIsEnabledLogFlag.Name),
		},
		BasicAuth: BasicAuth{
			Username: ctx.String(BasicAuthUsernameFlag.Name),
			Password: ctx.String(BasicAuthPasswordFlag.Name),
		},
	}

	configOnce.Do(func() {
		singletonConfig = conf
	})

	return singletonConfig
}

func GetConfig() *Config {
	return singletonConfig
}
