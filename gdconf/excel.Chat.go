package gdconf

import (
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/excel"
	"gucooing/lolo/protocol/proto"
)

type Chat struct {
	all                *excel.AllChatDatas
	ChatEmotionMap     map[int32]*excel.ChatEmotionConfigure
	ChatEmotionTypeMap map[proto.EEmotionType][]*excel.ChatEmotionConfigure
}

func (g *GameConfig) loadChat() {
	info := &Chat{
		all:                new(excel.AllChatDatas),
		ChatEmotionMap:     make(map[int32]*excel.ChatEmotionConfigure),
		ChatEmotionTypeMap: make(map[proto.EEmotionType][]*excel.ChatEmotionConfigure),
	}
	g.Excel.Chat = info
	name := "Chat.json"
	ReadJson(g.excelPath, name, &info.all)

	getChatEmotionType := func(t proto.EEmotionType) *[]*excel.ChatEmotionConfigure {
		list, ok := info.ChatEmotionTypeMap[t]
		if !ok {
			list = make([]*excel.ChatEmotionConfigure, 0)
			info.ChatEmotionTypeMap[t] = list
		}
		return &list
	}

	for _, v := range info.all.GetChatEmotion().GetDatas() {
		info.ChatEmotionMap[v.ID] = v

		et := getChatEmotionType(proto.EEmotionType(v.NewEmotionType))
		alg.AddList(et, v)
	}
}

func GetChatEmotionMap() map[int32]*excel.ChatEmotionConfigure {
	return cc.Excel.Chat.ChatEmotionMap
}

func GetChatEmotionType(t proto.EEmotionType) []*excel.ChatEmotionConfigure {
	return cc.Excel.Chat.ChatEmotionTypeMap[t]
}

func GetChatEmotion(id uint32) *excel.ChatEmotionConfigure {
	return cc.Excel.Chat.ChatEmotionMap[int32(id)]
}
