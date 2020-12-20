package Creational

type ShapeWithColor struct{}

type Shape1 interface {
	draw() string
}

type Color interface {
	fill() string
}

type Rectangle1 struct{}
type Square1 struct {
}
type Circle1 struct{}

type Red struct{}
type Yellow struct{}
type Blue struct{}

func (r Rectangle1) draw() string {
	return "rectangle1"
}

func (s Square1) draw() string {
	return "square1"
}

func (c Circle1) draw() string {
	return "circle1"
}

func (r Red) fill() string {
	return "red"
}

func (y Yellow) fill() string {
	return "yellow"
}

func (b Blue) fill() string {
	return "blue"
}

func (s ShapeWithColor) getShape(shapeKind string) Shape {
	switch shapeKind {
	case "rectangle":
		return &Rectangle{}
	case "square":
		return &Square{}
	case "circle":
		return &Circle{}
	default:
		return nil
	}
}

func (s ShapeWithColor) getColor(colrType string) Color {
	switch colrType {
	case "red":
		return Red{}
	case "yellow":
		return Yellow{}
	case "blue":
		return &Blue{}
	default:
		return nil
	}
}
