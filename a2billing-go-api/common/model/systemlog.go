package model

import (
	"database/sql"
	"time"
)

type SystemLog struct {
	ID           int            `bun:"column:id;primary_key" json:"id"`         //
	Iduser       int            `bun:"column:iduser" json:"iduser"`             //
	Loglevel     int            `bun:"column:loglevel" json:"loglevel"`         //
	Action       string         `bun:"column:action" json:"action"`             //
	Description  sql.NullString `bun:"column:description" json:"description"`   //
	Data         []byte         `bun:"column:data" json:"data"`                 //
	Tablename    sql.NullString `bun:"column:tablename" json:"tablename"`       //
	Pagename     sql.NullString `bun:"column:pagename" json:"pagename"`         //
	Ipaddress    sql.NullString `bun:"column:ipaddress" json:"ipaddress"`       //
	Creationdate time.Time      `bun:"column:creationdate" json:"creationdate"` //
	Agent        sql.NullInt64  `bun:"column:agent" json:"agent"`               //
}

// TableName sets the insert table name for this struct type
func (s *SystemLog) TableName() string {
	return "cc_system_log"
}
