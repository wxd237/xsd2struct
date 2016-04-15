package ooxml
type WordSchema  struct {
	Properties string  `xml:"Properties"`
}
type CT_Properties struct{
	property []CT_Property  `xml:"property"`
}
type CT_Property struct{
	fmtid string  `xml:"fmtid",attr`
	pid string  `xml:"pid",attr`
	name string  `xml:"name",attr`
	linkTarget string  `xml:"linkTarget",attr`
}
