package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) NpcTalk(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.NpcTalkReq)
	rsp := &proto.NpcTalkRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer g.send(s, cmd.NpcTalkRsp, msg.PacketId, rsp)
}

func (g *Game) Tutorial(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.TutorialReq)
	rsp := &proto.TutorialRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer g.send(s, cmd.TutorialRsp, msg.PacketId, rsp)
}
