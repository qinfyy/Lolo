package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) BattleEncounterInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.BattleEncounterInfoReq)
	rsp := &proto.BattleEncounterInfoRsp{
		Status:     proto.StatusCode_StatusCode_Ok,
		Encounters: make([]*proto.BattleEncounterData, 0),
	}
	defer g.send(s, msg.PacketId, rsp)
	for _, encounterId := range req.EncounterIds {
		alg.AddList(&rsp.Encounters, &proto.BattleEncounterData{
			BattleId: encounterId,
			State:    proto.BattleState_BattleState_Start,
			BoxId:    300000,
		})
	}
}

func (g *Game) BattleEncounterStateUpdate(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.BattleEncounterStateUpdateReq)
	rsp := &proto.BattleEncounterStateUpdateRsp{
		Status:                     proto.StatusCode_StatusCode_Ok,
		Encounter:                  nil,
		DynamicTreasureBoxBaseInfo: new(proto.DynamicTreasureBoxBaseData),
	}
	defer g.send(s, msg.PacketId, rsp)
	rsp.Encounter = &proto.BattleEncounterData{
		BattleId: req.EncounterId,
		State:    req.BattleState,
		BoxId:    300000,
	}
}

func (g *Game) MonsterDead(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.MonsterDeadReq)
	rsp := &proto.MonsterDeadRsp{
		Status:       proto.StatusCode_StatusCode_Ok,
		MonsterIndex: req.MonsterIndex,
		DropItem:     nil,
	}
	defer g.send(s, msg.PacketId, rsp)
	rsp.DropItem = &proto.DropItem{
		Index: req.MonsterIndex,
		Items: make([]*proto.ItemDetail, 0),
	}
}
