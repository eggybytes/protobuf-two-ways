package main

import (
	"bytes"
	"log"
	"text/template"

	"protos/annotations"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

var needsClientMockImports = map[string]bool{}

// ClientMockGenerator generates a mock struct that satisfies a service's client interface
type ClientMockGenerator struct{}

var _ GeneratorIface = (*ClientMockGenerator)(nil)

func NewClientMockGenerator() *ClientMockGenerator {
	return &ClientMockGenerator{}
}

func (cmg *ClientMockGenerator) Generate(g *protogen.GeneratedFile, file *protogen.File) {
	for _, svc := range file.Services {
		if shouldGenerateClientMock(svc) {
			needsClientMockImports[file.GeneratedFilenamePrefix] = true
			cmg.generateClientMockFactory(g, svc)
			cmg.generateClientMockMethods(g, svc, file)
		}
	}

	if needsClientMockImports[file.GeneratedFilenamePrefix] {
		g.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "context",
			GoImportPath: "context",
		})
		g.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "mock",
			GoImportPath: "github.com/stretchr/testify/mock",
		})
		g.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "grpc",
			GoImportPath: "google.golang.org/grpc",
		})
	}
}

// shouldGenerateClientMock determines whether or not the message is tagged with
// (client_mock=true)
func shouldGenerateClientMock(svc *protogen.Service) bool {
	opts := svc.Desc.Options()
	if opts == nil {
		return false
	}

	if proto.GetExtension(opts, annotations.E_ClientMock).(bool) {
		return true
	}

	return false
}

func (cmg *ClientMockGenerator) generateClientMockFactory(g *protogen.GeneratedFile, svc *protogen.Service) {
	factoryTemplate := `
// Mock{{.ServiceName}}Client is a mock {{.ServiceName}}Client which
// satisfies the {{.ServiceName}}Client interface.
type Mock{{.ServiceName}}Client struct {
	mock.Mock
}

func NewMock{{.ServiceName}}Client() *Mock{{.ServiceName}}Client {
	return &Mock{{.ServiceName}}Client{}
}
`
	tpl, err := template.New("factory").Parse(factoryTemplate)
	if err != nil {
		log.Fatal(err)
	}

	serviceName := svc.GoName
	var buf bytes.Buffer
	err = tpl.Execute(&buf, struct {
		ServiceName string
	}{
		ServiceName: serviceName,
	})

	if err != nil {
		log.Fatal(err)
	}

	g.P(buf.String())
}

func (cmg *ClientMockGenerator) generateClientMockMethods(g *protogen.GeneratedFile, svc *protogen.Service, file *protogen.File) {
	methodTemplate := `
func (c *Mock{{.ServiceName}}Client) {{.MethodName}}(ctx context.Context, in *{{.RequestType}}, opts ...grpc.CallOption) (*{{.ResponseType}}, error) {
	args := c.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*{{.ResponseType}}), args.Error(1)
}
`

	for _, m := range svc.Methods {
		tpl, err := template.New("method").Parse(methodTemplate)
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		err = tpl.Execute(&buf, struct {
			ServiceName  string
			MethodName   string
			RequestType  string
			ResponseType string
		}{
			ServiceName:  svc.GoName,
			MethodName:   m.GoName,
			RequestType:  m.Input.GoIdent.GoName,
			ResponseType: m.Output.GoIdent.GoName,
		})

		if err != nil {
			log.Fatal(err)
		}

		g.P(buf.String())
	}
}
