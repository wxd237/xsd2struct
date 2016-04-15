package ooxml
type WordSchema  struct {
	videoFile string  `xml:"videoFile"`
	graphic CT_GraphicalObject  `xml:"graphic"`
	blip CT_Blip  `xml:"blip"`
	theme CT_OfficeStyleSheet  `xml:"theme"`
	themeOverride string  `xml:"themeOverride"`
	themeManager CT_EmptyElement  `xml:"themeManager"`
	tbl CT_Table  `xml:"tbl"`
	tblStyleLst string  `xml:"tblStyleLst"`
}
type CT_AudioFile struct{
	extLst string  `xml:"extLst"`
	contentType string  `xml:"contentType",attr`
}
type CT_VideoFile struct{
	extLst string  `xml:"extLst"`
	contentType string  `xml:"contentType",attr`
}
type CT_QuickTimeFile struct{
	extLst string  `xml:"extLst"`
}
type CT_AudioCDTime struct{
	track string  `xml:"track",attr`
	time string  `xml:"time",attr`
}
type CT_AudioCD struct{
	st string  `xml:"st"`
	end string  `xml:"end"`
	extLst string  `xml:"extLst"`
}
type CT_ColorScheme struct{
	dk1 CT_Color  `xml:"dk1"`
	lt1 CT_Color  `xml:"lt1"`
	dk2 CT_Color  `xml:"dk2"`
	lt2 CT_Color  `xml:"lt2"`
	accent1 CT_Color  `xml:"accent1"`
	accent2 CT_Color  `xml:"accent2"`
	accent3 CT_Color  `xml:"accent3"`
	accent4 CT_Color  `xml:"accent4"`
	accent5 CT_Color  `xml:"accent5"`
	accent6 CT_Color  `xml:"accent6"`
	hlink CT_Color  `xml:"hlink"`
	folHlink CT_Color  `xml:"folHlink"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_CustomColor struct{
	 EG_ColorChoice  `xml:""`
	name string  `xml:"name",attr`
}
type CT_SupplementalFont struct{
	script string  `xml:"script",attr`
	typeface string  `xml:"typeface",attr`
}
type CT_CustomColorList struct{
	custClr []string  `xml:"custClr"`
}
type CT_FontCollection struct{
	latin string  `xml:"latin"`
	ea string  `xml:"ea"`
	cs string  `xml:"cs"`
	font []CT_SupplementalFont  `xml:"font"`
	extLst string  `xml:"extLst"`
}
type CT_EffectStyleItem struct{
	scene3d CT_Scene3D  `xml:"scene3d"`
	sp3d CT_Shape3D  `xml:"sp3d"`
	 EG_EffectProperties  `xml:""`
}
type CT_FontScheme struct{
	majorFont CT_FontCollection  `xml:"majorFont"`
	minorFont CT_FontCollection  `xml:"minorFont"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_FillStyleList struct{
	 []EG_FillProperties  `xml:""`
}
type CT_LineStyleList struct{
	ln []string  `xml:"ln"`
}
type CT_EffectStyleList struct{
	effectStyle []CT_EffectStyleItem  `xml:"effectStyle"`
}
type CT_BackgroundFillStyleList struct{
	 []EG_FillProperties  `xml:""`
}
type CT_StyleMatrix struct{
	fillStyleLst string  `xml:"fillStyleLst"`
	lnStyleLst string  `xml:"lnStyleLst"`
	effectStyleLst string  `xml:"effectStyleLst"`
	bgFillStyleLst string  `xml:"bgFillStyleLst"`
	name string  `xml:"name",attr`
}
type CT_BaseStyles struct{
	clrScheme CT_ColorScheme  `xml:"clrScheme"`
	fontScheme CT_FontScheme  `xml:"fontScheme"`
	fmtScheme string  `xml:"fmtScheme"`
	extLst string  `xml:"extLst"`
}
type CT_OfficeArtExtension struct{
	uri string  `xml:"uri",attr`
}
type CT_Angle struct{
	val ST_Angle  `xml:"val",attr`
}
type CT_PositiveFixedAngle struct{
	val string  `xml:"val",attr`
}
type CT_Percentage struct{
	val ST_Percentage  `xml:"val",attr`
}
type CT_PositivePercentage struct{
	val string  `xml:"val",attr`
}
type CT_FixedPercentage struct{
	val string  `xml:"val",attr`
}
type CT_PositiveFixedPercentage struct{
	val string  `xml:"val",attr`
}
type CT_Ratio struct{
	n string  `xml:"n",attr`
	d string  `xml:"d",attr`
}
type CT_Point2D struct{
	x string  `xml:"x",attr`
	y string  `xml:"y",attr`
}
type CT_PositiveSize2D struct{
	cx string  `xml:"cx",attr`
	cy string  `xml:"cy",attr`
}
type CT_ComplementTransform struct{
}
type CT_InverseTransform struct{
}
type CT_GrayscaleTransform struct{
}
type CT_GammaTransform struct{
}
type CT_InverseGammaTransform struct{
}
type CT_ScRgbColor struct{
	 []EG_ColorTransform  `xml:""`
	r ST_Percentage  `xml:"r",attr`
	g ST_Percentage  `xml:"g",attr`
	b ST_Percentage  `xml:"b",attr`
}
type CT_SRgbColor struct{
	 []EG_ColorTransform  `xml:""`
	val string  `xml:"val",attr`
}
type CT_HslColor struct{
	 []EG_ColorTransform  `xml:""`
	hue string  `xml:"hue",attr`
	sat ST_Percentage  `xml:"sat",attr`
	lum ST_Percentage  `xml:"lum",attr`
}
type CT_SystemColor struct{
	 []EG_ColorTransform  `xml:""`
	val string  `xml:"val",attr`
	lastClr string  `xml:"lastClr",attr`
}
type CT_SchemeColor struct{
	 []EG_ColorTransform  `xml:""`
	val ST_SchemeColorVal  `xml:"val",attr`
}
type CT_PresetColor struct{
	 []EG_ColorTransform  `xml:""`
	val string  `xml:"val",attr`
}
type CT_OfficeArtExtensionList struct{
	 EG_OfficeArtExtensionList  `xml:""`
}
type CT_Scale2D struct{
	sx CT_Ratio  `xml:"sx"`
	sy CT_Ratio  `xml:"sy"`
}
type CT_Transform2D struct{
	off CT_Point2D  `xml:"off"`
	ext string  `xml:"ext"`
	rot ST_Angle  `xml:"rot",attr`
	flipH string  `xml:"flipH",attr`
	flipV string  `xml:"flipV",attr`
}
type CT_GroupTransform2D struct{
	off CT_Point2D  `xml:"off"`
	ext string  `xml:"ext"`
	chOff CT_Point2D  `xml:"chOff"`
	chExt string  `xml:"chExt"`
	rot ST_Angle  `xml:"rot",attr`
	flipH string  `xml:"flipH",attr`
	flipV string  `xml:"flipV",attr`
}
type CT_Point3D struct{
	x string  `xml:"x",attr`
	y string  `xml:"y",attr`
	z string  `xml:"z",attr`
}
type CT_Vector3D struct{
	dx string  `xml:"dx",attr`
	dy string  `xml:"dy",attr`
	dz string  `xml:"dz",attr`
}
type CT_SphereCoords struct{
	lat string  `xml:"lat",attr`
	lon string  `xml:"lon",attr`
	rev string  `xml:"rev",attr`
}
type CT_RelativeRect struct{
	l ST_Percentage  `xml:"l",attr`
	t ST_Percentage  `xml:"t",attr`
	r ST_Percentage  `xml:"r",attr`
	b ST_Percentage  `xml:"b",attr`
}
type CT_Color struct{
	 EG_ColorChoice  `xml:""`
}
type CT_ColorMRU struct{
	 []EG_ColorChoice  `xml:""`
}
type CT_EmbeddedWAVAudioFile struct{
	name string  `xml:"name",attr`
}
type CT_Hyperlink struct{
	snd string  `xml:"snd"`
	extLst string  `xml:"extLst"`
	invalidUrl string  `xml:"invalidUrl",attr`
	action string  `xml:"action",attr`
	tgtFrame string  `xml:"tgtFrame",attr`
	tooltip string  `xml:"tooltip",attr`
	history string  `xml:"history",attr`
	highlightClick string  `xml:"highlightClick",attr`
	endSnd string  `xml:"endSnd",attr`
}
type CT_ConnectorLocking struct{
	extLst string  `xml:"extLst"`
}
type CT_ShapeLocking struct{
	extLst string  `xml:"extLst"`
	noTextEdit string  `xml:"noTextEdit",attr`
}
type CT_PictureLocking struct{
	extLst string  `xml:"extLst"`
	noCrop string  `xml:"noCrop",attr`
}
type CT_GroupLocking struct{
	extLst string  `xml:"extLst"`
	noGrp string  `xml:"noGrp",attr`
	noUngrp string  `xml:"noUngrp",attr`
	noSelect string  `xml:"noSelect",attr`
	noRot string  `xml:"noRot",attr`
	noChangeAspect string  `xml:"noChangeAspect",attr`
	noMove string  `xml:"noMove",attr`
	noResize string  `xml:"noResize",attr`
}
type CT_GraphicalObjectFrameLocking struct{
	extLst string  `xml:"extLst"`
	noGrp string  `xml:"noGrp",attr`
	noDrilldown string  `xml:"noDrilldown",attr`
	noSelect string  `xml:"noSelect",attr`
	noChangeAspect string  `xml:"noChangeAspect",attr`
	noMove string  `xml:"noMove",attr`
	noResize string  `xml:"noResize",attr`
}
type CT_ContentPartLocking struct{
	extLst string  `xml:"extLst"`
}
type CT_NonVisualDrawingProps struct{
	hlinkClick CT_Hyperlink  `xml:"hlinkClick"`
	hlinkHover CT_Hyperlink  `xml:"hlinkHover"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	name string  `xml:"name",attr`
	descr string  `xml:"descr",attr`
	hidden string  `xml:"hidden",attr`
	title string  `xml:"title",attr`
}
type CT_NonVisualDrawingShapeProps struct{
	spLocks CT_ShapeLocking  `xml:"spLocks"`
	extLst string  `xml:"extLst"`
	txBox string  `xml:"txBox",attr`
}
type CT_NonVisualConnectorProperties struct{
	cxnSpLocks CT_ConnectorLocking  `xml:"cxnSpLocks"`
	stCxn CT_Connection  `xml:"stCxn"`
	endCxn CT_Connection  `xml:"endCxn"`
	extLst string  `xml:"extLst"`
}
type CT_NonVisualPictureProperties struct{
	picLocks CT_PictureLocking  `xml:"picLocks"`
	extLst string  `xml:"extLst"`
	preferRelativeResize string  `xml:"preferRelativeResize",attr`
}
type CT_NonVisualGroupDrawingShapeProps struct{
	grpSpLocks CT_GroupLocking  `xml:"grpSpLocks"`
	extLst string  `xml:"extLst"`
}
type CT_NonVisualGraphicFrameProperties struct{
	graphicFrameLocks CT_GraphicalObjectFrameLocking  `xml:"graphicFrameLocks"`
	extLst string  `xml:"extLst"`
}
type CT_NonVisualContentPartProperties struct{
	cpLocks CT_ContentPartLocking  `xml:"cpLocks"`
	extLst string  `xml:"extLst"`
	isComment string  `xml:"isComment",attr`
}
type CT_GraphicalObjectData struct{
	uri string  `xml:"uri",attr`
}
type CT_GraphicalObject struct{
	graphicData CT_GraphicalObjectData  `xml:"graphicData"`
}
type CT_AnimationDgmElement struct{
	id string  `xml:"id",attr`
	bldStep string  `xml:"bldStep",attr`
}
type CT_AnimationChartElement struct{
	seriesIdx string  `xml:"seriesIdx",attr`
	categoryIdx string  `xml:"categoryIdx",attr`
	bldStep string  `xml:"bldStep",attr`
}
type CT_AnimationElementChoice struct{
	dgm CT_AnimationDgmElement  `xml:"dgm"`
	chart CT_AnimationChartElement  `xml:"chart"`
}
type CT_AnimationDgmBuildProperties struct{
	bld string  `xml:"bld",attr`
	rev string  `xml:"rev",attr`
}
type CT_AnimationChartBuildProperties struct{
	bld string  `xml:"bld",attr`
	animBg string  `xml:"animBg",attr`
}
type CT_AnimationGraphicalObjectBuildProperties struct{
	bldDgm string  `xml:"bldDgm"`
	bldChart string  `xml:"bldChart"`
}
type CT_BackgroundFormatting struct{
	 EG_EffectProperties  `xml:""`
}
type CT_WholeE2oFormatting struct{
	ln string  `xml:"ln"`
	 EG_EffectProperties  `xml:""`
}
type CT_GvmlUseShapeRectangle struct{
}
type CT_GvmlTextShape struct{
	txBody string  `xml:"txBody"`
	extLst string  `xml:"extLst"`
	useSpRect string  `xml:"useSpRect"`
	xfrm string  `xml:"xfrm"`
}
type CT_GvmlShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvSpPr string  `xml:"cNvSpPr"`
}
type CT_GvmlShape struct{
	nvSpPr string  `xml:"nvSpPr"`
	spPr string  `xml:"spPr"`
	txSp string  `xml:"txSp"`
	style CT_ShapeStyle  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_GvmlConnectorNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvCxnSpPr string  `xml:"cNvCxnSpPr"`
}
type CT_GvmlConnector struct{
	nvCxnSpPr string  `xml:"nvCxnSpPr"`
	spPr string  `xml:"spPr"`
	style CT_ShapeStyle  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_GvmlPictureNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvPicPr string  `xml:"cNvPicPr"`
}
type CT_GvmlPicture struct{
	nvPicPr string  `xml:"nvPicPr"`
	blipFill string  `xml:"blipFill"`
	spPr string  `xml:"spPr"`
	style CT_ShapeStyle  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_GvmlGraphicFrameNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
}
type CT_GvmlGraphicalObjectFrame struct{
	nvGraphicFramePr string  `xml:"nvGraphicFramePr"`
	xfrm string  `xml:"xfrm"`
	extLst string  `xml:"extLst"`
}
type CT_GvmlGroupShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGrpSpPr string  `xml:"cNvGrpSpPr"`
}
type CT_GvmlGroupShape struct{
	nvGrpSpPr string  `xml:"nvGrpSpPr"`
	grpSpPr string  `xml:"grpSpPr"`
	extLst string  `xml:"extLst"`
	txSp string  `xml:"txSp"`
	sp CT_GvmlShape  `xml:"sp"`
	cxnSp CT_GvmlConnector  `xml:"cxnSp"`
	pic CT_GvmlPicture  `xml:"pic"`
	graphicFrame CT_GvmlGraphicalObjectFrame  `xml:"graphicFrame"`
	grpSp CT_GvmlGroupShape  `xml:"grpSp"`
}
type CT_Camera struct{
	rot string  `xml:"rot"`
	prst string  `xml:"prst",attr`
	fov ST_FOVAngle  `xml:"fov",attr`
	zoom string  `xml:"zoom",attr`
}
type CT_LightRig struct{
	rot string  `xml:"rot"`
	rig ST_LightRigType  `xml:"rig",attr`
	dir ST_LightRigDirection  `xml:"dir",attr`
}
type CT_Scene3D struct{
	camera CT_Camera  `xml:"camera"`
	lightRig CT_LightRig  `xml:"lightRig"`
	backdrop string  `xml:"backdrop"`
	extLst string  `xml:"extLst"`
}
type CT_Backdrop struct{
	anchor CT_Point3D  `xml:"anchor"`
	norm CT_Vector3D  `xml:"norm"`
	up CT_Vector3D  `xml:"up"`
	extLst string  `xml:"extLst"`
}
type CT_Bevel struct{
	w string  `xml:"w",attr`
	h string  `xml:"h",attr`
	prst string  `xml:"prst",attr`
}
type CT_Shape3D struct{
	bevelT CT_Bevel  `xml:"bevelT"`
	bevelB CT_Bevel  `xml:"bevelB"`
	extrusionClr CT_Color  `xml:"extrusionClr"`
	contourClr CT_Color  `xml:"contourClr"`
	extLst string  `xml:"extLst"`
	z string  `xml:"z",attr`
	extrusionH string  `xml:"extrusionH",attr`
	contourW string  `xml:"contourW",attr`
	prstMaterial string  `xml:"prstMaterial",attr`
}
type CT_FlatText struct{
	z string  `xml:"z",attr`
}
type CT_AlphaBiLevelEffect struct{
	thresh string  `xml:"thresh",attr`
}
type CT_AlphaCeilingEffect struct{
}
type CT_AlphaFloorEffect struct{
}
type CT_AlphaInverseEffect struct{
	 EG_ColorChoice  `xml:""`
}
type CT_AlphaModulateFixedEffect struct{
	amt string  `xml:"amt",attr`
}
type CT_AlphaOutsetEffect struct{
	rad string  `xml:"rad",attr`
}
type CT_AlphaReplaceEffect struct{
	a string  `xml:"a",attr`
}
type CT_BiLevelEffect struct{
	thresh string  `xml:"thresh",attr`
}
type CT_BlurEffect struct{
	rad string  `xml:"rad",attr`
	grow string  `xml:"grow",attr`
}
type CT_ColorChangeEffect struct{
	clrFrom CT_Color  `xml:"clrFrom"`
	clrTo CT_Color  `xml:"clrTo"`
	useA string  `xml:"useA",attr`
}
type CT_ColorReplaceEffect struct{
	 EG_ColorChoice  `xml:""`
}
type CT_DuotoneEffect struct{
	 EG_ColorChoice  `xml:""`
}
type CT_GlowEffect struct{
	 EG_ColorChoice  `xml:""`
	rad string  `xml:"rad",attr`
}
type CT_GrayscaleEffect struct{
}
type CT_HSLEffect struct{
	hue string  `xml:"hue",attr`
	sat string  `xml:"sat",attr`
	lum string  `xml:"lum",attr`
}
type CT_InnerShadowEffect struct{
	 EG_ColorChoice  `xml:""`
	blurRad string  `xml:"blurRad",attr`
	dist string  `xml:"dist",attr`
	dir string  `xml:"dir",attr`
}
type CT_LuminanceEffect struct{
	bright string  `xml:"bright",attr`
	contrast string  `xml:"contrast",attr`
}
type CT_OuterShadowEffect struct{
	 EG_ColorChoice  `xml:""`
	blurRad string  `xml:"blurRad",attr`
	dist string  `xml:"dist",attr`
	dir string  `xml:"dir",attr`
	sx ST_Percentage  `xml:"sx",attr`
	sy ST_Percentage  `xml:"sy",attr`
	kx string  `xml:"kx",attr`
	ky string  `xml:"ky",attr`
	algn ST_RectAlignment  `xml:"algn",attr`
	rotWithShape string  `xml:"rotWithShape",attr`
}
type CT_PresetShadowEffect struct{
	 EG_ColorChoice  `xml:""`
	prst string  `xml:"prst",attr`
	dist string  `xml:"dist",attr`
	dir string  `xml:"dir",attr`
}
type CT_ReflectionEffect struct{
	blurRad string  `xml:"blurRad",attr`
	stA string  `xml:"stA",attr`
	stPos string  `xml:"stPos",attr`
	endA string  `xml:"endA",attr`
	endPos string  `xml:"endPos",attr`
	dist string  `xml:"dist",attr`
	dir string  `xml:"dir",attr`
	fadeDir string  `xml:"fadeDir",attr`
	sx ST_Percentage  `xml:"sx",attr`
	sy ST_Percentage  `xml:"sy",attr`
	kx string  `xml:"kx",attr`
	ky string  `xml:"ky",attr`
	algn ST_RectAlignment  `xml:"algn",attr`
	rotWithShape string  `xml:"rotWithShape",attr`
}
type CT_RelativeOffsetEffect struct{
	tx ST_Percentage  `xml:"tx",attr`
	ty ST_Percentage  `xml:"ty",attr`
}
type CT_SoftEdgesEffect struct{
	rad string  `xml:"rad",attr`
}
type CT_TintEffect struct{
	hue string  `xml:"hue",attr`
	amt string  `xml:"amt",attr`
}
type CT_TransformEffect struct{
	sx ST_Percentage  `xml:"sx",attr`
	sy ST_Percentage  `xml:"sy",attr`
	kx string  `xml:"kx",attr`
	ky string  `xml:"ky",attr`
	tx string  `xml:"tx",attr`
	ty string  `xml:"ty",attr`
}
type CT_NoFillProperties struct{
}
type CT_SolidColorFillProperties struct{
	 EG_ColorChoice  `xml:""`
}
type CT_LinearShadeProperties struct{
	ang string  `xml:"ang",attr`
	scaled string  `xml:"scaled",attr`
}
type CT_PathShadeProperties struct{
	fillToRect CT_RelativeRect  `xml:"fillToRect"`
	path string  `xml:"path",attr`
}
type CT_GradientStop struct{
	 EG_ColorChoice  `xml:""`
	pos string  `xml:"pos",attr`
}
type CT_GradientStopList struct{
	gs []string  `xml:"gs"`
}
type CT_GradientFillProperties struct{
	gsLst string  `xml:"gsLst"`
	tileRect CT_RelativeRect  `xml:"tileRect"`
	 EG_ShadeProperties  `xml:""`
	flip string  `xml:"flip",attr`
	rotWithShape string  `xml:"rotWithShape",attr`
}
type CT_TileInfoProperties struct{
	tx string  `xml:"tx",attr`
	ty string  `xml:"ty",attr`
	sx ST_Percentage  `xml:"sx",attr`
	sy ST_Percentage  `xml:"sy",attr`
	flip string  `xml:"flip",attr`
	algn ST_RectAlignment  `xml:"algn",attr`
}
type CT_StretchInfoProperties struct{
	fillRect CT_RelativeRect  `xml:"fillRect"`
}
type CT_Blip struct{
	extLst string  `xml:"extLst"`
	alphaBiLevel CT_AlphaBiLevelEffect  `xml:"alphaBiLevel"`
	alphaCeiling CT_AlphaCeilingEffect  `xml:"alphaCeiling"`
	alphaFloor CT_AlphaFloorEffect  `xml:"alphaFloor"`
	alphaInv string  `xml:"alphaInv"`
	alphaMod string  `xml:"alphaMod"`
	alphaModFix string  `xml:"alphaModFix"`
	alphaRepl CT_AlphaReplaceEffect  `xml:"alphaRepl"`
	biLevel CT_BiLevelEffect  `xml:"biLevel"`
	blur CT_BlurEffect  `xml:"blur"`
	clrChange CT_ColorChangeEffect  `xml:"clrChange"`
	clrRepl CT_ColorReplaceEffect  `xml:"clrRepl"`
	duotone CT_DuotoneEffect  `xml:"duotone"`
	fillOverlay CT_FillOverlayEffect  `xml:"fillOverlay"`
	grayscl string  `xml:"grayscl"`
	hsl CT_HSLEffect  `xml:"hsl"`
	lum CT_LuminanceEffect  `xml:"lum"`
	tint CT_TintEffect  `xml:"tint"`
	cstate string  `xml:"cstate",attr`
}
type CT_BlipFillProperties struct{
	blip CT_Blip  `xml:"blip"`
	srcRect CT_RelativeRect  `xml:"srcRect"`
	 EG_FillModeProperties  `xml:""`
	dpi string  `xml:"dpi",attr`
	rotWithShape string  `xml:"rotWithShape",attr`
}
type CT_PatternFillProperties struct{
	fgClr CT_Color  `xml:"fgClr"`
	bgClr CT_Color  `xml:"bgClr"`
	prst string  `xml:"prst",attr`
}
type CT_GroupFillProperties struct{
}
type CT_FillProperties struct{
	 EG_FillProperties  `xml:""`
}
type CT_FillEffect struct{
	 EG_FillProperties  `xml:""`
}
type CT_FillOverlayEffect struct{
	 EG_FillProperties  `xml:""`
	blend string  `xml:"blend",attr`
}
type CT_EffectReference struct{
	ref string  `xml:"ref",attr`
}
type CT_EffectContainer struct{
	type ST_EffectContainerType  `xml:"type",attr`
	name string  `xml:"name",attr`
}
type CT_AlphaModulateEffect struct{
	cont CT_EffectContainer  `xml:"cont"`
}
type CT_BlendEffect struct{
	cont CT_EffectContainer  `xml:"cont"`
	blend string  `xml:"blend",attr`
}
type CT_EffectList struct{
	blur CT_BlurEffect  `xml:"blur"`
	fillOverlay CT_FillOverlayEffect  `xml:"fillOverlay"`
	glow CT_GlowEffect  `xml:"glow"`
	innerShdw string  `xml:"innerShdw"`
	outerShdw string  `xml:"outerShdw"`
	prstShdw string  `xml:"prstShdw"`
	reflection CT_ReflectionEffect  `xml:"reflection"`
	softEdge string  `xml:"softEdge"`
}
type CT_EffectProperties struct{
	 EG_EffectProperties  `xml:""`
}
type CT_GeomGuide struct{
	name string  `xml:"name",attr`
	fmla string  `xml:"fmla",attr`
}
type CT_GeomGuideList struct{
	gd []string  `xml:"gd"`
}
type CT_AdjPoint2D struct{
	x string  `xml:"x",attr`
	y string  `xml:"y",attr`
}
type CT_GeomRect struct{
	l string  `xml:"l",attr`
	t string  `xml:"t",attr`
	r string  `xml:"r",attr`
	b string  `xml:"b",attr`
}
type CT_XYAdjustHandle struct{
	pos string  `xml:"pos"`
	gdRefX string  `xml:"gdRefX",attr`
	minX string  `xml:"minX",attr`
	maxX string  `xml:"maxX",attr`
	gdRefY string  `xml:"gdRefY",attr`
	minY string  `xml:"minY",attr`
	maxY string  `xml:"maxY",attr`
}
type CT_PolarAdjustHandle struct{
	pos string  `xml:"pos"`
	gdRefR string  `xml:"gdRefR",attr`
	minR string  `xml:"minR",attr`
	maxR string  `xml:"maxR",attr`
	gdRefAng string  `xml:"gdRefAng",attr`
	minAng string  `xml:"minAng",attr`
	maxAng string  `xml:"maxAng",attr`
}
type CT_ConnectionSite struct{
	pos string  `xml:"pos"`
	ang string  `xml:"ang",attr`
}
type CT_AdjustHandleList struct{
	ahXY string  `xml:"ahXY"`
	ahPolar string  `xml:"ahPolar"`
}
type CT_ConnectionSiteList struct{
	cxn []CT_ConnectionSite  `xml:"cxn"`
}
type CT_Connection struct{
	id string  `xml:"id",attr`
	idx string  `xml:"idx",attr`
}
type CT_Path2DMoveTo struct{
	pt string  `xml:"pt"`
}
type CT_Path2DLineTo struct{
	pt string  `xml:"pt"`
}
type CT_Path2DArcTo struct{
	wR string  `xml:"wR",attr`
	hR string  `xml:"hR",attr`
	stAng string  `xml:"stAng",attr`
	swAng string  `xml:"swAng",attr`
}
type CT_Path2DQuadBezierTo struct{
	pt string  `xml:"pt"`
}
type CT_Path2DCubicBezierTo struct{
	pt string  `xml:"pt"`
}
type CT_Path2DClose struct{
}
type CT_Path2D struct{
	close string  `xml:"close"`
	moveTo CT_Path2DMoveTo  `xml:"moveTo"`
	lnTo CT_Path2DLineTo  `xml:"lnTo"`
	arcTo CT_Path2DArcTo  `xml:"arcTo"`
	quadBezTo string  `xml:"quadBezTo"`
	cubicBezTo CT_Path2DCubicBezierTo  `xml:"cubicBezTo"`
	w string  `xml:"w",attr`
	h string  `xml:"h",attr`
	fill string  `xml:"fill",attr`
	stroke string  `xml:"stroke",attr`
	extrusionOk string  `xml:"extrusionOk",attr`
}
type CT_Path2DList struct{
	path []CT_Path2D  `xml:"path"`
}
type CT_PresetGeometry2D struct{
	avLst string  `xml:"avLst"`
	prst ST_ShapeType  `xml:"prst",attr`
}
type CT_PresetTextShape struct{
	avLst string  `xml:"avLst"`
	prst string  `xml:"prst",attr`
}
type CT_CustomGeometry2D struct{
	avLst string  `xml:"avLst"`
	gdLst string  `xml:"gdLst"`
	ahLst string  `xml:"ahLst"`
	cxnLst string  `xml:"cxnLst"`
	rect CT_GeomRect  `xml:"rect"`
	pathLst string  `xml:"pathLst"`
}
type CT_LineEndProperties struct{
	type string  `xml:"type",attr`
	w string  `xml:"w",attr`
	len string  `xml:"len",attr`
}
type CT_LineJoinBevel struct{
}
type CT_LineJoinRound struct{
}
type CT_LineJoinMiterProperties struct{
	lim string  `xml:"lim",attr`
}
type CT_PresetLineDashProperties struct{
	val string  `xml:"val",attr`
}
type CT_DashStop struct{
	d string  `xml:"d",attr`
	sp string  `xml:"sp",attr`
}
type CT_DashStopList struct{
	ds []string  `xml:"ds"`
}
type CT_LineProperties struct{
	headEnd string  `xml:"headEnd"`
	tailEnd string  `xml:"tailEnd"`
	extLst string  `xml:"extLst"`
	 EG_LineJoinProperties  `xml:""`
	w string  `xml:"w",attr`
	cap ST_LineCap  `xml:"cap",attr`
	cmpd string  `xml:"cmpd",attr`
	algn ST_PenAlignment  `xml:"algn",attr`
}
type CT_ShapeProperties struct{
	xfrm string  `xml:"xfrm"`
	ln string  `xml:"ln"`
	scene3d CT_Scene3D  `xml:"scene3d"`
	sp3d CT_Shape3D  `xml:"sp3d"`
	extLst string  `xml:"extLst"`
	 EG_EffectProperties  `xml:""`
	bwMode string  `xml:"bwMode",attr`
}
type CT_GroupShapeProperties struct{
	xfrm string  `xml:"xfrm"`
	scene3d CT_Scene3D  `xml:"scene3d"`
	extLst string  `xml:"extLst"`
	 EG_EffectProperties  `xml:""`
	bwMode string  `xml:"bwMode",attr`
}
type CT_StyleMatrixReference struct{
	 EG_ColorChoice  `xml:""`
	idx string  `xml:"idx",attr`
}
type CT_FontReference struct{
	 EG_ColorChoice  `xml:""`
	idx string  `xml:"idx",attr`
}
type CT_ShapeStyle struct{
	lnRef string  `xml:"lnRef"`
	fillRef string  `xml:"fillRef"`
	effectRef string  `xml:"effectRef"`
	fontRef CT_FontReference  `xml:"fontRef"`
}
type CT_DefaultShapeDefinition struct{
	spPr string  `xml:"spPr"`
	bodyPr string  `xml:"bodyPr"`
	lstStyle string  `xml:"lstStyle"`
	style CT_ShapeStyle  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_ObjectStyleDefaults struct{
	spDef CT_DefaultShapeDefinition  `xml:"spDef"`
	lnDef CT_DefaultShapeDefinition  `xml:"lnDef"`
	txDef CT_DefaultShapeDefinition  `xml:"txDef"`
	extLst string  `xml:"extLst"`
}
type CT_EmptyElement struct{
}
type CT_ColorMapping struct{
	extLst string  `xml:"extLst"`
	bg1 string  `xml:"bg1",attr`
	tx1 string  `xml:"tx1",attr`
	bg2 string  `xml:"bg2",attr`
	tx2 string  `xml:"tx2",attr`
	accent1 string  `xml:"accent1",attr`
	accent2 string  `xml:"accent2",attr`
	accent3 string  `xml:"accent3",attr`
	accent4 string  `xml:"accent4",attr`
	accent5 string  `xml:"accent5",attr`
	accent6 string  `xml:"accent6",attr`
	hlink string  `xml:"hlink",attr`
	folHlink string  `xml:"folHlink",attr`
}
type CT_ColorMappingOverride struct{
	masterClrMapping CT_EmptyElement  `xml:"masterClrMapping"`
	overrideClrMapping CT_ColorMapping  `xml:"overrideClrMapping"`
}
type CT_ColorSchemeAndMapping struct{
	clrScheme CT_ColorScheme  `xml:"clrScheme"`
	clrMap CT_ColorMapping  `xml:"clrMap"`
}
type CT_ColorSchemeList struct{
	extraClrScheme []string  `xml:"extraClrScheme"`
}
type CT_OfficeStyleSheet struct{
	themeElements string  `xml:"themeElements"`
	objectDefaults string  `xml:"objectDefaults"`
	extraClrSchemeLst string  `xml:"extraClrSchemeLst"`
	custClrLst string  `xml:"custClrLst"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_BaseStylesOverride struct{
	clrScheme CT_ColorScheme  `xml:"clrScheme"`
	fontScheme CT_FontScheme  `xml:"fontScheme"`
	fmtScheme string  `xml:"fmtScheme"`
}
type CT_ClipboardStyleSheet struct{
	themeElements string  `xml:"themeElements"`
	clrMap CT_ColorMapping  `xml:"clrMap"`
}
type CT_TableCellProperties struct{
	lnL string  `xml:"lnL"`
	lnR string  `xml:"lnR"`
	lnT string  `xml:"lnT"`
	lnB string  `xml:"lnB"`
	lnTlToBr string  `xml:"lnTlToBr"`
	lnBlToTr string  `xml:"lnBlToTr"`
	cell3D CT_Cell3D  `xml:"cell3D"`
	headers string  `xml:"headers"`
	extLst string  `xml:"extLst"`
	 EG_FillProperties  `xml:""`
	marL string  `xml:"marL",attr`
	marR string  `xml:"marR",attr`
	marT string  `xml:"marT",attr`
	marB string  `xml:"marB",attr`
	vert string  `xml:"vert",attr`
	anchor string  `xml:"anchor",attr`
	anchorCtr string  `xml:"anchorCtr",attr`
	horzOverflow string  `xml:"horzOverflow",attr`
}
type CT_Headers struct{
	header string  `xml:"header"`
}
type CT_TableCol struct{
	extLst string  `xml:"extLst"`
	w string  `xml:"w",attr`
}
type CT_TableGrid struct{
	gridCol []CT_TableCol  `xml:"gridCol"`
}
type CT_TableCell struct{
	txBody string  `xml:"txBody"`
	tcPr string  `xml:"tcPr"`
	extLst string  `xml:"extLst"`
	rowSpan string  `xml:"rowSpan",attr`
	gridSpan string  `xml:"gridSpan",attr`
	hMerge string  `xml:"hMerge",attr`
	vMerge string  `xml:"vMerge",attr`
	id string  `xml:"id",attr`
}
type CT_TableRow struct{
	tc []CT_TableCell  `xml:"tc"`
	extLst string  `xml:"extLst"`
	h string  `xml:"h",attr`
}
type CT_TableProperties struct{
	extLst string  `xml:"extLst"`
	 EG_EffectProperties  `xml:""`
	tableStyle CT_TableStyle  `xml:"tableStyle"`
	tableStyleId string  `xml:"tableStyleId"`
	rtl string  `xml:"rtl",attr`
	firstRow string  `xml:"firstRow",attr`
	firstCol string  `xml:"firstCol",attr`
	lastRow string  `xml:"lastRow",attr`
	lastCol string  `xml:"lastCol",attr`
	bandRow string  `xml:"bandRow",attr`
	bandCol string  `xml:"bandCol",attr`
}
type CT_Table struct{
	tblPr string  `xml:"tblPr"`
	tblGrid string  `xml:"tblGrid"`
	tr []CT_TableRow  `xml:"tr"`
}
type CT_Cell3D struct{
	bevel CT_Bevel  `xml:"bevel"`
	lightRig CT_LightRig  `xml:"lightRig"`
	extLst string  `xml:"extLst"`
	prstMaterial string  `xml:"prstMaterial",attr`
}
type CT_ThemeableLineStyle struct{
	ln string  `xml:"ln"`
	lnRef string  `xml:"lnRef"`
}
type CT_TableStyleTextStyle struct{
	extLst string  `xml:"extLst"`
	 EG_ColorChoice  `xml:""`
	b ST_OnOffStyleType  `xml:"b",attr`
	i ST_OnOffStyleType  `xml:"i",attr`
}
type CT_TableCellBorderStyle struct{
	left CT_ThemeableLineStyle  `xml:"left"`
	right CT_ThemeableLineStyle  `xml:"right"`
	top CT_ThemeableLineStyle  `xml:"top"`
	bottom CT_ThemeableLineStyle  `xml:"bottom"`
	insideH CT_ThemeableLineStyle  `xml:"insideH"`
	insideV CT_ThemeableLineStyle  `xml:"insideV"`
	tl2br CT_ThemeableLineStyle  `xml:"tl2br"`
	tr2bl CT_ThemeableLineStyle  `xml:"tr2bl"`
	extLst string  `xml:"extLst"`
}
type CT_TableBackgroundStyle struct{
	 EG_ThemeableEffectStyle  `xml:""`
}
type CT_TableStyleCellStyle struct{
	tcBdr string  `xml:"tcBdr"`
	cell3D CT_Cell3D  `xml:"cell3D"`
	 EG_ThemeableFillStyle  `xml:""`
}
type CT_TablePartStyle struct{
	tcTxStyle string  `xml:"tcTxStyle"`
	tcStyle CT_TableStyleCellStyle  `xml:"tcStyle"`
}
type CT_TableStyle struct{
	tblBg string  `xml:"tblBg"`
	wholeTbl CT_TablePartStyle  `xml:"wholeTbl"`
	band1H CT_TablePartStyle  `xml:"band1H"`
	band2H CT_TablePartStyle  `xml:"band2H"`
	band1V CT_TablePartStyle  `xml:"band1V"`
	band2V CT_TablePartStyle  `xml:"band2V"`
	lastCol CT_TablePartStyle  `xml:"lastCol"`
	firstCol CT_TablePartStyle  `xml:"firstCol"`
	lastRow CT_TablePartStyle  `xml:"lastRow"`
	seCell CT_TablePartStyle  `xml:"seCell"`
	swCell CT_TablePartStyle  `xml:"swCell"`
	firstRow CT_TablePartStyle  `xml:"firstRow"`
	neCell CT_TablePartStyle  `xml:"neCell"`
	nwCell CT_TablePartStyle  `xml:"nwCell"`
	extLst string  `xml:"extLst"`
	styleId string  `xml:"styleId",attr`
	styleName string  `xml:"styleName",attr`
}
type CT_TableStyleList struct{
	tblStyle []CT_TableStyle  `xml:"tblStyle"`
	def string  `xml:"def",attr`
}
type CT_TextParagraph struct{
	pPr string  `xml:"pPr"`
	endParaRPr string  `xml:"endParaRPr"`
	 []EG_TextRun  `xml:""`
}
type CT_TextListStyle struct{
	defPPr string  `xml:"defPPr"`
	lvl1pPr string  `xml:"lvl1pPr"`
	lvl2pPr string  `xml:"lvl2pPr"`
	lvl3pPr string  `xml:"lvl3pPr"`
	lvl4pPr string  `xml:"lvl4pPr"`
	lvl5pPr string  `xml:"lvl5pPr"`
	lvl6pPr string  `xml:"lvl6pPr"`
	lvl7pPr string  `xml:"lvl7pPr"`
	lvl8pPr string  `xml:"lvl8pPr"`
	lvl9pPr string  `xml:"lvl9pPr"`
	extLst string  `xml:"extLst"`
}
type CT_TextNormalAutofit struct{
	fontScale string  `xml:"fontScale",attr`
	lnSpcReduction string  `xml:"lnSpcReduction",attr`
}
type CT_TextShapeAutofit struct{
}
type CT_TextNoAutofit struct{
}
type CT_TextBodyProperties struct{
	prstTxWarp string  `xml:"prstTxWarp"`
	scene3d CT_Scene3D  `xml:"scene3d"`
	extLst string  `xml:"extLst"`
	 EG_Text3D  `xml:""`
	rot ST_Angle  `xml:"rot",attr`
	spcFirstLastPara string  `xml:"spcFirstLastPara",attr`
	vertOverflow string  `xml:"vertOverflow",attr`
	horzOverflow string  `xml:"horzOverflow",attr`
	vert string  `xml:"vert",attr`
	wrap string  `xml:"wrap",attr`
	lIns string  `xml:"lIns",attr`
	tIns string  `xml:"tIns",attr`
	rIns string  `xml:"rIns",attr`
	bIns string  `xml:"bIns",attr`
	numCol string  `xml:"numCol",attr`
	spcCol string  `xml:"spcCol",attr`
	rtlCol string  `xml:"rtlCol",attr`
	fromWordArt string  `xml:"fromWordArt",attr`
	anchor string  `xml:"anchor",attr`
	anchorCtr string  `xml:"anchorCtr",attr`
	forceAA string  `xml:"forceAA",attr`
	upright string  `xml:"upright",attr`
	compatLnSpc string  `xml:"compatLnSpc",attr`
}
type CT_TextBody struct{
	bodyPr string  `xml:"bodyPr"`
	lstStyle string  `xml:"lstStyle"`
	p []string  `xml:"p"`
}
type CT_TextBulletColorFollowText struct{
}
type CT_TextBulletSizeFollowText struct{
}
type CT_TextBulletSizePercent struct{
	val string  `xml:"val",attr`
}
type CT_TextBulletSizePoint struct{
	val string  `xml:"val",attr`
}
type CT_TextBulletTypefaceFollowText struct{
}
type CT_TextAutonumberBullet struct{
	type string  `xml:"type",attr`
	startAt string  `xml:"startAt",attr`
}
type CT_TextCharBullet struct{
	char string  `xml:"char",attr`
}
type CT_TextBlipBullet struct{
	blip CT_Blip  `xml:"blip"`
}
type CT_TextNoBullet struct{
}
type CT_TextFont struct{
	typeface string  `xml:"typeface",attr`
	panose string  `xml:"panose",attr`
	pitchFamily ST_PitchFamily  `xml:"pitchFamily",attr`
	charset string  `xml:"charset",attr`
}
type CT_TextUnderlineLineFollowText struct{
}
type CT_TextUnderlineFillFollowText struct{
}
type CT_TextUnderlineFillGroupWrapper struct{
}
type CT_TextCharacterProperties struct{
	ln string  `xml:"ln"`
	highlight CT_Color  `xml:"highlight"`
	latin string  `xml:"latin"`
	ea string  `xml:"ea"`
	cs string  `xml:"cs"`
	sym string  `xml:"sym"`
	hlinkClick CT_Hyperlink  `xml:"hlinkClick"`
	hlinkMouseOver CT_Hyperlink  `xml:"hlinkMouseOver"`
	rtl CT_Boolean  `xml:"rtl"`
	extLst string  `xml:"extLst"`
	 EG_TextUnderlineFill  `xml:""`
	kumimoji string  `xml:"kumimoji",attr`
	lang string  `xml:"lang",attr`
	altLang string  `xml:"altLang",attr`
	sz string  `xml:"sz",attr`
	b string  `xml:"b",attr`
	i string  `xml:"i",attr`
	u string  `xml:"u",attr`
	strike string  `xml:"strike",attr`
	kern string  `xml:"kern",attr`
	cap string  `xml:"cap",attr`
	spc string  `xml:"spc",attr`
	normalizeH string  `xml:"normalizeH",attr`
	baseline ST_Percentage  `xml:"baseline",attr`
	noProof string  `xml:"noProof",attr`
	dirty string  `xml:"dirty",attr`
	err string  `xml:"err",attr`
	smtClean string  `xml:"smtClean",attr`
	smtId string  `xml:"smtId",attr`
	bmk string  `xml:"bmk",attr`
}
type CT_Boolean struct{
	val string  `xml:"val",attr`
}
type CT_TextSpacingPercent struct{
	val string  `xml:"val",attr`
}
type CT_TextSpacingPoint struct{
	val string  `xml:"val",attr`
}
type CT_TextTabStop struct{
	pos string  `xml:"pos",attr`
	algn string  `xml:"algn",attr`
}
type CT_TextTabStopList struct{
	tab string  `xml:"tab"`
}
type CT_TextLineBreak struct{
	rPr string  `xml:"rPr"`
}
type CT_TextSpacing struct{
	spcPct string  `xml:"spcPct"`
	spcPts string  `xml:"spcPts"`
}
type CT_TextParagraphProperties struct{
	lnSpc string  `xml:"lnSpc"`
	spcBef string  `xml:"spcBef"`
	spcAft string  `xml:"spcAft"`
	tabLst string  `xml:"tabLst"`
	defRPr string  `xml:"defRPr"`
	extLst string  `xml:"extLst"`
	 EG_TextBullet  `xml:""`
	marL string  `xml:"marL",attr`
	marR string  `xml:"marR",attr`
	lvl string  `xml:"lvl",attr`
	indent string  `xml:"indent",attr`
	algn string  `xml:"algn",attr`
	defTabSz string  `xml:"defTabSz",attr`
	rtl string  `xml:"rtl",attr`
	eaLnBrk string  `xml:"eaLnBrk",attr`
	fontAlgn string  `xml:"fontAlgn",attr`
	latinLnBrk string  `xml:"latinLnBrk",attr`
	hangingPunct string  `xml:"hangingPunct",attr`
}
type CT_TextField struct{
	rPr string  `xml:"rPr"`
	pPr string  `xml:"pPr"`
	t string  `xml:"t"`
	id string  `xml:"id",attr`
	type string  `xml:"type",attr`
}
type CT_RegularTextRun struct{
	rPr string  `xml:"rPr"`
	t string  `xml:"t"`
}
