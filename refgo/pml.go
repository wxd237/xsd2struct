package ooxml
type WordSchema  struct {
	cmAuthorLst string  `xml:"cmAuthorLst"`
	cmLst string  `xml:"cmLst"`
	oleObj CT_OleObject  `xml:"oleObj"`
	presentation string  `xml:"presentation"`
	presentationPr string  `xml:"presentationPr"`
	sld string  `xml:"sld"`
	sldLayout string  `xml:"sldLayout"`
	sldMaster string  `xml:"sldMaster"`
	handoutMaster string  `xml:"handoutMaster"`
	notesMaster string  `xml:"notesMaster"`
	notes string  `xml:"notes"`
	sldSyncPr string  `xml:"sldSyncPr"`
	tagLst string  `xml:"tagLst"`
	viewPr string  `xml:"viewPr"`
}
type CT_SideDirectionTransition struct{
	dir string  `xml:"dir",attr`
}
type CT_CornerDirectionTransition struct{
	dir string  `xml:"dir",attr`
}
type CT_EightDirectionTransition struct{
	dir string  `xml:"dir",attr`
}
type CT_OrientationTransition struct{
	dir ST_Direction  `xml:"dir",attr`
}
type CT_InOutTransition struct{
	dir string  `xml:"dir",attr`
}
type CT_OptionalBlackTransition struct{
	thruBlk string  `xml:"thruBlk",attr`
}
type CT_SplitTransition struct{
	orient ST_Direction  `xml:"orient",attr`
	dir string  `xml:"dir",attr`
}
type CT_WheelTransition struct{
	spokes string  `xml:"spokes",attr`
}
type CT_TransitionStartSoundAction struct{
	snd string  `xml:"snd"`
	loop string  `xml:"loop",attr`
}
type CT_TransitionSoundAction struct{
	stSnd string  `xml:"stSnd"`
	endSnd CT_Empty  `xml:"endSnd"`
}
type CT_SlideTransition struct{
	sndAc string  `xml:"sndAc"`
	extLst string  `xml:"extLst"`
	blinds string  `xml:"blinds"`
	checker string  `xml:"checker"`
	circle CT_Empty  `xml:"circle"`
	dissolve CT_Empty  `xml:"dissolve"`
	comb string  `xml:"comb"`
	cover string  `xml:"cover"`
	cut string  `xml:"cut"`
	diamond CT_Empty  `xml:"diamond"`
	fade string  `xml:"fade"`
	newsflash CT_Empty  `xml:"newsflash"`
	plus CT_Empty  `xml:"plus"`
	pull string  `xml:"pull"`
	push string  `xml:"push"`
	random CT_Empty  `xml:"random"`
	randomBar string  `xml:"randomBar"`
	split string  `xml:"split"`
	strips string  `xml:"strips"`
	wedge CT_Empty  `xml:"wedge"`
	wheel string  `xml:"wheel"`
	wipe string  `xml:"wipe"`
	zoom string  `xml:"zoom"`
	spd string  `xml:"spd",attr`
	advClick string  `xml:"advClick",attr`
	advTm string  `xml:"advTm",attr`
}
type CT_TLIterateIntervalTime struct{
	val ST_TLTime  `xml:"val",attr`
}
type CT_TLIterateIntervalPercentage struct{
	val string  `xml:"val",attr`
}
type CT_TLIterateData struct{
	tmAbs CT_TLIterateIntervalTime  `xml:"tmAbs"`
	tmPct CT_TLIterateIntervalPercentage  `xml:"tmPct"`
	type ST_IterateType  `xml:"type",attr`
	backwards string  `xml:"backwards",attr`
}
type CT_TLSubShapeId struct{
	spid string  `xml:"spid",attr`
}
type CT_TLTextTargetElement struct{
	charRg string  `xml:"charRg"`
	pRg string  `xml:"pRg"`
}
type CT_TLOleChartTargetElement struct{
	type ST_TLChartSubelementType  `xml:"type",attr`
	lvl string  `xml:"lvl",attr`
}
type CT_TLShapeTargetElement struct{
	bg CT_Empty  `xml:"bg"`
	subSp string  `xml:"subSp"`
	oleChartEl CT_TLOleChartTargetElement  `xml:"oleChartEl"`
	txEl string  `xml:"txEl"`
	graphicEl string  `xml:"graphicEl"`
	spid string  `xml:"spid",attr`
}
type CT_TLTimeTargetElement struct{
	sldTgt CT_Empty  `xml:"sldTgt"`
	sndTgt string  `xml:"sndTgt"`
	spTgt CT_TLShapeTargetElement  `xml:"spTgt"`
	inkTgt string  `xml:"inkTgt"`
}
type CT_TLTriggerTimeNodeID struct{
	val string  `xml:"val",attr`
}
type CT_TLTriggerRuntimeNode struct{
	val string  `xml:"val",attr`
}
type CT_TLTimeCondition struct{
	tgtEl CT_TLTimeTargetElement  `xml:"tgtEl"`
	tn string  `xml:"tn"`
	rtn string  `xml:"rtn"`
	evt ST_TLTriggerEvent  `xml:"evt",attr`
	delay ST_TLTime  `xml:"delay",attr`
}
type CT_TLTimeConditionList struct{
	cond []string  `xml:"cond"`
}
type CT_TimeNodeList struct{
	par string  `xml:"par"`
	seq string  `xml:"seq"`
	excl string  `xml:"excl"`
	anim CT_TLAnimateBehavior  `xml:"anim"`
	animClr CT_TLAnimateColorBehavior  `xml:"animClr"`
	animEffect CT_TLAnimateEffectBehavior  `xml:"animEffect"`
	animMotion CT_TLAnimateMotionBehavior  `xml:"animMotion"`
	animRot CT_TLAnimateRotationBehavior  `xml:"animRot"`
	animScale CT_TLAnimateScaleBehavior  `xml:"animScale"`
	cmd string  `xml:"cmd"`
	set CT_TLSetBehavior  `xml:"set"`
	audio string  `xml:"audio"`
	video string  `xml:"video"`
}
type CT_TLCommonTimeNodeData struct{
	stCondLst string  `xml:"stCondLst"`
	endCondLst string  `xml:"endCondLst"`
	endSync string  `xml:"endSync"`
	iterate CT_TLIterateData  `xml:"iterate"`
	childTnLst string  `xml:"childTnLst"`
	subTnLst string  `xml:"subTnLst"`
	id string  `xml:"id",attr`
	presetID string  `xml:"presetID",attr`
	presetClass string  `xml:"presetClass",attr`
	presetSubtype string  `xml:"presetSubtype",attr`
	dur ST_TLTime  `xml:"dur",attr`
	repeatCount ST_TLTime  `xml:"repeatCount",attr`
	repeatDur ST_TLTime  `xml:"repeatDur",attr`
	spd string  `xml:"spd",attr`
	accel string  `xml:"accel",attr`
	decel string  `xml:"decel",attr`
	autoRev string  `xml:"autoRev",attr`
	restart string  `xml:"restart",attr`
	fill string  `xml:"fill",attr`
	syncBehavior string  `xml:"syncBehavior",attr`
	tmFilter string  `xml:"tmFilter",attr`
	evtFilter string  `xml:"evtFilter",attr`
	display string  `xml:"display",attr`
	masterRel string  `xml:"masterRel",attr`
	bldLvl string  `xml:"bldLvl",attr`
	grpId string  `xml:"grpId",attr`
	afterEffect string  `xml:"afterEffect",attr`
	nodeType string  `xml:"nodeType",attr`
	nodePh string  `xml:"nodePh",attr`
}
type CT_TLTimeNodeParallel struct{
	cTn string  `xml:"cTn"`
}
type CT_TLTimeNodeSequence struct{
	cTn string  `xml:"cTn"`
	prevCondLst string  `xml:"prevCondLst"`
	nextCondLst string  `xml:"nextCondLst"`
	concurrent string  `xml:"concurrent",attr`
	prevAc string  `xml:"prevAc",attr`
	nextAc string  `xml:"nextAc",attr`
}
type CT_TLTimeNodeExclusive struct{
	cTn string  `xml:"cTn"`
}
type CT_TLBehaviorAttributeNameList struct{
	attrName []string  `xml:"attrName"`
}
type CT_TLCommonBehaviorData struct{
	cTn string  `xml:"cTn"`
	tgtEl CT_TLTimeTargetElement  `xml:"tgtEl"`
	attrNameLst string  `xml:"attrNameLst"`
	additive string  `xml:"additive",attr`
	accumulate ST_TLBehaviorAccumulateType  `xml:"accumulate",attr`
	xfrmType string  `xml:"xfrmType",attr`
	from string  `xml:"from",attr`
	to string  `xml:"to",attr`
	by string  `xml:"by",attr`
	rctx string  `xml:"rctx",attr`
	override string  `xml:"override",attr`
}
type CT_TLAnimVariantBooleanVal struct{
	val string  `xml:"val",attr`
}
type CT_TLAnimVariantIntegerVal struct{
	val string  `xml:"val",attr`
}
type CT_TLAnimVariantFloatVal struct{
	val string  `xml:"val",attr`
}
type CT_TLAnimVariantStringVal struct{
	val string  `xml:"val",attr`
}
type CT_TLAnimVariant struct{
	boolVal CT_TLAnimVariantBooleanVal  `xml:"boolVal"`
	intVal CT_TLAnimVariantIntegerVal  `xml:"intVal"`
	fltVal CT_TLAnimVariantFloatVal  `xml:"fltVal"`
	strVal CT_TLAnimVariantStringVal  `xml:"strVal"`
	clrVal string  `xml:"clrVal"`
}
type CT_TLTimeAnimateValue struct{
	val CT_TLAnimVariant  `xml:"val"`
	tm ST_TLTimeAnimateValueTime  `xml:"tm",attr`
	fmla string  `xml:"fmla",attr`
}
type CT_TLTimeAnimateValueList struct{
	tav []CT_TLTimeAnimateValue  `xml:"tav"`
}
type CT_TLAnimateBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	tavLst string  `xml:"tavLst"`
	by string  `xml:"by",attr`
	from string  `xml:"from",attr`
	to string  `xml:"to",attr`
	calcmode string  `xml:"calcmode",attr`
	valueType ST_TLAnimateBehaviorValueType  `xml:"valueType",attr`
}
type CT_TLByRgbColorTransform struct{
	r string  `xml:"r",attr`
	g string  `xml:"g",attr`
	b string  `xml:"b",attr`
}
type CT_TLByHslColorTransform struct{
	h string  `xml:"h",attr`
	s string  `xml:"s",attr`
	l string  `xml:"l",attr`
}
type CT_TLByAnimateColorTransform struct{
	rgb string  `xml:"rgb"`
	hsl string  `xml:"hsl"`
}
type CT_TLAnimateColorBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	by string  `xml:"by"`
	from string  `xml:"from"`
	to string  `xml:"to"`
	clrSpc ST_TLAnimateColorSpace  `xml:"clrSpc",attr`
	dir ST_TLAnimateColorDirection  `xml:"dir",attr`
}
type CT_TLAnimateEffectBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	progress CT_TLAnimVariant  `xml:"progress"`
	transition string  `xml:"transition",attr`
	filter string  `xml:"filter",attr`
	prLst string  `xml:"prLst",attr`
}
type CT_TLPoint struct{
	x string  `xml:"x",attr`
	y string  `xml:"y",attr`
}
type CT_TLAnimateMotionBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	by CT_TLPoint  `xml:"by"`
	from CT_TLPoint  `xml:"from"`
	to CT_TLPoint  `xml:"to"`
	rCtr CT_TLPoint  `xml:"rCtr"`
	origin ST_TLAnimateMotionBehaviorOrigin  `xml:"origin",attr`
	path string  `xml:"path",attr`
	pathEditMode string  `xml:"pathEditMode",attr`
	rAng string  `xml:"rAng",attr`
	ptsTypes string  `xml:"ptsTypes",attr`
}
type CT_TLAnimateRotationBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	by string  `xml:"by",attr`
	from string  `xml:"from",attr`
	to string  `xml:"to",attr`
}
type CT_TLAnimateScaleBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	by CT_TLPoint  `xml:"by"`
	from CT_TLPoint  `xml:"from"`
	to CT_TLPoint  `xml:"to"`
	zoomContents string  `xml:"zoomContents",attr`
}
type CT_TLCommandBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	type string  `xml:"type",attr`
	cmd string  `xml:"cmd",attr`
}
type CT_TLSetBehavior struct{
	cBhvr CT_TLCommonBehaviorData  `xml:"cBhvr"`
	to CT_TLAnimVariant  `xml:"to"`
}
type CT_TLCommonMediaNodeData struct{
	cTn string  `xml:"cTn"`
	tgtEl CT_TLTimeTargetElement  `xml:"tgtEl"`
	vol string  `xml:"vol",attr`
	mute string  `xml:"mute",attr`
	numSld string  `xml:"numSld",attr`
	showWhenStopped string  `xml:"showWhenStopped",attr`
}
type CT_TLMediaNodeAudio struct{
	cMediaNode string  `xml:"cMediaNode"`
	isNarration string  `xml:"isNarration",attr`
}
type CT_TLMediaNodeVideo struct{
	cMediaNode string  `xml:"cMediaNode"`
	fullScrn string  `xml:"fullScrn",attr`
}
type CT_TLTemplate struct{
	tnLst string  `xml:"tnLst"`
	lvl string  `xml:"lvl",attr`
}
type CT_TLTemplateList struct{
	tmpl CT_TLTemplate  `xml:"tmpl"`
}
type CT_TLBuildParagraph struct{
	tmplLst string  `xml:"tmplLst"`
	build string  `xml:"build",attr`
	bldLvl string  `xml:"bldLvl",attr`
	animBg string  `xml:"animBg",attr`
	autoUpdateAnimBg string  `xml:"autoUpdateAnimBg",attr`
	rev string  `xml:"rev",attr`
	advAuto ST_TLTime  `xml:"advAuto",attr`
}
type CT_TLBuildDiagram struct{
	bld string  `xml:"bld",attr`
}
type CT_TLOleBuildChart struct{
	bld string  `xml:"bld",attr`
	animBg string  `xml:"animBg",attr`
}
type CT_TLGraphicalObjectBuild struct{
	bldAsOne CT_Empty  `xml:"bldAsOne"`
	bldSub string  `xml:"bldSub"`
}
type CT_BuildList struct{
	bldP string  `xml:"bldP"`
	bldDgm string  `xml:"bldDgm"`
	bldOleChart string  `xml:"bldOleChart"`
	bldGraphic string  `xml:"bldGraphic"`
}
type CT_SlideTiming struct{
	tnLst string  `xml:"tnLst"`
	bldLst string  `xml:"bldLst"`
	extLst string  `xml:"extLst"`
}
type CT_Empty struct{
}
type CT_IndexRange struct{
	st string  `xml:"st",attr`
	end string  `xml:"end",attr`
}
type CT_SlideRelationshipListEntry struct{
}
type CT_SlideRelationshipList struct{
	sld []string  `xml:"sld"`
}
type CT_CustomShowId struct{
	id string  `xml:"id",attr`
}
type CT_CustomerData struct{
}
type CT_TagsData struct{
}
type CT_CustomerDataList struct{
	custData []string  `xml:"custData"`
	tags string  `xml:"tags"`
}
type CT_Extension struct{
	uri string  `xml:"uri",attr`
}
type CT_ExtensionList struct{
	 EG_ExtensionList  `xml:""`
}
type CT_ExtensionListModify struct{
	 EG_ExtensionList  `xml:""`
	mod string  `xml:"mod",attr`
}
type CT_CommentAuthor struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	name ST_Name  `xml:"name",attr`
	initials ST_Name  `xml:"initials",attr`
	lastIdx string  `xml:"lastIdx",attr`
	clrIdx string  `xml:"clrIdx",attr`
}
type CT_CommentAuthorList struct{
	cmAuthor []CT_CommentAuthor  `xml:"cmAuthor"`
}
type CT_Comment struct{
	pos string  `xml:"pos"`
	text string  `xml:"text"`
	extLst string  `xml:"extLst"`
	authorId string  `xml:"authorId",attr`
	dt string  `xml:"dt",attr`
	idx string  `xml:"idx",attr`
}
type CT_CommentList struct{
	cm []CT_Comment  `xml:"cm"`
}
type CT_OleObjectEmbed struct{
	extLst string  `xml:"extLst"`
	followColorScheme ST_OleObjectFollowColorScheme  `xml:"followColorScheme",attr`
}
type CT_OleObjectLink struct{
	extLst string  `xml:"extLst"`
	updateAutomatic string  `xml:"updateAutomatic",attr`
}
type CT_OleObject struct{
	pic CT_Picture  `xml:"pic"`
	embed string  `xml:"embed"`
	link CT_OleObjectLink  `xml:"link"`
	progId string  `xml:"progId",attr`
}
type CT_Control struct{
	extLst string  `xml:"extLst"`
	pic CT_Picture  `xml:"pic"`
}
type CT_ControlList struct{
	control []CT_Control  `xml:"control"`
}
type CT_SlideIdListEntry struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
}
type CT_SlideIdList struct{
	sldId []string  `xml:"sldId"`
}
type CT_SlideMasterIdListEntry struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
}
type CT_SlideMasterIdList struct{
	sldMasterId []string  `xml:"sldMasterId"`
}
type CT_NotesMasterIdListEntry struct{
	extLst string  `xml:"extLst"`
}
type CT_NotesMasterIdList struct{
	notesMasterId string  `xml:"notesMasterId"`
}
type CT_HandoutMasterIdListEntry struct{
	extLst string  `xml:"extLst"`
}
type CT_HandoutMasterIdList struct{
	handoutMasterId string  `xml:"handoutMasterId"`
}
type CT_EmbeddedFontDataId struct{
}
type CT_EmbeddedFontListEntry struct{
	font string  `xml:"font"`
	regular string  `xml:"regular"`
	bold string  `xml:"bold"`
	italic string  `xml:"italic"`
	boldItalic string  `xml:"boldItalic"`
}
type CT_EmbeddedFontList struct{
	embeddedFont []string  `xml:"embeddedFont"`
}
type CT_SmartTags struct{
}
type CT_CustomShow struct{
	sldLst string  `xml:"sldLst"`
	extLst string  `xml:"extLst"`
	name ST_Name  `xml:"name",attr`
	id string  `xml:"id",attr`
}
type CT_CustomShowList struct{
	custShow []string  `xml:"custShow"`
}
type CT_PhotoAlbum struct{
	extLst string  `xml:"extLst"`
	bw string  `xml:"bw",attr`
	showCaptions string  `xml:"showCaptions",attr`
	layout ST_PhotoAlbumLayout  `xml:"layout",attr`
	frame ST_PhotoAlbumFrameShape  `xml:"frame",attr`
}
type CT_SlideSize struct{
	cx string  `xml:"cx",attr`
	cy string  `xml:"cy",attr`
	type string  `xml:"type",attr`
}
type CT_Kinsoku struct{
	lang string  `xml:"lang",attr`
	invalStChars string  `xml:"invalStChars",attr`
	invalEndChars string  `xml:"invalEndChars",attr`
}
type CT_ModifyVerifier struct{
	algorithmName string  `xml:"algorithmName",attr`
	hashValue string  `xml:"hashValue",attr`
	saltValue string  `xml:"saltValue",attr`
	spinValue string  `xml:"spinValue",attr`
	cryptProviderType string  `xml:"cryptProviderType",attr`
	cryptAlgorithmClass string  `xml:"cryptAlgorithmClass",attr`
	cryptAlgorithmType string  `xml:"cryptAlgorithmType",attr`
	cryptAlgorithmSid string  `xml:"cryptAlgorithmSid",attr`
	spinCount string  `xml:"spinCount",attr`
	saltData string  `xml:"saltData",attr`
	hashData string  `xml:"hashData",attr`
	cryptProvider string  `xml:"cryptProvider",attr`
	algIdExt string  `xml:"algIdExt",attr`
	algIdExtSource string  `xml:"algIdExtSource",attr`
	cryptProviderTypeExt string  `xml:"cryptProviderTypeExt",attr`
	cryptProviderTypeExtSource string  `xml:"cryptProviderTypeExtSource",attr`
}
type CT_Presentation struct{
	sldMasterIdLst string  `xml:"sldMasterIdLst"`
	notesMasterIdLst string  `xml:"notesMasterIdLst"`
	handoutMasterIdLst string  `xml:"handoutMasterIdLst"`
	sldIdLst string  `xml:"sldIdLst"`
	sldSz string  `xml:"sldSz"`
	notesSz string  `xml:"notesSz"`
	smartTags string  `xml:"smartTags"`
	embeddedFontLst string  `xml:"embeddedFontLst"`
	custShowLst string  `xml:"custShowLst"`
	photoAlbum CT_PhotoAlbum  `xml:"photoAlbum"`
	custDataLst string  `xml:"custDataLst"`
	kinsoku string  `xml:"kinsoku"`
	defaultTextStyle string  `xml:"defaultTextStyle"`
	modifyVerifier string  `xml:"modifyVerifier"`
	extLst string  `xml:"extLst"`
	serverZoom string  `xml:"serverZoom",attr`
	firstSlideNum string  `xml:"firstSlideNum",attr`
	showSpecialPlsOnTitleSld string  `xml:"showSpecialPlsOnTitleSld",attr`
	rtl string  `xml:"rtl",attr`
	removePersonalInfoOnSave string  `xml:"removePersonalInfoOnSave",attr`
	compatMode string  `xml:"compatMode",attr`
	strictFirstAndLastChars string  `xml:"strictFirstAndLastChars",attr`
	embedTrueTypeFonts string  `xml:"embedTrueTypeFonts",attr`
	saveSubsetFonts string  `xml:"saveSubsetFonts",attr`
	autoCompressPictures string  `xml:"autoCompressPictures",attr`
	bookmarkIdSeed string  `xml:"bookmarkIdSeed",attr`
	conformance string  `xml:"conformance",attr`
}
type CT_HtmlPublishProperties struct{
	extLst string  `xml:"extLst"`
	 EG_SlideListChoice  `xml:""`
	showSpeakerNotes string  `xml:"showSpeakerNotes",attr`
	target string  `xml:"target",attr`
	title string  `xml:"title",attr`
}
type CT_WebProperties struct{
	extLst string  `xml:"extLst"`
	showAnimation string  `xml:"showAnimation",attr`
	resizeGraphics string  `xml:"resizeGraphics",attr`
	allowPng string  `xml:"allowPng",attr`
	relyOnVml string  `xml:"relyOnVml",attr`
	organizeInFolders string  `xml:"organizeInFolders",attr`
	useLongFilenames string  `xml:"useLongFilenames",attr`
	imgSz ST_WebScreenSize  `xml:"imgSz",attr`
	encoding string  `xml:"encoding",attr`
	clr ST_WebColorType  `xml:"clr",attr`
}
type CT_PrintProperties struct{
	extLst string  `xml:"extLst"`
	prnWhat ST_PrintWhat  `xml:"prnWhat",attr`
	clrMode string  `xml:"clrMode",attr`
	hiddenSlides string  `xml:"hiddenSlides",attr`
	scaleToFitPaper string  `xml:"scaleToFitPaper",attr`
	frameSlides string  `xml:"frameSlides",attr`
}
type CT_ShowInfoBrowse struct{
	showScrollbar string  `xml:"showScrollbar",attr`
}
type CT_ShowInfoKiosk struct{
	restart string  `xml:"restart",attr`
}
type CT_ShowProperties struct{
	penClr string  `xml:"penClr"`
	extLst string  `xml:"extLst"`
	 EG_SlideListChoice  `xml:""`
	loop string  `xml:"loop",attr`
	showNarration string  `xml:"showNarration",attr`
	showAnimation string  `xml:"showAnimation",attr`
	useTimings string  `xml:"useTimings",attr`
}
type CT_PresentationProperties struct{
	htmlPubPr string  `xml:"htmlPubPr"`
	webPr string  `xml:"webPr"`
	prnPr string  `xml:"prnPr"`
	showPr string  `xml:"showPr"`
	clrMru string  `xml:"clrMru"`
	extLst string  `xml:"extLst"`
}
type CT_HeaderFooter struct{
	extLst string  `xml:"extLst"`
	sldNum string  `xml:"sldNum",attr`
	hdr string  `xml:"hdr",attr`
	ftr string  `xml:"ftr",attr`
	dt string  `xml:"dt",attr`
}
type CT_Placeholder struct{
	extLst string  `xml:"extLst"`
	type string  `xml:"type",attr`
	orient ST_Direction  `xml:"orient",attr`
	sz string  `xml:"sz",attr`
	idx string  `xml:"idx",attr`
	hasCustomPrompt string  `xml:"hasCustomPrompt",attr`
}
type CT_ApplicationNonVisualDrawingProps struct{
	ph string  `xml:"ph"`
	custDataLst string  `xml:"custDataLst"`
	extLst string  `xml:"extLst"`
	 a:EG_Media  `xml:""`
	isPhoto string  `xml:"isPhoto",attr`
	userDrawn string  `xml:"userDrawn",attr`
}
type CT_ShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvSpPr string  `xml:"cNvSpPr"`
	nvPr string  `xml:"nvPr"`
}
type CT_Shape struct{
	nvSpPr string  `xml:"nvSpPr"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	txBody string  `xml:"txBody"`
	extLst string  `xml:"extLst"`
	useBgFill string  `xml:"useBgFill",attr`
}
type CT_ConnectorNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvCxnSpPr string  `xml:"cNvCxnSpPr"`
	nvPr string  `xml:"nvPr"`
}
type CT_Connector struct{
	nvCxnSpPr string  `xml:"nvCxnSpPr"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_PictureNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvPicPr string  `xml:"cNvPicPr"`
	nvPr string  `xml:"nvPr"`
}
type CT_Picture struct{
	nvPicPr string  `xml:"nvPicPr"`
	blipFill string  `xml:"blipFill"`
	spPr string  `xml:"spPr"`
	style string  `xml:"style"`
	extLst string  `xml:"extLst"`
}
type CT_GraphicalObjectFrameNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGraphicFramePr string  `xml:"cNvGraphicFramePr"`
	nvPr string  `xml:"nvPr"`
}
type CT_GraphicalObjectFrame struct{
	nvGraphicFramePr string  `xml:"nvGraphicFramePr"`
	xfrm string  `xml:"xfrm"`
	extLst string  `xml:"extLst"`
	bwMode string  `xml:"bwMode",attr`
}
type CT_GroupShapeNonVisual struct{
	cNvPr string  `xml:"cNvPr"`
	cNvGrpSpPr string  `xml:"cNvGrpSpPr"`
	nvPr string  `xml:"nvPr"`
}
type CT_GroupShape struct{
	nvGrpSpPr string  `xml:"nvGrpSpPr"`
	grpSpPr string  `xml:"grpSpPr"`
	extLst string  `xml:"extLst"`
	sp CT_Shape  `xml:"sp"`
	grpSp CT_GroupShape  `xml:"grpSp"`
	graphicFrame CT_GraphicalObjectFrame  `xml:"graphicFrame"`
	cxnSp CT_Connector  `xml:"cxnSp"`
	pic CT_Picture  `xml:"pic"`
	contentPart CT_Rel  `xml:"contentPart"`
}
type CT_Rel struct{
}
type CT_BackgroundProperties struct{
	extLst string  `xml:"extLst"`
	 a:EG_EffectProperties  `xml:""`
	shadeToTitle string  `xml:"shadeToTitle",attr`
}
type CT_Background struct{
	 EG_Background  `xml:""`
	bwMode string  `xml:"bwMode",attr`
}
type CT_CommonSlideData struct{
	bg string  `xml:"bg"`
	spTree CT_GroupShape  `xml:"spTree"`
	custDataLst string  `xml:"custDataLst"`
	controls string  `xml:"controls"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_Slide struct{
	cSld string  `xml:"cSld"`
	transition string  `xml:"transition"`
	timing string  `xml:"timing"`
	extLst string  `xml:"extLst"`
	 EG_ChildSlide  `xml:""`
	show string  `xml:"show",attr`
}
type CT_SlideLayout struct{
	cSld string  `xml:"cSld"`
	transition string  `xml:"transition"`
	timing string  `xml:"timing"`
	hf string  `xml:"hf"`
	extLst string  `xml:"extLst"`
	 EG_ChildSlide  `xml:""`
	matchingName string  `xml:"matchingName",attr`
	type string  `xml:"type",attr`
	preserve string  `xml:"preserve",attr`
	userDrawn string  `xml:"userDrawn",attr`
}
type CT_SlideMasterTextStyles struct{
	titleStyle string  `xml:"titleStyle"`
	bodyStyle string  `xml:"bodyStyle"`
	otherStyle string  `xml:"otherStyle"`
	extLst string  `xml:"extLst"`
}
type CT_SlideLayoutIdListEntry struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
}
type CT_SlideLayoutIdList struct{
	sldLayoutId []string  `xml:"sldLayoutId"`
}
type CT_SlideMaster struct{
	cSld string  `xml:"cSld"`
	sldLayoutIdLst string  `xml:"sldLayoutIdLst"`
	transition string  `xml:"transition"`
	timing string  `xml:"timing"`
	hf string  `xml:"hf"`
	txStyles string  `xml:"txStyles"`
	extLst string  `xml:"extLst"`
	 EG_TopLevelSlide  `xml:""`
	preserve string  `xml:"preserve",attr`
}
type CT_HandoutMaster struct{
	cSld string  `xml:"cSld"`
	hf string  `xml:"hf"`
	extLst string  `xml:"extLst"`
	 EG_TopLevelSlide  `xml:""`
}
type CT_NotesMaster struct{
	cSld string  `xml:"cSld"`
	hf string  `xml:"hf"`
	notesStyle string  `xml:"notesStyle"`
	extLst string  `xml:"extLst"`
	 EG_TopLevelSlide  `xml:""`
}
type CT_NotesSlide struct{
	cSld string  `xml:"cSld"`
	extLst string  `xml:"extLst"`
	 EG_ChildSlide  `xml:""`
}
type CT_SlideSyncProperties struct{
	extLst string  `xml:"extLst"`
	serverSldId string  `xml:"serverSldId",attr`
	serverSldModifiedTime string  `xml:"serverSldModifiedTime",attr`
	clientInsertedTime string  `xml:"clientInsertedTime",attr`
}
type CT_StringTag struct{
	name string  `xml:"name",attr`
	val string  `xml:"val",attr`
}
type CT_TagList struct{
	tag []CT_StringTag  `xml:"tag"`
}
type CT_NormalViewPortion struct{
	sz string  `xml:"sz",attr`
	autoAdjust string  `xml:"autoAdjust",attr`
}
type CT_NormalViewProperties struct{
	restoredLeft CT_NormalViewPortion  `xml:"restoredLeft"`
	restoredTop CT_NormalViewPortion  `xml:"restoredTop"`
	extLst string  `xml:"extLst"`
	showOutlineIcons string  `xml:"showOutlineIcons",attr`
	snapVertSplitter string  `xml:"snapVertSplitter",attr`
	vertBarState ST_SplitterBarState  `xml:"vertBarState",attr`
	horzBarState ST_SplitterBarState  `xml:"horzBarState",attr`
	preferSingleView string  `xml:"preferSingleView",attr`
}
type CT_CommonViewProperties struct{
	scale string  `xml:"scale"`
	origin string  `xml:"origin"`
	varScale string  `xml:"varScale",attr`
}
type CT_NotesTextViewProperties struct{
	cViewPr string  `xml:"cViewPr"`
	extLst string  `xml:"extLst"`
}
type CT_OutlineViewSlideEntry struct{
	collapse string  `xml:"collapse",attr`
}
type CT_OutlineViewSlideList struct{
	sld []string  `xml:"sld"`
}
type CT_OutlineViewProperties struct{
	cViewPr string  `xml:"cViewPr"`
	sldLst string  `xml:"sldLst"`
	extLst string  `xml:"extLst"`
}
type CT_SlideSorterViewProperties struct{
	cViewPr string  `xml:"cViewPr"`
	extLst string  `xml:"extLst"`
	showFormatting string  `xml:"showFormatting",attr`
}
type CT_Guide struct{
	orient ST_Direction  `xml:"orient",attr`
	pos string  `xml:"pos",attr`
}
type CT_GuideList struct{
	guide []string  `xml:"guide"`
}
type CT_CommonSlideViewProperties struct{
	cViewPr string  `xml:"cViewPr"`
	guideLst string  `xml:"guideLst"`
	snapToGrid string  `xml:"snapToGrid",attr`
	snapToObjects string  `xml:"snapToObjects",attr`
	showGuides string  `xml:"showGuides",attr`
}
type CT_SlideViewProperties struct{
	cSldViewPr string  `xml:"cSldViewPr"`
	extLst string  `xml:"extLst"`
}
type CT_NotesViewProperties struct{
	cSldViewPr string  `xml:"cSldViewPr"`
	extLst string  `xml:"extLst"`
}
type CT_ViewProperties struct{
	normalViewPr string  `xml:"normalViewPr"`
	slideViewPr string  `xml:"slideViewPr"`
	outlineViewPr string  `xml:"outlineViewPr"`
	notesTextViewPr string  `xml:"notesTextViewPr"`
	sorterViewPr string  `xml:"sorterViewPr"`
	notesViewPr string  `xml:"notesViewPr"`
	gridSpacing string  `xml:"gridSpacing"`
	extLst string  `xml:"extLst"`
	lastView ST_ViewType  `xml:"lastView",attr`
	showComments string  `xml:"showComments",attr`
}
