package Creational

import (
	"reflect"
	"testing"
)

//func isMapEqual(m1 map,m2 map) bool {
//	if(len(m1)!=len(m2)) {
//		return false;
//	}
//
//}

func TestSingleton(t *testing.T) {
	maps1 := New()
	maps1["key1"] = "value1"
	maps2 := New()
	if !reflect.DeepEqual(maps1, maps2) {
		t.Errorf("not one instance")
	}
}
