//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

func deleteFiles(pattern string) error {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Failed to delete file: %s, Error: %v\n", file, err)
		} else {
			fmt.Printf("Deleted file: %s\n", file)
		}
	}

	return nil
}

// Clean removes the specified files and directories.
func Clean() {
	sh.RunV("rm", "./maker/generator/main.go")
	sh.RunV("rm", "./maker/service/handler.yaml")
	deleteFiles("client_*.go")
	deleteFiles("handler_*.go")
	deleteFiles("router_*.go")
	deleteFiles("dto_*.go")
	sh.RunV("rm", "main.go")
	sh.RunV("rm", "router.go")
}
