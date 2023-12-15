package config

import (
	"github.com/huuhoait/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/huuhoait/simple-admin-common/config"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	RedisConf    redis.RedisConf
}
