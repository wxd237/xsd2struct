package ooxml
type WordSchema  struct {
	iscomment CT_Empty  `xml:"iscomment"`
	textdata CT_Rel  `xml:"textdata"`
}
type CT_Empty struct{
}
type CT_Rel struct{
	id string  `xml:"id",attr`
}
