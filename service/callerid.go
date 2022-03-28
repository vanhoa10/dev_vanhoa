package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	"a2billing-go-api/repository"
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

type CallerIdService struct {
}

func NewCallerIdService() CallerIdService {
	return CallerIdService{}
}

func (service *CallerIdService) AddCallerIdToCard(agentId, cardId, cid string) (int, interface{}) {
	log.Debug("CallerIdService", "UpdateCallerIdToCard", cardId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, cardId)
	if err != nil {
		log.Error("CallerIdService", "UpdateCallerIdToCard", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if cardRes == nil {
		return response.BadRequestMsg("user_id is not exists")
	}
	card, _ := cardRes.(model.CardInfo)
	callerIdRes, err := repository.CallerIdRepo.GetCallerIdByCid(agentId, cid)
	if err != nil {
		log.Error("CallerIdService", "UpdateCallerIdToCard", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerIdRes != nil {
		return response.BadRequestMsg("caller_id is already exists")
	}
	if callerId, err := repository.CallerIdRepo.CreateCallerId(model.CallerId{Cid: cid, IDCcCard: card.ID, Activated: "t"}); err != nil {
		log.Error("CardService", "GetCallerIdByCid", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerId.ID < 1 {
		return response.ServiceUnavailableMsg("create customer invalid")
	} else {
		userId, _ := strconv.Atoi(agentId)
		repository.SystemLogRepo.CreateLog(model.SystemLog{
			Iduser:   userId,
			Loglevel: 1,
			Action:   "API EXECUTE",
			Description: sql.NullString{
				Valid:  true,
				String: "Create callerid " + cid + " card_id : " + cardId,
			},
			Tablename: sql.NullString{
				Valid:  true,
				String: "cc_callerid",
			},
			Creationdate: time.Now(),
			Ipaddress: sql.NullString{
				Valid:  true,
				String: "API",
			},
			Pagename: sql.NullString{
				Valid:  true,
				String: "Caller API",
			},
			Agent: sql.NullInt64{
				Valid: true,
				Int64: int64(userId),
			},
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"cid":     cid,
		"user_id": cardId,
	})
}

func (service *CallerIdService) UpdateCallerIdToCard(agentId, cid, cardId string) (int, interface{}) {
	log.Debug("CallerIdService", "UpdateCallerIdToCard", cardId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, cardId)
	if err != nil {
		log.Error("CallerIdService", "UpdateCallerIdToCard", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if cardRes == nil {
		return response.BadRequestMsg("user_id is not exists")
	}
	card, _ := cardRes.(model.CardInfo)

	callerIdRes, err := repository.CallerIdRepo.GetCallerIdByCid(agentId, cid)
	if err != nil {
		log.Error("CallerIdService", "UpdateCallerIdToCard", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerIdRes == nil {
		return response.BadRequestMsg("caller_id is not exists")
	}
	callerId, _ := callerIdRes.(model.CallerId)
	isUpdated, err := repository.CallerIdRepo.UpdateCallerIdToCard(int(callerId.ID), int(card.ID))
	if err != nil {
		log.Error("CallerIdService", "UpdateCallerIdToCard", err.Error())
		return response.ServiceUnavailableMsg("update callerid invalid")
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		repository.SystemLogRepo.CreateLog(model.SystemLog{
			Iduser:   userId,
			Loglevel: 1,
			Action:   "API EXECUTE",
			Description: sql.NullString{
				Valid:  true,
				String: "Update callerid " + cid + " card_id : " + cardId,
			},
			Tablename: sql.NullString{
				Valid:  true,
				String: "cc_callerid",
			},
			Creationdate: time.Now(),
			Ipaddress: sql.NullString{
				Valid:  true,
				String: "API",
			},
			Pagename: sql.NullString{
				Valid:  true,
				String: "Caller API",
			},
			Agent: sql.NullInt64{
				Valid: true,
				Int64: int64(userId),
			},
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"cid":     cid,
		"user_id": cardId,
	})
}
