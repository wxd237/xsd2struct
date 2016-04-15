package xsd2struct

import (
	"testing"
)

var s schmea

func TestNewXSDFile(t *testing.T) {

}

func TestDump(t *testing.T) {

}

func TestParseTo(t *testing.T) {
	tmap := map[string]string{
		"ref/dml-chart.xsd":                             "DmlChart",
		"ref/dml-chartDrawing.xsd":                      "DmlChartdrawing",
		"ref/dml-diagram.xsd":                           "DmlDiagram",
		"ref/dml-lockedCanvas.xsd":                      "DmlLockedcanvas",
		"ref/dml-main.xsd":                              "DmlMain",
		"ref/dml-picture.xsd":                           "DmlPicture",
		"ref/dml-spreadsheetDrawing.xsd":                "DmlSpreadsheetdrawing",
		"ref/dml-wordprocessingDrawing.xsd":             "DmlWordprocessingdrawing",
		"ref/pml.xsd":                                   "Pml",
		"ref/shared-additionalCharacteristics.xsd":      "SharedAdditionalcharacteristics",
		"ref/shared-bibliography.xsd":                   "SharedBibliography",
		"ref/shared-commonSimpleTypes.xsd":              "SharedCommonsimpletypes",
		"ref/shared-customXmlDataProperties.xsd":        "SharedCustomxmldataproperties",
		"ref/shared-customXmlSchemaProperties.xsd":      "SharedCustomxmlschemaproperties",
		"ref/shared-documentPropertiesCustom.xsd":       "SharedDocumentpropertiescustom",
		"ref/shared-documentPropertiesExtended.xsd":     "SharedDocumentpropertiesextended",
		"ref/shared-documentPropertiesVariantTypes.xsd": "SharedDocumentpropertiesvarianttypes",
		"ref/shared-math.xsd":                           "SharedMath",
		"ref/shared-relationshipReference.xsd":          "SharedRelationshipreference",
		"ref/sml.xsd":                                   "Sml",
		"ref/vml-main.xsd":                              "VmMain",
		"ref/vml-officeDrawing.xsd":                     "VmlOfficedrawing",
		"ref/vml-presentationDrawing.xsd":               "VmlPresentationdrawing",
		"ref/vml-spreadsheetDrawing.xsd":                "VmlSpreadsheetdrawing",
		"ref/vml-wordprocessingDrawing.xsd":             "VmlWordprocessingdrawing",
		"ref/wml.xsd":                                   "Wml",
	}

	ParseTo("ref/wml.xsd", "Wml")
	//ParseTo("ref/shared-math.xsd", "SharedMath")
	return
	for i, k := range tmap {
		ParseTo(i, k)
	}

}
