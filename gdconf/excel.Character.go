package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Character struct {
	all             *excel.AllCharacterDatas
	CharacterAllMap map[uint32]*CharacterAllInfo
}

type CharacterAllInfo struct {
	CharacterId   uint32
	CharacterInfo *excel.CharacterConfigure
	LevelRules    []*excel.CharacterLevelRuleInfo
}

func (g *GameConfig) loadCharacter() {
	info := &Character{
		all:             new(excel.AllCharacterDatas),
		CharacterAllMap: make(map[uint32]*CharacterAllInfo),
	}
	g.Excel.Character = info
	name := "Character.json"
	ReadJson(g.excelPath, name, &info.all)

	getCharacterAllMap := func(id int32) *CharacterAllInfo {
		if info.CharacterAllMap[uint32(id)] == nil {
			info.CharacterAllMap[uint32(id)] = &CharacterAllInfo{
				CharacterId: uint32(id),
			}
		}
		return info.CharacterAllMap[uint32(id)]
	}

	for _, v := range info.all.GetCharacter().GetDatas() {
		getCharacterAllMap(v.ID).CharacterInfo = v
	}
	for _, v := range info.all.GetLevelRule().GetDatas() {
		getCharacterAllMap(v.ID).LevelRules = v.LevelRuleInfo
	}
}

func GetCharacterAllMap() map[uint32]*CharacterAllInfo {
	return cc.Excel.Character.CharacterAllMap
}

func GetCharacterAll(id uint32) *CharacterAllInfo {
	return cc.Excel.Character.CharacterAllMap[id]
}
