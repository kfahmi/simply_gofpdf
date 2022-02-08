package logics

type Coordinates struct {
	x float64
	y float64
}

type RGB struct {
	R int
	G int
	B int
}

type TextCfg struct {
	Alignment  string
	Text       string
	FontFamily string
	FontStyle  string
	FontSize   float64
	Margin     float64
	Color      RGB
}

type AlignmentCfg struct {
	Margin float64
}
