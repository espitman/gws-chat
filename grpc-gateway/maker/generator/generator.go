package main

import (
	"github.com/espitman/gws-chat/grpc-gateway/maker/_types"
	util "github.com/espitman/gws-chat/grpc-gateway/maker/_util"
)

func createMain(s _types.Services) {
	tmplFile := "./maker/generator/main.tmpl"
	outputFile := "./maker/generator/main.go"
	util.Render(tmplFile, outputFile, s)
}

func main() {
	var services _types.Services
	var data struct {
		Service []_types.Service `yaml:"services"`
	}
	util.YamlReader("./maker/service.yaml", &data)

	for _, service := range data.Service {
		services.Services = append(services.Services, _types.Service{
			Name: service.Name,
			Port: service.Port,
			Path: service.Path,
		})
	}

	createMain(services)
}
