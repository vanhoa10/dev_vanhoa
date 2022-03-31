package sqlclient

import (
	"github.com/uptrace/bun"
)

type ISqlClientConn interface {
	GetDB() *bun.DB
}

var SqlClientConn ISqlClientConn
