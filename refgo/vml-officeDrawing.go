package ooxml
type WordSchema  struct {
	shapedefaults string  `xml:"shapedefaults"`
	shapelayout CT_ShapeLayout  `xml:"shapelayout"`
	signatureline CT_SignatureLine  `xml:"signatureline"`
	ink CT_Ink  `xml:"ink"`
	diagram CT_Diagram  `xml:"diagram"`
	equationxml CT_EquationXml  `xml:"equationxml"`
	skew CT_Skew  `xml:"skew"`
	extrusion string  `xml:"extrusion"`
	callout CT_Callout  `xml:"callout"`
	lock CT_Lock  `xml:"lock"`
	OLEObject CT_OLEObject  `xml:"OLEObject"`
	complex string  `xml:"complex"`
	left string  `xml:"left"`
	top string  `xml:"top"`
	right string  `xml:"right"`
	bottom string  `xml:"bottom"`
	column string  `xml:"column"`
	clippath CT_ClipPath  `xml:"clippath"`
	fill CT_Fill  `xml:"fill"`
}
type CT_ShapeDefaults struct{
	spidmax string  `xml:"spidmax",attr`
	style string  `xml:"style",attr`
	fill string  `xml:"fill",attr`
	fillcolor string  `xml:"fillcolor",attr`
	stroke string  `xml:"stroke",attr`
	strokecolor string  `xml:"strokecolor",attr`
	allowincell string  `xml:"allowincell",attr`
	colormru CT_ColorMru  `xml:"colormru"`
	colormenu CT_ColorMenu  `xml:"colormenu"`
}
type CT_Ink struct{
	i string  `xml:"i",attr`
	annotation string  `xml:"annotation",attr`
	contentType ST_ContentType  `xml:"contentType",attr`
}
type CT_SignatureLine struct{
	issignatureline string  `xml:"issignatureline",attr`
	id string  `xml:"id",attr`
	provid string  `xml:"provid",attr`
	signinginstructionsset string  `xml:"signinginstructionsset",attr`
	allowcomments string  `xml:"allowcomments",attr`
	showsigndate string  `xml:"showsigndate",attr`
	suggestedsigner string  `xml:"suggestedsigner",attr`
	suggestedsigner2 string  `xml:"suggestedsigner2",attr`
	suggestedsigneremail string  `xml:"suggestedsigneremail",attr`
	signinginstructions string  `xml:"signinginstructions",attr`
	addlxml string  `xml:"addlxml",attr`
	sigprovurl string  `xml:"sigprovurl",attr`
}
type CT_ShapeLayout struct{
	idmap string  `xml:"idmap"`
	regrouptable CT_RegroupTable  `xml:"regrouptable"`
	rules string  `xml:"rules"`
}
type CT_IdMap struct{
	data string  `xml:"data",attr`
}
type CT_RegroupTable struct{
	entry []CT_Entry  `xml:"entry"`
}
type CT_Entry struct{
	new string  `xml:"new",attr`
	old string  `xml:"old",attr`
}
type CT_Rules struct{
	r []CT_R  `xml:"r"`
}
type CT_R struct{
	proxy []string  `xml:"proxy"`
	id string  `xml:"id",attr`
	type ST_RType  `xml:"type",attr`
	how ST_How  `xml:"how",attr`
	idref string  `xml:"idref",attr`
}
type CT_Proxy struct{
	start string  `xml:"start",attr`
	end string  `xml:"end",attr`
	idref string  `xml:"idref",attr`
	connectloc string  `xml:"connectloc",attr`
}
type CT_Diagram struct{
	relationtable CT_RelationTable  `xml:"relationtable"`
	dgmstyle string  `xml:"dgmstyle",attr`
	autoformat string  `xml:"autoformat",attr`
	reverse string  `xml:"reverse",attr`
	autolayout string  `xml:"autolayout",attr`
	dgmscalex string  `xml:"dgmscalex",attr`
	dgmscaley string  `xml:"dgmscaley",attr`
	dgmfontsize string  `xml:"dgmfontsize",attr`
	constrainbounds string  `xml:"constrainbounds",attr`
	dgmbasetextscale string  `xml:"dgmbasetextscale",attr`
}
type CT_EquationXml struct{
	contentType ST_AlternateMathContentType  `xml:"contentType",attr`
}
type CT_RelationTable struct{
	rel []CT_Relation  `xml:"rel"`
}
type CT_Relation struct{
	idsrc string  `xml:"idsrc",attr`
	iddest string  `xml:"iddest",attr`
	idcntr string  `xml:"idcntr",attr`
}
type CT_ColorMru struct{
	colors string  `xml:"colors",attr`
}
type CT_ColorMenu struct{
	strokecolor string  `xml:"strokecolor",attr`
	fillcolor string  `xml:"fillcolor",attr`
	shadowcolor string  `xml:"shadowcolor",attr`
	extrusioncolor string  `xml:"extrusioncolor",attr`
}
type CT_Skew struct{
	id string  `xml:"id",attr`
	on string  `xml:"on",attr`
	offset string  `xml:"offset",attr`
	origin string  `xml:"origin",attr`
	matrix string  `xml:"matrix",attr`
}
type CT_Extrusion struct{
	on string  `xml:"on",attr`
	type string  `xml:"type",attr`
	render string  `xml:"render",attr`
	viewpointorigin string  `xml:"viewpointorigin",attr`
	viewpoint string  `xml:"viewpoint",attr`
	plane string  `xml:"plane",attr`
	skewangle string  `xml:"skewangle",attr`
	skewamt string  `xml:"skewamt",attr`
	foredepth string  `xml:"foredepth",attr`
	backdepth string  `xml:"backdepth",attr`
	orientation string  `xml:"orientation",attr`
	orientationangle string  `xml:"orientationangle",attr`
	lockrotationcenter string  `xml:"lockrotationcenter",attr`
	autorotationcenter string  `xml:"autorotationcenter",attr`
	rotationcenter string  `xml:"rotationcenter",attr`
	rotationangle string  `xml:"rotationangle",attr`
	colormode string  `xml:"colormode",attr`
	color string  `xml:"color",attr`
	shininess string  `xml:"shininess",attr`
	specularity string  `xml:"specularity",attr`
	diffusity string  `xml:"diffusity",attr`
	metal string  `xml:"metal",attr`
	edge string  `xml:"edge",attr`
	facet string  `xml:"facet",attr`
	lightface string  `xml:"lightface",attr`
	brightness string  `xml:"brightness",attr`
	lightposition string  `xml:"lightposition",attr`
	lightlevel string  `xml:"lightlevel",attr`
	lightharsh string  `xml:"lightharsh",attr`
	lightposition2 string  `xml:"lightposition2",attr`
	lightlevel2 string  `xml:"lightlevel2",attr`
	lightharsh2 string  `xml:"lightharsh2",attr`
}
type CT_Callout struct{
	on string  `xml:"on",attr`
	type string  `xml:"type",attr`
	gap string  `xml:"gap",attr`
	angle ST_Angle  `xml:"angle",attr`
	dropauto string  `xml:"dropauto",attr`
	drop ST_CalloutDrop  `xml:"drop",attr`
	distance string  `xml:"distance",attr`
	lengthspecified string  `xml:"lengthspecified",attr`
	length string  `xml:"length",attr`
	accentbar string  `xml:"accentbar",attr`
	textborder string  `xml:"textborder",attr`
	minusx string  `xml:"minusx",attr`
	minusy string  `xml:"minusy",attr`
}
type CT_Lock struct{
	position string  `xml:"position",attr`
	selection string  `xml:"selection",attr`
	grouping string  `xml:"grouping",attr`
	ungrouping string  `xml:"ungrouping",attr`
	rotation string  `xml:"rotation",attr`
	cropping string  `xml:"cropping",attr`
	verticies string  `xml:"verticies",attr`
	adjusthandles string  `xml:"adjusthandles",attr`
	text string  `xml:"text",attr`
	aspectratio string  `xml:"aspectratio",attr`
	shapetype string  `xml:"shapetype",attr`
}
type CT_OLEObject struct{
	LinkType ST_OLELinkType  `xml:"LinkType"`
	LockedField string  `xml:"LockedField"`
	FieldCodes string  `xml:"FieldCodes"`
	Type ST_OLEType  `xml:"Type",attr`
	ProgID string  `xml:"ProgID",attr`
	ShapeID string  `xml:"ShapeID",attr`
	DrawAspect string  `xml:"DrawAspect",attr`
	ObjectID string  `xml:"ObjectID",attr`
	UpdateMode string  `xml:"UpdateMode",attr`
}
type CT_Complex struct{
}
type CT_StrokeChild struct{
	on string  `xml:"on",attr`
	weight string  `xml:"weight",attr`
	color string  `xml:"color",attr`
	color2 string  `xml:"color2",attr`
	opacity string  `xml:"opacity",attr`
	linestyle string  `xml:"linestyle",attr`
	miterlimit string  `xml:"miterlimit",attr`
	joinstyle string  `xml:"joinstyle",attr`
	endcap string  `xml:"endcap",attr`
	dashstyle string  `xml:"dashstyle",attr`
	insetpen string  `xml:"insetpen",attr`
	filltype string  `xml:"filltype",attr`
	src string  `xml:"src",attr`
	imageaspect string  `xml:"imageaspect",attr`
	imagesize string  `xml:"imagesize",attr`
	imagealignshape string  `xml:"imagealignshape",attr`
	startarrow string  `xml:"startarrow",attr`
	startarrowwidth string  `xml:"startarrowwidth",attr`
	startarrowlength string  `xml:"startarrowlength",attr`
	endarrow string  `xml:"endarrow",attr`
	endarrowwidth string  `xml:"endarrowwidth",attr`
	endarrowlength string  `xml:"endarrowlength",attr`
}
type CT_ClipPath struct{
	v string  `xml:"v",attr`
}
type CT_Fill struct{
	type ST_FillType  `xml:"type",attr`
}
