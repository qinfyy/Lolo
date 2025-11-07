package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GetAllCharacterEquip(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GetAllCharacterEquipRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Items:  make([]*proto.ItemDetail, 0),
	}
	defer g.send(s, cmd.GetAllCharacterEquipRsp, msg.PacketId, rsp)
}

func (g *Game) GetCharacterAchievementList(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.GetCharacterAchievementListReq)
	rsp := &proto.GetCharacterAchievementListRsp{
		Status:                  proto.StatusCode_StatusCode_OK,
		CharacterAchievementLst: make([]*proto.Achieve, 0),
		HasRewardedIds:          make([]uint32, 0),
		IsUnlockedPayment:       false,
		CharacterId:             req.CharacterId,
		RewardedIdLst:           make([]uint32, 0),
	}
	defer g.send(s, cmd.GetCharacterAchievementListRsp, msg.PacketId, rsp)
}

func (g *Game) OutfitPresetUpdate(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.OutfitPresetUpdateReq)
	rsp := &proto.OutfitPresetUpdateRsp{
		Status: proto.StatusCode_StatusCode_OK,
		CharId: req.CharId,
		Preset: req.Preset,
	}
	defer g.send(s, cmd.OutfitPresetUpdateRsp, msg.PacketId, rsp)
	characterInfo := s.GetCharacterInfo(req.CharId)
	if characterInfo == nil {
		log.Game.Warnf("保存角色预设装扮失败,角色%v不存在", req.CharId)
		return
	}
	outfitPreset := s.GetOutfitPreset(characterInfo, req.Preset.PresetIndex)

	outfitPreset.Hair = req.Preset.Hair
	outfitPreset.Hair = req.Preset.Hair
	outfitPreset.Clothes = req.Preset.Clothes
	outfitPreset.Ornament = req.Preset.Ornament
	outfitPreset.HatDyeSchemeIndex = req.Preset.HatDyeSchemeIndex
	outfitPreset.HairDyeSchemeIndex = req.Preset.HairDyeSchemeIndex
	outfitPreset.ClothesDyeSchemeIndex = req.Preset.ClothesDyeSchemeIndex
	outfitPreset.OrnamentDyeSchemeIndex = req.Preset.OrnamentDyeSchemeIndex
	outfitPreset.OutfitHideInfo = &model.OutfitHideInfo{
		HideOrn:   req.Preset.OutfitHideInfo.HideOrn,
		HideBraid: req.Preset.OutfitHideInfo.HideBraid,
	}
}
