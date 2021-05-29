package sancus

import (
	"go.sancus.dev/web"
)

type Handler interface {
	web.Handler
	web.RouterPageInfo
}
