package fighter

import "fmt"

type Succubus struct {
	Id   FighterType
	Suck int
	Life int
	Mana int
}

func NewSuccubus(suck, life, mana int) Succubus {
	return Succubus{SuccubusType, suck, life, mana}
}

func (s Succubus) GetId()FighterType{
	return  s.Id
}
func (s Succubus) Clone() Fighter {
	return NewSuccubus(s.Suck, s.Life, s.Mana)
}

func (s Succubus) Fight()  {
	fmt.Println("🏀🏀 Succubus Sucking 💦 ",s.Suck)
}