package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DBConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func gormOpen() (gormDB *gorm.DB, err error) {
	var infoDb infoDB

	err = infoDb.getDriverConn("Databases.pgsql.auth")

	if err != nil {
		return nil, err
	}

	gormDB, err = gorm.Open(postgres.Open(infoDb.write.driverConn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return
	}

	dialector := postgres.New(postgres.Config{
		DSN: infoDb.read.driverConn,
	})

	err = gormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{dialector},
	}))

	if err != nil {
		return nil, err
	}

	var result int
	
	if err = gormDB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return nil, err
	}

	return
}
