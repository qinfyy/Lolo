package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Shop struct {
	all  *excel.AllShopDatas
	Shop map[uint32]*excel.ShopConfigure
	Grid map[uint32]map[uint32]*excel.ShopGridConfigureItem
	Pool map[uint32]map[uint32]*excel.ShopPoolConfigureItem
}

func (g *GameConfig) loadShop() {
	info := &Shop{
		all:  new(excel.AllShopDatas),
		Shop: make(map[uint32]*excel.ShopConfigure),
		Grid: make(map[uint32]map[uint32]*excel.ShopGridConfigureItem),
		Pool: make(map[uint32]map[uint32]*excel.ShopPoolConfigureItem),
	}
	g.Excel.Shop = info
	name := "Shop.json"
	ReadJson(g.excelPath, name, &info.all)
	for _, v := range info.all.GetShop().GetDatas() {
		info.Shop[uint32(v.ID)] = v
	}

	for _, v := range info.all.GetGrid().GetDatas() {
		grid := make(map[uint32]*excel.ShopGridConfigureItem)
		for _, v2 := range v.GetItems() {
			grid[uint32(v2.GridID)] = v2
		}
		info.Grid[uint32(v.ID)] = grid
	}

	for _, v := range info.all.GetPool().GetDatas() {
		pool := make(map[uint32]*excel.ShopPoolConfigureItem)
		for _, v2 := range v.GetItems() {
			pool[uint32(v2.Index)] = v2
		}
		info.Pool[uint32(v.ID)] = pool
	}
}
