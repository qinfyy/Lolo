package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GachaList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GachaListRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Gachas: make([]*proto.GachaInfo, 0),
	}
	defer g.send(s, cmd.GachaListRsp, msg.PacketId, rsp)
	for _, v := range gdconf.GetAllGacha().GetInfo().GetDatas() {
		info := &proto.GachaInfo{
			GachaId:        uint32(v.ID),
			GachaTimes:     0,
			HasFullPick:    false,
			IsFree:         false,
			OptionalUpItem: 0,
			OptionalValue:  0,
			Guarantee:      0,
		}
		alg.AddList(&rsp.Gachas, info)
	}
}
