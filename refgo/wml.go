package ooxml

type WmlSchema struct {
	Recipients       string       `xml:"recipients"`
	TxbxContent      string       `xml:"txbxContent"`
	Comments         string       `xml:"comments"`
	Footnotes        string       `xml:"footnotes"`
	Endnotes         string       `xml:"endnotes"`
	Hdr              string       `xml:"hdr"`
	Ftr              string       `xml:"ftr"`
	Settings         string       `xml:"settings"`
	WebSettings      string       `xml:"webSettings"`
	Fonts            string       `xml:"fonts"`
	Numbering        CT_Numbering `xml:"numbering"`
	Styles           string       `xml:"styles"`
	Document         CT_Document  `xml:"document"`
	GlossaryDocument string       `xml:"glossaryDocument"`
}
type EG_PContentMath struct {
	PContentBase          []EG_PContentBase
	ContentRunContentBase []EG_ContentRunContentBase
}
type EG_PContentBase struct {
	CustomXml string       `xml:"customXml"`
	FldSimple []string     `xml:"fldSimple"`
	Hyperlink CT_Hyperlink `xml:"hyperlink"`
}
type EG_ContentRunContentBase struct {
	RunLevelElts []EG_RunLevelElts
	SmartTag     CT_SmartTagRun `xml:"smartTag"`
	Sdt          string         `xml:"sdt"`
}
type EG_CellMarkupElements struct {
	CellIns   CT_TrackChange          `xml:"cellIns"`
	CellDel   CT_TrackChange          `xml:"cellDel"`
	CellMerge CT_CellMergeTrackChange `xml:"cellMerge"`
}
type EG_RangeMarkupElements struct {
	BookmarkStart               CT_Bookmark     `xml:"bookmarkStart"`
	BookmarkEnd                 CT_MarkupRange  `xml:"bookmarkEnd"`
	MoveFromRangeStart          CT_MoveBookmark `xml:"moveFromRangeStart"`
	MoveFromRangeEnd            CT_MarkupRange  `xml:"moveFromRangeEnd"`
	MoveToRangeStart            CT_MoveBookmark `xml:"moveToRangeStart"`
	MoveToRangeEnd              CT_MarkupRange  `xml:"moveToRangeEnd"`
	CommentRangeStart           CT_MarkupRange  `xml:"commentRangeStart"`
	CommentRangeEnd             CT_MarkupRange  `xml:"commentRangeEnd"`
	CustomXmlInsRangeStart      CT_TrackChange  `xml:"customXmlInsRangeStart"`
	CustomXmlInsRangeEnd        CT_Markup       `xml:"customXmlInsRangeEnd"`
	CustomXmlDelRangeStart      CT_TrackChange  `xml:"customXmlDelRangeStart"`
	CustomXmlDelRangeEnd        CT_Markup       `xml:"customXmlDelRangeEnd"`
	CustomXmlMoveFromRangeStart CT_TrackChange  `xml:"customXmlMoveFromRangeStart"`
	CustomXmlMoveFromRangeEnd   CT_Markup       `xml:"customXmlMoveFromRangeEnd"`
	CustomXmlMoveToRangeStart   CT_TrackChange  `xml:"customXmlMoveToRangeStart"`
	CustomXmlMoveToRangeEnd     CT_Markup       `xml:"customXmlMoveToRangeEnd"`
}
type EG_HdrFtrReferences struct {
	HeaderReference string `xml:"headerReference"`
	FooterReference string `xml:"footerReference"`
}
type EG_SectPrContents struct {
	FootnotePr      string         `xml:"footnotePr"`
	EndnotePr       string         `xml:"endnotePr"`
	Type            CT_SectType    `xml:"type"`
	PgSz            CT_PageSz      `xml:"pgSz"`
	PgMar           CT_PageMar     `xml:"pgMar"`
	PaperSrc        CT_PaperSource `xml:"paperSrc"`
	PgBorders       string         `xml:"pgBorders"`
	LnNumType       CT_LineNumber  `xml:"lnNumType"`
	PgNumType       CT_PageNumber  `xml:"pgNumType"`
	Cols            string         `xml:"cols"`
	FormProt        CT_OnOff       `xml:"formProt"`
	VAlign          CT_VerticalJc  `xml:"vAlign"`
	NoEndnote       CT_OnOff       `xml:"noEndnote"`
	TitlePg         CT_OnOff       `xml:"titlePg"`
	TextDirection   string         `xml:"textDirection"`
	Bidi            CT_OnOff       `xml:"bidi"`
	RtlGutter       CT_OnOff       `xml:"rtlGutter"`
	DocGrid         string         `xml:"docGrid"`
	PrinterSettings CT_Rel         `xml:"printerSettings"`
}
type EG_RunInnerContent struct {
	Br                    CT_Br      `xml:"br"`
	T                     string     `xml:"t"`
	ContentPart           CT_Rel     `xml:"contentPart"`
	DelText               string     `xml:"delText"`
	InstrText             string     `xml:"instrText"`
	DelInstrText          string     `xml:"delInstrText"`
	NoBreakHyphen         CT_Empty   `xml:"noBreakHyphen"`
	SoftHyphen            CT_Empty   `xml:"softHyphen"`
	DayShort              CT_Empty   `xml:"dayShort"`
	MonthShort            CT_Empty   `xml:"monthShort"`
	YearShort             CT_Empty   `xml:"yearShort"`
	DayLong               CT_Empty   `xml:"dayLong"`
	MonthLong             CT_Empty   `xml:"monthLong"`
	YearLong              CT_Empty   `xml:"yearLong"`
	AnnotationRef         CT_Empty   `xml:"annotationRef"`
	FootnoteRef           CT_Empty   `xml:"footnoteRef"`
	EndnoteRef            CT_Empty   `xml:"endnoteRef"`
	Separator             CT_Empty   `xml:"separator"`
	ContinuationSeparator CT_Empty   `xml:"continuationSeparator"`
	Sym                   CT_Sym     `xml:"sym"`
	PgNum                 CT_Empty   `xml:"pgNum"`
	Cr                    CT_Empty   `xml:"cr"`
	Tab                   CT_Empty   `xml:"tab"`
	Object                CT_Object  `xml:"object"`
	Pict                  CT_Picture `xml:"pict"`
	FldChar               string     `xml:"fldChar"`
	Ruby                  CT_Ruby    `xml:"ruby"`
	FootnoteReference     string     `xml:"footnoteReference"`
	EndnoteReference      string     `xml:"endnoteReference"`
	CommentReference      CT_Markup  `xml:"commentReference"`
	Drawing               CT_Drawing `xml:"drawing"`
	Ptab                  CT_PTab    `xml:"ptab"`
	LastRenderedPageBreak CT_Empty   `xml:"lastRenderedPageBreak"`
}
type EG_RPrBase struct {
	RStyle          CT_String           `xml:"rStyle"`
	RFonts          string              `xml:"rFonts"`
	B               CT_OnOff            `xml:"b"`
	BCs             CT_OnOff            `xml:"bCs"`
	I               CT_OnOff            `xml:"i"`
	ICs             CT_OnOff            `xml:"iCs"`
	Caps            CT_OnOff            `xml:"caps"`
	SmallCaps       CT_OnOff            `xml:"smallCaps"`
	Strike          CT_OnOff            `xml:"strike"`
	Dstrike         CT_OnOff            `xml:"dstrike"`
	Outline         CT_OnOff            `xml:"outline"`
	Shadow          CT_OnOff            `xml:"shadow"`
	Emboss          CT_OnOff            `xml:"emboss"`
	Imprint         CT_OnOff            `xml:"imprint"`
	NoProof         CT_OnOff            `xml:"noProof"`
	SnapToGrid      CT_OnOff            `xml:"snapToGrid"`
	Vanish          CT_OnOff            `xml:"vanish"`
	WebHidden       CT_OnOff            `xml:"webHidden"`
	Color           CT_Color            `xml:"color"`
	Spacing         string              `xml:"spacing"`
	W               string              `xml:"w"`
	Kern            string              `xml:"kern"`
	Position        string              `xml:"position"`
	Sz              string              `xml:"sz"`
	SzCs            string              `xml:"szCs"`
	Highlight       CT_Highlight        `xml:"highlight"`
	U               string              `xml:"u"`
	Effect          string              `xml:"effect"`
	Bdr             string              `xml:"bdr"`
	Shd             string              `xml:"shd"`
	FitText         string              `xml:"fitText"`
	VertAlign       CT_VerticalAlignRun `xml:"vertAlign"`
	Rtl             CT_OnOff            `xml:"rtl"`
	Cs              CT_OnOff            `xml:"cs"`
	Em              CT_Em               `xml:"em"`
	Lang            CT_Language         `xml:"lang"`
	EastAsianLayout string              `xml:"eastAsianLayout"`
	SpecVanish      CT_OnOff            `xml:"specVanish"`
	OMath           CT_OnOff            `xml:"oMath"`
}
type EG_RPrContent struct {
	RPrChange CT_RPrChange `xml:"rPrChange"`
	RPrBase   []EG_RPrBase
}
type EG_RPr struct {
	RPr CT_RPr `xml:"rPr"`
}
type EG_RPrMath struct {
	RPr EG_RPr
	Ins string         `xml:"ins"`
	Del CT_MathCtrlDel `xml:"del"`
}
type EG_ParaRPrTrackChanges struct {
	Ins      CT_TrackChange `xml:"ins"`
	Del      CT_TrackChange `xml:"del"`
	MoveFrom CT_TrackChange `xml:"moveFrom"`
	MoveTo   CT_TrackChange `xml:"moveTo"`
}
type EG_RubyContent struct {
	RunLevelElts []EG_RunLevelElts
	R            CT_R `xml:"r"`
}
type EG_ContentRunContent struct {
	RunLevelElts []EG_RunLevelElts
	CustomXml    string           `xml:"customXml"`
	SmartTag     CT_SmartTagRun   `xml:"smartTag"`
	Sdt          string           `xml:"sdt"`
	Dir          CT_DirContentRun `xml:"dir"`
	Bdo          string           `xml:"bdo"`
	R            CT_R             `xml:"r"`
}
type EG_ContentBlockContent struct {
	RunLevelElts []EG_RunLevelElts
	CustomXml    string   `xml:"customXml"`
	Sdt          string   `xml:"sdt"`
	P            []CT_P   `xml:"p"`
	Tbl          []CT_Tbl `xml:"tbl"`
}
type EG_ContentRowContent struct {
	RunLevelElts []EG_RunLevelElts
	Tr           []CT_Row `xml:"tr"`
	CustomXml    string   `xml:"customXml"`
	Sdt          string   `xml:"sdt"`
}
type EG_ContentCellContent struct {
	RunLevelElts []EG_RunLevelElts
	Tc           []CT_Tc `xml:"tc"`
	CustomXml    string  `xml:"customXml"`
	Sdt          string  `xml:"sdt"`
}
type EG_PContent struct {
	ContentRunContent []EG_ContentRunContent
	FldSimple         []string     `xml:"fldSimple"`
	Hyperlink         CT_Hyperlink `xml:"hyperlink"`
	SubDoc            CT_Rel       `xml:"subDoc"`
}
type EG_FtnEdnNumProps struct {
	NumStart   CT_DecimalNumber `xml:"numStart"`
	NumRestart string           `xml:"numRestart"`
}
type EG_MathContent struct {
}
type EG_BlockLevelChunkElts struct {
	ContentBlockContent []EG_ContentBlockContent
}
type EG_BlockLevelElts struct {
	BlockLevelChunkElts []EG_BlockLevelChunkElts
	AltChunk            []CT_AltChunk `xml:"altChunk"`
}
type EG_RunLevelElts struct {
	RangeMarkupElements []EG_RangeMarkupElements
	MathContent         []EG_MathContent
	ProofErr            CT_ProofErr       `xml:"proofErr"`
	PermStart           CT_PermStart      `xml:"permStart"`
	PermEnd             CT_Perm           `xml:"permEnd"`
	Ins                 CT_RunTrackChange `xml:"ins"`
	Del                 CT_RunTrackChange `xml:"del"`
	MoveFrom            CT_RunTrackChange `xml:"moveFrom"`
	MoveTo              CT_RunTrackChange `xml:"moveTo"`
}
type CT_Empty struct {
}
type CT_OnOff struct {
	Val string `xml:"val",attr`
}
type CT_LongHexNumber struct {
	Val string `xml:"val",attr`
}
type CT_Charset struct {
	Val          string `xml:"val",attr`
	CharacterSet string `xml:"characterSet",attr`
}
type CT_DecimalNumber struct {
	Val ST_DecimalNumber `xml:"val",attr`
}
type CT_UnsignedDecimalNumber struct {
	Val string `xml:"val",attr`
}
type CT_DecimalNumberOrPrecent struct {
	Val ST_DecimalNumberOrPercent `xml:"val",attr`
}
type CT_TwipsMeasure struct {
	Val string `xml:"val",attr`
}
type CT_SignedTwipsMeasure struct {
	Val string `xml:"val",attr`
}
type CT_PixelsMeasure struct {
	Val string `xml:"val",attr`
}
type CT_HpsMeasure struct {
	Val string `xml:"val",attr`
}
type CT_SignedHpsMeasure struct {
	Val string `xml:"val",attr`
}
type CT_MacroName struct {
	Val ST_MacroName `xml:"val",attr`
}
type CT_String struct {
	Val string `xml:"val",attr`
}
type CT_TextScale struct {
	Val string `xml:"val",attr`
}
type CT_Highlight struct {
	Val ST_HighlightColor `xml:"val",attr`
}
type CT_Color struct {
	Val        string        `xml:"val",attr`
	ThemeColor ST_ThemeColor `xml:"themeColor",attr`
	ThemeTint  string        `xml:"themeTint",attr`
	ThemeShade string        `xml:"themeShade",attr`
}
type CT_Lang struct {
	Val string `xml:"val",attr`
}
type CT_Guid struct {
	Val string `xml:"val",attr`
}
type CT_Underline struct {
	Val        string        `xml:"val",attr`
	Color      string        `xml:"color",attr`
	ThemeColor ST_ThemeColor `xml:"themeColor",attr`
	ThemeTint  string        `xml:"themeTint",attr`
	ThemeShade string        `xml:"themeShade",attr`
}
type CT_TextEffect struct {
	Val string `xml:"val",attr`
}
type CT_Border struct {
	Val        string        `xml:"val",attr`
	Color      string        `xml:"color",attr`
	ThemeColor ST_ThemeColor `xml:"themeColor",attr`
	ThemeTint  string        `xml:"themeTint",attr`
	ThemeShade string        `xml:"themeShade",attr`
	Sz         string        `xml:"sz",attr`
	Space      string        `xml:"space",attr`
	Shadow     string        `xml:"shadow",attr`
	Frame      string        `xml:"frame",attr`
}
type CT_Shd struct {
	Val            string        `xml:"val",attr`
	Color          string        `xml:"color",attr`
	ThemeColor     ST_ThemeColor `xml:"themeColor",attr`
	ThemeTint      string        `xml:"themeTint",attr`
	ThemeShade     string        `xml:"themeShade",attr`
	Fill           string        `xml:"fill",attr`
	ThemeFill      ST_ThemeColor `xml:"themeFill",attr`
	ThemeFillTint  string        `xml:"themeFillTint",attr`
	ThemeFillShade string        `xml:"themeFillShade",attr`
}
type CT_VerticalAlignRun struct {
	Val string `xml:"val",attr`
}
type CT_FitText struct {
	Val string           `xml:"val",attr`
	Id  ST_DecimalNumber `xml:"id",attr`
}
type CT_Em struct {
	Val ST_Em `xml:"val",attr`
}
type CT_Language struct {
	Val      string `xml:"val",attr`
	EastAsia string `xml:"eastAsia",attr`
	Bidi     string `xml:"bidi",attr`
}
type CT_EastAsianLayout struct {
	Id              ST_DecimalNumber `xml:"id",attr`
	Combine         string           `xml:"combine",attr`
	CombineBrackets string           `xml:"combineBrackets",attr`
	Vert            string           `xml:"vert",attr`
	VertCompress    string           `xml:"vertCompress",attr`
}
type CT_FramePr struct {
	DropCap    ST_DropCap       `xml:"dropCap",attr`
	Lines      ST_DecimalNumber `xml:"lines",attr`
	W          string           `xml:"w",attr`
	H          string           `xml:"h",attr`
	VSpace     string           `xml:"vSpace",attr`
	HSpace     string           `xml:"hSpace",attr`
	Wrap       ST_Wrap          `xml:"wrap",attr`
	HAnchor    ST_HAnchor       `xml:"hAnchor",attr`
	VAnchor    ST_VAnchor       `xml:"vAnchor",attr`
	X          string           `xml:"x",attr`
	XAlign     string           `xml:"xAlign",attr`
	Y          string           `xml:"y",attr`
	YAlign     string           `xml:"yAlign",attr`
	HRule      ST_HeightRule    `xml:"hRule",attr`
	AnchorLock string           `xml:"anchorLock",attr`
}
type CT_TabStop struct {
	Val    ST_TabJc  `xml:"val",attr`
	Leader ST_TabTlc `xml:"leader",attr`
	Pos    string    `xml:"pos",attr`
}
type CT_Spacing struct {
	Before            string             `xml:"before",attr`
	BeforeLines       ST_DecimalNumber   `xml:"beforeLines",attr`
	BeforeAutospacing string             `xml:"beforeAutospacing",attr`
	After             string             `xml:"after",attr`
	AfterLines        ST_DecimalNumber   `xml:"afterLines",attr`
	AfterAutospacing  string             `xml:"afterAutospacing",attr`
	Line              string             `xml:"line",attr`
	LineRule          ST_LineSpacingRule `xml:"lineRule",attr`
}
type CT_Ind struct {
	Start          string           `xml:"start",attr`
	StartChars     ST_DecimalNumber `xml:"startChars",attr`
	End            string           `xml:"end",attr`
	EndChars       ST_DecimalNumber `xml:"endChars",attr`
	Left           string           `xml:"left",attr`
	LeftChars      ST_DecimalNumber `xml:"leftChars",attr`
	Right          string           `xml:"right",attr`
	RightChars     ST_DecimalNumber `xml:"rightChars",attr`
	Hanging        string           `xml:"hanging",attr`
	HangingChars   ST_DecimalNumber `xml:"hangingChars",attr`
	FirstLine      string           `xml:"firstLine",attr`
	FirstLineChars ST_DecimalNumber `xml:"firstLineChars",attr`
}
type CT_Jc struct {
	Val ST_Jc `xml:"val",attr`
}
type CT_JcTable struct {
	Val ST_JcTable `xml:"val",attr`
}
type CT_View struct {
	Val ST_View `xml:"val",attr`
}
type CT_Zoom struct {
	Val     ST_Zoom                   `xml:"val",attr`
	Percent ST_DecimalNumberOrPercent `xml:"percent",attr`
}
type CT_WritingStyle struct {
	Lang       string `xml:"lang",attr`
	VendorID   string `xml:"vendorID",attr`
	DllVersion string `xml:"dllVersion",attr`
	NlCheck    string `xml:"nlCheck",attr`
	CheckStyle string `xml:"checkStyle",attr`
	AppName    string `xml:"appName",attr`
}
type CT_Proof struct {
	Spelling ST_Proof `xml:"spelling",attr`
	Grammar  ST_Proof `xml:"grammar",attr`
}
type CT_DocType struct {
	Val ST_DocType `xml:"val",attr`
}
type CT_DocProtect struct {
	Edit        ST_DocProtect `xml:"edit",attr`
	Formatting  string        `xml:"formatting",attr`
	Enforcement string        `xml:"enforcement",attr`
}
type CT_MailMergeDocType struct {
	Val ST_MailMergeDocType `xml:"val",attr`
}
type CT_MailMergeDataType struct {
	Val ST_MailMergeDataType `xml:"val",attr`
}
type CT_MailMergeDest struct {
	Val string `xml:"val",attr`
}
type CT_MailMergeOdsoFMDFieldType struct {
	Val string `xml:"val",attr`
}
type CT_TrackChangesView struct {
	Markup         string `xml:"markup",attr`
	Comments       string `xml:"comments",attr`
	InsDel         string `xml:"insDel",attr`
	Formatting     string `xml:"formatting",attr`
	InkAnnotations string `xml:"inkAnnotations",attr`
}
type CT_Kinsoku struct {
	Lang string `xml:"lang",attr`
	Val  string `xml:"val",attr`
}
type CT_TextDirection struct {
	Val string `xml:"val",attr`
}
type CT_TextAlignment struct {
	Val string `xml:"val",attr`
}
type CT_Markup struct {
	Id ST_DecimalNumber `xml:"id",attr`
}
type CT_TrackChange struct {
	CT_Markup
	Author string      `xml:"author",attr`
	Date   ST_DateTime `xml:"date",attr`
}
type CT_CellMergeTrackChange struct {
	CT_TrackChange
	VMerge     ST_AnnotationVMerge `xml:"vMerge",attr`
	VMergeOrig ST_AnnotationVMerge `xml:"vMergeOrig",attr`
}
type CT_TrackChangeRange struct {
	CT_TrackChange
	DisplacedByCustomXml string `xml:"displacedByCustomXml",attr`
}
type CT_MarkupRange struct {
	CT_Markup
	DisplacedByCustomXml string `xml:"displacedByCustomXml",attr`
}
type CT_BookmarkRange struct {
	CT_MarkupRange
	ColFirst ST_DecimalNumber `xml:"colFirst",attr`
	ColLast  ST_DecimalNumber `xml:"colLast",attr`
}
type CT_Bookmark struct {
	CT_BookmarkRange
	Name string `xml:"name",attr`
}
type CT_MoveBookmark struct {
	CT_Bookmark
	Author string      `xml:"author",attr`
	Date   ST_DateTime `xml:"date",attr`
}
type CT_Comment struct {
	CT_TrackChange
	BlockLevelElts []EG_BlockLevelElts
	Initials       string `xml:"initials",attr`
}
type CT_TrackChangeNumbering struct {
	CT_TrackChange
	Original string `xml:"original",attr`
}
type CT_TblPrExChange struct {
	CT_TrackChange
	TblPrEx string `xml:"tblPrEx"`
}
type CT_TcPrChange struct {
	CT_TrackChange
	TcPr CT_TcPrInner `xml:"tcPr"`
}
type CT_TrPrChange struct {
	CT_TrackChange
	TrPr string `xml:"trPr"`
}
type CT_TblGridChange struct {
	CT_Markup
	TblGrid string `xml:"tblGrid"`
}
type CT_TblPrChange struct {
	CT_TrackChange
	TblPr string `xml:"tblPr"`
}
type CT_SectPrChange struct {
	CT_TrackChange
	SectPr string `xml:"sectPr"`
}
type CT_PPrChange struct {
	CT_TrackChange
	PPr string `xml:"pPr"`
}
type CT_RPrChange struct {
	CT_TrackChange
	RPr CT_RPrOriginal `xml:"rPr"`
}
type CT_ParaRPrChange struct {
	CT_TrackChange
	RPr CT_ParaRPrOriginal `xml:"rPr"`
}
type CT_RunTrackChange struct {
	CT_TrackChange
	ContentRunContent EG_ContentRunContent
	//	OMathMathElements EG_OMathMathElements
}
type CT_NumPr struct {
	Ilvl            CT_DecimalNumber        `xml:"ilvl"`
	NumId           CT_DecimalNumber        `xml:"numId"`
	NumberingChange CT_TrackChangeNumbering `xml:"numberingChange"`
	Ins             CT_TrackChange          `xml:"ins"`
}
type CT_PBdr struct {
	Top     string `xml:"top"`
	Left    string `xml:"left"`
	Bottom  string `xml:"bottom"`
	Right   string `xml:"right"`
	Between string `xml:"between"`
	Bar     string `xml:"bar"`
}
type CT_Tabs struct {
	Tab []CT_TabStop `xml:"tab"`
}
type CT_TextboxTightWrap struct {
	Val string `xml:"val",attr`
}
type CT_PPr struct {
	CT_PPrBase
	RPr       CT_ParaRPr   `xml:"rPr"`
	SectPr    CT_SectPr    `xml:"sectPr"`
	PPrChange CT_PPrChange `xml:"pPrChange"`
}
type CT_PPrBase struct {
	PStyle              CT_String        `xml:"pStyle"`
	KeepNext            CT_OnOff         `xml:"keepNext"`
	KeepLines           CT_OnOff         `xml:"keepLines"`
	PageBreakBefore     CT_OnOff         `xml:"pageBreakBefore"`
	FramePr             CT_FramePr       `xml:"framePr"`
	WidowControl        CT_OnOff         `xml:"widowControl"`
	NumPr               CT_NumPr         `xml:"numPr"`
	SuppressLineNumbers CT_OnOff         `xml:"suppressLineNumbers"`
	PBdr                string           `xml:"pBdr"`
	Shd                 string           `xml:"shd"`
	Tabs                string           `xml:"tabs"`
	SuppressAutoHyphens CT_OnOff         `xml:"suppressAutoHyphens"`
	Kinsoku             CT_OnOff         `xml:"kinsoku"`
	WordWrap            CT_OnOff         `xml:"wordWrap"`
	OverflowPunct       CT_OnOff         `xml:"overflowPunct"`
	TopLinePunct        CT_OnOff         `xml:"topLinePunct"`
	AutoSpaceDE         CT_OnOff         `xml:"autoSpaceDE"`
	AutoSpaceDN         CT_OnOff         `xml:"autoSpaceDN"`
	Bidi                CT_OnOff         `xml:"bidi"`
	AdjustRightInd      CT_OnOff         `xml:"adjustRightInd"`
	SnapToGrid          CT_OnOff         `xml:"snapToGrid"`
	Spacing             CT_Spacing       `xml:"spacing"`
	Ind                 string           `xml:"ind"`
	ContextualSpacing   CT_OnOff         `xml:"contextualSpacing"`
	MirrorIndents       CT_OnOff         `xml:"mirrorIndents"`
	SuppressOverlap     CT_OnOff         `xml:"suppressOverlap"`
	Jc                  CT_Jc            `xml:"jc"`
	TextDirection       string           `xml:"textDirection"`
	TextAlignment       string           `xml:"textAlignment"`
	TextboxTightWrap    string           `xml:"textboxTightWrap"`
	OutlineLvl          CT_DecimalNumber `xml:"outlineLvl"`
	DivId               CT_DecimalNumber `xml:"divId"`
	CnfStyle            CT_Cnf           `xml:"cnfStyle"`
}
type CT_PPrGeneral struct {
	CT_PPrBase
	PPrChange CT_PPrChange `xml:"pPrChange"`
}
type CT_Control struct {
	Name    string `xml:"name",attr`
	Shapeid string `xml:"shapeid",attr`
}
type CT_Background struct {
	Drawing    CT_Drawing    `xml:"drawing"`
	Color      string        `xml:"color",attr`
	ThemeColor ST_ThemeColor `xml:"themeColor",attr`
	ThemeTint  string        `xml:"themeTint",attr`
	ThemeShade string        `xml:"themeShade",attr`
}
type CT_Rel struct {
}
type CT_Object struct {
	Drawing     CT_Drawing    `xml:"drawing"`
	Control     CT_Control    `xml:"control"`
	ObjectLink  CT_ObjectLink `xml:"objectLink"`
	ObjectEmbed string        `xml:"objectEmbed"`
	Movie       CT_Rel        `xml:"movie"`
	DxaOrig     string        `xml:"dxaOrig",attr`
	DyaOrig     string        `xml:"dyaOrig",attr`
}
type CT_Picture struct {
	Movie   CT_Rel     `xml:"movie"`
	Control CT_Control `xml:"control"`
}
type CT_ObjectEmbed struct {
	DrawAspect string `xml:"drawAspect",attr`
	ProgId     string `xml:"progId",attr`
	ShapeId    string `xml:"shapeId",attr`
	FieldCodes string `xml:"fieldCodes",attr`
}
type CT_ObjectLink struct {
	CT_ObjectEmbed
	UpdateMode  string `xml:"updateMode",attr`
	LockedField string `xml:"lockedField",attr`
}
type CT_Drawing struct {
}
type CT_SimpleField struct {
	FldData  string `xml:"fldData"`
	PContent []EG_PContent
	Instr    string `xml:"instr",attr`
	FldLock  string `xml:"fldLock",attr`
	Dirty    string `xml:"dirty",attr`
}
type CT_FFTextType struct {
	Val string `xml:"val",attr`
}
type CT_FFName struct {
	Val ST_FFName `xml:"val",attr`
}
type CT_FldChar struct {
	FldData         string                  `xml:"fldData"`
	FfData          CT_FFData               `xml:"ffData"`
	NumberingChange CT_TrackChangeNumbering `xml:"numberingChange"`
	FldCharType     string                  `xml:"fldCharType",attr`
	FldLock         string                  `xml:"fldLock",attr`
	Dirty           string                  `xml:"dirty",attr`
}
type CT_Hyperlink struct {
	TgtFrame    string `xml:"tgtFrame",attr`
	Tooltip     string `xml:"tooltip",attr`
	DocLocation string `xml:"docLocation",attr`
	History     string `xml:"history",attr`
	Anchor      string `xml:"anchor",attr`
}
type CT_FFData struct {
	Name       CT_FFName        `xml:"name"`
	Label      CT_DecimalNumber `xml:"label"`
	TabIndex   string           `xml:"tabIndex"`
	Enabled    CT_OnOff         `xml:"enabled"`
	CalcOnExit CT_OnOff         `xml:"calcOnExit"`
	EntryMacro CT_MacroName     `xml:"entryMacro"`
	ExitMacro  CT_MacroName     `xml:"exitMacro"`
	HelpText   string           `xml:"helpText"`
	StatusText string           `xml:"statusText"`
}
type CT_FFHelpText struct {
	Type string `xml:"type",attr`
	Val  string `xml:"val",attr`
}
type CT_FFStatusText struct {
	Type string `xml:"type",attr`
	Val  string `xml:"val",attr`
}
type CT_FFCheckBox struct {
	Default  CT_OnOff `xml:"default"`
	Checked  CT_OnOff `xml:"checked"`
	Size     string   `xml:"size"`
	SizeAuto CT_OnOff `xml:"sizeAuto"`
}
type CT_FFDDList struct {
	Result    CT_DecimalNumber `xml:"result"`
	Default   CT_DecimalNumber `xml:"default"`
	ListEntry []CT_String      `xml:"listEntry"`
}
type CT_FFTextInput struct {
	Type      string           `xml:"type"`
	Default   CT_String        `xml:"default"`
	MaxLength CT_DecimalNumber `xml:"maxLength"`
	Format    CT_String        `xml:"format"`
}
type CT_SectType struct {
	Val ST_SectionMark `xml:"val",attr`
}
type CT_PaperSource struct {
	First ST_DecimalNumber `xml:"first",attr`
	Other ST_DecimalNumber `xml:"other",attr`
}
type CT_PageSz struct {
	W      string             `xml:"w",attr`
	H      string             `xml:"h",attr`
	Orient ST_PageOrientation `xml:"orient",attr`
	Code   ST_DecimalNumber   `xml:"code",attr`
}
type CT_PageMar struct {
	Top    string `xml:"top",attr`
	Right  string `xml:"right",attr`
	Bottom string `xml:"bottom",attr`
	Left   string `xml:"left",attr`
	Header string `xml:"header",attr`
	Footer string `xml:"footer",attr`
	Gutter string `xml:"gutter",attr`
}
type CT_PageBorders struct {
	Top        string `xml:"top"`
	Left       string `xml:"left"`
	Bottom     string `xml:"bottom"`
	Right      string `xml:"right"`
	ZOrder     string `xml:"zOrder",attr`
	Display    string `xml:"display",attr`
	OffsetFrom string `xml:"offsetFrom",attr`
}
type CT_PageBorder struct {
	CT_Border
}
type CT_BottomPageBorder struct {
	CT_PageBorder
}
type CT_TopPageBorder struct {
	CT_PageBorder
}
type CT_LineNumber struct {
	CountBy  ST_DecimalNumber `xml:"countBy",attr`
	Start    ST_DecimalNumber `xml:"start",attr`
	Distance string           `xml:"distance",attr`
	Restart  string           `xml:"restart",attr`
}
type CT_PageNumber struct {
	Fmt       ST_NumberFormat  `xml:"fmt",attr`
	Start     ST_DecimalNumber `xml:"start",attr`
	ChapStyle ST_DecimalNumber `xml:"chapStyle",attr`
	ChapSep   ST_ChapterSep    `xml:"chapSep",attr`
}
type CT_Column struct {
	W     string `xml:"w",attr`
	Space string `xml:"space",attr`
}
type CT_Columns struct {
	Col        CT_Column        `xml:"col"`
	EqualWidth string           `xml:"equalWidth",attr`
	Space      string           `xml:"space",attr`
	Num        ST_DecimalNumber `xml:"num",attr`
	Sep        string           `xml:"sep",attr`
}
type CT_VerticalJc struct {
	Val ST_VerticalJc `xml:"val",attr`
}
type CT_DocGrid struct {
	Type      string           `xml:"type",attr`
	LinePitch ST_DecimalNumber `xml:"linePitch",attr`
	CharSpace ST_DecimalNumber `xml:"charSpace",attr`
}
type CT_HdrFtrRef struct {
	CT_Rel
	Type string `xml:"type",attr`
}
type CT_HdrFtr struct {
}
type CT_SectPrBase struct {
	SectPrContents EG_SectPrContents
}
type CT_SectPr struct {
	SectPrChange     CT_SectPrChange `xml:"sectPrChange"`
	HdrFtrReferences EG_HdrFtrReferences
	SectPrContents   EG_SectPrContents
}
type CT_Br struct {
	Type  ST_BrType  `xml:"type",attr`
	Clear ST_BrClear `xml:"clear",attr`
}
type CT_PTab struct {
	Alignment  ST_PTabAlignment  `xml:"alignment",attr`
	RelativeTo ST_PTabRelativeTo `xml:"relativeTo",attr`
	Leader     string            `xml:"leader",attr`
}
type CT_Sym struct {
	Font string `xml:"font",attr`
	Char string `xml:"char",attr`
}
type CT_ProofErr struct {
	Type ST_ProofErr `xml:"type",attr`
}
type CT_Perm struct {
	Id                   string `xml:"id",attr`
	DisplacedByCustomXml string `xml:"displacedByCustomXml",attr`
}
type CT_PermStart struct {
	CT_Perm
	EdGrp    string           `xml:"edGrp",attr`
	Ed       string           `xml:"ed",attr`
	ColFirst ST_DecimalNumber `xml:"colFirst",attr`
	ColLast  ST_DecimalNumber `xml:"colLast",attr`
}
type CT_Text struct {
}
type CT_R struct {
	RPr             EG_RPr
	RunInnerContent []EG_RunInnerContent
	RsidRPr         string `xml:"rsidRPr",attr`
	RsidDel         string `xml:"rsidDel",attr`
	RsidR           string `xml:"rsidR",attr`
}
type CT_Fonts struct {
	Hint          ST_Hint  `xml:"hint",attr`
	Ascii         string   `xml:"ascii",attr`
	HAnsi         string   `xml:"hAnsi",attr`
	EastAsia      string   `xml:"eastAsia",attr`
	Cs            string   `xml:"cs",attr`
	AsciiTheme    ST_Theme `xml:"asciiTheme",attr`
	HAnsiTheme    ST_Theme `xml:"hAnsiTheme",attr`
	EastAsiaTheme ST_Theme `xml:"eastAsiaTheme",attr`
	Cstheme       ST_Theme `xml:"cstheme",attr`
}
type CT_RPr struct {
	RPrContent EG_RPrContent
}
type CT_MathCtrlIns struct {
	CT_TrackChange
	Del CT_RPrChange `xml:"del"`
	RPr CT_RPr       `xml:"rPr"`
}
type CT_MathCtrlDel struct {
	CT_TrackChange
	RPr CT_RPr `xml:"rPr"`
}
type CT_RPrOriginal struct {
	RPrBase []EG_RPrBase
}
type CT_ParaRPrOriginal struct {
	ParaRPrTrackChanges EG_ParaRPrTrackChanges
	RPrBase             []EG_RPrBase
}
type CT_ParaRPr struct {
	RPrChange           CT_ParaRPrChange `xml:"rPrChange"`
	ParaRPrTrackChanges EG_ParaRPrTrackChanges
	RPrBase             []EG_RPrBase
}
type CT_AltChunk struct {
	AltChunkPr CT_AltChunkPr `xml:"altChunkPr"`
}
type CT_AltChunkPr struct {
	MatchSrc CT_OnOff `xml:"matchSrc"`
}
type CT_RubyAlign struct {
	Val ST_RubyAlign `xml:"val",attr`
}
type CT_RubyPr struct {
	RubyAlign   CT_RubyAlign `xml:"rubyAlign"`
	Hps         string       `xml:"hps"`
	HpsRaise    string       `xml:"hpsRaise"`
	HpsBaseText string       `xml:"hpsBaseText"`
	Lid         CT_Lang      `xml:"lid"`
	Dirty       CT_OnOff     `xml:"dirty"`
}
type CT_RubyContent struct {
}
type CT_Ruby struct {
	RubyPr   CT_RubyPr      `xml:"rubyPr"`
	Rt       CT_RubyContent `xml:"rt"`
	RubyBase CT_RubyContent `xml:"rubyBase"`
}
type CT_Lock struct {
	Val ST_Lock `xml:"val",attr`
}
type CT_SdtListItem struct {
	DisplayText string `xml:"displayText",attr`
	Value       string `xml:"value",attr`
}
type CT_SdtDateMappingType struct {
	Val string `xml:"val",attr`
}
type CT_CalendarType struct {
	Val string `xml:"val",attr`
}
type CT_SdtDate struct {
	DateFormat        CT_String   `xml:"dateFormat"`
	Lid               CT_Lang     `xml:"lid"`
	StoreMappedDataAs string      `xml:"storeMappedDataAs"`
	Calendar          string      `xml:"calendar"`
	FullDate          ST_DateTime `xml:"fullDate",attr`
}
type CT_SdtComboBox struct {
	ListItem  []string `xml:"listItem"`
	LastValue string   `xml:"lastValue",attr`
}
type CT_SdtDocPart struct {
	DocPartGallery  CT_String `xml:"docPartGallery"`
	DocPartCategory CT_String `xml:"docPartCategory"`
	DocPartUnique   CT_OnOff  `xml:"docPartUnique"`
}
type CT_SdtDropDownList struct {
	ListItem  []string `xml:"listItem"`
	LastValue string   `xml:"lastValue",attr`
}
type CT_Placeholder struct {
	DocPart CT_String `xml:"docPart"`
}
type CT_SdtText struct {
	MultiLine string `xml:"multiLine",attr`
}
type CT_DataBinding struct {
	PrefixMappings string `xml:"prefixMappings",attr`
	Xpath          string `xml:"xpath",attr`
	StoreItemID    string `xml:"storeItemID",attr`
}
type CT_SdtPr struct {
	RPr           CT_RPr           `xml:"rPr"`
	Alias         CT_String        `xml:"alias"`
	Tag           CT_String        `xml:"tag"`
	Id            CT_DecimalNumber `xml:"id"`
	Lock          CT_Lock          `xml:"lock"`
	Placeholder   string           `xml:"placeholder"`
	Temporary     CT_OnOff         `xml:"temporary"`
	ShowingPlcHdr CT_OnOff         `xml:"showingPlcHdr"`
	DataBinding   string           `xml:"dataBinding"`
	Label         CT_DecimalNumber `xml:"label"`
	TabIndex      string           `xml:"tabIndex"`
	Equation      CT_Empty         `xml:"equation"`
	ComboBox      string           `xml:"comboBox"`
	Date          string           `xml:"date"`
	DocPartObj    string           `xml:"docPartObj"`
	DocPartList   string           `xml:"docPartList"`
	DropDownList  string           `xml:"dropDownList"`
	Picture       CT_Empty         `xml:"picture"`
	RichText      CT_Empty         `xml:"richText"`
	Text          string           `xml:"text"`
	Citation      CT_Empty         `xml:"citation"`
	Group         CT_Empty         `xml:"group"`
	Bibliography  CT_Empty         `xml:"bibliography"`
}
type CT_SdtEndPr struct {
	RPr CT_RPr `xml:"rPr"`
}
type CT_DirContentRun struct {
	Val ST_Direction `xml:"val",attr`
}
type CT_BdoContentRun struct {
	Val ST_Direction `xml:"val",attr`
}
type CT_SdtContentRun struct {
}
type CT_SdtContentBlock struct {
}
type CT_SdtContentRow struct {
}
type CT_SdtContentCell struct {
}
type CT_SdtBlock struct {
	SdtPr      string `xml:"sdtPr"`
	SdtEndPr   string `xml:"sdtEndPr"`
	SdtContent string `xml:"sdtContent"`
}
type CT_SdtRun struct {
	SdtPr      string `xml:"sdtPr"`
	SdtEndPr   string `xml:"sdtEndPr"`
	SdtContent string `xml:"sdtContent"`
}
type CT_SdtCell struct {
	SdtPr      string `xml:"sdtPr"`
	SdtEndPr   string `xml:"sdtEndPr"`
	SdtContent string `xml:"sdtContent"`
}
type CT_SdtRow struct {
	SdtPr      string `xml:"sdtPr"`
	SdtEndPr   string `xml:"sdtEndPr"`
	SdtContent string `xml:"sdtContent"`
}
type CT_Attr struct {
	Uri  string `xml:"uri",attr`
	Name string `xml:"name",attr`
	Val  string `xml:"val",attr`
}
type CT_CustomXmlRun struct {
	CustomXmlPr string `xml:"customXmlPr"`
	PContent    []EG_PContent
	Uri         string `xml:"uri",attr`
	Element     string `xml:"element",attr`
}
type CT_SmartTagRun struct {
	SmartTagPr CT_SmartTagPr `xml:"smartTagPr"`
	PContent   []EG_PContent
	Uri        string `xml:"uri",attr`
	Element    string `xml:"element",attr`
}
type CT_CustomXmlBlock struct {
	CustomXmlPr         string `xml:"customXmlPr"`
	ContentBlockContent []EG_ContentBlockContent
	Uri                 string `xml:"uri",attr`
	Element             string `xml:"element",attr`
}
type CT_CustomXmlPr struct {
	Placeholder CT_String `xml:"placeholder"`
	Attr        []CT_Attr `xml:"attr"`
}
type CT_CustomXmlRow struct {
	CustomXmlPr       string `xml:"customXmlPr"`
	ContentRowContent []EG_ContentRowContent
	Uri               string `xml:"uri",attr`
	Element           string `xml:"element",attr`
}
type CT_CustomXmlCell struct {
	CustomXmlPr        string `xml:"customXmlPr"`
	ContentCellContent []EG_ContentCellContent
	Uri                string `xml:"uri",attr`
	Element            string `xml:"element",attr`
}
type CT_SmartTagPr struct {
	Attr []CT_Attr `xml:"attr"`
}
type CT_P struct {
	PPr          CT_PPr `xml:"pPr"`
	PContent     []EG_PContent
	RsidRPr      string `xml:"rsidRPr",attr`
	RsidR        string `xml:"rsidR",attr`
	RsidDel      string `xml:"rsidDel",attr`
	RsidP        string `xml:"rsidP",attr`
	RsidRDefault string `xml:"rsidRDefault",attr`
}
type CT_Height struct {
	Val   string        `xml:"val",attr`
	HRule ST_HeightRule `xml:"hRule",attr`
}
type CT_TblWidth struct {
	W    string `xml:"w",attr`
	Type string `xml:"type",attr`
}
type CT_TblGridCol struct {
	W string `xml:"w",attr`
}
type CT_TblGridBase struct {
	GridCol []string `xml:"gridCol"`
}
type CT_TblGrid struct {
	CT_TblGridBase
	TblGridChange string `xml:"tblGridChange"`
}
type CT_TcBorders struct {
	Top     string `xml:"top"`
	Start   string `xml:"start"`
	Left    string `xml:"left"`
	Bottom  string `xml:"bottom"`
	End     string `xml:"end"`
	Right   string `xml:"right"`
	InsideH string `xml:"insideH"`
	InsideV string `xml:"insideV"`
	Tl2br   string `xml:"tl2br"`
	Tr2bl   string `xml:"tr2bl"`
}
type CT_TcMar struct {
	Top    string `xml:"top"`
	Start  string `xml:"start"`
	Left   string `xml:"left"`
	Bottom string `xml:"bottom"`
	End    string `xml:"end"`
	Right  string `xml:"right"`
}
type CT_VMerge struct {
	Val ST_Merge `xml:"val",attr`
}
type CT_HMerge struct {
	Val ST_Merge `xml:"val",attr`
}
type CT_TcPrBase struct {
	CnfStyle      CT_Cnf           `xml:"cnfStyle"`
	TcW           string           `xml:"tcW"`
	GridSpan      CT_DecimalNumber `xml:"gridSpan"`
	HMerge        CT_HMerge        `xml:"hMerge"`
	VMerge        CT_VMerge        `xml:"vMerge"`
	TcBorders     string           `xml:"tcBorders"`
	Shd           string           `xml:"shd"`
	NoWrap        CT_OnOff         `xml:"noWrap"`
	TcMar         CT_TcMar         `xml:"tcMar"`
	TextDirection string           `xml:"textDirection"`
	TcFitText     CT_OnOff         `xml:"tcFitText"`
	VAlign        CT_VerticalJc    `xml:"vAlign"`
	HideMark      CT_OnOff         `xml:"hideMark"`
	Headers       string           `xml:"headers"`
}
type CT_TcPr struct {
	CT_TcPrInner
	TcPrChange CT_TcPrChange `xml:"tcPrChange"`
}
type CT_TcPrInner struct {
	CT_TcPrBase
	CellMarkupElements EG_CellMarkupElements
}
type CT_Tc struct {
	TcPr           CT_TcPr `xml:"tcPr"`
	BlockLevelElts []EG_BlockLevelElts
	Id             string `xml:"id",attr`
}
type CT_Cnf struct {
	Val                 ST_Cnf `xml:"val",attr`
	FirstRow            string `xml:"firstRow",attr`
	LastRow             string `xml:"lastRow",attr`
	FirstColumn         string `xml:"firstColumn",attr`
	LastColumn          string `xml:"lastColumn",attr`
	OddVBand            string `xml:"oddVBand",attr`
	EvenVBand           string `xml:"evenVBand",attr`
	OddHBand            string `xml:"oddHBand",attr`
	EvenHBand           string `xml:"evenHBand",attr`
	FirstRowFirstColumn string `xml:"firstRowFirstColumn",attr`
	FirstRowLastColumn  string `xml:"firstRowLastColumn",attr`
	LastRowFirstColumn  string `xml:"lastRowFirstColumn",attr`
	LastRowLastColumn   string `xml:"lastRowLastColumn",attr`
}
type CT_Headers struct {
	Header CT_String `xml:"header"`
}
type CT_TrPrBase struct {
	CnfStyle       CT_Cnf           `xml:"cnfStyle"`
	DivId          CT_DecimalNumber `xml:"divId"`
	GridBefore     CT_DecimalNumber `xml:"gridBefore"`
	GridAfter      CT_DecimalNumber `xml:"gridAfter"`
	WBefore        string           `xml:"wBefore"`
	WAfter         string           `xml:"wAfter"`
	CantSplit      CT_OnOff         `xml:"cantSplit"`
	TrHeight       CT_Height        `xml:"trHeight"`
	TblHeader      CT_OnOff         `xml:"tblHeader"`
	TblCellSpacing string           `xml:"tblCellSpacing"`
	Jc             CT_JcTable       `xml:"jc"`
	Hidden         CT_OnOff         `xml:"hidden"`
}
type CT_TrPr struct {
	CT_TrPrBase
	Ins        CT_TrackChange `xml:"ins"`
	Del        CT_TrackChange `xml:"del"`
	TrPrChange CT_TrPrChange  `xml:"trPrChange"`
}
type CT_Row struct {
	TblPrEx            string  `xml:"tblPrEx"`
	TrPr               CT_TrPr `xml:"trPr"`
	ContentCellContent []EG_ContentCellContent
	RsidRPr            string `xml:"rsidRPr",attr`
	RsidR              string `xml:"rsidR",attr`
	RsidDel            string `xml:"rsidDel",attr`
	RsidTr             string `xml:"rsidTr",attr`
}
type CT_TblLayoutType struct {
	Type ST_TblLayoutType `xml:"type",attr`
}
type CT_TblOverlap struct {
	Val ST_TblOverlap `xml:"val",attr`
}
type CT_TblPPr struct {
	LeftFromText   string     `xml:"leftFromText",attr`
	RightFromText  string     `xml:"rightFromText",attr`
	TopFromText    string     `xml:"topFromText",attr`
	BottomFromText string     `xml:"bottomFromText",attr`
	VertAnchor     ST_VAnchor `xml:"vertAnchor",attr`
	HorzAnchor     ST_HAnchor `xml:"horzAnchor",attr`
	TblpXSpec      string     `xml:"tblpXSpec",attr`
	TblpX          string     `xml:"tblpX",attr`
	TblpYSpec      string     `xml:"tblpYSpec",attr`
	TblpY          string     `xml:"tblpY",attr`
}
type CT_TblCellMar struct {
	Top    string `xml:"top"`
	Start  string `xml:"start"`
	Left   string `xml:"left"`
	Bottom string `xml:"bottom"`
	End    string `xml:"end"`
	Right  string `xml:"right"`
}
type CT_TblBorders struct {
	Top     string `xml:"top"`
	Start   string `xml:"start"`
	Left    string `xml:"left"`
	Bottom  string `xml:"bottom"`
	End     string `xml:"end"`
	Right   string `xml:"right"`
	InsideH string `xml:"insideH"`
	InsideV string `xml:"insideV"`
}
type CT_TblPrBase struct {
	TblStyle            CT_String        `xml:"tblStyle"`
	TblpPr              CT_TblPPr        `xml:"tblpPr"`
	TblOverlap          CT_TblOverlap    `xml:"tblOverlap"`
	BidiVisual          CT_OnOff         `xml:"bidiVisual"`
	TblStyleRowBandSize CT_DecimalNumber `xml:"tblStyleRowBandSize"`
	TblStyleColBandSize CT_DecimalNumber `xml:"tblStyleColBandSize"`
	TblW                string           `xml:"tblW"`
	Jc                  CT_JcTable       `xml:"jc"`
	TblCellSpacing      string           `xml:"tblCellSpacing"`
	TblInd              string           `xml:"tblInd"`
	TblBorders          string           `xml:"tblBorders"`
	Shd                 string           `xml:"shd"`
	TblLayout           CT_TblLayoutType `xml:"tblLayout"`
	TblCellMar          CT_TblCellMar    `xml:"tblCellMar"`
	TblLook             CT_TblLook       `xml:"tblLook"`
	TblCaption          CT_String        `xml:"tblCaption"`
	TblDescription      CT_String        `xml:"tblDescription"`
}
type CT_TblPr struct {
	CT_TblPrBase
	TblPrChange CT_TblPrChange `xml:"tblPrChange"`
}
type CT_TblPrExBase struct {
	TblW           string           `xml:"tblW"`
	Jc             CT_JcTable       `xml:"jc"`
	TblCellSpacing string           `xml:"tblCellSpacing"`
	TblInd         string           `xml:"tblInd"`
	TblBorders     string           `xml:"tblBorders"`
	Shd            string           `xml:"shd"`
	TblLayout      CT_TblLayoutType `xml:"tblLayout"`
	TblCellMar     CT_TblCellMar    `xml:"tblCellMar"`
	TblLook        CT_TblLook       `xml:"tblLook"`
}
type CT_TblPrEx struct {
	CT_TblPrExBase
	TblPrExChange string `xml:"tblPrExChange"`
}
type CT_Tbl struct {
	TblPr               CT_TblPr `xml:"tblPr"`
	TblGrid             string   `xml:"tblGrid"`
	RangeMarkupElements []EG_RangeMarkupElements
	ContentRowContent   []EG_ContentRowContent
}
type CT_TblLook struct {
	FirstRow    string `xml:"firstRow",attr`
	LastRow     string `xml:"lastRow",attr`
	FirstColumn string `xml:"firstColumn",attr`
	LastColumn  string `xml:"lastColumn",attr`
	NoHBand     string `xml:"noHBand",attr`
	NoVBand     string `xml:"noVBand",attr`
	Val         string `xml:"val",attr`
}
type CT_FtnPos struct {
	Val string `xml:"val",attr`
}
type CT_EdnPos struct {
	Val string `xml:"val",attr`
}
type CT_NumFmt struct {
	Val    ST_NumberFormat `xml:"val",attr`
	Format string          `xml:"format",attr`
}
type CT_NumRestart struct {
	Val string `xml:"val",attr`
}
type CT_FtnEdnRef struct {
	CustomMarkFollows string           `xml:"customMarkFollows",attr`
	Id                ST_DecimalNumber `xml:"id",attr`
}
type CT_FtnEdnSepRef struct {
	Id ST_DecimalNumber `xml:"id",attr`
}
type CT_FtnEdn struct {
	BlockLevelElts []EG_BlockLevelElts
	Type           string           `xml:"type",attr`
	Id             ST_DecimalNumber `xml:"id",attr`
}
type CT_FtnProps struct {
	Pos            string    `xml:"pos"`
	NumFmt         CT_NumFmt `xml:"numFmt"`
	FtnEdnNumProps EG_FtnEdnNumProps
}
type CT_EdnProps struct {
	Pos            string    `xml:"pos"`
	NumFmt         CT_NumFmt `xml:"numFmt"`
	FtnEdnNumProps EG_FtnEdnNumProps
}
type CT_FtnDocProps struct {
	CT_FtnProps
	Footnote string `xml:"footnote"`
}
type CT_EdnDocProps struct {
	CT_EdnProps
	Endnote string `xml:"endnote"`
}
type CT_RecipientData struct {
	Active    CT_OnOff         `xml:"active"`
	Column    CT_DecimalNumber `xml:"column"`
	UniqueTag string           `xml:"uniqueTag"`
}
type CT_Base64Binary struct {
	Val string `xml:"val",attr`
}
type CT_Recipients struct {
	RecipientData []CT_RecipientData `xml:"recipientData"`
}
type CT_OdsoFieldMapData struct {
	Type           string           `xml:"type"`
	Name           CT_String        `xml:"name"`
	MappedName     CT_String        `xml:"mappedName"`
	Column         CT_DecimalNumber `xml:"column"`
	Lid            CT_Lang          `xml:"lid"`
	DynamicAddress CT_OnOff         `xml:"dynamicAddress"`
}
type CT_MailMergeSourceType struct {
	Val ST_MailMergeSourceType `xml:"val",attr`
}
type CT_Odso struct {
	Udl           CT_String              `xml:"udl"`
	Table         CT_String              `xml:"table"`
	Src           CT_Rel                 `xml:"src"`
	ColDelim      CT_DecimalNumber       `xml:"colDelim"`
	Type          CT_MailMergeSourceType `xml:"type"`
	FHdr          CT_OnOff               `xml:"fHdr"`
	FieldMapData  []string               `xml:"fieldMapData"`
	RecipientData []CT_Rel               `xml:"recipientData"`
}
type CT_MailMerge struct {
	MainDocumentType        CT_MailMergeDocType  `xml:"mainDocumentType"`
	LinkToQuery             CT_OnOff             `xml:"linkToQuery"`
	DataType                CT_MailMergeDataType `xml:"dataType"`
	ConnectString           CT_String            `xml:"connectString"`
	Query                   CT_String            `xml:"query"`
	DataSource              CT_Rel               `xml:"dataSource"`
	HeaderSource            CT_Rel               `xml:"headerSource"`
	DoNotSuppressBlankLines CT_OnOff             `xml:"doNotSuppressBlankLines"`
	Destination             string               `xml:"destination"`
	AddressFieldName        CT_String            `xml:"addressFieldName"`
	MailSubject             CT_String            `xml:"mailSubject"`
	MailAsAttachment        CT_OnOff             `xml:"mailAsAttachment"`
	ViewMergedData          CT_OnOff             `xml:"viewMergedData"`
	ActiveRecord            CT_DecimalNumber     `xml:"activeRecord"`
	CheckErrors             CT_DecimalNumber     `xml:"checkErrors"`
	Odso                    string               `xml:"odso"`
}
type CT_TargetScreenSz struct {
	Val ST_TargetScreenSz `xml:"val",attr`
}
type CT_Compat struct {
	UseSingleBorderforContiguousCells CT_OnOff           `xml:"useSingleBorderforContiguousCells"`
	WpJustification                   CT_OnOff           `xml:"wpJustification"`
	NoTabHangInd                      CT_OnOff           `xml:"noTabHangInd"`
	NoLeading                         CT_OnOff           `xml:"noLeading"`
	SpaceForUL                        CT_OnOff           `xml:"spaceForUL"`
	NoColumnBalance                   CT_OnOff           `xml:"noColumnBalance"`
	BalanceSingleByteDoubleByteWidth  CT_OnOff           `xml:"balanceSingleByteDoubleByteWidth"`
	NoExtraLineSpacing                CT_OnOff           `xml:"noExtraLineSpacing"`
	DoNotLeaveBackslashAlone          CT_OnOff           `xml:"doNotLeaveBackslashAlone"`
	UlTrailSpace                      CT_OnOff           `xml:"ulTrailSpace"`
	DoNotExpandShiftReturn            CT_OnOff           `xml:"doNotExpandShiftReturn"`
	SpacingInWholePoints              CT_OnOff           `xml:"spacingInWholePoints"`
	LineWrapLikeWord6                 CT_OnOff           `xml:"lineWrapLikeWord6"`
	PrintBodyTextBeforeHeader         CT_OnOff           `xml:"printBodyTextBeforeHeader"`
	PrintColBlack                     CT_OnOff           `xml:"printColBlack"`
	WpSpaceWidth                      CT_OnOff           `xml:"wpSpaceWidth"`
	ShowBreaksInFrames                CT_OnOff           `xml:"showBreaksInFrames"`
	SubFontBySize                     CT_OnOff           `xml:"subFontBySize"`
	SuppressBottomSpacing             CT_OnOff           `xml:"suppressBottomSpacing"`
	SuppressTopSpacing                CT_OnOff           `xml:"suppressTopSpacing"`
	SuppressSpacingAtTopOfPage        CT_OnOff           `xml:"suppressSpacingAtTopOfPage"`
	SuppressTopSpacingWP              CT_OnOff           `xml:"suppressTopSpacingWP"`
	SuppressSpBfAfterPgBrk            CT_OnOff           `xml:"suppressSpBfAfterPgBrk"`
	SwapBordersFacingPages            CT_OnOff           `xml:"swapBordersFacingPages"`
	ConvMailMergeEsc                  CT_OnOff           `xml:"convMailMergeEsc"`
	TruncateFontHeightsLikeWP6        CT_OnOff           `xml:"truncateFontHeightsLikeWP6"`
	MwSmallCaps                       CT_OnOff           `xml:"mwSmallCaps"`
	UsePrinterMetrics                 CT_OnOff           `xml:"usePrinterMetrics"`
	DoNotSuppressParagraphBorders     CT_OnOff           `xml:"doNotSuppressParagraphBorders"`
	WrapTrailSpaces                   CT_OnOff           `xml:"wrapTrailSpaces"`
	FootnoteLayoutLikeWW8             CT_OnOff           `xml:"footnoteLayoutLikeWW8"`
	ShapeLayoutLikeWW8                CT_OnOff           `xml:"shapeLayoutLikeWW8"`
	AlignTablesRowByRow               CT_OnOff           `xml:"alignTablesRowByRow"`
	ForgetLastTabAlignment            CT_OnOff           `xml:"forgetLastTabAlignment"`
	AdjustLineHeightInTable           CT_OnOff           `xml:"adjustLineHeightInTable"`
	AutoSpaceLikeWord95               CT_OnOff           `xml:"autoSpaceLikeWord95"`
	NoSpaceRaiseLower                 CT_OnOff           `xml:"noSpaceRaiseLower"`
	DoNotUseHTMLParagraphAutoSpacing  CT_OnOff           `xml:"doNotUseHTMLParagraphAutoSpacing"`
	LayoutRawTableWidth               CT_OnOff           `xml:"layoutRawTableWidth"`
	LayoutTableRowsApart              CT_OnOff           `xml:"layoutTableRowsApart"`
	UseWord97LineBreakRules           CT_OnOff           `xml:"useWord97LineBreakRules"`
	DoNotBreakWrappedTables           CT_OnOff           `xml:"doNotBreakWrappedTables"`
	DoNotSnapToGridInCell             CT_OnOff           `xml:"doNotSnapToGridInCell"`
	SelectFldWithFirstOrLastChar      CT_OnOff           `xml:"selectFldWithFirstOrLastChar"`
	ApplyBreakingRules                CT_OnOff           `xml:"applyBreakingRules"`
	DoNotWrapTextWithPunct            CT_OnOff           `xml:"doNotWrapTextWithPunct"`
	DoNotUseEastAsianBreakRules       CT_OnOff           `xml:"doNotUseEastAsianBreakRules"`
	UseWord2002TableStyleRules        CT_OnOff           `xml:"useWord2002TableStyleRules"`
	GrowAutofit                       CT_OnOff           `xml:"growAutofit"`
	UseFELayout                       CT_OnOff           `xml:"useFELayout"`
	UseNormalStyleForList             CT_OnOff           `xml:"useNormalStyleForList"`
	DoNotUseIndentAsNumberingTabStop  CT_OnOff           `xml:"doNotUseIndentAsNumberingTabStop"`
	UseAltKinsokuLineBreakRules       CT_OnOff           `xml:"useAltKinsokuLineBreakRules"`
	AllowSpaceOfSameStyleInTable      CT_OnOff           `xml:"allowSpaceOfSameStyleInTable"`
	DoNotSuppressIndentation          CT_OnOff           `xml:"doNotSuppressIndentation"`
	DoNotAutofitConstrainedTables     CT_OnOff           `xml:"doNotAutofitConstrainedTables"`
	AutofitToFirstFixedWidthCell      CT_OnOff           `xml:"autofitToFirstFixedWidthCell"`
	UnderlineTabInNumList             CT_OnOff           `xml:"underlineTabInNumList"`
	DisplayHangulFixedWidth           CT_OnOff           `xml:"displayHangulFixedWidth"`
	SplitPgBreakAndParaMark           CT_OnOff           `xml:"splitPgBreakAndParaMark"`
	DoNotVertAlignCellWithSp          CT_OnOff           `xml:"doNotVertAlignCellWithSp"`
	DoNotBreakConstrainedForcedTable  CT_OnOff           `xml:"doNotBreakConstrainedForcedTable"`
	DoNotVertAlignInTxbx              CT_OnOff           `xml:"doNotVertAlignInTxbx"`
	UseAnsiKerningPairs               CT_OnOff           `xml:"useAnsiKerningPairs"`
	CachedColBalance                  CT_OnOff           `xml:"cachedColBalance"`
	CompatSetting                     []CT_CompatSetting `xml:"compatSetting"`
}
type CT_CompatSetting struct {
	Name string `xml:"name",attr`
	Uri  string `xml:"uri",attr`
	Val  string `xml:"val",attr`
}
type CT_DocVar struct {
	Name string `xml:"name",attr`
	Val  string `xml:"val",attr`
}
type CT_DocVars struct {
	DocVar []CT_DocVar `xml:"docVar"`
}
type CT_DocRsids struct {
	RsidRoot string   `xml:"rsidRoot"`
	Rsid     []string `xml:"rsid"`
}
type CT_CharacterSpacing struct {
	Val ST_CharacterSpacing `xml:"val",attr`
}
type CT_SaveThroughXslt struct {
	SolutionID string `xml:"solutionID",attr`
}
type CT_RPrDefault struct {
	RPr CT_RPr `xml:"rPr"`
}
type CT_PPrDefault struct {
	PPr CT_PPrGeneral `xml:"pPr"`
}
type CT_DocDefaults struct {
	RPrDefault CT_RPrDefault `xml:"rPrDefault"`
	PPrDefault CT_PPrDefault `xml:"pPrDefault"`
}
type CT_ColorSchemeMapping struct {
	Bg1               string `xml:"bg1",attr`
	T1                string `xml:"t1",attr`
	Bg2               string `xml:"bg2",attr`
	T2                string `xml:"t2",attr`
	Accent1           string `xml:"accent1",attr`
	Accent2           string `xml:"accent2",attr`
	Accent3           string `xml:"accent3",attr`
	Accent4           string `xml:"accent4",attr`
	Accent5           string `xml:"accent5",attr`
	Accent6           string `xml:"accent6",attr`
	Hyperlink         string `xml:"hyperlink",attr`
	FollowedHyperlink string `xml:"followedHyperlink",attr`
}
type CT_ReadingModeInkLockDown struct {
	ActualPg string                    `xml:"actualPg",attr`
	W        string                    `xml:"w",attr`
	H        string                    `xml:"h",attr`
	FontSz   ST_DecimalNumberOrPercent `xml:"fontSz",attr`
}
type CT_WriteProtection struct {
	Recommended string `xml:"recommended",attr`
}
type CT_Settings struct {
	WriteProtection                     CT_WriteProtection        `xml:"writeProtection"`
	View                                CT_View                   `xml:"view"`
	Zoom                                CT_Zoom                   `xml:"zoom"`
	RemovePersonalInformation           CT_OnOff                  `xml:"removePersonalInformation"`
	RemoveDateAndTime                   CT_OnOff                  `xml:"removeDateAndTime"`
	DoNotDisplayPageBoundaries          CT_OnOff                  `xml:"doNotDisplayPageBoundaries"`
	DisplayBackgroundShape              CT_OnOff                  `xml:"displayBackgroundShape"`
	PrintPostScriptOverText             CT_OnOff                  `xml:"printPostScriptOverText"`
	PrintFractionalCharacterWidth       CT_OnOff                  `xml:"printFractionalCharacterWidth"`
	PrintFormsData                      CT_OnOff                  `xml:"printFormsData"`
	EmbedTrueTypeFonts                  CT_OnOff                  `xml:"embedTrueTypeFonts"`
	EmbedSystemFonts                    CT_OnOff                  `xml:"embedSystemFonts"`
	SaveSubsetFonts                     CT_OnOff                  `xml:"saveSubsetFonts"`
	SaveFormsData                       CT_OnOff                  `xml:"saveFormsData"`
	MirrorMargins                       CT_OnOff                  `xml:"mirrorMargins"`
	AlignBordersAndEdges                CT_OnOff                  `xml:"alignBordersAndEdges"`
	BordersDoNotSurroundHeader          CT_OnOff                  `xml:"bordersDoNotSurroundHeader"`
	BordersDoNotSurroundFooter          CT_OnOff                  `xml:"bordersDoNotSurroundFooter"`
	GutterAtTop                         CT_OnOff                  `xml:"gutterAtTop"`
	HideSpellingErrors                  CT_OnOff                  `xml:"hideSpellingErrors"`
	HideGrammaticalErrors               CT_OnOff                  `xml:"hideGrammaticalErrors"`
	ActiveWritingStyle                  []CT_WritingStyle         `xml:"activeWritingStyle"`
	ProofState                          CT_Proof                  `xml:"proofState"`
	FormsDesign                         CT_OnOff                  `xml:"formsDesign"`
	AttachedTemplate                    CT_Rel                    `xml:"attachedTemplate"`
	LinkStyles                          CT_OnOff                  `xml:"linkStyles"`
	StylePaneFormatFilter               CT_StylePaneFilter        `xml:"stylePaneFormatFilter"`
	StylePaneSortMethod                 CT_StyleSort              `xml:"stylePaneSortMethod"`
	DocumentType                        CT_DocType                `xml:"documentType"`
	MailMerge                           CT_MailMerge              `xml:"mailMerge"`
	RevisionView                        string                    `xml:"revisionView"`
	TrackRevisions                      CT_OnOff                  `xml:"trackRevisions"`
	DoNotTrackMoves                     CT_OnOff                  `xml:"doNotTrackMoves"`
	DoNotTrackFormatting                CT_OnOff                  `xml:"doNotTrackFormatting"`
	DocumentProtection                  CT_DocProtect             `xml:"documentProtection"`
	AutoFormatOverride                  CT_OnOff                  `xml:"autoFormatOverride"`
	StyleLockTheme                      CT_OnOff                  `xml:"styleLockTheme"`
	StyleLockQFSet                      CT_OnOff                  `xml:"styleLockQFSet"`
	DefaultTabStop                      string                    `xml:"defaultTabStop"`
	AutoHyphenation                     CT_OnOff                  `xml:"autoHyphenation"`
	ConsecutiveHyphenLimit              CT_DecimalNumber          `xml:"consecutiveHyphenLimit"`
	HyphenationZone                     string                    `xml:"hyphenationZone"`
	DoNotHyphenateCaps                  CT_OnOff                  `xml:"doNotHyphenateCaps"`
	ShowEnvelope                        CT_OnOff                  `xml:"showEnvelope"`
	SummaryLength                       CT_DecimalNumberOrPrecent `xml:"summaryLength"`
	ClickAndTypeStyle                   CT_String                 `xml:"clickAndTypeStyle"`
	DefaultTableStyle                   CT_String                 `xml:"defaultTableStyle"`
	EvenAndOddHeaders                   CT_OnOff                  `xml:"evenAndOddHeaders"`
	BookFoldRevPrinting                 CT_OnOff                  `xml:"bookFoldRevPrinting"`
	BookFoldPrinting                    CT_OnOff                  `xml:"bookFoldPrinting"`
	BookFoldPrintingSheets              CT_DecimalNumber          `xml:"bookFoldPrintingSheets"`
	DrawingGridHorizontalSpacing        string                    `xml:"drawingGridHorizontalSpacing"`
	DrawingGridVerticalSpacing          string                    `xml:"drawingGridVerticalSpacing"`
	DisplayHorizontalDrawingGridEvery   CT_DecimalNumber          `xml:"displayHorizontalDrawingGridEvery"`
	DisplayVerticalDrawingGridEvery     CT_DecimalNumber          `xml:"displayVerticalDrawingGridEvery"`
	DoNotUseMarginsForDrawingGridOrigin CT_OnOff                  `xml:"doNotUseMarginsForDrawingGridOrigin"`
	DrawingGridHorizontalOrigin         string                    `xml:"drawingGridHorizontalOrigin"`
	DrawingGridVerticalOrigin           string                    `xml:"drawingGridVerticalOrigin"`
	DoNotShadeFormData                  CT_OnOff                  `xml:"doNotShadeFormData"`
	NoPunctuationKerning                CT_OnOff                  `xml:"noPunctuationKerning"`
	CharacterSpacingControl             CT_CharacterSpacing       `xml:"characterSpacingControl"`
	PrintTwoOnOne                       CT_OnOff                  `xml:"printTwoOnOne"`
	StrictFirstAndLastChars             CT_OnOff                  `xml:"strictFirstAndLastChars"`
	NoLineBreaksAfter                   string                    `xml:"noLineBreaksAfter"`
	NoLineBreaksBefore                  string                    `xml:"noLineBreaksBefore"`
	SavePreviewPicture                  CT_OnOff                  `xml:"savePreviewPicture"`
	DoNotValidateAgainstSchema          CT_OnOff                  `xml:"doNotValidateAgainstSchema"`
	SaveInvalidXml                      CT_OnOff                  `xml:"saveInvalidXml"`
	IgnoreMixedContent                  CT_OnOff                  `xml:"ignoreMixedContent"`
	AlwaysShowPlaceholderText           CT_OnOff                  `xml:"alwaysShowPlaceholderText"`
	DoNotDemarcateInvalidXml            CT_OnOff                  `xml:"doNotDemarcateInvalidXml"`
	SaveXmlDataOnly                     CT_OnOff                  `xml:"saveXmlDataOnly"`
	UseXSLTWhenSaving                   CT_OnOff                  `xml:"useXSLTWhenSaving"`
	SaveThroughXslt                     string                    `xml:"saveThroughXslt"`
	ShowXMLTags                         CT_OnOff                  `xml:"showXMLTags"`
	AlwaysMergeEmptyNamespace           CT_OnOff                  `xml:"alwaysMergeEmptyNamespace"`
	UpdateFields                        CT_OnOff                  `xml:"updateFields"`
	HdrShapeDefaults                    string                    `xml:"hdrShapeDefaults"`
	FootnotePr                          string                    `xml:"footnotePr"`
	EndnotePr                           string                    `xml:"endnotePr"`
	Compat                              CT_Compat                 `xml:"compat"`
	DocVars                             string                    `xml:"docVars"`
	Rsids                               string                    `xml:"rsids"`
	AttachedSchema                      []CT_String               `xml:"attachedSchema"`
	ThemeFontLang                       CT_Language               `xml:"themeFontLang"`
	ClrSchemeMapping                    CT_ColorSchemeMapping     `xml:"clrSchemeMapping"`
	DoNotIncludeSubdocsInStats          CT_OnOff                  `xml:"doNotIncludeSubdocsInStats"`
	DoNotAutoCompressPictures           CT_OnOff                  `xml:"doNotAutoCompressPictures"`
	ForceUpgrade                        CT_Empty                  `xml:"forceUpgrade"`
	Captions                            string                    `xml:"captions"`
	ReadModeInkLockDown                 string                    `xml:"readModeInkLockDown"`
	SmartTagType                        []CT_SmartTagType         `xml:"smartTagType"`
	ShapeDefaults                       string                    `xml:"shapeDefaults"`
	DoNotEmbedSmartTags                 CT_OnOff                  `xml:"doNotEmbedSmartTags"`
	DecimalSymbol                       CT_String                 `xml:"decimalSymbol"`
	ListSeparator                       CT_String                 `xml:"listSeparator"`
}
type CT_StyleSort struct {
	Val ST_StyleSort `xml:"val",attr`
}
type CT_StylePaneFilter struct {
	AllStyles                    string `xml:"allStyles",attr`
	CustomStyles                 string `xml:"customStyles",attr`
	LatentStyles                 string `xml:"latentStyles",attr`
	StylesInUse                  string `xml:"stylesInUse",attr`
	HeadingStyles                string `xml:"headingStyles",attr`
	NumberingStyles              string `xml:"numberingStyles",attr`
	TableStyles                  string `xml:"tableStyles",attr`
	DirectFormattingOnRuns       string `xml:"directFormattingOnRuns",attr`
	DirectFormattingOnParagraphs string `xml:"directFormattingOnParagraphs",attr`
	DirectFormattingOnNumbering  string `xml:"directFormattingOnNumbering",attr`
	DirectFormattingOnTables     string `xml:"directFormattingOnTables",attr`
	ClearFormatting              string `xml:"clearFormatting",attr`
	Top3HeadingStyles            string `xml:"top3HeadingStyles",attr`
	VisibleStyles                string `xml:"visibleStyles",attr`
	AlternateStyleNames          string `xml:"alternateStyleNames",attr`
	Val                          string `xml:"val",attr`
}
type CT_WebSettings struct {
	Frameset              string            `xml:"frameset"`
	Divs                  string            `xml:"divs"`
	Encoding              CT_String         `xml:"encoding"`
	OptimizeForBrowser    string            `xml:"optimizeForBrowser"`
	RelyOnVML             CT_OnOff          `xml:"relyOnVML"`
	AllowPNG              CT_OnOff          `xml:"allowPNG"`
	DoNotRelyOnCSS        CT_OnOff          `xml:"doNotRelyOnCSS"`
	DoNotSaveAsSingleFile CT_OnOff          `xml:"doNotSaveAsSingleFile"`
	DoNotOrganizeInFolder CT_OnOff          `xml:"doNotOrganizeInFolder"`
	DoNotUseLongFileNames CT_OnOff          `xml:"doNotUseLongFileNames"`
	PixelsPerInch         CT_DecimalNumber  `xml:"pixelsPerInch"`
	TargetScreenSz        CT_TargetScreenSz `xml:"targetScreenSz"`
	SaveSmartTagsAsXml    CT_OnOff          `xml:"saveSmartTagsAsXml"`
}
type CT_FrameScrollbar struct {
	Val ST_FrameScrollbar `xml:"val",attr`
}
type CT_OptimizeForBrowser struct {
	CT_OnOff
	Target string `xml:"target",attr`
}
type CT_Frame struct {
	Sz              CT_String         `xml:"sz"`
	Name            CT_String         `xml:"name"`
	Title           CT_String         `xml:"title"`
	LongDesc        CT_Rel            `xml:"longDesc"`
	SourceFileName  CT_Rel            `xml:"sourceFileName"`
	MarW            string            `xml:"marW"`
	MarH            string            `xml:"marH"`
	Scrollbar       CT_FrameScrollbar `xml:"scrollbar"`
	NoResizeAllowed CT_OnOff          `xml:"noResizeAllowed"`
	LinkedToFile    CT_OnOff          `xml:"linkedToFile"`
}
type CT_FrameLayout struct {
	Val ST_FrameLayout `xml:"val",attr`
}
type CT_FramesetSplitbar struct {
	W           string   `xml:"w"`
	Color       CT_Color `xml:"color"`
	NoBorder    CT_OnOff `xml:"noBorder"`
	FlatBorders CT_OnOff `xml:"flatBorders"`
}
type CT_Frameset struct {
	Sz               CT_String      `xml:"sz"`
	FramesetSplitbar string         `xml:"framesetSplitbar"`
	FrameLayout      CT_FrameLayout `xml:"frameLayout"`
	Title            CT_String      `xml:"title"`
	Frameset         []string       `xml:"frameset"`
	Frame            []CT_Frame     `xml:"frame"`
}
type CT_NumPicBullet struct {
	Pict           CT_Picture       `xml:"pict"`
	Drawing        CT_Drawing       `xml:"drawing"`
	NumPicBulletId ST_DecimalNumber `xml:"numPicBulletId",attr`
}
type CT_LevelSuffix struct {
	Val string `xml:"val",attr`
}
type CT_LevelText struct {
	Val  string `xml:"val",attr`
	Null string `xml:"null",attr`
}
type CT_LvlLegacy struct {
	Legacy       string `xml:"legacy",attr`
	LegacySpace  string `xml:"legacySpace",attr`
	LegacyIndent string `xml:"legacyIndent",attr`
}
type CT_Lvl struct {
	Start          CT_DecimalNumber `xml:"start"`
	NumFmt         CT_NumFmt        `xml:"numFmt"`
	LvlRestart     CT_DecimalNumber `xml:"lvlRestart"`
	PStyle         CT_String        `xml:"pStyle"`
	IsLgl          CT_OnOff         `xml:"isLgl"`
	Suff           string           `xml:"suff"`
	LvlText        string           `xml:"lvlText"`
	LvlPicBulletId CT_DecimalNumber `xml:"lvlPicBulletId"`
	Legacy         CT_LvlLegacy     `xml:"legacy"`
	LvlJc          CT_Jc            `xml:"lvlJc"`
	PPr            CT_PPrGeneral    `xml:"pPr"`
	RPr            CT_RPr           `xml:"rPr"`
	Ilvl           ST_DecimalNumber `xml:"ilvl",attr`
	Tplc           string           `xml:"tplc",attr`
	Tentative      string           `xml:"tentative",attr`
}
type CT_MultiLevelType struct {
	Val ST_MultiLevelType `xml:"val",attr`
}
type CT_AbstractNum struct {
	Nsid           string            `xml:"nsid"`
	MultiLevelType CT_MultiLevelType `xml:"multiLevelType"`
	Tmpl           string            `xml:"tmpl"`
	Name           CT_String         `xml:"name"`
	StyleLink      CT_String         `xml:"styleLink"`
	NumStyleLink   CT_String         `xml:"numStyleLink"`
	Lvl            CT_Lvl            `xml:"lvl"`
	AbstractNumId  ST_DecimalNumber  `xml:"abstractNumId",attr`
}
type CT_NumLvl struct {
	StartOverride CT_DecimalNumber `xml:"startOverride"`
	Lvl           CT_Lvl           `xml:"lvl"`
	Ilvl          ST_DecimalNumber `xml:"ilvl",attr`
}
type CT_Num struct {
	AbstractNumId CT_DecimalNumber `xml:"abstractNumId"`
	LvlOverride   CT_NumLvl        `xml:"lvlOverride"`
	NumId         ST_DecimalNumber `xml:"numId",attr`
}
type CT_Numbering struct {
	NumPicBullet      []CT_NumPicBullet `xml:"numPicBullet"`
	AbstractNum       []string          `xml:"abstractNum"`
	Num               []CT_Num          `xml:"num"`
	NumIdMacAtCleanup CT_DecimalNumber  `xml:"numIdMacAtCleanup"`
}
type CT_TblStylePr struct {
	PPr   CT_PPrGeneral `xml:"pPr"`
	RPr   CT_RPr        `xml:"rPr"`
	TblPr string        `xml:"tblPr"`
	TrPr  CT_TrPr       `xml:"trPr"`
	TcPr  CT_TcPr       `xml:"tcPr"`
	Type  string        `xml:"type",attr`
}
type CT_Style struct {
	Name            CT_String        `xml:"name"`
	Aliases         CT_String        `xml:"aliases"`
	BasedOn         CT_String        `xml:"basedOn"`
	Next            CT_String        `xml:"next"`
	Link            CT_String        `xml:"link"`
	AutoRedefine    CT_OnOff         `xml:"autoRedefine"`
	Hidden          CT_OnOff         `xml:"hidden"`
	UiPriority      CT_DecimalNumber `xml:"uiPriority"`
	SemiHidden      CT_OnOff         `xml:"semiHidden"`
	UnhideWhenUsed  CT_OnOff         `xml:"unhideWhenUsed"`
	QFormat         CT_OnOff         `xml:"qFormat"`
	Locked          CT_OnOff         `xml:"locked"`
	Personal        CT_OnOff         `xml:"personal"`
	PersonalCompose CT_OnOff         `xml:"personalCompose"`
	PersonalReply   CT_OnOff         `xml:"personalReply"`
	Rsid            string           `xml:"rsid"`
	PPr             CT_PPrGeneral    `xml:"pPr"`
	RPr             CT_RPr           `xml:"rPr"`
	TblPr           string           `xml:"tblPr"`
	TrPr            CT_TrPr          `xml:"trPr"`
	TcPr            CT_TcPr          `xml:"tcPr"`
	TblStylePr      []CT_TblStylePr  `xml:"tblStylePr"`
	Type            ST_StyleType     `xml:"type",attr`
	StyleId         string           `xml:"styleId",attr`
	Default         string           `xml:"default",attr`
	CustomStyle     string           `xml:"customStyle",attr`
}
type CT_LsdException struct {
	Name           string           `xml:"name",attr`
	Locked         string           `xml:"locked",attr`
	UiPriority     ST_DecimalNumber `xml:"uiPriority",attr`
	SemiHidden     string           `xml:"semiHidden",attr`
	UnhideWhenUsed string           `xml:"unhideWhenUsed",attr`
	QFormat        string           `xml:"qFormat",attr`
}
type CT_LatentStyles struct {
	LsdException      []string         `xml:"lsdException"`
	DefLockedState    string           `xml:"defLockedState",attr`
	DefUIPriority     ST_DecimalNumber `xml:"defUIPriority",attr`
	DefSemiHidden     string           `xml:"defSemiHidden",attr`
	DefUnhideWhenUsed string           `xml:"defUnhideWhenUsed",attr`
	DefQFormat        string           `xml:"defQFormat",attr`
	Count             ST_DecimalNumber `xml:"count",attr`
}
type CT_Styles struct {
	DocDefaults  string     `xml:"docDefaults"`
	LatentStyles string     `xml:"latentStyles"`
	Style        []CT_Style `xml:"style"`
}
type CT_Panose struct {
	Val string `xml:"val",attr`
}
type CT_FontFamily struct {
	Val ST_FontFamily `xml:"val",attr`
}
type CT_Pitch struct {
	Val ST_Pitch `xml:"val",attr`
}
type CT_FontSig struct {
	Usb0 string `xml:"usb0",attr`
	Usb1 string `xml:"usb1",attr`
	Usb2 string `xml:"usb2",attr`
	Usb3 string `xml:"usb3",attr`
	Csb0 string `xml:"csb0",attr`
	Csb1 string `xml:"csb1",attr`
}
type CT_FontRel struct {
	CT_Rel
	FontKey   string `xml:"fontKey",attr`
	Subsetted string `xml:"subsetted",attr`
}
type CT_Font struct {
	AltName         CT_String     `xml:"altName"`
	Panose1         string        `xml:"panose1"`
	Charset         string        `xml:"charset"`
	Family          CT_FontFamily `xml:"family"`
	NotTrueType     CT_OnOff      `xml:"notTrueType"`
	Pitch           CT_Pitch      `xml:"pitch"`
	Sig             CT_FontSig    `xml:"sig"`
	EmbedRegular    CT_FontRel    `xml:"embedRegular"`
	EmbedBold       CT_FontRel    `xml:"embedBold"`
	EmbedItalic     CT_FontRel    `xml:"embedItalic"`
	EmbedBoldItalic CT_FontRel    `xml:"embedBoldItalic"`
	Name            string        `xml:"name",attr`
}
type CT_FontsList struct {
	Font []CT_Font `xml:"font"`
}
type CT_DivBdr struct {
	Top    string `xml:"top"`
	Left   string `xml:"left"`
	Bottom string `xml:"bottom"`
	Right  string `xml:"right"`
}
type CT_Div struct {
	BlockQuote CT_OnOff         `xml:"blockQuote"`
	BodyDiv    CT_OnOff         `xml:"bodyDiv"`
	MarLeft    string           `xml:"marLeft"`
	MarRight   string           `xml:"marRight"`
	MarTop     string           `xml:"marTop"`
	MarBottom  string           `xml:"marBottom"`
	DivBdr     string           `xml:"divBdr"`
	DivsChild  []string         `xml:"divsChild"`
	Id         ST_DecimalNumber `xml:"id",attr`
}
type CT_Divs struct {
	Div CT_Div `xml:"div"`
}
type CT_TxbxContent struct {
}
type CT_Body struct {
	SectPr         CT_SectPr `xml:"sectPr"`
	BlockLevelElts []EG_BlockLevelElts
}
type CT_ShapeDefaults struct {
}
type CT_Comments struct {
	Comment []CT_Comment `xml:"comment"`
}
type CT_Footnotes struct {
	Footnote string `xml:"footnote"`
}
type CT_Endnotes struct {
	Endnote string `xml:"endnote"`
}
type CT_SmartTagType struct {
	Namespaceuri string `xml:"namespaceuri",attr`
	Name         string `xml:"name",attr`
	Url          string `xml:"url",attr`
}
type CT_DocPartBehavior struct {
	Val ST_DocPartBehavior `xml:"val",attr`
}
type CT_DocPartBehaviors struct {
	Behavior []CT_DocPartBehavior `xml:"behavior"`
}
type CT_DocPartType struct {
	Val ST_DocPartType `xml:"val",attr`
}
type CT_DocPartTypes struct {
	Type []CT_DocPartType `xml:"type"`
	All  string           `xml:"all",attr`
}
type CT_DocPartGallery struct {
	Val ST_DocPartGallery `xml:"val",attr`
}
type CT_DocPartCategory struct {
	Name    CT_String         `xml:"name"`
	Gallery CT_DocPartGallery `xml:"gallery"`
}
type CT_DocPartName struct {
	Val       string `xml:"val",attr`
	Decorated string `xml:"decorated",attr`
}
type CT_DocPartPr struct {
	Name        CT_DocPartName     `xml:"name"`
	Style       CT_String          `xml:"style"`
	Category    CT_DocPartCategory `xml:"category"`
	Types       string             `xml:"types"`
	Behaviors   string             `xml:"behaviors"`
	Description CT_String          `xml:"description"`
	Guid        string             `xml:"guid"`
}
type CT_DocPart struct {
	DocPartPr   CT_DocPartPr `xml:"docPartPr"`
	DocPartBody string       `xml:"docPartBody"`
}
type CT_DocParts struct {
	DocPart []CT_DocPart `xml:"docPart"`
}
type CT_Caption struct {
	Name    string           `xml:"name",attr`
	Pos     string           `xml:"pos",attr`
	ChapNum string           `xml:"chapNum",attr`
	Heading ST_DecimalNumber `xml:"heading",attr`
	NoLabel string           `xml:"noLabel",attr`
	NumFmt  ST_NumberFormat  `xml:"numFmt",attr`
	Sep     ST_ChapterSep    `xml:"sep",attr`
}
type CT_AutoCaption struct {
	Name    string `xml:"name",attr`
	Caption string `xml:"caption",attr`
}
type CT_AutoCaptions struct {
	AutoCaption []CT_AutoCaption `xml:"autoCaption"`
}
type CT_Captions struct {
	Caption      []CT_Caption `xml:"caption"`
	AutoCaptions string       `xml:"autoCaptions"`
}
type CT_DocumentBase struct {
	Background string `xml:"background"`
}
type CT_Document struct {
	CT_DocumentBase
	Body        string `xml:"body"`
	Conformance string `xml:"conformance",attr`
}
type CT_GlossaryDocument struct {
	CT_DocumentBase
	DocParts string `xml:"docParts"`
}
type ST_LongHexNumber struct {
	ST_LongHexNumber string `xml:"ST_LongHexNumber"`
}
type ST_ShortHexNumber struct {
	ST_ShortHexNumber string `xml:"ST_ShortHexNumber"`
}
type ST_UcharHexNumber struct {
	ST_UcharHexNumber string `xml:"ST_UcharHexNumber"`
}
type ST_DecimalNumberOrPercent struct {
	ST_DecimalNumberOrPercent string `xml:"ST_DecimalNumberOrPercent"`
}
type ST_UnqualifiedPercentage struct {
	ST_UnqualifiedPercentage string `xml:"ST_UnqualifiedPercentage"`
}
type ST_DecimalNumber struct {
	ST_DecimalNumber string `xml:"ST_DecimalNumber"`
}
type ST_SignedTwipsMeasure struct {
	ST_SignedTwipsMeasure string `xml:"ST_SignedTwipsMeasure"`
}
type ST_PixelsMeasure struct {
	ST_PixelsMeasure string `xml:"ST_PixelsMeasure"`
}
type ST_HpsMeasure struct {
	ST_HpsMeasure string `xml:"ST_HpsMeasure"`
}
type ST_SignedHpsMeasure struct {
	ST_SignedHpsMeasure string `xml:"ST_SignedHpsMeasure"`
}
type ST_DateTime struct {
	ST_DateTime string `xml:"ST_DateTime"`
}
type ST_MacroName struct {
	ST_MacroName string `xml:"ST_MacroName"`
}
type ST_EighthPointMeasure struct {
	ST_EighthPointMeasure string `xml:"ST_EighthPointMeasure"`
}
type ST_PointMeasure struct {
	ST_PointMeasure string `xml:"ST_PointMeasure"`
}
type ST_TextScale struct {
	ST_TextScale string `xml:"ST_TextScale"`
}
type ST_TextScalePercent struct {
	ST_TextScalePercent string `xml:"ST_TextScalePercent"`
}
type ST_TextScaleDecimal struct {
	ST_TextScaleDecimal string `xml:"ST_TextScaleDecimal"`
}
type ST_HighlightColor struct {
	ST_HighlightColor string `xml:"ST_HighlightColor"`
}
type ST_HexColorAuto struct {
	ST_HexColorAuto string `xml:"ST_HexColorAuto"`
}
type ST_HexColor struct {
	ST_HexColor string `xml:"ST_HexColor"`
}
type ST_Underline struct {
	ST_Underline string `xml:"ST_Underline"`
}
type ST_TextEffect struct {
	ST_TextEffect string `xml:"ST_TextEffect"`
}
type ST_Border struct {
	ST_Border string `xml:"ST_Border"`
}
type ST_Shd struct {
	ST_Shd string `xml:"ST_Shd"`
}
type ST_Em struct {
	ST_Em string `xml:"ST_Em"`
}
type ST_CombineBrackets struct {
	ST_CombineBrackets string `xml:"ST_CombineBrackets"`
}
type ST_HeightRule struct {
	ST_HeightRule string `xml:"ST_HeightRule"`
}
type ST_Wrap struct {
	ST_Wrap string `xml:"ST_Wrap"`
}
type ST_VAnchor struct {
	ST_VAnchor string `xml:"ST_VAnchor"`
}
type ST_HAnchor struct {
	ST_HAnchor string `xml:"ST_HAnchor"`
}
type ST_DropCap struct {
	ST_DropCap string `xml:"ST_DropCap"`
}
type ST_TabJc struct {
	ST_TabJc string `xml:"ST_TabJc"`
}
type ST_TabTlc struct {
	ST_TabTlc string `xml:"ST_TabTlc"`
}
type ST_LineSpacingRule struct {
	ST_LineSpacingRule string `xml:"ST_LineSpacingRule"`
}
type ST_Jc struct {
	ST_Jc string `xml:"ST_Jc"`
}
type ST_JcTable struct {
	ST_JcTable string `xml:"ST_JcTable"`
}
type ST_View struct {
	ST_View string `xml:"ST_View"`
}
type ST_Zoom struct {
	ST_Zoom string `xml:"ST_Zoom"`
}
type ST_Proof struct {
	ST_Proof string `xml:"ST_Proof"`
}
type ST_DocType struct {
	ST_DocType string `xml:"ST_DocType"`
}
type ST_DocProtect struct {
	ST_DocProtect string `xml:"ST_DocProtect"`
}
type ST_MailMergeDocType struct {
	ST_MailMergeDocType string `xml:"ST_MailMergeDocType"`
}
type ST_MailMergeDataType struct {
	ST_MailMergeDataType string `xml:"ST_MailMergeDataType"`
}
type ST_MailMergeDest struct {
	ST_MailMergeDest string `xml:"ST_MailMergeDest"`
}
type ST_MailMergeOdsoFMDFieldType struct {
	ST_MailMergeOdsoFMDFieldType string `xml:"ST_MailMergeOdsoFMDFieldType"`
}
type ST_TextDirection struct {
	ST_TextDirection string `xml:"ST_TextDirection"`
}
type ST_TextAlignment struct {
	ST_TextAlignment string `xml:"ST_TextAlignment"`
}
type ST_DisplacedByCustomXml struct {
	ST_DisplacedByCustomXml string `xml:"ST_DisplacedByCustomXml"`
}
type ST_AnnotationVMerge struct {
	ST_AnnotationVMerge string `xml:"ST_AnnotationVMerge"`
}
type ST_TextboxTightWrap struct {
	ST_TextboxTightWrap string `xml:"ST_TextboxTightWrap"`
}
type ST_ObjectDrawAspect struct {
	ST_ObjectDrawAspect string `xml:"ST_ObjectDrawAspect"`
}
type ST_ObjectUpdateMode struct {
	ST_ObjectUpdateMode string `xml:"ST_ObjectUpdateMode"`
}
type ST_FldCharType struct {
	ST_FldCharType string `xml:"ST_FldCharType"`
}
type ST_InfoTextType struct {
	ST_InfoTextType string `xml:"ST_InfoTextType"`
}
type ST_FFHelpTextVal struct {
	ST_FFHelpTextVal string `xml:"ST_FFHelpTextVal"`
}
type ST_FFStatusTextVal struct {
	ST_FFStatusTextVal string `xml:"ST_FFStatusTextVal"`
}
type ST_FFName struct {
	ST_FFName string `xml:"ST_FFName"`
}
type ST_FFTextType struct {
	ST_FFTextType string `xml:"ST_FFTextType"`
}
type ST_SectionMark struct {
	ST_SectionMark string `xml:"ST_SectionMark"`
}
type ST_NumberFormat struct {
	ST_NumberFormat string `xml:"ST_NumberFormat"`
}
type ST_PageOrientation struct {
	ST_PageOrientation string `xml:"ST_PageOrientation"`
}
type ST_PageBorderZOrder struct {
	ST_PageBorderZOrder string `xml:"ST_PageBorderZOrder"`
}
type ST_PageBorderDisplay struct {
	ST_PageBorderDisplay string `xml:"ST_PageBorderDisplay"`
}
type ST_PageBorderOffset struct {
	ST_PageBorderOffset string `xml:"ST_PageBorderOffset"`
}
type ST_ChapterSep struct {
	ST_ChapterSep string `xml:"ST_ChapterSep"`
}
type ST_LineNumberRestart struct {
	ST_LineNumberRestart string `xml:"ST_LineNumberRestart"`
}
type ST_VerticalJc struct {
	ST_VerticalJc string `xml:"ST_VerticalJc"`
}
type ST_DocGrid struct {
	ST_DocGrid string `xml:"ST_DocGrid"`
}
type ST_HdrFtr struct {
	ST_HdrFtr string `xml:"ST_HdrFtr"`
}
type ST_FtnEdn struct {
	ST_FtnEdn string `xml:"ST_FtnEdn"`
}
type ST_BrType struct {
	ST_BrType string `xml:"ST_BrType"`
}
type ST_BrClear struct {
	ST_BrClear string `xml:"ST_BrClear"`
}
type ST_PTabAlignment struct {
	ST_PTabAlignment string `xml:"ST_PTabAlignment"`
}
type ST_PTabRelativeTo struct {
	ST_PTabRelativeTo string `xml:"ST_PTabRelativeTo"`
}
type ST_PTabLeader struct {
	ST_PTabLeader string `xml:"ST_PTabLeader"`
}
type ST_ProofErr struct {
	ST_ProofErr string `xml:"ST_ProofErr"`
}
type ST_EdGrp struct {
	ST_EdGrp string `xml:"ST_EdGrp"`
}
type ST_Hint struct {
	ST_Hint string `xml:"ST_Hint"`
}
type ST_Theme struct {
	ST_Theme string `xml:"ST_Theme"`
}
type ST_RubyAlign struct {
	ST_RubyAlign string `xml:"ST_RubyAlign"`
}
type ST_Lock struct {
	ST_Lock string `xml:"ST_Lock"`
}
type ST_SdtDateMappingType struct {
	ST_SdtDateMappingType string `xml:"ST_SdtDateMappingType"`
}
type ST_Direction struct {
	ST_Direction string `xml:"ST_Direction"`
}
type ST_TblWidth struct {
	ST_TblWidth string `xml:"ST_TblWidth"`
}
type ST_MeasurementOrPercent struct {
	ST_MeasurementOrPercent string `xml:"ST_MeasurementOrPercent"`
}
type ST_Merge struct {
	ST_Merge string `xml:"ST_Merge"`
}
type ST_Cnf struct {
	ST_Cnf string `xml:"ST_Cnf"`
}
type ST_TblLayoutType struct {
	ST_TblLayoutType string `xml:"ST_TblLayoutType"`
}
type ST_TblOverlap struct {
	ST_TblOverlap string `xml:"ST_TblOverlap"`
}
type ST_FtnPos struct {
	ST_FtnPos string `xml:"ST_FtnPos"`
}
type ST_EdnPos struct {
	ST_EdnPos string `xml:"ST_EdnPos"`
}
type ST_RestartNumber struct {
	ST_RestartNumber string `xml:"ST_RestartNumber"`
}
type ST_MailMergeSourceType struct {
	ST_MailMergeSourceType string `xml:"ST_MailMergeSourceType"`
}
type ST_TargetScreenSz struct {
	ST_TargetScreenSz string `xml:"ST_TargetScreenSz"`
}
type ST_CharacterSpacing struct {
	ST_CharacterSpacing string `xml:"ST_CharacterSpacing"`
}
type ST_WmlColorSchemeIndex struct {
	ST_WmlColorSchemeIndex string `xml:"ST_WmlColorSchemeIndex"`
}
type ST_StyleSort struct {
	ST_StyleSort string `xml:"ST_StyleSort"`
}
type ST_FrameScrollbar struct {
	ST_FrameScrollbar string `xml:"ST_FrameScrollbar"`
}
type ST_FrameLayout struct {
	ST_FrameLayout string `xml:"ST_FrameLayout"`
}
type ST_LevelSuffix struct {
	ST_LevelSuffix string `xml:"ST_LevelSuffix"`
}
type ST_MultiLevelType struct {
	ST_MultiLevelType string `xml:"ST_MultiLevelType"`
}
type ST_TblStyleOverrideType struct {
	ST_TblStyleOverrideType string `xml:"ST_TblStyleOverrideType"`
}
type ST_StyleType struct {
	ST_StyleType string `xml:"ST_StyleType"`
}
type ST_FontFamily struct {
	ST_FontFamily string `xml:"ST_FontFamily"`
}
type ST_Pitch struct {
	ST_Pitch string `xml:"ST_Pitch"`
}
type ST_ThemeColor struct {
	ST_ThemeColor string `xml:"ST_ThemeColor"`
}
type ST_DocPartBehavior struct {
	ST_DocPartBehavior string `xml:"ST_DocPartBehavior"`
}
type ST_DocPartType struct {
	ST_DocPartType string `xml:"ST_DocPartType"`
}
type ST_DocPartGallery struct {
	ST_DocPartGallery string `xml:"ST_DocPartGallery"`
}
type ST_CaptionPos struct {
	ST_CaptionPos string `xml:"ST_CaptionPos"`
}
