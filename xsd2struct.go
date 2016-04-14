package xsd2struct

import (
	"fmt"
	"log"

	"github.com/jteeuwen/xmlx"
)

type simpleType struct {
	name string
}

func (t *simpleType) makePkg() string {
	line := fmt.Sprintf("type %s string", t.name)
	return line
}

type XSDFile struct {
	filename string
	doc      *xmlx.Document
}

func NewXSDFile(filename string) *XSDFile {
	var x XSDFile
	x.doc = xmlx.New()

	if err := x.doc.LoadFile("./ref/wml.xsd", nil); err != nil {
		log.Fatal(err)

	}
	if len(x.doc.Root.Children) == 0 {
		log.Fatalf("Root node has no children.")
	}
	return &x
}

func (this *XSDFile) ParsesimpleType() map[simpleType]string {
	ret := make(map[simpleType]string)

	snode := this.doc.SelectNodes("*", "simpleType")
	for _, k := range snode {
		var s simpleType
		s.name = k.Attributes[0].Value
		ret[s] = "string"
		fmt.Printf("simpleType:%s \n", k.Attributes)
	}
	return ret
}
