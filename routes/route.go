package route

import (
	"video_api/handlers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Router *gin.Engine
}

func NewRouter() *Routes {
	r := gin.Default()
	return &Routes{
		Router: r,
	}
}

func (h *Routes) RoutesFunc() {
	h.Router.GET("/videoscount/:id", handlers.ViewCountById)
	h.Router.PUT("/videoview", handlers.ViewVideo)
}
