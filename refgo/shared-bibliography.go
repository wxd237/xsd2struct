package ooxml
type WordSchema  struct {
	Sources string  `xml:"Sources"`
}
type CT_NameListType struct{
	Person []string  `xml:"Person"`
}
type CT_PersonType struct{
	Last []string  `xml:"Last"`
	First []string  `xml:"First"`
	Middle []string  `xml:"Middle"`
}
type CT_NameType struct{
	NameList string  `xml:"NameList"`
}
type CT_NameOrCorporateType struct{
	NameList string  `xml:"NameList"`
	Corporate string  `xml:"Corporate"`
}
type CT_AuthorType struct{
	Artist CT_NameType  `xml:"Artist"`
	Author CT_NameOrCorporateType  `xml:"Author"`
	BookAuthor CT_NameType  `xml:"BookAuthor"`
	Compiler CT_NameType  `xml:"Compiler"`
	Composer CT_NameType  `xml:"Composer"`
	Conductor CT_NameType  `xml:"Conductor"`
	Counsel CT_NameType  `xml:"Counsel"`
	Director CT_NameType  `xml:"Director"`
	Editor CT_NameType  `xml:"Editor"`
	Interviewee CT_NameType  `xml:"Interviewee"`
	Interviewer CT_NameType  `xml:"Interviewer"`
	Inventor CT_NameType  `xml:"Inventor"`
	Performer CT_NameOrCorporateType  `xml:"Performer"`
	ProducerName CT_NameType  `xml:"ProducerName"`
	Translator CT_NameType  `xml:"Translator"`
	Writer CT_NameType  `xml:"Writer"`
}
type CT_SourceType struct{
	AbbreviatedCaseNumber string  `xml:"AbbreviatedCaseNumber"`
	AlbumTitle string  `xml:"AlbumTitle"`
	Author CT_AuthorType  `xml:"Author"`
	BookTitle string  `xml:"BookTitle"`
	Broadcaster string  `xml:"Broadcaster"`
	BroadcastTitle string  `xml:"BroadcastTitle"`
	CaseNumber string  `xml:"CaseNumber"`
	ChapterNumber string  `xml:"ChapterNumber"`
	City string  `xml:"City"`
	Comments string  `xml:"Comments"`
	ConferenceName string  `xml:"ConferenceName"`
	CountryRegion string  `xml:"CountryRegion"`
	Court string  `xml:"Court"`
	Day string  `xml:"Day"`
	DayAccessed string  `xml:"DayAccessed"`
	Department string  `xml:"Department"`
	Distributor string  `xml:"Distributor"`
	Edition string  `xml:"Edition"`
	Guid string  `xml:"Guid"`
	Institution string  `xml:"Institution"`
	InternetSiteTitle string  `xml:"InternetSiteTitle"`
	Issue string  `xml:"Issue"`
	JournalName string  `xml:"JournalName"`
	LCID string  `xml:"LCID"`
	Medium string  `xml:"Medium"`
	Month string  `xml:"Month"`
	MonthAccessed string  `xml:"MonthAccessed"`
	NumberVolumes string  `xml:"NumberVolumes"`
	Pages string  `xml:"Pages"`
	PatentNumber string  `xml:"PatentNumber"`
	PeriodicalTitle string  `xml:"PeriodicalTitle"`
	ProductionCompany string  `xml:"ProductionCompany"`
	PublicationTitle string  `xml:"PublicationTitle"`
	Publisher string  `xml:"Publisher"`
	RecordingNumber string  `xml:"RecordingNumber"`
	RefOrder string  `xml:"RefOrder"`
	Reporter string  `xml:"Reporter"`
	SourceType ST_SourceType  `xml:"SourceType"`
	ShortTitle string  `xml:"ShortTitle"`
	StandardNumber string  `xml:"StandardNumber"`
	StateProvince string  `xml:"StateProvince"`
	Station string  `xml:"Station"`
	Tag string  `xml:"Tag"`
	Theater string  `xml:"Theater"`
	ThesisType string  `xml:"ThesisType"`
	Title string  `xml:"Title"`
	Type string  `xml:"Type"`
	URL string  `xml:"URL"`
	Version string  `xml:"Version"`
	Volume string  `xml:"Volume"`
	Year string  `xml:"Year"`
	YearAccessed string  `xml:"YearAccessed"`
}
type CT_Sources struct{
	Source []CT_SourceType  `xml:"Source"`
	SelectedStyle string  `xml:"SelectedStyle",attr`
	StyleName string  `xml:"StyleName",attr`
	URI string  `xml:"URI",attr`
}
