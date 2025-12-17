package db

import (
	"time"

	"gorm.io/gorm"
)

// 聊天消息基础结构
type OFChatMsg struct {
	SendTime   int64 `gorm:"not null"`
	Text       string
	Expression uint32
}

// 玩家私聊数据库
type OFChatPrivate struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	UserID1   uint32    `gorm:"uniqueIndex:user"`
	UserID2   uint32    `gorm:"uniqueIndex:user"`
	IsNewMsg  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (p *OFChatPrivate) GetSubUserID(mainUserId uint32) uint32 {
	if p.UserID1 == mainUserId {
		return p.UserID2
	}
	return p.UserID1
}

func getPrivateIndex(userId1, userId2 uint32) (uint32, uint32) {
	if userId1 > userId2 {
		return userId2, userId1
	}
	return userId1, userId2
}

// 获取私聊目标频道
func GetChatPrivate(userId1, userId2 uint32) (*OFChatPrivate, error) {
	userId1, userId2 = getPrivateIndex(userId1, userId2)
	info := &OFChatPrivate{
		UserID1: userId1,
		UserID2: userId2,
	}
	err := db.FirstOrCreate(info).Error
	return info, err
}

// 获取已创建的全部私聊频道
func GetAllChatPrivate(userId uint32) ([]*OFChatPrivate, error) {
	list := make([]*OFChatPrivate, 0)
	err := db.Where("user_id1 = ? OR user_id2 = ?", userId, userId).
		Find(&list).Error
	return list, err
}

// 私聊记录表
type OFChatPrivateMsg struct {
	ID        int64  `gorm:"primary_key;auto_increment"`
	PrivateID int64  `gorm:"not null;uniqueIndex:user_msg"` // 房间号
	UserId    uint32 `gorm:"not null;uniqueIndex:user_msg"` // 发送用户
	*OFChatMsg
}

// 获取私聊全部聊天记录
func GetAllChatPrivateMsg(userId1, userId2 uint32) ([]*OFChatPrivateMsg, error) {
	list := make([]*OFChatPrivateMsg, 0)
	err := db.Transaction(func(tx *gorm.DB) error {
		private, err := GetChatPrivate(userId1, userId2)
		if err != nil {
			return err
		}
		if private.IsNewMsg {
			private.IsNewMsg = false
			if err = tx.Save(private).Error; err != nil {
				return err
			}
		}
		err = tx.Where("private_id = ?", private.ID).
			Limit(100).
			Find(&list).Error
		return err
	})

	return list, err
}

/*
写入私聊消息
UserId 接收方
Msg 消息内容
*/
func CreateChatPrivateMsg(userId uint32, msg *OFChatPrivateMsg) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		private, err := GetChatPrivate(msg.UserId, userId)
		if err != nil {
			return err
		}
		if !private.IsNewMsg {
			private.IsNewMsg = true
			if err = tx.Save(private).Error; err != nil {
				return err
			}
		}
		msg.PrivateID = private.ID
		err = tx.Create(msg).Error
		return err
	})
	return err
}

// 系统消息记录表
type OFChatSystemMsg struct {
	ID int64 `gorm:"primary_key;auto_increment;index:system_msg_id"`
	*OFChatMsg
}
