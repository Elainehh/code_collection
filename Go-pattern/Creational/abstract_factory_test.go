package Creational

import "testing"

func TestAbstractFactory(t *testing.T) {
	s := ShapeWithColor{}
	shapes := []string{"rectangle", "circle", "square"}
	for _, shape := range shapes {
		r1 := s.getShape(shape)
		if r1.draw() != shape {
			t.Fatal("not right kind")
		}
	}

	colors := []string{"red", "yellow", "blue"}
	for _, color := range colors {
		c := s.getColor(color)
		if c.fill() != color {
			t.Fatal("not right color")
		}
	}
}
