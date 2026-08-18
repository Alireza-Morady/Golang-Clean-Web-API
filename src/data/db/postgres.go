package db

import (
	"fmt"
	// "log"
	"time"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())
func InitDb(cfg *config.Config)error{
	pg := &cfg.Postgres
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
	pg.Host,pg.Port,pg.User,pg.Password,pg.DbName,pg.SSLMode)
	dbClient, err := gorm.Open(postgres.Open(cnn),&gorm.Config{})
	if err != nil{
		return err
	}
	sqlDb,_ := dbClient.DB()

	err = sqlDb.Ping()
	if err != nil{
		return err
	}
	sqlDb.SetMaxIdleConns(pg.MaxIdleConns)
	sqlDb.SetMaxOpenConns(pg.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(pg.ConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres,logging.Startup,"Db connection established",nil)
	return nil
}
func GetDb()*gorm.DB{
	return dbClient
}

func CloseDb(){
	con,_ := dbClient.DB()
	con.Close()
}