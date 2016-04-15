package ooxml
type WordSchema  struct {
	calcChain CT_CalcChain  `xml:"calcChain"`
	comments string  `xml:"comments"`
	MapInfo CT_MapInfo  `xml:"MapInfo"`
	connections string  `xml:"connections"`
	pivotCacheDefinition CT_PivotCacheDefinition  `xml:"pivotCacheDefinition"`
	pivotCacheRecords string  `xml:"pivotCacheRecords"`
	pivotTableDefinition CT_pivotTableDefinition  `xml:"pivotTableDefinition"`
	queryTable CT_QueryTable  `xml:"queryTable"`
	sst string  `xml:"sst"`
	headers string  `xml:"headers"`
	revisions string  `xml:"revisions"`
	users string  `xml:"users"`
	worksheet string  `xml:"worksheet"`
	chartsheet string  `xml:"chartsheet"`
	dialogsheet string  `xml:"dialogsheet"`
	metadata string  `xml:"metadata"`
	singleXmlCells string  `xml:"singleXmlCells"`
	styleSheet string  `xml:"styleSheet"`
	externalLink string  `xml:"externalLink"`
	table CT_Table  `xml:"table"`
	volTypes string  `xml:"volTypes"`
	workbook CT_Workbook  `xml:"workbook"`
}
type CT_AutoFilter struct{
	filterColumn []CT_FilterColumn  `xml:"filterColumn"`
	sortState CT_SortState  `xml:"sortState"`
	extLst string  `xml:"extLst"`
	ref ST_Ref  `xml:"ref",attr`
}
type CT_FilterColumn struct{
	filters string  `xml:"filters"`
	top10 CT_Top10  `xml:"top10"`
	customFilters string  `xml:"customFilters"`
	dynamicFilter CT_DynamicFilter  `xml:"dynamicFilter"`
	colorFilter CT_ColorFilter  `xml:"colorFilter"`
	iconFilter CT_IconFilter  `xml:"iconFilter"`
	extLst string  `xml:"extLst"`
	colId string  `xml:"colId",attr`
	hiddenButton string  `xml:"hiddenButton",attr`
	showButton string  `xml:"showButton",attr`
}
type CT_Filters struct{
	filter []CT_Filter  `xml:"filter"`
	dateGroupItem []CT_DateGroupItem  `xml:"dateGroupItem"`
	blank string  `xml:"blank",attr`
	calendarType string  `xml:"calendarType",attr`
}
type CT_Filter struct{
	val string  `xml:"val",attr`
}
type CT_CustomFilters struct{
	customFilter string  `xml:"customFilter"`
	and string  `xml:"and",attr`
}
type CT_CustomFilter struct{
	operator ST_FilterOperator  `xml:"operator",attr`
	val string  `xml:"val",attr`
}
type CT_Top10 struct{
	top string  `xml:"top",attr`
	percent string  `xml:"percent",attr`
	val string  `xml:"val",attr`
	filterVal string  `xml:"filterVal",attr`
}
type CT_ColorFilter struct{
	dxfId string  `xml:"dxfId",attr`
	cellColor string  `xml:"cellColor",attr`
}
type CT_IconFilter struct{
	iconSet ST_IconSetType  `xml:"iconSet",attr`
	iconId string  `xml:"iconId",attr`
}
type CT_DynamicFilter struct{
	type ST_DynamicFilterType  `xml:"type",attr`
	val string  `xml:"val",attr`
	valIso string  `xml:"valIso",attr`
	maxVal string  `xml:"maxVal",attr`
	maxValIso string  `xml:"maxValIso",attr`
}
type CT_SortState struct{
	sortCondition string  `xml:"sortCondition"`
	extLst string  `xml:"extLst"`
	columnSort string  `xml:"columnSort",attr`
	caseSensitive string  `xml:"caseSensitive",attr`
	sortMethod string  `xml:"sortMethod",attr`
	ref ST_Ref  `xml:"ref",attr`
}
type CT_SortCondition struct{
	descending string  `xml:"descending",attr`
	sortBy ST_SortBy  `xml:"sortBy",attr`
	ref ST_Ref  `xml:"ref",attr`
	customList string  `xml:"customList",attr`
	dxfId string  `xml:"dxfId",attr`
	iconSet ST_IconSetType  `xml:"iconSet",attr`
	iconId string  `xml:"iconId",attr`
}
type CT_DateGroupItem struct{
	year string  `xml:"year",attr`
	month string  `xml:"month",attr`
	day string  `xml:"day",attr`
	hour string  `xml:"hour",attr`
	minute string  `xml:"minute",attr`
	second string  `xml:"second",attr`
	dateTimeGrouping ST_DateTimeGrouping  `xml:"dateTimeGrouping",attr`
}
type CT_XStringElement struct{
	v string  `xml:"v",attr`
}
type CT_Extension struct{
	uri string  `xml:"uri",attr`
}
type CT_ObjectAnchor struct{
	moveWithCells string  `xml:"moveWithCells",attr`
	sizeWithCells string  `xml:"sizeWithCells",attr`
}
type CT_ExtensionList struct{
	 EG_ExtensionList  `xml:""`
}
type CT_CalcChain struct{
	c []CT_CalcCell  `xml:"c"`
	extLst string  `xml:"extLst"`
}
type CT_CalcCell struct{
	r ST_CellRef  `xml:"r",attr`
	ref ST_CellRef  `xml:"ref",attr`
	i string  `xml:"i",attr`
	s string  `xml:"s",attr`
	l string  `xml:"l",attr`
	t string  `xml:"t",attr`
	a string  `xml:"a",attr`
}
type CT_Comments struct{
	authors string  `xml:"authors"`
	commentList string  `xml:"commentList"`
	extLst string  `xml:"extLst"`
}
type CT_Authors struct{
	author []string  `xml:"author"`
}
type CT_CommentList struct{
	comment []CT_Comment  `xml:"comment"`
}
type CT_Comment struct{
	text string  `xml:"text"`
	commentPr CT_CommentPr  `xml:"commentPr"`
	ref ST_Ref  `xml:"ref",attr`
	authorId string  `xml:"authorId",attr`
	guid string  `xml:"guid",attr`
	shapeId string  `xml:"shapeId",attr`
}
type CT_CommentPr struct{
	anchor CT_ObjectAnchor  `xml:"anchor"`
	locked string  `xml:"locked",attr`
	defaultSize string  `xml:"defaultSize",attr`
	print string  `xml:"print",attr`
	disabled string  `xml:"disabled",attr`
	autoFill string  `xml:"autoFill",attr`
	autoLine string  `xml:"autoLine",attr`
	altText string  `xml:"altText",attr`
	textHAlign string  `xml:"textHAlign",attr`
	textVAlign string  `xml:"textVAlign",attr`
	lockText string  `xml:"lockText",attr`
	justLastX string  `xml:"justLastX",attr`
	autoScale string  `xml:"autoScale",attr`
}
type CT_MapInfo struct{
	Schema []CT_Schema  `xml:"Schema"`
	Map []CT_Map  `xml:"Map"`
	SelectionNamespaces string  `xml:"SelectionNamespaces",attr`
}
type CT_Schema struct{
	ID string  `xml:"ID",attr`
	SchemaRef string  `xml:"SchemaRef",attr`
	Namespace string  `xml:"Namespace",attr`
	SchemaLanguage string  `xml:"SchemaLanguage",attr`
}
type CT_Map struct{
	DataBinding string  `xml:"DataBinding"`
	ID string  `xml:"ID",attr`
	Name string  `xml:"Name",attr`
	RootElement string  `xml:"RootElement",attr`
	SchemaID string  `xml:"SchemaID",attr`
	ShowImportExportValidationErrors string  `xml:"ShowImportExportValidationErrors",attr`
	AutoFit string  `xml:"AutoFit",attr`
	Append string  `xml:"Append",attr`
	PreserveSortAFLayout string  `xml:"PreserveSortAFLayout",attr`
	PreserveFormat string  `xml:"PreserveFormat",attr`
}
type CT_DataBinding struct{
	DataBindingName string  `xml:"DataBindingName",attr`
	FileBinding string  `xml:"FileBinding",attr`
	ConnectionID string  `xml:"ConnectionID",attr`
	FileBindingName string  `xml:"FileBindingName",attr`
	DataBindingLoadMode string  `xml:"DataBindingLoadMode",attr`
}
type CT_Connections struct{
	connection []CT_Connection  `xml:"connection"`
}
type CT_Connection struct{
	dbPr CT_DbPr  `xml:"dbPr"`
	olapPr CT_OlapPr  `xml:"olapPr"`
	webPr CT_WebPr  `xml:"webPr"`
	textPr string  `xml:"textPr"`
	parameters string  `xml:"parameters"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	sourceFile string  `xml:"sourceFile",attr`
	odcFile string  `xml:"odcFile",attr`
	keepAlive string  `xml:"keepAlive",attr`
	interval string  `xml:"interval",attr`
	name string  `xml:"name",attr`
	description string  `xml:"description",attr`
	type string  `xml:"type",attr`
	reconnectionMethod string  `xml:"reconnectionMethod",attr`
	refreshedVersion string  `xml:"refreshedVersion",attr`
	minRefreshableVersion string  `xml:"minRefreshableVersion",attr`
	savePassword string  `xml:"savePassword",attr`
	new string  `xml:"new",attr`
	deleted string  `xml:"deleted",attr`
	onlyUseConnectionFile string  `xml:"onlyUseConnectionFile",attr`
	background string  `xml:"background",attr`
	refreshOnLoad string  `xml:"refreshOnLoad",attr`
	saveData string  `xml:"saveData",attr`
	credentials string  `xml:"credentials",attr`
	singleSignOnId string  `xml:"singleSignOnId",attr`
}
type CT_DbPr struct{
	connection string  `xml:"connection",attr`
	command string  `xml:"command",attr`
	serverCommand string  `xml:"serverCommand",attr`
	commandType string  `xml:"commandType",attr`
}
type CT_OlapPr struct{
	local string  `xml:"local",attr`
	localConnection string  `xml:"localConnection",attr`
	localRefresh string  `xml:"localRefresh",attr`
	sendLocale string  `xml:"sendLocale",attr`
	rowDrillCount string  `xml:"rowDrillCount",attr`
	serverFill string  `xml:"serverFill",attr`
	serverNumberFormat string  `xml:"serverNumberFormat",attr`
	serverFont string  `xml:"serverFont",attr`
	serverFontColor string  `xml:"serverFontColor",attr`
}
type CT_WebPr struct{
	tables string  `xml:"tables"`
	xml string  `xml:"xml",attr`
	sourceData string  `xml:"sourceData",attr`
	parsePre string  `xml:"parsePre",attr`
	consecutive string  `xml:"consecutive",attr`
	firstRow string  `xml:"firstRow",attr`
	xl97 string  `xml:"xl97",attr`
	textDates string  `xml:"textDates",attr`
	xl2000 string  `xml:"xl2000",attr`
	url string  `xml:"url",attr`
	post string  `xml:"post",attr`
	htmlTables string  `xml:"htmlTables",attr`
	htmlFormat ST_HtmlFmt  `xml:"htmlFormat",attr`
	editPage string  `xml:"editPage",attr`
}
type CT_Parameters struct{
	parameter []CT_Parameter  `xml:"parameter"`
	count string  `xml:"count",attr`
}
type CT_Parameter struct{
	name string  `xml:"name",attr`
	sqlType string  `xml:"sqlType",attr`
	parameterType ST_ParameterType  `xml:"parameterType",attr`
	refreshOnChange string  `xml:"refreshOnChange",attr`
	prompt string  `xml:"prompt",attr`
	boolean string  `xml:"boolean",attr`
	double string  `xml:"double",attr`
	integer string  `xml:"integer",attr`
	string string  `xml:"string",attr`
	cell string  `xml:"cell",attr`
}
type CT_Tables struct{
	m string  `xml:"m"`
	s CT_XStringElement  `xml:"s"`
	x string  `xml:"x"`
	count string  `xml:"count",attr`
}
type CT_TableMissing struct{
}
type CT_TextPr struct{
	textFields string  `xml:"textFields"`
	prompt string  `xml:"prompt",attr`
	fileType ST_FileType  `xml:"fileType",attr`
	codePage string  `xml:"codePage",attr`
	characterSet string  `xml:"characterSet",attr`
	firstRow string  `xml:"firstRow",attr`
	sourceFile string  `xml:"sourceFile",attr`
	delimited string  `xml:"delimited",attr`
	decimal string  `xml:"decimal",attr`
	thousands string  `xml:"thousands",attr`
	tab string  `xml:"tab",attr`
	space string  `xml:"space",attr`
	comma string  `xml:"comma",attr`
	semicolon string  `xml:"semicolon",attr`
	consecutive string  `xml:"consecutive",attr`
	qualifier ST_Qualifier  `xml:"qualifier",attr`
	delimiter string  `xml:"delimiter",attr`
}
type CT_TextFields struct{
	textField []string  `xml:"textField"`
	count string  `xml:"count",attr`
}
type CT_TextField struct{
	type string  `xml:"type",attr`
	position string  `xml:"position",attr`
}
type CT_PivotCacheDefinition struct{
	cacheSource CT_CacheSource  `xml:"cacheSource"`
	cacheFields string  `xml:"cacheFields"`
	cacheHierarchies string  `xml:"cacheHierarchies"`
	kpis string  `xml:"kpis"`
	tupleCache CT_TupleCache  `xml:"tupleCache"`
	calculatedItems string  `xml:"calculatedItems"`
	calculatedMembers string  `xml:"calculatedMembers"`
	dimensions string  `xml:"dimensions"`
	measureGroups string  `xml:"measureGroups"`
	maps string  `xml:"maps"`
	extLst string  `xml:"extLst"`
	invalid string  `xml:"invalid",attr`
	saveData string  `xml:"saveData",attr`
	refreshOnLoad string  `xml:"refreshOnLoad",attr`
	optimizeMemory string  `xml:"optimizeMemory",attr`
	enableRefresh string  `xml:"enableRefresh",attr`
	refreshedBy string  `xml:"refreshedBy",attr`
	refreshedDate string  `xml:"refreshedDate",attr`
	refreshedDateIso string  `xml:"refreshedDateIso",attr`
	backgroundQuery string  `xml:"backgroundQuery",attr`
	missingItemsLimit string  `xml:"missingItemsLimit",attr`
	createdVersion string  `xml:"createdVersion",attr`
	refreshedVersion string  `xml:"refreshedVersion",attr`
	minRefreshableVersion string  `xml:"minRefreshableVersion",attr`
	recordCount string  `xml:"recordCount",attr`
	upgradeOnRefresh string  `xml:"upgradeOnRefresh",attr`
	tupleCache string  `xml:"tupleCache",attr`
	supportSubquery string  `xml:"supportSubquery",attr`
	supportAdvancedDrill string  `xml:"supportAdvancedDrill",attr`
}
type CT_CacheFields struct{
	cacheField []string  `xml:"cacheField"`
	count string  `xml:"count",attr`
}
type CT_CacheField struct{
	sharedItems string  `xml:"sharedItems"`
	fieldGroup string  `xml:"fieldGroup"`
	mpMap []CT_X  `xml:"mpMap"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	caption string  `xml:"caption",attr`
	propertyName string  `xml:"propertyName",attr`
	serverField string  `xml:"serverField",attr`
	uniqueList string  `xml:"uniqueList",attr`
	numFmtId string  `xml:"numFmtId",attr`
	formula string  `xml:"formula",attr`
	sqlType string  `xml:"sqlType",attr`
	hierarchy string  `xml:"hierarchy",attr`
	level string  `xml:"level",attr`
	databaseField string  `xml:"databaseField",attr`
	mappingCount string  `xml:"mappingCount",attr`
	memberPropertyField string  `xml:"memberPropertyField",attr`
}
type CT_CacheSource struct{
	worksheetSource string  `xml:"worksheetSource"`
	consolidation string  `xml:"consolidation"`
	extLst string  `xml:"extLst"`
	type ST_SourceType  `xml:"type",attr`
	connectionId string  `xml:"connectionId",attr`
}
type CT_WorksheetSource struct{
	ref ST_Ref  `xml:"ref",attr`
	name string  `xml:"name",attr`
	sheet string  `xml:"sheet",attr`
}
type CT_Consolidation struct{
	pages string  `xml:"pages"`
	rangeSets string  `xml:"rangeSets"`
	autoPage string  `xml:"autoPage",attr`
}
type CT_Pages struct{
	page CT_PCDSCPage  `xml:"page"`
	count string  `xml:"count",attr`
}
type CT_PCDSCPage struct{
	pageItem []CT_PageItem  `xml:"pageItem"`
	count string  `xml:"count",attr`
}
type CT_PageItem struct{
	name string  `xml:"name",attr`
}
type CT_RangeSets struct{
	rangeSet []CT_RangeSet  `xml:"rangeSet"`
	count string  `xml:"count",attr`
}
type CT_RangeSet struct{
	i1 string  `xml:"i1",attr`
	i2 string  `xml:"i2",attr`
	i3 string  `xml:"i3",attr`
	i4 string  `xml:"i4",attr`
	ref ST_Ref  `xml:"ref",attr`
	name string  `xml:"name",attr`
	sheet string  `xml:"sheet",attr`
}
type CT_SharedItems struct{
	m string  `xml:"m"`
	n CT_Number  `xml:"n"`
	b CT_Boolean  `xml:"b"`
	e CT_Error  `xml:"e"`
	s CT_String  `xml:"s"`
	d CT_DateTime  `xml:"d"`
	containsSemiMixedTypes string  `xml:"containsSemiMixedTypes",attr`
	containsNonDate string  `xml:"containsNonDate",attr`
	containsDate string  `xml:"containsDate",attr`
	containsString string  `xml:"containsString",attr`
	containsBlank string  `xml:"containsBlank",attr`
	containsMixedTypes string  `xml:"containsMixedTypes",attr`
	containsNumber string  `xml:"containsNumber",attr`
	containsInteger string  `xml:"containsInteger",attr`
	minValue string  `xml:"minValue",attr`
	maxValue string  `xml:"maxValue",attr`
	minDate string  `xml:"minDate",attr`
	maxDate string  `xml:"maxDate",attr`
	count string  `xml:"count",attr`
	longText string  `xml:"longText",attr`
}
type CT_Missing struct{
	tpls []string  `xml:"tpls"`
	x []CT_X  `xml:"x"`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
	in string  `xml:"in",attr`
	bc string  `xml:"bc",attr`
	fc string  `xml:"fc",attr`
	i string  `xml:"i",attr`
	un string  `xml:"un",attr`
	st string  `xml:"st",attr`
	b string  `xml:"b",attr`
}
type CT_Number struct{
	tpls []string  `xml:"tpls"`
	x []CT_X  `xml:"x"`
	v string  `xml:"v",attr`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
	in string  `xml:"in",attr`
	bc string  `xml:"bc",attr`
	fc string  `xml:"fc",attr`
	i string  `xml:"i",attr`
	un string  `xml:"un",attr`
	st string  `xml:"st",attr`
	b string  `xml:"b",attr`
}
type CT_Boolean struct{
	x []CT_X  `xml:"x"`
	v string  `xml:"v",attr`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
}
type CT_Error struct{
	tpls string  `xml:"tpls"`
	x []CT_X  `xml:"x"`
	v string  `xml:"v",attr`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
	in string  `xml:"in",attr`
	bc string  `xml:"bc",attr`
	fc string  `xml:"fc",attr`
	i string  `xml:"i",attr`
	un string  `xml:"un",attr`
	st string  `xml:"st",attr`
	b string  `xml:"b",attr`
}
type CT_String struct{
	tpls []string  `xml:"tpls"`
	x []CT_X  `xml:"x"`
	v string  `xml:"v",attr`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
	in string  `xml:"in",attr`
	bc string  `xml:"bc",attr`
	fc string  `xml:"fc",attr`
	i string  `xml:"i",attr`
	un string  `xml:"un",attr`
	st string  `xml:"st",attr`
	b string  `xml:"b",attr`
}
type CT_DateTime struct{
	x []CT_X  `xml:"x"`
	v string  `xml:"v",attr`
	u string  `xml:"u",attr`
	f string  `xml:"f",attr`
	c string  `xml:"c",attr`
	cp string  `xml:"cp",attr`
}
type CT_FieldGroup struct{
	rangePr CT_RangePr  `xml:"rangePr"`
	discretePr string  `xml:"discretePr"`
	groupItems string  `xml:"groupItems"`
	par string  `xml:"par",attr`
	base string  `xml:"base",attr`
}
type CT_RangePr struct{
	autoStart string  `xml:"autoStart",attr`
	autoEnd string  `xml:"autoEnd",attr`
	groupBy ST_GroupBy  `xml:"groupBy",attr`
	startNum string  `xml:"startNum",attr`
	endNum string  `xml:"endNum",attr`
	startDate string  `xml:"startDate",attr`
	endDate string  `xml:"endDate",attr`
	groupInterval string  `xml:"groupInterval",attr`
}
type CT_DiscretePr struct{
	x []string  `xml:"x"`
	count string  `xml:"count",attr`
}
type CT_GroupItems struct{
	m string  `xml:"m"`
	n CT_Number  `xml:"n"`
	b CT_Boolean  `xml:"b"`
	e CT_Error  `xml:"e"`
	s CT_String  `xml:"s"`
	d CT_DateTime  `xml:"d"`
	count string  `xml:"count",attr`
}
type CT_PivotCacheRecords struct{
	r []string  `xml:"r"`
	extLst string  `xml:"extLst"`
	count string  `xml:"count",attr`
}
type CT_Record struct{
	m string  `xml:"m"`
	n CT_Number  `xml:"n"`
	b CT_Boolean  `xml:"b"`
	e CT_Error  `xml:"e"`
	s CT_String  `xml:"s"`
	d CT_DateTime  `xml:"d"`
	x string  `xml:"x"`
}
type CT_PCDKPIs struct{
	kpi []CT_PCDKPI  `xml:"kpi"`
	count string  `xml:"count",attr`
}
type CT_PCDKPI struct{
	uniqueName string  `xml:"uniqueName",attr`
	caption string  `xml:"caption",attr`
	displayFolder string  `xml:"displayFolder",attr`
	measureGroup string  `xml:"measureGroup",attr`
	parent string  `xml:"parent",attr`
	value string  `xml:"value",attr`
	goal string  `xml:"goal",attr`
	status string  `xml:"status",attr`
	trend string  `xml:"trend",attr`
	weight string  `xml:"weight",attr`
	time string  `xml:"time",attr`
}
type CT_CacheHierarchies struct{
	cacheHierarchy []CT_CacheHierarchy  `xml:"cacheHierarchy"`
	count string  `xml:"count",attr`
}
type CT_CacheHierarchy struct{
	fieldsUsage string  `xml:"fieldsUsage"`
	groupLevels string  `xml:"groupLevels"`
	extLst string  `xml:"extLst"`
	uniqueName string  `xml:"uniqueName",attr`
	caption string  `xml:"caption",attr`
	measure string  `xml:"measure",attr`
	set string  `xml:"set",attr`
	parentSet string  `xml:"parentSet",attr`
	iconSet string  `xml:"iconSet",attr`
	attribute string  `xml:"attribute",attr`
	time string  `xml:"time",attr`
	keyAttribute string  `xml:"keyAttribute",attr`
	defaultMemberUniqueName string  `xml:"defaultMemberUniqueName",attr`
	allUniqueName string  `xml:"allUniqueName",attr`
	allCaption string  `xml:"allCaption",attr`
	dimensionUniqueName string  `xml:"dimensionUniqueName",attr`
	displayFolder string  `xml:"displayFolder",attr`
	measureGroup string  `xml:"measureGroup",attr`
	measures string  `xml:"measures",attr`
	count string  `xml:"count",attr`
	oneField string  `xml:"oneField",attr`
	memberValueDatatype string  `xml:"memberValueDatatype",attr`
	unbalanced string  `xml:"unbalanced",attr`
	unbalancedGroup string  `xml:"unbalancedGroup",attr`
	hidden string  `xml:"hidden",attr`
}
type CT_FieldsUsage struct{
	fieldUsage []string  `xml:"fieldUsage"`
	count string  `xml:"count",attr`
}
type CT_FieldUsage struct{
	x string  `xml:"x",attr`
}
type CT_GroupLevels struct{
	groupLevel []CT_GroupLevel  `xml:"groupLevel"`
	count string  `xml:"count",attr`
}
type CT_GroupLevel struct{
	groups string  `xml:"groups"`
	extLst string  `xml:"extLst"`
	uniqueName string  `xml:"uniqueName",attr`
	caption string  `xml:"caption",attr`
	user string  `xml:"user",attr`
	customRollUp string  `xml:"customRollUp",attr`
}
type CT_Groups struct{
	group []CT_LevelGroup  `xml:"group"`
	count string  `xml:"count",attr`
}
type CT_LevelGroup struct{
	groupMembers string  `xml:"groupMembers"`
	name string  `xml:"name",attr`
	uniqueName string  `xml:"uniqueName",attr`
	caption string  `xml:"caption",attr`
	uniqueParent string  `xml:"uniqueParent",attr`
	id string  `xml:"id",attr`
}
type CT_GroupMembers struct{
	groupMember []CT_GroupMember  `xml:"groupMember"`
	count string  `xml:"count",attr`
}
type CT_GroupMember struct{
	uniqueName string  `xml:"uniqueName",attr`
	group string  `xml:"group",attr`
}
type CT_TupleCache struct{
	entries string  `xml:"entries"`
	sets string  `xml:"sets"`
	queryCache CT_QueryCache  `xml:"queryCache"`
	serverFormats string  `xml:"serverFormats"`
	extLst string  `xml:"extLst"`
}
type CT_ServerFormat struct{
	culture string  `xml:"culture",attr`
	format string  `xml:"format",attr`
}
type CT_ServerFormats struct{
	serverFormat []CT_ServerFormat  `xml:"serverFormat"`
	count string  `xml:"count",attr`
}
type CT_PCDSDTCEntries struct{
	m string  `xml:"m"`
	n CT_Number  `xml:"n"`
	e CT_Error  `xml:"e"`
	s CT_String  `xml:"s"`
	count string  `xml:"count",attr`
}
type CT_Tuples struct{
	tpl []CT_Tuple  `xml:"tpl"`
	c string  `xml:"c",attr`
}
type CT_Tuple struct{
	fld string  `xml:"fld",attr`
	hier string  `xml:"hier",attr`
	item string  `xml:"item",attr`
}
type CT_Sets struct{
	set []CT_Set  `xml:"set"`
	count string  `xml:"count",attr`
}
type CT_Set struct{
	tpls []string  `xml:"tpls"`
	sortByTuple string  `xml:"sortByTuple"`
	count string  `xml:"count",attr`
	maxRank string  `xml:"maxRank",attr`
	setDefinition string  `xml:"setDefinition",attr`
	sortType ST_SortType  `xml:"sortType",attr`
	queryFailed string  `xml:"queryFailed",attr`
}
type CT_QueryCache struct{
	query []CT_Query  `xml:"query"`
	count string  `xml:"count",attr`
}
type CT_Query struct{
	tpls string  `xml:"tpls"`
	mdx string  `xml:"mdx",attr`
}
type CT_CalculatedItems struct{
	calculatedItem []string  `xml:"calculatedItem"`
	count string  `xml:"count",attr`
}
type CT_CalculatedItem struct{
	pivotArea CT_PivotArea  `xml:"pivotArea"`
	extLst string  `xml:"extLst"`
	field string  `xml:"field",attr`
	formula string  `xml:"formula",attr`
}
type CT_CalculatedMembers struct{
	calculatedMember []string  `xml:"calculatedMember"`
	count string  `xml:"count",attr`
}
type CT_CalculatedMember struct{
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	mdx string  `xml:"mdx",attr`
	memberName string  `xml:"memberName",attr`
	hierarchy string  `xml:"hierarchy",attr`
	parent string  `xml:"parent",attr`
	solveOrder string  `xml:"solveOrder",attr`
	set string  `xml:"set",attr`
}
type CT_pivotTableDefinition struct{
	location CT_Location  `xml:"location"`
	pivotFields string  `xml:"pivotFields"`
	rowFields string  `xml:"rowFields"`
	rowItems string  `xml:"rowItems"`
	colFields string  `xml:"colFields"`
	colItems string  `xml:"colItems"`
	pageFields string  `xml:"pageFields"`
	dataFields string  `xml:"dataFields"`
	formats string  `xml:"formats"`
	conditionalFormats string  `xml:"conditionalFormats"`
	chartFormats string  `xml:"chartFormats"`
	pivotHierarchies string  `xml:"pivotHierarchies"`
	pivotTableStyleInfo CT_PivotTableStyle  `xml:"pivotTableStyleInfo"`
	filters string  `xml:"filters"`
	rowHierarchiesUsage string  `xml:"rowHierarchiesUsage"`
	colHierarchiesUsage string  `xml:"colHierarchiesUsage"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	cacheId string  `xml:"cacheId",attr`
	dataOnRows string  `xml:"dataOnRows",attr`
	dataPosition string  `xml:"dataPosition",attr`
	dataCaption string  `xml:"dataCaption",attr`
	grandTotalCaption string  `xml:"grandTotalCaption",attr`
	errorCaption string  `xml:"errorCaption",attr`
	showError string  `xml:"showError",attr`
	missingCaption string  `xml:"missingCaption",attr`
	showMissing string  `xml:"showMissing",attr`
	pageStyle string  `xml:"pageStyle",attr`
	pivotTableStyle string  `xml:"pivotTableStyle",attr`
	vacatedStyle string  `xml:"vacatedStyle",attr`
	tag string  `xml:"tag",attr`
	updatedVersion string  `xml:"updatedVersion",attr`
	minRefreshableVersion string  `xml:"minRefreshableVersion",attr`
	asteriskTotals string  `xml:"asteriskTotals",attr`
	showItems string  `xml:"showItems",attr`
	editData string  `xml:"editData",attr`
	disableFieldList string  `xml:"disableFieldList",attr`
	showCalcMbrs string  `xml:"showCalcMbrs",attr`
	visualTotals string  `xml:"visualTotals",attr`
	showMultipleLabel string  `xml:"showMultipleLabel",attr`
	showDataDropDown string  `xml:"showDataDropDown",attr`
	showDrill string  `xml:"showDrill",attr`
	printDrill string  `xml:"printDrill",attr`
	showMemberPropertyTips string  `xml:"showMemberPropertyTips",attr`
	showDataTips string  `xml:"showDataTips",attr`
	enableWizard string  `xml:"enableWizard",attr`
	enableDrill string  `xml:"enableDrill",attr`
	enableFieldProperties string  `xml:"enableFieldProperties",attr`
	preserveFormatting string  `xml:"preserveFormatting",attr`
	useAutoFormatting string  `xml:"useAutoFormatting",attr`
	pageWrap string  `xml:"pageWrap",attr`
	pageOverThenDown string  `xml:"pageOverThenDown",attr`
	subtotalHiddenItems string  `xml:"subtotalHiddenItems",attr`
	rowGrandTotals string  `xml:"rowGrandTotals",attr`
	colGrandTotals string  `xml:"colGrandTotals",attr`
	fieldPrintTitles string  `xml:"fieldPrintTitles",attr`
	itemPrintTitles string  `xml:"itemPrintTitles",attr`
	mergeItem string  `xml:"mergeItem",attr`
	showDropZones string  `xml:"showDropZones",attr`
	createdVersion string  `xml:"createdVersion",attr`
	indent string  `xml:"indent",attr`
	showEmptyRow string  `xml:"showEmptyRow",attr`
	showEmptyCol string  `xml:"showEmptyCol",attr`
	showHeaders string  `xml:"showHeaders",attr`
	compact string  `xml:"compact",attr`
	outline string  `xml:"outline",attr`
	outlineData string  `xml:"outlineData",attr`
	compactData string  `xml:"compactData",attr`
	published string  `xml:"published",attr`
	gridDropZones string  `xml:"gridDropZones",attr`
	immersive string  `xml:"immersive",attr`
	multipleFieldFilters string  `xml:"multipleFieldFilters",attr`
	chartFormat string  `xml:"chartFormat",attr`
	rowHeaderCaption string  `xml:"rowHeaderCaption",attr`
	colHeaderCaption string  `xml:"colHeaderCaption",attr`
	fieldListSortAscending string  `xml:"fieldListSortAscending",attr`
	mdxSubqueries string  `xml:"mdxSubqueries",attr`
	customListSort string  `xml:"customListSort",attr`
}
type CT_Location struct{
	ref ST_Ref  `xml:"ref",attr`
	firstHeaderRow string  `xml:"firstHeaderRow",attr`
	firstDataRow string  `xml:"firstDataRow",attr`
	firstDataCol string  `xml:"firstDataCol",attr`
	rowPageCount string  `xml:"rowPageCount",attr`
	colPageCount string  `xml:"colPageCount",attr`
}
type CT_PivotFields struct{
	pivotField []string  `xml:"pivotField"`
	count string  `xml:"count",attr`
}
type CT_PivotField struct{
	items string  `xml:"items"`
	autoSortScope CT_AutoSortScope  `xml:"autoSortScope"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	axis string  `xml:"axis",attr`
	dataField string  `xml:"dataField",attr`
	subtotalCaption string  `xml:"subtotalCaption",attr`
	showDropDowns string  `xml:"showDropDowns",attr`
	hiddenLevel string  `xml:"hiddenLevel",attr`
	uniqueMemberProperty string  `xml:"uniqueMemberProperty",attr`
	compact string  `xml:"compact",attr`
	allDrilled string  `xml:"allDrilled",attr`
	numFmtId string  `xml:"numFmtId",attr`
	outline string  `xml:"outline",attr`
	subtotalTop string  `xml:"subtotalTop",attr`
	dragToRow string  `xml:"dragToRow",attr`
	dragToCol string  `xml:"dragToCol",attr`
	multipleItemSelectionAllowed string  `xml:"multipleItemSelectionAllowed",attr`
	dragToPage string  `xml:"dragToPage",attr`
	dragToData string  `xml:"dragToData",attr`
	dragOff string  `xml:"dragOff",attr`
	showAll string  `xml:"showAll",attr`
	insertBlankRow string  `xml:"insertBlankRow",attr`
	serverField string  `xml:"serverField",attr`
	insertPageBreak string  `xml:"insertPageBreak",attr`
	autoShow string  `xml:"autoShow",attr`
	topAutoShow string  `xml:"topAutoShow",attr`
	hideNewItems string  `xml:"hideNewItems",attr`
	measureFilter string  `xml:"measureFilter",attr`
	includeNewItemsInFilter string  `xml:"includeNewItemsInFilter",attr`
	itemPageCount string  `xml:"itemPageCount",attr`
	sortType string  `xml:"sortType",attr`
	dataSourceSort string  `xml:"dataSourceSort",attr`
	nonAutoSortDefault string  `xml:"nonAutoSortDefault",attr`
	rankBy string  `xml:"rankBy",attr`
	defaultSubtotal string  `xml:"defaultSubtotal",attr`
	sumSubtotal string  `xml:"sumSubtotal",attr`
	countASubtotal string  `xml:"countASubtotal",attr`
	avgSubtotal string  `xml:"avgSubtotal",attr`
	maxSubtotal string  `xml:"maxSubtotal",attr`
	minSubtotal string  `xml:"minSubtotal",attr`
	productSubtotal string  `xml:"productSubtotal",attr`
	countSubtotal string  `xml:"countSubtotal",attr`
	stdDevSubtotal string  `xml:"stdDevSubtotal",attr`
	stdDevPSubtotal string  `xml:"stdDevPSubtotal",attr`
	varSubtotal string  `xml:"varSubtotal",attr`
	varPSubtotal string  `xml:"varPSubtotal",attr`
	showPropCell string  `xml:"showPropCell",attr`
	showPropTip string  `xml:"showPropTip",attr`
	showPropAsCaption string  `xml:"showPropAsCaption",attr`
	defaultAttributeDrillState string  `xml:"defaultAttributeDrillState",attr`
}
type CT_AutoSortScope struct{
	pivotArea CT_PivotArea  `xml:"pivotArea"`
}
type CT_Items struct{
	item []CT_Item  `xml:"item"`
	count string  `xml:"count",attr`
}
type CT_Item struct{
	n string  `xml:"n",attr`
	t ST_ItemType  `xml:"t",attr`
	h string  `xml:"h",attr`
	s string  `xml:"s",attr`
	sd string  `xml:"sd",attr`
	f string  `xml:"f",attr`
	m string  `xml:"m",attr`
	c string  `xml:"c",attr`
	x string  `xml:"x",attr`
	d string  `xml:"d",attr`
	e string  `xml:"e",attr`
}
type CT_PageFields struct{
	pageField []string  `xml:"pageField"`
	count string  `xml:"count",attr`
}
type CT_PageField struct{
	extLst string  `xml:"extLst"`
	fld string  `xml:"fld",attr`
	item string  `xml:"item",attr`
	hier string  `xml:"hier",attr`
	name string  `xml:"name",attr`
	cap string  `xml:"cap",attr`
}
type CT_DataFields struct{
	dataField []string  `xml:"dataField"`
	count string  `xml:"count",attr`
}
type CT_DataField struct{
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	fld string  `xml:"fld",attr`
	subtotal string  `xml:"subtotal",attr`
	showDataAs string  `xml:"showDataAs",attr`
	baseField string  `xml:"baseField",attr`
	baseItem string  `xml:"baseItem",attr`
	numFmtId string  `xml:"numFmtId",attr`
}
type CT_rowItems struct{
	i []CT_I  `xml:"i"`
	count string  `xml:"count",attr`
}
type CT_colItems struct{
	i []CT_I  `xml:"i"`
	count string  `xml:"count",attr`
}
type CT_I struct{
	x []CT_X  `xml:"x"`
	t ST_ItemType  `xml:"t",attr`
	r string  `xml:"r",attr`
	i string  `xml:"i",attr`
}
type CT_X struct{
	v string  `xml:"v",attr`
}
type CT_RowFields struct{
	field []string  `xml:"field"`
	count string  `xml:"count",attr`
}
type CT_ColFields struct{
	field []string  `xml:"field"`
	count string  `xml:"count",attr`
}
type CT_Field struct{
	x string  `xml:"x",attr`
}
type CT_Formats struct{
	format []CT_Format  `xml:"format"`
	count string  `xml:"count",attr`
}
type CT_Format struct{
	pivotArea CT_PivotArea  `xml:"pivotArea"`
	extLst string  `xml:"extLst"`
	action ST_FormatAction  `xml:"action",attr`
	dxfId string  `xml:"dxfId",attr`
}
type CT_ConditionalFormats struct{
	conditionalFormat []string  `xml:"conditionalFormat"`
	count string  `xml:"count",attr`
}
type CT_ConditionalFormat struct{
	pivotAreas string  `xml:"pivotAreas"`
	extLst string  `xml:"extLst"`
	scope ST_Scope  `xml:"scope",attr`
	type ST_Type  `xml:"type",attr`
	priority string  `xml:"priority",attr`
}
type CT_PivotAreas struct{
	pivotArea []CT_PivotArea  `xml:"pivotArea"`
	count string  `xml:"count",attr`
}
type CT_ChartFormats struct{
	chartFormat []CT_ChartFormat  `xml:"chartFormat"`
	count string  `xml:"count",attr`
}
type CT_ChartFormat struct{
	pivotArea CT_PivotArea  `xml:"pivotArea"`
	chart string  `xml:"chart",attr`
	format string  `xml:"format",attr`
	series string  `xml:"series",attr`
}
type CT_PivotHierarchies struct{
	pivotHierarchy []CT_PivotHierarchy  `xml:"pivotHierarchy"`
	count string  `xml:"count",attr`
}
type CT_PivotHierarchy struct{
	mps string  `xml:"mps"`
	members []string  `xml:"members"`
	extLst string  `xml:"extLst"`
	outline string  `xml:"outline",attr`
	multipleItemSelectionAllowed string  `xml:"multipleItemSelectionAllowed",attr`
	subtotalTop string  `xml:"subtotalTop",attr`
	showInFieldList string  `xml:"showInFieldList",attr`
	dragToRow string  `xml:"dragToRow",attr`
	dragToCol string  `xml:"dragToCol",attr`
	dragToPage string  `xml:"dragToPage",attr`
	dragToData string  `xml:"dragToData",attr`
	dragOff string  `xml:"dragOff",attr`
	includeNewItemsInFilter string  `xml:"includeNewItemsInFilter",attr`
	caption string  `xml:"caption",attr`
}
type CT_RowHierarchiesUsage struct{
	rowHierarchyUsage []string  `xml:"rowHierarchyUsage"`
	count string  `xml:"count",attr`
}
type CT_ColHierarchiesUsage struct{
	colHierarchyUsage []string  `xml:"colHierarchyUsage"`
	count string  `xml:"count",attr`
}
type CT_HierarchyUsage struct{
	hierarchyUsage string  `xml:"hierarchyUsage",attr`
}
type CT_MemberProperties struct{
	mp []CT_MemberProperty  `xml:"mp"`
	count string  `xml:"count",attr`
}
type CT_MemberProperty struct{
	name string  `xml:"name",attr`
	showCell string  `xml:"showCell",attr`
	showTip string  `xml:"showTip",attr`
	showAsCaption string  `xml:"showAsCaption",attr`
	nameLen string  `xml:"nameLen",attr`
	pPos string  `xml:"pPos",attr`
	pLen string  `xml:"pLen",attr`
	level string  `xml:"level",attr`
	field string  `xml:"field",attr`
}
type CT_Members struct{
	member []CT_Member  `xml:"member"`
	count string  `xml:"count",attr`
	level string  `xml:"level",attr`
}
type CT_Member struct{
	name string  `xml:"name",attr`
}
type CT_Dimensions struct{
	dimension []string  `xml:"dimension"`
	count string  `xml:"count",attr`
}
type CT_PivotDimension struct{
	measure string  `xml:"measure",attr`
	name string  `xml:"name",attr`
	uniqueName string  `xml:"uniqueName",attr`
	caption string  `xml:"caption",attr`
}
type CT_MeasureGroups struct{
	measureGroup []string  `xml:"measureGroup"`
	count string  `xml:"count",attr`
}
type CT_MeasureDimensionMaps struct{
	map []string  `xml:"map"`
	count string  `xml:"count",attr`
}
type CT_MeasureGroup struct{
	name string  `xml:"name",attr`
	caption string  `xml:"caption",attr`
}
type CT_MeasureDimensionMap struct{
	measureGroup string  `xml:"measureGroup",attr`
	dimension string  `xml:"dimension",attr`
}
type CT_PivotTableStyle struct{
	name string  `xml:"name",attr`
	showRowHeaders string  `xml:"showRowHeaders",attr`
	showColHeaders string  `xml:"showColHeaders",attr`
	showRowStripes string  `xml:"showRowStripes",attr`
	showColStripes string  `xml:"showColStripes",attr`
	showLastColumn string  `xml:"showLastColumn",attr`
}
type CT_PivotFilters struct{
	filter []CT_PivotFilter  `xml:"filter"`
	count string  `xml:"count",attr`
}
type CT_PivotFilter struct{
	autoFilter CT_AutoFilter  `xml:"autoFilter"`
	extLst string  `xml:"extLst"`
	fld string  `xml:"fld",attr`
	mpFld string  `xml:"mpFld",attr`
	type ST_PivotFilterType  `xml:"type",attr`
	evalOrder string  `xml:"evalOrder",attr`
	id string  `xml:"id",attr`
	iMeasureHier string  `xml:"iMeasureHier",attr`
	iMeasureFld string  `xml:"iMeasureFld",attr`
	name string  `xml:"name",attr`
	description string  `xml:"description",attr`
	stringValue1 string  `xml:"stringValue1",attr`
	stringValue2 string  `xml:"stringValue2",attr`
}
type CT_PivotArea struct{
	references string  `xml:"references"`
	extLst string  `xml:"extLst"`
	field string  `xml:"field",attr`
	type ST_PivotAreaType  `xml:"type",attr`
	dataOnly string  `xml:"dataOnly",attr`
	labelOnly string  `xml:"labelOnly",attr`
	grandRow string  `xml:"grandRow",attr`
	grandCol string  `xml:"grandCol",attr`
	cacheIndex string  `xml:"cacheIndex",attr`
	outline string  `xml:"outline",attr`
	offset ST_Ref  `xml:"offset",attr`
	collapsedLevelsAreSubtotals string  `xml:"collapsedLevelsAreSubtotals",attr`
	axis string  `xml:"axis",attr`
	fieldPosition string  `xml:"fieldPosition",attr`
}
type CT_PivotAreaReferences struct{
	reference []CT_PivotAreaReference  `xml:"reference"`
	count string  `xml:"count",attr`
}
type CT_PivotAreaReference struct{
	x []string  `xml:"x"`
	extLst string  `xml:"extLst"`
	field string  `xml:"field",attr`
	count string  `xml:"count",attr`
	selected string  `xml:"selected",attr`
	byPosition string  `xml:"byPosition",attr`
	relative string  `xml:"relative",attr`
	defaultSubtotal string  `xml:"defaultSubtotal",attr`
	sumSubtotal string  `xml:"sumSubtotal",attr`
	countASubtotal string  `xml:"countASubtotal",attr`
	avgSubtotal string  `xml:"avgSubtotal",attr`
	maxSubtotal string  `xml:"maxSubtotal",attr`
	minSubtotal string  `xml:"minSubtotal",attr`
	productSubtotal string  `xml:"productSubtotal",attr`
	countSubtotal string  `xml:"countSubtotal",attr`
	stdDevSubtotal string  `xml:"stdDevSubtotal",attr`
	stdDevPSubtotal string  `xml:"stdDevPSubtotal",attr`
	varSubtotal string  `xml:"varSubtotal",attr`
	varPSubtotal string  `xml:"varPSubtotal",attr`
}
type CT_Index struct{
	v string  `xml:"v",attr`
}
type CT_QueryTable struct{
	queryTableRefresh string  `xml:"queryTableRefresh"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	headers string  `xml:"headers",attr`
	rowNumbers string  `xml:"rowNumbers",attr`
	disableRefresh string  `xml:"disableRefresh",attr`
	backgroundRefresh string  `xml:"backgroundRefresh",attr`
	firstBackgroundRefresh string  `xml:"firstBackgroundRefresh",attr`
	refreshOnLoad string  `xml:"refreshOnLoad",attr`
	growShrinkType ST_GrowShrinkType  `xml:"growShrinkType",attr`
	fillFormulas string  `xml:"fillFormulas",attr`
	removeDataOnSave string  `xml:"removeDataOnSave",attr`
	disableEdit string  `xml:"disableEdit",attr`
	preserveFormatting string  `xml:"preserveFormatting",attr`
	adjustColumnWidth string  `xml:"adjustColumnWidth",attr`
	intermediate string  `xml:"intermediate",attr`
	connectionId string  `xml:"connectionId",attr`
}
type CT_QueryTableRefresh struct{
	queryTableFields string  `xml:"queryTableFields"`
	queryTableDeletedFields string  `xml:"queryTableDeletedFields"`
	sortState CT_SortState  `xml:"sortState"`
	extLst string  `xml:"extLst"`
	preserveSortFilterLayout string  `xml:"preserveSortFilterLayout",attr`
	fieldIdWrapped string  `xml:"fieldIdWrapped",attr`
	headersInLastRefresh string  `xml:"headersInLastRefresh",attr`
	minimumVersion string  `xml:"minimumVersion",attr`
	nextId string  `xml:"nextId",attr`
	unboundColumnsLeft string  `xml:"unboundColumnsLeft",attr`
	unboundColumnsRight string  `xml:"unboundColumnsRight",attr`
}
type CT_QueryTableDeletedFields struct{
	deletedField []string  `xml:"deletedField"`
	count string  `xml:"count",attr`
}
type CT_DeletedField struct{
	name string  `xml:"name",attr`
}
type CT_QueryTableFields struct{
	queryTableField []string  `xml:"queryTableField"`
	count string  `xml:"count",attr`
}
type CT_QueryTableField struct{
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	name string  `xml:"name",attr`
	dataBound string  `xml:"dataBound",attr`
	rowNumbers string  `xml:"rowNumbers",attr`
	fillFormulas string  `xml:"fillFormulas",attr`
	clipped string  `xml:"clipped",attr`
	tableColumnId string  `xml:"tableColumnId",attr`
}
type CT_Sst struct{
	si []string  `xml:"si"`
	extLst string  `xml:"extLst"`
	count string  `xml:"count",attr`
	uniqueCount string  `xml:"uniqueCount",attr`
}
type CT_PhoneticRun struct{
	t string  `xml:"t"`
	sb string  `xml:"sb",attr`
	eb string  `xml:"eb",attr`
}
type CT_RElt struct{
	rPr CT_RPrElt  `xml:"rPr"`
	t string  `xml:"t"`
}
type CT_RPrElt struct{
	rFont CT_FontName  `xml:"rFont"`
	charset CT_IntProperty  `xml:"charset"`
	family CT_IntProperty  `xml:"family"`
	b CT_BooleanProperty  `xml:"b"`
	i CT_BooleanProperty  `xml:"i"`
	strike CT_BooleanProperty  `xml:"strike"`
	outline CT_BooleanProperty  `xml:"outline"`
	shadow CT_BooleanProperty  `xml:"shadow"`
	condense CT_BooleanProperty  `xml:"condense"`
	extend CT_BooleanProperty  `xml:"extend"`
	color CT_Color  `xml:"color"`
	sz CT_FontSize  `xml:"sz"`
	u string  `xml:"u"`
	vertAlign CT_VerticalAlignFontProperty  `xml:"vertAlign"`
	scheme CT_FontScheme  `xml:"scheme"`
}
type CT_Rst struct{
	t string  `xml:"t"`
	r []CT_RElt  `xml:"r"`
	rPh []CT_PhoneticRun  `xml:"rPh"`
	phoneticPr CT_PhoneticPr  `xml:"phoneticPr"`
}
type CT_PhoneticPr struct{
	fontId string  `xml:"fontId",attr`
	type ST_PhoneticType  `xml:"type",attr`
	alignment ST_PhoneticAlignment  `xml:"alignment",attr`
}
type CT_RevisionHeaders struct{
	header []string  `xml:"header"`
	guid string  `xml:"guid",attr`
	lastGuid string  `xml:"lastGuid",attr`
	shared string  `xml:"shared",attr`
	diskRevisions string  `xml:"diskRevisions",attr`
	history string  `xml:"history",attr`
	trackRevisions string  `xml:"trackRevisions",attr`
	exclusive string  `xml:"exclusive",attr`
	revisionId string  `xml:"revisionId",attr`
	version string  `xml:"version",attr`
	keepChangeHistory string  `xml:"keepChangeHistory",attr`
	protected string  `xml:"protected",attr`
	preserveHistory string  `xml:"preserveHistory",attr`
}
type CT_Revisions struct{
	rrc []string  `xml:"rrc"`
	rm []string  `xml:"rm"`
	rcv []string  `xml:"rcv"`
	rsnm []string  `xml:"rsnm"`
	ris []string  `xml:"ris"`
	rcc []string  `xml:"rcc"`
	rfmt []string  `xml:"rfmt"`
	raf []string  `xml:"raf"`
	rdn []string  `xml:"rdn"`
	rcmt []string  `xml:"rcmt"`
	rqt []string  `xml:"rqt"`
	rcft []string  `xml:"rcft"`
}
type CT_RevisionHeader struct{
	sheetIdMap string  `xml:"sheetIdMap"`
	reviewedList string  `xml:"reviewedList"`
	extLst string  `xml:"extLst"`
	guid string  `xml:"guid",attr`
	dateTime string  `xml:"dateTime",attr`
	maxSheetId string  `xml:"maxSheetId",attr`
	userName string  `xml:"userName",attr`
	minRId string  `xml:"minRId",attr`
	maxRId string  `xml:"maxRId",attr`
}
type CT_SheetIdMap struct{
	sheetId []string  `xml:"sheetId"`
	count string  `xml:"count",attr`
}
type CT_SheetId struct{
	val string  `xml:"val",attr`
}
type CT_ReviewedRevisions struct{
	reviewed []string  `xml:"reviewed"`
	count string  `xml:"count",attr`
}
type CT_Reviewed struct{
	rId string  `xml:"rId",attr`
}
type CT_UndoInfo struct{
	index string  `xml:"index",attr`
	exp string  `xml:"exp",attr`
	ref3D string  `xml:"ref3D",attr`
	array string  `xml:"array",attr`
	v string  `xml:"v",attr`
	nf string  `xml:"nf",attr`
	cs string  `xml:"cs",attr`
	dr ST_RefA  `xml:"dr",attr`
	dn string  `xml:"dn",attr`
	r ST_CellRef  `xml:"r",attr`
	sId string  `xml:"sId",attr`
}
type CT_RevisionRowColumn struct{
	undo []string  `xml:"undo"`
	rcc []string  `xml:"rcc"`
	rfmt []string  `xml:"rfmt"`
	sId string  `xml:"sId",attr`
	eol string  `xml:"eol",attr`
	ref ST_Ref  `xml:"ref",attr`
	action ST_rwColActionType  `xml:"action",attr`
	edge string  `xml:"edge",attr`
}
type CT_RevisionMove struct{
	undo []string  `xml:"undo"`
	rcc []string  `xml:"rcc"`
	rfmt []string  `xml:"rfmt"`
	sheetId string  `xml:"sheetId",attr`
	source ST_Ref  `xml:"source",attr`
	destination ST_Ref  `xml:"destination",attr`
	sourceSheetId string  `xml:"sourceSheetId",attr`
}
type CT_RevisionCustomView struct{
	guid string  `xml:"guid",attr`
	action string  `xml:"action",attr`
}
type CT_RevisionSheetRename struct{
	extLst string  `xml:"extLst"`
	sheetId string  `xml:"sheetId",attr`
	oldName string  `xml:"oldName",attr`
	newName string  `xml:"newName",attr`
}
type CT_RevisionInsertSheet struct{
	sheetId string  `xml:"sheetId",attr`
	name string  `xml:"name",attr`
	sheetPosition string  `xml:"sheetPosition",attr`
}
type CT_RevisionCellChange struct{
	oc CT_Cell  `xml:"oc"`
	nc CT_Cell  `xml:"nc"`
	odxf string  `xml:"odxf"`
	ndxf string  `xml:"ndxf"`
	extLst string  `xml:"extLst"`
	sId string  `xml:"sId",attr`
	odxf string  `xml:"odxf",attr`
	xfDxf string  `xml:"xfDxf",attr`
	s string  `xml:"s",attr`
	dxf string  `xml:"dxf",attr`
	numFmtId string  `xml:"numFmtId",attr`
	quotePrefix string  `xml:"quotePrefix",attr`
	oldQuotePrefix string  `xml:"oldQuotePrefix",attr`
	ph string  `xml:"ph",attr`
	oldPh string  `xml:"oldPh",attr`
	endOfListFormulaUpdate string  `xml:"endOfListFormulaUpdate",attr`
}
type CT_RevisionFormatting struct{
	dxf string  `xml:"dxf"`
	extLst string  `xml:"extLst"`
	sheetId string  `xml:"sheetId",attr`
	xfDxf string  `xml:"xfDxf",attr`
	s string  `xml:"s",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
	start string  `xml:"start",attr`
	length string  `xml:"length",attr`
}
type CT_RevisionAutoFormatting struct{
	sheetId string  `xml:"sheetId",attr`
	ref ST_Ref  `xml:"ref",attr`
}
type CT_RevisionComment struct{
	sheetId string  `xml:"sheetId",attr`
	cell ST_CellRef  `xml:"cell",attr`
	guid string  `xml:"guid",attr`
	action string  `xml:"action",attr`
	alwaysShow string  `xml:"alwaysShow",attr`
	old string  `xml:"old",attr`
	hiddenRow string  `xml:"hiddenRow",attr`
	hiddenColumn string  `xml:"hiddenColumn",attr`
	author string  `xml:"author",attr`
	oldLength string  `xml:"oldLength",attr`
	newLength string  `xml:"newLength",attr`
}
type CT_RevisionDefinedName struct{
	formula ST_Formula  `xml:"formula"`
	oldFormula ST_Formula  `xml:"oldFormula"`
	extLst string  `xml:"extLst"`
	localSheetId string  `xml:"localSheetId",attr`
	customView string  `xml:"customView",attr`
	name string  `xml:"name",attr`
	function string  `xml:"function",attr`
	oldFunction string  `xml:"oldFunction",attr`
	functionGroupId string  `xml:"functionGroupId",attr`
	oldFunctionGroupId string  `xml:"oldFunctionGroupId",attr`
	shortcutKey string  `xml:"shortcutKey",attr`
	oldShortcutKey string  `xml:"oldShortcutKey",attr`
	hidden string  `xml:"hidden",attr`
	oldHidden string  `xml:"oldHidden",attr`
	customMenu string  `xml:"customMenu",attr`
	oldCustomMenu string  `xml:"oldCustomMenu",attr`
	description string  `xml:"description",attr`
	oldDescription string  `xml:"oldDescription",attr`
	help string  `xml:"help",attr`
	oldHelp string  `xml:"oldHelp",attr`
	statusBar string  `xml:"statusBar",attr`
	oldStatusBar string  `xml:"oldStatusBar",attr`
	comment string  `xml:"comment",attr`
	oldComment string  `xml:"oldComment",attr`
}
type CT_RevisionConflict struct{
	sheetId string  `xml:"sheetId",attr`
}
type CT_RevisionQueryTableField struct{
	sheetId string  `xml:"sheetId",attr`
	ref ST_Ref  `xml:"ref",attr`
	fieldId string  `xml:"fieldId",attr`
}
type CT_Users struct{
	userInfo string  `xml:"userInfo"`
	count string  `xml:"count",attr`
}
type CT_SharedUser struct{
	extLst string  `xml:"extLst"`
	guid string  `xml:"guid",attr`
	name string  `xml:"name",attr`
	id string  `xml:"id",attr`
	dateTime string  `xml:"dateTime",attr`
}
type CT_Macrosheet struct{
	sheetPr CT_SheetPr  `xml:"sheetPr"`
	dimension string  `xml:"dimension"`
	sheetViews string  `xml:"sheetViews"`
	sheetFormatPr CT_SheetFormatPr  `xml:"sheetFormatPr"`
	cols []string  `xml:"cols"`
	sheetData CT_SheetData  `xml:"sheetData"`
	sheetProtection CT_SheetProtection  `xml:"sheetProtection"`
	autoFilter CT_AutoFilter  `xml:"autoFilter"`
	sortState CT_SortState  `xml:"sortState"`
	dataConsolidate string  `xml:"dataConsolidate"`
	customSheetViews string  `xml:"customSheetViews"`
	phoneticPr CT_PhoneticPr  `xml:"phoneticPr"`
	conditionalFormatting []string  `xml:"conditionalFormatting"`
	printOptions string  `xml:"printOptions"`
	pageMargins string  `xml:"pageMargins"`
	pageSetup CT_PageSetup  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	rowBreaks CT_PageBreak  `xml:"rowBreaks"`
	colBreaks CT_PageBreak  `xml:"colBreaks"`
	customProperties string  `xml:"customProperties"`
	drawing CT_Drawing  `xml:"drawing"`
	legacyDrawing CT_LegacyDrawing  `xml:"legacyDrawing"`
	legacyDrawingHF CT_LegacyDrawing  `xml:"legacyDrawingHF"`
	drawingHF CT_DrawingHF  `xml:"drawingHF"`
	picture string  `xml:"picture"`
	oleObjects string  `xml:"oleObjects"`
	extLst string  `xml:"extLst"`
}
type CT_Dialogsheet struct{
	sheetPr CT_SheetPr  `xml:"sheetPr"`
	sheetViews string  `xml:"sheetViews"`
	sheetFormatPr CT_SheetFormatPr  `xml:"sheetFormatPr"`
	sheetProtection CT_SheetProtection  `xml:"sheetProtection"`
	customSheetViews string  `xml:"customSheetViews"`
	printOptions string  `xml:"printOptions"`
	pageMargins string  `xml:"pageMargins"`
	pageSetup CT_PageSetup  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	drawing CT_Drawing  `xml:"drawing"`
	legacyDrawing CT_LegacyDrawing  `xml:"legacyDrawing"`
	legacyDrawingHF CT_LegacyDrawing  `xml:"legacyDrawingHF"`
	drawingHF CT_DrawingHF  `xml:"drawingHF"`
	oleObjects string  `xml:"oleObjects"`
	controls string  `xml:"controls"`
	extLst string  `xml:"extLst"`
}
type CT_Worksheet struct{
	sheetPr CT_SheetPr  `xml:"sheetPr"`
	dimension string  `xml:"dimension"`
	sheetViews string  `xml:"sheetViews"`
	sheetFormatPr CT_SheetFormatPr  `xml:"sheetFormatPr"`
	cols []string  `xml:"cols"`
	sheetData CT_SheetData  `xml:"sheetData"`
	sheetCalcPr CT_SheetCalcPr  `xml:"sheetCalcPr"`
	sheetProtection CT_SheetProtection  `xml:"sheetProtection"`
	protectedRanges string  `xml:"protectedRanges"`
	scenarios string  `xml:"scenarios"`
	autoFilter CT_AutoFilter  `xml:"autoFilter"`
	sortState CT_SortState  `xml:"sortState"`
	dataConsolidate string  `xml:"dataConsolidate"`
	customSheetViews string  `xml:"customSheetViews"`
	mergeCells string  `xml:"mergeCells"`
	phoneticPr CT_PhoneticPr  `xml:"phoneticPr"`
	conditionalFormatting []string  `xml:"conditionalFormatting"`
	dataValidations string  `xml:"dataValidations"`
	hyperlinks string  `xml:"hyperlinks"`
	printOptions string  `xml:"printOptions"`
	pageMargins string  `xml:"pageMargins"`
	pageSetup CT_PageSetup  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	rowBreaks CT_PageBreak  `xml:"rowBreaks"`
	colBreaks CT_PageBreak  `xml:"colBreaks"`
	customProperties string  `xml:"customProperties"`
	cellWatches string  `xml:"cellWatches"`
	ignoredErrors string  `xml:"ignoredErrors"`
	smartTags string  `xml:"smartTags"`
	drawing CT_Drawing  `xml:"drawing"`
	legacyDrawing CT_LegacyDrawing  `xml:"legacyDrawing"`
	legacyDrawingHF CT_LegacyDrawing  `xml:"legacyDrawingHF"`
	drawingHF CT_DrawingHF  `xml:"drawingHF"`
	picture string  `xml:"picture"`
	oleObjects string  `xml:"oleObjects"`
	controls string  `xml:"controls"`
	webPublishItems string  `xml:"webPublishItems"`
	tableParts string  `xml:"tableParts"`
	extLst string  `xml:"extLst"`
}
type CT_SheetData struct{
	row []CT_Row  `xml:"row"`
}
type CT_SheetCalcPr struct{
	fullCalcOnLoad string  `xml:"fullCalcOnLoad",attr`
}
type CT_SheetFormatPr struct{
	baseColWidth string  `xml:"baseColWidth",attr`
	defaultColWidth string  `xml:"defaultColWidth",attr`
	defaultRowHeight string  `xml:"defaultRowHeight",attr`
	customHeight string  `xml:"customHeight",attr`
	zeroHeight string  `xml:"zeroHeight",attr`
	thickTop string  `xml:"thickTop",attr`
	thickBottom string  `xml:"thickBottom",attr`
	outlineLevelRow string  `xml:"outlineLevelRow",attr`
	outlineLevelCol string  `xml:"outlineLevelCol",attr`
}
type CT_Cols struct{
	col []CT_Col  `xml:"col"`
}
type CT_Col struct{
	min string  `xml:"min",attr`
	max string  `xml:"max",attr`
	width string  `xml:"width",attr`
	style string  `xml:"style",attr`
	hidden string  `xml:"hidden",attr`
	bestFit string  `xml:"bestFit",attr`
	customWidth string  `xml:"customWidth",attr`
	phonetic string  `xml:"phonetic",attr`
	outlineLevel string  `xml:"outlineLevel",attr`
	collapsed string  `xml:"collapsed",attr`
}
type CT_Row struct{
	c []CT_Cell  `xml:"c"`
	extLst string  `xml:"extLst"`
	r string  `xml:"r",attr`
	spans string  `xml:"spans",attr`
	s string  `xml:"s",attr`
	customFormat string  `xml:"customFormat",attr`
	ht string  `xml:"ht",attr`
	hidden string  `xml:"hidden",attr`
	customHeight string  `xml:"customHeight",attr`
	outlineLevel string  `xml:"outlineLevel",attr`
	collapsed string  `xml:"collapsed",attr`
	thickTop string  `xml:"thickTop",attr`
	thickBot string  `xml:"thickBot",attr`
	ph string  `xml:"ph",attr`
}
type CT_Cell struct{
	f CT_CellFormula  `xml:"f"`
	v string  `xml:"v"`
	is string  `xml:"is"`
	extLst string  `xml:"extLst"`
	r ST_CellRef  `xml:"r",attr`
	s string  `xml:"s",attr`
	t ST_CellType  `xml:"t",attr`
	cm string  `xml:"cm",attr`
	vm string  `xml:"vm",attr`
	ph string  `xml:"ph",attr`
}
type CT_SheetPr struct{
	tabColor CT_Color  `xml:"tabColor"`
	outlinePr CT_OutlinePr  `xml:"outlinePr"`
	pageSetUpPr CT_PageSetUpPr  `xml:"pageSetUpPr"`
	syncHorizontal string  `xml:"syncHorizontal",attr`
	syncVertical string  `xml:"syncVertical",attr`
	syncRef ST_Ref  `xml:"syncRef",attr`
	transitionEvaluation string  `xml:"transitionEvaluation",attr`
	transitionEntry string  `xml:"transitionEntry",attr`
	published string  `xml:"published",attr`
	codeName string  `xml:"codeName",attr`
	filterMode string  `xml:"filterMode",attr`
	enableFormatConditionsCalculation string  `xml:"enableFormatConditionsCalculation",attr`
}
type CT_SheetDimension struct{
	ref ST_Ref  `xml:"ref",attr`
}
type CT_SheetViews struct{
	sheetView []CT_SheetView  `xml:"sheetView"`
	extLst string  `xml:"extLst"`
}
type CT_SheetView struct{
	pane CT_Pane  `xml:"pane"`
	selection CT_Selection  `xml:"selection"`
	pivotSelection CT_PivotSelection  `xml:"pivotSelection"`
	extLst string  `xml:"extLst"`
	windowProtection string  `xml:"windowProtection",attr`
	showFormulas string  `xml:"showFormulas",attr`
	showGridLines string  `xml:"showGridLines",attr`
	showRowColHeaders string  `xml:"showRowColHeaders",attr`
	showZeros string  `xml:"showZeros",attr`
	rightToLeft string  `xml:"rightToLeft",attr`
	tabSelected string  `xml:"tabSelected",attr`
	showRuler string  `xml:"showRuler",attr`
	showOutlineSymbols string  `xml:"showOutlineSymbols",attr`
	defaultGridColor string  `xml:"defaultGridColor",attr`
	showWhiteSpace string  `xml:"showWhiteSpace",attr`
	view ST_SheetViewType  `xml:"view",attr`
	topLeftCell ST_CellRef  `xml:"topLeftCell",attr`
	colorId string  `xml:"colorId",attr`
	zoomScale string  `xml:"zoomScale",attr`
	zoomScaleNormal string  `xml:"zoomScaleNormal",attr`
	zoomScaleSheetLayoutView string  `xml:"zoomScaleSheetLayoutView",attr`
	zoomScalePageLayoutView string  `xml:"zoomScalePageLayoutView",attr`
	workbookViewId string  `xml:"workbookViewId",attr`
}
type CT_Pane struct{
	xSplit string  `xml:"xSplit",attr`
	ySplit string  `xml:"ySplit",attr`
	topLeftCell ST_CellRef  `xml:"topLeftCell",attr`
	activePane ST_Pane  `xml:"activePane",attr`
	state ST_PaneState  `xml:"state",attr`
}
type CT_PivotSelection struct{
	pivotArea CT_PivotArea  `xml:"pivotArea"`
	pane ST_Pane  `xml:"pane",attr`
	showHeader string  `xml:"showHeader",attr`
	label string  `xml:"label",attr`
	data string  `xml:"data",attr`
	extendable string  `xml:"extendable",attr`
	count string  `xml:"count",attr`
	axis string  `xml:"axis",attr`
	dimension string  `xml:"dimension",attr`
	start string  `xml:"start",attr`
	min string  `xml:"min",attr`
	max string  `xml:"max",attr`
	activeRow string  `xml:"activeRow",attr`
	activeCol string  `xml:"activeCol",attr`
	previousRow string  `xml:"previousRow",attr`
	previousCol string  `xml:"previousCol",attr`
	click string  `xml:"click",attr`
}
type CT_Selection struct{
	pane ST_Pane  `xml:"pane",attr`
	activeCell ST_CellRef  `xml:"activeCell",attr`
	activeCellId string  `xml:"activeCellId",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
}
type CT_PageBreak struct{
	brk []CT_Break  `xml:"brk"`
	count string  `xml:"count",attr`
	manualBreakCount string  `xml:"manualBreakCount",attr`
}
type CT_Break struct{
	id string  `xml:"id",attr`
	min string  `xml:"min",attr`
	max string  `xml:"max",attr`
	man string  `xml:"man",attr`
	pt string  `xml:"pt",attr`
}
type CT_OutlinePr struct{
	applyStyles string  `xml:"applyStyles",attr`
	summaryBelow string  `xml:"summaryBelow",attr`
	summaryRight string  `xml:"summaryRight",attr`
	showOutlineSymbols string  `xml:"showOutlineSymbols",attr`
}
type CT_PageSetUpPr struct{
	autoPageBreaks string  `xml:"autoPageBreaks",attr`
	fitToPage string  `xml:"fitToPage",attr`
}
type CT_DataConsolidate struct{
	dataRefs string  `xml:"dataRefs"`
	function string  `xml:"function",attr`
	startLabels string  `xml:"startLabels",attr`
	leftLabels string  `xml:"leftLabels",attr`
	topLabels string  `xml:"topLabels",attr`
	link string  `xml:"link",attr`
}
type CT_DataRefs struct{
	dataRef []CT_DataRef  `xml:"dataRef"`
	count string  `xml:"count",attr`
}
type CT_DataRef struct{
	ref ST_Ref  `xml:"ref",attr`
	name string  `xml:"name",attr`
	sheet string  `xml:"sheet",attr`
}
type CT_MergeCells struct{
	mergeCell []CT_MergeCell  `xml:"mergeCell"`
	count string  `xml:"count",attr`
}
type CT_MergeCell struct{
	ref ST_Ref  `xml:"ref",attr`
}
type CT_SmartTags struct{
	cellSmartTags []string  `xml:"cellSmartTags"`
}
type CT_CellSmartTags struct{
	cellSmartTag []CT_CellSmartTag  `xml:"cellSmartTag"`
	r ST_CellRef  `xml:"r",attr`
}
type CT_CellSmartTag struct{
	cellSmartTagPr []CT_CellSmartTagPr  `xml:"cellSmartTagPr"`
	type string  `xml:"type",attr`
	deleted string  `xml:"deleted",attr`
	xmlBased string  `xml:"xmlBased",attr`
}
type CT_CellSmartTagPr struct{
	key string  `xml:"key",attr`
	val string  `xml:"val",attr`
}
type CT_Drawing struct{
}
type CT_LegacyDrawing struct{
}
type CT_DrawingHF struct{
	lho string  `xml:"lho",attr`
	lhe string  `xml:"lhe",attr`
	lhf string  `xml:"lhf",attr`
	cho string  `xml:"cho",attr`
	che string  `xml:"che",attr`
	chf string  `xml:"chf",attr`
	rho string  `xml:"rho",attr`
	rhe string  `xml:"rhe",attr`
	rhf string  `xml:"rhf",attr`
	lfo string  `xml:"lfo",attr`
	lfe string  `xml:"lfe",attr`
	lff string  `xml:"lff",attr`
	cfo string  `xml:"cfo",attr`
	cfe string  `xml:"cfe",attr`
	cff string  `xml:"cff",attr`
	rfo string  `xml:"rfo",attr`
	rfe string  `xml:"rfe",attr`
	rff string  `xml:"rff",attr`
}
type CT_CustomSheetViews struct{
	customSheetView []string  `xml:"customSheetView"`
}
type CT_CustomSheetView struct{
	pane CT_Pane  `xml:"pane"`
	selection CT_Selection  `xml:"selection"`
	rowBreaks CT_PageBreak  `xml:"rowBreaks"`
	colBreaks CT_PageBreak  `xml:"colBreaks"`
	pageMargins string  `xml:"pageMargins"`
	printOptions string  `xml:"printOptions"`
	pageSetup CT_PageSetup  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	autoFilter CT_AutoFilter  `xml:"autoFilter"`
	extLst string  `xml:"extLst"`
	guid string  `xml:"guid",attr`
	scale string  `xml:"scale",attr`
	colorId string  `xml:"colorId",attr`
	showPageBreaks string  `xml:"showPageBreaks",attr`
	showFormulas string  `xml:"showFormulas",attr`
	showGridLines string  `xml:"showGridLines",attr`
	showRowCol string  `xml:"showRowCol",attr`
	outlineSymbols string  `xml:"outlineSymbols",attr`
	zeroValues string  `xml:"zeroValues",attr`
	fitToPage string  `xml:"fitToPage",attr`
	printArea string  `xml:"printArea",attr`
	filter string  `xml:"filter",attr`
	showAutoFilter string  `xml:"showAutoFilter",attr`
	hiddenRows string  `xml:"hiddenRows",attr`
	hiddenColumns string  `xml:"hiddenColumns",attr`
	state ST_SheetState  `xml:"state",attr`
	filterUnique string  `xml:"filterUnique",attr`
	view ST_SheetViewType  `xml:"view",attr`
	showRuler string  `xml:"showRuler",attr`
	topLeftCell ST_CellRef  `xml:"topLeftCell",attr`
}
type CT_DataValidations struct{
	dataValidation []string  `xml:"dataValidation"`
	disablePrompts string  `xml:"disablePrompts",attr`
	xWindow string  `xml:"xWindow",attr`
	yWindow string  `xml:"yWindow",attr`
	count string  `xml:"count",attr`
}
type CT_DataValidation struct{
	formula1 ST_Formula  `xml:"formula1"`
	formula2 ST_Formula  `xml:"formula2"`
	type string  `xml:"type",attr`
	errorStyle string  `xml:"errorStyle",attr`
	imeMode string  `xml:"imeMode",attr`
	operator string  `xml:"operator",attr`
	allowBlank string  `xml:"allowBlank",attr`
	showDropDown string  `xml:"showDropDown",attr`
	showInputMessage string  `xml:"showInputMessage",attr`
	showErrorMessage string  `xml:"showErrorMessage",attr`
	errorTitle string  `xml:"errorTitle",attr`
	error string  `xml:"error",attr`
	promptTitle string  `xml:"promptTitle",attr`
	prompt string  `xml:"prompt",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
}
type CT_ConditionalFormatting struct{
	cfRule []CT_CfRule  `xml:"cfRule"`
	extLst string  `xml:"extLst"`
	pivot string  `xml:"pivot",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
}
type CT_CfRule struct{
	formula ST_Formula  `xml:"formula"`
	colorScale CT_ColorScale  `xml:"colorScale"`
	dataBar CT_DataBar  `xml:"dataBar"`
	iconSet CT_IconSet  `xml:"iconSet"`
	extLst string  `xml:"extLst"`
	type ST_CfType  `xml:"type",attr`
	dxfId string  `xml:"dxfId",attr`
	priority string  `xml:"priority",attr`
	stopIfTrue string  `xml:"stopIfTrue",attr`
	aboveAverage string  `xml:"aboveAverage",attr`
	percent string  `xml:"percent",attr`
	bottom string  `xml:"bottom",attr`
	operator string  `xml:"operator",attr`
	text string  `xml:"text",attr`
	timePeriod string  `xml:"timePeriod",attr`
	rank string  `xml:"rank",attr`
	stdDev string  `xml:"stdDev",attr`
	equalAverage string  `xml:"equalAverage",attr`
}
type CT_Hyperlinks struct{
	hyperlink []CT_Hyperlink  `xml:"hyperlink"`
}
type CT_Hyperlink struct{
	ref ST_Ref  `xml:"ref",attr`
	location string  `xml:"location",attr`
	tooltip string  `xml:"tooltip",attr`
	display string  `xml:"display",attr`
}
type CT_CellFormula struct{
}
type CT_ColorScale struct{
	cfvo []CT_Cfvo  `xml:"cfvo"`
	color []CT_Color  `xml:"color"`
}
type CT_DataBar struct{
	cfvo CT_Cfvo  `xml:"cfvo"`
	color CT_Color  `xml:"color"`
	minLength string  `xml:"minLength",attr`
	maxLength string  `xml:"maxLength",attr`
	showValue string  `xml:"showValue",attr`
}
type CT_IconSet struct{
	cfvo []CT_Cfvo  `xml:"cfvo"`
	iconSet ST_IconSetType  `xml:"iconSet",attr`
	showValue string  `xml:"showValue",attr`
	percent string  `xml:"percent",attr`
	reverse string  `xml:"reverse",attr`
}
type CT_Cfvo struct{
	extLst string  `xml:"extLst"`
	type ST_CfvoType  `xml:"type",attr`
	val string  `xml:"val",attr`
	gte string  `xml:"gte",attr`
}
type CT_PageMargins struct{
	left string  `xml:"left",attr`
	right string  `xml:"right",attr`
	top string  `xml:"top",attr`
	bottom string  `xml:"bottom",attr`
	header string  `xml:"header",attr`
	footer string  `xml:"footer",attr`
}
type CT_PrintOptions struct{
	horizontalCentered string  `xml:"horizontalCentered",attr`
	verticalCentered string  `xml:"verticalCentered",attr`
	headings string  `xml:"headings",attr`
	gridLines string  `xml:"gridLines",attr`
	gridLinesSet string  `xml:"gridLinesSet",attr`
}
type CT_PageSetup struct{
	paperSize string  `xml:"paperSize",attr`
	paperHeight string  `xml:"paperHeight",attr`
	paperWidth string  `xml:"paperWidth",attr`
	scale string  `xml:"scale",attr`
	firstPageNumber string  `xml:"firstPageNumber",attr`
	fitToWidth string  `xml:"fitToWidth",attr`
	fitToHeight string  `xml:"fitToHeight",attr`
	pageOrder string  `xml:"pageOrder",attr`
	orientation ST_Orientation  `xml:"orientation",attr`
	usePrinterDefaults string  `xml:"usePrinterDefaults",attr`
	blackAndWhite string  `xml:"blackAndWhite",attr`
	draft string  `xml:"draft",attr`
	cellComments string  `xml:"cellComments",attr`
	useFirstPageNumber string  `xml:"useFirstPageNumber",attr`
	errors ST_PrintError  `xml:"errors",attr`
	horizontalDpi string  `xml:"horizontalDpi",attr`
	verticalDpi string  `xml:"verticalDpi",attr`
	copies string  `xml:"copies",attr`
}
type CT_HeaderFooter struct{
	oddHeader string  `xml:"oddHeader"`
	oddFooter string  `xml:"oddFooter"`
	evenHeader string  `xml:"evenHeader"`
	evenFooter string  `xml:"evenFooter"`
	firstHeader string  `xml:"firstHeader"`
	firstFooter string  `xml:"firstFooter"`
	differentOddEven string  `xml:"differentOddEven",attr`
	differentFirst string  `xml:"differentFirst",attr`
	scaleWithDoc string  `xml:"scaleWithDoc",attr`
	alignWithMargins string  `xml:"alignWithMargins",attr`
}
type CT_Scenarios struct{
	scenario []CT_Scenario  `xml:"scenario"`
	current string  `xml:"current",attr`
	show string  `xml:"show",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
}
type CT_SheetProtection struct{
	password string  `xml:"password",attr`
	algorithmName string  `xml:"algorithmName",attr`
	hashValue string  `xml:"hashValue",attr`
	saltValue string  `xml:"saltValue",attr`
	spinCount string  `xml:"spinCount",attr`
	sheet string  `xml:"sheet",attr`
	objects string  `xml:"objects",attr`
	scenarios string  `xml:"scenarios",attr`
	formatCells string  `xml:"formatCells",attr`
	formatColumns string  `xml:"formatColumns",attr`
	formatRows string  `xml:"formatRows",attr`
	insertColumns string  `xml:"insertColumns",attr`
	insertRows string  `xml:"insertRows",attr`
	insertHyperlinks string  `xml:"insertHyperlinks",attr`
	deleteColumns string  `xml:"deleteColumns",attr`
	deleteRows string  `xml:"deleteRows",attr`
	selectLockedCells string  `xml:"selectLockedCells",attr`
	sort string  `xml:"sort",attr`
	autoFilter string  `xml:"autoFilter",attr`
	pivotTables string  `xml:"pivotTables",attr`
	selectUnlockedCells string  `xml:"selectUnlockedCells",attr`
}
type CT_ProtectedRanges struct{
	protectedRange []string  `xml:"protectedRange"`
}
type CT_ProtectedRange struct{
	securityDescriptor []string  `xml:"securityDescriptor"`
	password string  `xml:"password",attr`
	sqref ST_Sqref  `xml:"sqref",attr`
	name string  `xml:"name",attr`
	securityDescriptor string  `xml:"securityDescriptor",attr`
	algorithmName string  `xml:"algorithmName",attr`
	hashValue string  `xml:"hashValue",attr`
	saltValue string  `xml:"saltValue",attr`
	spinCount string  `xml:"spinCount",attr`
}
type CT_Scenario struct{
	inputCells []string  `xml:"inputCells"`
	name string  `xml:"name",attr`
	locked string  `xml:"locked",attr`
	hidden string  `xml:"hidden",attr`
	count string  `xml:"count",attr`
	user string  `xml:"user",attr`
	comment string  `xml:"comment",attr`
}
type CT_InputCells struct{
	r ST_CellRef  `xml:"r",attr`
	deleted string  `xml:"deleted",attr`
	undone string  `xml:"undone",attr`
	val string  `xml:"val",attr`
	numFmtId string  `xml:"numFmtId",attr`
}
type CT_CellWatches struct{
	cellWatch []CT_CellWatch  `xml:"cellWatch"`
}
type CT_CellWatch struct{
	r ST_CellRef  `xml:"r",attr`
}
type CT_Chartsheet struct{
	sheetPr string  `xml:"sheetPr"`
	sheetViews string  `xml:"sheetViews"`
	sheetProtection string  `xml:"sheetProtection"`
	customSheetViews string  `xml:"customSheetViews"`
	pageMargins string  `xml:"pageMargins"`
	pageSetup string  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	drawing CT_Drawing  `xml:"drawing"`
	legacyDrawing CT_LegacyDrawing  `xml:"legacyDrawing"`
	legacyDrawingHF CT_LegacyDrawing  `xml:"legacyDrawingHF"`
	drawingHF CT_DrawingHF  `xml:"drawingHF"`
	picture string  `xml:"picture"`
	webPublishItems string  `xml:"webPublishItems"`
	extLst string  `xml:"extLst"`
}
type CT_ChartsheetPr struct{
	tabColor CT_Color  `xml:"tabColor"`
	published string  `xml:"published",attr`
	codeName string  `xml:"codeName",attr`
}
type CT_ChartsheetViews struct{
	sheetView []string  `xml:"sheetView"`
	extLst string  `xml:"extLst"`
}
type CT_ChartsheetView struct{
	extLst string  `xml:"extLst"`
	tabSelected string  `xml:"tabSelected",attr`
	zoomScale string  `xml:"zoomScale",attr`
	workbookViewId string  `xml:"workbookViewId",attr`
	zoomToFit string  `xml:"zoomToFit",attr`
}
type CT_ChartsheetProtection struct{
	password string  `xml:"password",attr`
	algorithmName string  `xml:"algorithmName",attr`
	hashValue string  `xml:"hashValue",attr`
	saltValue string  `xml:"saltValue",attr`
	spinCount string  `xml:"spinCount",attr`
	content string  `xml:"content",attr`
	objects string  `xml:"objects",attr`
}
type CT_CsPageSetup struct{
	paperSize string  `xml:"paperSize",attr`
	paperHeight string  `xml:"paperHeight",attr`
	paperWidth string  `xml:"paperWidth",attr`
	firstPageNumber string  `xml:"firstPageNumber",attr`
	orientation ST_Orientation  `xml:"orientation",attr`
	usePrinterDefaults string  `xml:"usePrinterDefaults",attr`
	blackAndWhite string  `xml:"blackAndWhite",attr`
	draft string  `xml:"draft",attr`
	useFirstPageNumber string  `xml:"useFirstPageNumber",attr`
	horizontalDpi string  `xml:"horizontalDpi",attr`
	verticalDpi string  `xml:"verticalDpi",attr`
	copies string  `xml:"copies",attr`
}
type CT_CustomChartsheetViews struct{
	customSheetView []string  `xml:"customSheetView"`
}
type CT_CustomChartsheetView struct{
	pageMargins string  `xml:"pageMargins"`
	pageSetup string  `xml:"pageSetup"`
	headerFooter string  `xml:"headerFooter"`
	guid string  `xml:"guid",attr`
	scale string  `xml:"scale",attr`
	state ST_SheetState  `xml:"state",attr`
	zoomToFit string  `xml:"zoomToFit",attr`
}
type CT_CustomProperties struct{
	customPr []string  `xml:"customPr"`
}
type CT_CustomProperty struct{
	name string  `xml:"name",attr`
}
type CT_OleObjects struct{
	oleObject []CT_OleObject  `xml:"oleObject"`
}
type CT_OleObject struct{
	objectPr CT_ObjectPr  `xml:"objectPr"`
	progId string  `xml:"progId",attr`
	dvAspect string  `xml:"dvAspect",attr`
	link string  `xml:"link",attr`
	oleUpdate string  `xml:"oleUpdate",attr`
	autoLoad string  `xml:"autoLoad",attr`
	shapeId string  `xml:"shapeId",attr`
}
type CT_ObjectPr struct{
	anchor CT_ObjectAnchor  `xml:"anchor"`
	locked string  `xml:"locked",attr`
	defaultSize string  `xml:"defaultSize",attr`
	print string  `xml:"print",attr`
	disabled string  `xml:"disabled",attr`
	uiObject string  `xml:"uiObject",attr`
	autoFill string  `xml:"autoFill",attr`
	autoLine string  `xml:"autoLine",attr`
	autoPict string  `xml:"autoPict",attr`
	macro ST_Formula  `xml:"macro",attr`
	altText string  `xml:"altText",attr`
	dde string  `xml:"dde",attr`
}
type CT_WebPublishItems struct{
	webPublishItem []string  `xml:"webPublishItem"`
	count string  `xml:"count",attr`
}
type CT_WebPublishItem struct{
	id string  `xml:"id",attr`
	divId string  `xml:"divId",attr`
	sourceType ST_WebSourceType  `xml:"sourceType",attr`
	sourceRef ST_Ref  `xml:"sourceRef",attr`
	sourceObject string  `xml:"sourceObject",attr`
	destinationFile string  `xml:"destinationFile",attr`
	title string  `xml:"title",attr`
	autoRepublish string  `xml:"autoRepublish",attr`
}
type CT_Controls struct{
	control []CT_Control  `xml:"control"`
}
type CT_Control struct{
	controlPr CT_ControlPr  `xml:"controlPr"`
	shapeId string  `xml:"shapeId",attr`
	name string  `xml:"name",attr`
}
type CT_ControlPr struct{
	anchor CT_ObjectAnchor  `xml:"anchor"`
	locked string  `xml:"locked",attr`
	defaultSize string  `xml:"defaultSize",attr`
	print string  `xml:"print",attr`
	disabled string  `xml:"disabled",attr`
	recalcAlways string  `xml:"recalcAlways",attr`
	uiObject string  `xml:"uiObject",attr`
	autoFill string  `xml:"autoFill",attr`
	autoLine string  `xml:"autoLine",attr`
	autoPict string  `xml:"autoPict",attr`
	macro ST_Formula  `xml:"macro",attr`
	altText string  `xml:"altText",attr`
	linkedCell ST_Formula  `xml:"linkedCell",attr`
	listFillRange ST_Formula  `xml:"listFillRange",attr`
	cf string  `xml:"cf",attr`
}
type CT_IgnoredErrors struct{
	ignoredError []string  `xml:"ignoredError"`
	extLst string  `xml:"extLst"`
}
type CT_IgnoredError struct{
	sqref ST_Sqref  `xml:"sqref",attr`
	evalError string  `xml:"evalError",attr`
	twoDigitTextYear string  `xml:"twoDigitTextYear",attr`
	numberStoredAsText string  `xml:"numberStoredAsText",attr`
	formula string  `xml:"formula",attr`
	formulaRange string  `xml:"formulaRange",attr`
	unlockedFormula string  `xml:"unlockedFormula",attr`
	emptyCellReference string  `xml:"emptyCellReference",attr`
	listDataValidation string  `xml:"listDataValidation",attr`
	calculatedColumn string  `xml:"calculatedColumn",attr`
}
type CT_TableParts struct{
	tablePart []CT_TablePart  `xml:"tablePart"`
	count string  `xml:"count",attr`
}
type CT_TablePart struct{
}
type CT_Metadata struct{
	metadataTypes string  `xml:"metadataTypes"`
	metadataStrings string  `xml:"metadataStrings"`
	mdxMetadata string  `xml:"mdxMetadata"`
	futureMetadata []string  `xml:"futureMetadata"`
	cellMetadata string  `xml:"cellMetadata"`
	valueMetadata string  `xml:"valueMetadata"`
	extLst string  `xml:"extLst"`
}
type CT_MetadataTypes struct{
	metadataType []string  `xml:"metadataType"`
	count string  `xml:"count",attr`
}
type CT_MetadataType struct{
	name string  `xml:"name",attr`
	minSupportedVersion string  `xml:"minSupportedVersion",attr`
	ghostRow string  `xml:"ghostRow",attr`
	ghostCol string  `xml:"ghostCol",attr`
	edit string  `xml:"edit",attr`
	delete string  `xml:"delete",attr`
	copy string  `xml:"copy",attr`
	pasteAll string  `xml:"pasteAll",attr`
	pasteFormulas string  `xml:"pasteFormulas",attr`
	pasteValues string  `xml:"pasteValues",attr`
	pasteFormats string  `xml:"pasteFormats",attr`
	pasteComments string  `xml:"pasteComments",attr`
	pasteDataValidation string  `xml:"pasteDataValidation",attr`
	pasteBorders string  `xml:"pasteBorders",attr`
	pasteColWidths string  `xml:"pasteColWidths",attr`
	pasteNumberFormats string  `xml:"pasteNumberFormats",attr`
	merge string  `xml:"merge",attr`
	splitFirst string  `xml:"splitFirst",attr`
	splitAll string  `xml:"splitAll",attr`
	rowColShift string  `xml:"rowColShift",attr`
	clearAll string  `xml:"clearAll",attr`
	clearFormats string  `xml:"clearFormats",attr`
	clearContents string  `xml:"clearContents",attr`
	clearComments string  `xml:"clearComments",attr`
	assign string  `xml:"assign",attr`
	coerce string  `xml:"coerce",attr`
	adjust string  `xml:"adjust",attr`
	cellMeta string  `xml:"cellMeta",attr`
}
type CT_MetadataBlocks struct{
	bk []string  `xml:"bk"`
	count string  `xml:"count",attr`
}
type CT_MetadataBlock struct{
	rc []string  `xml:"rc"`
}
type CT_MetadataRecord struct{
	t string  `xml:"t",attr`
	v string  `xml:"v",attr`
}
type CT_FutureMetadata struct{
	bk []string  `xml:"bk"`
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	count string  `xml:"count",attr`
}
type CT_FutureMetadataBlock struct{
	extLst string  `xml:"extLst"`
}
type CT_MdxMetadata struct{
	mdx []string  `xml:"mdx"`
	count string  `xml:"count",attr`
}
type CT_Mdx struct{
	t string  `xml:"t"`
	ms string  `xml:"ms"`
	p string  `xml:"p"`
	k string  `xml:"k"`
	n string  `xml:"n",attr`
	f string  `xml:"f",attr`
}
type CT_MdxTuple struct{
	n []string  `xml:"n"`
	c string  `xml:"c",attr`
	ct string  `xml:"ct",attr`
	si string  `xml:"si",attr`
	fi string  `xml:"fi",attr`
	bc string  `xml:"bc",attr`
	fc string  `xml:"fc",attr`
	i string  `xml:"i",attr`
	u string  `xml:"u",attr`
	st string  `xml:"st",attr`
	b string  `xml:"b",attr`
}
type CT_MdxSet struct{
	n []string  `xml:"n"`
	ns string  `xml:"ns",attr`
	c string  `xml:"c",attr`
	o string  `xml:"o",attr`
}
type CT_MdxMemeberProp struct{
	n string  `xml:"n",attr`
	np string  `xml:"np",attr`
}
type CT_MdxKPI struct{
	n string  `xml:"n",attr`
	np string  `xml:"np",attr`
	p string  `xml:"p",attr`
}
type CT_MetadataStringIndex struct{
	x string  `xml:"x",attr`
	s string  `xml:"s",attr`
}
type CT_MetadataStrings struct{
	s []CT_XStringElement  `xml:"s"`
	count string  `xml:"count",attr`
}
type CT_SingleXmlCells struct{
	singleXmlCell []CT_SingleXmlCell  `xml:"singleXmlCell"`
}
type CT_SingleXmlCell struct{
	xmlCellPr CT_XmlCellPr  `xml:"xmlCellPr"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	r ST_CellRef  `xml:"r",attr`
	connectionId string  `xml:"connectionId",attr`
}
type CT_XmlCellPr struct{
	xmlPr CT_XmlPr  `xml:"xmlPr"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	uniqueName string  `xml:"uniqueName",attr`
}
type CT_XmlPr struct{
	extLst string  `xml:"extLst"`
	mapId string  `xml:"mapId",attr`
	xpath string  `xml:"xpath",attr`
	xmlDataType ST_XmlDataType  `xml:"xmlDataType",attr`
}
type CT_Stylesheet struct{
	numFmts string  `xml:"numFmts"`
	fonts string  `xml:"fonts"`
	fills string  `xml:"fills"`
	borders string  `xml:"borders"`
	cellStyleXfs string  `xml:"cellStyleXfs"`
	cellXfs string  `xml:"cellXfs"`
	cellStyles string  `xml:"cellStyles"`
	dxfs string  `xml:"dxfs"`
	tableStyles string  `xml:"tableStyles"`
	colors string  `xml:"colors"`
	extLst string  `xml:"extLst"`
}
type CT_CellAlignment struct{
	horizontal ST_HorizontalAlignment  `xml:"horizontal",attr`
	vertical ST_VerticalAlignment  `xml:"vertical",attr`
	textRotation string  `xml:"textRotation",attr`
	wrapText string  `xml:"wrapText",attr`
	indent string  `xml:"indent",attr`
	relativeIndent string  `xml:"relativeIndent",attr`
	justifyLastLine string  `xml:"justifyLastLine",attr`
	shrinkToFit string  `xml:"shrinkToFit",attr`
	readingOrder string  `xml:"readingOrder",attr`
}
type CT_Borders struct{
	border []string  `xml:"border"`
	count string  `xml:"count",attr`
}
type CT_Border struct{
	start string  `xml:"start"`
	end string  `xml:"end"`
	left string  `xml:"left"`
	right string  `xml:"right"`
	top string  `xml:"top"`
	bottom string  `xml:"bottom"`
	diagonal string  `xml:"diagonal"`
	vertical string  `xml:"vertical"`
	horizontal string  `xml:"horizontal"`
	diagonalUp string  `xml:"diagonalUp",attr`
	diagonalDown string  `xml:"diagonalDown",attr`
	outline string  `xml:"outline",attr`
}
type CT_BorderPr struct{
	color CT_Color  `xml:"color"`
	style string  `xml:"style",attr`
}
type CT_CellProtection struct{
	locked string  `xml:"locked",attr`
	hidden string  `xml:"hidden",attr`
}
type CT_Fonts struct{
	font []CT_Font  `xml:"font"`
	count string  `xml:"count",attr`
}
type CT_Fills struct{
	fill []CT_Fill  `xml:"fill"`
	count string  `xml:"count",attr`
}
type CT_Fill struct{
	patternFill CT_PatternFill  `xml:"patternFill"`
	gradientFill string  `xml:"gradientFill"`
}
type CT_PatternFill struct{
	fgColor CT_Color  `xml:"fgColor"`
	bgColor CT_Color  `xml:"bgColor"`
	patternType ST_PatternType  `xml:"patternType",attr`
}
type CT_Color struct{
	auto string  `xml:"auto",attr`
	indexed string  `xml:"indexed",attr`
	rgb string  `xml:"rgb",attr`
	theme string  `xml:"theme",attr`
	tint string  `xml:"tint",attr`
}
type CT_GradientFill struct{
	stop []string  `xml:"stop"`
	type string  `xml:"type",attr`
	degree string  `xml:"degree",attr`
	left string  `xml:"left",attr`
	right string  `xml:"right",attr`
	top string  `xml:"top",attr`
	bottom string  `xml:"bottom",attr`
}
type CT_GradientStop struct{
	color CT_Color  `xml:"color"`
	position string  `xml:"position",attr`
}
type CT_NumFmts struct{
	numFmt []CT_NumFmt  `xml:"numFmt"`
	count string  `xml:"count",attr`
}
type CT_NumFmt struct{
	numFmtId string  `xml:"numFmtId",attr`
	formatCode string  `xml:"formatCode",attr`
}
type CT_CellStyleXfs struct{
	xf []CT_Xf  `xml:"xf"`
	count string  `xml:"count",attr`
}
type CT_CellXfs struct{
	xf []CT_Xf  `xml:"xf"`
	count string  `xml:"count",attr`
}
type CT_Xf struct{
	alignment CT_CellAlignment  `xml:"alignment"`
	protection CT_CellProtection  `xml:"protection"`
	extLst string  `xml:"extLst"`
	numFmtId string  `xml:"numFmtId",attr`
	fontId string  `xml:"fontId",attr`
	fillId string  `xml:"fillId",attr`
	borderId string  `xml:"borderId",attr`
	xfId string  `xml:"xfId",attr`
	quotePrefix string  `xml:"quotePrefix",attr`
	pivotButton string  `xml:"pivotButton",attr`
	applyNumberFormat string  `xml:"applyNumberFormat",attr`
	applyFont string  `xml:"applyFont",attr`
	applyFill string  `xml:"applyFill",attr`
	applyBorder string  `xml:"applyBorder",attr`
	applyAlignment string  `xml:"applyAlignment",attr`
	applyProtection string  `xml:"applyProtection",attr`
}
type CT_CellStyles struct{
	cellStyle []CT_CellStyle  `xml:"cellStyle"`
	count string  `xml:"count",attr`
}
type CT_CellStyle struct{
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	xfId string  `xml:"xfId",attr`
	builtinId string  `xml:"builtinId",attr`
	iLevel string  `xml:"iLevel",attr`
	hidden string  `xml:"hidden",attr`
	customBuiltin string  `xml:"customBuiltin",attr`
}
type CT_Dxfs struct{
	dxf []string  `xml:"dxf"`
	count string  `xml:"count",attr`
}
type CT_Dxf struct{
	font CT_Font  `xml:"font"`
	numFmt CT_NumFmt  `xml:"numFmt"`
	fill CT_Fill  `xml:"fill"`
	alignment CT_CellAlignment  `xml:"alignment"`
	border string  `xml:"border"`
	protection CT_CellProtection  `xml:"protection"`
	extLst string  `xml:"extLst"`
}
type CT_Colors struct{
	indexedColors string  `xml:"indexedColors"`
	mruColors string  `xml:"mruColors"`
}
type CT_IndexedColors struct{
	rgbColor []CT_RgbColor  `xml:"rgbColor"`
}
type CT_MRUColors struct{
	color []CT_Color  `xml:"color"`
}
type CT_RgbColor struct{
	rgb string  `xml:"rgb",attr`
}
type CT_TableStyles struct{
	tableStyle []CT_TableStyle  `xml:"tableStyle"`
	count string  `xml:"count",attr`
	defaultTableStyle string  `xml:"defaultTableStyle",attr`
	defaultPivotStyle string  `xml:"defaultPivotStyle",attr`
}
type CT_TableStyle struct{
	tableStyleElement []CT_TableStyleElement  `xml:"tableStyleElement"`
	name string  `xml:"name",attr`
	pivot string  `xml:"pivot",attr`
	table string  `xml:"table",attr`
	count string  `xml:"count",attr`
}
type CT_TableStyleElement struct{
	type ST_TableStyleType  `xml:"type",attr`
	size string  `xml:"size",attr`
	dxfId string  `xml:"dxfId",attr`
}
type CT_BooleanProperty struct{
	val string  `xml:"val",attr`
}
type CT_FontSize struct{
	val string  `xml:"val",attr`
}
type CT_IntProperty struct{
	val string  `xml:"val",attr`
}
type CT_FontName struct{
	val string  `xml:"val",attr`
}
type CT_VerticalAlignFontProperty struct{
	val string  `xml:"val",attr`
}
type CT_FontScheme struct{
	val ST_FontScheme  `xml:"val",attr`
}
type CT_UnderlineProperty struct{
	val string  `xml:"val",attr`
}
type CT_Font struct{
	name CT_FontName  `xml:"name"`
	charset CT_IntProperty  `xml:"charset"`
	family CT_FontFamily  `xml:"family"`
	b CT_BooleanProperty  `xml:"b"`
	i CT_BooleanProperty  `xml:"i"`
	strike CT_BooleanProperty  `xml:"strike"`
	outline CT_BooleanProperty  `xml:"outline"`
	shadow CT_BooleanProperty  `xml:"shadow"`
	condense CT_BooleanProperty  `xml:"condense"`
	extend CT_BooleanProperty  `xml:"extend"`
	color CT_Color  `xml:"color"`
	sz CT_FontSize  `xml:"sz"`
	u string  `xml:"u"`
	vertAlign CT_VerticalAlignFontProperty  `xml:"vertAlign"`
	scheme CT_FontScheme  `xml:"scheme"`
}
type CT_FontFamily struct{
	val ST_FontFamily  `xml:"val",attr`
}
type CT_ExternalLink struct{
	extLst string  `xml:"extLst"`
	externalBook string  `xml:"externalBook"`
	ddeLink string  `xml:"ddeLink"`
	oleLink CT_OleLink  `xml:"oleLink"`
}
type CT_ExternalBook struct{
	sheetNames string  `xml:"sheetNames"`
	definedNames string  `xml:"definedNames"`
	sheetDataSet string  `xml:"sheetDataSet"`
}
type CT_ExternalSheetNames struct{
	sheetName []string  `xml:"sheetName"`
}
type CT_ExternalSheetName struct{
	val string  `xml:"val",attr`
}
type CT_ExternalDefinedNames struct{
	definedName []string  `xml:"definedName"`
}
type CT_ExternalDefinedName struct{
	name string  `xml:"name",attr`
	refersTo string  `xml:"refersTo",attr`
	sheetId string  `xml:"sheetId",attr`
}
type CT_ExternalSheetDataSet struct{
	sheetData []string  `xml:"sheetData"`
}
type CT_ExternalSheetData struct{
	row []string  `xml:"row"`
	sheetId string  `xml:"sheetId",attr`
	refreshError string  `xml:"refreshError",attr`
}
type CT_ExternalRow struct{
	cell []string  `xml:"cell"`
	r string  `xml:"r",attr`
}
type CT_ExternalCell struct{
	v string  `xml:"v"`
	r ST_CellRef  `xml:"r",attr`
	t ST_CellType  `xml:"t",attr`
	vm string  `xml:"vm",attr`
}
type CT_DdeLink struct{
	ddeItems string  `xml:"ddeItems"`
	ddeService string  `xml:"ddeService",attr`
	ddeTopic string  `xml:"ddeTopic",attr`
}
type CT_DdeItems struct{
	ddeItem []string  `xml:"ddeItem"`
}
type CT_DdeItem struct{
	values string  `xml:"values"`
	name string  `xml:"name",attr`
	ole string  `xml:"ole",attr`
	advise string  `xml:"advise",attr`
	preferPic string  `xml:"preferPic",attr`
}
type CT_DdeValues struct{
	value []string  `xml:"value"`
	rows string  `xml:"rows",attr`
	cols string  `xml:"cols",attr`
}
type CT_DdeValue struct{
	val string  `xml:"val"`
	t string  `xml:"t",attr`
}
type CT_OleLink struct{
	oleItems string  `xml:"oleItems"`
	progId string  `xml:"progId",attr`
}
type CT_OleItems struct{
	oleItem []CT_OleItem  `xml:"oleItem"`
}
type CT_OleItem struct{
	name string  `xml:"name",attr`
	icon string  `xml:"icon",attr`
	advise string  `xml:"advise",attr`
	preferPic string  `xml:"preferPic",attr`
}
type CT_Table struct{
	autoFilter CT_AutoFilter  `xml:"autoFilter"`
	sortState CT_SortState  `xml:"sortState"`
	tableColumns string  `xml:"tableColumns"`
	tableStyleInfo CT_TableStyleInfo  `xml:"tableStyleInfo"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	name string  `xml:"name",attr`
	displayName string  `xml:"displayName",attr`
	comment string  `xml:"comment",attr`
	ref ST_Ref  `xml:"ref",attr`
	tableType ST_TableType  `xml:"tableType",attr`
	headerRowCount string  `xml:"headerRowCount",attr`
	insertRow string  `xml:"insertRow",attr`
	insertRowShift string  `xml:"insertRowShift",attr`
	totalsRowCount string  `xml:"totalsRowCount",attr`
	totalsRowShown string  `xml:"totalsRowShown",attr`
	published string  `xml:"published",attr`
	headerRowDxfId string  `xml:"headerRowDxfId",attr`
	dataDxfId string  `xml:"dataDxfId",attr`
	totalsRowDxfId string  `xml:"totalsRowDxfId",attr`
	headerRowBorderDxfId string  `xml:"headerRowBorderDxfId",attr`
	tableBorderDxfId string  `xml:"tableBorderDxfId",attr`
	totalsRowBorderDxfId string  `xml:"totalsRowBorderDxfId",attr`
	headerRowCellStyle string  `xml:"headerRowCellStyle",attr`
	dataCellStyle string  `xml:"dataCellStyle",attr`
	totalsRowCellStyle string  `xml:"totalsRowCellStyle",attr`
	connectionId string  `xml:"connectionId",attr`
}
type CT_TableStyleInfo struct{
	name string  `xml:"name",attr`
	showFirstColumn string  `xml:"showFirstColumn",attr`
	showLastColumn string  `xml:"showLastColumn",attr`
	showRowStripes string  `xml:"showRowStripes",attr`
	showColumnStripes string  `xml:"showColumnStripes",attr`
}
type CT_TableColumns struct{
	tableColumn []CT_TableColumn  `xml:"tableColumn"`
	count string  `xml:"count",attr`
}
type CT_TableColumn struct{
	calculatedColumnFormula CT_TableFormula  `xml:"calculatedColumnFormula"`
	totalsRowFormula CT_TableFormula  `xml:"totalsRowFormula"`
	xmlColumnPr CT_XmlColumnPr  `xml:"xmlColumnPr"`
	extLst string  `xml:"extLst"`
	id string  `xml:"id",attr`
	uniqueName string  `xml:"uniqueName",attr`
	name string  `xml:"name",attr`
	totalsRowFunction string  `xml:"totalsRowFunction",attr`
	totalsRowLabel string  `xml:"totalsRowLabel",attr`
	queryTableFieldId string  `xml:"queryTableFieldId",attr`
	headerRowDxfId string  `xml:"headerRowDxfId",attr`
	dataDxfId string  `xml:"dataDxfId",attr`
	totalsRowDxfId string  `xml:"totalsRowDxfId",attr`
	headerRowCellStyle string  `xml:"headerRowCellStyle",attr`
	dataCellStyle string  `xml:"dataCellStyle",attr`
	totalsRowCellStyle string  `xml:"totalsRowCellStyle",attr`
}
type CT_TableFormula struct{
}
type CT_XmlColumnPr struct{
	extLst string  `xml:"extLst"`
	mapId string  `xml:"mapId",attr`
	xpath string  `xml:"xpath",attr`
	denormalized string  `xml:"denormalized",attr`
	xmlDataType ST_XmlDataType  `xml:"xmlDataType",attr`
}
type CT_VolTypes struct{
	volType []CT_VolType  `xml:"volType"`
	extLst string  `xml:"extLst"`
}
type CT_VolType struct{
	main []CT_VolMain  `xml:"main"`
	type ST_VolDepType  `xml:"type",attr`
}
type CT_VolMain struct{
	tp []CT_VolTopic  `xml:"tp"`
	first string  `xml:"first",attr`
}
type CT_VolTopic struct{
	v string  `xml:"v"`
	stp []string  `xml:"stp"`
	tr []CT_VolTopicRef  `xml:"tr"`
	t ST_VolValueType  `xml:"t",attr`
}
type CT_VolTopicRef struct{
	r ST_CellRef  `xml:"r",attr`
	s string  `xml:"s",attr`
}
type CT_Workbook struct{
	fileVersion string  `xml:"fileVersion"`
	fileSharing CT_FileSharing  `xml:"fileSharing"`
	workbookPr CT_WorkbookPr  `xml:"workbookPr"`
	workbookProtection CT_WorkbookProtection  `xml:"workbookProtection"`
	bookViews string  `xml:"bookViews"`
	sheets string  `xml:"sheets"`
	functionGroups string  `xml:"functionGroups"`
	externalReferences string  `xml:"externalReferences"`
	definedNames string  `xml:"definedNames"`
	calcPr CT_CalcPr  `xml:"calcPr"`
	oleSize CT_OleSize  `xml:"oleSize"`
	customWorkbookViews string  `xml:"customWorkbookViews"`
	pivotCaches string  `xml:"pivotCaches"`
	smartTagPr CT_SmartTagPr  `xml:"smartTagPr"`
	smartTagTypes string  `xml:"smartTagTypes"`
	webPublishing string  `xml:"webPublishing"`
	fileRecoveryPr []CT_FileRecoveryPr  `xml:"fileRecoveryPr"`
	webPublishObjects string  `xml:"webPublishObjects"`
	extLst string  `xml:"extLst"`
	conformance string  `xml:"conformance",attr`
}
type CT_FileVersion struct{
	appName string  `xml:"appName",attr`
	lastEdited string  `xml:"lastEdited",attr`
	lowestEdited string  `xml:"lowestEdited",attr`
	rupBuild string  `xml:"rupBuild",attr`
	codeName string  `xml:"codeName",attr`
}
type CT_BookViews struct{
	workbookView []CT_BookView  `xml:"workbookView"`
}
type CT_BookView struct{
	extLst string  `xml:"extLst"`
	visibility string  `xml:"visibility",attr`
	minimized string  `xml:"minimized",attr`
	showHorizontalScroll string  `xml:"showHorizontalScroll",attr`
	showVerticalScroll string  `xml:"showVerticalScroll",attr`
	showSheetTabs string  `xml:"showSheetTabs",attr`
	xWindow string  `xml:"xWindow",attr`
	yWindow string  `xml:"yWindow",attr`
	windowWidth string  `xml:"windowWidth",attr`
	windowHeight string  `xml:"windowHeight",attr`
	tabRatio string  `xml:"tabRatio",attr`
	firstSheet string  `xml:"firstSheet",attr`
	activeTab string  `xml:"activeTab",attr`
	autoFilterDateGrouping string  `xml:"autoFilterDateGrouping",attr`
}
type CT_CustomWorkbookViews struct{
	customWorkbookView []string  `xml:"customWorkbookView"`
}
type CT_CustomWorkbookView struct{
	extLst string  `xml:"extLst"`
	name string  `xml:"name",attr`
	guid string  `xml:"guid",attr`
	autoUpdate string  `xml:"autoUpdate",attr`
	mergeInterval string  `xml:"mergeInterval",attr`
	changesSavedWin string  `xml:"changesSavedWin",attr`
	onlySync string  `xml:"onlySync",attr`
	personalView string  `xml:"personalView",attr`
	includePrintSettings string  `xml:"includePrintSettings",attr`
	includeHiddenRowCol string  `xml:"includeHiddenRowCol",attr`
	maximized string  `xml:"maximized",attr`
	minimized string  `xml:"minimized",attr`
	showHorizontalScroll string  `xml:"showHorizontalScroll",attr`
	showVerticalScroll string  `xml:"showVerticalScroll",attr`
	showSheetTabs string  `xml:"showSheetTabs",attr`
	xWindow string  `xml:"xWindow",attr`
	yWindow string  `xml:"yWindow",attr`
	windowWidth string  `xml:"windowWidth",attr`
	windowHeight string  `xml:"windowHeight",attr`
	tabRatio string  `xml:"tabRatio",attr`
	activeSheetId string  `xml:"activeSheetId",attr`
	showFormulaBar string  `xml:"showFormulaBar",attr`
	showStatusbar string  `xml:"showStatusbar",attr`
	showComments string  `xml:"showComments",attr`
	showObjects string  `xml:"showObjects",attr`
}
type CT_Sheets struct{
	sheet []CT_Sheet  `xml:"sheet"`
}
type CT_Sheet struct{
	name string  `xml:"name",attr`
	sheetId string  `xml:"sheetId",attr`
	state ST_SheetState  `xml:"state",attr`
}
type CT_WorkbookPr struct{
	date1904 string  `xml:"date1904",attr`
	showObjects string  `xml:"showObjects",attr`
	showBorderUnselectedTables string  `xml:"showBorderUnselectedTables",attr`
	filterPrivacy string  `xml:"filterPrivacy",attr`
	promptedSolutions string  `xml:"promptedSolutions",attr`
	showInkAnnotation string  `xml:"showInkAnnotation",attr`
	backupFile string  `xml:"backupFile",attr`
	saveExternalLinkValues string  `xml:"saveExternalLinkValues",attr`
	updateLinks string  `xml:"updateLinks",attr`
	codeName string  `xml:"codeName",attr`
	hidePivotFieldList string  `xml:"hidePivotFieldList",attr`
	showPivotChartFilter string  `xml:"showPivotChartFilter",attr`
	allowRefreshQuery string  `xml:"allowRefreshQuery",attr`
	publishItems string  `xml:"publishItems",attr`
	checkCompatibility string  `xml:"checkCompatibility",attr`
	autoCompressPictures string  `xml:"autoCompressPictures",attr`
	refreshAllConnections string  `xml:"refreshAllConnections",attr`
	defaultThemeVersion string  `xml:"defaultThemeVersion",attr`
}
type CT_SmartTagPr struct{
	embed string  `xml:"embed",attr`
	show ST_SmartTagShow  `xml:"show",attr`
}
type CT_SmartTagTypes struct{
	smartTagType []CT_SmartTagType  `xml:"smartTagType"`
}
type CT_SmartTagType struct{
	namespaceUri string  `xml:"namespaceUri",attr`
	name string  `xml:"name",attr`
	url string  `xml:"url",attr`
}
type CT_FileRecoveryPr struct{
	autoRecover string  `xml:"autoRecover",attr`
	crashSave string  `xml:"crashSave",attr`
	dataExtractLoad string  `xml:"dataExtractLoad",attr`
	repairLoad string  `xml:"repairLoad",attr`
}
type CT_CalcPr struct{
	calcId string  `xml:"calcId",attr`
	calcMode string  `xml:"calcMode",attr`
	fullCalcOnLoad string  `xml:"fullCalcOnLoad",attr`
	refMode string  `xml:"refMode",attr`
	iterate string  `xml:"iterate",attr`
	iterateCount string  `xml:"iterateCount",attr`
	iterateDelta string  `xml:"iterateDelta",attr`
	fullPrecision string  `xml:"fullPrecision",attr`
	calcCompleted string  `xml:"calcCompleted",attr`
	calcOnSave string  `xml:"calcOnSave",attr`
	concurrentCalc string  `xml:"concurrentCalc",attr`
	concurrentManualCount string  `xml:"concurrentManualCount",attr`
	forceFullCalc string  `xml:"forceFullCalc",attr`
}
type CT_DefinedNames struct{
	definedName []string  `xml:"definedName"`
}
type CT_DefinedName struct{
}
type CT_ExternalReferences struct{
	externalReference []string  `xml:"externalReference"`
}
type CT_ExternalReference struct{
}
type CT_SheetBackgroundPicture struct{
}
type CT_PivotCaches struct{
	pivotCache []CT_PivotCache  `xml:"pivotCache"`
}
type CT_PivotCache struct{
	cacheId string  `xml:"cacheId",attr`
}
type CT_FileSharing struct{
	readOnlyRecommended string  `xml:"readOnlyRecommended",attr`
	userName string  `xml:"userName",attr`
	reservationPassword string  `xml:"reservationPassword",attr`
	algorithmName string  `xml:"algorithmName",attr`
	hashValue string  `xml:"hashValue",attr`
	saltValue string  `xml:"saltValue",attr`
	spinCount string  `xml:"spinCount",attr`
}
type CT_OleSize struct{
	ref ST_Ref  `xml:"ref",attr`
}
type CT_WorkbookProtection struct{
	workbookPassword string  `xml:"workbookPassword",attr`
	workbookPasswordCharacterSet string  `xml:"workbookPasswordCharacterSet",attr`
	revisionsPassword string  `xml:"revisionsPassword",attr`
	revisionsPasswordCharacterSet string  `xml:"revisionsPasswordCharacterSet",attr`
	lockStructure string  `xml:"lockStructure",attr`
	lockWindows string  `xml:"lockWindows",attr`
	lockRevision string  `xml:"lockRevision",attr`
	revisionsAlgorithmName string  `xml:"revisionsAlgorithmName",attr`
	revisionsHashValue string  `xml:"revisionsHashValue",attr`
	revisionsSaltValue string  `xml:"revisionsSaltValue",attr`
	revisionsSpinCount string  `xml:"revisionsSpinCount",attr`
	workbookAlgorithmName string  `xml:"workbookAlgorithmName",attr`
	workbookHashValue string  `xml:"workbookHashValue",attr`
	workbookSaltValue string  `xml:"workbookSaltValue",attr`
	workbookSpinCount string  `xml:"workbookSpinCount",attr`
}
type CT_WebPublishing struct{
	css string  `xml:"css",attr`
	thicket string  `xml:"thicket",attr`
	longFileNames string  `xml:"longFileNames",attr`
	vml string  `xml:"vml",attr`
	allowPng string  `xml:"allowPng",attr`
	targetScreenSize ST_TargetScreenSize  `xml:"targetScreenSize",attr`
	dpi string  `xml:"dpi",attr`
	codePage string  `xml:"codePage",attr`
	characterSet string  `xml:"characterSet",attr`
}
type CT_FunctionGroups struct{
	functionGroup CT_FunctionGroup  `xml:"functionGroup"`
	builtInGroupCount string  `xml:"builtInGroupCount",attr`
}
type CT_FunctionGroup struct{
	name string  `xml:"name",attr`
}
type CT_WebPublishObjects struct{
	webPublishObject []string  `xml:"webPublishObject"`
	count string  `xml:"count",attr`
}
type CT_WebPublishObject struct{
	id string  `xml:"id",attr`
	divId string  `xml:"divId",attr`
	sourceObject string  `xml:"sourceObject",attr`
	destinationFile string  `xml:"destinationFile",attr`
	title string  `xml:"title",attr`
	autoRepublish string  `xml:"autoRepublish",attr`
}
