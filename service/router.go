package service

import (
	"kit/net/router"
)

func initInner(r *router.Router) {
	r.VerifyGet("/x/internal/add", add)
	r.VerifyGet("/x/internal/list", list)
}

func initOutter(r *router.Router) {
	r.GuestGet("/x/outter/hello", hello)
}
