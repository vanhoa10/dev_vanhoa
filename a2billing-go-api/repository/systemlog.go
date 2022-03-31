package repository

import (
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
	"errors"
)

type SystemLogRepository struct {
}

func NewSystemLogRepository() SystemLogRepository {
	repo := SystemLogRepository{}
	return repo
}

var SystemLogRepo SystemLogRepository

func (repo *SystemLogRepository) CreateLog(systemLog model.SystemLog) (bool, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&systemLog).Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, _ := resp.RowsAffected()
	if affected < 1 {
		return false, errors.New("create failed")
	}
	return affected == 1, nil
}
