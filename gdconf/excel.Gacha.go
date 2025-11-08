package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Gacha struct {
	all *excel.AllGachaDatas
}

func (g *GameConfig) loadGacha() {
	info := &Gacha{
		all: new(excel.AllGachaDatas),
	}
	g.Excel.Gacha = info
	name := "Gacha.json"
	ReadJson(g.excelPath, name, &info.all)
}

func GetAllGacha() *excel.AllGachaDatas {
	return cc.Excel.Gacha.all
}
