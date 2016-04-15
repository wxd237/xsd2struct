package ooxml
type WordSchema  struct {
	bordertop string  `xml:"bordertop"`
	borderleft string  `xml:"borderleft"`
	borderright string  `xml:"borderright"`
	borderbottom string  `xml:"borderbottom"`
	wrap CT_Wrap  `xml:"wrap"`
	anchorlock CT_AnchorLock  `xml:"anchorlock"`
}
type CT_Border struct{
	type string  `xml:"type",attr`
	width string  `xml:"width",attr`
	shadow string  `xml:"shadow",attr`
}
type CT_Wrap struct{
	type ST_WrapType  `xml:"type",attr`
	side string  `xml:"side",attr`
	anchorx ST_HorizontalAnchor  `xml:"anchorx",attr`
	anchory ST_VerticalAnchor  `xml:"anchory",attr`
}
type CT_AnchorLock struct{
}
