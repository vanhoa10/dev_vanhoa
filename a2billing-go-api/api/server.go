package api

import (
	"net/http"
	"time"

	responseTimeMdw "a2billing-go-api/middleware/response_time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	serviceName = "a2billing-go-api"
	version     = "v1.2.0"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.New()
	// engine.Use(cors.Default())
	engine.Use(gin.Recovery())
	engine.MaxMultipartMemory = 100 << 20
	engine.Use(CORSMiddleware())
	engine.Use(responseTimeMdw.Handler)
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": serviceName,
			"version": version,
			"time":    time.Now().Unix(),
		})
	})
	server := &Server{Engine: engine}
	return server
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func (server *Server) Start(port string) {
	v := make(chan struct{})
	go func() {
		if err := server.Engine.Run(":" + port); err != nil {
			log.WithError(err).Error("failed to start service")
			close(v)
		}
	}()
	log.Infof("service %v listening on port %v", serviceName, port)
	<-v
}
