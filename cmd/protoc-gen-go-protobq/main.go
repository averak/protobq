package main

import (
	"github.com/averak/protobq/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			g := internal.NewCodeGenerator(plugin, file)
			if err := g.Gen(); err != nil {
				return err
			}
		}
		return nil
	})
}
