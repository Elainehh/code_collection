package Creational

import (
	"testing"
)

func TestFactory(t *testing.T) {
	items := []string{"rectangle", "circle", "square"}
	for _, item := range items {
		var r1 Shape = createShape(item)
		if r1.draw() != item {
			t.Fatal("not right kind")
		}
	}

}
