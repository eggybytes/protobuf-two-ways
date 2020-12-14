package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"protos/annotations"

	"github.com/eggybytes/protobuf-two-ways/go/helpers"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var needsClientMockImports = map[string]bool{}

// ClientMockGenerator generates a mock struct that satisfies a service's client interface
type ClientMockGenerator struct{}

var _ GeneratorIface = (*ClientMockGenerator)(nil)

func NewClientMockGenerator() *ClientMockGenerator {
	return &ClientMockGenerator{}
}

func (cmg *ClientMockGenerator) Generate(g *protogen.GeneratedFile, file *protogen.File) {
	for _, pb := range file.Proto.GetService() {
		if shouldGenerateClientMock(pb) {
			needsClientMockImports[file.GeneratedFilenamePrefix] = true
			cmg.generateClientMockFactory(g, pb)
			cmg.generateClientMockMethods(g, pb, file)
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
func shouldGenerateClientMock(pb *descriptorpb.ServiceDescriptorProto) bool {
	if pb.GetOptions() == nil {
		return false
	}

	if proto.GetExtension(pb.GetOptions(), annotations.E_ClientMock).(bool) {
		return true
	}

	return false
}

func (cmg *ClientMockGenerator) generateClientMockFactory(g *protogen.GeneratedFile, pb *descriptorpb.ServiceDescriptorProto) {
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

	serviceName := helpers.Camel(pb.GetName())
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

func (cmg *ClientMockGenerator) generateClientMockMethods(g *protogen.GeneratedFile, pb *descriptorpb.ServiceDescriptorProto, file *protogen.File) {
	methodTemplate := `
func (c *Mock{{.ServiceName}}Client) {{.MethodName}}(ctx context.Context, in *{{.RequestType}}, opts ...grpc.CallOption) (*{{.ResponseType}}, error) {
	args := c.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*{{.ResponseType}}), args.Error(1)
}
`

	for _, m := range pb.Method {
		tpl, err := template.New("method").Parse(methodTemplate)
		if err != nil {
			log.Fatal(err)
		}

		origMethodName := m.GetName()
		var buf bytes.Buffer
		err = tpl.Execute(&buf, struct {
			ServiceName  string
			MethodName   string
			RequestType  string
			ResponseType string
		}{
			ServiceName:  helpers.Camel(pb.GetName()),
			MethodName:   helpers.Camel(origMethodName),
			RequestType:  typeName(file.GoPackageName, m.GetInputType()),
			ResponseType: typeName(file.GoPackageName, m.GetOutputType()),
		})

		if err != nil {
			log.Fatal(err)
		}

		g.P(buf.String())
	}
}

// Since the types we're referring to are in the same package, we drop the leading `.` and the package name
func typeName(trimPackageName protogen.GoPackageName, typeName string) string {
	return strings.TrimLeft(typeName, fmt.Sprintf(".%s", trimPackageName))
}
