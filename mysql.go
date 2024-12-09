package pass_common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"user"`
	Password string `json:"passwd"`
	Database string `json:"database"`
}

func GetMysqlConfig(config config.Config, opt ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(opt...).Scan(mysqlConfig)
	return mysqlConfig
}
