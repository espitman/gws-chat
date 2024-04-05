//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Add mg.Namespace
type Generate mg.Namespace

func Run() error {
	return sh.RunV("go", "run", ".")
}
