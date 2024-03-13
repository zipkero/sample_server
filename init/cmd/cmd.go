package cmd

import "server/config"

type Cmd struct {
	config *config.Config
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}
	return c
}
