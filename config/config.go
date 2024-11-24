package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Server struct {
		Port string
	}
}

func New(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	return c
}

func (c *Config) PrintEnvVarDbg() {
	dbgstr := "\tConfig.PrintEnvVarDbg()> (need to be deleted in production)\n"
	dbgstr += fmt.Sprintf("- port= %s", c.Server.Port)
	dbgstr += "\n\n"

	fmt.Println(dbgstr)
}
