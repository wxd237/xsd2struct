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
		"ref/dml-chart.xsd":                             "refgo/dml-chart.go",
		"ref/dml-chartDrawing.xsd":                      "refgo/dml-chartDrawing.go",
		"ref/dml-diagram.xsd":                           "refgo/dml-diagram.go",
		"ref/dml-lockedCanvas.xsd":                      "refgo/dml-lockedCanvas.go",
		"ref/dml-main.xsd":                              "refgo/dml-main.go",
		"ref/dml-picture.xsd":                           "refgo/dml-picture.go",
		"ref/dml-spreadsheetDrawing.xsd":                "refgo/dml-spreadsheetDrawing.go",
		"ref/dml-wordprocessingDrawing.xsd":             "refgo/dml-wordprocessingDrawing.go",
		"ref/pml.xsd":                                   "refgo/pml.go",
		"ref/shared-additionalCharacteristics.xsd":      "refgo/shared-additionalCharacteristics.go",
		"ref/shared-bibliography.xsd":                   "refgo/shared-bibliography.go",
		"ref/shared-commonSimpleTypes.xsd":              "refgo/shared-commonSimpleTypes.go",
		"ref/shared-customXmlDataProperties.xsd":        "refgo/shared-customXmlDataProperties.go",
		"ref/shared-customXmlSchemaProperties.xsd":      "refgo/shared-customXmlSchemaProperties.go",
		"ref/shared-documentPropertiesCustom.xsd":       "refgo/shared-documentPropertiesCustom.go",
		"ref/shared-documentPropertiesExtended.xsd":     "refgo/shared-documentPropertiesExtended.go",
		"ref/shared-documentPropertiesVariantTypes.xsd": "refgo/shared-documentPropertiesVariantTypes.go",
		"ref/shared-math.xsd":                           "refgo/shared-math.go",
		"ref/shared-relationshipReference.xsd":          "refgo/shared-relationshipReference.go",
		"ref/sml.xsd":                                   "refgo/sml.go",
		"ref/vml-main.xsd":                              "refgo/vml-main.go",
		"ref/vml-officeDrawing.xsd":                     "refgo/vml-officeDrawing.go",
		"ref/vml-presentationDrawing.xsd":               "refgo/vml-presentationDrawing.go",
		"ref/vml-spreadsheetDrawing.xsd":                "refgo/vml-spreadsheetDrawing.go",
		"ref/vml-wordprocessingDrawing.xsd":             "refgo/vml-wordprocessingDrawing.go",
		"ref/wml.xsd":                                   "refgo/wml.go",
	}

	for i, k := range tmap {
		ParseTo(i, k)
	}

}
