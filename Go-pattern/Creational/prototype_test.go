package Creational

import "testing"

func TestPrototype(t *testing.T) {
	cache := LoadCache()

	if GetShape(cache["rectangle"]).(*Rectangle2).Shape.GetID() != 1 {
		t.Fatal("wrong ID")
	}
}
