package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type PlayerUnlock struct {
	all             *excel.AllPlayerUnlockDatas
	PlayerUnlockMap map[int32]*excel.PlayerUnlockConfigure
}

func (g *GameConfig) loadPlayerUnlock() {
	info := &PlayerUnlock{
		all:             new(excel.AllPlayerUnlockDatas),
		PlayerUnlockMap: make(map[int32]*excel.PlayerUnlockConfigure),
	}
	g.Excel.PlayerUnlock = info
	name := "PlayerUnlock.json"
	ReadJson(g.excelPath, name, &info.all)
	for _, v := range info.all.GetPlayerUnlock().GetDatas() {
		info.PlayerUnlockMap[v.ID] = v
	}
}

func GetPlayerUnlockMap() map[int32]*excel.PlayerUnlockConfigure {
	return cc.Excel.PlayerUnlock.PlayerUnlockMap
}

func GetPlayerUnlockConfigure(id int32) *excel.PlayerUnlockConfigure {
	return cc.Excel.PlayerUnlock.PlayerUnlockMap[id]
}
