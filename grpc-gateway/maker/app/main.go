package main

import (
	types "github.com/espitman/gws-chat/grpc-gateway/maker/_types"
	util "github.com/espitman/gws-chat/grpc-gateway/maker/_util"
)

func createRouter(h types.Handlers) {
	tmplFile := "./maker/app/router.tmpl"
	outputFile := "./maker/_files/router.go"
	util.Render(tmplFile, outputFile, h)
}

func createMain(h types.Handlers) {
	tmplFile := "./maker/app/main.tmpl"
	outputFile := "./maker/_files/main.go"
	util.Render(tmplFile, outputFile, h)
}

func main() {
	var data struct {
		Service []types.Service `yaml:"services"`
	}
	util.YamlReader("./maker/service.yaml", &data)

	var handlers []types.Handler
	for _, service := range data.Service {
		handlers = append(handlers, types.Handler{
			Name:        service.Name,
			PB:          service.Name + "pb",
			PBPath:      service.Path,
			HandlerName: service.Name + "ServiceHandler",
			ClientName:  service.Name + "ServiceClient",
			Methods:     nil,
			RouterName:  service.Name + "ServiceRouter",
		})
	}
	h := types.Handlers{
		Handlers: handlers,
	}
	createRouter(h)
	createMain(h)
}
