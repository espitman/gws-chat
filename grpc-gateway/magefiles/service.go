//go:build mage
// +build mage

package main

import (
	"bufio"
	"fmt"
	cli "github.com/espitman/go-super-cli"
	"github.com/espitman/gws-chat/grpc-gateway/maker/_types"
	util "github.com/espitman/gws-chat/grpc-gateway/maker/_util"
	"os"
	"strconv"
)

func addGRPCService() error {
	serviceName := cli.TextInput("Enter Service name:", "", false)
	servicePort := cli.TextInput("Enter Service port:", "", false)
	servicePath := cli.TextInput("Enter Service path:", "", false)

	yamlPath := "./maker/service.yaml"

	var data struct {
		Service []_types.Service `yaml:"services"`
	}

	util.YamlReader(yamlPath, &data)
	port, _ := strconv.Atoi(servicePort)
	newService := _types.Service{
		Name:   serviceName,
		Port:   port,
		Path:   servicePath,
		Method: nil,
	}
	data.Service = append(data.Service, newService)
	err := util.YamlWriter(yamlPath, data)
	if err != nil {
		return err
	}
	return nil
}

func readInput(prompt string) (string, error) {

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return input, err
}

func (Add) Service() error {
	addGRPCService()
	return nil
}
