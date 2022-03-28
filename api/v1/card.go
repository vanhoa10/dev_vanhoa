package api

import (
	dataUtil "a2billing-go-api/common/data"
	"a2billing-go-api/common/model"
	"a2billing-go-api/common/response"
	"a2billing-go-api/common/validator"
	mdw "a2billing-go-api/internal/middleware"
	"a2billing-go-api/service"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	CardService service.CardService
}

func NewCardHandler(r *gin.Engine, CardService service.CardService) {
	handler := &CardHandler{
		CardService: CardService,
	}
	Group := r.Group("v1/customer")
	{
		Group.GET("", mdw.AuthMiddleware(), handler.GetCards)
		Group.GET(":id", mdw.AuthMiddleware(), handler.GetCardById)
		Group.PUT(":id/credit", mdw.AuthMiddleware(), handler.UpdateCardCredit)
		Group.PUT(":id/credit/add", mdw.AuthMiddleware(), handler.AddCardCredit)
		Group.PUT(":id/status", mdw.AuthMiddleware(), handler.UpdateCardStatus)
		Group.POST("", mdw.AuthMiddleware(), handler.CreateCard)

	}
}

func (handler *CardHandler) GetCards(c *gin.Context) {
	id, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	limit := dataUtil.ParseLimit(c.Query("limit"))
	offset := dataUtil.ParseOffset(c.Query("offset"))
	code, result := handler.CardService.GetCardsOfAgent(id.(string), limit, offset)
	c.JSON(code, result)
}

func (handler *CardHandler) GetCardById(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	id := c.Param("id")
	code, result := handler.CardService.GetCardsOfAgentById(userId.(string), id)
	c.JSON(code, result)
}

func (handler *CardHandler) UpdateCardCredit(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	id := c.Param("id")
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("update-card-credit.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	credit, _ := jsonBody["credit"].(float64)
	code, result := handler.CardService.UpdateCardCreditOfAgent(userId.(string), id, credit)
	c.JSON(code, result)
}

func (handler *CardHandler) AddCardCredit(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	id := c.Param("id")
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("update-card-credit.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	credit, _ := jsonBody["credit"].(float64)
	code, result := handler.CardService.AddCardCreditOfAgent(userId.(string), id, credit)
	c.JSON(code, result)
}

func (handler *CardHandler) UpdateCardStatus(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	id := c.Param("id")
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("update-card-status.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	status, _ := jsonBody["status"].(string)
	statusInt := dataUtil.STATUS_MAP[status]
	code, result := handler.CardService.UpdateCardStatusOfAgent(userId.(string), id, statusInt)
	c.JSON(code, result)
}

func (handler *CardHandler) CreateCard(c *gin.Context) {
	userId, ok := mdw.GetUserId(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	jsonBody := make(map[string]interface{})
	if err := c.BindJSON(&jsonBody); err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, validSchema := validator.CheckSchema("create-card.json", jsonBody)
	if code != http.StatusOK {
		c.JSON(code, validSchema)
		return
	}
	card := model.Card{
		Address:        "",
		Expirationdate: time.Time{},
		Creationdate:   time.Now(),
		Firstusedate:   time.Time{},
		Enableexpire: sql.NullInt64{
			Valid: true,
			Int64: 0,
		},
		Expiredays: sql.NullInt64{
			Valid: true,
			Int64: 0,
		},
		Tariff: sql.NullInt64{
			Valid: true,
			Int64: 1,
		},
		IDDidgroup: sql.NullInt64{
			Valid: true,
			Int64: 0,
		},
		Activated: "1",
		Status:    1,
		Lastname:  fmt.Sprintf("AndDuong%d", time.Now().Unix()),
		Firstname: "SmartCallCenter",
		City:      "",
		State:     "",
		Country:   "USA",
		Zipcode:   "",
		Phone:     "",
		Email:     "",
		Fax:       "",
		Inuse: sql.NullInt64{
			Valid: true,
			Int64: 0,
		},
		Simultaccess: sql.NullInt64{
			Valid: true,
			Int64: 1,
		},
		Currency: sql.NullString{
			Valid:  true,
			String: "VND",
		},
		CreditNotification: -1,
		IDGroup:            0,
		Language: sql.NullString{
			Valid:  true,
			String: "en",
		},
		SipBuddy: sql.NullInt64{
			Valid: true,
			Int64: 1,
		},
		IaxBuddy: sql.NullInt64{
			Valid: true,
			Int64: 1,
		},
		Invoiceday: sql.NullInt64{
			Valid: true,
			Int64: 1,
		},
		MaxConcurrent: 10,
	}
	card.Username, _ = jsonBody["username"].(string)
	card.Useralias, _ = jsonBody["username"].(string)
	card.Uipass, _ = jsonBody["password"].(string)
	card.Credit, _ = jsonBody["credit"].(float64)
	typePaidFloat, _ := jsonBody["type_paid"].(float64)
	typePaid := int64(typePaidFloat)
	card.Typepaid = sql.NullInt64{
		Int64: typePaid,
		Valid: true,
	}
	if typePaid == 1 {
		creditLimitFloat, _ := jsonBody["credit_limit"].(float64)
		creditLimit := int64(creditLimitFloat)
		card.Creditlimit.Valid = true
		card.Creditlimit.Int64 = creditLimit
	} else {
		card.Creditlimit.Valid = true
		card.Creditlimit.Int64 = 0
	}

	if tariffGroup, ok := jsonBody["call_plan"].(float64); ok {
		card.Tariff = sql.NullInt64{
			Valid: true,
			Int64: int64(tariffGroup),
		}
	}
	cid, _ := jsonBody["cid"].(string)
	if cid == "" {
		c.JSON(response.BadRequestMsg("cid is missing"))
		return
	}
	code, result := handler.CardService.CreateCardAndSip(userId.(string), card, cid)
	c.JSON(code, result)
}
