package bootstrap

import (
	"dip/bootstrap/conf"
	"dip/bootstrap/dip_logger"
)

func InitApplication() error {
	err := conf.LoadAppConfig("conf.yaml")
	if err != nil {
		return err
	}
	dip_logger.InitAppLogger()
	return nil
}
