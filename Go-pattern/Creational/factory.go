package Creational

type Shape interface {
	draw() string
}

type Rectangle struct {
	Shape
}

type Square struct {
	Shape
}

type Circle struct {
	Shape
}

func (r Rectangle) draw() string {
	return "rectangle"
}

func (s Square) draw() string {
	return "square"
}

func (c Circle) draw() string {
	return "circle"
}

func createShape(kind string) Shape {
	switch kind {
	case "rectangle":
		return Rectangle{}
	case "square":
		return &Square{}
	case "circle":
		return &Circle{}
	}

	return nil
}
