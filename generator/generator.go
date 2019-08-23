package generator

import (
	"bytes"
	"fmt"
	"github.com/emicklei/proto"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	PROTOC_GRPC_COMMAND = "protoc --proto_path=%s --go_out=plugins=grpc:%s %s"
)

var (
	generator *Generator = &Generator{
		protoInfo: new(ProtoInfo),
		baseInfo: new(BaseInfo),
	}
	SERVER_DIR_LIST = []string {"proto", "handler", "router"}
	CLIENT_DIR_LIST = []string {}
)

type ProtoInfo struct {
	Service *proto.Service
	Messages []*proto.Message
	Rpcs []*proto.RPC
	Package *proto.Package
}

type BaseInfo struct {
	Gopath string
	Rpath string
	Fpath string
}


type Generator struct {
	options *Options
	protoInfo *ProtoInfo
	baseInfo *BaseInfo
}


func NewGenerator (opts ...Option) *Generator {
	generator.options = &Options{}
	for _, opt := range opts {
		opt(generator.options)
	}

	generator.ParseProto()
	generator.GetBaseInfo()

	fmt.Println(generator.protoInfo)
	fmt.Println(generator.baseInfo)

	return generator
}


func (g *Generator) Gen () {

	g.GenerateDir()
	g.GenerateGrpc()
	g.GenerateServer()
	g.GenerateRouter()
	g.GenerateHandler()
}


func (g *Generator) GenerateDir () error {

	fmt.Printf("Generator dirs\n")

	var (
		dirs []string
	)

	if g.options.SrvFlag {
		dirs = SERVER_DIR_LIST
	}

	if g.options.CliFlag {
		dirs = CLIENT_DIR_LIST
	}

	for _, dir := range dirs {
		genPath := filepath.Join(g.baseInfo.Fpath, dir)

		fmt.Printf("MKDIR: %s\n", genPath)

		if err := os.MkdirAll(genPath, 0775); err != nil {
			fmt.Printf("MKDIR %s ERROR: %v\n", genPath, err)
			continue
		}
	}

	return nil
}


func (g *Generator) GenerateGrpc () error {

	fmt.Printf("Generate grpc code\n")

	if _, err := os.Stat(g.options.ProtoFile); err != nil {
		fmt.Printf("PROTO_FILE_ERROR: %v\n", err)
		return err
	}

	protoPath := filepath.Join(g.baseInfo.Fpath, "proto", g.protoInfo.Package)
	if err := os.MkdirAll(protoPath, 0775); err != nil {
		fmt.Printf("MKDIR %s ERROR: %v\n", err)
		return err
	}

	commandLine := fmt.Sprintf(PROTOC_GRPC_COMMAND, filepath.Dir(g.options.ProtoFile), protoPath, g.options.ProtoFile)
	command := strings.Split(commandLine, " ")[0]
	args := strings.Split(commandLine, " ")[1:]
	fmt.Printf("COMMAND: %s, ARGS: %v\n", command, args)
	fmt.Printf("COMMAND_LINE: %s\n", commandLine)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmdl := exec.Command(command, args...)

	cmdl.Stdout = &stdout
	cmdl.Stderr = &stderr

	if err := cmdl.Run(); err != nil {
		fmt.Printf("CMD_RUN_ERROR: %v\n", err)
		return err
	}

	fmt.Printf("%s\n", stdout.String())
	fmt.Printf("%s\n", stderr.String())

	return nil
}


func (g *Generator) GetBaseInfo () error {

	g.baseInfo.Gopath = filepath.Join(os.Getenv("GOPATH"), "src")
	if strings.HasPrefix(g.options.Output, g.baseInfo.Gopath) {
		g.baseInfo.Rpath = strings.Replace(g.options.Output, g.baseInfo.Gopath, "", 1)
		g.baseInfo.Rpath = g.options.Output
	} else {
		g.baseInfo.Rpath = filepath.Join(g.baseInfo.Gopath, g.options.Output)
	}

	return nil
}


func (g *Generator) ParseProto () error {

	fmt.Printf("Parse proto file\n")

	if _, err := os.Stat(g.options.ProtoFile); err != nil {
		fmt.Printf("PROTO_FILE_ERROR: %v\n", err)
		return err
	}

	protoReader, _ := os.Open(g.options.ProtoFile)
	defer protoReader.Close()

	protoParser := proto.NewParser(protoReader)
	definition, _ := protoParser.Parse()

	proto.Walk(definition,
		proto.WithService(g.handleProtoService),
		proto.WithMessage(g.handleProtoMessage),
		proto.WithRPC(g.handleProtoRPC),
		proto.WithPackage(g.handleProtoPackage),
	)

	return nil
}


func (g *Generator)handleProtoService (s *proto.Service) {
	fmt.Printf("proto.Service: %s\n", s.Name)
	g.protoInfo.Service = s
}


func (g *Generator)handleProtoMessage (m *proto.Message) {
	fmt.Printf("proto.Message: %s\n", m.Name)
	g.protoInfo.Messages = append(g.protoInfo.Messages, m)
}


func (g *Generator) handleProtoRPC (r *proto.RPC) {
	fmt.Printf("proto.RPC: %s\n", r.Name)
	g.protoInfo.Rpcs = append(g.protoInfo.Rpcs, r)
}

func (g *Generator) handleProtoPackage (r *proto.Package) {
	fmt.Printf("proto.Package: %s\n", r.Name)
	g.protoInfo.Package = r
}



func (g *Generator) GenerateServer () error {

	fmt.Printf("Generate server code\n")

	srvFile := filepath.Join(g.options.Output, "server.go")
	fp, err := os.OpenFile(srvFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("OPEN_FILE %s ERROR: %v\n", srvFile, err)
		return err
	}
	defer fp.Close()


	t := template.New("server")
	t, err = t.Parse(serverTemplate)
	if err != nil {
		fmt.Printf("TEMPLATE_PARSE_ERROR: %v\n", err)
		return err
	}

	if err = t.Execute(fp, g.protoInfo); err != nil {
		fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
		return err
	}

	return nil
}


func (g *Generator) GenerateHandler () error {

	fmt.Printf("Generate handler code\n")

	hdlFile := filepath.Join(g.options.Output, "handler", "handler.go")
	fp, err := os.OpenFile(hdlFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("OPEN_FILE %s ERROR: %v\n", hdlFile, err)
		return err
	}
	defer fp.Close()


	t := template.New("handler")
	t, err = t.Parse(handlerTemplate)
	if err != nil {
		fmt.Printf("TEMPLATE_PARSE_ERROR: %v\n", err)
		return err
	}

	if err = t.Execute(fp, g.protoInfo); err != nil {
		fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
		return err
	}


	return nil
}


func (g *Generator) GenerateRouter () error {

	fmt.Printf("Generate router code\n")

	rtFile := filepath.Join(g.options.Output, "router", "router.go")
	fp, err := os.OpenFile(rtFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("OPEN_FILE %s ERROR: %v\n", rtFile, err)
		return err
	}
	defer fp.Close()


	t := template.New("router")
	t, err = t.Parse(routerTemplate)
	if err != nil {
		fmt.Printf("TEMPLATE_PARSE_ERROR: %v\n", err)
		return err
	}

	if err = t.Execute(fp, g.protoInfo); err != nil {
		fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
		return err
	}

	return nil
}





