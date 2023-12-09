package router

import (
	"github.com/EasyCode-Platform/app-backend/src/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		Controller: controller,
	}
}

func (r *Router) RegisterRouters(engine *gin.Engine) {
	// config
	engine.UseRawPath = true

	// init route
	routerGroup := engine.Group("/api/v1")

	postgresRouter := routerGroup.Group("/postgres")
	imageRouter := routerGroup.Group("/image")

	//relational database request routers
	postgresRouter.POST("/excecute/:sql", r.Controller.ExecutePostgresSql)
	postgresRouter.POST("/validate/:sql", r.Controller.ValidatePostgresSql)
	postgresRouter.POST("/createTable", r.Controller.CreateTable)
	postgresRouter.POST("/insert", r.Controller.InsertRecord)
	postgresRouter.POST("/records", r.Controller.DisplayTable)
	postgresRouter.POST("/remove", r.Controller.RemoveRecord)

	//image router
	imageRouter.POST("/upload", r.Controller.UploadImage)
}
