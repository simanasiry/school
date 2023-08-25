package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"time"
)

func (a *App) initEngine() {
	a.gin = gin.Default()
	a.httpServer = &http.Server{
		Addr:    a.config.Http.Address,
		Handler: a.gin,
	}
}

func (a *App) initDB() {

	if a.config.Sql.Dialect == "" {
		panic("sql dialect is not determined")
	}
	connStr := a.config.Sql.ConnectionString
	if connStr == "" {
		connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			a.config.Sql.Host,
			a.config.Sql.Port,
			a.config.Sql.Username,
			a.config.Sql.DbName,
			a.config.Sql.Password,
		)
	}
	var dialect gorm.Dialector
	switch a.config.Sql.Dialect {
	case "mysql":
		dialect = mysql.Open(connStr)

	case "postgres":
		//dialect = postgres.New(postgres.Config{
		//	DSN:                  connStr,
		//	PreferSimpleProtocol: true,
		//})
	}
	d, err := gorm.Open(dialect, &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	a.db = d

}

func (a *App) initConfig() {
	a.panicOnErr(a.registry.Unmarshal(&a.config))
	a.config.Initialize()
}
