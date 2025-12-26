package model

import (
	"gucooing/lolo/db"
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

func (s *Player) GetPbPlayerDropRateInfo() *proto.PlayerDropRateInfo {
	info := &proto.PlayerDropRateInfo{
		KillDropRate:     100,
		TreasureDropRate: 100,
	}
	return info
}

func (s *Player) GetUnlockFunctions() []uint32 {
	list := make([]uint32, 0)
	for _, v := range gdconf.GetPlayerUnlockMap() {
		list = append(list, uint32(v.ID))
	}
	return list
}

func GetPlayerAppearance(userId uint32) *proto.PlayerAppearance {
	basic, err := db.GetGameBasic(userId)
	if err != nil {
		log.Game.Errorf("userId:%v func db.GetGameBasic err:%v", userId, err)
		return nil
	}
	return &proto.PlayerAppearance{
		AvatarFrame: basic.AvatarFrame,
		Pendant:     basic.Pendant,
	}
}
