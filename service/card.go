package service

import (
	"a2billing-go-api/common/log"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	IMySql "a2billing-go-api/internal/sqldb/mysql"
	"a2billing-go-api/repository"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type CardService struct {
}

func NewCardService() CardService {
	return CardService{}
}

func (service *CardService) GetCardsOfAgent(agentId string, limit, offset int) (int, interface{}) {
	log.Debug("CardService", "GetCardsOfAgent", agentId)
	cards, total, err := repository.CardRepo.GetCardsOfAgent(agentId, limit, offset)
	if err != nil {
		log.Error("CardService", "GetCardsOfAgent", err.Error())
	}
	return response.NewBaseResponsePagination(cards, limit, offset, int(total))
}

func (service *CardService) GetCardsOfAgentById(agentId string, id string) (int, interface{}) {
	log.Debug("CardService", "GetCardsOfAgent", agentId)
	card, err := repository.CardRepo.GetCardOfAgentById(agentId, id)
	if err != nil {
		log.Error("CardService", "GetCardsOfAgent", err.Error())
	}
	if card == nil {
		return response.NotFound()
	}
	return response.NewResponse(http.StatusOK, card)
}

func (service *CardService) UpdateCardCreditOfAgent(agentId, id string, credit float64) (int, interface{}) {
	log.Debug("CardService", "UpdateCCardCreditOfAgent", agentId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, id)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if cardRes == nil {
		return response.NotFound()
	}
	card, _ := cardRes.(model.CardInfo)
	isUpdated, err := repository.CardRepo.UpdateCCardCreditOfAgent(fmt.Sprintf("%v", card.ID), credit)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		repository.SystemLogRepo.CreateLog(model.SystemLog{
			Iduser:   userId,
			Loglevel: 1,
			Action:   "API EXECUTE",
			Description: sql.NullString{
				Valid:  true,
				String: "Update card " + id + " credit : " + fmt.Sprintf("%v", credit),
			},
			Tablename: sql.NullString{
				Valid:  true,
				String: "cc_card",
			},
			Creationdate: time.Now(),
			Ipaddress: sql.NullString{
				Valid:  true,
				String: "API",
			},
			Pagename: sql.NullString{
				Valid:  true,
				String: "Customer API",
			},
			Agent: sql.NullInt64{
				Valid: true,
				Int64: int64(userId),
			},
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"credit":  credit,
	})
}
func (service *CardService) AddCardCreditOfAgent(agentId, id string, credit float64) (int, interface{}) {
	log.Debug("CardService", "UpdateCCardCreditOfAgent", agentId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, id)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if cardRes == nil {
		return response.NotFound()
	}
	card, _ := cardRes.(model.CardInfo)
	credit = card.Credit + credit
	isUpdated, err := repository.CardRepo.UpdateCCardCreditOfAgent(fmt.Sprintf("%v", card.ID), credit)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		repository.SystemLogRepo.CreateLog(model.SystemLog{
			Iduser:   userId,
			Loglevel: 1,
			Action:   "API EXECUTE",
			Description: sql.NullString{
				Valid:  true,
				String: "Update card " + id + " credit : " + fmt.Sprintf("%v", credit),
			},
			Tablename: sql.NullString{
				Valid:  true,
				String: "cc_card",
			},
			Creationdate: time.Now(),
			Ipaddress: sql.NullString{
				Valid:  true,
				String: "API",
			},
			Pagename: sql.NullString{
				Valid:  true,
				String: "Customer API",
			},
			Agent: sql.NullInt64{
				Valid: true,
				Int64: int64(userId),
			},
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"credit":  credit,
	})
}

func (service *CardService) UpdateCardStatusOfAgent(agentId, id string, status int) (int, interface{}) {
	log.Debug("CardService", "UpdateCCardCreditOfAgent", agentId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, id)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if cardRes == nil {
		return response.NotFound()
	}
	card, _ := cardRes.(model.CardInfo)
	isUpdated, err := repository.CardRepo.UpdateCCardStatusOfAgent(fmt.Sprintf("%v", card.ID), status)
	if err != nil {
		log.Error("CardService", "UpdateCCardCreditOfAgent", err.Error())
		return response.ServiceUnavailableMsg(err.Error())
	}
	if isUpdated == true {
		userId, _ := strconv.Atoi(agentId)
		repository.SystemLogRepo.CreateLog(model.SystemLog{
			Iduser:   userId,
			Loglevel: 1,
			Action:   "API EXECUTE",
			Description: sql.NullString{
				Valid:  true,
				String: "Update card " + id + " status : " + fmt.Sprintf("%v", status),
			},
			Tablename: sql.NullString{
				Valid:  true,
				String: "cc_card",
			},
			Creationdate: time.Now(),
			Ipaddress: sql.NullString{
				Valid:  true,
				String: "API",
			},
			Pagename: sql.NullString{
				Valid:  true,
				String: "Customer API",
			},
			Agent: sql.NullInt64{
				Valid: true,
				Int64: int64(userId),
			},
		})
	}
	return response.NewResponse(http.StatusOK, map[string]interface{}{
		"message": "successfully",
		"id":      id,
		"status":  status,
	})
}

func (service *CardService) CreateCardAndSip(agentId string, card model.Card, cid string) (int, interface{}) {
	log.Debug("CardService", "CreateCardAndSip", agentId)
	cardRes, err := repository.CardRepo.GetCardOfAgentById(agentId, card.Username)
	if err != nil {
		log.Error("CardService", "CreateCardAndSip", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if cardRes != nil {
		return response.BadRequestMsg("username is already exists")
	}
	if groupId, err := repository.AgentRepo.GetGroupIdById(agentId); err != nil {
		log.Error("CardService", "GetGroupIdById", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if groupId < 1 {
		return response.BadRequestMsg("please check group configuration")
	} else {
		card.IDGroup = groupId
	}
	if callerId, err := repository.CallerIdRepo.GetCallerIdByCid(agentId, cid); err != nil {
		log.Error("CardService", "GetCallerIdByCid", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if callerId != nil {
		return response.BadRequestMsg("caller_id is already exists")
	}
	if tariffGroup, err := repository.TariffGroupRepo.GetTariffGroupById(card.Tariff.Int64); err != nil {
		log.Error("CardService", "GetTariffGroupById", err.Error())
		return response.ServiceUnavailableMsg("create customer invalid")
	} else if tariffGroup == nil {
		return response.BadRequestMsg("call_plan is not exists")
	}
	tx := IMySql.MySqlConnector.GetConn().Begin()
	card, err = repository.CardRepo.CreateCardTransaction(tx, card)
	if err != nil {
		log.Error("CardService", "CreateCardAndSip - CreateCard", err.Error())
		tx.Rollback()
		return response.ServiceUnavailableMsg("create customer invalid")
	}
	result := map[string]interface{}{
		"message":   "successfully",
		"id":        int(card.ID),
		"username":  card.Username,
		"password":  card.Uipass,
		"status":    card.Status,
		"call_plan": card.Tariff.Int64,
	}
	if callerId, err := repository.CallerIdRepo.CreateCallerIdTransaction(tx, model.CallerId{Cid: cid, IDCcCard: card.ID, Activated: "t"}); err != nil {
		log.Error("CardService", "GetCallerIdByCid", err.Error())
		tx.Rollback()
		return response.ServiceUnavailableMsg("create customer invalid")
	} else {
		result["cid"] = callerId.Cid
	}
	if card.SipBuddy.Int64 == 1 {
		sipBuddies := model.SipBuddies{
			IDCcCard:    int(card.ID),
			Name:        card.Username,
			Accountcode: card.Username,
			Regexten:    card.Username,
			Amaflags: sql.NullString{
				Valid:  true,
				String: "billing",
			},
			Canreinvite: "YES",
			Context:     "a2billing",
			Dtmfmode:    "RFC2833",
			Host:        "dynamic",
			Nat: sql.NullString{
				Valid:  true,
				String: "no",
			},
			Qualify: sql.NullString{
				Valid:  true,
				String: "no",
			},
			Secret:     card.Uipass,
			Type:       "friend",
			Username:   card.Username,
			Disallow:   "ALL",
			Allow:      "ulaw,alaw,gsm,g729",
			Regseconds: 0,
			Cancallforward: sql.NullString{
				Valid:  true,
				String: "yes",
			},
			Rtpkeepalive: "0",
		}
		sipBuddies, err := repository.SipBuddiesRepo.CreateSipBuddiesTransaction(tx, sipBuddies)
		if err != nil {
			log.Error("CardService", "CreateCardAndSip - CreateSipBuddies", err.Error())
			tx.Rollback()
			return response.ServiceUnavailableMsg("create customer invalid")
		}
		result["sip"] = "created"
	}
	if card.IaxBuddy.Int64 == 1 {
		iaxBuddies := model.IaxBuddies{
			IDCcCard:    int(card.ID),
			Name:        card.Username,
			Accountcode: card.Username,
			Regexten:    card.Username,
			Amaflags: sql.NullString{
				Valid:  true,
				String: "billing",
			},
			Context: "a2billing",
			Host:    "dynamic",
			Qualify: sql.NullString{
				Valid:  true,
				String: "no",
			},
			Secret:     card.Uipass,
			Type:       "friend",
			Username:   card.Username,
			Disallow:   "",
			Allow:      "ulaw,alaw,gsm,g729",
			Regseconds: 0,
			Trunk: sql.NullString{
				Valid:  true,
				String: "no",
			},
		}
		iaxBuddies, err := repository.IaxBuddiesRepo.CreateIaxBuddiesTransaction(tx, iaxBuddies)
		if err != nil {
			log.Error("CardService", "CreateCardAndSip - CreateIaxBuddies", err.Error())
			tx.Rollback()
			return response.ServiceUnavailableMsg("create customer invalid")
		}
		result["iax"] = "created"
	}
	tx.Commit()
	return response.NewResponse(http.StatusOK, result)
}
