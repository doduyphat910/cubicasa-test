package app

import (
	"fmt"
	"github.com/doduyphat910/cubicasa-test/backend/app/config"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
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

	db, err := utils.OpenDBConnection(dsn, gorm.Config{Logger: logMode})
	if err != nil {
		log.Fatalf("can't connect to database, err:%s", err)
		return nil
	}
	gormDB, err := db.DB()
	if err != nil {
		log.Fatalf("get gormDB, err:%s", err)
		return nil
	}
	gormDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	gormDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	gormDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnMaxLifetime) * time.Minute)

	return db
}
