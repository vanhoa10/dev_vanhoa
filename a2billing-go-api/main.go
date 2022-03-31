package main

import (
	"a2billing-go-api/api"
	"a2billing-go-api/common/cache"
	IRedis "a2billing-go-api/internal/redis"
	redis "a2billing-go-api/internal/redis/driver"
	sqlclient "a2billing-go-api/internal/sql-client"
	"a2billing-go-api/middleware/auth/goauth"
	"a2billing-go-api/repository"
	"io"
	"os"
	"path/filepath"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Dir      string `env:"CONFIG_DIR" envDefault:"config/config.json"`
	Port     string
	LogType  string
	LogLevel string
	LogFile  string
	LogAddr  string
	DB       string
	Redis    string
	Auth     string
}

var config Config

func init() {
	if err := env.Parse(&config); err != nil {
		log.Error("Get environment values fail")
		log.Fatal(err)
	}
	viper.SetConfigFile(config.Dir)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
		panic(err)
	}
	cfg := Config{
		Dir:      config.Dir,
		Port:     viper.GetString(`main.port`),
		LogType:  viper.GetString(`main.log_type`),
		LogLevel: viper.GetString(`main.log_level`),
		LogFile:  viper.GetString(`main.log_file`),
		LogAddr:  viper.GetString(`main.log_addr`),
		DB:       viper.GetString(`main.db`),
		Redis:    viper.GetString(`main.redis`),
		Auth:     viper.GetString(`main.auth`),
	}
	if cfg.DB == "enabled" {
		switch viper.GetString(`db.driver`) {
		case "postgresql":
			sqlClientConfig := sqlclient.SqlConfig{
				Driver:       "postgresql",
				Host:         viper.GetString(`db.host`),
				Database:     viper.GetString(`db.database`),
				Username:     viper.GetString(`db.username`),
				Password:     viper.GetString(`db.password`),
				Port:         viper.GetInt(`db.port`),
				DialTimeout:  20,
				ReadTimeout:  30,
				WriteTimeout: 30,
				Timeout:      30,
				PoolSize:     10,
				MaxIdleConns: 10,
				MaxOpenConns: 10,
			}
			repository.SqlClient = sqlclient.NewSqlClient(sqlClientConfig)
		case "mysql":
			sqlClientConfig := sqlclient.SqlConfig{
				Driver:       "mysql",
				Host:         viper.GetString(`db.host`),
				Database:     viper.GetString(`db.database`),
				Username:     viper.GetString(`db.username`),
				Password:     viper.GetString(`db.password`),
				Port:         viper.GetInt(`db.port`),
				DialTimeout:  20,
				ReadTimeout:  30,
				WriteTimeout: 30,
				Timeout:      30,
				PoolSize:     10,
				MaxIdleConns: 10,
				MaxOpenConns: 10,
			}
			repository.SqlClient = sqlclient.NewSqlClient(sqlClientConfig)
		}
	}
	if cfg.Redis == "enabled" {
		var err error
		IRedis.Redis, err = redis.NewRedis(redis.Config{
			Addr:         viper.GetString(`redis.address`),
			Password:     viper.GetString(`redis.password`),
			DB:           viper.GetInt(`redis.database`),
			PoolSize:     30,
			PoolTimeout:  20,
			IdleTimeout:  10,
			ReadTimeout:  20,
			WriteTimeout: 15,
		})
		if err != nil {
			panic(err)
		}
	}
	switch cfg.Auth {
	case "oauth":
		var err error
		goauth.GoAuthClient, err = goauth.NewGoAuth(goauth.GoAuth{
			RedisExpiredIn: viper.GetInt(`oauth.expired_in`),
			TokenType:      viper.GetString(`oauth.tokenType`),
			RedisTokenKey:  "access_token_key",
			RedisUserKey:   "access_token_user",
			RedisClient:    IRedis.Redis.GetClient(),
		})
		if err != nil {
			panic(err)
		}
	}
	config = cfg
}

func main() {
	_ = os.Mkdir(filepath.Dir(config.LogFile), 0755)
	file, _ := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	cache.MemCache = cache.NewCache()
	defer cache.MemCache.Close()
	setAppLogger(config, file)
	server := api.NewServer()
	server.Start(config.Port)
}

func setAppLogger(cfg Config, file *os.File) {
	// log.SetFormatter(&log.TextFormatter{
	// 	FullTimestamp: true,
	// })
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
	switch cfg.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	switch cfg.LogType {
	case "DEFAULT":
		log.SetOutput(os.Stdout)
	case "FILE":
		if file != nil {
			log.SetOutput(io.MultiWriter(os.Stdout, file))
		} else {
			log.Error("main ", "Log File "+cfg.LogFile+" error")
			log.SetOutput(os.Stdout)
		}
	default:
		log.SetOutput(os.Stdout)
	}
}
