package repository

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
	"errors"

	"github.com/uptrace/bun"
)

type CallerIdRepository struct {
}

var CallerIdRepo CallerIdRepository

func NewCallerIdRepository() CallerIdRepository {
	repo := CallerIdRepository{}
	return repo
}

func (repo *CallerIdRepository) GetCallerIdByCid(agentId, cid string) (interface{}, error) {
	log.Debug("CallerIdRepository", "GetCallerIdByCid", cid)
	var callerId model.CallerId
	resp, err := IMySql.SqlClientConn.GetDB().NewSelect().Model(&model.CallerId{}).
		Table("cc_callerid ccid").
		Column("ccid.id, ccid.cid, ccid.id_cc_card, ccid.activated").
		Join("INNER JOIN cc_card c ON c.id = ccid.id_cc_card").
		Join("INNER JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId).
		Where("ccid.cid = ?", cid).
		Limit(1).Exec(ctx)
	if resp != nil {
		return nil, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return nil, nil
	}
	return callerId, nil
}

func (repo *CallerIdRepository) CreateCallerId(callerId model.CallerId) (model.CallerId, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&callerId).Exec(ctx)
	if err != nil {
		return callerId, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return callerId, errors.New("create callerId failed")
	}
	return callerId, nil
}
func (repo *CallerIdRepository) CreateCallerIdTransaction(tx *bun.Tx, callerId model.CallerId) (model.CallerId, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&callerId).Exec(ctx)
	if err != nil {
		return callerId, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return callerId, errors.New("create callerId failed")
	}
	return callerId, nil
}
func (repo *CallerIdRepository) UpdateCallerIdToCard(id int, cardId int) (bool, error) {
	log.Debug("CardRepository", "UpdateCCardCreditOfAgent", id)
	resp, err := IMySql.SqlClientConn.GetDB().NewUpdate().Model(&model.CardInfo{}).
		Table("cc_callerid").
		Where("id = ?", id).Set("id_cc_card", cardId).Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, _ := resp.RowsAffected()
	if affected < 1 {
		return false, errors.New("update failed")
	}
	return affected == 1, nil
}
