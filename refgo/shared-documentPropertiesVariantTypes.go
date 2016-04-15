package ooxml
type WordSchema  struct {
	variant CT_Variant  `xml:"variant"`
	vector CT_Vector  `xml:"vector"`
	array CT_Array  `xml:"array"`
	blob string  `xml:"blob"`
	oblob string  `xml:"oblob"`
	empty CT_Empty  `xml:"empty"`
	null CT_Null  `xml:"null"`
	i1 string  `xml:"i1"`
	i2 string  `xml:"i2"`
	i4 string  `xml:"i4"`
	i8 string  `xml:"i8"`
	int string  `xml:"int"`
	ui1 string  `xml:"ui1"`
	ui2 string  `xml:"ui2"`
	ui4 string  `xml:"ui4"`
	ui8 string  `xml:"ui8"`
	uint string  `xml:"uint"`
	r4 string  `xml:"r4"`
	r8 string  `xml:"r8"`
	decimal string  `xml:"decimal"`
	lpstr string  `xml:"lpstr"`
	lpwstr string  `xml:"lpwstr"`
	bstr string  `xml:"bstr"`
	date string  `xml:"date"`
	filetime string  `xml:"filetime"`
	bool string  `xml:"bool"`
	cy ST_Cy  `xml:"cy"`
	error ST_Error  `xml:"error"`
	stream string  `xml:"stream"`
	ostream string  `xml:"ostream"`
	storage string  `xml:"storage"`
	ostorage string  `xml:"ostorage"`
	vstream string  `xml:"vstream"`
	clsid string  `xml:"clsid"`
}
type CT_Empty struct{
}
type CT_Null struct{
}
type CT_Vector struct{
	baseType string  `xml:"baseType",attr`
	size string  `xml:"size",attr`
}
type CT_Array struct{
	lBounds string  `xml:"lBounds",attr`
	uBounds string  `xml:"uBounds",attr`
	baseType string  `xml:"baseType",attr`
}
type CT_Variant struct{
}
type CT_Vstream struct{
}
