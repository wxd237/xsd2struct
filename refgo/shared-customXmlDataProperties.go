package ooxml
type WordSchema  struct {
	datastoreItem string  `xml:"datastoreItem"`
}
type CT_DatastoreSchemaRef struct{
	uri string  `xml:"uri",attr`
}
type CT_DatastoreSchemaRefs struct{
	schemaRef []string  `xml:"schemaRef"`
}
type CT_DatastoreItem struct{
	schemaRefs string  `xml:"schemaRefs"`
	itemID string  `xml:"itemID",attr`
}
