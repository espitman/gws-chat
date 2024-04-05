package main

import (
	types "github.com/espitman/gws-chat/grpc-gateway/maker/_types"
	util "github.com/espitman/gws-chat/grpc-gateway/maker/_util"
	"strconv"
	"strings"
)

func createHandler(h types.Handler) {
	tmplFile := "./maker/service/handler.tmpl"
	outputFile := "./maker/_files/handler_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func createRouter(h types.Handler) {
	tmplFile := "./maker/service/router.tmpl"
	outputFile := "./maker/_files/router_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func createDto(h types.Handler) {
	tmplFile := "./maker/service/dto.tmpl"
	outputFile := "./maker/_files/dto_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func createClient(h types.Handler) {
	tmplFile := "./maker/service/client.tmpl"
	outputFile := "./maker/_files/client_" + strings.ToLower(h.Name) + "_service.go"
	util.Render(tmplFile, outputFile, h)
}

func getCustomHandler() map[string]map[string]types.Method {
	custom := make(map[string]map[string]types.Method)
	var data struct {
		Service []types.Service `yaml:"services"`
	}
	util.YamlReader("./maker/service/handler_custom.yaml", &data)
	for _, service := range data.Service {
		if custom[service.Name] == nil {
			custom[service.Name] = make(map[string]types.Method)
		}
		for _, method := range service.Method {
			custom[service.Name][method.Name] = method
		}
	}
	return custom
}

func main() {
	customHandler := getCustomHandler()

	var data struct {
		Service []types.Service `yaml:"services"`
	}
	util.YamlReader("./maker/service/handler.yaml", &data)

	for _, service := range data.Service {
		var methods []types.Method
		for _, method := range service.Method {
			cm, ok := customHandler[service.Name][method.Name]
			if ok {
				method = cm
			}
			methods = append(methods, method)
		}

		h := types.Handler{
			Name:        service.Name,
			PB:          service.Name + "pb",
			PBPath:      service.Path,
			Port:        strconv.Itoa(service.Port),
			HandlerName: service.Name + "ServiceHandler",
			ClientName:  service.Name + "ServiceClient",
			RouterName:  service.Name + "ServiceRouter",
			Methods:     methods,
		}

		createHandler(h)
		createRouter(h)
		createDto(h)
		createClient(h)
	}

}
