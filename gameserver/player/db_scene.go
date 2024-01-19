package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetScene() *spb.Scene {
	if g.PlayerPb.Scene == nil {
		g.PlayerPb.Scene = &spb.Scene{
			EntryId: 1010101,
			PlaneId: 10101,
			FloorId: 10101001,
		}
	}
	return g.PlayerPb.Scene
}

func (g *GamePlayer) GetPos() *spb.VectorBin {
	if g.PlayerPb.Pos == nil {
		g.PlayerPb.Pos = &spb.VectorBin{
			X: -43300,
			Y: 6,
			Z: -37960,
		}
	}
	return g.PlayerPb.Pos
}

func (g *GamePlayer) GetRot() *spb.VectorBin {
	if g.PlayerPb.Rot == nil {
		g.PlayerPb.Rot = &spb.VectorBin{
			X: 0,
			Y: 90000,
			Z: 0,
		}
	}
	return g.PlayerPb.Rot
}

func (g *GamePlayer) GetPropByID(sceneGroup *gdconf.LevelGroup, groupID uint32) *proto.SceneEntityGroupInfo {
	entityGroupLists := &proto.SceneEntityGroupInfo{
		GroupId:    groupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, propList := range sceneGroup.PropList {
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,     // 文件名后那个G
			InstId:   propList.ID, // ID
			EntityId: uint32(g.GetNextGameObjectGuid()),
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(propList.PosX * 1000),
					Y: int32(propList.PosY * 1000),
					Z: int32(propList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(propList.RotY * 1000),
					Z: 0,
				},
			},
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: gdconf.GetStateValue(propList.State),
			},
		}
		entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
	}
	return entityGroupLists
}

func (g *GamePlayer) GetNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup, groupID uint32, entityMap map[uint32]*EntityList) (*proto.SceneEntityGroupInfo, map[uint32]*EntityList) {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(monsterList.RotY * 1000),
					Z: 0,
				},
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.PlayerPb.WorldLevel,
				MonsterId:  monsterList.NPCMonsterID,
				EventId:    monsterList.EventID,
			},
		}
		// 添加实体
		entityMap[entityId] = &EntityList{
			Entity:  monsterList.EventID,
			GroupId: groupID,
			Pos: &Vector{
				X: int32(monsterList.PosX * 1000),
				Y: int32(monsterList.PosY * 1000),
				Z: int32(monsterList.PosZ * 1000),
			},
			Rot: &Vector{
				X: 0,
				Y: int32(monsterList.RotY * 1000),
				Z: 0,
			},
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList, entityMap
}

func (g *GamePlayer) GetNPCByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup, groupID uint32) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   npcList.ID,
			EntityId: uint32(g.GetNextGameObjectGuid()),
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(npcList.PosX * 1000),
					Y: int32(npcList.PosY * 1000),
					Z: int32(npcList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(npcList.RotY * 1000),
					Z: 0,
				},
			},
			Npc: &proto.SceneNpcInfo{
				ExtraInfo: nil,
				NpcId:     npcList.NPCID,
			},
		}
		if npcList.FirstDialogueGroupID != 0 {
			g.GetSceneNpcList()[npcList.NPCID] = npcList.FirstDialogueGroupID
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}
