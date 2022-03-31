package repository

import (
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
	"errors"

	"github.com/uptrace/bun"
)

type SipBuddiesRepository struct {
}

var SipBuddiesRepo SipBuddiesRepository

func NewSipBuddiesRepository() SipBuddiesRepository {
	repo := SipBuddiesRepository{}
	return repo
}
func (repo *SipBuddiesRepository) CreateSipBuddies(sipBuddies model.SipBuddies) (model.SipBuddies, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&sipBuddies).Exec(ctx)
	if err != nil {
		return sipBuddies, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return sipBuddies, errors.New("create sipBuddies failed")
	}
	return sipBuddies, nil
}
func (repo *SipBuddiesRepository) CreateSipBuddiesTransaction(tx *bun.Tx, sipBuddies model.SipBuddies) (model.SipBuddies, error) {
	resp, err := tx.NewInsert().Model(&sipBuddies).Exec(ctx)
	if err != nil {
		return sipBuddies, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return sipBuddies, errors.New("create sipBuddies failed")
	}
	return sipBuddies, nil
}
