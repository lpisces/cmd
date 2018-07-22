package boot

import (
	"gopkg.in/ini.v1"
	"gopkg.in/urfave/cli.v1"
)

var (
	Conf *Config
)

type (
	Config struct {
		Debug      bool
		ConfigFile string
	}
)

func init() {
	Conf = Default()
}

// DefaultConfig get default config
func Default() *Config {

	return &Config{
		false,
		"./config.ini",
	}

}

// LoadFromIni load config from ini override default config
func (config *Config) LoadFromIni() (err error) {
	return ini.MapTo(config, config.ConfigFile)
}

// Load load config from command line param
func (config *Config) Load(c *cli.Context) (err error) {

	if c.String("config") != "" {
		Conf.ConfigFile = c.String("config")
		if err = Conf.LoadFromIni(); err != nil {
			return
		}
	}

	if c.Bool("debug") {
		Conf.Debug = true
	}

	return
}
