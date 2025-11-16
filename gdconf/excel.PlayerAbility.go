package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type PlayerAbility struct {
	all              *excel.AllPlayerAbilityDatas
	PlayerAbilityMap map[int32]*excel.PlayerAbilityConfigure
}

func (g *GameConfig) loadPlayerAbility() {
	info := &PlayerAbility{
		all:              new(excel.AllPlayerAbilityDatas),
		PlayerAbilityMap: make(map[int32]*excel.PlayerAbilityConfigure),
	}
	g.Excel.PlayerAbility = info
	name := "PlayerAbility.json"
	ReadJson(g.excelPath, name, &info.all)
	for _, v := range info.all.GetPlayerAbility().GetDatas() {
		info.PlayerAbilityMap[v.ID] = v
	}
}

func GetPlayerAbilityMap() map[int32]*excel.PlayerAbilityConfigure {
	return cc.Excel.PlayerAbility.PlayerAbilityMap
}

func GetPlayerAbilityConfigure(id int32) *excel.PlayerAbilityConfigure {
	return cc.Excel.PlayerAbility.PlayerAbilityMap[id]
}
