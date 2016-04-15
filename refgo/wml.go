package ooxml
type WordSchema  struct {
	recipients string  `xml:"recipients"`
	txbxContent string  `xml:"txbxContent"`
	comments string  `xml:"comments"`
	footnotes string  `xml:"footnotes"`
	endnotes string  `xml:"endnotes"`
	hdr string  `xml:"hdr"`
	ftr string  `xml:"ftr"`
	settings string  `xml:"settings"`
	webSettings string  `xml:"webSettings"`
	fonts string  `xml:"fonts"`
	numbering CT_Numbering  `xml:"numbering"`
	styles string  `xml:"styles"`
	document CT_Document  `xml:"document"`
	glossaryDocument string  `xml:"glossaryDocument"`
}
type CT_Empty struct{
}
type CT_OnOff struct{
	val string  `xml:"val",attr`
}
type CT_LongHexNumber struct{
	val string  `xml:"val",attr`
}
type CT_Charset struct{
	val string  `xml:"val",attr`
	characterSet string  `xml:"characterSet",attr`
}
type CT_DecimalNumber struct{
	val ST_DecimalNumber  `xml:"val",attr`
}
type CT_UnsignedDecimalNumber struct{
	val string  `xml:"val",attr`
}
type CT_DecimalNumberOrPrecent struct{
	val ST_DecimalNumberOrPercent  `xml:"val",attr`
}
type CT_TwipsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_SignedTwipsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_PixelsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_HpsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_SignedHpsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_MacroName struct{
	val ST_MacroName  `xml:"val",attr`
}
type CT_String struct{
	val string  `xml:"val",attr`
}
type CT_TextScale struct{
	val string  `xml:"val",attr`
}
type CT_Highlight struct{
	val ST_HighlightColor  `xml:"val",attr`
}
type CT_Color struct{
	val string  `xml:"val",attr`
	themeColor ST_ThemeColor  `xml:"themeColor",attr`
	themeTint string  `xml:"themeTint",attr`
	themeShade string  `xml:"themeShade",attr`
}
type CT_Lang struct{
	val string  `xml:"val",attr`
}
type CT_Guid struct{
	val string  `xml:"val",attr`
}
type CT_Underline struct{
	val string  `xml:"val",attr`
	color string  `xml:"color",attr`
	themeColor ST_ThemeColor  `xml:"themeColor",attr`
	themeTint string  `xml:"themeTint",attr`
	themeShade string  `xml:"themeShade",attr`
}
type CT_TextEffect struct{
	val string  `xml:"val",attr`
}
type CT_Border struct{
	val string  `xml:"val",attr`
	color string  `xml:"color",attr`
	themeColor ST_ThemeColor  `xml:"themeColor",attr`
	themeTint string  `xml:"themeTint",attr`
	themeShade string  `xml:"themeShade",attr`
	sz string  `xml:"sz",attr`
	space string  `xml:"space",attr`
	shadow string  `xml:"shadow",attr`
	frame string  `xml:"frame",attr`
}
type CT_Shd struct{
	val string  `xml:"val",attr`
	color string  `xml:"color",attr`
	themeColor ST_ThemeColor  `xml:"themeColor",attr`
	themeTint string  `xml:"themeTint",attr`
	themeShade string  `xml:"themeShade",attr`
	fill string  `xml:"fill",attr`
	themeFill ST_ThemeColor  `xml:"themeFill",attr`
	themeFillTint string  `xml:"themeFillTint",attr`
	themeFillShade string  `xml:"themeFillShade",attr`
}
type CT_VerticalAlignRun struct{
	val string  `xml:"val",attr`
}
type CT_FitText struct{
	val string  `xml:"val",attr`
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_Em struct{
	val ST_Em  `xml:"val",attr`
}
type CT_Language struct{
	val string  `xml:"val",attr`
	eastAsia string  `xml:"eastAsia",attr`
	bidi string  `xml:"bidi",attr`
}
type CT_EastAsianLayout struct{
	id ST_DecimalNumber  `xml:"id",attr`
	combine string  `xml:"combine",attr`
	combineBrackets string  `xml:"combineBrackets",attr`
	vert string  `xml:"vert",attr`
	vertCompress string  `xml:"vertCompress",attr`
}
type CT_FramePr struct{
	dropCap ST_DropCap  `xml:"dropCap",attr`
	lines ST_DecimalNumber  `xml:"lines",attr`
	w string  `xml:"w",attr`
	h string  `xml:"h",attr`
	vSpace string  `xml:"vSpace",attr`
	hSpace string  `xml:"hSpace",attr`
	wrap ST_Wrap  `xml:"wrap",attr`
	hAnchor ST_HAnchor  `xml:"hAnchor",attr`
	vAnchor ST_VAnchor  `xml:"vAnchor",attr`
	x string  `xml:"x",attr`
	xAlign string  `xml:"xAlign",attr`
	y string  `xml:"y",attr`
	yAlign string  `xml:"yAlign",attr`
	hRule ST_HeightRule  `xml:"hRule",attr`
	anchorLock string  `xml:"anchorLock",attr`
}
type CT_TabStop struct{
	val ST_TabJc  `xml:"val",attr`
	leader ST_TabTlc  `xml:"leader",attr`
	pos string  `xml:"pos",attr`
}
type CT_Spacing struct{
	before string  `xml:"before",attr`
	beforeLines ST_DecimalNumber  `xml:"beforeLines",attr`
	beforeAutospacing string  `xml:"beforeAutospacing",attr`
	after string  `xml:"after",attr`
	afterLines ST_DecimalNumber  `xml:"afterLines",attr`
	afterAutospacing string  `xml:"afterAutospacing",attr`
	line string  `xml:"line",attr`
	lineRule ST_LineSpacingRule  `xml:"lineRule",attr`
}
type CT_Ind struct{
	start string  `xml:"start",attr`
	startChars ST_DecimalNumber  `xml:"startChars",attr`
	end string  `xml:"end",attr`
	endChars ST_DecimalNumber  `xml:"endChars",attr`
	left string  `xml:"left",attr`
	leftChars ST_DecimalNumber  `xml:"leftChars",attr`
	right string  `xml:"right",attr`
	rightChars ST_DecimalNumber  `xml:"rightChars",attr`
	hanging string  `xml:"hanging",attr`
	hangingChars ST_DecimalNumber  `xml:"hangingChars",attr`
	firstLine string  `xml:"firstLine",attr`
	firstLineChars ST_DecimalNumber  `xml:"firstLineChars",attr`
}
type CT_Jc struct{
	val ST_Jc  `xml:"val",attr`
}
type CT_JcTable struct{
	val ST_JcTable  `xml:"val",attr`
}
type CT_View struct{
	val ST_View  `xml:"val",attr`
}
type CT_Zoom struct{
	val ST_Zoom  `xml:"val",attr`
	percent ST_DecimalNumberOrPercent  `xml:"percent",attr`
}
type CT_WritingStyle struct{
	lang string  `xml:"lang",attr`
	vendorID string  `xml:"vendorID",attr`
	dllVersion string  `xml:"dllVersion",attr`
	nlCheck string  `xml:"nlCheck",attr`
	checkStyle string  `xml:"checkStyle",attr`
	appName string  `xml:"appName",attr`
}
type CT_Proof struct{
	spelling ST_Proof  `xml:"spelling",attr`
	grammar ST_Proof  `xml:"grammar",attr`
}
type CT_DocType struct{
	val ST_DocType  `xml:"val",attr`
}
type CT_DocProtect struct{
	edit ST_DocProtect  `xml:"edit",attr`
	formatting string  `xml:"formatting",attr`
	enforcement string  `xml:"enforcement",attr`
}
type CT_MailMergeDocType struct{
	val ST_MailMergeDocType  `xml:"val",attr`
}
type CT_MailMergeDataType struct{
	val ST_MailMergeDataType  `xml:"val",attr`
}
type CT_MailMergeDest struct{
	val string  `xml:"val",attr`
}
type CT_MailMergeOdsoFMDFieldType struct{
	val string  `xml:"val",attr`
}
type CT_TrackChangesView struct{
	markup string  `xml:"markup",attr`
	comments string  `xml:"comments",attr`
	insDel string  `xml:"insDel",attr`
	formatting string  `xml:"formatting",attr`
	inkAnnotations string  `xml:"inkAnnotations",attr`
}
type CT_Kinsoku struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_TextDirection struct{
	val string  `xml:"val",attr`
}
type CT_TextAlignment struct{
	val string  `xml:"val",attr`
}
type CT_Markup struct{
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_TrackChange struct{
	CT_Markup
	author string  `xml:"author",attr`
	date ST_DateTime  `xml:"date",attr`
}
type CT_CellMergeTrackChange struct{
	CT_TrackChange
	vMerge ST_AnnotationVMerge  `xml:"vMerge",attr`
	vMergeOrig ST_AnnotationVMerge  `xml:"vMergeOrig",attr`
}
type CT_TrackChangeRange struct{
	CT_TrackChange
	displacedByCustomXml string  `xml:"displacedByCustomXml",attr`
}
type CT_MarkupRange struct{
	CT_Markup
	displacedByCustomXml string  `xml:"displacedByCustomXml",attr`
}
type CT_BookmarkRange struct{
	CT_MarkupRange
	colFirst ST_DecimalNumber  `xml:"colFirst",attr`
	colLast ST_DecimalNumber  `xml:"colLast",attr`
}
type CT_Bookmark struct{
	CT_BookmarkRange
	name string  `xml:"name",attr`
}
type CT_MoveBookmark struct{
	CT_Bookmark
	author string  `xml:"author",attr`
	date ST_DateTime  `xml:"date",attr`
}
type CT_Comment struct{
	CT_TrackChange
	 []EG_BlockLevelElts  `xml:""`
	initials string  `xml:"initials",attr`
}
type CT_TrackChangeNumbering struct{
	CT_TrackChange
	original string  `xml:"original",attr`
}
type CT_TblPrExChange struct{
	CT_TrackChange
	tblPrEx string  `xml:"tblPrEx"`
}
type CT_TcPrChange struct{
	CT_TrackChange
	tcPr CT_TcPrInner  `xml:"tcPr"`
}
type CT_TrPrChange struct{
	CT_TrackChange
	trPr string  `xml:"trPr"`
}
type CT_TblGridChange struct{
	CT_Markup
	tblGrid string  `xml:"tblGrid"`
}
type CT_TblPrChange struct{
	CT_TrackChange
	tblPr string  `xml:"tblPr"`
}
type CT_SectPrChange struct{
	CT_TrackChange
	sectPr string  `xml:"sectPr"`
}
type CT_PPrChange struct{
	CT_TrackChange
	pPr string  `xml:"pPr"`
}
type CT_RPrChange struct{
	CT_TrackChange
	rPr CT_RPrOriginal  `xml:"rPr"`
}
type CT_ParaRPrChange struct{
	CT_TrackChange
	rPr CT_ParaRPrOriginal  `xml:"rPr"`
}
type CT_RunTrackChange struct{
	CT_TrackChange
	 EG_ContentRunContent  `xml:""`
	 m:EG_OMathMathElements  `xml:""`
}
type CT_NumPr struct{
	ilvl CT_DecimalNumber  `xml:"ilvl"`
	numId CT_DecimalNumber  `xml:"numId"`
	numberingChange CT_TrackChangeNumbering  `xml:"numberingChange"`
	ins CT_TrackChange  `xml:"ins"`
}
type CT_PBdr struct{
	top string  `xml:"top"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	right string  `xml:"right"`
	between string  `xml:"between"`
	bar string  `xml:"bar"`
}
type CT_Tabs struct{
	tab []CT_TabStop  `xml:"tab"`
}
type CT_TextboxTightWrap struct{
	val string  `xml:"val",attr`
}
type CT_PPr struct{
	CT_PPrBase
	rPr CT_ParaRPr  `xml:"rPr"`
	sectPr CT_SectPr  `xml:"sectPr"`
	pPrChange CT_PPrChange  `xml:"pPrChange"`
}
type CT_PPrBase struct{
	pStyle CT_String  `xml:"pStyle"`
	keepNext CT_OnOff  `xml:"keepNext"`
	keepLines CT_OnOff  `xml:"keepLines"`
	pageBreakBefore CT_OnOff  `xml:"pageBreakBefore"`
	framePr CT_FramePr  `xml:"framePr"`
	widowControl CT_OnOff  `xml:"widowControl"`
	numPr CT_NumPr  `xml:"numPr"`
	suppressLineNumbers CT_OnOff  `xml:"suppressLineNumbers"`
	pBdr string  `xml:"pBdr"`
	shd string  `xml:"shd"`
	tabs string  `xml:"tabs"`
	suppressAutoHyphens CT_OnOff  `xml:"suppressAutoHyphens"`
	kinsoku CT_OnOff  `xml:"kinsoku"`
	wordWrap CT_OnOff  `xml:"wordWrap"`
	overflowPunct CT_OnOff  `xml:"overflowPunct"`
	topLinePunct CT_OnOff  `xml:"topLinePunct"`
	autoSpaceDE CT_OnOff  `xml:"autoSpaceDE"`
	autoSpaceDN CT_OnOff  `xml:"autoSpaceDN"`
	bidi CT_OnOff  `xml:"bidi"`
	adjustRightInd CT_OnOff  `xml:"adjustRightInd"`
	snapToGrid CT_OnOff  `xml:"snapToGrid"`
	spacing CT_Spacing  `xml:"spacing"`
	ind string  `xml:"ind"`
	contextualSpacing CT_OnOff  `xml:"contextualSpacing"`
	mirrorIndents CT_OnOff  `xml:"mirrorIndents"`
	suppressOverlap CT_OnOff  `xml:"suppressOverlap"`
	jc CT_Jc  `xml:"jc"`
	textDirection string  `xml:"textDirection"`
	textAlignment string  `xml:"textAlignment"`
	textboxTightWrap string  `xml:"textboxTightWrap"`
	outlineLvl CT_DecimalNumber  `xml:"outlineLvl"`
	divId CT_DecimalNumber  `xml:"divId"`
	cnfStyle CT_Cnf  `xml:"cnfStyle"`
}
type CT_PPrGeneral struct{
	CT_PPrBase
	pPrChange CT_PPrChange  `xml:"pPrChange"`
}
type CT_Control struct{
	name string  `xml:"name",attr`
	shapeid string  `xml:"shapeid",attr`
}
type CT_Background struct{
	drawing CT_Drawing  `xml:"drawing"`
	color string  `xml:"color",attr`
	themeColor ST_ThemeColor  `xml:"themeColor",attr`
	themeTint string  `xml:"themeTint",attr`
	themeShade string  `xml:"themeShade",attr`
}
type CT_Rel struct{
}
type CT_Object struct{
	drawing CT_Drawing  `xml:"drawing"`
	control CT_Control  `xml:"control"`
	objectLink CT_ObjectLink  `xml:"objectLink"`
	objectEmbed string  `xml:"objectEmbed"`
	movie CT_Rel  `xml:"movie"`
	dxaOrig string  `xml:"dxaOrig",attr`
	dyaOrig string  `xml:"dyaOrig",attr`
}
type CT_Picture struct{
	movie CT_Rel  `xml:"movie"`
	control CT_Control  `xml:"control"`
}
type CT_ObjectEmbed struct{
	drawAspect string  `xml:"drawAspect",attr`
	progId string  `xml:"progId",attr`
	shapeId string  `xml:"shapeId",attr`
	fieldCodes string  `xml:"fieldCodes",attr`
}
type CT_ObjectLink struct{
	CT_ObjectEmbed
	updateMode string  `xml:"updateMode",attr`
	lockedField string  `xml:"lockedField",attr`
}
type CT_Drawing struct{
}
type CT_SimpleField struct{
	fldData string  `xml:"fldData"`
	 []EG_PContent  `xml:""`
	instr string  `xml:"instr",attr`
	fldLock string  `xml:"fldLock",attr`
	dirty string  `xml:"dirty",attr`
}
type CT_FFTextType struct{
	val string  `xml:"val",attr`
}
type CT_FFName struct{
	val ST_FFName  `xml:"val",attr`
}
type CT_FldChar struct{
	fldData string  `xml:"fldData"`
	ffData CT_FFData  `xml:"ffData"`
	numberingChange CT_TrackChangeNumbering  `xml:"numberingChange"`
	fldCharType string  `xml:"fldCharType",attr`
	fldLock string  `xml:"fldLock",attr`
	dirty string  `xml:"dirty",attr`
}
type CT_Hyperlink struct{
	tgtFrame string  `xml:"tgtFrame",attr`
	tooltip string  `xml:"tooltip",attr`
	docLocation string  `xml:"docLocation",attr`
	history string  `xml:"history",attr`
	anchor string  `xml:"anchor",attr`
}
type CT_FFData struct{
	name CT_FFName  `xml:"name"`
	label CT_DecimalNumber  `xml:"label"`
	tabIndex string  `xml:"tabIndex"`
	enabled CT_OnOff  `xml:"enabled"`
	calcOnExit CT_OnOff  `xml:"calcOnExit"`
	entryMacro CT_MacroName  `xml:"entryMacro"`
	exitMacro CT_MacroName  `xml:"exitMacro"`
	helpText string  `xml:"helpText"`
	statusText string  `xml:"statusText"`
}
type CT_FFHelpText struct{
	type string  `xml:"type",attr`
	val string  `xml:"val",attr`
}
type CT_FFStatusText struct{
	type string  `xml:"type",attr`
	val string  `xml:"val",attr`
}
type CT_FFCheckBox struct{
	default CT_OnOff  `xml:"default"`
	checked CT_OnOff  `xml:"checked"`
	size string  `xml:"size"`
	sizeAuto CT_OnOff  `xml:"sizeAuto"`
}
type CT_FFDDList struct{
	result CT_DecimalNumber  `xml:"result"`
	default CT_DecimalNumber  `xml:"default"`
	listEntry []CT_String  `xml:"listEntry"`
}
type CT_FFTextInput struct{
	type string  `xml:"type"`
	default CT_String  `xml:"default"`
	maxLength CT_DecimalNumber  `xml:"maxLength"`
	format CT_String  `xml:"format"`
}
type CT_SectType struct{
	val ST_SectionMark  `xml:"val",attr`
}
type CT_PaperSource struct{
	first ST_DecimalNumber  `xml:"first",attr`
	other ST_DecimalNumber  `xml:"other",attr`
}
type CT_PageSz struct{
	w string  `xml:"w",attr`
	h string  `xml:"h",attr`
	orient ST_PageOrientation  `xml:"orient",attr`
	code ST_DecimalNumber  `xml:"code",attr`
}
type CT_PageMar struct{
	top string  `xml:"top",attr`
	right string  `xml:"right",attr`
	bottom string  `xml:"bottom",attr`
	left string  `xml:"left",attr`
	header string  `xml:"header",attr`
	footer string  `xml:"footer",attr`
	gutter string  `xml:"gutter",attr`
}
type CT_PageBorders struct{
	top string  `xml:"top"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	right string  `xml:"right"`
	zOrder string  `xml:"zOrder",attr`
	display string  `xml:"display",attr`
	offsetFrom string  `xml:"offsetFrom",attr`
}
type CT_PageBorder struct{
	CT_Border
}
type CT_BottomPageBorder struct{
	CT_PageBorder
}
type CT_TopPageBorder struct{
	CT_PageBorder
}
type CT_LineNumber struct{
	countBy ST_DecimalNumber  `xml:"countBy",attr`
	start ST_DecimalNumber  `xml:"start",attr`
	distance string  `xml:"distance",attr`
	restart string  `xml:"restart",attr`
}
type CT_PageNumber struct{
	fmt ST_NumberFormat  `xml:"fmt",attr`
	start ST_DecimalNumber  `xml:"start",attr`
	chapStyle ST_DecimalNumber  `xml:"chapStyle",attr`
	chapSep ST_ChapterSep  `xml:"chapSep",attr`
}
type CT_Column struct{
	w string  `xml:"w",attr`
	space string  `xml:"space",attr`
}
type CT_Columns struct{
	col CT_Column  `xml:"col"`
	equalWidth string  `xml:"equalWidth",attr`
	space string  `xml:"space",attr`
	num ST_DecimalNumber  `xml:"num",attr`
	sep string  `xml:"sep",attr`
}
type CT_VerticalJc struct{
	val ST_VerticalJc  `xml:"val",attr`
}
type CT_DocGrid struct{
	type string  `xml:"type",attr`
	linePitch ST_DecimalNumber  `xml:"linePitch",attr`
	charSpace ST_DecimalNumber  `xml:"charSpace",attr`
}
type CT_HdrFtrRef struct{
	CT_Rel
	type string  `xml:"type",attr`
}
type CT_HdrFtr struct{
}
type CT_SectPrBase struct{
	 EG_SectPrContents  `xml:""`
}
type CT_SectPr struct{
	sectPrChange CT_SectPrChange  `xml:"sectPrChange"`
	 EG_SectPrContents  `xml:""`
}
type CT_Br struct{
	type ST_BrType  `xml:"type",attr`
	clear ST_BrClear  `xml:"clear",attr`
}
type CT_PTab struct{
	alignment ST_PTabAlignment  `xml:"alignment",attr`
	relativeTo ST_PTabRelativeTo  `xml:"relativeTo",attr`
	leader string  `xml:"leader",attr`
}
type CT_Sym struct{
	font string  `xml:"font",attr`
	char string  `xml:"char",attr`
}
type CT_ProofErr struct{
	type ST_ProofErr  `xml:"type",attr`
}
type CT_Perm struct{
	id string  `xml:"id",attr`
	displacedByCustomXml string  `xml:"displacedByCustomXml",attr`
}
type CT_PermStart struct{
	CT_Perm
	edGrp string  `xml:"edGrp",attr`
	ed string  `xml:"ed",attr`
	colFirst ST_DecimalNumber  `xml:"colFirst",attr`
	colLast ST_DecimalNumber  `xml:"colLast",attr`
}
type CT_Text struct{
}
type CT_R struct{
	 []EG_RunInnerContent  `xml:""`
	rsidRPr string  `xml:"rsidRPr",attr`
	rsidDel string  `xml:"rsidDel",attr`
	rsidR string  `xml:"rsidR",attr`
}
type CT_Fonts struct{
	hint ST_Hint  `xml:"hint",attr`
	ascii string  `xml:"ascii",attr`
	hAnsi string  `xml:"hAnsi",attr`
	eastAsia string  `xml:"eastAsia",attr`
	cs string  `xml:"cs",attr`
	asciiTheme ST_Theme  `xml:"asciiTheme",attr`
	hAnsiTheme ST_Theme  `xml:"hAnsiTheme",attr`
	eastAsiaTheme ST_Theme  `xml:"eastAsiaTheme",attr`
	cstheme ST_Theme  `xml:"cstheme",attr`
}
type CT_RPr struct{
	 EG_RPrContent  `xml:""`
}
type CT_MathCtrlIns struct{
	CT_TrackChange
	del CT_RPrChange  `xml:"del"`
	rPr CT_RPr  `xml:"rPr"`
}
type CT_MathCtrlDel struct{
	CT_TrackChange
	rPr CT_RPr  `xml:"rPr"`
}
type CT_RPrOriginal struct{
	 []EG_RPrBase  `xml:""`
}
type CT_ParaRPrOriginal struct{
	 []EG_RPrBase  `xml:""`
}
type CT_ParaRPr struct{
	rPrChange CT_ParaRPrChange  `xml:"rPrChange"`
	 []EG_RPrBase  `xml:""`
}
type CT_AltChunk struct{
	altChunkPr CT_AltChunkPr  `xml:"altChunkPr"`
}
type CT_AltChunkPr struct{
	matchSrc CT_OnOff  `xml:"matchSrc"`
}
type CT_RubyAlign struct{
	val ST_RubyAlign  `xml:"val",attr`
}
type CT_RubyPr struct{
	rubyAlign CT_RubyAlign  `xml:"rubyAlign"`
	hps string  `xml:"hps"`
	hpsRaise string  `xml:"hpsRaise"`
	hpsBaseText string  `xml:"hpsBaseText"`
	lid CT_Lang  `xml:"lid"`
	dirty CT_OnOff  `xml:"dirty"`
}
type CT_RubyContent struct{
}
type CT_Ruby struct{
	rubyPr CT_RubyPr  `xml:"rubyPr"`
	rt CT_RubyContent  `xml:"rt"`
	rubyBase CT_RubyContent  `xml:"rubyBase"`
}
type CT_Lock struct{
	val ST_Lock  `xml:"val",attr`
}
type CT_SdtListItem struct{
	displayText string  `xml:"displayText",attr`
	value string  `xml:"value",attr`
}
type CT_SdtDateMappingType struct{
	val string  `xml:"val",attr`
}
type CT_CalendarType struct{
	val string  `xml:"val",attr`
}
type CT_SdtDate struct{
	dateFormat CT_String  `xml:"dateFormat"`
	lid CT_Lang  `xml:"lid"`
	storeMappedDataAs string  `xml:"storeMappedDataAs"`
	calendar string  `xml:"calendar"`
	fullDate ST_DateTime  `xml:"fullDate",attr`
}
type CT_SdtComboBox struct{
	listItem []string  `xml:"listItem"`
	lastValue string  `xml:"lastValue",attr`
}
type CT_SdtDocPart struct{
	docPartGallery CT_String  `xml:"docPartGallery"`
	docPartCategory CT_String  `xml:"docPartCategory"`
	docPartUnique CT_OnOff  `xml:"docPartUnique"`
}
type CT_SdtDropDownList struct{
	listItem []string  `xml:"listItem"`
	lastValue string  `xml:"lastValue",attr`
}
type CT_Placeholder struct{
	docPart CT_String  `xml:"docPart"`
}
type CT_SdtText struct{
	multiLine string  `xml:"multiLine",attr`
}
type CT_DataBinding struct{
	prefixMappings string  `xml:"prefixMappings",attr`
	xpath string  `xml:"xpath",attr`
	storeItemID string  `xml:"storeItemID",attr`
}
type CT_SdtPr struct{
	rPr CT_RPr  `xml:"rPr"`
	alias CT_String  `xml:"alias"`
	tag CT_String  `xml:"tag"`
	id CT_DecimalNumber  `xml:"id"`
	lock CT_Lock  `xml:"lock"`
	placeholder string  `xml:"placeholder"`
	temporary CT_OnOff  `xml:"temporary"`
	showingPlcHdr CT_OnOff  `xml:"showingPlcHdr"`
	dataBinding string  `xml:"dataBinding"`
	label CT_DecimalNumber  `xml:"label"`
	tabIndex string  `xml:"tabIndex"`
	equation CT_Empty  `xml:"equation"`
	comboBox string  `xml:"comboBox"`
	date string  `xml:"date"`
	docPartObj string  `xml:"docPartObj"`
	docPartList string  `xml:"docPartList"`
	dropDownList string  `xml:"dropDownList"`
	picture CT_Empty  `xml:"picture"`
	richText CT_Empty  `xml:"richText"`
	text string  `xml:"text"`
	citation CT_Empty  `xml:"citation"`
	group CT_Empty  `xml:"group"`
	bibliography CT_Empty  `xml:"bibliography"`
}
type CT_SdtEndPr struct{
	rPr CT_RPr  `xml:"rPr"`
}
type CT_DirContentRun struct{
	val ST_Direction  `xml:"val",attr`
}
type CT_BdoContentRun struct{
	val ST_Direction  `xml:"val",attr`
}
type CT_SdtContentRun struct{
}
type CT_SdtContentBlock struct{
}
type CT_SdtContentRow struct{
}
type CT_SdtContentCell struct{
}
type CT_SdtBlock struct{
	sdtPr string  `xml:"sdtPr"`
	sdtEndPr string  `xml:"sdtEndPr"`
	sdtContent string  `xml:"sdtContent"`
}
type CT_SdtRun struct{
	sdtPr string  `xml:"sdtPr"`
	sdtEndPr string  `xml:"sdtEndPr"`
	sdtContent string  `xml:"sdtContent"`
}
type CT_SdtCell struct{
	sdtPr string  `xml:"sdtPr"`
	sdtEndPr string  `xml:"sdtEndPr"`
	sdtContent string  `xml:"sdtContent"`
}
type CT_SdtRow struct{
	sdtPr string  `xml:"sdtPr"`
	sdtEndPr string  `xml:"sdtEndPr"`
	sdtContent string  `xml:"sdtContent"`
}
type CT_Attr struct{
	uri string  `xml:"uri",attr`
	name string  `xml:"name",attr`
	val string  `xml:"val",attr`
}
type CT_CustomXmlRun struct{
	customXmlPr string  `xml:"customXmlPr"`
	 []EG_PContent  `xml:""`
	uri string  `xml:"uri",attr`
	element string  `xml:"element",attr`
}
type CT_SmartTagRun struct{
	smartTagPr CT_SmartTagPr  `xml:"smartTagPr"`
	 []EG_PContent  `xml:""`
	uri string  `xml:"uri",attr`
	element string  `xml:"element",attr`
}
type CT_CustomXmlBlock struct{
	customXmlPr string  `xml:"customXmlPr"`
	 []EG_ContentBlockContent  `xml:""`
	uri string  `xml:"uri",attr`
	element string  `xml:"element",attr`
}
type CT_CustomXmlPr struct{
	placeholder CT_String  `xml:"placeholder"`
	attr []CT_Attr  `xml:"attr"`
}
type CT_CustomXmlRow struct{
	customXmlPr string  `xml:"customXmlPr"`
	 []EG_ContentRowContent  `xml:""`
	uri string  `xml:"uri",attr`
	element string  `xml:"element",attr`
}
type CT_CustomXmlCell struct{
	customXmlPr string  `xml:"customXmlPr"`
	 []EG_ContentCellContent  `xml:""`
	uri string  `xml:"uri",attr`
	element string  `xml:"element",attr`
}
type CT_SmartTagPr struct{
	attr []CT_Attr  `xml:"attr"`
}
type CT_P struct{
	pPr CT_PPr  `xml:"pPr"`
	 []EG_PContent  `xml:""`
	rsidRPr string  `xml:"rsidRPr",attr`
	rsidR string  `xml:"rsidR",attr`
	rsidDel string  `xml:"rsidDel",attr`
	rsidP string  `xml:"rsidP",attr`
	rsidRDefault string  `xml:"rsidRDefault",attr`
}
type CT_Height struct{
	val string  `xml:"val",attr`
	hRule ST_HeightRule  `xml:"hRule",attr`
}
type CT_TblWidth struct{
	w string  `xml:"w",attr`
	type string  `xml:"type",attr`
}
type CT_TblGridCol struct{
	w string  `xml:"w",attr`
}
type CT_TblGridBase struct{
	gridCol []string  `xml:"gridCol"`
}
type CT_TblGrid struct{
	CT_TblGridBase
	tblGridChange string  `xml:"tblGridChange"`
}
type CT_TcBorders struct{
	top string  `xml:"top"`
	start string  `xml:"start"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	end string  `xml:"end"`
	right string  `xml:"right"`
	insideH string  `xml:"insideH"`
	insideV string  `xml:"insideV"`
	tl2br string  `xml:"tl2br"`
	tr2bl string  `xml:"tr2bl"`
}
type CT_TcMar struct{
	top string  `xml:"top"`
	start string  `xml:"start"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	end string  `xml:"end"`
	right string  `xml:"right"`
}
type CT_VMerge struct{
	val ST_Merge  `xml:"val",attr`
}
type CT_HMerge struct{
	val ST_Merge  `xml:"val",attr`
}
type CT_TcPrBase struct{
	cnfStyle CT_Cnf  `xml:"cnfStyle"`
	tcW string  `xml:"tcW"`
	gridSpan CT_DecimalNumber  `xml:"gridSpan"`
	hMerge CT_HMerge  `xml:"hMerge"`
	vMerge CT_VMerge  `xml:"vMerge"`
	tcBorders string  `xml:"tcBorders"`
	shd string  `xml:"shd"`
	noWrap CT_OnOff  `xml:"noWrap"`
	tcMar CT_TcMar  `xml:"tcMar"`
	textDirection string  `xml:"textDirection"`
	tcFitText CT_OnOff  `xml:"tcFitText"`
	vAlign CT_VerticalJc  `xml:"vAlign"`
	hideMark CT_OnOff  `xml:"hideMark"`
	headers string  `xml:"headers"`
}
type CT_TcPr struct{
	CT_TcPrInner
	tcPrChange CT_TcPrChange  `xml:"tcPrChange"`
}
type CT_TcPrInner struct{
	CT_TcPrBase
	 EG_CellMarkupElements  `xml:""`
}
type CT_Tc struct{
	tcPr CT_TcPr  `xml:"tcPr"`
	 []EG_BlockLevelElts  `xml:""`
	id string  `xml:"id",attr`
}
type CT_Cnf struct{
	val ST_Cnf  `xml:"val",attr`
	firstRow string  `xml:"firstRow",attr`
	lastRow string  `xml:"lastRow",attr`
	firstColumn string  `xml:"firstColumn",attr`
	lastColumn string  `xml:"lastColumn",attr`
	oddVBand string  `xml:"oddVBand",attr`
	evenVBand string  `xml:"evenVBand",attr`
	oddHBand string  `xml:"oddHBand",attr`
	evenHBand string  `xml:"evenHBand",attr`
	firstRowFirstColumn string  `xml:"firstRowFirstColumn",attr`
	firstRowLastColumn string  `xml:"firstRowLastColumn",attr`
	lastRowFirstColumn string  `xml:"lastRowFirstColumn",attr`
	lastRowLastColumn string  `xml:"lastRowLastColumn",attr`
}
type CT_Headers struct{
	header CT_String  `xml:"header"`
}
type CT_TrPrBase struct{
	cnfStyle CT_Cnf  `xml:"cnfStyle"`
	divId CT_DecimalNumber  `xml:"divId"`
	gridBefore CT_DecimalNumber  `xml:"gridBefore"`
	gridAfter CT_DecimalNumber  `xml:"gridAfter"`
	wBefore string  `xml:"wBefore"`
	wAfter string  `xml:"wAfter"`
	cantSplit CT_OnOff  `xml:"cantSplit"`
	trHeight CT_Height  `xml:"trHeight"`
	tblHeader CT_OnOff  `xml:"tblHeader"`
	tblCellSpacing string  `xml:"tblCellSpacing"`
	jc CT_JcTable  `xml:"jc"`
	hidden CT_OnOff  `xml:"hidden"`
}
type CT_TrPr struct{
	CT_TrPrBase
	ins CT_TrackChange  `xml:"ins"`
	del CT_TrackChange  `xml:"del"`
	trPrChange CT_TrPrChange  `xml:"trPrChange"`
}
type CT_Row struct{
	tblPrEx string  `xml:"tblPrEx"`
	trPr CT_TrPr  `xml:"trPr"`
	 []EG_ContentCellContent  `xml:""`
	rsidRPr string  `xml:"rsidRPr",attr`
	rsidR string  `xml:"rsidR",attr`
	rsidDel string  `xml:"rsidDel",attr`
	rsidTr string  `xml:"rsidTr",attr`
}
type CT_TblLayoutType struct{
	type ST_TblLayoutType  `xml:"type",attr`
}
type CT_TblOverlap struct{
	val ST_TblOverlap  `xml:"val",attr`
}
type CT_TblPPr struct{
	leftFromText string  `xml:"leftFromText",attr`
	rightFromText string  `xml:"rightFromText",attr`
	topFromText string  `xml:"topFromText",attr`
	bottomFromText string  `xml:"bottomFromText",attr`
	vertAnchor ST_VAnchor  `xml:"vertAnchor",attr`
	horzAnchor ST_HAnchor  `xml:"horzAnchor",attr`
	tblpXSpec string  `xml:"tblpXSpec",attr`
	tblpX string  `xml:"tblpX",attr`
	tblpYSpec string  `xml:"tblpYSpec",attr`
	tblpY string  `xml:"tblpY",attr`
}
type CT_TblCellMar struct{
	top string  `xml:"top"`
	start string  `xml:"start"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	end string  `xml:"end"`
	right string  `xml:"right"`
}
type CT_TblBorders struct{
	top string  `xml:"top"`
	start string  `xml:"start"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	end string  `xml:"end"`
	right string  `xml:"right"`
	insideH string  `xml:"insideH"`
	insideV string  `xml:"insideV"`
}
type CT_TblPrBase struct{
	tblStyle CT_String  `xml:"tblStyle"`
	tblpPr CT_TblPPr  `xml:"tblpPr"`
	tblOverlap CT_TblOverlap  `xml:"tblOverlap"`
	bidiVisual CT_OnOff  `xml:"bidiVisual"`
	tblStyleRowBandSize CT_DecimalNumber  `xml:"tblStyleRowBandSize"`
	tblStyleColBandSize CT_DecimalNumber  `xml:"tblStyleColBandSize"`
	tblW string  `xml:"tblW"`
	jc CT_JcTable  `xml:"jc"`
	tblCellSpacing string  `xml:"tblCellSpacing"`
	tblInd string  `xml:"tblInd"`
	tblBorders string  `xml:"tblBorders"`
	shd string  `xml:"shd"`
	tblLayout CT_TblLayoutType  `xml:"tblLayout"`
	tblCellMar CT_TblCellMar  `xml:"tblCellMar"`
	tblLook CT_TblLook  `xml:"tblLook"`
	tblCaption CT_String  `xml:"tblCaption"`
	tblDescription CT_String  `xml:"tblDescription"`
}
type CT_TblPr struct{
	CT_TblPrBase
	tblPrChange CT_TblPrChange  `xml:"tblPrChange"`
}
type CT_TblPrExBase struct{
	tblW string  `xml:"tblW"`
	jc CT_JcTable  `xml:"jc"`
	tblCellSpacing string  `xml:"tblCellSpacing"`
	tblInd string  `xml:"tblInd"`
	tblBorders string  `xml:"tblBorders"`
	shd string  `xml:"shd"`
	tblLayout CT_TblLayoutType  `xml:"tblLayout"`
	tblCellMar CT_TblCellMar  `xml:"tblCellMar"`
	tblLook CT_TblLook  `xml:"tblLook"`
}
type CT_TblPrEx struct{
	CT_TblPrExBase
	tblPrExChange string  `xml:"tblPrExChange"`
}
type CT_Tbl struct{
	tblPr CT_TblPr  `xml:"tblPr"`
	tblGrid string  `xml:"tblGrid"`
	 []EG_ContentRowContent  `xml:""`
}
type CT_TblLook struct{
	firstRow string  `xml:"firstRow",attr`
	lastRow string  `xml:"lastRow",attr`
	firstColumn string  `xml:"firstColumn",attr`
	lastColumn string  `xml:"lastColumn",attr`
	noHBand string  `xml:"noHBand",attr`
	noVBand string  `xml:"noVBand",attr`
	val string  `xml:"val",attr`
}
type CT_FtnPos struct{
	val string  `xml:"val",attr`
}
type CT_EdnPos struct{
	val string  `xml:"val",attr`
}
type CT_NumFmt struct{
	val ST_NumberFormat  `xml:"val",attr`
	format string  `xml:"format",attr`
}
type CT_NumRestart struct{
	val string  `xml:"val",attr`
}
type CT_FtnEdnRef struct{
	customMarkFollows string  `xml:"customMarkFollows",attr`
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_FtnEdnSepRef struct{
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_FtnEdn struct{
	 []EG_BlockLevelElts  `xml:""`
	type string  `xml:"type",attr`
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_FtnProps struct{
	pos string  `xml:"pos"`
	numFmt CT_NumFmt  `xml:"numFmt"`
	 EG_FtnEdnNumProps  `xml:""`
}
type CT_EdnProps struct{
	pos string  `xml:"pos"`
	numFmt CT_NumFmt  `xml:"numFmt"`
	 EG_FtnEdnNumProps  `xml:""`
}
type CT_FtnDocProps struct{
	CT_FtnProps
	footnote string  `xml:"footnote"`
}
type CT_EdnDocProps struct{
	CT_EdnProps
	endnote string  `xml:"endnote"`
}
type CT_RecipientData struct{
	active CT_OnOff  `xml:"active"`
	column CT_DecimalNumber  `xml:"column"`
	uniqueTag string  `xml:"uniqueTag"`
}
type CT_Base64Binary struct{
	val string  `xml:"val",attr`
}
type CT_Recipients struct{
	recipientData []CT_RecipientData  `xml:"recipientData"`
}
type CT_OdsoFieldMapData struct{
	type string  `xml:"type"`
	name CT_String  `xml:"name"`
	mappedName CT_String  `xml:"mappedName"`
	column CT_DecimalNumber  `xml:"column"`
	lid CT_Lang  `xml:"lid"`
	dynamicAddress CT_OnOff  `xml:"dynamicAddress"`
}
type CT_MailMergeSourceType struct{
	val ST_MailMergeSourceType  `xml:"val",attr`
}
type CT_Odso struct{
	udl CT_String  `xml:"udl"`
	table CT_String  `xml:"table"`
	src CT_Rel  `xml:"src"`
	colDelim CT_DecimalNumber  `xml:"colDelim"`
	type CT_MailMergeSourceType  `xml:"type"`
	fHdr CT_OnOff  `xml:"fHdr"`
	fieldMapData []string  `xml:"fieldMapData"`
	recipientData []CT_Rel  `xml:"recipientData"`
}
type CT_MailMerge struct{
	mainDocumentType CT_MailMergeDocType  `xml:"mainDocumentType"`
	linkToQuery CT_OnOff  `xml:"linkToQuery"`
	dataType CT_MailMergeDataType  `xml:"dataType"`
	connectString CT_String  `xml:"connectString"`
	query CT_String  `xml:"query"`
	dataSource CT_Rel  `xml:"dataSource"`
	headerSource CT_Rel  `xml:"headerSource"`
	doNotSuppressBlankLines CT_OnOff  `xml:"doNotSuppressBlankLines"`
	destination string  `xml:"destination"`
	addressFieldName CT_String  `xml:"addressFieldName"`
	mailSubject CT_String  `xml:"mailSubject"`
	mailAsAttachment CT_OnOff  `xml:"mailAsAttachment"`
	viewMergedData CT_OnOff  `xml:"viewMergedData"`
	activeRecord CT_DecimalNumber  `xml:"activeRecord"`
	checkErrors CT_DecimalNumber  `xml:"checkErrors"`
	odso string  `xml:"odso"`
}
type CT_TargetScreenSz struct{
	val ST_TargetScreenSz  `xml:"val",attr`
}
type CT_Compat struct{
	useSingleBorderforContiguousCells CT_OnOff  `xml:"useSingleBorderforContiguousCells"`
	wpJustification CT_OnOff  `xml:"wpJustification"`
	noTabHangInd CT_OnOff  `xml:"noTabHangInd"`
	noLeading CT_OnOff  `xml:"noLeading"`
	spaceForUL CT_OnOff  `xml:"spaceForUL"`
	noColumnBalance CT_OnOff  `xml:"noColumnBalance"`
	balanceSingleByteDoubleByteWidth CT_OnOff  `xml:"balanceSingleByteDoubleByteWidth"`
	noExtraLineSpacing CT_OnOff  `xml:"noExtraLineSpacing"`
	doNotLeaveBackslashAlone CT_OnOff  `xml:"doNotLeaveBackslashAlone"`
	ulTrailSpace CT_OnOff  `xml:"ulTrailSpace"`
	doNotExpandShiftReturn CT_OnOff  `xml:"doNotExpandShiftReturn"`
	spacingInWholePoints CT_OnOff  `xml:"spacingInWholePoints"`
	lineWrapLikeWord6 CT_OnOff  `xml:"lineWrapLikeWord6"`
	printBodyTextBeforeHeader CT_OnOff  `xml:"printBodyTextBeforeHeader"`
	printColBlack CT_OnOff  `xml:"printColBlack"`
	wpSpaceWidth CT_OnOff  `xml:"wpSpaceWidth"`
	showBreaksInFrames CT_OnOff  `xml:"showBreaksInFrames"`
	subFontBySize CT_OnOff  `xml:"subFontBySize"`
	suppressBottomSpacing CT_OnOff  `xml:"suppressBottomSpacing"`
	suppressTopSpacing CT_OnOff  `xml:"suppressTopSpacing"`
	suppressSpacingAtTopOfPage CT_OnOff  `xml:"suppressSpacingAtTopOfPage"`
	suppressTopSpacingWP CT_OnOff  `xml:"suppressTopSpacingWP"`
	suppressSpBfAfterPgBrk CT_OnOff  `xml:"suppressSpBfAfterPgBrk"`
	swapBordersFacingPages CT_OnOff  `xml:"swapBordersFacingPages"`
	convMailMergeEsc CT_OnOff  `xml:"convMailMergeEsc"`
	truncateFontHeightsLikeWP6 CT_OnOff  `xml:"truncateFontHeightsLikeWP6"`
	mwSmallCaps CT_OnOff  `xml:"mwSmallCaps"`
	usePrinterMetrics CT_OnOff  `xml:"usePrinterMetrics"`
	doNotSuppressParagraphBorders CT_OnOff  `xml:"doNotSuppressParagraphBorders"`
	wrapTrailSpaces CT_OnOff  `xml:"wrapTrailSpaces"`
	footnoteLayoutLikeWW8 CT_OnOff  `xml:"footnoteLayoutLikeWW8"`
	shapeLayoutLikeWW8 CT_OnOff  `xml:"shapeLayoutLikeWW8"`
	alignTablesRowByRow CT_OnOff  `xml:"alignTablesRowByRow"`
	forgetLastTabAlignment CT_OnOff  `xml:"forgetLastTabAlignment"`
	adjustLineHeightInTable CT_OnOff  `xml:"adjustLineHeightInTable"`
	autoSpaceLikeWord95 CT_OnOff  `xml:"autoSpaceLikeWord95"`
	noSpaceRaiseLower CT_OnOff  `xml:"noSpaceRaiseLower"`
	doNotUseHTMLParagraphAutoSpacing CT_OnOff  `xml:"doNotUseHTMLParagraphAutoSpacing"`
	layoutRawTableWidth CT_OnOff  `xml:"layoutRawTableWidth"`
	layoutTableRowsApart CT_OnOff  `xml:"layoutTableRowsApart"`
	useWord97LineBreakRules CT_OnOff  `xml:"useWord97LineBreakRules"`
	doNotBreakWrappedTables CT_OnOff  `xml:"doNotBreakWrappedTables"`
	doNotSnapToGridInCell CT_OnOff  `xml:"doNotSnapToGridInCell"`
	selectFldWithFirstOrLastChar CT_OnOff  `xml:"selectFldWithFirstOrLastChar"`
	applyBreakingRules CT_OnOff  `xml:"applyBreakingRules"`
	doNotWrapTextWithPunct CT_OnOff  `xml:"doNotWrapTextWithPunct"`
	doNotUseEastAsianBreakRules CT_OnOff  `xml:"doNotUseEastAsianBreakRules"`
	useWord2002TableStyleRules CT_OnOff  `xml:"useWord2002TableStyleRules"`
	growAutofit CT_OnOff  `xml:"growAutofit"`
	useFELayout CT_OnOff  `xml:"useFELayout"`
	useNormalStyleForList CT_OnOff  `xml:"useNormalStyleForList"`
	doNotUseIndentAsNumberingTabStop CT_OnOff  `xml:"doNotUseIndentAsNumberingTabStop"`
	useAltKinsokuLineBreakRules CT_OnOff  `xml:"useAltKinsokuLineBreakRules"`
	allowSpaceOfSameStyleInTable CT_OnOff  `xml:"allowSpaceOfSameStyleInTable"`
	doNotSuppressIndentation CT_OnOff  `xml:"doNotSuppressIndentation"`
	doNotAutofitConstrainedTables CT_OnOff  `xml:"doNotAutofitConstrainedTables"`
	autofitToFirstFixedWidthCell CT_OnOff  `xml:"autofitToFirstFixedWidthCell"`
	underlineTabInNumList CT_OnOff  `xml:"underlineTabInNumList"`
	displayHangulFixedWidth CT_OnOff  `xml:"displayHangulFixedWidth"`
	splitPgBreakAndParaMark CT_OnOff  `xml:"splitPgBreakAndParaMark"`
	doNotVertAlignCellWithSp CT_OnOff  `xml:"doNotVertAlignCellWithSp"`
	doNotBreakConstrainedForcedTable CT_OnOff  `xml:"doNotBreakConstrainedForcedTable"`
	doNotVertAlignInTxbx CT_OnOff  `xml:"doNotVertAlignInTxbx"`
	useAnsiKerningPairs CT_OnOff  `xml:"useAnsiKerningPairs"`
	cachedColBalance CT_OnOff  `xml:"cachedColBalance"`
	compatSetting []CT_CompatSetting  `xml:"compatSetting"`
}
type CT_CompatSetting struct{
	name string  `xml:"name",attr`
	uri string  `xml:"uri",attr`
	val string  `xml:"val",attr`
}
type CT_DocVar struct{
	name string  `xml:"name",attr`
	val string  `xml:"val",attr`
}
type CT_DocVars struct{
	docVar []CT_DocVar  `xml:"docVar"`
}
type CT_DocRsids struct{
	rsidRoot string  `xml:"rsidRoot"`
	rsid []string  `xml:"rsid"`
}
type CT_CharacterSpacing struct{
	val ST_CharacterSpacing  `xml:"val",attr`
}
type CT_SaveThroughXslt struct{
	solutionID string  `xml:"solutionID",attr`
}
type CT_RPrDefault struct{
	rPr CT_RPr  `xml:"rPr"`
}
type CT_PPrDefault struct{
	pPr CT_PPrGeneral  `xml:"pPr"`
}
type CT_DocDefaults struct{
	rPrDefault CT_RPrDefault  `xml:"rPrDefault"`
	pPrDefault CT_PPrDefault  `xml:"pPrDefault"`
}
type CT_ColorSchemeMapping struct{
	bg1 string  `xml:"bg1",attr`
	t1 string  `xml:"t1",attr`
	bg2 string  `xml:"bg2",attr`
	t2 string  `xml:"t2",attr`
	accent1 string  `xml:"accent1",attr`
	accent2 string  `xml:"accent2",attr`
	accent3 string  `xml:"accent3",attr`
	accent4 string  `xml:"accent4",attr`
	accent5 string  `xml:"accent5",attr`
	accent6 string  `xml:"accent6",attr`
	hyperlink string  `xml:"hyperlink",attr`
	followedHyperlink string  `xml:"followedHyperlink",attr`
}
type CT_ReadingModeInkLockDown struct{
	actualPg string  `xml:"actualPg",attr`
	w string  `xml:"w",attr`
	h string  `xml:"h",attr`
	fontSz ST_DecimalNumberOrPercent  `xml:"fontSz",attr`
}
type CT_WriteProtection struct{
	recommended string  `xml:"recommended",attr`
}
type CT_Settings struct{
	writeProtection CT_WriteProtection  `xml:"writeProtection"`
	view CT_View  `xml:"view"`
	zoom CT_Zoom  `xml:"zoom"`
	removePersonalInformation CT_OnOff  `xml:"removePersonalInformation"`
	removeDateAndTime CT_OnOff  `xml:"removeDateAndTime"`
	doNotDisplayPageBoundaries CT_OnOff  `xml:"doNotDisplayPageBoundaries"`
	displayBackgroundShape CT_OnOff  `xml:"displayBackgroundShape"`
	printPostScriptOverText CT_OnOff  `xml:"printPostScriptOverText"`
	printFractionalCharacterWidth CT_OnOff  `xml:"printFractionalCharacterWidth"`
	printFormsData CT_OnOff  `xml:"printFormsData"`
	embedTrueTypeFonts CT_OnOff  `xml:"embedTrueTypeFonts"`
	embedSystemFonts CT_OnOff  `xml:"embedSystemFonts"`
	saveSubsetFonts CT_OnOff  `xml:"saveSubsetFonts"`
	saveFormsData CT_OnOff  `xml:"saveFormsData"`
	mirrorMargins CT_OnOff  `xml:"mirrorMargins"`
	alignBordersAndEdges CT_OnOff  `xml:"alignBordersAndEdges"`
	bordersDoNotSurroundHeader CT_OnOff  `xml:"bordersDoNotSurroundHeader"`
	bordersDoNotSurroundFooter CT_OnOff  `xml:"bordersDoNotSurroundFooter"`
	gutterAtTop CT_OnOff  `xml:"gutterAtTop"`
	hideSpellingErrors CT_OnOff  `xml:"hideSpellingErrors"`
	hideGrammaticalErrors CT_OnOff  `xml:"hideGrammaticalErrors"`
	activeWritingStyle []CT_WritingStyle  `xml:"activeWritingStyle"`
	proofState CT_Proof  `xml:"proofState"`
	formsDesign CT_OnOff  `xml:"formsDesign"`
	attachedTemplate CT_Rel  `xml:"attachedTemplate"`
	linkStyles CT_OnOff  `xml:"linkStyles"`
	stylePaneFormatFilter CT_StylePaneFilter  `xml:"stylePaneFormatFilter"`
	stylePaneSortMethod CT_StyleSort  `xml:"stylePaneSortMethod"`
	documentType CT_DocType  `xml:"documentType"`
	mailMerge CT_MailMerge  `xml:"mailMerge"`
	revisionView string  `xml:"revisionView"`
	trackRevisions CT_OnOff  `xml:"trackRevisions"`
	doNotTrackMoves CT_OnOff  `xml:"doNotTrackMoves"`
	doNotTrackFormatting CT_OnOff  `xml:"doNotTrackFormatting"`
	documentProtection CT_DocProtect  `xml:"documentProtection"`
	autoFormatOverride CT_OnOff  `xml:"autoFormatOverride"`
	styleLockTheme CT_OnOff  `xml:"styleLockTheme"`
	styleLockQFSet CT_OnOff  `xml:"styleLockQFSet"`
	defaultTabStop string  `xml:"defaultTabStop"`
	autoHyphenation CT_OnOff  `xml:"autoHyphenation"`
	consecutiveHyphenLimit CT_DecimalNumber  `xml:"consecutiveHyphenLimit"`
	hyphenationZone string  `xml:"hyphenationZone"`
	doNotHyphenateCaps CT_OnOff  `xml:"doNotHyphenateCaps"`
	showEnvelope CT_OnOff  `xml:"showEnvelope"`
	summaryLength CT_DecimalNumberOrPrecent  `xml:"summaryLength"`
	clickAndTypeStyle CT_String  `xml:"clickAndTypeStyle"`
	defaultTableStyle CT_String  `xml:"defaultTableStyle"`
	evenAndOddHeaders CT_OnOff  `xml:"evenAndOddHeaders"`
	bookFoldRevPrinting CT_OnOff  `xml:"bookFoldRevPrinting"`
	bookFoldPrinting CT_OnOff  `xml:"bookFoldPrinting"`
	bookFoldPrintingSheets CT_DecimalNumber  `xml:"bookFoldPrintingSheets"`
	drawingGridHorizontalSpacing string  `xml:"drawingGridHorizontalSpacing"`
	drawingGridVerticalSpacing string  `xml:"drawingGridVerticalSpacing"`
	displayHorizontalDrawingGridEvery CT_DecimalNumber  `xml:"displayHorizontalDrawingGridEvery"`
	displayVerticalDrawingGridEvery CT_DecimalNumber  `xml:"displayVerticalDrawingGridEvery"`
	doNotUseMarginsForDrawingGridOrigin CT_OnOff  `xml:"doNotUseMarginsForDrawingGridOrigin"`
	drawingGridHorizontalOrigin string  `xml:"drawingGridHorizontalOrigin"`
	drawingGridVerticalOrigin string  `xml:"drawingGridVerticalOrigin"`
	doNotShadeFormData CT_OnOff  `xml:"doNotShadeFormData"`
	noPunctuationKerning CT_OnOff  `xml:"noPunctuationKerning"`
	characterSpacingControl CT_CharacterSpacing  `xml:"characterSpacingControl"`
	printTwoOnOne CT_OnOff  `xml:"printTwoOnOne"`
	strictFirstAndLastChars CT_OnOff  `xml:"strictFirstAndLastChars"`
	noLineBreaksAfter string  `xml:"noLineBreaksAfter"`
	noLineBreaksBefore string  `xml:"noLineBreaksBefore"`
	savePreviewPicture CT_OnOff  `xml:"savePreviewPicture"`
	doNotValidateAgainstSchema CT_OnOff  `xml:"doNotValidateAgainstSchema"`
	saveInvalidXml CT_OnOff  `xml:"saveInvalidXml"`
	ignoreMixedContent CT_OnOff  `xml:"ignoreMixedContent"`
	alwaysShowPlaceholderText CT_OnOff  `xml:"alwaysShowPlaceholderText"`
	doNotDemarcateInvalidXml CT_OnOff  `xml:"doNotDemarcateInvalidXml"`
	saveXmlDataOnly CT_OnOff  `xml:"saveXmlDataOnly"`
	useXSLTWhenSaving CT_OnOff  `xml:"useXSLTWhenSaving"`
	saveThroughXslt string  `xml:"saveThroughXslt"`
	showXMLTags CT_OnOff  `xml:"showXMLTags"`
	alwaysMergeEmptyNamespace CT_OnOff  `xml:"alwaysMergeEmptyNamespace"`
	updateFields CT_OnOff  `xml:"updateFields"`
	hdrShapeDefaults string  `xml:"hdrShapeDefaults"`
	footnotePr string  `xml:"footnotePr"`
	endnotePr string  `xml:"endnotePr"`
	compat CT_Compat  `xml:"compat"`
	docVars string  `xml:"docVars"`
	rsids string  `xml:"rsids"`
	attachedSchema []CT_String  `xml:"attachedSchema"`
	themeFontLang CT_Language  `xml:"themeFontLang"`
	clrSchemeMapping CT_ColorSchemeMapping  `xml:"clrSchemeMapping"`
	doNotIncludeSubdocsInStats CT_OnOff  `xml:"doNotIncludeSubdocsInStats"`
	doNotAutoCompressPictures CT_OnOff  `xml:"doNotAutoCompressPictures"`
	forceUpgrade CT_Empty  `xml:"forceUpgrade"`
	captions string  `xml:"captions"`
	readModeInkLockDown string  `xml:"readModeInkLockDown"`
	smartTagType []CT_SmartTagType  `xml:"smartTagType"`
	shapeDefaults string  `xml:"shapeDefaults"`
	doNotEmbedSmartTags CT_OnOff  `xml:"doNotEmbedSmartTags"`
	decimalSymbol CT_String  `xml:"decimalSymbol"`
	listSeparator CT_String  `xml:"listSeparator"`
}
type CT_StyleSort struct{
	val ST_StyleSort  `xml:"val",attr`
}
type CT_StylePaneFilter struct{
	allStyles string  `xml:"allStyles",attr`
	customStyles string  `xml:"customStyles",attr`
	latentStyles string  `xml:"latentStyles",attr`
	stylesInUse string  `xml:"stylesInUse",attr`
	headingStyles string  `xml:"headingStyles",attr`
	numberingStyles string  `xml:"numberingStyles",attr`
	tableStyles string  `xml:"tableStyles",attr`
	directFormattingOnRuns string  `xml:"directFormattingOnRuns",attr`
	directFormattingOnParagraphs string  `xml:"directFormattingOnParagraphs",attr`
	directFormattingOnNumbering string  `xml:"directFormattingOnNumbering",attr`
	directFormattingOnTables string  `xml:"directFormattingOnTables",attr`
	clearFormatting string  `xml:"clearFormatting",attr`
	top3HeadingStyles string  `xml:"top3HeadingStyles",attr`
	visibleStyles string  `xml:"visibleStyles",attr`
	alternateStyleNames string  `xml:"alternateStyleNames",attr`
	val string  `xml:"val",attr`
}
type CT_WebSettings struct{
	frameset string  `xml:"frameset"`
	divs string  `xml:"divs"`
	encoding CT_String  `xml:"encoding"`
	optimizeForBrowser string  `xml:"optimizeForBrowser"`
	relyOnVML CT_OnOff  `xml:"relyOnVML"`
	allowPNG CT_OnOff  `xml:"allowPNG"`
	doNotRelyOnCSS CT_OnOff  `xml:"doNotRelyOnCSS"`
	doNotSaveAsSingleFile CT_OnOff  `xml:"doNotSaveAsSingleFile"`
	doNotOrganizeInFolder CT_OnOff  `xml:"doNotOrganizeInFolder"`
	doNotUseLongFileNames CT_OnOff  `xml:"doNotUseLongFileNames"`
	pixelsPerInch CT_DecimalNumber  `xml:"pixelsPerInch"`
	targetScreenSz CT_TargetScreenSz  `xml:"targetScreenSz"`
	saveSmartTagsAsXml CT_OnOff  `xml:"saveSmartTagsAsXml"`
}
type CT_FrameScrollbar struct{
	val ST_FrameScrollbar  `xml:"val",attr`
}
type CT_OptimizeForBrowser struct{
	CT_OnOff
	target string  `xml:"target",attr`
}
type CT_Frame struct{
	sz CT_String  `xml:"sz"`
	name CT_String  `xml:"name"`
	title CT_String  `xml:"title"`
	longDesc CT_Rel  `xml:"longDesc"`
	sourceFileName CT_Rel  `xml:"sourceFileName"`
	marW string  `xml:"marW"`
	marH string  `xml:"marH"`
	scrollbar CT_FrameScrollbar  `xml:"scrollbar"`
	noResizeAllowed CT_OnOff  `xml:"noResizeAllowed"`
	linkedToFile CT_OnOff  `xml:"linkedToFile"`
}
type CT_FrameLayout struct{
	val ST_FrameLayout  `xml:"val",attr`
}
type CT_FramesetSplitbar struct{
	w string  `xml:"w"`
	color CT_Color  `xml:"color"`
	noBorder CT_OnOff  `xml:"noBorder"`
	flatBorders CT_OnOff  `xml:"flatBorders"`
}
type CT_Frameset struct{
	sz CT_String  `xml:"sz"`
	framesetSplitbar string  `xml:"framesetSplitbar"`
	frameLayout CT_FrameLayout  `xml:"frameLayout"`
	title CT_String  `xml:"title"`
	frameset []string  `xml:"frameset"`
	frame []CT_Frame  `xml:"frame"`
}
type CT_NumPicBullet struct{
	pict CT_Picture  `xml:"pict"`
	drawing CT_Drawing  `xml:"drawing"`
	numPicBulletId ST_DecimalNumber  `xml:"numPicBulletId",attr`
}
type CT_LevelSuffix struct{
	val string  `xml:"val",attr`
}
type CT_LevelText struct{
	val string  `xml:"val",attr`
	null string  `xml:"null",attr`
}
type CT_LvlLegacy struct{
	legacy string  `xml:"legacy",attr`
	legacySpace string  `xml:"legacySpace",attr`
	legacyIndent string  `xml:"legacyIndent",attr`
}
type CT_Lvl struct{
	start CT_DecimalNumber  `xml:"start"`
	numFmt CT_NumFmt  `xml:"numFmt"`
	lvlRestart CT_DecimalNumber  `xml:"lvlRestart"`
	pStyle CT_String  `xml:"pStyle"`
	isLgl CT_OnOff  `xml:"isLgl"`
	suff string  `xml:"suff"`
	lvlText string  `xml:"lvlText"`
	lvlPicBulletId CT_DecimalNumber  `xml:"lvlPicBulletId"`
	legacy CT_LvlLegacy  `xml:"legacy"`
	lvlJc CT_Jc  `xml:"lvlJc"`
	pPr CT_PPrGeneral  `xml:"pPr"`
	rPr CT_RPr  `xml:"rPr"`
	ilvl ST_DecimalNumber  `xml:"ilvl",attr`
	tplc string  `xml:"tplc",attr`
	tentative string  `xml:"tentative",attr`
}
type CT_MultiLevelType struct{
	val ST_MultiLevelType  `xml:"val",attr`
}
type CT_AbstractNum struct{
	nsid string  `xml:"nsid"`
	multiLevelType CT_MultiLevelType  `xml:"multiLevelType"`
	tmpl string  `xml:"tmpl"`
	name CT_String  `xml:"name"`
	styleLink CT_String  `xml:"styleLink"`
	numStyleLink CT_String  `xml:"numStyleLink"`
	lvl CT_Lvl  `xml:"lvl"`
	abstractNumId ST_DecimalNumber  `xml:"abstractNumId",attr`
}
type CT_NumLvl struct{
	startOverride CT_DecimalNumber  `xml:"startOverride"`
	lvl CT_Lvl  `xml:"lvl"`
	ilvl ST_DecimalNumber  `xml:"ilvl",attr`
}
type CT_Num struct{
	abstractNumId CT_DecimalNumber  `xml:"abstractNumId"`
	lvlOverride CT_NumLvl  `xml:"lvlOverride"`
	numId ST_DecimalNumber  `xml:"numId",attr`
}
type CT_Numbering struct{
	numPicBullet []CT_NumPicBullet  `xml:"numPicBullet"`
	abstractNum []string  `xml:"abstractNum"`
	num []CT_Num  `xml:"num"`
	numIdMacAtCleanup CT_DecimalNumber  `xml:"numIdMacAtCleanup"`
}
type CT_TblStylePr struct{
	pPr CT_PPrGeneral  `xml:"pPr"`
	rPr CT_RPr  `xml:"rPr"`
	tblPr string  `xml:"tblPr"`
	trPr CT_TrPr  `xml:"trPr"`
	tcPr CT_TcPr  `xml:"tcPr"`
	type string  `xml:"type",attr`
}
type CT_Style struct{
	name CT_String  `xml:"name"`
	aliases CT_String  `xml:"aliases"`
	basedOn CT_String  `xml:"basedOn"`
	next CT_String  `xml:"next"`
	link CT_String  `xml:"link"`
	autoRedefine CT_OnOff  `xml:"autoRedefine"`
	hidden CT_OnOff  `xml:"hidden"`
	uiPriority CT_DecimalNumber  `xml:"uiPriority"`
	semiHidden CT_OnOff  `xml:"semiHidden"`
	unhideWhenUsed CT_OnOff  `xml:"unhideWhenUsed"`
	qFormat CT_OnOff  `xml:"qFormat"`
	locked CT_OnOff  `xml:"locked"`
	personal CT_OnOff  `xml:"personal"`
	personalCompose CT_OnOff  `xml:"personalCompose"`
	personalReply CT_OnOff  `xml:"personalReply"`
	rsid string  `xml:"rsid"`
	pPr CT_PPrGeneral  `xml:"pPr"`
	rPr CT_RPr  `xml:"rPr"`
	tblPr string  `xml:"tblPr"`
	trPr CT_TrPr  `xml:"trPr"`
	tcPr CT_TcPr  `xml:"tcPr"`
	tblStylePr []CT_TblStylePr  `xml:"tblStylePr"`
	type ST_StyleType  `xml:"type",attr`
	styleId string  `xml:"styleId",attr`
	default string  `xml:"default",attr`
	customStyle string  `xml:"customStyle",attr`
}
type CT_LsdException struct{
	name string  `xml:"name",attr`
	locked string  `xml:"locked",attr`
	uiPriority ST_DecimalNumber  `xml:"uiPriority",attr`
	semiHidden string  `xml:"semiHidden",attr`
	unhideWhenUsed string  `xml:"unhideWhenUsed",attr`
	qFormat string  `xml:"qFormat",attr`
}
type CT_LatentStyles struct{
	lsdException []string  `xml:"lsdException"`
	defLockedState string  `xml:"defLockedState",attr`
	defUIPriority ST_DecimalNumber  `xml:"defUIPriority",attr`
	defSemiHidden string  `xml:"defSemiHidden",attr`
	defUnhideWhenUsed string  `xml:"defUnhideWhenUsed",attr`
	defQFormat string  `xml:"defQFormat",attr`
	count ST_DecimalNumber  `xml:"count",attr`
}
type CT_Styles struct{
	docDefaults string  `xml:"docDefaults"`
	latentStyles string  `xml:"latentStyles"`
	style []CT_Style  `xml:"style"`
}
type CT_Panose struct{
	val string  `xml:"val",attr`
}
type CT_FontFamily struct{
	val ST_FontFamily  `xml:"val",attr`
}
type CT_Pitch struct{
	val ST_Pitch  `xml:"val",attr`
}
type CT_FontSig struct{
	usb0 string  `xml:"usb0",attr`
	usb1 string  `xml:"usb1",attr`
	usb2 string  `xml:"usb2",attr`
	usb3 string  `xml:"usb3",attr`
	csb0 string  `xml:"csb0",attr`
	csb1 string  `xml:"csb1",attr`
}
type CT_FontRel struct{
	CT_Rel
	fontKey string  `xml:"fontKey",attr`
	subsetted string  `xml:"subsetted",attr`
}
type CT_Font struct{
	altName CT_String  `xml:"altName"`
	panose1 string  `xml:"panose1"`
	charset string  `xml:"charset"`
	family CT_FontFamily  `xml:"family"`
	notTrueType CT_OnOff  `xml:"notTrueType"`
	pitch CT_Pitch  `xml:"pitch"`
	sig CT_FontSig  `xml:"sig"`
	embedRegular CT_FontRel  `xml:"embedRegular"`
	embedBold CT_FontRel  `xml:"embedBold"`
	embedItalic CT_FontRel  `xml:"embedItalic"`
	embedBoldItalic CT_FontRel  `xml:"embedBoldItalic"`
	name string  `xml:"name",attr`
}
type CT_FontsList struct{
	font []CT_Font  `xml:"font"`
}
type CT_DivBdr struct{
	top string  `xml:"top"`
	left string  `xml:"left"`
	bottom string  `xml:"bottom"`
	right string  `xml:"right"`
}
type CT_Div struct{
	blockQuote CT_OnOff  `xml:"blockQuote"`
	bodyDiv CT_OnOff  `xml:"bodyDiv"`
	marLeft string  `xml:"marLeft"`
	marRight string  `xml:"marRight"`
	marTop string  `xml:"marTop"`
	marBottom string  `xml:"marBottom"`
	divBdr string  `xml:"divBdr"`
	divsChild []string  `xml:"divsChild"`
	id ST_DecimalNumber  `xml:"id",attr`
}
type CT_Divs struct{
	div CT_Div  `xml:"div"`
}
type CT_TxbxContent struct{
}
type CT_Body struct{
	sectPr CT_SectPr  `xml:"sectPr"`
	 []EG_BlockLevelElts  `xml:""`
}
type CT_ShapeDefaults struct{
}
type CT_Comments struct{
	comment []CT_Comment  `xml:"comment"`
}
type CT_Footnotes struct{
	footnote string  `xml:"footnote"`
}
type CT_Endnotes struct{
	endnote string  `xml:"endnote"`
}
type CT_SmartTagType struct{
	namespaceuri string  `xml:"namespaceuri",attr`
	name string  `xml:"name",attr`
	url string  `xml:"url",attr`
}
type CT_DocPartBehavior struct{
	val ST_DocPartBehavior  `xml:"val",attr`
}
type CT_DocPartBehaviors struct{
	behavior []CT_DocPartBehavior  `xml:"behavior"`
}
type CT_DocPartType struct{
	val ST_DocPartType  `xml:"val",attr`
}
type CT_DocPartTypes struct{
	type []CT_DocPartType  `xml:"type"`
	all string  `xml:"all",attr`
}
type CT_DocPartGallery struct{
	val ST_DocPartGallery  `xml:"val",attr`
}
type CT_DocPartCategory struct{
	name CT_String  `xml:"name"`
	gallery CT_DocPartGallery  `xml:"gallery"`
}
type CT_DocPartName struct{
	val string  `xml:"val",attr`
	decorated string  `xml:"decorated",attr`
}
type CT_DocPartPr struct{
	name CT_DocPartName  `xml:"name"`
	style CT_String  `xml:"style"`
	category CT_DocPartCategory  `xml:"category"`
	types string  `xml:"types"`
	behaviors string  `xml:"behaviors"`
	description CT_String  `xml:"description"`
	guid string  `xml:"guid"`
}
type CT_DocPart struct{
	docPartPr CT_DocPartPr  `xml:"docPartPr"`
	docPartBody string  `xml:"docPartBody"`
}
type CT_DocParts struct{
	docPart []CT_DocPart  `xml:"docPart"`
}
type CT_Caption struct{
	name string  `xml:"name",attr`
	pos string  `xml:"pos",attr`
	chapNum string  `xml:"chapNum",attr`
	heading ST_DecimalNumber  `xml:"heading",attr`
	noLabel string  `xml:"noLabel",attr`
	numFmt ST_NumberFormat  `xml:"numFmt",attr`
	sep ST_ChapterSep  `xml:"sep",attr`
}
type CT_AutoCaption struct{
	name string  `xml:"name",attr`
	caption string  `xml:"caption",attr`
}
type CT_AutoCaptions struct{
	autoCaption []CT_AutoCaption  `xml:"autoCaption"`
}
type CT_Captions struct{
	caption []CT_Caption  `xml:"caption"`
	autoCaptions string  `xml:"autoCaptions"`
}
type CT_DocumentBase struct{
	background string  `xml:"background"`
}
type CT_Document struct{
	CT_DocumentBase
	body string  `xml:"body"`
	conformance string  `xml:"conformance",attr`
}
type CT_GlossaryDocument struct{
	CT_DocumentBase
	docParts string  `xml:"docParts"`
}
