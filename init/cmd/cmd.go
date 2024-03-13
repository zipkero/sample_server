package cmd

import (
	"fmt"
	"sample_server/config"
	"sample_server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	fmt.Println("Server is running on port", c.config.Server.Port)

	if err := c.network.Run(c.config.Server.Port); err != nil {
		panic(err)
	}
	return c
}
