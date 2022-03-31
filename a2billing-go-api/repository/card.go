package repository

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
	"errors"

	"github.com/uptrace/bun"
)

type CardRepository struct {
}

func NewCardRepository() CardRepository {
	repo := CardRepository{}
	return repo
}

var CardRepo CardRepository

func (repo *CardRepository) GetCardsOfAgent(agentId string, limit, offset int) (interface{}, int, error) {
	log.Debug("CardRepository", "GetCardsOfAgent", agentId)
	cards := make([]model.CardInfo, 0)
	count, err := IMySql.SqlClientConn.GetDB().NewSelect().Model(&model.CardInfo{}).
		Table("cc_card c").
		Column("c.id, c.creationdate, c.firstusedate, c.expirationdate, c.enableexpire, c.expiredays, c.username, c.useralias, c.credit, c.activated, c.status, c.lastuse, c.creditlimit, c.id_group, c.tariff").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId).
		Order("c.id ASC").
		Limit(limit).Offset(offset).
		Count(ctx)
	if err != nil {
		return nil, count, err
	}
	return cards, count, nil
}

func (repo *CardRepository) UpdateCCardCreditOfAgent(id string, credit float64) (bool, error) {
	log.Debug("CardRepository", "UpdateCCardCreditOfAgent", id)
	resp, err := IMySql.SqlClientConn.GetDB().NewUpdate().Model(&model.CardInfo{}).
		Table("cc_card").
		Where("id = ?", id).Set("credit", credit).Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, _ := resp.RowsAffected()
	if affected < 1 {
		return false, errors.New("update failed")
	}
	return affected == 1, nil
}

func (repo *CardRepository) UpdateCCardStatusOfAgent(id string, status int) (bool, error) {
	log.Debug("CardRepository", "UpdateCCardCreditOfAgent", id)
	resp, err := IMySql.SqlClientConn.GetDB().NewUpdate().Model(&model.CardInfo{}).
		Table("cc_card").
		Where("id = ?", id).Set("status", status).Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, _ := resp.RowsAffected()
	if affected < 1 {
		return false, errors.New("update failed")
	}
	return affected == 1, nil
}

func (repo *CardRepository) GetCardOfAgentById(agentId, id string) (interface{}, error) {
	log.Debug("CardRepository", "GetCardOfAgentById", agentId)
	var card model.CardInfo
	resp, err := IMySql.SqlClientConn.GetDB().NewSelect().Model(&model.CardInfo{}).
		Table("cc_card c").
		Column("c.id, c.creationdate, c.firstusedate, c.expirationdate, c.enableexpire, c.expiredays, c.username, c.useralias, c.credit, c.activated, c.status, c.lastuse, c.creditlimit, c.id_group, c.tariff").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Join("JOIN cc_callerid ccid ON ccid.id_cc_card = c.id").
		Where("cg.id_agent = ?", agentId).
		Where("c.id = ? OR c.username = ? OR ccid.cid = ?", id, id, id).
		Limit(1).Exec(ctx)
	if err != nil {
		return nil, err
	}
	if affected, _ := resp.RowsAffected(); affected == 0 {
		return nil, nil
	}
	return card, nil
}
func (repo *CardRepository) CreateCard(card model.Card) (model.Card, error) {
	resp, err := IMySql.SqlClientConn.GetDB().NewInsert().Model(&card).Exec(ctx)
	if err != nil {
		return card, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return card, errors.New("create card failed")
	}
	return card, nil
}
func (repo *CardRepository) CreateCardTransaction(tx *bun.Tx, card model.Card) (model.Card, error) {
	resp, err := tx.NewInsert().Model(&card).Exec(ctx)
	if err != nil {
		return card, err
	}
	if affected, _ := resp.RowsAffected(); affected < 1 {
		return card, errors.New("create card failed")
	}
	return card, nil
}
