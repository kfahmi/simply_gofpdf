package logics

// Return Coordinate X and Y depends on interface you choose
type Alignment struct {
}
type IAlignment interface {
	Center(data interface{}) (Coordinates, error)
	Left(data interface{}) (Coordinates, error)
	Right(data interface{}) (Coordinates, error)
}

func InitAlignment() *Alignment {
	service := Alignment{}
	return &service
}

func (service *Alignment) Center(data interface{}) (Coordinates, error) {
	coord := Coordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}

func (service *Alignment) Left(data interface{}) (Coordinates, error) {
	coord := Coordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}

func (service *Alignment) Right(data interface{}) (Coordinates, error) {
	coord := Coordinates{
		x: 10,
		y: 10,
	}
	return coord, nil
}
