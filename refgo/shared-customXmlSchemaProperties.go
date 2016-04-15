package ooxml
type WordSchema  struct {
	schemaLibrary CT_SchemaLibrary  `xml:"schemaLibrary"`
}
type CT_Schema struct{
	uri string  `xml:"uri",attr`
	manifestLocation string  `xml:"manifestLocation",attr`
	schemaLocation string  `xml:"schemaLocation",attr`
	schemaLanguage string  `xml:"schemaLanguage",attr`
}
type CT_SchemaLibrary struct{
	schema []CT_Schema  `xml:"schema"`
}
