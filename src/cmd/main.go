package main

import (
	"log"

	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/api"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/config"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/data/cache"
	"github.com/Alireza-Morady/Golang-Clean-Web-API.git/data/db"
)


func main(){
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
