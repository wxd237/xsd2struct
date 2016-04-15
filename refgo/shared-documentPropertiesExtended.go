package ooxml
type WordSchema  struct {
	Properties string  `xml:"Properties"`
}
type CT_Properties struct{
	Template string  `xml:"Template"`
	Manager string  `xml:"Manager"`
	Company string  `xml:"Company"`
	Pages string  `xml:"Pages"`
	Words string  `xml:"Words"`
	Characters string  `xml:"Characters"`
	PresentationFormat string  `xml:"PresentationFormat"`
	Lines string  `xml:"Lines"`
	Paragraphs string  `xml:"Paragraphs"`
	Slides string  `xml:"Slides"`
	Notes string  `xml:"Notes"`
	TotalTime string  `xml:"TotalTime"`
	HiddenSlides string  `xml:"HiddenSlides"`
	MMClips string  `xml:"MMClips"`
	ScaleCrop string  `xml:"ScaleCrop"`
	HeadingPairs CT_VectorVariant  `xml:"HeadingPairs"`
	TitlesOfParts string  `xml:"TitlesOfParts"`
	LinksUpToDate string  `xml:"LinksUpToDate"`
	CharactersWithSpaces string  `xml:"CharactersWithSpaces"`
	SharedDoc string  `xml:"SharedDoc"`
	HyperlinkBase string  `xml:"HyperlinkBase"`
	HLinks CT_VectorVariant  `xml:"HLinks"`
	HyperlinksChanged string  `xml:"HyperlinksChanged"`
	DigSig CT_DigSigBlob  `xml:"DigSig"`
	Application string  `xml:"Application"`
	AppVersion string  `xml:"AppVersion"`
	DocSecurity string  `xml:"DocSecurity"`
}
type CT_VectorVariant struct{
}
type CT_VectorLpstr struct{
}
type CT_DigSigBlob struct{
}
