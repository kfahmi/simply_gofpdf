package logics

import (
	"github.com/jung-kurt/gofpdf"
)

type SFPDFText struct {
	alignment IAlignment
	pdfObj    *gofpdf.Fpdf
}
type ISPDFText interface {
	CreateSFPDFText(data interface{}) error
}

func InitSFPDFText(pdf *gofpdf.Fpdf) *SFPDFText {
	align := InitAlignment()
	service := SFPDFText{
		pdfObj:    pdf,
		alignment: align,
	}
	return &service
}

func (service *SFPDFText) CreateSFPDFText(data interface{}) error {
	coord := VarCoordinates{}
	d := data.(VarTextCfg)

	alignCfg := VarAlignmentCfg{
		Margin: d.Margin,
	}

	//FONT TEXT
	service.pdfObj.SetFont(d.FontFamily, d.FontStyle, d.FontSize)
	service.pdfObj.SetTextColor(d.Color.R, d.Color.G, d.Color.B)

	//Alignment Text
	switch d.Alignment {
	case "center":
		coord, _ = service.alignment.Center(alignCfg)
	case "left":
		coord, _ = service.alignment.Left(alignCfg)
	default:
		coord, _ = service.alignment.Right(alignCfg)
	}

	service.pdfObj.Text(coord.x, coord.y, d.Text)

	return nil
}
