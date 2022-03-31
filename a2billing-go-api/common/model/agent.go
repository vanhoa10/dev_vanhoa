package model

import (
	"database/sql"
	"time"
)

type Agent struct {
	ID           int64          `bun:"column:id;primary_key" json:"id"`
	DateCreation time.Time      `bun:"column:datecreation" json:"datecreation"`
	Active       string         `bun:"column:active" json:"active"`
	Login        string         `bun:"column:login" json:"login"`
	Passwd       sql.NullString `bun:"column:passwd" json:"passwd"`
	ApiKey       sql.NullString `json:"api_key" bun:"column:api_key;type:varchar(200);null"`
}

func (a *Agent) TableName() string {
	return "cc_agent"
}
