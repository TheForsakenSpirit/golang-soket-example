package server

import (
	"github.com/gin-gonic/gin"

	controllers "example/socket/Controller"
	model "example/socket/Model"
)

func BuildRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	builder := GetRouterBuilder(api)
	builder.
		addController(new(controllers.User)).
		addController(new(controllers.Ping))

}

func GetRouterBuilder(group *gin.RouterGroup) *routerBuilder {
	db := make(map[string]model.User)
	testUser := model.User{
		Username: "test1",
		Password: "test1",
		Key:      model.GenerateKey(32),
	}

	db[testUser.Key] = testUser

	if group == nil || db == nil {
		panic("Router group or db not initialized")
	}

	return &routerBuilder{
		api: group,
		db:  db,
	}
}

type routerBuilder struct {
	api *gin.RouterGroup
	db  map[string]model.User
}

type Controller interface {
	Init(group *gin.RouterGroup, db map[string]model.User)
}

func (builder *routerBuilder) addController(controller Controller) *routerBuilder {
	controller.Init(builder.api, builder.db)
	return builder
}
