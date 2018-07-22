package boot

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
)

func Run(c *cli.Context) (err error) {
	Conf.Load(c)
	log.Info(Conf)
	return
}
