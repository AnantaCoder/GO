package fighter

import "fmt"

type Dwarf struct {
	Id       FighterType
	AxePoint int
	Life     int
	Mana     int
}

func NewDwarf(axePoint, life, mana int) Dwarf {
	return Dwarf{DwarfType, axePoint, life, mana}
}

func (d Dwarf) GetId() FighterType {
	return d.Id
}

func (d Dwarf) Clone() Fighter {
	return NewDwarf(d.AxePoint, d.Life, d.Mana)
}

func (d Dwarf) Fight() {
	fmt.Printf("🪓 Dwarf swings with AxePoint %d!\n", d.AxePoint)
}
