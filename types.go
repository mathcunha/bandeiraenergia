package bandeira

type Color int

const (
	Red Color = iota
	Green
	Yellow
)

type Region int

const (
	Norte Region = iota
	Nordeste
	Sudeste
	CentroOeste
	Sul
)

type Status struct {
	Color
	Region
}
