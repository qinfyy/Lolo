package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type PlayerUnlock struct {
	all *excel.AllPlayerUnlockDatas
}

func (g *GameConfig) loadPlayerUnlock() {
	info := &PlayerUnlock{
		all: new(excel.AllPlayerUnlockDatas),
	}
	g.Excel.PlayerUnlock = info
	name := "PlayerUnlock.json"
	ReadJson(g.excelPath, name, &info.all)
}

func GetAllPlayerUnlock() *excel.AllPlayerUnlockDatas {
	return cc.Excel.PlayerUnlock.all
}
