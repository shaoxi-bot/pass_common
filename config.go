package pass_common

import (
	"strconv"

	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心的地址
		consul.WithAddress(host + strconv.FormatInt(port, 10)),
		// 设置前缀
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	err = conf.Load(consulSource)
	if err != nil {
		return nil, err
	}
	return conf, err
}
