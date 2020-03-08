package ginApp

import (
	"github.com/chenyingdi/gin-app/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	Addr   string
	DB     *gorm.DB
	Engine *gin.Engine
	Cache  *redis.Client
}

func New() *App {
	return &App{}
}

func (a *App) Init(cfg *config.Config) error {
	err := a.initDB(&cfg.DBConfig)
	if err != nil {
		return err
	}

	err = a.initCache(&cfg.CacheConfig)
	if err != nil {
		return err
	}

	a.Engine = gin.Default()
	a.Addr = cfg.SrvConfig.ParseUrl()
	gin.SetMode(cfg.SrvConfig.Mode)

	return nil
}

func (a *App) initDB(cfg *config.DBConfig) error {
	var err error
	a.DB, err = gorm.Open(
		cfg.Dialect,
		cfg.ParseUrl(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initCache(cfg *config.CacheConfig) error {
	a.Cache = redis.NewClient(&redis.Options{
		DB:       cfg.DB,
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	_, err := a.Cache.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Run() error {
	return a.Engine.Run(a.Addr)
}

