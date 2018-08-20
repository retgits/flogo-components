// +build example
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jhump/protoreflect/desc/protoparse"
)

const methodTemplate = `func (s *server) {{.methodname}}(ctx context.Context, in *{{.inputmessage}}) (*{{.outputmessage}}, error) {
	// Transform the incoming struct to a map and assign it to a map
	message := structs.Map(in)
	triggerData := make(map[string]interface{})
	triggerData["message"] = message

	// Find the handler for the flow
	handler := handlerMap["{{.servicename}}-{{.methodname}}"]

	// Execute the flow
	out, err := handler.Handle(context.Background(), triggerData)
	if err != nil {
		log.Infof("Error: %s", err.Error())
	}

	// Create the output
	output := &{{.outputmessage}}{}
	err = fillStruct(out["data"].Value().(map[string]interface{}), output)
	if err != nil {
		return nil, err
	}

	return output, nil
}`

const serverFileTemplate = `// Package grpctrigger implements a trigger to receive messages over gRPC.
package grpctrigger

import (
	"context"

	"github.com/fatih/structs"
	"google.golang.org/grpc"
)

// Method to register the gRPC server methods
func registerServerMethods(s *grpc.Server) {
	{{.servermethods}}
}

// Methods from the protobuf
{{.protomethods}}
`

func main() {
	fmt.Println("Running build script for the gRPC trigger")

	// appdir is the directory where the app is stored. For example if you app is called
	// lambda this would be <path>/lambda/src/lambda
	appDir := os.Args[1]
	fmt.Printf("appDir has been set to: %s\n", appDir)

	// Variables
	grpcTriggerLocation := filepath.Join(appDir, "vendor", "github.com", "retgits", "flogo-components", "trigger", "grpctrigger")
	flogoJSON := filepath.Join(appDir, "..", "..", "flogo.json")
	var protoFileLocation, protoFileName string
	fmt.Printf("gRPC trigger location has been set to: %s\n", grpcTriggerLocation)
	fmt.Printf("flogo.json has been set to: %s\n", flogoJSON)

	// Read the flogo.json file to find the protoFileLocation
	input, err := ioutil.ReadFile(flogoJSON)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	triggers := data["triggers"].([]interface{})
	for _, trigger := range triggers {
		trigger := trigger.(map[string]interface{})
		if trigger["ref"].(string) == "github.com/retgits/flogo-components/trigger/grpctrigger" {
			settings := trigger["settings"].(map[string]interface{})
			protoFileLocation, protoFileName = path.Split(settings["protofileLocation"].(string))
		}
	}

	fmt.Printf("proto file location has been set to: %s\n", protoFileLocation)
	fmt.Printf("proto filename has been set to: %s\n", protoFileName)

	// Run protoc to generate the protobuffer file
	args := []string{
		"-I",
		protoFileLocation,
		fmt.Sprintf("%s/%s", protoFileLocation, protoFileName),
		fmt.Sprintf("--go_out=plugins=grpc:%s", grpcTriggerLocation),
	}
	cmd := exec.Command("protoc", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	if len(string(out)) > 3 {
		fmt.Printf("protoc result: %s\n", string(out))
	}

	// Parse the proto file
	p := &protoparse.Parser{}
	fds, err := p.ParseFiles(filepath.Join(protoFileLocation, protoFileName))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// For now the trigger only supports one proto file
	protoFile := fds[0]

	// Override the package name in the proto file. This has no consequences on the server or the client, but will make the code easier
	fmt.Printf("Override generated package name to [grpctrigger]\n")
	input, err = ioutil.ReadFile(filepath.Join(grpcTriggerLocation, strings.Replace(protoFileName, ".proto", ".pb.go", 1)))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	output := strings.Replace(string(input), fmt.Sprintf("package %s", protoFile.GetPackage()), "package grpctrigger", 1)
	err = ioutil.WriteFile(filepath.Join(grpcTriggerLocation, strings.Replace(protoFileName, ".proto", ".pb.go", 1)), []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// The methods that need to be registered for the server to work
	var methodsToRegister string

	// The functions that need to be present for the server to work
	var functionsToRegister string

	// Find all services and methods
	for _, service := range protoFile.GetServices() {
		methodsToRegister = fmt.Sprintf("%s\nRegister%sServer(s, &server{})\n", methodsToRegister, service.GetName())
		for _, method := range service.GetMethods() {
			t := template.New("top")
			data := make(map[string]interface{})
			data["servicename"] = service.GetName()
			data["methodname"] = method.GetName()
			data["inputmessage"] = method.GetInputType().GetName()
			data["outputmessage"] = method.GetOutputType().GetName()
			template.Must(t.Parse(methodTemplate))
			buf := &bytes.Buffer{}
			t.Execute(buf, data)
			functionsToRegister = fmt.Sprintf("%s\n%s\n", functionsToRegister, buf.String())
		}
	}

	// Generate the contents of the registration.go file
	t := template.New("top")
	data = make(map[string]interface{})
	data["servermethods"] = html.UnescapeString(methodsToRegister)
	data["protomethods"] = html.UnescapeString(functionsToRegister)
	template.Must(t.Parse(serverFileTemplate))
	buf := &bytes.Buffer{}
	t.Execute(buf, data)

	// Write the new registration.go file
	err = ioutil.WriteFile(filepath.Join(grpcTriggerLocation, "registration.go"), buf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Okay, this is kinda weird but it seems the build.go is copied too so we have to remove it
	// TODO: Figure out why the build.go is copied, because that shouldn't be the case
	os.Remove(filepath.Join(appDir, "build.go"))
}
