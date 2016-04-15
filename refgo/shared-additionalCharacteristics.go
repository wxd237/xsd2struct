package ooxml
type WordSchema  struct {
	additionalCharacteristics string  `xml:"additionalCharacteristics"`
}
type CT_AdditionalCharacteristics struct{
	characteristic []string  `xml:"characteristic"`
}
type CT_Characteristic struct{
	name string  `xml:"name",attr`
	relation ST_Relation  `xml:"relation",attr`
	val string  `xml:"val",attr`
	vocabulary string  `xml:"vocabulary",attr`
}
