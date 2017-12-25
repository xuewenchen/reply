package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"kit/config"
	"kit/log"
	"kit/time"
)

// define your config
type Config struct {
	Common   *config.Common
	Log      *config.Log
	Mysql    *config.Mysql
	Redis    *config.Redis
	Memcache *config.Memcache

	Grpc  *config.Grpc
	Mhttp *config.Mhttp

	// trace
	Trace *config.Trace
	Reply *Reply
}

type Reply struct {
	ExpireRedis time.Duration
	ExpireMc    time.Duration
}

var (
	Conf     = &Config{}
	ConfPath = flag.String("conf", "./reply.toml", "config path")
)

func init() {
	// get config Path
	if _, err := toml.DecodeFile(*ConfPath, &Conf); err != nil {
		log.Error("toml.DecodeFile error(%v)", err)
		panic(err)
	}
}
