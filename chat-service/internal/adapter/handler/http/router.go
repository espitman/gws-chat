package http

//
//import (
//	"github.com/gofiber/fiber/v2"
//)
//
//type Router struct {
//	fiberRouter  fiber.Router
//	messageHandler *MessageHandler
//	// +salvation RouterType
//}
//
//func newRouter(
//	messageHandler *MessageHandler,
//	// +salvation NewRouterType
//) Router {
//	return Router{
//		messageHandler: messageHandler,
//		// +salvation RouterHandler
//	}
//}
//
//func (r *Router) serve(app *fiber.App) {
//	r.fiberRouter = app.Group("/")
//	api := app.Group("/api")
//	v1 := api.Group("/v1")
//
//	r.swaggerRouter()
//	r.messageV1Router(v1)
//	// +salvation RouterServe
//}
