package config

import (
	"github.com/huuhoait/zero-tools/config"
	"github.com/huuhoait/zero-tools/plugins/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	JobRpc       zrpc.RpcClientConf
	Captcha      Captcha
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
}

type Captcha struct {
	KeyLong   int // captcha length
	ImgWidth  int // captcha width
	ImgHeight int // captcha height
}
