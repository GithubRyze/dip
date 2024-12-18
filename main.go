package main

import (
	"dip/bootstrap"
	"dip/bootstrap/conf"
	"dip/bootstrap/dip_logger"
	"dip/cmd/configuration"
	"dip/cmd/dip"
	"log"
)

func main() {
	err := bootstrap.InitApplication()
	if err != nil {
		log.Fatalf("bootstrap app error %s", err.Error())
	}
	go func() {
		dip_logger.Infof("start web server in %s", conf.Config.WebPort)
		err = configuration.StartDipWebConfigurationServer(conf.Config.WebPort)
		if err != nil {
			log.Fatal("start web server failed")
		}
	}()
	dip_logger.Infof("start dip server in %s", conf.Config.EndpointTlsPort)
	err = dip.StartDipHttpServer(conf.Config.EndpointTlsPort)
	if err != nil {
		log.Fatal("start dip server failed")
	}
}
