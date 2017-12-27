package service

import (
	"context"
	"os"
	"reply/config"
	"reply/model"
	"testing"
)

var (
	testSvr  *service
	err      error
	ok       bool
	ctx      = context.Background()
	SOURCEID = int64(1)
	TYPEID   = model.NOTE_TYPE
)

func TestMain(m *testing.M) {
	testSvr, _ = NewService(config.Conf)
	os.Exit(m.Run())
}

// go test service.go init_test.go reply.go service_test.go -v
