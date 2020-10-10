package main

import (
	"github.com/beldmian/habitFarmer/pkg/config"
	"github.com/beldmian/habitFarmer/pkg/server"
	"log"
)

func main() {
	localConfig, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := server.New(localConfig.BindAddr, localConfig.DbUrl)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}