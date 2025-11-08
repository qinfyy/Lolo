package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Head struct {
	all *excel.AllHeadDatas
}

func (g *GameConfig) loadHead() {
	info := &Head{
		all: new(excel.AllHeadDatas),
	}
	g.Excel.Head = info
	name := "Head.json"
	ReadJson(g.excelPath, name, &info.all)
}
