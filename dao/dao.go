package dao

import (
	"kit/db/memcache"
	"kit/db/mysql"
	"kit/db/redis"
	"kit/log"
	"reply/config"
	"time"
)

type Dao struct {
	conf        *config.Config
	db          *mysql.DB
	mc          *memcache.Pool
	redis       *redis.Pool
	expireRedis int
	expireMc    int
}

func NewDao(c *config.Config) (d *Dao) {
	d = &Dao{
		conf: c,
	}
	if db, err := mysql.NewMysql(c.Mysql); err != nil {
		log.Error("mysql.NewMysql error(%v)", err)
		return
	} else {
		d.db = db
	}
	d.redis = redis.NewRedisPool(c.Redis)
	d.mc = memcache.NewMemcachePool(c.Memcache)

	// expire
	d.expireRedis = int(time.Duration(c.Reply.ExpireRedis) / time.Second)
	d.expireMc = int(time.Duration(c.Reply.ExpireMc) / time.Second)
	return
}
