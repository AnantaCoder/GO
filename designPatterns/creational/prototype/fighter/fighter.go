package fighter

//categorical definition id
type FighterType int

const (
	SolderType   FighterType = 1
	ElfType      FighterType = 2
	DwarfType    FighterType = 3
	SuccubusType FighterType = 4
)

// Fighter is the Prototype interface: every fighter can GetId, Fight, and Clone itself.
type Fighter interface {
	GetId() FighterType
	Fight()
	Clone() Fighter
}
