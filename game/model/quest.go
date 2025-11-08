package model

import (
	"gucooing/lolo/protocol/proto"
)

func (s *Player) GetQuestDetail() *proto.QuestDetail {
	info := &proto.QuestDetail{
		Chapters:                make([]*proto.Chapter, 0),
		DailyQuestBonusDayLeft:  nil,
		DailyQuestBonusWeekLeft: nil,
		RandomQuestBonusLeft:    nil,
		Quests:                  make([]*proto.Quest, 0),
	}
	// for _, v := range gdconf.GetAllQuest().GetQuest().GetDatas() {
	// 	quest := &proto.Quest{
	// 		QuestId:       uint32(v.ID),
	// 		Conditions:    make([]*proto.Condition, 0),
	// 		Status:        proto.QuestStatus_QuestStatus_InProgress,
	// 		CompleteCount: 1,
	// 		BonusTimes:    1,
	// 		ActivityId:    0,
	// 	}
	// 	for _, v2 := range gdconf.GetAllQuest().GetConditionSetGroup().GetDatas() {
	// 		if v2.ID == v.ID {
	// 			for _, set := range v2.QuestConditionSet {
	// 				for _, conditionId := range set.GetAchieveConditionID() {
	// 					alg.AddList(&quest.Conditions, &proto.Condition{
	// 						ConditionId: uint32(conditionId),
	// 						Progress:    1,
	// 						Status:      proto.QuestStatus_QuestStatus_InProgress,
	// 					})
	// 				}
	// 			}
	// 			break
	// 		}
	// 	}
	// 	alg.AddList(&info.Quests, quest)
	// }
	return info
}

func (s *Player) GetPlayerQuestionnaireInfo() *proto.PlayerQuestionnaireInfo {
	return &proto.PlayerQuestionnaireInfo{
		ToFill: make([]*proto.QuestionnaireBrief, 0),
	}
}
