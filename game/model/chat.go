package model

import (
	"gucooing/lolo/db"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

type ChatModel struct {
	UnLockExpression []uint32 // 已解锁的表情
}

func DefaultChatModel() *ChatModel {
	info := &ChatModel{
		UnLockExpression: make([]uint32, 0),
	}
	return info
}

func (s *Player) GetChatModel() *ChatModel {
	if s.Chat == nil {
		s.Chat = DefaultChatModel()
	}
	return s.Chat
}

func (c *ChatModel) GetUnLockExpression() []uint32 {
	if c.UnLockExpression == nil {
		c.UnLockExpression = make([]uint32, 0)
	}
	return c.UnLockExpression
}

func (s *Player) GetPrivateChatOffline(private *db.OFChatPrivate) *proto.PrivateChatOffline {
	userId := private.GetSubUserID(s.UserId)
	basic, err := db.GetGameBasic(userId)
	if err != nil {
		log.Game.Warnf("UserId:%v func db.GetGameBasic err:%v", userId, err)
		return nil
	}
	return &proto.PrivateChatOffline{
		PlayerId:    basic.UserId,
		Name:        basic.NickName,
		Head:        basic.Head,
		IsNewMsg:    private.IsNewMsg,
		AvatarFrame: basic.AvatarFrame,
	}
}

func GetUserChatMsgData(chatMsg *db.OFChatMsg, userId uint32) *proto.ChatMsgData {
	basic, err := db.GetGameBasic(userId)
	if err != nil {
		log.Game.Warnf("UserId:%v func db.GetGameBasic err:%v", userId, err)
		return nil
	}
	return &proto.ChatMsgData{
		PlayerId:    basic.UserId,
		Head:        basic.Head,
		Badge:       basic.Head,
		Name:        basic.NickName,
		Text:        chatMsg.Text,
		Expression:  chatMsg.Expression,
		SendTime:    chatMsg.SendTime,
		AvatarFrame: basic.AvatarFrame,
	}
}
