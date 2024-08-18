
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(ToTheMoonQueryPathReq) },
		func() ProtoMessage { return new(ToTheMoonQueryPathRsp) },
		func() ProtoMessage { return new(ToTheMoonPingNotify) },
		func() ProtoMessage { return new(ToTheMoonEnterSceneReq) },
		func() ProtoMessage { return new(ToTheMoonEnterSceneRsp) },
		func() ProtoMessage { return new(ToTheMoonAddObstacleReq) },
		func() ProtoMessage { return new(ToTheMoonAddObstacleRsp) },
		func() ProtoMessage { return new(ToTheMoonRemoveObstacleReq) },
		func() ProtoMessage { return new(ToTheMoonRemoveObstacleRsp) },
		func() ProtoMessage { return new(ToTheMoonObstaclesModifyNotify) },
	)
}

const (
	ProtoCommandToTheMoonQueryPathReq          ProtoCommand = 6172
	ProtoCommandToTheMoonQueryPathRsp          ProtoCommand = 6198
	ProtoCommandToTheMoonPingNotify            ProtoCommand = 6112
	ProtoCommandToTheMoonEnterSceneReq         ProtoCommand = 6135
	ProtoCommandToTheMoonEnterSceneRsp         ProtoCommand = 6107
	ProtoCommandToTheMoonAddObstacleReq        ProtoCommand = 6121
	ProtoCommandToTheMoonAddObstacleRsp        ProtoCommand = 6103
	ProtoCommandToTheMoonRemoveObstacleReq     ProtoCommand = 6190
	ProtoCommandToTheMoonRemoveObstacleRsp     ProtoCommand = 6173
	ProtoCommandToTheMoonObstaclesModifyNotify ProtoCommand = 6199
)

func (*ToTheMoonQueryPathReq) ProtoCommand() ProtoCommand         { return ProtoCommandToTheMoonQueryPathReq }
func (*ToTheMoonQueryPathReq) ProtoMessageType() ProtoMessageType { return "ToTheMoonQueryPathReq" }

func (*ToTheMoonQueryPathRsp) ProtoCommand() ProtoCommand         { return ProtoCommandToTheMoonQueryPathRsp }
func (*ToTheMoonQueryPathRsp) ProtoMessageType() ProtoMessageType { return "ToTheMoonQueryPathRsp" }

func (*ToTheMoonPingNotify) ProtoCommand() ProtoCommand         { return ProtoCommandToTheMoonPingNotify }
func (*ToTheMoonPingNotify) ProtoMessageType() ProtoMessageType { return "ToTheMoonPingNotify" }

func (*ToTheMoonEnterSceneReq) ProtoCommand() ProtoCommand         { return ProtoCommandToTheMoonEnterSceneReq }
func (*ToTheMoonEnterSceneReq) ProtoMessageType() ProtoMessageType { return "ToTheMoonEnterSceneReq" }

func (*ToTheMoonEnterSceneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandToTheMoonEnterSceneRsp }
func (*ToTheMoonEnterSceneRsp) ProtoMessageType() ProtoMessageType { return "ToTheMoonEnterSceneRsp" }

func (*ToTheMoonAddObstacleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandToTheMoonAddObstacleReq
}
func (*ToTheMoonAddObstacleReq) ProtoMessageType() ProtoMessageType { return "ToTheMoonAddObstacleReq" }

func (*ToTheMoonAddObstacleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandToTheMoonAddObstacleRsp
}
func (*ToTheMoonAddObstacleRsp) ProtoMessageType() ProtoMessageType { return "ToTheMoonAddObstacleRsp" }

func (*ToTheMoonRemoveObstacleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandToTheMoonRemoveObstacleReq
}
func (*ToTheMoonRemoveObstacleReq) ProtoMessageType() ProtoMessageType {
	return "ToTheMoonRemoveObstacleReq"
}

func (*ToTheMoonRemoveObstacleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandToTheMoonRemoveObstacleRsp
}
func (*ToTheMoonRemoveObstacleRsp) ProtoMessageType() ProtoMessageType {
	return "ToTheMoonRemoveObstacleRsp"
}

func (*ToTheMoonObstaclesModifyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandToTheMoonObstaclesModifyNotify
}
func (*ToTheMoonObstaclesModifyNotify) ProtoMessageType() ProtoMessageType {
	return "ToTheMoonObstaclesModifyNotify"
}