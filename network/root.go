package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	n := &Network{
		engin: gin.Default(),
	}
	return n
}

func (n *Network) Run(port string) error {
	return n.engin.Run(port)
}
