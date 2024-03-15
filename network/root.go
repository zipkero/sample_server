package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin  *gin.Engine
	router *Router
}

func NewNetwork() *Network {
	engin := gin.Default()
	n := &Network{
		engin:  engin,
		router: NewRouter(engin),
	}

	return n
}

func (n *Network) Run(port string) error {
	return n.engin.Run(port)
}
