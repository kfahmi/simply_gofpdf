package sgofpdf

import (
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/kfahmi/simply_gofpdf/logics"
)

type Init struct {
	Width          float64
	Heigth         float64
	OrientationStr string
	UnitStr        string
	PdfObj         *gofpdf.Fpdf
	FontPath       string
	sFPDFText      logics.ISPDFText
}

func (init *Init) Test() {
	fmt.Println("TestFunction")
}

//start docs init function sgofpdf
func (init *Init) InitSPDFDoc() error {
	fmt.Println("initiate InitSPDFDoc()")
	init.PdfObj = gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: init.OrientationStr,
		UnitStr:        init.UnitStr,
		Size:           gofpdf.SizeType{Wd: init.Width, Ht: init.Heigth},
		FontDirStr:     init.FontPath,
	})
	init.sFPDFText = logics.InitSFPDFText(init.PdfObj)
	return nil
}

// Other Function
func (init *Init) NewPage(cfg *logics.VarPageCfg) error {
	init.PdfObj.AddPage()
	if cfg.RullerPartition > 0 {
		init.PdfObj = init.DrawGridRuller(cfg.RullerPartition)
	}
	return nil
}

func (init *Init) DrawGridRuller(partition float64) *gofpdf.Fpdf {
	pdf := init.PdfObj
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 36)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / partition) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / partition) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
	return pdf
}

func (init *Init) RegisterFont(fontList []string) error {
	for _, fontType := range fontList {
		lower := strings.ToLower(fontType)
		registerName := strings.Split(lower, ".")[0]
		init.PdfObj.AddUTF8Font(registerName, "", fontType)
	}
	return nil
}

func (init *Init) WriteText(data interface{}) error {
	err := init.sFPDFText.CreateSFPDFText(data)
	return err
}

func (init *Init) RenderImage(data interface{}) error {
	return nil
}

func (init *Init) GeneratePDF(path string) error {
	err := init.PdfObj.OutputFileAndClose(path)
	return err
}
