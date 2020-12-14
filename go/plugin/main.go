package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

type GeneratorIface interface {
	Generate(g *protogen.GeneratedFile, file *protogen.File)
}

var plugins = []GeneratorIface{
	NewEggGenerator(),
	NewClientMockGenerator(),
	// Add any other generators here
}

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			filename := fmt.Sprintf("%s.pb.custom.go", f.GeneratedFilenamePrefix)
			g := plugin.NewGeneratedFile(filename, f.GoImportPath)

			for _, p := range plugins {
				p.Generate(g, f)
			}
		}

		return nil
	})
}
