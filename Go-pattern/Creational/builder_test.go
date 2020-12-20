package Creational

import "testing"

func TestBuilder(t *testing.T) {
	m := Manufacturer{}

	homePCBuilder := HomeEditionPCBuilder{}
	m.SetBuilder(&homePCBuilder)

	m.constructPC()
	homePC := m.builder.GetPC()
	if homePC.os != "Win 7" {
		t.Fatal("Wrong CPU")
	}

	// TODO add checks for OS/CPU

	gamePCBuilder := GameEditionPCBuilder{}
	m.SetBuilder(&gamePCBuilder)
	m.constructPC()
	gamePC := m.builder.GetPC()
	if gamePC.diskType != "ssd" {
		t.Errorf("Wrong type")
	}

}
