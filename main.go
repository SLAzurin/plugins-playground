package main

import (
	"errors"
	"log"
	"os"
	"plugin"

	"github.com/slazurin/plugins-playground/pkg"
)

func main() {
	binPath := "./plugins/bin/"
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		panic(err)
	}
	files, _ := os.ReadDir(binPath)
	for _, f := range files {
		cmd, err := plugin.Open(binPath + f.Name())
		if err != nil {
			panic(err)
		}
		symbol, err := cmd.Lookup("Fn")
		if err != nil {
			panic(err)
		}
		fn, ok := symbol.(*pkg.Fn)
		if !ok {
			panic(errors.New("cannot cast symbol to pkg.Fn for plugin " + f.Name()))
		}
		log.Println("Executing " + f.Name() + "...")
		(*fn)()
	}
}
