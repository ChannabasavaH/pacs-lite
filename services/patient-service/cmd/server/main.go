package main

import (
	"patient-service/internal/database"
	"patient-service/internal/config"
)

func main(){
	cfg := config.DBLoadConfig()
	db := database.ConnectPostgres(cfg)
}