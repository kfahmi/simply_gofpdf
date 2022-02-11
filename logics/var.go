package logics

type VarCoordinates struct {
	x float64
	y float64
}

type VarRGB struct {
	R int
	G int
	B int
}

type VarTextCfg struct {
	Alignment  string
	Text       string
	FontFamily string
	FontStyle  string
	FontSize   float64
	Margin     float64
	Color      VarRGB
}

type VarAlignmentCfg struct {
	Margin          float64
	TextWidth       float64
	TextHeight      float64
	ContainerWidth  float64
	ContainerHeigth float64
}

type VarDataGrid struct {
	Row    int
	Col    int
	Header bool
}
