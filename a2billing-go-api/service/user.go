package service

import (
	"a2billing-go-api/common/auth"
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AgentService struct {
}

func NewAgentService() AgentService {
	return AgentService{}
}
func (service *AgentService) GenerateTokenByApiKey(apiKey string, isRefresh bool) (int, interface{}) {
	log.Debug("AgentService", "GenerateTokenByApiKey", apiKey)
	AgentRes, err := repository.AgentRepo.GetAgentByApiKey(apiKey)
	if err != nil {
		log.Error("AgentService", "GenerateTokenByApiKey", err.Error())
		return response.NotFound()
	}
	if AgentRes == nil {
		log.Error("AgentService", "GenerateTokenByApiKey", "Agent is null")
		return response.NotFound()
	}
	Agent := AgentRes.(model.Agent)

	clientAuth := auth.AuthClient{
		ClienId:  apiKey,
		UserId:   fmt.Sprintf("%d", Agent.ID),
		UserData: map[string]string{},
	}
	client, err := auth.GoAuthClient.ClientCredential(clientAuth, false)
	if err != nil {
		log.Error("AgentService", "GenerateTokenByApiKey", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if err != nil {
		log.Error("AgentService", "GenerateTokenByApiKey", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	token := gin.H{
		"client_id":     client.ClienId,
		"user_id":       client.UserId,
		"token":         client.Token,
		"refresh_token": client.RefreshToken,
		"expired_in":    client.ExpiredIn,
		"token_type":    client.TokenType,
	}
	return response.OK(token)
}
