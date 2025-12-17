package model

import (
	"gucooing/lolo/protocol/proto"
)

const (
	ChannelTypePublic = iota
	ChannelTypePrivate
)

// 房间花园数据
type SceneGardenData struct {
	GardenName             string                                `json:"gardenName;omitempty"`             // 花园名称
	LikesNum               int64                                 `json:"likesNum;omitempty"`               // 点赞数
	AccessPlayerNum        int64                                 `json:"accessPlayerNum;omitempty"`        // 访问数
	GardenFurnitureInfoMap map[int64]*proto.FurnitureDetailsInfo `json:"gardenFurnitureInfoMap;omitempty"` // 花园家具信息
}

func GetSceneGardenData(userId, sceneId uint32, channelType int) *SceneGardenData {
	if sceneId != 9999 || channelType == ChannelTypePublic {
		return &SceneGardenData{}
	}
	return &SceneGardenData{}
}

func (s *SceneGardenData) SceneGardenData() *proto.SceneGardenData {
	info := &proto.SceneGardenData{
		GardenFurnitureInfoMap:        make(map[int64]*proto.FurnitureDetailsInfo),
		LikesNum:                      0,
		AccessPlayerNum:               0,
		LeftLikeNum:                   0,
		GardenName:                    "",
		FurniturePlayerMap:            make(map[int64]uint32),
		OtherPlayerFurnitureInfoMap:   make(map[int64]*proto.SceneGardenOtherPlayerData),
		FurnitureCurrentPointNum:      0,
		PlayerHandingFurnitureInfoMap: make(map[int64]*proto.SceneGardenOtherPlayerData),
	}
	return info
}
