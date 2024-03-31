package dal

import (
	"sync"

	log "go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"kbds-ref-restapi/models"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Connect(l *log.Logger) (*gorm.DB, error) {
	var dbErr error
	//var db *gorm.DB

	once.Do(func() {
		db, dbErr = gorm.Open(
			mysql.Open(models.GlobalConfig.ConnStr),
			&gorm.Config{Logger: *LogConfig(l, logger.Info)},
		)
		if dbErr != nil {
			return
		}
		sqldb, err := db.DB()
		if err != nil {
			dbErr = err
			return
		}
		sqldb.SetMaxIdleConns(6)
		sqldb.SetMaxOpenConns(100)
	})

	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}

func StartTran(l *log.Logger, db *gorm.DB) *gorm.DB {
	//opt := new(sql.TxOptions)
	//opt.Isolation = sql.LevelSerializable
	if db == nil {
		newDb, err := Connect(l)
		if err != nil {
			panic("Unable to establish connection to database.")
		}
		return newDb.Begin()
	}
	return db.Begin()
}

func CommitTran(db *gorm.DB) {
	db.Commit()
}

func Rollback(db *gorm.DB) {
	db.Rollback()
}
