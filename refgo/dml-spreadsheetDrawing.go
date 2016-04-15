package ooxml
type WordSchema  struct {
	from CT_Marker  `xml:"from"`
	to CT_Marker  `xml:"to"`
	wsDr CT_Drawing  `xml:"wsDr"`
}
type CT_AnchorClientData struct{
	fLocksWithSheet string  `xml:"fLocksWithSheet",attr`
	fPrintsWithSheet string  `xml:"fPrintsWithSheet",attr`
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
type CT_GraphicalObjectFrameNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
}
type CT_GraphicalObjectFrame struct{
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
	graphicFrame CT_GraphicalObjectFrame  `xml:"graphicFrame"`
	cxnSp CT_Connector  `xml:"cxnSp"`
	pic CT_Picture  `xml:"pic"`
}
type CT_Rel struct{
}
type CT_Marker struct{
	col ST_ColID  `xml:"col"`
	colOff string  `xml:"colOff"`
	row ST_RowID  `xml:"row"`
	rowOff string  `xml:"rowOff"`
}
type CT_TwoCellAnchor struct{
	from CT_Marker  `xml:"from"`
	to CT_Marker  `xml:"to"`
	clientData CT_AnchorClientData  `xml:"clientData"`
	 EG_ObjectChoices  `xml:""`
	editAs string  `xml:"editAs",attr`
}
type CT_OneCellAnchor struct{
	from CT_Marker  `xml:"from"`
	ext string  `xml:"ext"`
	clientData CT_AnchorClientData  `xml:"clientData"`
	 EG_ObjectChoices  `xml:""`
}
type CT_AbsoluteAnchor struct{
	pos string  `xml:"pos"`
	ext string  `xml:"ext"`
	clientData CT_AnchorClientData  `xml:"clientData"`
	 EG_ObjectChoices  `xml:""`
}
type CT_Drawing struct{
	 []EG_Anchor  `xml:""`
}
