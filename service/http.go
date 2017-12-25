package service

import (
	"kit/net/context"
)

func hello(c context.Context) {
	res := c.Result()
	r, err := svr.Get(c, int64(1))
	res["code"] = err
	res["data"] = r
	return
}
