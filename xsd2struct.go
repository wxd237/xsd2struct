package xsd2struct

import (
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
)

type simpleType struct {
	XmlName xml.Name
	Name   string  `xml:"name,attr"`
}

type restriction struct{
	XmlName xml.Name
	Base    string  `xml:"base,attr"`
}

type  element struct{
		XmlName xml.Name
		Name 			string	`xml:"name,attr"`
		Type			string	`xml:"type,attr"`
		MinOccurs		string	`xml:"minOccurs,attr"`
		MaxOccurs		string	`xml:"maxOccurs,attr"`
		
}

type attribute struct{
		XmlName xml.Name
		Name 			string	`xml:"name,attr"`
		Type			string	`xml:"type,attr"`
		Use				string	`xml:"use,attr"`	
}	

type  complexType struct{
	XmlName xml.Name
	Name 			string			`xml:"name,attr"`
	Attributettribute []attribute   		`xml:"attribute"`
	Element   []element       		`xml:"element"`
}

func(t *complexType) makepkg(){
	
}

type extension struct{
		XmlName xml.Name
		Base    string  `xml:"base,attr"`
}


type group struct{
	XmlName xml.Name
	Ref      		string	`xml:"ref,attr"`
	MinOccurs		string	`xml:"minOccurs,attr"`
	MaxOccurs		string	`xml:"maxOccurs,attr"`   	
}


type schmea struct{
	XmlName xml.Name
	ComplexType	[]complexType	`xml:"complexType"`
	Element     []element 		`xml:"element"`
}



func(t *schmea) makepkg(){
	fmt.Printf("schmea:%s\n",t.XmlName.Local)
	for i,k:=range t.ComplexType {
		fmt.Printf("schmea:%s\n",t.XmlName.Local)
	}
}

func (t *simpleType) makePkg() string {
	line := fmt.Sprintf("type %s string")
	return line
}



func NewXSDFile(filename string) schmea {
	var s schmea
	b,err:=ioutil.ReadFile(filename)
	if err!=nil{
		log.Fatal(err)
	}
	xml.Unmarshal(b,&s)
	
	return s
}
