package ooxml
type WordSchema  struct {
	chartSpace CT_ChartSpace  `xml:"chartSpace"`
	userShapes string  `xml:"userShapes"`
	chart string  `xml:"chart"`
}
type CT_Boolean struct{
	val string  `xml:"val",attr`
}
type CT_Double struct{
	val string  `xml:"val",attr`
}
type CT_UnsignedInt struct{
	val string  `xml:"val",attr`
}
type CT_RelId struct{
}
type CT_Extension struct{
	uri string  `xml:"uri",attr`
}
type CT_ExtensionList struct{
	ext []string  `xml:"ext"`
}
type CT_NumVal struct{
	v string  `xml:"v"`
	idx string  `xml:"idx",attr`
	formatCode string  `xml:"formatCode",attr`
}
type CT_NumData struct{
	formatCode string  `xml:"formatCode"`
	ptCount string  `xml:"ptCount"`
	pt []CT_NumVal  `xml:"pt"`
	extLst string  `xml:"extLst"`
}
type CT_NumRef struct{
	f string  `xml:"f"`
	numCache CT_NumData  `xml:"numCache"`
	extLst string  `xml:"extLst"`
}
type CT_NumDataSource struct{
	numRef CT_NumRef  `xml:"numRef"`
	numLit CT_NumData  `xml:"numLit"`
}
type CT_StrVal struct{
	v string  `xml:"v"`
	idx string  `xml:"idx",attr`
}
type CT_StrData struct{
	ptCount string  `xml:"ptCount"`
	pt []CT_StrVal  `xml:"pt"`
	extLst string  `xml:"extLst"`
}
type CT_StrRef struct{
	f string  `xml:"f"`
	strCache CT_StrData  `xml:"strCache"`
	extLst string  `xml:"extLst"`
}
type CT_Tx struct{
	strRef CT_StrRef  `xml:"strRef"`
	rich string  `xml:"rich"`
}
type CT_TextLanguageID struct{
	val string  `xml:"val",attr`
}
type CT_Lvl struct{
	pt []CT_StrVal  `xml:"pt"`
}
type CT_MultiLvlStrData struct{
	ptCount string  `xml:"ptCount"`
	lvl []CT_Lvl  `xml:"lvl"`
	extLst string  `xml:"extLst"`
}
type CT_MultiLvlStrRef struct{
	f string  `xml:"f"`
	multiLvlStrCache CT_MultiLvlStrData  `xml:"multiLvlStrCache"`
	extLst string  `xml:"extLst"`
}
type CT_AxDataSource struct{
	multiLvlStrRef CT_MultiLvlStrRef  `xml:"multiLvlStrRef"`
	numRef CT_NumRef  `xml:"numRef"`
	numLit CT_NumData  `xml:"numLit"`
	strRef CT_StrRef  `xml:"strRef"`
	strLit CT_StrData  `xml:"strLit"`
}
type CT_SerTx struct{
	strRef CT_StrRef  `xml:"strRef"`
	v string  `xml:"v"`
}
type CT_LayoutTarget struct{
	val ST_LayoutTarget  `xml:"val",attr`
}
type CT_LayoutMode struct{
	val string  `xml:"val",attr`
}
type CT_ManualLayout struct{
	layoutTarget CT_LayoutTarget  `xml:"layoutTarget"`
	xMode string  `xml:"xMode"`
	yMode string  `xml:"yMode"`
	wMode string  `xml:"wMode"`
	hMode string  `xml:"hMode"`
	x CT_Double  `xml:"x"`
	y CT_Double  `xml:"y"`
	w CT_Double  `xml:"w"`
	h CT_Double  `xml:"h"`
	extLst string  `xml:"extLst"`
}
type CT_Layout struct{
	manualLayout CT_ManualLayout  `xml:"manualLayout"`
	extLst string  `xml:"extLst"`
}
type CT_Title struct{
	tx string  `xml:"tx"`
	layout CT_Layout  `xml:"layout"`
	overlay CT_Boolean  `xml:"overlay"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	extLst string  `xml:"extLst"`
}
type CT_RotX struct{
	val ST_RotX  `xml:"val",attr`
}
type CT_HPercent struct{
	val ST_HPercent  `xml:"val",attr`
}
type CT_RotY struct{
	val ST_RotY  `xml:"val",attr`
}
type CT_DepthPercent struct{
	val ST_DepthPercent  `xml:"val",attr`
}
type CT_Perspective struct{
	val string  `xml:"val",attr`
}
type CT_View3D struct{
	rotX CT_RotX  `xml:"rotX"`
	hPercent CT_HPercent  `xml:"hPercent"`
	rotY CT_RotY  `xml:"rotY"`
	depthPercent CT_DepthPercent  `xml:"depthPercent"`
	rAngAx CT_Boolean  `xml:"rAngAx"`
	perspective string  `xml:"perspective"`
	extLst string  `xml:"extLst"`
}
type CT_Surface struct{
	thickness string  `xml:"thickness"`
	spPr string  `xml:"spPr"`
	pictureOptions string  `xml:"pictureOptions"`
	extLst string  `xml:"extLst"`
}
type CT_Thickness struct{
	val string  `xml:"val",attr`
}
type CT_DTable struct{
	showHorzBorder CT_Boolean  `xml:"showHorzBorder"`
	showVertBorder CT_Boolean  `xml:"showVertBorder"`
	showOutline CT_Boolean  `xml:"showOutline"`
	showKeys CT_Boolean  `xml:"showKeys"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	extLst string  `xml:"extLst"`
}
type CT_GapAmount struct{
	val ST_GapAmount  `xml:"val",attr`
}
type CT_Overlap struct{
	val ST_Overlap  `xml:"val",attr`
}
type CT_BubbleScale struct{
	val ST_BubbleScale  `xml:"val",attr`
}
type CT_SizeRepresents struct{
	val string  `xml:"val",attr`
}
type CT_FirstSliceAng struct{
	val string  `xml:"val",attr`
}
type CT_HoleSize struct{
	val ST_HoleSize  `xml:"val",attr`
}
type CT_SplitType struct{
	val ST_SplitType  `xml:"val",attr`
}
type CT_CustSplit struct{
	secondPiePt []string  `xml:"secondPiePt"`
}
type CT_SecondPieSize struct{
	val string  `xml:"val",attr`
}
type CT_NumFmt struct{
	formatCode string  `xml:"formatCode",attr`
	sourceLinked string  `xml:"sourceLinked",attr`
}
type CT_LblAlgn struct{
	val ST_LblAlgn  `xml:"val",attr`
}
type CT_DLblPos struct{
	val string  `xml:"val",attr`
}
type CT_DLbl struct{
	idx string  `xml:"idx"`
	extLst string  `xml:"extLst"`
	 Group_DLbl  `xml:""`
	delete CT_Boolean  `xml:"delete"`
}
type CT_DLbls struct{
	dLbl []CT_DLbl  `xml:"dLbl"`
	extLst string  `xml:"extLst"`
	 Group_DLbls  `xml:""`
	delete CT_Boolean  `xml:"delete"`
}
type CT_MarkerStyle struct{
	val ST_MarkerStyle  `xml:"val",attr`
}
type CT_MarkerSize struct{
	val ST_MarkerSize  `xml:"val",attr`
}
type CT_Marker struct{
	symbol CT_MarkerStyle  `xml:"symbol"`
	size CT_MarkerSize  `xml:"size"`
	spPr string  `xml:"spPr"`
	extLst string  `xml:"extLst"`
}
type CT_DPt struct{
	idx string  `xml:"idx"`
	invertIfNegative CT_Boolean  `xml:"invertIfNegative"`
	marker CT_Marker  `xml:"marker"`
	bubble3D CT_Boolean  `xml:"bubble3D"`
	explosion string  `xml:"explosion"`
	spPr string  `xml:"spPr"`
	pictureOptions string  `xml:"pictureOptions"`
	extLst string  `xml:"extLst"`
}
type CT_TrendlineType struct{
	val string  `xml:"val",attr`
}
type CT_Order struct{
	val string  `xml:"val",attr`
}
type CT_Period struct{
	val string  `xml:"val",attr`
}
type CT_TrendlineLbl struct{
	layout CT_Layout  `xml:"layout"`
	tx string  `xml:"tx"`
	numFmt CT_NumFmt  `xml:"numFmt"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	extLst string  `xml:"extLst"`
}
type CT_Trendline struct{
	name string  `xml:"name"`
	spPr string  `xml:"spPr"`
	trendlineType string  `xml:"trendlineType"`
	order string  `xml:"order"`
	period string  `xml:"period"`
	forward CT_Double  `xml:"forward"`
	backward CT_Double  `xml:"backward"`
	intercept CT_Double  `xml:"intercept"`
	dispRSqr CT_Boolean  `xml:"dispRSqr"`
	dispEq CT_Boolean  `xml:"dispEq"`
	trendlineLbl string  `xml:"trendlineLbl"`
	extLst string  `xml:"extLst"`
}
type CT_ErrDir struct{
	val ST_ErrDir  `xml:"val",attr`
}
type CT_ErrBarType struct{
	val ST_ErrBarType  `xml:"val",attr`
}
type CT_ErrValType struct{
	val ST_ErrValType  `xml:"val",attr`
}
type CT_ErrBars struct{
	errDir CT_ErrDir  `xml:"errDir"`
	errBarType CT_ErrBarType  `xml:"errBarType"`
	errValType CT_ErrValType  `xml:"errValType"`
	noEndCap CT_Boolean  `xml:"noEndCap"`
	plus CT_NumDataSource  `xml:"plus"`
	minus CT_NumDataSource  `xml:"minus"`
	val CT_Double  `xml:"val"`
	spPr string  `xml:"spPr"`
	extLst string  `xml:"extLst"`
}
type CT_UpDownBar struct{
	spPr string  `xml:"spPr"`
}
type CT_UpDownBars struct{
	gapWidth CT_GapAmount  `xml:"gapWidth"`
	upBars CT_UpDownBar  `xml:"upBars"`
	downBars CT_UpDownBar  `xml:"downBars"`
	extLst string  `xml:"extLst"`
}
type CT_LineSer struct{
	marker CT_Marker  `xml:"marker"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	trendline []string  `xml:"trendline"`
	errBars string  `xml:"errBars"`
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	smooth CT_Boolean  `xml:"smooth"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_ScatterSer struct{
	marker CT_Marker  `xml:"marker"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	trendline []string  `xml:"trendline"`
	errBars string  `xml:"errBars"`
	xVal string  `xml:"xVal"`
	yVal CT_NumDataSource  `xml:"yVal"`
	smooth CT_Boolean  `xml:"smooth"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_RadarSer struct{
	marker CT_Marker  `xml:"marker"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_BarSer struct{
	invertIfNegative CT_Boolean  `xml:"invertIfNegative"`
	pictureOptions string  `xml:"pictureOptions"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	trendline []string  `xml:"trendline"`
	errBars string  `xml:"errBars"`
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	shape CT_Shape  `xml:"shape"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_AreaSer struct{
	pictureOptions string  `xml:"pictureOptions"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	trendline []string  `xml:"trendline"`
	errBars string  `xml:"errBars"`
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_PieSer struct{
	explosion string  `xml:"explosion"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_BubbleSer struct{
	invertIfNegative CT_Boolean  `xml:"invertIfNegative"`
	dPt []CT_DPt  `xml:"dPt"`
	dLbls string  `xml:"dLbls"`
	trendline []string  `xml:"trendline"`
	errBars string  `xml:"errBars"`
	xVal string  `xml:"xVal"`
	yVal CT_NumDataSource  `xml:"yVal"`
	bubbleSize CT_NumDataSource  `xml:"bubbleSize"`
	bubble3D CT_Boolean  `xml:"bubble3D"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_SurfaceSer struct{
	cat string  `xml:"cat"`
	val CT_NumDataSource  `xml:"val"`
	extLst string  `xml:"extLst"`
	 EG_SerShared  `xml:""`
}
type CT_Grouping struct{
	val ST_Grouping  `xml:"val",attr`
}
type CT_ChartLines struct{
	spPr string  `xml:"spPr"`
}
type CT_LineChart struct{
	hiLowLines string  `xml:"hiLowLines"`
	upDownBars string  `xml:"upDownBars"`
	marker CT_Boolean  `xml:"marker"`
	smooth CT_Boolean  `xml:"smooth"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_LineChartShared  `xml:""`
}
type CT_Line3DChart struct{
	gapDepth CT_GapAmount  `xml:"gapDepth"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_LineChartShared  `xml:""`
}
type CT_StockChart struct{
	ser CT_LineSer  `xml:"ser"`
	dLbls string  `xml:"dLbls"`
	dropLines string  `xml:"dropLines"`
	hiLowLines string  `xml:"hiLowLines"`
	upDownBars string  `xml:"upDownBars"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
}
type CT_ScatterStyle struct{
	val ST_ScatterStyle  `xml:"val",attr`
}
type CT_ScatterChart struct{
	scatterStyle CT_ScatterStyle  `xml:"scatterStyle"`
	varyColors CT_Boolean  `xml:"varyColors"`
	ser []CT_ScatterSer  `xml:"ser"`
	dLbls string  `xml:"dLbls"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
}
type CT_RadarStyle struct{
	val string  `xml:"val",attr`
}
type CT_RadarChart struct{
	radarStyle string  `xml:"radarStyle"`
	varyColors CT_Boolean  `xml:"varyColors"`
	ser []string  `xml:"ser"`
	dLbls string  `xml:"dLbls"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
}
type CT_BarGrouping struct{
	val ST_BarGrouping  `xml:"val",attr`
}
type CT_BarDir struct{
	val ST_BarDir  `xml:"val",attr`
}
type CT_Shape struct{
	val ST_Shape  `xml:"val",attr`
}
type CT_BarChart struct{
	gapWidth CT_GapAmount  `xml:"gapWidth"`
	overlap CT_Overlap  `xml:"overlap"`
	serLines []string  `xml:"serLines"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_BarChartShared  `xml:""`
}
type CT_Bar3DChart struct{
	gapWidth CT_GapAmount  `xml:"gapWidth"`
	gapDepth CT_GapAmount  `xml:"gapDepth"`
	shape CT_Shape  `xml:"shape"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_BarChartShared  `xml:""`
}
type CT_AreaChart struct{
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_AreaChartShared  `xml:""`
}
type CT_Area3DChart struct{
	gapDepth CT_GapAmount  `xml:"gapDepth"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_AreaChartShared  `xml:""`
}
type CT_PieChart struct{
	firstSliceAng string  `xml:"firstSliceAng"`
	extLst string  `xml:"extLst"`
	 EG_PieChartShared  `xml:""`
}
type CT_Pie3DChart struct{
	extLst string  `xml:"extLst"`
	 EG_PieChartShared  `xml:""`
}
type CT_DoughnutChart struct{
	firstSliceAng string  `xml:"firstSliceAng"`
	holeSize CT_HoleSize  `xml:"holeSize"`
	extLst string  `xml:"extLst"`
	 EG_PieChartShared  `xml:""`
}
type CT_OfPieType struct{
	val ST_OfPieType  `xml:"val",attr`
}
type CT_OfPieChart struct{
	ofPieType CT_OfPieType  `xml:"ofPieType"`
	gapWidth CT_GapAmount  `xml:"gapWidth"`
	splitType CT_SplitType  `xml:"splitType"`
	splitPos CT_Double  `xml:"splitPos"`
	custSplit string  `xml:"custSplit"`
	secondPieSize string  `xml:"secondPieSize"`
	serLines []string  `xml:"serLines"`
	extLst string  `xml:"extLst"`
	 EG_PieChartShared  `xml:""`
}
type CT_BubbleChart struct{
	varyColors CT_Boolean  `xml:"varyColors"`
	ser []CT_BubbleSer  `xml:"ser"`
	dLbls string  `xml:"dLbls"`
	bubble3D CT_Boolean  `xml:"bubble3D"`
	bubbleScale CT_BubbleScale  `xml:"bubbleScale"`
	showNegBubbles CT_Boolean  `xml:"showNegBubbles"`
	sizeRepresents string  `xml:"sizeRepresents"`
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
}
type CT_BandFmt struct{
	idx string  `xml:"idx"`
	spPr string  `xml:"spPr"`
}
type CT_BandFmts struct{
	bandFmt []string  `xml:"bandFmt"`
}
type CT_SurfaceChart struct{
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_SurfaceChartShared  `xml:""`
}
type CT_Surface3DChart struct{
	axId string  `xml:"axId"`
	extLst string  `xml:"extLst"`
	 EG_SurfaceChartShared  `xml:""`
}
type CT_AxPos struct{
	val string  `xml:"val",attr`
}
type CT_Crosses struct{
	val string  `xml:"val",attr`
}
type CT_CrossBetween struct{
	val string  `xml:"val",attr`
}
type CT_TickMark struct{
	val ST_TickMark  `xml:"val",attr`
}
type CT_TickLblPos struct{
	val string  `xml:"val",attr`
}
type CT_Skip struct{
	val ST_Skip  `xml:"val",attr`
}
type CT_TimeUnit struct{
	val ST_TimeUnit  `xml:"val",attr`
}
type CT_AxisUnit struct{
	val string  `xml:"val",attr`
}
type CT_BuiltInUnit struct{
	val ST_BuiltInUnit  `xml:"val",attr`
}
type CT_PictureFormat struct{
	val ST_PictureFormat  `xml:"val",attr`
}
type CT_PictureStackUnit struct{
	val ST_PictureStackUnit  `xml:"val",attr`
}
type CT_PictureOptions struct{
	applyToFront CT_Boolean  `xml:"applyToFront"`
	applyToSides CT_Boolean  `xml:"applyToSides"`
	applyToEnd CT_Boolean  `xml:"applyToEnd"`
	pictureFormat CT_PictureFormat  `xml:"pictureFormat"`
	pictureStackUnit CT_PictureStackUnit  `xml:"pictureStackUnit"`
}
type CT_DispUnitsLbl struct{
	layout CT_Layout  `xml:"layout"`
	tx string  `xml:"tx"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
}
type CT_DispUnits struct{
	dispUnitsLbl string  `xml:"dispUnitsLbl"`
	extLst string  `xml:"extLst"`
	custUnit CT_Double  `xml:"custUnit"`
	builtInUnit CT_BuiltInUnit  `xml:"builtInUnit"`
}
type CT_Orientation struct{
	val ST_Orientation  `xml:"val",attr`
}
type CT_LogBase struct{
	val string  `xml:"val",attr`
}
type CT_Scaling struct{
	logBase string  `xml:"logBase"`
	orientation CT_Orientation  `xml:"orientation"`
	max CT_Double  `xml:"max"`
	min CT_Double  `xml:"min"`
	extLst string  `xml:"extLst"`
}
type CT_LblOffset struct{
	val string  `xml:"val",attr`
}
type CT_CatAx struct{
	auto CT_Boolean  `xml:"auto"`
	lblAlgn CT_LblAlgn  `xml:"lblAlgn"`
	lblOffset string  `xml:"lblOffset"`
	tickLblSkip CT_Skip  `xml:"tickLblSkip"`
	tickMarkSkip CT_Skip  `xml:"tickMarkSkip"`
	noMultiLvlLbl CT_Boolean  `xml:"noMultiLvlLbl"`
	extLst string  `xml:"extLst"`
	 EG_AxShared  `xml:""`
}
type CT_DateAx struct{
	auto CT_Boolean  `xml:"auto"`
	lblOffset string  `xml:"lblOffset"`
	baseTimeUnit CT_TimeUnit  `xml:"baseTimeUnit"`
	majorUnit string  `xml:"majorUnit"`
	majorTimeUnit CT_TimeUnit  `xml:"majorTimeUnit"`
	minorUnit string  `xml:"minorUnit"`
	minorTimeUnit CT_TimeUnit  `xml:"minorTimeUnit"`
	extLst string  `xml:"extLst"`
	 EG_AxShared  `xml:""`
}
type CT_SerAx struct{
	tickLblSkip CT_Skip  `xml:"tickLblSkip"`
	tickMarkSkip CT_Skip  `xml:"tickMarkSkip"`
	extLst string  `xml:"extLst"`
	 EG_AxShared  `xml:""`
}
type CT_ValAx struct{
	crossBetween string  `xml:"crossBetween"`
	majorUnit string  `xml:"majorUnit"`
	minorUnit string  `xml:"minorUnit"`
	dispUnits string  `xml:"dispUnits"`
	extLst string  `xml:"extLst"`
	 EG_AxShared  `xml:""`
}
type CT_PlotArea struct{
	layout CT_Layout  `xml:"layout"`
	dTable CT_DTable  `xml:"dTable"`
	spPr string  `xml:"spPr"`
	extLst string  `xml:"extLst"`
	areaChart CT_AreaChart  `xml:"areaChart"`
	area3DChart CT_Area3DChart  `xml:"area3DChart"`
	lineChart CT_LineChart  `xml:"lineChart"`
	line3DChart CT_Line3DChart  `xml:"line3DChart"`
	stockChart CT_StockChart  `xml:"stockChart"`
	radarChart string  `xml:"radarChart"`
	scatterChart CT_ScatterChart  `xml:"scatterChart"`
	pieChart CT_PieChart  `xml:"pieChart"`
	pie3DChart CT_Pie3DChart  `xml:"pie3DChart"`
	doughnutChart CT_DoughnutChart  `xml:"doughnutChart"`
	barChart CT_BarChart  `xml:"barChart"`
	bar3DChart CT_Bar3DChart  `xml:"bar3DChart"`
	ofPieChart CT_OfPieChart  `xml:"ofPieChart"`
	surfaceChart CT_SurfaceChart  `xml:"surfaceChart"`
	surface3DChart CT_Surface3DChart  `xml:"surface3DChart"`
	bubbleChart CT_BubbleChart  `xml:"bubbleChart"`
	valAx string  `xml:"valAx"`
	catAx string  `xml:"catAx"`
	dateAx string  `xml:"dateAx"`
	serAx string  `xml:"serAx"`
}
type CT_PivotFmt struct{
	idx string  `xml:"idx"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	marker CT_Marker  `xml:"marker"`
	dLbl CT_DLbl  `xml:"dLbl"`
	extLst string  `xml:"extLst"`
}
type CT_PivotFmts struct{
	pivotFmt []CT_PivotFmt  `xml:"pivotFmt"`
}
type CT_LegendPos struct{
	val string  `xml:"val",attr`
}
type CT_LegendEntry struct{
	idx string  `xml:"idx"`
	extLst string  `xml:"extLst"`
	 EG_LegendEntryData  `xml:""`
	delete CT_Boolean  `xml:"delete"`
}
type CT_Legend struct{
	legendPos string  `xml:"legendPos"`
	legendEntry []string  `xml:"legendEntry"`
	layout CT_Layout  `xml:"layout"`
	overlay CT_Boolean  `xml:"overlay"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	extLst string  `xml:"extLst"`
}
type CT_DispBlanksAs struct{
	val string  `xml:"val",attr`
}
type CT_Chart struct{
	title CT_Title  `xml:"title"`
	autoTitleDeleted CT_Boolean  `xml:"autoTitleDeleted"`
	pivotFmts string  `xml:"pivotFmts"`
	view3D CT_View3D  `xml:"view3D"`
	floor CT_Surface  `xml:"floor"`
	sideWall CT_Surface  `xml:"sideWall"`
	backWall CT_Surface  `xml:"backWall"`
	plotArea CT_PlotArea  `xml:"plotArea"`
	legend string  `xml:"legend"`
	plotVisOnly CT_Boolean  `xml:"plotVisOnly"`
	dispBlanksAs string  `xml:"dispBlanksAs"`
	showDLblsOverMax CT_Boolean  `xml:"showDLblsOverMax"`
	extLst string  `xml:"extLst"`
}
type CT_Style struct{
	val ST_Style  `xml:"val",attr`
}
type CT_PivotSource struct{
	name string  `xml:"name"`
	fmtId string  `xml:"fmtId"`
	extLst []string  `xml:"extLst"`
}
type CT_Protection struct{
	chartObject CT_Boolean  `xml:"chartObject"`
	data CT_Boolean  `xml:"data"`
	formatting CT_Boolean  `xml:"formatting"`
	selection CT_Boolean  `xml:"selection"`
	userInterface CT_Boolean  `xml:"userInterface"`
}
type CT_HeaderFooter struct{
	oddHeader string  `xml:"oddHeader"`
	oddFooter string  `xml:"oddFooter"`
	evenHeader string  `xml:"evenHeader"`
	evenFooter string  `xml:"evenFooter"`
	firstHeader string  `xml:"firstHeader"`
	firstFooter string  `xml:"firstFooter"`
	alignWithMargins string  `xml:"alignWithMargins",attr`
	differentOddEven string  `xml:"differentOddEven",attr`
	differentFirst string  `xml:"differentFirst",attr`
}
type CT_PageMargins struct{
	l string  `xml:"l",attr`
	r string  `xml:"r",attr`
	t string  `xml:"t",attr`
	b string  `xml:"b",attr`
	header string  `xml:"header",attr`
	footer string  `xml:"footer",attr`
}
type CT_ExternalData struct{
	autoUpdate CT_Boolean  `xml:"autoUpdate"`
}
type CT_PageSetup struct{
	paperSize string  `xml:"paperSize",attr`
	paperHeight string  `xml:"paperHeight",attr`
	paperWidth string  `xml:"paperWidth",attr`
	firstPageNumber string  `xml:"firstPageNumber",attr`
	orientation ST_PageSetupOrientation  `xml:"orientation",attr`
	blackAndWhite string  `xml:"blackAndWhite",attr`
	draft string  `xml:"draft",attr`
	useFirstPageNumber string  `xml:"useFirstPageNumber",attr`
	horizontalDpi string  `xml:"horizontalDpi",attr`
	verticalDpi string  `xml:"verticalDpi",attr`
	copies string  `xml:"copies",attr`
}
type CT_PrintSettings struct{
	headerFooter string  `xml:"headerFooter"`
	pageMargins string  `xml:"pageMargins"`
	pageSetup CT_PageSetup  `xml:"pageSetup"`
	legacyDrawingHF string  `xml:"legacyDrawingHF"`
}
type CT_ChartSpace struct{
	date1904 CT_Boolean  `xml:"date1904"`
	lang string  `xml:"lang"`
	roundedCorners CT_Boolean  `xml:"roundedCorners"`
	style CT_Style  `xml:"style"`
	clrMapOvr string  `xml:"clrMapOvr"`
	pivotSource CT_PivotSource  `xml:"pivotSource"`
	protection CT_Protection  `xml:"protection"`
	chart CT_Chart  `xml:"chart"`
	spPr string  `xml:"spPr"`
	txPr string  `xml:"txPr"`
	externalData string  `xml:"externalData"`
	printSettings string  `xml:"printSettings"`
	userShapes string  `xml:"userShapes"`
	extLst string  `xml:"extLst"`
}
