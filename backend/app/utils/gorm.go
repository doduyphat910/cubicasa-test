package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	onceDB      sync.Once
	singletonDB *gorm.DB
)

func OpenDBConnection(dsn string, config gorm.Config) (*gorm.DB, error) {
	var err error
	onceDB.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &config)
		if err == nil {
			singletonDB = db
		}
	})
	if err != nil {
		return nil, err
	}
	return singletonDB, nil
}

func GetDB() *gorm.DB {
	return singletonDB
}

const maxPageSize = 100

type Paging struct{ Size, Number uint64 }

func (paging *Paging) ParsePaging() {
	switch {
	case paging.Size == 0:
		paging.Size = maxPageSize
	case paging.Size > maxPageSize:
		paging.Size = maxPageSize
	}
	if paging.Number == 0 {
		paging.Number = 1
	}
}

func Paginate(paging Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		paging.ParsePaging()
		offset := (paging.Number - 1) * paging.Size
		return db.Limit(int(paging.Size)).Offset(int(offset))
	}
}
