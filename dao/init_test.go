package dao

import (
	"context"
	"os"
	"reply/config"
	"reply/model"
	"testing"
)

var (
	d        *Dao
	err      error
	ok       bool
	ctx      = context.Background()
	SOURCEID = int64(1)
	TYPEID   = model.NOTE_TYPE
)

func TestMain(m *testing.M) {
	d = NewDao(config.Conf)
	os.Exit(m.Run())
}

// test db     <===>   go test dao.go reply_mysql.go init_test.go reply_mysql_test.go -v
// test mc     <===>   go test dao.go reply_mc.go init_test.go reply_mc_test.go -v
// test redis  <===>   go test dao.go reply_redis.go init_test.go reply_redis_test.go -v
