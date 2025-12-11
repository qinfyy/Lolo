package game

import (
	"gucooing/lolo/db"
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) ChatUnLockExpressionNotice(s *model.Player) {
	notice := &proto.ChatUnLockExpressionNotice{
		Status:       proto.StatusCode_StatusCode_OK,
		ExpressionId: s.GetChatModel().GetUnLockExpression(),
	}
	defer g.send(s, 0, notice)
}

func (g *Game) PrivateChatOfflineNotice(s *model.Player) {
	notice := &proto.PrivateChatOfflineNotice{
		Status:     proto.StatusCode_StatusCode_OK,
		OfflineMsg: make([]*proto.PrivateChatOffline, 0),
	}
	defer g.send(s, 0, notice)
	privates, err := db.GetAllChatPrivate(s.UserId)
	if err != nil {
		notice.Status = proto.StatusCode_StatusCode_CHAT_CHANNEL_NOT_EXIST
		log.Game.Warnf("UserID:%v func db.GetAllChatPrivate err:%v", s.UserId, err)
		return
	}
	for _, private := range privates {
		alg.AddList(&notice.OfflineMsg, s.GetPrivateChatOffline(private))
	}
}

func (g *Game) ChatMsgRecordInitNotice(s *model.Player) {
	notice := &proto.ChatMsgRecordInitNotice{
		Status: proto.StatusCode_StatusCode_OK,
		Type:   0,
		Msg:    make([]*proto.ChatMsgData, 0),
	}
	defer g.send(s, 0, notice)
}

func (g *Game) PrivateChatMsgRecord(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.PrivateChatMsgRecordReq)
	rsp := &proto.PrivateChatMsgRecordRsp{
		Status:    proto.StatusCode_StatusCode_OK,
		MsgRecord: make([]*proto.ChatMsgData, 0),
	}
	defer g.send(s, msg.PacketId, rsp)
	// 好友判断
	if count, err := db.GetIsFiend(s.UserId, req.TargetPlayerId); err != nil {
		log.Game.Warnf("UserId:%v db.GetIsFiend err:%v", s.UserId, err)
		return
	} else if count == 0 {
		return
	}
	// 获取聊天内容
	privateMsgs, err := db.GetAllChatPrivateMsg(s.UserId, req.TargetPlayerId)
	if err != nil {
		log.Game.Warnf("UserId:%v db.GetAllChatPrivateMsg err:%v", s.UserId, err)
		return
	}
	for _, privateMsg := range privateMsgs {
		alg.AddList(&rsp.MsgRecord,
			model.GetUserChatMsgData(privateMsg.OFChatMsg, privateMsg.UserId))
	}
}

func (g *Game) SendChatMsg(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.SendChatMsgReq)
	rsp := &proto.SendChatMsgRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Text:   req.Text,
	}
	defer g.send(s, msg.PacketId, rsp)
	switch req.Type {
	case proto.ChatChannelType_ChatChannel_Default: // 默认消息是房间消息
	case proto.ChatChannelType_ChatChannel_ChatRoom: // 聊天房间
	case proto.ChatChannelType_ChatChannel_Private: // 私聊
		// 好友判断
		if count, err := db.GetIsFiend(s.UserId, req.PlayerId); err != nil {
			log.Game.Warnf("UserId:%v db.GetIsFiend err:%v", s.UserId, err)
			return
		} else if count == 0 {
			return
		}
		// 写入数据库
		// 如果在线就通知过去
	}
}
