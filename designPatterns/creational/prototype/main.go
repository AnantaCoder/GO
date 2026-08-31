package main

import "main/fighter"

var dataHistoryList map[int]fighter.Fighter

func loadCache() {
	dataHistoryList = make(map[int]fighter.Fighter)

	soldier := fighter.NewSoldier(30, 20, 5)
	dwarf := fighter.NewDwarf(50, 40, 15)
	succubus := fighter.NewSuccubus(70, 20, 30)

	dataHistoryList[int(soldier.GetId())] = soldier
	dataHistoryList[int(dwarf.GetId())] = dwarf
	dataHistoryList[int(succubus.GetId())] = succubus
}

func main() {
	loadCache()

	// 1) Soldier Clone
	if proto, found := dataHistoryList[int(fighter.SolderType)]; found {
		if s, ok := proto.(fighter.Soldier); ok {
			clone := s.Clone().(fighter.Soldier)
			clone.GunHit = 45
			clone.Fight()
		}
	}

	// 2) Elf Handling
	if proto, found := dataHistoryList[int(fighter.ElfType)]; found {
		e := proto.Clone()
		e.Fight()
	} else {
		newElf := fighter.NewElf(15, 30, 3)
		dataHistoryList[int(newElf.GetId())] = newElf
		newElf.Fight()
	}

	if proto, found := dataHistoryList[int(fighter.ElfType)]; found {
		if e2, ok := proto.(fighter.Elf); ok {
			clone2 := e2.Clone().(fighter.Elf)
			clone2.ArrowHit = 35
			clone2.Fight()
		}
	}

	// 3) Dwarf Handling
	if proto, found := dataHistoryList[int(fighter.DwarfType)]; found {
		d := proto.Clone()
		d.Fight()
	} else {
		newDwarf := fighter.NewDwarf(50, 40, 15)
		dataHistoryList[int(newDwarf.GetId())] = newDwarf
		newDwarf.Fight()
	}

	// 4) Succubus Handling
	if proto, found := dataHistoryList[int(fighter.SuccubusType)]; found {
		if s, ok := proto.(fighter.Succubus); ok {
			clone := s.Clone().(fighter.Succubus)
			clone.Suck = 999 // 😈 Modified power
			clone.Fight()
		}
	} else {
		newSuccubus := fighter.NewSuccubus(70, 20, 30)
		dataHistoryList[int(newSuccubus.GetId())] = newSuccubus
		newSuccubus.Fight()
	}
}
