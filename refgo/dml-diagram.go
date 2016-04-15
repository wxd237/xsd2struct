package ooxml
type WordSchema  struct {
	colorsDef string  `xml:"colorsDef"`
	colorsDefHdr string  `xml:"colorsDefHdr"`
	colorsDefHdrLst string  `xml:"colorsDefHdrLst"`
	dataModel string  `xml:"dataModel"`
	layoutDef CT_DiagramDefinition  `xml:"layoutDef"`
	layoutDefHdr string  `xml:"layoutDefHdr"`
	layoutDefHdrLst string  `xml:"layoutDefHdrLst"`
	relIds string  `xml:"relIds"`
	styleDef CT_StyleDefinition  `xml:"styleDef"`
	styleDefHdr string  `xml:"styleDefHdr"`
	styleDefHdrLst string  `xml:"styleDefHdrLst"`
}
type CT_CTName struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_CTDescription struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_CTCategory struct{
	type string  `xml:"type",attr`
	pri string  `xml:"pri",attr`
}
type CT_CTCategories struct{
	cat []CT_CTCategory  `xml:"cat"`
}
type CT_Colors struct{
	 []a:EG_ColorChoice  `xml:""`
	meth string  `xml:"meth",attr`
	hueDir ST_HueDir  `xml:"hueDir",attr`
}
type CT_CTStyleLabel struct{
	fillClrLst string  `xml:"fillClrLst"`
	linClrLst string  `xml:"linClrLst"`
	effectClrLst string  `xml:"effectClrLst"`
	txLinClrLst string  `xml:"txLinClrLst"`
	txFillClrLst string  `xml:"txFillClrLst"`
	txEffectClrLst string  `xml:"txEffectClrLst"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_ColorTransform struct{
	title []CT_CTName  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	styleLbl []CT_CTStyleLabel  `xml:"styleLbl"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
}
type CT_ColorTransformHeader struct{
	title []CT_CTName  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
	resId string  `xml:"resId",attr`
}
type CT_ColorTransformHeaderLst struct{
	colorsDefHdr []string  `xml:"colorsDefHdr"`
}
type CT_Pt struct{
	prSet CT_ElemPropSet  `xml:"prSet"`
	spPr string  `xml:"spPr"`
	t string  `xml:"t"`
	extLst string  `xml:"extLst"`
	modelId string  `xml:"modelId",attr`
	type ST_PtType  `xml:"type",attr`
	cxnId string  `xml:"cxnId",attr`
}
type CT_PtList struct{
	pt []CT_Pt  `xml:"pt"`
}
type CT_Cxn struct{
	extLst string  `xml:"extLst"`
	modelId string  `xml:"modelId",attr`
	type string  `xml:"type",attr`
	srcId string  `xml:"srcId",attr`
	destId string  `xml:"destId",attr`
	srcOrd string  `xml:"srcOrd",attr`
	destOrd string  `xml:"destOrd",attr`
	parTransId string  `xml:"parTransId",attr`
	sibTransId string  `xml:"sibTransId",attr`
	presId string  `xml:"presId",attr`
}
type CT_CxnList struct{
	cxn []string  `xml:"cxn"`
}
type CT_DataModel struct{
	ptLst string  `xml:"ptLst"`
	cxnLst string  `xml:"cxnLst"`
	bg string  `xml:"bg"`
	whole string  `xml:"whole"`
	extLst string  `xml:"extLst"`
}
type CT_Constraint struct{
	extLst string  `xml:"extLst"`
	op ST_BoolOperator  `xml:"op",attr`
	val string  `xml:"val",attr`
	fact string  `xml:"fact",attr`
}
type CT_Constraints struct{
	constr []string  `xml:"constr"`
}
type CT_NumericRule struct{
	extLst string  `xml:"extLst"`
	val string  `xml:"val",attr`
	fact string  `xml:"fact",attr`
	max string  `xml:"max",attr`
}
type CT_Rules struct{
	rule []CT_NumericRule  `xml:"rule"`
}
type CT_PresentationOf struct{
	extLst string  `xml:"extLst"`
}
type CT_Adj struct{
	idx string  `xml:"idx",attr`
	val string  `xml:"val",attr`
}
type CT_AdjLst struct{
	adj []string  `xml:"adj"`
}
type CT_Shape struct{
	adjLst string  `xml:"adjLst"`
	extLst string  `xml:"extLst"`
	rot string  `xml:"rot",attr`
	type ST_LayoutShapeType  `xml:"type",attr`
	zOrderOff string  `xml:"zOrderOff",attr`
	hideGeom string  `xml:"hideGeom",attr`
	lkTxEntry string  `xml:"lkTxEntry",attr`
	blipPhldr string  `xml:"blipPhldr",attr`
}
type CT_Parameter struct{
	type string  `xml:"type",attr`
	val ST_ParameterVal  `xml:"val",attr`
}
type CT_Algorithm struct{
	param []CT_Parameter  `xml:"param"`
	extLst string  `xml:"extLst"`
	type ST_AlgorithmType  `xml:"type",attr`
	rev string  `xml:"rev",attr`
}
type CT_LayoutNode struct{
	alg CT_Algorithm  `xml:"alg"`
	shape CT_Shape  `xml:"shape"`
	presOf string  `xml:"presOf"`
	constrLst string  `xml:"constrLst"`
	ruleLst string  `xml:"ruleLst"`
	varLst CT_LayoutVariablePropertySet  `xml:"varLst"`
	forEach CT_ForEach  `xml:"forEach"`
	layoutNode string  `xml:"layoutNode"`
	choose string  `xml:"choose"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	styleLbl string  `xml:"styleLbl",attr`
	chOrder string  `xml:"chOrder",attr`
	moveWith string  `xml:"moveWith",attr`
}
type CT_ForEach struct{
	alg CT_Algorithm  `xml:"alg"`
	shape CT_Shape  `xml:"shape"`
	presOf string  `xml:"presOf"`
	constrLst string  `xml:"constrLst"`
	ruleLst string  `xml:"ruleLst"`
	forEach CT_ForEach  `xml:"forEach"`
	layoutNode string  `xml:"layoutNode"`
	choose string  `xml:"choose"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	ref string  `xml:"ref",attr`
}
type CT_When struct{
	alg CT_Algorithm  `xml:"alg"`
	shape CT_Shape  `xml:"shape"`
	presOf string  `xml:"presOf"`
	constrLst string  `xml:"constrLst"`
	ruleLst string  `xml:"ruleLst"`
	forEach CT_ForEach  `xml:"forEach"`
	layoutNode string  `xml:"layoutNode"`
	choose string  `xml:"choose"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	func ST_FunctionType  `xml:"func",attr`
	arg ST_FunctionArgument  `xml:"arg",attr`
	op ST_FunctionOperator  `xml:"op",attr`
	val ST_FunctionValue  `xml:"val",attr`
}
type CT_Otherwise struct{
	alg CT_Algorithm  `xml:"alg"`
	shape CT_Shape  `xml:"shape"`
	presOf string  `xml:"presOf"`
	constrLst string  `xml:"constrLst"`
	ruleLst string  `xml:"ruleLst"`
	forEach CT_ForEach  `xml:"forEach"`
	layoutNode string  `xml:"layoutNode"`
	choose string  `xml:"choose"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_Choose struct{
	if []CT_When  `xml:"if"`
	else string  `xml:"else"`
	name string  `xml:"name",attr`
}
type CT_SampleData struct{
	dataModel string  `xml:"dataModel"`
	useDef string  `xml:"useDef",attr`
}
type CT_Category struct{
	type string  `xml:"type",attr`
	pri string  `xml:"pri",attr`
}
type CT_Categories struct{
	cat []CT_Category  `xml:"cat"`
}
type CT_Name struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_Description struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_DiagramDefinition struct{
	title []CT_Name  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	sampData CT_SampleData  `xml:"sampData"`
	styleData CT_SampleData  `xml:"styleData"`
	clrData CT_SampleData  `xml:"clrData"`
	layoutNode string  `xml:"layoutNode"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
	defStyle string  `xml:"defStyle",attr`
}
type CT_DiagramDefinitionHeader struct{
	title []CT_Name  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
	defStyle string  `xml:"defStyle",attr`
	resId string  `xml:"resId",attr`
}
type CT_DiagramDefinitionHeaderLst struct{
	layoutDefHdr []string  `xml:"layoutDefHdr"`
}
type CT_RelIds struct{
}
type CT_ElemPropSet struct{
	presLayoutVars CT_LayoutVariablePropertySet  `xml:"presLayoutVars"`
	style string  `xml:"style"`
	presAssocID string  `xml:"presAssocID",attr`
	presName string  `xml:"presName",attr`
	presStyleLbl string  `xml:"presStyleLbl",attr`
	presStyleIdx string  `xml:"presStyleIdx",attr`
	presStyleCnt string  `xml:"presStyleCnt",attr`
	loTypeId string  `xml:"loTypeId",attr`
	loCatId string  `xml:"loCatId",attr`
	qsTypeId string  `xml:"qsTypeId",attr`
	qsCatId string  `xml:"qsCatId",attr`
	csTypeId string  `xml:"csTypeId",attr`
	csCatId string  `xml:"csCatId",attr`
	coherent3DOff string  `xml:"coherent3DOff",attr`
	phldrT string  `xml:"phldrT",attr`
	phldr string  `xml:"phldr",attr`
	custAng string  `xml:"custAng",attr`
	custFlipVert string  `xml:"custFlipVert",attr`
	custFlipHor string  `xml:"custFlipHor",attr`
	custSzX string  `xml:"custSzX",attr`
	custSzY string  `xml:"custSzY",attr`
	custScaleX string  `xml:"custScaleX",attr`
	custScaleY string  `xml:"custScaleY",attr`
	custT string  `xml:"custT",attr`
	custLinFactX string  `xml:"custLinFactX",attr`
	custLinFactY string  `xml:"custLinFactY",attr`
	custLinFactNeighborX string  `xml:"custLinFactNeighborX",attr`
	custLinFactNeighborY string  `xml:"custLinFactNeighborY",attr`
	custRadScaleRad string  `xml:"custRadScaleRad",attr`
	custRadScaleInc string  `xml:"custRadScaleInc",attr`
}
type CT_OrgChart struct{
	val string  `xml:"val",attr`
}
type CT_ChildMax struct{
	val string  `xml:"val",attr`
}
type CT_ChildPref struct{
	val string  `xml:"val",attr`
}
type CT_BulletEnabled struct{
	val string  `xml:"val",attr`
}
type CT_Direction struct{
	val ST_Direction  `xml:"val",attr`
}
type CT_HierBranchStyle struct{
	val ST_HierBranchStyle  `xml:"val",attr`
}
type CT_AnimOne struct{
	val ST_AnimOneStr  `xml:"val",attr`
}
type CT_AnimLvl struct{
	val ST_AnimLvlStr  `xml:"val",attr`
}
type CT_ResizeHandles struct{
	val string  `xml:"val",attr`
}
type CT_LayoutVariablePropertySet struct{
	orgChart CT_OrgChart  `xml:"orgChart"`
	chMax string  `xml:"chMax"`
	chPref string  `xml:"chPref"`
	bulletEnabled string  `xml:"bulletEnabled"`
	dir CT_Direction  `xml:"dir"`
	hierBranch CT_HierBranchStyle  `xml:"hierBranch"`
	animOne CT_AnimOne  `xml:"animOne"`
	animLvl CT_AnimLvl  `xml:"animLvl"`
	resizeHandles string  `xml:"resizeHandles"`
}
type CT_SDName struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_SDDescription struct{
	lang string  `xml:"lang",attr`
	val string  `xml:"val",attr`
}
type CT_SDCategory struct{
	type string  `xml:"type",attr`
	pri string  `xml:"pri",attr`
}
type CT_SDCategories struct{
	cat []CT_SDCategory  `xml:"cat"`
}
type CT_TextProps struct{
	 a:EG_Text3D  `xml:""`
}
type CT_StyleLabel struct{
	scene3d string  `xml:"scene3d"`
	sp3d string  `xml:"sp3d"`
	txPr string  `xml:"txPr"`
	style string  `xml:"style"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
}
type CT_StyleDefinition struct{
	title []CT_SDName  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	scene3d string  `xml:"scene3d"`
	styleLbl []CT_StyleLabel  `xml:"styleLbl"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
}
type CT_StyleDefinitionHeader struct{
	title []CT_SDName  `xml:"title"`
	desc []string  `xml:"desc"`
	catLst string  `xml:"catLst"`
	extLst string  `xml:"extLst"`
	uniqueId string  `xml:"uniqueId",attr`
	minVer string  `xml:"minVer",attr`
	resId string  `xml:"resId",attr`
}
type CT_StyleDefinitionHeaderLst struct{
	styleDefHdr []string  `xml:"styleDefHdr"`
}
