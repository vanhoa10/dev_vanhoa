package repository

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	IMySql "dev_vanhoa/internal/sql-client"
)

type AgentRepository struct {
}

func NewAgentRepository() AgentRepository {
	repo := AgentRepository{}
	repo.SyncTable()
	return repo
}

var AgentRepo AgentRepository

func (repo *AgentRepository) SyncTable() error {
	err := IMySql.SqlClientConn.Set("gorm:table_options", "ENGINE=InnoDB COLLATE utf8_general_ci").AutoMigrate(&model.Agent{})
	if err != nil {
		log.Error("AgentRepository", "SyncTable", err.Error())
		return err
	}
	return nil
}

func (repo *AgentRepository) GetAgentByApiKey(apiKey string) (interface{}, error) {
	log.Debug("AgentRepository", "GetAgentByApiKey", apiKey)
	var Agent model.Agent
	_, err := IMySql.SqlClientConn.GetDB().NewSelect().Model(&model.Agent{}).Where("api_key = ?", apiKey).Limit(1).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return Agent, nil
}

func (repo *AgentRepository) GetGroupIdById(id string) (int, error) {
	log.Debug("AgentRepository", "GetAgentGroupById", id)
	var groupId int
	_, err := IMySql.MySqlConnector.GetConn().NewSelect().Column("id").Table("cc_card_group").Where("id_agent=?", id).Exec(ctx)
	if err != nil {
		return -1, err
	}
	return groupId, nil
}
