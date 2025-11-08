package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GetArchiveInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.GetArchiveInfoReq)
	rsp := &proto.GetArchiveInfoRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Key:    req.Key,
		Value:  "",
	}
	g.send(s, cmd.GetArchiveInfoRsp, msg.PacketId, rsp)
}

func (g *Game) SetArchiveInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.SetArchiveInfoReq)
	rsp := &proto.SetArchiveInfoRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Key:    req.Key,
		Value:  req.Value,
	}
	g.send(s, cmd.SetArchiveInfoRsp, msg.PacketId, rsp)
}
