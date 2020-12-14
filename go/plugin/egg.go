package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// EggGenerator generates a package header and an egg
type EggGenerator struct{}

var _ GeneratorIface = (*EggGenerator)(nil)

func NewEggGenerator() *EggGenerator {
	return &EggGenerator{}
}

func (aeg *EggGenerator) Generate(g *protogen.GeneratedFile, file *protogen.File) {
	g.P("package ", file.GoPackageName)
	g.P()
	g.P(`
//             ████      
//           ██░░░░██    
//         ██░░░░░░░░██  
//         ██░░░░░░░░██  
//       ██░░░░░░░░░░░░██
//       ██░░  ░░░░  ░░██
//       ██░░░░    ░░░░██
//         ██░░░░░░░░██  
//           ████▓▓██
`)
}
