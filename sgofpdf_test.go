package sgofpdf_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/jung-kurt/gofpdf"
	sgofpdf "github.com/kfahmi/simply_gofpdf"
	sgofpdf_logics "github.com/kfahmi/simply_gofpdf/logics"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	rootdir, _ := os.Getwd()
	fontPath := path.Join(rootdir, "font")
	simplypdf := sgofpdf.Init{
		Heigth:         800,
		Width:          800,
		OrientationStr: "P",
		UnitStr:        "mm",
		FontPath:       fontPath,
	}
	simplypdf.InitSPDFDoc()
	//register font
	fontList := []string{
		"Birthstone-Regular.ttf",
	}
	simplypdf.RegisterFont(fontList)

	//page1
	pageCfg := &sgofpdf_logics.VarPageCfg{
		RullerPartition: 20,
	}
	simplypdf.NewPage(pageCfg)
	data := &sgofpdf_logics.VarTextCfg{
		Alignment:  "left",
		Text:       "SimpyGofpdf by Kfahmi",
		FontFamily: "birthstone-regular",
		FontStyle:  "",
		FontSize:   200,
		Margin:     10,
		Color: sgofpdf_logics.VarRGB{
			R: 0,
			G: 0,
			B: 0,
		},
	}
	simplypdf.WriteText(data)

	//page3
	pageCfg = &sgofpdf_logics.VarPageCfg{
		RullerPartition: 10,
	}
	simplypdf.NewPage(pageCfg)
	data = &sgofpdf_logics.VarTextCfg{
		Alignment:  "center",
		Text:       "SimpyGofpdf by Kfahmi",
		FontFamily: "birthstone-regular",
		FontStyle:  "",
		FontSize:   200,
		Margin:     10,
		Color: sgofpdf_logics.VarRGB{
			R: 0,
			G: 0,
			B: 0,
		},
	}
	simplypdf.WriteText(data)

	//GENERATE PDF
	filePath := path.Join(rootdir, "files")
	prefixFile := fmt.Sprintf("simpy-gofpdf-%s", "test")
	pdfName := fmt.Sprintf("%s.%s", prefixFile, "pdf")
	pdfPath := path.Join(filePath, pdfName)
	err := simplypdf.GeneratePDF(pdfPath)

	assert.Nil(t, err)
}

func TestCreateFileGofpdf(t *testing.T) {
	rootdir, _ := os.Getwd()
	fontPath := path.Join(rootdir, "font")
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 842, Ht: 595}, //282-211(3,2.8)
		FontDirStr:     fontPath,
	})

	//page 1
	pdf.AddPage()
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 36)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	//ruller
	for x := 0.0; x < w; x = x + (w / 20) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
	//font
	pdf.AddUTF8Font("birthStoneRegular", "", "Birthstone-Regular.ttf")
	pdf.SetFont("birthStoneRegular", "", 144) //figma 48px
	pdf.SetTextColor(51, 51, 51)
	pdf.Text(0, 400, "GOFPDF Package manual left")

	//page2 text auto center
	pdf.AddPage()
	w, h = pdf.GetPageSize()
	pdf.SetFont("courier", "", 36)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	//ruller
	for x := 0.0; x < w; x = x + (w / 20) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

	//TEXT
	pdf.AddUTF8Font("birthStoneRegular", "", "Birthstone-Regular.ttf")
	pdf.SetFont("birthStoneRegular", "", 144) //figma 48px
	pdf.SetTextColor(51, 51, 51)

	//htung luas container page
	cWidth, cHeigth := pdf.GetPageSize()
	centerContainerX := cWidth / 2
	centerContainerY := cHeigth / 2

	//hitung lebar tinggi text
	text := "GOFPDF Package CENTER text"
	textWidth := pdf.GetStringWidth(text)
	_, textHeight := pdf.GetFontSize()

	//bagi dua lebar dan tinggi buat nyari center
	divideWidthText := textWidth / 2
	divideHeightText := textHeight / 2
	//tentuin start x dan y biar center
	startX := centerContainerX - divideWidthText
	startY := centerContainerY - divideHeightText

	pdf.Text(startX, startY, text)

	//GENERATE PDF
	filePath := path.Join(rootdir, "files")
	prefixFile := fmt.Sprintf("gofpdf-%s", "test")
	pdfName := fmt.Sprintf("%s.%s", prefixFile, "pdf")
	pdfPath := path.Join(filePath, pdfName)
	err := pdf.OutputFileAndClose(pdfPath)
	assert.Nil(t, err)
}
