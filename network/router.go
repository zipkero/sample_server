package network

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	routerInit     sync.Once
	routerInstance *Router
)

type Router struct {
	engin *gin.Engine
}

func NewRouter(engin *gin.Engine) *Router {
	routerInit.Do(func() {
		routerInstance = &Router{
			engin: engin,
		}

		routerInstance.engin.POST("/", routerInstance.create)
		routerInstance.engin.PUT("/", routerInstance.update)
		routerInstance.engin.DELETE("/", routerInstance.delete)
		routerInstance.engin.GET("/", routerInstance.get)
	})
	return routerInstance
}

func (r *Router) create(c *gin.Context) {

}

func (r *Router) update(c *gin.Context) {

}

func (r *Router) get(c *gin.Context) {

}

func (r *Router) delete(c *gin.Context) {

}
