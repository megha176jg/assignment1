package main

import (
	"video_api/dummy"
	route "video_api/routes"
)

func main() {

	dummy.DummyData()
	r := route.NewRouter()
	r.RoutesFunc()
	r.Router.Run("localhost:8000")
}
