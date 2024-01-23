package player

import (
	"net"
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

var SNOWFLAKE *alg.SnowflakeWorker // 雪花唯一id生成器

type GamePlayer struct {
	IsToken        bool // 是否通过token验证
	Uid            uint32
	LastActiveTime int64 // 最近一次的活跃时间
	// 玩家数据
	Player   *PlayerData
	PlayerPb *spb.PlayerBasicCompBin // 玩家pb数据
	GateConn net.Conn
}

type NetMsg struct {
	G         *GamePlayer
	CmdId     uint16
	PlayerMsg pb.Message
	Type      int
}

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

var blacklist = []uint16{cmd.SceneEntityMoveScRsp, cmd.SceneEntityMoveCsReq, cmd.PlayerHeartBeatCsReq, cmd.PlayerHeartBeatScRsp} // 黑名单
func isValid(cmdid uint16) bool {
	for _, value := range blacklist {
		if cmdid == value {
			return false
		}
	}
	return true
}

func (g *GamePlayer) Send(cmdId uint16, playerMsg pb.Message) {
	// 打印需要的数据包
	if isValid(cmdId) {
		data := protojson.Format(playerMsg)
		logger.Debug("[UID:%v] S --> C : CmdId: %v KcpMsg: \n%s\n", g.Uid, cmdId, data)
	}
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)

	gtgMsg := &spb.PlayerToGameByGateRsp{
		MessageType: 0,
		PlayerBin:   binMsg,
	}

	g.sendGate(cmd.PlayerToGameByGateRsp, gtgMsg)
}

func (g *GamePlayer) DecodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	// 打印需要的数据包
	if isValid(cmdId) {
		data := protojson.Format(protoObj)
		logger.Debug("[UID:%v] C --> S : NAME: %s KcpMsg: \n%s\n", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
	return protoObj
}

func (g *GamePlayer) KickPlayer() {
	/*
		TODO
		1.通知node game玩家下线
		2.通知gate game玩家下线
		3.保存数据到数据库
		4.断开gate-game连接
	*/
	UpDataPlayer(g)
}

func UpDataPlayer(g *GamePlayer) error {
	var err error
	if g.PlayerPb == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	dbDate := new(db.Player)
	dbDate.AccountUid = g.Uid

	dbDate.PlayerDataPb, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
	}

	if err = db.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return err
	}

	logger.Info("数据库账号:%v 数据更新", g.Uid)
	return nil
}

/*

func (g *GamePlayer) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		if g.Seed == 0 {
			return
		}
		lastActiveTime := g.LastActiveTime
		timestamp := time.Now().Unix()
		if timestamp-lastActiveTime >= 120 {
			g.KickPlayer()
			return
		}
	}
}
*/

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
