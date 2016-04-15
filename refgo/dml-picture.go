package ooxml
type WordSchema  struct {
	pic CT_Picture  `xml:"pic"`
}
type CT_PictureNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvPicPr string  `xml:"cNvPicPr"`
}
type CT_Picture struct{
	nvPicPr string  `xml:"nvPicPr"`
	blipFill string  `xml:"blipFill"`
	spPr string  `xml:"spPr"`
}
