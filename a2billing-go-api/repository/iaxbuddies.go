package repository

import (
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
	"errors"

	"github.com/uptrace/bun"
)

type IaxBuddiesRepository struct {
}

var IaxBuddiesRepo IaxBuddiesRepository

func NewIaxBuddiesRepository() IaxBuddiesRepository {
	repo := IaxBuddiesRepository{}
	return repo
}
func (repo *IaxBuddiesRepository) CreateIaxBuddies(iaxBuddies model.IaxBuddies) (model.IaxBuddies, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&iaxBuddies).Exec(ctx)
	if err != nil {
		return iaxBuddies, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return iaxBuddies, errors.New("create iaxBuddies failed")
	}
	return iaxBuddies, nil
}
func (repo *IaxBuddiesRepository) CreateIaxBuddiesTransaction(tx *bun.Tx, iaxBuddies model.IaxBuddies) (model.IaxBuddies, error) {
	resp, err := tx.NewInsert().Model(&iaxBuddies).Exec(ctx)
	if err != nil {
		return iaxBuddies, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return iaxBuddies, errors.New("create iaxBuddies failed")
	}
	return iaxBuddies, nil
}
