package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Quest struct {
	all *excel.AllQuestDatas
}

func (g *GameConfig) loadQuest() {
	info := &Quest{
		all: new(excel.AllQuestDatas),
	}
	g.Excel.Quest = info
	name := "Quest.json"
	ReadJson(g.excelPath, name, &info.all)
}

func GetAllQuest() *excel.AllQuestDatas {
	return cc.Excel.Quest.all
}
