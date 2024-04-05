//go:build mage
// +build mage

package main

import (
	"github.com/espitman/go-super-cli"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

//-------------------Service

func generateService() error {
	return sh.RunV("go", "run", "./maker/generator/generator.go")
}

func generateMain() error {
	return sh.RunV("go", "run", "./maker/generator/main.go")
}

func moveHandler() error {
	return sh.RunV("mv", "./maker/generator/handler.yaml", "./maker/service/handler.yaml")
}

func generateMainService() error {
	return sh.RunV("go", "run", "./maker/service/main.go")
}

//-------------------Ask

func ask() {
	_ = cli.TextInput("Customize handler_custom.yaml and press enter", "", false)
}

//-------------------App

func generateApp() error {
	return sh.RunV("go", "run", "./maker/app/main.go")
}

func formatFiles() error {
	return sh.RunV("sh", "-c", "cd ./maker/_files && gofmt -w *.go")
}

func moveFiles() error {
	return sh.RunV("sh", "-c", "cd ./maker/_files && mv *.go ../..")
}

//-------------------Swagger

func generateSwagger() error {
	return sh.RunV("swag", "init", "-g", "./main.go", "-o", "./docs", "--parseDependency")
}

func (Generate) Server() error {
	mg.SerialDeps(generateService, generateMain, moveHandler, generateMainService)
	ask()
	mg.SerialDeps(generateApp, formatFiles, moveFiles)
	mg.SerialDeps(generateSwagger)

	return nil
}
