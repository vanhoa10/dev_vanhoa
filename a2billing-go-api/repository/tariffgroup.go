package repository

import (
	"a2billing-go-api/common/model"
	IMySql "a2billing-go-api/internal/sql-client"
)

type TariffGroup struct {
}

var TariffGroupRepo TariffGroup

func NewTariffGroupRepository() TariffGroup {
	repo := TariffGroup{}
	return repo
}

func (repo *TariffGroup) GetTariffGroupById(id int64) (*model.TariffGroup, error) {
	tariffGroup := new(model.TariffGroup)
	resp, err := IMySql.SqlClientConn.GetDB().NewSelect().Model(tariffGroup).
		Table("cc_tariffgroup tg").
		Column("tg.id, tg.iduser, tg.tariffgroupname").
		Where("tg.id = ?", id).
		Limit(1).Exec(ctx)
	if err != nil {
		return nil, err
	} else if affected, _ := resp.RowsAffected(); affected == 0 {
		return nil, nil
	}
	return tariffGroup, nil
}
