package xsd2struct

import (
	"fmt"
	"testing"
)

var s schmea

func TestNewXSDFile(t *testing.T) {

	s := NewXSDFile("ref/wml.xsd")
	fmt.Printf("%+v\n", s)
	s.makepkg()
}

func TestDump(t *testing.T) {
	//	var e element
	//	e.Name = "fafdadf"
	//	e.MaxOccurs = "2"
	//	e.Type = "CT_ELE"
	//	e.makepkg()
}
