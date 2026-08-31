package fighter

import "fmt"

type Soldier struct {
	Id     FighterType
	GunHit int
	Life   int
	Mana   int
}

// constructor for clone

func NewSoldier(gunhit, life, mana int) Soldier {
	return Soldier{SolderType, gunhit, life, mana}
}

// GetId returns the enum ID for this prototype.
func (s Soldier) GetId() FighterType {
	return s.Id
}

// Clone makes a deep copy by calling NewSoldier with the same data.
func (s Soldier) Clone() Fighter {
	return NewSoldier(s.GunHit, s.Life, s.Mana)
}

// Fight is just a demo method to show the fighter in action.
func (s Soldier) Fight() {
	fmt.Printf("💥 Soldier attacks with GunHit %d!\n", s.GunHit)
}