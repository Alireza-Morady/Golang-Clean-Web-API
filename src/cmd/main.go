package main

import (
	// "log"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/data/cache"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/data/db"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/pkg/logging"
)


func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil{
		logger.Fatal(logging.Redis,logging.Startup,err.Error(),nil)
	}

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres,logging.Startup,err.Error(),nil)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
