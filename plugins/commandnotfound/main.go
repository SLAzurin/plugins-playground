package main

import (
	"log"

	"github.com/slazurin/plugins-playground/pkg"
)

var Fn pkg.Fn = func() {
	log.Println("command not found")
}
