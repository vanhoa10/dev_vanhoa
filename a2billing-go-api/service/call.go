package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository"
)

type CallService struct {
}

func NewCallService() CallService {
	return CallService{}
}

func (service *CallService) GetCallLogs(agentId, cardId, source, fromDate, toDate string, limit, offset int) (int, interface{}) {
	log.Debug("CallService", "GetCallLogs", cardId)
	callLogs, total, err := repository.CallRepo.GetCallLogs(agentId, cardId, source, fromDate, toDate, limit, offset)
	if err != nil {
		log.Error("ReportService", "GetAgentCallLogs", err.Error())
	}
	return response.Pagination(callLogs, limit, offset, int(total))
}
