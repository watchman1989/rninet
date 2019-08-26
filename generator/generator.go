package generator

import (
	"bytes"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/watchman1989/rninet/common/global"
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
		meta: new(Meta),
	}
	SERVER_DIR_LIST = []string {"proto", "handler", "router"}
	CLIENT_DIR_LIST = []string {}
)

type Meta struct {
	Service *proto.Service
	Messages []*proto.Message
	Rpcs []*proto.RPC
	Package *proto.Package
	Rpath string
	Fpath string
}


type RpcMeta struct {
	Rpc *proto.RPC
	Package *proto.Package
	Rpath string
	Fpath string
}

type Generator struct {
	options *Options
	meta *Meta
}


func NewGenerator (opts ...Option) *Generator {
	generator.options = &Options{}
	for _, opt := range opts {
		opt(generator.options)
	}

	_ = generator.ParseProto()
	_ = generator.GetPathInfo()

	//fmt.Println(generator.meta)

	return generator
}


func (g *Generator) Gen () {

	_ = g.GenerateDir()
	_ = g.GenerateGrpc()
	_ = g.GenerateServer()
	_ = g.GenerateRouter()
	_ = g.GenerateHandler()
}


func (g *Generator) GetPathInfo () error {

	srcPath := filepath.Join(os.Getenv("GOPATH"), "src") + global.SLASH
	fmt.Printf("++++++++++++SRC_PATH: %s\n", srcPath)
	if strings.HasPrefix(g.options.Output, srcPath) {
		g.meta.Rpath = strings.Replace(strings.Replace(g.options.Output, srcPath, "", 1), "\\", "/", -1)
		g.meta.Fpath = g.options.Output
	} else {
		g.meta.Rpath = strings.Replace(g.options.Output, "\\", "/", -1)
		g.meta.Fpath = filepath.Join(srcPath, g.options.Output)
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
	g.meta.Service = s
}


func (g *Generator)handleProtoMessage (m *proto.Message) {
	fmt.Printf("proto.Message: %s\n", m.Name)
	g.meta.Messages = append(g.meta.Messages, m)
}


func (g *Generator) handleProtoRPC (r *proto.RPC) {
	fmt.Printf("proto.RPC: %s\n", r.Name)
	g.meta.Rpcs = append(g.meta.Rpcs, r)
}

func (g *Generator) handleProtoPackage (r *proto.Package) {
	fmt.Printf("proto.Package: %s\n", r.Name)
	g.meta.Package = r
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
		genPath := filepath.Join(g.meta.Fpath, dir)

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

	protoPath := filepath.Join(g.meta.Fpath, "proto", g.meta.Package.Name)
	if err := os.MkdirAll(protoPath, 0775); err != nil {
		fmt.Printf("MKDIR %s ERROR: %v\n", err)
		return err
	}

	commandLine := fmt.Sprintf(PROTOC_GRPC_COMMAND, filepath.Dir(g.options.ProtoFile), protoPath, g.options.ProtoFile)
	command := strings.Split(commandLine, " ")[0]
	args := strings.Split(commandLine, " ")[1:]
	//fmt.Printf("COMMAND: %s, ARGS: %v\n", command, args)
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


func (g *Generator) GenerateServer () error {

	fmt.Printf("Generate server code\n")

	srvFile := filepath.Join(g.meta.Fpath, "server.go")
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

	if err = t.Execute(fp, g.meta); err != nil {
		fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
		return err
	}

	return nil
}


func (g *Generator) GenerateHandler () error {

	fmt.Printf("Generate handler code\n")

	for _, rpc := range g.meta.Rpcs {

		hdlFile := filepath.Join(g.meta.Fpath, "handler", fmt.Sprintf("handle_%s.go", strings.ToLower(rpc.Name)))
		fp, err := os.OpenFile(hdlFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Printf("OPEN_FILE %s ERROR: %v\n", hdlFile, err)
			continue
		}
		defer fp.Close()

		rpcMeta := &RpcMeta{}
		rpcMeta.Rpc = rpc
		rpcMeta.Package = g.meta.Package
		rpcMeta.Fpath = g.meta.Fpath
		rpcMeta.Rpath = g.meta.Rpath

		t := template.New(fmt.Sprintf("handle_%s.go", rpc.Name))
		t, err = t.Parse(handlerTemplate)
		if err != nil {
			fmt.Printf("TEMPLATE_PARSE_ERROR: %v\n", err)
			return err
		}

		if err = t.Execute(fp, rpcMeta); err != nil {
			fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
			return err
		}

	}

	return nil
}


func (g *Generator) GenerateRouter () error {

	fmt.Printf("Generate router code\n")

	rtFile := filepath.Join(g.meta.Fpath, "router", "router.go")
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

	if err = t.Execute(fp, g.meta); err != nil {
		fmt.Printf("TEMPLATE_EXECUTE_ERROR: %v\n", err)
		return err
	}

	return nil
}
