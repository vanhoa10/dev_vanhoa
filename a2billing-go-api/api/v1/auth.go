package api

import (
	_ "a2billing-go-api/common/data"
	"a2billing-go-api/common/response"
	mdw "a2billing-go-api/middleware/auth"
	Service "a2billing-go-api/service"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	AgentService Service.AgentService
}

func NewAuthHandler(r *gin.Engine, agentService Service.AgentService) {
	handler := &AuthHandler{
		AgentService: agentService,
	}
	Group := r.Group("v1/auth")
	{
		Group.GET("check", mdw.AuthMiddleware(), handler.CheckAuthen)
		Group.POST("token", handler.GenerateToken)
	}
}

func (handler *AuthHandler) CheckAuthen(c *gin.Context) {
	user, isExisted := c.Get("user")
	c.JSON(200, gin.H{
		"isExisted": isExisted,
		"user":      user,
	})
}

func (handler *AuthHandler) GenerateToken(c *gin.Context) {
	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(response.BadRequest())
	}
	apiKey, ok := body["api_key"].(string)
	if !ok || apiKey == "" {
		code, result := response.BadRequestMsg("api_key must not be null")
		c.JSON(code, result)
		return
	}
	isRefresh, ok := body["refresh"].(bool)
	if !ok {
		isRefresh = false
	}
	code, result := handler.AgentService.GenerateTokenByApiKey(apiKey, isRefresh)
	c.JSON(code, result)
}
