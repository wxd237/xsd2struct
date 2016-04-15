package ooxml
type WordSchema  struct {
	shape CT_Shape  `xml:"shape"`
	shapetype CT_Shapetype  `xml:"shapetype"`
	group CT_Group  `xml:"group"`
	background string  `xml:"background"`
	fill CT_Fill  `xml:"fill"`
	formulas string  `xml:"formulas"`
	handles string  `xml:"handles"`
	imagedata CT_ImageData  `xml:"imagedata"`
	path CT_Path  `xml:"path"`
	textbox string  `xml:"textbox"`
	shadow string  `xml:"shadow"`
	stroke CT_Stroke  `xml:"stroke"`
	textpath string  `xml:"textpath"`
	arc CT_Arc  `xml:"arc"`
	curve CT_Curve  `xml:"curve"`
	image CT_Image  `xml:"image"`
	line CT_Line  `xml:"line"`
	oval CT_Oval  `xml:"oval"`
	polyline CT_PolyLine  `xml:"polyline"`
	rect CT_Rect  `xml:"rect"`
	roundrect string  `xml:"roundrect"`
}
type CT_Shape struct{
	 EG_ShapeElements  `xml:""`
	equationxml string  `xml:"equationxml",attr`
}
type CT_Shapetype struct{
	 []EG_ShapeElements  `xml:""`
}
type CT_Group struct{
	 EG_ShapeElements  `xml:""`
	editas string  `xml:"editas",attr`
}
type CT_Background struct{
}
type CT_Fill struct{
	type ST_FillType  `xml:"type",attr`
	on string  `xml:"on",attr`
	color string  `xml:"color",attr`
	opacity string  `xml:"opacity",attr`
	color2 string  `xml:"color2",attr`
	src string  `xml:"src",attr`
	size string  `xml:"size",attr`
	origin string  `xml:"origin",attr`
	position string  `xml:"position",attr`
	aspect string  `xml:"aspect",attr`
	colors string  `xml:"colors",attr`
	angle string  `xml:"angle",attr`
	alignshape string  `xml:"alignshape",attr`
	focus string  `xml:"focus",attr`
	focussize string  `xml:"focussize",attr`
	focusposition string  `xml:"focusposition",attr`
	method string  `xml:"method",attr`
	recolor string  `xml:"recolor",attr`
	rotate string  `xml:"rotate",attr`
}
type CT_Formulas struct{
	f []CT_F  `xml:"f"`
}
type CT_F struct{
	eqn string  `xml:"eqn",attr`
}
type CT_Handles struct{
	h []CT_H  `xml:"h"`
}
type CT_H struct{
	position string  `xml:"position",attr`
	polar string  `xml:"polar",attr`
	map string  `xml:"map",attr`
	invx string  `xml:"invx",attr`
	invy string  `xml:"invy",attr`
	switch string  `xml:"switch",attr`
	xrange string  `xml:"xrange",attr`
	yrange string  `xml:"yrange",attr`
	radiusrange string  `xml:"radiusrange",attr`
}
type CT_ImageData struct{
	embosscolor string  `xml:"embosscolor",attr`
	recolortarget string  `xml:"recolortarget",attr`
}
type CT_Path struct{
	v string  `xml:"v",attr`
	limo string  `xml:"limo",attr`
	textboxrect string  `xml:"textboxrect",attr`
	fillok string  `xml:"fillok",attr`
	strokeok string  `xml:"strokeok",attr`
	shadowok string  `xml:"shadowok",attr`
	arrowok string  `xml:"arrowok",attr`
	gradientshapeok string  `xml:"gradientshapeok",attr`
	textpathok string  `xml:"textpathok",attr`
	insetpenok string  `xml:"insetpenok",attr`
}
type CT_Shadow struct{
	on string  `xml:"on",attr`
	type string  `xml:"type",attr`
	obscured string  `xml:"obscured",attr`
	color string  `xml:"color",attr`
	opacity string  `xml:"opacity",attr`
	offset string  `xml:"offset",attr`
	color2 string  `xml:"color2",attr`
	offset2 string  `xml:"offset2",attr`
	origin string  `xml:"origin",attr`
	matrix string  `xml:"matrix",attr`
}
type CT_Stroke struct{
}
type CT_Textbox struct{
	inset string  `xml:"inset",attr`
}
type CT_TextPath struct{
	on string  `xml:"on",attr`
	fitshape string  `xml:"fitshape",attr`
	fitpath string  `xml:"fitpath",attr`
	trim string  `xml:"trim",attr`
	xscale string  `xml:"xscale",attr`
	string string  `xml:"string",attr`
}
type CT_Arc struct{
	 []EG_ShapeElements  `xml:""`
	startAngle string  `xml:"startAngle",attr`
	endAngle string  `xml:"endAngle",attr`
}
type CT_Curve struct{
	 []EG_ShapeElements  `xml:""`
	from string  `xml:"from",attr`
	control1 string  `xml:"control1",attr`
	control2 string  `xml:"control2",attr`
	to string  `xml:"to",attr`
}
type CT_Image struct{
	 []EG_ShapeElements  `xml:""`
}
type CT_Line struct{
	 []EG_ShapeElements  `xml:""`
	from string  `xml:"from",attr`
	to string  `xml:"to",attr`
}
type CT_Oval struct{
	 []EG_ShapeElements  `xml:""`
}
type CT_PolyLine struct{
	 EG_ShapeElements  `xml:""`
	points string  `xml:"points",attr`
}
type CT_Rect struct{
	 []EG_ShapeElements  `xml:""`
}
type CT_RoundRect struct{
	 []EG_ShapeElements  `xml:""`
	arcsize string  `xml:"arcsize",attr`
}
