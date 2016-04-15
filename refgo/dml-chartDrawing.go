package ooxml
type WordSchema  struct {
}
type CT_ShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvSpPr string  `xml:"cNvSpPr"`
}
type CT_Shape struct{
	nvSpPr string  `xml:"nvSpPr"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	txBody string  `xml:"txBody"`
	macro string  `xml:"macro",attr`
	textlink string  `xml:"textlink",attr`
	fLocksText string  `xml:"fLocksText",attr`
	fPublished string  `xml:"fPublished",attr`
}
type CT_ConnectorNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvCxnSpPr string  `xml:"cNvCxnSpPr"`
}
type CT_Connector struct{
	nvCxnSpPr string  `xml:"nvCxnSpPr"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	macro string  `xml:"macro",attr`
	fPublished string  `xml:"fPublished",attr`
}
type CT_PictureNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvPicPr string  `xml:"cNvPicPr"`
}
type CT_Picture struct{
	nvPicPr string  `xml:"nvPicPr"`
	blipFill string  `xml:"blipFill"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	macro string  `xml:"macro",attr`
	fPublished string  `xml:"fPublished",attr`
}
type CT_GraphicFrameNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
}
type CT_GraphicFrame struct{
	nvGraphicFramePr string  `xml:"nvGraphicFramePr"`
	xfrm string  `xml:"xfrm"`
	macro string  `xml:"macro",attr`
	fPublished string  `xml:"fPublished",attr`
}
type CT_GroupShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGrpSpPr string  `xml:"cNvGrpSpPr"`
}
type CT_GroupShape struct{
	nvGrpSpPr string  `xml:"nvGrpSpPr"`
	grpSpPr string  `xml:"grpSpPr"`
	sp CT_Shape  `xml:"sp"`
	grpSp CT_GroupShape  `xml:"grpSp"`
	graphicFrame CT_GraphicFrame  `xml:"graphicFrame"`
	cxnSp CT_Connector  `xml:"cxnSp"`
	pic CT_Picture  `xml:"pic"`
}
type CT_Marker struct{
	x string  `xml:"x"`
	y string  `xml:"y"`
}
type CT_RelSizeAnchor struct{
	from CT_Marker  `xml:"from"`
	to CT_Marker  `xml:"to"`
	 EG_ObjectChoices  `xml:""`
}
type CT_AbsSizeAnchor struct{
	from CT_Marker  `xml:"from"`
	ext string  `xml:"ext"`
	 EG_ObjectChoices  `xml:""`
}
type CT_Drawing struct{
	 []EG_Anchor  `xml:""`
}
