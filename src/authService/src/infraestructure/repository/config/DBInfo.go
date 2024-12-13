package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type infoDB struct {
	read struct {
		hostname   string
		name       string
		username   string
		password   string
		port       string
		parameter  string
		driverConn string
	}

	write struct {
		hostname   string
		name       string
		username   string
		password   string
		port       string
		parameter  string
		driverConn string
	}
}

func (infoDb *infoDB) getDriverConn(nameMap string) (err error) {
	viper.SetConfigFile("config.json")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDb)

	if err != nil {
		return
	}

	infoDb.read.driverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", infoDb.read.username, infoDb.read.password, infoDb.read.hostname, infoDb.read.port, infoDb.read.name)

	infoDb.write.driverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", infoDb.read.username, infoDb.read.password, infoDb.read.hostname, infoDb.read.port, infoDb.read.name)
	return
}
