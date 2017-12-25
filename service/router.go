package service

import (
	"kit/net/router"
)

func initOutter(r *router.Router) {
	r.GuestGet("/x/outter/hello", hello)
}

func initInner(r *router.Router) {
	r.GuestGet("/x/internal/hello", hello)
}
