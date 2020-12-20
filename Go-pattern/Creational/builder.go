package Creational

type PersonalComputer struct {
	cpu      string
	os       string
	diskType string
}

type PCBuilder interface {
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetDiskType() PCBuilder
	GetPC() PersonalComputer
}

// It should implement all interfaces of PCBuilder
type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

func (h *HomeEditionPCBuilder) SetCPU() PCBuilder {
	h.pc.cpu = "i5"
	return h
}

func (h *HomeEditionPCBuilder) SetOS() PCBuilder {
	h.pc.os="Win 7"
	return h
}

func (h *HomeEditionPCBuilder) SetDiskType() PCBuilder {
	h.pc.diskType = "hdd"
	return h
}

func (h *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return h.pc
}

type GameEditionPCBuilder struct {
	pc PersonalComputer
}

func (g *GameEditionPCBuilder) SetCPU() PCBuilder {
	g.pc.cpu = "i7"
	return g
}

func (g *GameEditionPCBuilder) SetOS() PCBuilder {
	g.pc.os = "Win 10"
	return g
}

func (g *GameEditionPCBuilder) SetDiskType() PCBuilder {
	g.pc.diskType = "ssd"
	return g
}

func (g *GameEditionPCBuilder) GetPC() PersonalComputer {
	return g.pc
}

type Manufacturer struct {
	builder PCBuilder
}

func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.builder = builder
}

func (m *Manufacturer) constructPC() {
	m.builder.SetCPU().SetOS().SetDiskType()
}

