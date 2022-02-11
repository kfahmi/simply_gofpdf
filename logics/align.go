package logics

import "github.com/jung-kurt/gofpdf"

// Return Coordinate X and Y depends on interface you choose
type Alignment struct {
	PdfObj *gofpdf.Fpdf
}
type IAlignment interface {
	Center(data interface{}) (*VarCoordinates, error)
	Left(data interface{}) (*VarCoordinates, error)
	Right(data interface{}) (*VarCoordinates, error)
}

func InitAlignment(pdf *gofpdf.Fpdf) *Alignment {
	service := Alignment{
		PdfObj: pdf,
	}
	return &service
}

func (service *Alignment) Center(data interface{}) (*VarCoordinates, error) {
	d := data.(*VarAlignmentCfg)

	centerContainerX := d.ContainerWidth / 2
	centerContainerY := d.ContainerHeigth / 2

	divideWidthText := d.TextWidth / 2
	divideHeightText := d.TextHeight / 2

	startX := centerContainerX - divideWidthText
	startY := centerContainerY - divideHeightText

	coord := &VarCoordinates{
		x: startX,
		y: startY,
	}
	return coord, nil
}

func (service *Alignment) Left(data interface{}) (*VarCoordinates, error) {
	d := data.(*VarAlignmentCfg)

	centerContainerY := d.ContainerHeigth / 2

	divideHeightText := d.TextHeight / 2

	startX := 0 + d.Margin
	startY := centerContainerY - divideHeightText

	coord := &VarCoordinates{
		x: startX,
		y: startY,
	}
	return coord, nil
}

func (service *Alignment) Right(data interface{}) (*VarCoordinates, error) {
	d := data.(*VarAlignmentCfg)

	centerContainerY := d.ContainerHeigth / 2

	divideHeightText := d.TextHeight / 2

	startX := d.ContainerWidth - d.TextWidth - d.Margin
	startY := centerContainerY - divideHeightText

	coord := &VarCoordinates{
		x: startX,
		y: startY,
	}
	return coord, nil
}
