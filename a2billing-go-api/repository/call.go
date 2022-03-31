package repository

import (
	"a2billing-go-api/common/log"
	IMySql "a2billing-go-api/internal/sql-client"
)

type CallRepository struct {
}

func NewCallRepository() CallRepository {
	repo := CallRepository{}
	return repo
}

var CallRepo CallRepository

func (repo *CallRepository) GetCallLogs(agentId, cardId, source string, fromDate, toDate string, limit, offset int) (interface{}, int, error) {
	log.Debug("CallRepository", "GetCallLogs", agentId)
	result := []map[string]interface{}{}
	var count int64
	field := "ca.id, ca.uniqueid as uniqueid, ca.starttime as start_time, ca.stoptime as end_time, ca.sessiontime as duration, ca.sessionbill as amount, ca.src as source, ca.calledstation as destination"
	query := IMySql.SqlClientConn.GetDB().NewSelect().
		Table("cc_call ca").
		Column(field).
		Join("JOIN cc_card c ON c.id = ca.card_id").
		Join("JOIN cc_card_group cg ON c.id_group = cg.id").
		Where("cg.id_agent = ?", agentId)
	if len(cardId) > 0 {
		query = query.Where("ca.card_id = ?", cardId)
	}
	if len(source) > 0 {
		query = query.Where("ca.src = ?", source)
	}
	if fromDate != "" {
		query = query.Where("ca.starttime >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("ca.starttime <= ?", toDate)
	}
	_, err := query.NewSelect().Limit(limit).Offset(offset).Order("ca.starttime DESC").Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return result, int(count), nil
}
