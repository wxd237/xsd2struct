package xsd2struct

import (
	"bufio"
	//"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type simpleType struct {
	XmlName xml.Name
	Name    string `xml:"name,attr"`
}

func (t *simpleType) makepkg() {
	w := bufio.NewWriter(os.Stdout)
	line1 := fmt.Sprintf("type %s struct{\n", t.Name)
	w.WriteString(line1)
	line2 := fmt.Sprintf("\t%s %s  `xml:\"%s\"`\n", t.Name, "string", t.Name)
	w.WriteString(line2)
	line3 := fmt.Sprintf("}\n")
	w.WriteString(line3)
	w.Flush()
}

type enumeration struct {
	XmlName xml.Name
	Value   string `xml:"value,attr"`
}

type union struct {
}

type restriction struct {
	XmlName xml.Name
	Base    string `xml:"base,attr"`
}

type element struct {
	XmlName   xml.Name
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	MinOccurs string `xml:"MinOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

func (t *element) makepkg() {
	if t.Name == "" {
		return
	}
	w := bufio.NewWriter(os.Stdout)
	var line string
	if t.MaxOccurs != "unbounded" {
		line = fmt.Sprintf("\t%s %s  `xml:\"%s\"`\n", t.Name, t.Type, t.Name)
		//w.WriteString(line)
	} else {
		line = fmt.Sprintf("\t%s []%s  `xml:\"%s\"`\n", t.Name, t.Type, t.Name)
		//w.WriteString(line)
	}
	w.WriteString(line)
	w.Flush()

}

type attribute struct {
	XmlName xml.Name
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Use     string `xml:"use,attr"`
}

func (t *attribute) makepkg() {
	if t.Name == "" {
		return
	}
	w := bufio.NewWriter(os.Stdout)
	line := fmt.Sprintf("\t%s %s  `xml:\"%s\",attr`\n", t.Name, t.Type, t.Name)
	w.WriteString(line)
	w.Flush()

}

type attributeGroup struct {
	Attribute []attribute `xml:"attribute"`
}

type choice struct {
	Group   []group   `xml:"group"`
	Element []element `xml:"element"`
}

func (t *choice) makepkg() {

	for _, k := range t.Group {
		k.makepkg()
	}

	for _, k := range t.Element {
		k.makepkg()
	}

}

type complexContent struct {
	Extension extension `xml:"extension"`
}

func (t *complexContent) makepkg() {
	t.Extension.makepkg()
}

type extension struct {
	XmlName   xml.Name
	Base      string      `xml:"base,attr"`
	Name      string      `xml:"name,attr"`
	Choice    choice      `xml:"choice"`
	Sequence  sequence    `xml:"sequence"`
	Attribute []attribute `xml:"attribute"`
}

func (t *extension) makepkg() {
	if t.Base == "" {
		return
	}
	w := bufio.NewWriter(os.Stdout)
	line1 := fmt.Sprintf("\t%s\n", t.Base)
	w.WriteString(line1)
	w.Flush()
	t.Choice.makepkg()
	t.Sequence.makepkg()

	for _, k := range t.Attribute {
		k.makepkg()
	}
}

type sequence struct {
	MaxOccurs string    `xml:"maxOccurs,attr"`
	Element   []element `xml:"element"`
	Group     group     `xml:"group"`
	Choice    []choice  `xml:"choice"`
}

func (t *sequence) makepkg() {

	for _, k := range t.Element {
		k.makepkg()
	}

	t.Group.makepkg()

	for _, k := range t.Choice {
		k.makepkg()
	}

}

type all struct {
	Element []element `xml:"element"`
}

func (t *all) makepkg() {
	for _, k := range t.Element {
		k.makepkg()
	}
}

type complexType struct {
	XmlName        xml.Name
	Name           string         `xml:"name,attr"`
	Attribute      []attribute    `xml:"attribute"`
	Element        []element      `xml:"element"`
	All            all            `xml:"all"`
	Sequence       sequence       `xml:"sequence"`
	Choice         choice         `xml:"choice"`
	ComplexContent complexContent `xml:"complexContent"`
}

func (t *complexType) makepkg() {
	w := bufio.NewWriter(os.Stdout)
	line1 := fmt.Sprintf("type %s struct{\n", t.Name)
	w.WriteString(line1)
	w.Flush()

	t.Choice.makepkg()
	t.ComplexContent.makepkg()

	t.Sequence.makepkg()

	for _, k := range t.Element {
		k.makepkg()
	}

	w.Flush()

	for _, k := range t.Attribute {
		k.makepkg()
	}

	t.All.makepkg()

	line3 := fmt.Sprintf("}\n")
	w.WriteString(line3)
	w.Flush()
}

type group struct {
	XmlName   xml.Name
	Name      string `xml:"name,attr"`
	Ref       string `xml:"ref,attr"`
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

func (t *group) makepkg() {
	if t.Name == "" {
		return
	}
	w := bufio.NewWriter(os.Stdout)
	var line string
	if t.MaxOccurs != "unbounded" {

		line = fmt.Sprintf("\t%s %s  `xml:\"%s\"`\n", t.Name, t.Ref, t.Name)
		//w.WriteString(line)
	} else {
		line = fmt.Sprintf("\t%s []%s  `xml:\"%s\"`\n", t.Name, t.Ref, t.Name)
		//w.WriteString(line)
	}
	w.WriteString(line)
	w.Flush()

}

type schmea struct {
	XmlName     xml.Name
	ComplexType []complexType `xml:"complexType"`
	Element     []element     `xml:"element"`
}

func (t *schmea) makepkg() {
	fmt.Printf("type WordSchema  struct {\n")
	for _, k := range t.Element {
		k.makepkg()
	}
	fmt.Printf("}\n")

	for _, k := range t.ComplexType {
		k.makepkg()
	}
}

func (t *simpleType) makePkg() string {
	line := fmt.Sprintf("type %s string")
	return line
}

func NewXSDFile(filename string) schmea {
	var s schmea
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	xml.Unmarshal(b, &s)

	return s
}
