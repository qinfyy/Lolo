package model

import (
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

type ItemModel struct {
	InstanceIndex  uint32                      // 物品索引生成器
	ItemBaseInfo   map[uint32]*ItemBaseInfo    // 基础物品 徽章 伞
	ItemWeaponMap  map[uint32]*ItemWeaponInfo  // 武器
	ItemFashionMap map[uint32]*ItemFashionInfo // 服装
	ItemArmorMap   map[uint32]*ItemArmorInfo   // 盔甲
	ItemPosterMap  map[uint32]*ItemPosterInfo  // 海报
}

func DefaultItemModel() *ItemModel {
	return &ItemModel{}
}

func (s *Player) GetItemModel() *ItemModel {
	if s == nil {
		return nil
	}
	if s.Item == nil {
		s.Item = DefaultItemModel()
	}
	return s.Item
}

func (i *ItemModel) NextInstanceIndex() uint32 {
	if i == nil {
		return 0
	}
	i.InstanceIndex++
	return i.InstanceIndex
}

func (i *ItemModel) AllItemModel() {
	for tag, confList := range gdconf.GetItemByNewBagItemTagAll() {
		switch tag {
		case proto.EBagItemTag_EBagItemTag_Badge,
			proto.EBagItemTag_EBagItemTag_Umbrella:
			for _, conf := range confList {
				i.AddItemBase(uint32(conf.ID), 999)
			}
		case proto.EBagItemTag_EBagItemTag_Weapon:
			for _, conf := range confList {
				i.AddItemWeaponInfo(uint32(conf.ID))
			}
		case proto.EBagItemTag_EBagItemTag_Fashion:
			for _, conf := range confList {
				i.AddItemFashionInfo(uint32(conf.ID))
			}
		case proto.EBagItemTag_EBagItemTag_Armor:
		case proto.EBagItemTag_EBagItemTag_Poster:
		}
	}
}

type EBagItemTag interface {
	GetPbItemDetail() *proto.ItemDetail
}

type ItemBaseInfo struct {
	ItemId   uint32
	Num      int64
	ItemType proto.EBagItemTag
}

func (i *ItemModel) GetItemBaseMap() map[uint32]*ItemBaseInfo {
	if i == nil {
		return nil
	}
	if i.ItemBaseInfo == nil {
		i.ItemBaseInfo = make(map[uint32]*ItemBaseInfo)
	}
	return i.ItemBaseInfo
}

func (i *ItemModel) AddItemBase(itemId uint32, num int64) {
	conf := gdconf.GetItemConfigure(itemId)
	list := i.GetItemBaseMap()
	if conf == nil || list == nil {
		log.Game.Warnf("添加基础物品失败,数据异常或不存在ID:%v", itemId)
		return
	}
	info := list[itemId]
	if info == nil {
		info = &ItemBaseInfo{
			ItemId:   itemId,
			Num:      0,
			ItemType: proto.EBagItemTag(conf.NewBagItemTag),
		}
		list[itemId] = info
	}
	info.Num += num
}

func (i *ItemBaseInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.ItemId,
			ItemTag: i.ItemType,
			Item: &proto.ItemInfo_BaseItem{
				BaseItem: &proto.BaseItem{
					ItemId: i.ItemId,
					Num:    i.Num,
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

type ItemWeaponInfo struct {
	WeaponId         uint32                  // 武器id
	InstanceId       uint32                  // 索引id
	WeaponSystemType proto.EWeaponSystemType // 装备类型
	Attack           uint32                  // 攻击力
	DamageBalance    uint32                  // 伤害平衡
	CriticalRatio    uint32                  // 临界比率
	RandomProperty   []*RandomProperty       // 随机属性
	WearerId         uint32                  // 装备者id
	Level            uint32                  // 等级
	StrengthLevel    uint32                  // 强度等级
	StrengthExp      uint32                  // 强度经验
	Star             uint32                  // 星数
	Inscription1     uint32                  // 铭文1
	Durability       uint32                  // 耐用性
	PropertyIndex    uint32                  // ?指数
	IsLock           bool                    // 是否锁
}

type RandomProperty struct {
	PropertyType proto.EPropertyType // 类型
	Value        uint32              // 值
}

func (i *ItemModel) GetItemWeaponMap() map[uint32]*ItemWeaponInfo {
	if i == nil {
		return nil
	}
	if i.ItemWeaponMap == nil {
		i.ItemWeaponMap = make(map[uint32]*ItemWeaponInfo)
	}
	return i.ItemWeaponMap
}

func (i *ItemModel) GetItemWeaponInfo(instanceId uint32) *ItemWeaponInfo {
	list := i.GetItemWeaponMap()
	return list[instanceId]
}

func (i *ItemModel) AddItemWeaponInfo(weaponId uint32) *ItemWeaponInfo {
	conf := gdconf.GetWeaponAllInfo(weaponId)
	list := i.GetItemWeaponMap()
	if conf == nil || list == nil {
		log.Game.Warnf("添加Weapon失败,数据异常或不存在ID:%v", weaponId)
		return nil
	}
	instanceId := i.NextInstanceIndex()
	info := &ItemWeaponInfo{
		WeaponId:         conf.WeaponId,
		InstanceId:       instanceId,
		WeaponSystemType: proto.EWeaponSystemType(conf.WeaponInfo.NewWeaponSystemType),
		Attack:           1,
		DamageBalance:    1,
		CriticalRatio:    1,
		RandomProperty:   make([]*RandomProperty, 0),
		WearerId:         0,
		Level:            1,
		StrengthLevel:    0,
		StrengthExp:      0,
		Star:             1,
		Inscription1:     0,
		Durability:       0,
		PropertyIndex:    1,
		IsLock:           false,
	}
	list[instanceId] = info

	return info
}

func (i *ItemWeaponInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.WeaponId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Weapon,
			Item: &proto.ItemInfo_Weapon{
				Weapon: i.GetPbWeaponInstance(),
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

func (i *ItemWeaponInfo) SetWearerId(id uint32) {
	i.WearerId = id
}

func (i *ItemWeaponInfo) GetPbWeaponInstance() *proto.WeaponInstance {
	info := &proto.WeaponInstance{
		WeaponId:       i.WeaponId,
		InstanceId:     i.InstanceId,
		Attack:         i.Attack,
		DamageBalance:  i.DamageBalance,
		CriticalRatio:  i.CriticalRatio,
		RandomProperty: make([]*proto.RandomProperty, 0),
		WearerId:       i.WearerId,
		Level:          i.Level,
		StrengthLevel:  i.StrengthLevel,
		StrengthExp:    i.StrengthExp,
		Star:           i.Star,
		Inscription1:   i.Inscription1,
		Durability:     i.Durability,
		PropertyIndex:  i.PropertyIndex,
		IsLock:         i.IsLock,
	}

	return info
}

type ItemFashionInfo struct {
	ItemId uint32
}

func (i *ItemModel) GetItemFashionMap() map[uint32]*ItemFashionInfo {
	if i == nil {
		return nil
	}
	if i.ItemFashionMap == nil {
		i.ItemFashionMap = make(map[uint32]*ItemFashionInfo)
	}
	return i.ItemFashionMap
}

func (i *ItemModel) AddItemFashionInfo(itemId uint32) bool {
	conf := gdconf.GetItemConfigure(itemId)
	list := i.GetItemFashionMap()
	if conf == nil || list == nil ||
		conf.NewBagItemTag != int32(proto.EBagItemTag_EBagItemTag_Fashion) {
		log.Game.Warnf("添加Fashion失败,数据异常或不存在ID:%v", itemId)
		return false
	}
	if list[itemId] != nil {
		return true
	}
	list[itemId] = &ItemFashionInfo{
		ItemId: itemId,
	}
	return true
}

func (i *ItemFashionInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.ItemId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Fashion,
			Item: &proto.ItemInfo_Outfit{
				Outfit: &proto.Outfit{
					OutfitId: i.ItemId,
					DyeSchemes: []*proto.OutfitDyeScheme{
						{
							SchemeIndex: 0,
							Colors:      make([]*proto.PosColor, 0),
							IsUnLock:    true,
						},
					},
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

type ItemArmorInfo struct {
	WeaponSystemType proto.EWeaponSystemType //  装备类型
}

func (i *ItemModel) GetItemArmorMap() map[uint32]*ItemArmorInfo {
	if i == nil {
		return nil
	}
	if i.ItemArmorMap == nil {
		i.ItemArmorMap = make(map[uint32]*ItemArmorInfo)
	}
	return i.ItemArmorMap
}

func (i *ItemModel) AddItemArmorInfo(itemId uint32) bool {
	return false
}

func (i *ItemArmorInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  0,
			ItemTag: proto.EBagItemTag_EBagItemTag_Armor,
			Item: &proto.ItemInfo_Armor{
				Armor: &proto.ArmorInstance{
					ArmorId:          0,
					InstanceId:       0,
					MainPropertyType: 0,
					MainPropertyVal:  0,
					RandomProperty:   nil,
					WearerId:         0,
					Level:            0,
					StrengthLevel:    0,
					StrengthExp:      0,
					PropertyIndex:    0,
					IsLock:           false,
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

func (i *ItemArmorInfo) GetPbArmorInstance() *proto.ArmorInstance {
	info := &proto.ArmorInstance{
		ArmorId:          0,
		InstanceId:       0,
		MainPropertyType: 0,
		MainPropertyVal:  0,
		RandomProperty:   make([]*proto.RandomProperty, 0),
		WearerId:         0,
		Level:            0,
		StrengthLevel:    0,
		StrengthExp:      0,
		PropertyIndex:    0,
		IsLock:           false,
	}

	return info
}

type ItemPosterInfo struct {
}

func (i *ItemModel) GetItemPosterMap() map[uint32]*ItemPosterInfo {
	if i == nil {
		return nil
	}
	if i.ItemPosterMap == nil {
		i.ItemPosterMap = make(map[uint32]*ItemPosterInfo)
	}
	return i.ItemPosterMap
}

func (i *ItemPosterInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  0,
			ItemTag: proto.EBagItemTag_EBagItemTag_Poster,
			Item: &proto.ItemInfo_Poster{
				Poster: i.GetPbPosterInstance(),
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

func (i *ItemPosterInfo) GetPbPosterInstance() *proto.PosterInstance {
	info := &proto.PosterInstance{
		PosterId:   0,
		InstanceId: 0,
		WearerId:   0,
		Star:       0,
	}
	return info
}
