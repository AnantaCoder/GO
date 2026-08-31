package fighter

import "fmt"

// Elf is another concrete Prototype.
type Elf struct {
	Id       FighterType
	ArrowHit int
	Life     int
	Mana     int
}

// NewElf constructor for Elf.
func NewElf(arrowHit, life, mana int) Elf {
	return Elf{ElfType, arrowHit, life, mana}
}

func (e Elf) GetId() FighterType {
	return e.Id
}

func (e Elf) Clone() Fighter {
	return NewElf(e.ArrowHit, e.Life, e.Mana)
}

func (e Elf) Fight() {
	fmt.Printf("🏹 Elf shoots with ArrowHit %d!\n", e.ArrowHit)
}
