
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(NormalUidOpNotify) },
		func() ProtoMessage { return new(ServerMessageNotify) },
	)
}

const (
	ProtoCommandNormalUidOpNotify   ProtoCommand = 5726
	ProtoCommandServerMessageNotify ProtoCommand = 5718
)

func (*NormalUidOpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandNormalUidOpNotify }
func (*NormalUidOpNotify) ProtoMessageType() ProtoMessageType { return "NormalUidOpNotify" }

func (*ServerMessageNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerMessageNotify }
func (*ServerMessageNotify) ProtoMessageType() ProtoMessageType { return "ServerMessageNotify" }