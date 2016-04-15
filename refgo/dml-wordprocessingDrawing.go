package ooxml
type WordSchema  struct {
	wpc string  `xml:"wpc"`
	wgp string  `xml:"wgp"`
	wsp string  `xml:"wsp"`
	inline CT_Inline  `xml:"inline"`
	anchor CT_Anchor  `xml:"anchor"`
}
type CT_EffectExtent struct{
	l string  `xml:"l",attr`
	t string  `xml:"t",attr`
	r string  `xml:"r",attr`
	b string  `xml:"b",attr`
}
type CT_Inline struct{
	extent string  `xml:"extent"`
	effectExtent string  `xml:"effectExtent"`
	docPr string  `xml:"docPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
	distT string  `xml:"distT",attr`
	distB string  `xml:"distB",attr`
	distL string  `xml:"distL",attr`
	distR string  `xml:"distR",attr`
}
type CT_WrapPath struct{
	start string  `xml:"start"`
	lineTo []string  `xml:"lineTo"`
	edited string  `xml:"edited",attr`
}
type CT_WrapNone struct{
}
type CT_WrapSquare struct{
	effectExtent string  `xml:"effectExtent"`
	wrapText string  `xml:"wrapText",attr`
	distT string  `xml:"distT",attr`
	distB string  `xml:"distB",attr`
	distL string  `xml:"distL",attr`
	distR string  `xml:"distR",attr`
}
type CT_WrapTight struct{
	wrapPolygon CT_WrapPath  `xml:"wrapPolygon"`
	wrapText string  `xml:"wrapText",attr`
	distL string  `xml:"distL",attr`
	distR string  `xml:"distR",attr`
}
type CT_WrapThrough struct{
	wrapPolygon CT_WrapPath  `xml:"wrapPolygon"`
	wrapText string  `xml:"wrapText",attr`
	distL string  `xml:"distL",attr`
	distR string  `xml:"distR",attr`
}
type CT_WrapTopBottom struct{
	effectExtent string  `xml:"effectExtent"`
	distT string  `xml:"distT",attr`
	distB string  `xml:"distB",attr`
}
type CT_PosH struct{
	align ST_AlignH  `xml:"align"`
	posOffset string  `xml:"posOffset"`
	relativeFrom ST_RelFromH  `xml:"relativeFrom",attr`
}
type CT_PosV struct{
	align ST_AlignV  `xml:"align"`
	posOffset string  `xml:"posOffset"`
	relativeFrom ST_RelFromV  `xml:"relativeFrom",attr`
}
type CT_Anchor struct{
	simplePos string  `xml:"simplePos"`
	positionH string  `xml:"positionH"`
	positionV string  `xml:"positionV"`
	extent string  `xml:"extent"`
	effectExtent string  `xml:"effectExtent"`
	docPr string  `xml:"docPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
	 EG_WrapType  `xml:""`
	distT string  `xml:"distT",attr`
	distB string  `xml:"distB",attr`
	distL string  `xml:"distL",attr`
	distR string  `xml:"distR",attr`
	simplePos string  `xml:"simplePos",attr`
	relativeHeight string  `xml:"relativeHeight",attr`
	behindDoc string  `xml:"behindDoc",attr`
	locked string  `xml:"locked",attr`
	layoutInCell string  `xml:"layoutInCell",attr`
	hidden string  `xml:"hidden",attr`
	allowOverlap string  `xml:"allowOverlap",attr`
}
type CT_TxbxContent struct{
}
type CT_TextboxInfo struct{
	txbxContent string  `xml:"txbxContent"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
}
type CT_LinkedTextboxInformation struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	seq string  `xml:"seq",attr`
}
type CT_WordprocessingShape struct{
	cNvPr string  `xml:"cNvPr"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	extLst string  `xml:"extLst"`
	bodyPr string  `xml:"bodyPr"`
	cNvSpPr string  `xml:"cNvSpPr"`
	cNvCnPr string  `xml:"cNvCnPr"`
	txbx string  `xml:"txbx"`
	linkedTxbx string  `xml:"linkedTxbx"`
	normalEastAsianFlow string  `xml:"normalEastAsianFlow",attr`
}
type CT_GraphicFrame struct{
	cNvPr string  `xml:"cNvPr"`
	cNvFrPr string  `xml:"cNvFrPr"`
	xfrm string  `xml:"xfrm"`
	extLst string  `xml:"extLst"`
}
type CT_WordprocessingContentPartNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvContentPartPr string  `xml:"cNvContentPartPr"`
}
type CT_WordprocessingContentPart struct{
	nvContentPartPr string  `xml:"nvContentPartPr"`
	xfrm string  `xml:"xfrm"`
	extLst string  `xml:"extLst"`
	bwMode string  `xml:"bwMode",attr`
}
type CT_WordprocessingGroup struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGrpSpPr string  `xml:"cNvGrpSpPr"`
	grpSpPr string  `xml:"grpSpPr"`
	extLst string  `xml:"extLst"`
	grpSp string  `xml:"grpSp"`
	graphicFrame CT_GraphicFrame  `xml:"graphicFrame"`
	contentPart string  `xml:"contentPart"`
}
type CT_WordprocessingCanvas struct{
	bg string  `xml:"bg"`
	whole string  `xml:"whole"`
	extLst string  `xml:"extLst"`
	contentPart string  `xml:"contentPart"`
	graphicFrame CT_GraphicFrame  `xml:"graphicFrame"`
}
