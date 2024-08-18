
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(OnlinePlayerNumReq) },
		func() ProtoMessage { return new(OnlinePlayerNumRsp) },
		func() ProtoMessage { return new(KickoutPlayerNotify) },
		func() ProtoMessage { return new(CheckOnlinePlayerReq) },
		func() ProtoMessage { return new(CheckOnlinePlayerRsp) },
		func() ProtoMessage { return new(PlayerCombatForceReq) },
		func() ProtoMessage { return new(PlayerCombatForceRsp) },
		func() ProtoMessage { return new(CheckGameVersionReq) },
		func() ProtoMessage { return new(CheckGameVersionRsp) },
		func() ProtoMessage { return new(PlatformPlayerNumReq) },
		func() ProtoMessage { return new(PlatformPlayerNumRsp) },
		func() ProtoMessage { return new(QueryPlayerMemDataByMuipReq) },
		func() ProtoMessage { return new(QueryPlayerMemDataByMuipRsp) },
		func() ProtoMessage { return new(BindGmUidNotify) },
		func() ProtoMessage { return new(UnbindGmUidNotify) },
		func() ProtoMessage { return new(GetBindGmUidReq) },
		func() ProtoMessage { return new(GetBindGmUidRsp) },
		func() ProtoMessage { return new(PlatformAntiAddictNotify) },
		func() ProtoMessage { return new(PlayerLoginPerSecondReq) },
		func() ProtoMessage { return new(PlayerLoginPerSecondRsp) },
		func() ProtoMessage { return new(FineGrainedPlayerNumReq) },
		func() ProtoMessage { return new(FineGrainedPlayerNumRsp) },
		func() ProtoMessage { return new(CheckGameCrcVersionReq) },
		func() ProtoMessage { return new(CheckGameCrcVersionRsp) },
		func() ProtoMessage { return new(UpdateRedPointByMuipNotify) },
		func() ProtoMessage { return new(SendConcertProductReq) },
		func() ProtoMessage { return new(SendConcertProductRsp) },
		func() ProtoMessage { return new(QueryConcertProductInfoReq) },
		func() ProtoMessage { return new(QueryConcertProductInfoRsp) },
		func() ProtoMessage { return new(PlayerMpModeReq) },
		func() ProtoMessage { return new(PlayerMpModeRsp) },
	)
}

const (
	ProtoCommandOnlinePlayerNumReq          ProtoCommand = 10272
	ProtoCommandOnlinePlayerNumRsp          ProtoCommand = 10298
	ProtoCommandKickoutPlayerNotify         ProtoCommand = 10212
	ProtoCommandCheckOnlinePlayerReq        ProtoCommand = 10235
	ProtoCommandCheckOnlinePlayerRsp        ProtoCommand = 10207
	ProtoCommandPlayerCombatForceReq        ProtoCommand = 10221
	ProtoCommandPlayerCombatForceRsp        ProtoCommand = 10203
	ProtoCommandCheckGameVersionReq         ProtoCommand = 10290
	ProtoCommandCheckGameVersionRsp         ProtoCommand = 10273
	ProtoCommandPlatformPlayerNumReq        ProtoCommand = 10299
	ProtoCommandPlatformPlayerNumRsp        ProtoCommand = 10231
	ProtoCommandQueryPlayerMemDataByMuipReq ProtoCommand = 10275
	ProtoCommandQueryPlayerMemDataByMuipRsp ProtoCommand = 10248
	ProtoCommandBindGmUidNotify             ProtoCommand = 10297
	ProtoCommandUnbindGmUidNotify           ProtoCommand = 10281
	ProtoCommandGetBindGmUidReq             ProtoCommand = 10205
	ProtoCommandGetBindGmUidRsp             ProtoCommand = 10282
	ProtoCommandPlatformAntiAddictNotify    ProtoCommand = 10247
	ProtoCommandPlayerLoginPerSecondReq     ProtoCommand = 10239
	ProtoCommandPlayerLoginPerSecondRsp     ProtoCommand = 10253
	ProtoCommandFineGrainedPlayerNumReq     ProtoCommand = 10222
	ProtoCommandFineGrainedPlayerNumRsp     ProtoCommand = 10265
	ProtoCommandCheckGameCrcVersionReq      ProtoCommand = 10204
	ProtoCommandCheckGameCrcVersionRsp      ProtoCommand = 10293
	ProtoCommandUpdateRedPointByMuipNotify  ProtoCommand = 10227
	ProtoCommandSendConcertProductReq       ProtoCommand = 10294
	ProtoCommandSendConcertProductRsp       ProtoCommand = 10288
	ProtoCommandQueryConcertProductInfoReq  ProtoCommand = 10226
	ProtoCommandQueryConcertProductInfoRsp  ProtoCommand = 10263
	ProtoCommandPlayerMpModeReq             ProtoCommand = 10295
	ProtoCommandPlayerMpModeRsp             ProtoCommand = 10206
)

func (*OnlinePlayerNumReq) ProtoCommand() ProtoCommand         { return ProtoCommandOnlinePlayerNumReq }
func (*OnlinePlayerNumReq) ProtoMessageType() ProtoMessageType { return "OnlinePlayerNumReq" }

func (*OnlinePlayerNumRsp) ProtoCommand() ProtoCommand         { return ProtoCommandOnlinePlayerNumRsp }
func (*OnlinePlayerNumRsp) ProtoMessageType() ProtoMessageType { return "OnlinePlayerNumRsp" }

func (*KickoutPlayerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandKickoutPlayerNotify }
func (*KickoutPlayerNotify) ProtoMessageType() ProtoMessageType { return "KickoutPlayerNotify" }

func (*CheckOnlinePlayerReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckOnlinePlayerReq }
func (*CheckOnlinePlayerReq) ProtoMessageType() ProtoMessageType { return "CheckOnlinePlayerReq" }

func (*CheckOnlinePlayerRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckOnlinePlayerRsp }
func (*CheckOnlinePlayerRsp) ProtoMessageType() ProtoMessageType { return "CheckOnlinePlayerRsp" }

func (*PlayerCombatForceReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCombatForceReq }
func (*PlayerCombatForceReq) ProtoMessageType() ProtoMessageType { return "PlayerCombatForceReq" }

func (*PlayerCombatForceRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCombatForceRsp }
func (*PlayerCombatForceRsp) ProtoMessageType() ProtoMessageType { return "PlayerCombatForceRsp" }

func (*CheckGameVersionReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGameVersionReq }
func (*CheckGameVersionReq) ProtoMessageType() ProtoMessageType { return "CheckGameVersionReq" }

func (*CheckGameVersionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGameVersionRsp }
func (*CheckGameVersionRsp) ProtoMessageType() ProtoMessageType { return "CheckGameVersionRsp" }

func (*PlatformPlayerNumReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlatformPlayerNumReq }
func (*PlatformPlayerNumReq) ProtoMessageType() ProtoMessageType { return "PlatformPlayerNumReq" }

func (*PlatformPlayerNumRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlatformPlayerNumRsp }
func (*PlatformPlayerNumRsp) ProtoMessageType() ProtoMessageType { return "PlatformPlayerNumRsp" }

func (*QueryPlayerMemDataByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryPlayerMemDataByMuipReq
}
func (*QueryPlayerMemDataByMuipReq) ProtoMessageType() ProtoMessageType {
	return "QueryPlayerMemDataByMuipReq"
}

func (*QueryPlayerMemDataByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryPlayerMemDataByMuipRsp
}
func (*QueryPlayerMemDataByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "QueryPlayerMemDataByMuipRsp"
}

func (*BindGmUidNotify) ProtoCommand() ProtoCommand         { return ProtoCommandBindGmUidNotify }
func (*BindGmUidNotify) ProtoMessageType() ProtoMessageType { return "BindGmUidNotify" }

func (*UnbindGmUidNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUnbindGmUidNotify }
func (*UnbindGmUidNotify) ProtoMessageType() ProtoMessageType { return "UnbindGmUidNotify" }

func (*GetBindGmUidReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetBindGmUidReq }
func (*GetBindGmUidReq) ProtoMessageType() ProtoMessageType { return "GetBindGmUidReq" }

func (*GetBindGmUidRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetBindGmUidRsp }
func (*GetBindGmUidRsp) ProtoMessageType() ProtoMessageType { return "GetBindGmUidRsp" }

func (*PlatformAntiAddictNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlatformAntiAddictNotify
}
func (*PlatformAntiAddictNotify) ProtoMessageType() ProtoMessageType {
	return "PlatformAntiAddictNotify"
}

func (*PlayerLoginPerSecondReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerLoginPerSecondReq
}
func (*PlayerLoginPerSecondReq) ProtoMessageType() ProtoMessageType { return "PlayerLoginPerSecondReq" }

func (*PlayerLoginPerSecondRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerLoginPerSecondRsp
}
func (*PlayerLoginPerSecondRsp) ProtoMessageType() ProtoMessageType { return "PlayerLoginPerSecondRsp" }

func (*FineGrainedPlayerNumReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFineGrainedPlayerNumReq
}
func (*FineGrainedPlayerNumReq) ProtoMessageType() ProtoMessageType { return "FineGrainedPlayerNumReq" }

func (*FineGrainedPlayerNumRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFineGrainedPlayerNumRsp
}
func (*FineGrainedPlayerNumRsp) ProtoMessageType() ProtoMessageType { return "FineGrainedPlayerNumRsp" }

func (*CheckGameCrcVersionReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGameCrcVersionReq }
func (*CheckGameCrcVersionReq) ProtoMessageType() ProtoMessageType { return "CheckGameCrcVersionReq" }

func (*CheckGameCrcVersionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGameCrcVersionRsp }
func (*CheckGameCrcVersionRsp) ProtoMessageType() ProtoMessageType { return "CheckGameCrcVersionRsp" }

func (*UpdateRedPointByMuipNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdateRedPointByMuipNotify
}
func (*UpdateRedPointByMuipNotify) ProtoMessageType() ProtoMessageType {
	return "UpdateRedPointByMuipNotify"
}

func (*SendConcertProductReq) ProtoCommand() ProtoCommand         { return ProtoCommandSendConcertProductReq }
func (*SendConcertProductReq) ProtoMessageType() ProtoMessageType { return "SendConcertProductReq" }

func (*SendConcertProductRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSendConcertProductRsp }
func (*SendConcertProductRsp) ProtoMessageType() ProtoMessageType { return "SendConcertProductRsp" }

func (*QueryConcertProductInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryConcertProductInfoReq
}
func (*QueryConcertProductInfoReq) ProtoMessageType() ProtoMessageType {
	return "QueryConcertProductInfoReq"
}

func (*QueryConcertProductInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryConcertProductInfoRsp
}
func (*QueryConcertProductInfoRsp) ProtoMessageType() ProtoMessageType {
	return "QueryConcertProductInfoRsp"
}

func (*PlayerMpModeReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerMpModeReq }
func (*PlayerMpModeReq) ProtoMessageType() ProtoMessageType { return "PlayerMpModeReq" }

func (*PlayerMpModeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerMpModeRsp }
func (*PlayerMpModeRsp) ProtoMessageType() ProtoMessageType { return "PlayerMpModeRsp" }