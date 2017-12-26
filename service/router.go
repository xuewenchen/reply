package service

import (
	"kit/net/router"
)

func initInner(r *router.Router) {
	r.VerifyGet("/x/internal/hello", hello)
}

func initOutter(r *router.Router) {
	r.GuestGet("/x/outter/hello", hello)
}
