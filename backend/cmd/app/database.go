package app

import (
	"fmt"
	"github.com/doduyphat910/cubicasa-test/backend/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var (
	onceDB      sync.Once
	singletonDB *gorm.DB
)

func initDBConnection(dbCfg config.PGSQL) *gorm.DB {
	logMode := logger.Default.LogMode(logger.Error)
	if dbCfg.IsEnabledLog {
		logMode = logger.Default.LogMode(logger.Info)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.Name, dbCfg.Port,
	)

	var err error
	onceDB.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logMode,
		})
		if err == nil {
			singletonDB = db
		}
	})
	if err != nil {
		log.Fatalf("can't connect to database, err:%s", err)
		return nil
	}

	gormDB, err := singletonDB.DB()
	if err != nil {
		log.Fatalf("get gormDB, err:%s", err)
		return nil
	}
	gormDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	gormDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	gormDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnMaxLifetime) * time.Minute)

	return singletonDB
}
