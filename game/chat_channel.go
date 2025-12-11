package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/protocol/proto"
)

// 初始化玩家聊天
func (g *Game) chatInit(s *model.Player) {
	// 同步解锁的表情
	g.ChatUnLockExpressionNotice(s)
	// 获取私聊情况
	g.PrivateChatOfflineNotice(s)
	// 公共聊天频道
	// - 系统频道
	// -
	// proto.ChangeChatChannelNotice

	/*
		s.ChangeChatChannel()
	*/
	g.ChatMsgRecordInitNotice(s)
}

type ChatInfo struct {
	noticeChan    *ChatChannel            // 通知频道
	privateChat   *ChatChannel            // 私聊频道
	allSystemChat map[uint32]*ChatChannel // 系统频道
}

func (g *Game) getChatInfo() *ChatInfo {
	if g.chatInfo == nil {
		chatInfo := &ChatInfo{
			noticeChan:    nil,
			allSystemChat: make(map[uint32]*ChatChannel),
		}
		g.chatInfo = chatInfo
	}
	return g.chatInfo
}

// 获取通知频道
func (c *ChatInfo) getNoticeChan() *ChatChannel {
	return c.noticeChan
}

// ChatChannel 聊天房间对象
type ChatChannel struct {
	sendMsgChan chan *proto.SendChatMsgReq // 发送消息通道
}

func newChatChannel() *ChatChannel {
	info := &ChatChannel{
		sendMsgChan: make(chan *proto.SendChatMsgReq, 100),
	}
	return info
}

func (c *ChatChannel) SendMsg(msg *proto.SendChatMsgReq) {}
