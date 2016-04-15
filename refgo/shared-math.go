package ooxml
type WordSchema  struct {
	mathPr CT_MathPr  `xml:"mathPr"`
	oMathPara CT_OMathPara  `xml:"oMathPara"`
	oMath CT_OMath  `xml:"oMath"`
}
type CT_Integer255 struct{
	val ST_Integer255  `xml:"val",attr`
}
type CT_Integer2 struct{
	val ST_Integer2  `xml:"val",attr`
}
type CT_SpacingRule struct{
	val ST_SpacingRule  `xml:"val",attr`
}
type CT_UnSignedInteger struct{
	val string  `xml:"val",attr`
}
type CT_Char struct{
	val ST_Char  `xml:"val",attr`
}
type CT_OnOff struct{
	val string  `xml:"val",attr`
}
type CT_String struct{
	val string  `xml:"val",attr`
}
type CT_XAlign struct{
	val string  `xml:"val",attr`
}
type CT_YAlign struct{
	val string  `xml:"val",attr`
}
type CT_Shp struct{
	val ST_Shp  `xml:"val",attr`
}
type CT_FType struct{
	val ST_FType  `xml:"val",attr`
}
type CT_LimLoc struct{
	val ST_LimLoc  `xml:"val",attr`
}
type CT_TopBot struct{
	val ST_TopBot  `xml:"val",attr`
}
type CT_Script struct{
	val ST_Script  `xml:"val",attr`
}
type CT_Style struct{
	val ST_Style  `xml:"val",attr`
}
type CT_ManualBreak struct{
	alnAt ST_Integer255  `xml:"alnAt",attr`
}
type CT_RPR struct{
	lit CT_OnOff  `xml:"lit"`
	brk CT_ManualBreak  `xml:"brk"`
	aln CT_OnOff  `xml:"aln"`
	nor CT_OnOff  `xml:"nor"`
}
type CT_Text struct{
}
type CT_R struct{
	rPr CT_RPR  `xml:"rPr"`
	 w:EG_RPr  `xml:""`
	 w:EG_RunInnerContent  `xml:""`
	t string  `xml:"t"`
}
type CT_CtrlPr struct{
	 w:EG_RPrMath  `xml:""`
}
type CT_AccPr struct{
	chr CT_Char  `xml:"chr"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Acc struct{
	accPr CT_AccPr  `xml:"accPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_BarPr struct{
	pos CT_TopBot  `xml:"pos"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Bar struct{
	barPr CT_BarPr  `xml:"barPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_BoxPr struct{
	opEmu CT_OnOff  `xml:"opEmu"`
	noBreak CT_OnOff  `xml:"noBreak"`
	diff CT_OnOff  `xml:"diff"`
	brk CT_ManualBreak  `xml:"brk"`
	aln CT_OnOff  `xml:"aln"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Box struct{
	boxPr string  `xml:"boxPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_BorderBoxPr struct{
	hideTop CT_OnOff  `xml:"hideTop"`
	hideBot CT_OnOff  `xml:"hideBot"`
	hideLeft CT_OnOff  `xml:"hideLeft"`
	hideRight CT_OnOff  `xml:"hideRight"`
	strikeH CT_OnOff  `xml:"strikeH"`
	strikeV CT_OnOff  `xml:"strikeV"`
	strikeBLTR CT_OnOff  `xml:"strikeBLTR"`
	strikeTLBR CT_OnOff  `xml:"strikeTLBR"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_BorderBox struct{
	borderBoxPr string  `xml:"borderBoxPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_DPr struct{
	begChr CT_Char  `xml:"begChr"`
	sepChr CT_Char  `xml:"sepChr"`
	endChr CT_Char  `xml:"endChr"`
	grow CT_OnOff  `xml:"grow"`
	shp CT_Shp  `xml:"shp"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_D struct{
	dPr CT_DPr  `xml:"dPr"`
	e []CT_OMathArg  `xml:"e"`
}
type CT_EqArrPr struct{
	baseJc CT_YAlign  `xml:"baseJc"`
	maxDist CT_OnOff  `xml:"maxDist"`
	objDist CT_OnOff  `xml:"objDist"`
	rSpRule CT_SpacingRule  `xml:"rSpRule"`
	rSp string  `xml:"rSp"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_EqArr struct{
	eqArrPr CT_EqArrPr  `xml:"eqArrPr"`
	e []CT_OMathArg  `xml:"e"`
}
type CT_FPr struct{
	type CT_FType  `xml:"type"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_F struct{
	fPr CT_FPr  `xml:"fPr"`
	num CT_OMathArg  `xml:"num"`
	den CT_OMathArg  `xml:"den"`
}
type CT_FuncPr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Func struct{
	funcPr CT_FuncPr  `xml:"funcPr"`
	fName CT_OMathArg  `xml:"fName"`
	e CT_OMathArg  `xml:"e"`
}
type CT_GroupChrPr struct{
	chr CT_Char  `xml:"chr"`
	pos CT_TopBot  `xml:"pos"`
	vertJc CT_TopBot  `xml:"vertJc"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_GroupChr struct{
	groupChrPr CT_GroupChrPr  `xml:"groupChrPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_LimLowPr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_LimLow struct{
	limLowPr CT_LimLowPr  `xml:"limLowPr"`
	e CT_OMathArg  `xml:"e"`
	lim CT_OMathArg  `xml:"lim"`
}
type CT_LimUppPr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_LimUpp struct{
	limUppPr CT_LimUppPr  `xml:"limUppPr"`
	e CT_OMathArg  `xml:"e"`
	lim CT_OMathArg  `xml:"lim"`
}
type CT_MCPr struct{
	count CT_Integer255  `xml:"count"`
	mcJc CT_XAlign  `xml:"mcJc"`
}
type CT_MC struct{
	mcPr CT_MCPr  `xml:"mcPr"`
}
type CT_MCS struct{
	mc []CT_MC  `xml:"mc"`
}
type CT_MPr struct{
	baseJc CT_YAlign  `xml:"baseJc"`
	plcHide CT_OnOff  `xml:"plcHide"`
	rSpRule CT_SpacingRule  `xml:"rSpRule"`
	cGpRule CT_SpacingRule  `xml:"cGpRule"`
	rSp string  `xml:"rSp"`
	cSp string  `xml:"cSp"`
	cGp string  `xml:"cGp"`
	mcs CT_MCS  `xml:"mcs"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_MR struct{
	e []CT_OMathArg  `xml:"e"`
}
type CT_M struct{
	mPr CT_MPr  `xml:"mPr"`
	mr []CT_MR  `xml:"mr"`
}
type CT_NaryPr struct{
	chr CT_Char  `xml:"chr"`
	limLoc CT_LimLoc  `xml:"limLoc"`
	grow CT_OnOff  `xml:"grow"`
	subHide CT_OnOff  `xml:"subHide"`
	supHide CT_OnOff  `xml:"supHide"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Nary struct{
	naryPr CT_NaryPr  `xml:"naryPr"`
	sub CT_OMathArg  `xml:"sub"`
	sup CT_OMathArg  `xml:"sup"`
	e CT_OMathArg  `xml:"e"`
}
type CT_PhantPr struct{
	show CT_OnOff  `xml:"show"`
	zeroWid CT_OnOff  `xml:"zeroWid"`
	zeroAsc CT_OnOff  `xml:"zeroAsc"`
	zeroDesc CT_OnOff  `xml:"zeroDesc"`
	transp CT_OnOff  `xml:"transp"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Phant struct{
	phantPr CT_PhantPr  `xml:"phantPr"`
	e CT_OMathArg  `xml:"e"`
}
type CT_RadPr struct{
	degHide CT_OnOff  `xml:"degHide"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_Rad struct{
	radPr string  `xml:"radPr"`
	deg CT_OMathArg  `xml:"deg"`
	e CT_OMathArg  `xml:"e"`
}
type CT_SPrePr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_SPre struct{
	sPrePr CT_SPrePr  `xml:"sPrePr"`
	sub CT_OMathArg  `xml:"sub"`
	sup CT_OMathArg  `xml:"sup"`
	e CT_OMathArg  `xml:"e"`
}
type CT_SSubPr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_SSub struct{
	sSubPr CT_SSubPr  `xml:"sSubPr"`
	e CT_OMathArg  `xml:"e"`
	sub CT_OMathArg  `xml:"sub"`
}
type CT_SSubSupPr struct{
	alnScr CT_OnOff  `xml:"alnScr"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_SSubSup struct{
	sSubSupPr CT_SSubSupPr  `xml:"sSubSupPr"`
	e CT_OMathArg  `xml:"e"`
	sub CT_OMathArg  `xml:"sub"`
	sup CT_OMathArg  `xml:"sup"`
}
type CT_SSupPr struct{
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
}
type CT_SSup struct{
	sSupPr CT_SSupPr  `xml:"sSupPr"`
	e CT_OMathArg  `xml:"e"`
	sup CT_OMathArg  `xml:"sup"`
}
type CT_OMathArgPr struct{
	argSz CT_Integer2  `xml:"argSz"`
}
type CT_OMathArg struct{
	argPr CT_OMathArgPr  `xml:"argPr"`
	ctrlPr CT_CtrlPr  `xml:"ctrlPr"`
	 []EG_OMathElements  `xml:""`
}
type CT_OMathJc struct{
	val ST_Jc  `xml:"val",attr`
}
type CT_OMathParaPr struct{
	jc CT_OMathJc  `xml:"jc"`
}
type CT_TwipsMeasure struct{
	val string  `xml:"val",attr`
}
type CT_BreakBin struct{
	val ST_BreakBin  `xml:"val",attr`
}
type CT_BreakBinSub struct{
	val ST_BreakBinSub  `xml:"val",attr`
}
type CT_MathPr struct{
	mathFont CT_String  `xml:"mathFont"`
	brkBin CT_BreakBin  `xml:"brkBin"`
	brkBinSub CT_BreakBinSub  `xml:"brkBinSub"`
	smallFrac CT_OnOff  `xml:"smallFrac"`
	dispDef CT_OnOff  `xml:"dispDef"`
	lMargin string  `xml:"lMargin"`
	rMargin string  `xml:"rMargin"`
	defJc CT_OMathJc  `xml:"defJc"`
	preSp string  `xml:"preSp"`
	postSp string  `xml:"postSp"`
	interSp string  `xml:"interSp"`
	intraSp string  `xml:"intraSp"`
	intLim CT_LimLoc  `xml:"intLim"`
	naryLim CT_LimLoc  `xml:"naryLim"`
	wrapIndent string  `xml:"wrapIndent"`
	wrapRight CT_OnOff  `xml:"wrapRight"`
}
type CT_OMathPara struct{
	oMathParaPr CT_OMathParaPr  `xml:"oMathParaPr"`
	oMath []CT_OMath  `xml:"oMath"`
}
type CT_OMath struct{
	 []EG_OMathElements  `xml:""`
}
