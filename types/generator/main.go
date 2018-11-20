package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type MyMsg struct {
	PackageOrig        string
	SubfolderOrig      string
	MsgNameOrig        string
	PackageCamelCase   string
	SubfolderCamelCase string
	MsgNameCamelCase   string
	Typename           string
}

func replaceUnderscoreAndCamelCase(input string) string {
	splittedString := strings.Split(input, "_")
	retValue := ""

	for _, splitted := range splittedString {
		retValue = retValue + strings.Title(splitted)
	}
	return retValue
}

func getFileAsString(myMsg *MyMsg) string {

	camelCase := strings.Title(myMsg.PackageOrig) + strings.Title(myMsg.MsgNameOrig)

	myMsg.Typename = camelCase
	myMsg.PackageCamelCase = replaceUnderscoreAndCamelCase(myMsg.PackageOrig)
	myMsg.SubfolderCamelCase = replaceUnderscoreAndCamelCase(myMsg.SubfolderOrig)
	myMsg.MsgNameCamelCase = replaceUnderscoreAndCamelCase(myMsg.MsgNameOrig)

	retString :=
		`package types
// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -l{{.PackageOrig}}__rosidl_generator_c -l{{.PackageOrig}}__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type {{.PackageCamelCase}}{{.MsgNameCamelCase}} struct {
	data    *C.{{.PackageOrig}}__{{.SubfolderOrig}}__{{.MsgNameCamelCase}}
	MsgType MessageTypeSupport
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) InitMessage() {
	msg.data = C.init_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}()
	msg.MsgType = GetMessageTypeFrom{{.PackageCamelCase}}{{.MsgNameCamelCase}}()
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) Set{{.MsgNameCamelCase}}(data {{.MsgNameCamelCase}}) {
	//TODO: to implement the setter
	//msg.data.data = whatever
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) Get{{.MsgNameCamelCase}}() {{.MsgNameCamelCase}} {
	return {{.MsgNameCamelCase}}(msg.data.data)
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) DestroyMessage() {
	C.destroy_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}(msg.data)
}

func GetMessageTypeFrom{{.PackageCamelCase}}{{.MsgNameCamelCase}}() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}()}
}
`
	retValue2 := strings.Trim(retString, "\t")
	return retValue2
}

func main() {
	t := template.New("MyGoROS2Package")

	msgs := []MyMsg{
		{"std_msgs", "msg", "string", "", "", "", ""},
	}
	text := getFileAsString(&msgs[0])

	t.Parse(text)

	for _, msg := range msgs {
		f, err := os.OpenFile("../msg_"+msg.PackageOrig+"_"+msg.MsgNameOrig+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(f, msg)

		defer f.Close()
	}
}
