package logics

// Return Coordinate X and Y depends on interface you choose
type Alignment struct {
}
type IAlignment interface {
	Center(data interface{}) (VarCoordinates, error)
	Left(data interface{}) (VarCoordinates, error)
	Right(data interface{}) (VarCoordinates, error)
}

func InitAlignment() *Alignment {
	service := Alignment{}
	return &service
}

func (service *Alignment) Center(data interface{}) (VarCoordinates, error) {
	coord := VarCoordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}

func (service *Alignment) Left(data interface{}) (VarCoordinates, error) {
	coord := VarCoordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}

func (service *Alignment) Right(data interface{}) (VarCoordinates, error) {
	coord := VarCoordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}
