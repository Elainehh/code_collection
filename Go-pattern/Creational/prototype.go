package Creational

const (
	rectangle = 1
	square    = 2
	circle    = 3
)

type Shape2 struct {
	ID   int
	Type string
}

func (s *Shape2) SetID(ID int) {
	s.ID = ID
}

func (s *Shape2) GetID() int {
	return s.ID
}

func Clone(input *Shape2) *Shape2 {
	return &Shape2{
		ID:   input.ID,
		Type: input.Type,
	}
}

type Rectangle2 struct {
	Shape *Shape2
}

func NewRectangle() *Rectangle2 {
	return &Rectangle2{
		Shape: &Shape2{
			Type: "Rectangle",
		},
	}
}

func LoadCache() map[string]interface{} {
	hash := map[string]interface{}{}
	rectangle := NewRectangle()
	rectangle.Shape.SetID(1)
	hash["rectangle"] = rectangle

	return hash
}

func GetShape(shape interface{}) interface{} {
	switch v := shape.(type) {
	case *Rectangle2:
		return &Rectangle2{Shape: Clone(v.Shape)}
	default:
		return nil
	}
}
