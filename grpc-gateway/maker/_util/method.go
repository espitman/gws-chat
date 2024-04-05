package _util

import (
	"github.com/espitman/gws-chat/grpc-gateway/maker/_types"
	"reflect"
	"strings"
)

func GetMethods(ifs reflect.Type, service string) []_types.Method {
	var methods []_types.Method
	for i := 0; i < ifs.NumMethod(); i++ {
		method := ifs.Method(i)
		params := make([]_types.Params, 1)
		params[0] = _types.Params{
			Name:     "",
			Type:     "",
			Kind:     "",
			Required: false,
		}
		if method.IsExported() {
			old := "*" + service + "_service."
			methods = append(methods, _types.Method{
				Name:         method.Name,
				Route:        method.Name,
				SwaggerRoute: method.Name,
				Enable:       true,
				Authorize:    true,
				In:           strings.Replace(method.Type.In(1).String(), old, "", 1),
				Out:          strings.Replace(method.Type.Out(0).String(), old, "", 1),
				Method:       "Post",
				Query:        true,
				Param:        true,
				Body:         true,
				Params:       params,
			})
		}
	}
	return methods
}
