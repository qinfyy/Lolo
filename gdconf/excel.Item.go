package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Item struct {
	all     *excel.AllItemDatas
	ItemMap map[uint32]*excel.ItemConfigure
}

func (g *GameConfig) loadItem() {
	info := &Item{
		all:     new(excel.AllItemDatas),
		ItemMap: make(map[uint32]*excel.ItemConfigure),
	}
	g.Excel.Item = info
	name := "Item.json"
	ReadJson(g.excelPath, name, &info.all)
	for _, v := range info.all.GetItem().GetDatas() {
		info.ItemMap[uint32(v.ID)] = v
	}
}

func GetItemConfigure(id uint32) *excel.ItemConfigure {
	return cc.Excel.Item.ItemMap[id]
}
