package xsd2struct

import (
	"bytes"
	"os"
	"strings"
	//"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type simpleType struct {
	XmlName xml.Name
	Name    string `xml:"name,attr"`
}

func (t *simpleType) makepkg() string {

	var w bytes.Buffer
	line1 := fmt.Sprintf("type %s struct{\n", t.Name)
	w.WriteString(line1)
	line2 := fmt.Sprintf("\t%s %s  `xml:\"%s\"`\n", t.Name, "string", t.Name)
	w.WriteString(line2)
	line3 := fmt.Sprintf("}\n")
	w.WriteString(line3)
	return w.String()
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

func (t *element) makepkg() string {
	if t.Name == "" {
		return ""
	}
	var w bytes.Buffer
	var line string
	Curtype := t.Type
	if strings.ContainsAny(Curtype, "xsd:") {
		Curtype = "string"
	}
	if strings.ContainsAny(Curtype, "s:") {
		Curtype = strings.Replace(Curtype, "s:", "", 1)
	}
	if t.MaxOccurs != "unbounded" {
		line = fmt.Sprintf("\t%s %s  `xml:\"%s\"`\n", strings.Title(t.Name), Curtype, t.Name)
		//w.WriteString(line)
	} else {
		line = fmt.Sprintf("\t%s []%s  `xml:\"%s\"`\n", strings.Title(t.Name), Curtype, t.Name)
		//w.WriteString(line)
	}
	w.WriteString(line)
	return w.String()

}

type attribute struct {
	XmlName xml.Name
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Use     string `xml:"use,attr"`
}

func (t *attribute) makepkg() string {
	if t.Name == "" {
		return ""
	}
	Curtype := t.Type
	if strings.ContainsAny(Curtype, "xsd:") {
		Curtype = "string"
	}
	if strings.ContainsAny(Curtype, "s:") {
		Curtype = strings.Replace(Curtype, "s:", "", 1)
	}

	var w bytes.Buffer
	line := fmt.Sprintf("\t%s %s  `xml:\"%s\",attr`\n", strings.Title(t.Name), Curtype, t.Name)
	w.WriteString(line)
	return w.String()

}

type attributeGroup struct {
	Attribute []attribute `xml:"attribute"`
}

type choice struct {
	Group   []groupb  `xml:"group"`
	Element []element `xml:"element"`
}

func (t *choice) makepkg() string {
	var w bytes.Buffer
	for _, k := range t.Group {
		w.WriteString(k.makepkg())
	}

	for _, k := range t.Element {
		w.WriteString(k.makepkg())

	}
	return w.String()
}

type complexContent struct {
	Extension extension `xml:"extension"`
}

func (t *complexContent) makepkg() string {
	return t.Extension.makepkg()
}

type extension struct {
	XmlName   xml.Name
	Base      string      `xml:"base,attr"`
	Name      string      `xml:"name,attr"`
	Choice    choice      `xml:"choice"`
	Sequence  sequence    `xml:"sequence"`
	Attribute []attribute `xml:"attribute"`
}

func (t *extension) makepkg() string {
	if t.Base == "" {
		return ""
	}
	var w bytes.Buffer
	line1 := fmt.Sprintf("\t%s\n", t.Base)
	w.WriteString(line1)

	w.WriteString(t.Choice.makepkg())

	w.WriteString(t.Sequence.makepkg())

	for _, k := range t.Attribute {
		w.WriteString(k.makepkg())
	}
	return w.String()
}

type sequence struct {
	MaxOccurs string    `xml:"maxOccurs,attr"`
	Element   []element `xml:"element"`
	Group     []groupb  `xml:"group"`
	Choice    []choice  `xml:"choice"`
}

func (t *sequence) makepkg() string {
	var w bytes.Buffer
	for _, k := range t.Element {
		w.WriteString(k.makepkg())
	}

	for _, k := range t.Group {
		w.WriteString(k.makepkg())
	}

	for _, k := range t.Choice {
		w.WriteString(k.makepkg())
	}
	return w.String()
}

type all struct {
	Element []element `xml:"element"`
}

func (t *all) makepkg() string {
	var w bytes.Buffer
	for _, k := range t.Element {
		w.WriteString(k.makepkg())
	}
	return w.String()
}

type complexType struct {
	XmlName        xml.Name
	Group          []groupb
	Name           string         `xml:"name,attr"`
	Attribute      []attribute    `xml:"attribute"`
	Element        []element      `xml:"element"`
	All            all            `xml:"all"`
	Sequence       sequence       `xml:"sequence"`
	Choice         choice         `xml:"choice"`
	ComplexContent complexContent `xml:"complexContent"`
}

func (t *complexType) makepkg() string {
	var w bytes.Buffer
	line1 := fmt.Sprintf("type %s struct{\n", t.Name)
	w.WriteString(line1)

	w.WriteString(t.Choice.makepkg())
	w.WriteString(t.ComplexContent.makepkg())

	w.WriteString(t.Sequence.makepkg())

	for _, k := range t.Element {
		w.WriteString(k.makepkg())

	}

	for _, k := range t.Attribute {
		w.WriteString(k.makepkg())
	}

	w.WriteString(t.All.makepkg())

	line3 := fmt.Sprintf("}\n")
	w.WriteString(line3)
	return w.String()
}

type group struct {
	XmlName  xml.Name
	Name     string   `xml:"name,attr"`
	Groupb   groupb   `xml:"group"`
	Sequence sequence `xml:"sequence"`
	Choice   choice   `xml:"choice"`
}

func (t *group) makepkg() string {

	if t.Name == "" {
		return ""
	}
	var w bytes.Buffer
	var line string

	line1 := fmt.Sprintf("type %s struct{\n", t.Name)
	w.WriteString(line1)
	w.WriteString(t.Sequence.makepkg())
	w.WriteString(t.Groupb.makepkg())
	w.WriteString(t.Choice.makepkg())
	//fname := strings.Replace(t.Name, "EG_", "", 1)
	//line = fmt.Sprintf("\t%s %s  \n", fname, t.Name)

	w.WriteString(line)

	w.WriteString("}\n")
	return w.String()

}

type groupb struct {
	XmlName   xml.Name
	Ref       string `xml:"ref,attr"`
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

func (t *groupb) makepkg() string {

	if t.Ref == "" {
		return ""
	}
	var w bytes.Buffer
	var line string
	fname := strings.Replace(t.Ref, "EG_", "", 1)
	if t.MaxOccurs != "unbounded" {

		line = fmt.Sprintf("\t%s %s  \n", fname, t.Ref)

	} else {
		line = fmt.Sprintf("\t%s []%s \n", fname, t.Ref)

	}
	w.WriteString(line)
	return w.String()

}

type schmea struct {
	Prefix      string
	XmlName     xml.Name
	ComplexType []complexType `xml:"complexType"`
	Element     []element     `xml:"element"`
	SimpleType  []simpleType  `xml:"simpleType"`
	Group       []group       `xml:"group"`
}

func (t *schmea) makepkg() string {
	var w bytes.Buffer
	line1 := fmt.Sprintf("type %sSchema  struct {\n", t.Prefix)
	w.WriteString(line1)
	for _, k := range t.Element {
		w.WriteString(k.makepkg())

	}

	w.WriteString("}\n")

	for _, k := range t.Group {
		w.WriteString(k.makepkg())
	}

	for _, k := range t.ComplexType {
		w.WriteString(k.makepkg())
	}

	for _, k := range t.SimpleType {
		w.WriteString(k.makepkg())
	}

	return w.String()
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

func ParseTo(infile, schmeaname string) {
	s := NewXSDFile(infile)
	s.Prefix = strings.Title(schmeaname)
	rets := s.makepkg()
	log.Printf("cur hand %s:%s \n", infile, schmeaname)
	w, _ := os.Create("refgo/" + s.Prefix + ".go")
	defer w.Close()

	w.WriteString("package ooxml\n")

	w.WriteString(rets)
	w.Close()
}
