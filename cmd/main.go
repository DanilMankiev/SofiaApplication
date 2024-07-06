package main

import (
	"github.com/DanilMankiev/SofiaApplication/config"
	"github.com/DanilMankiev/SofiaApplication/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.InitConfig()
	if err!=nil{
		logrus.Fatalf("Config error: %s",err)
	}

	app.Run(cfg)
}