package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type GlobalConfig struct {
	all             *excel.AllGlobalConfigDatas
	GlobalConfigKey map[string]*excel.GlobalConfigConfigure
}

func (g *GameConfig) loadGlobalConfig() {
	info := &GlobalConfig{
		all:             new(excel.AllGlobalConfigDatas),
		GlobalConfigKey: make(map[string]*excel.GlobalConfigConfigure),
	}
	g.Excel.GlobalConfig = info
	name := "GlobalConfig.json"
	ReadJson(g.excelPath, name, &info.all)

	for _, v := range info.all.GetConfig().GetDatas() {
		info.GlobalConfigKey[v.Key] = v
	}
}

func GetGlobalConfigConfigure(key string) *excel.GlobalConfigConfigure {
	return cc.Excel.GlobalConfig.GlobalConfigKey[key]
}
