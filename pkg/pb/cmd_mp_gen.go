
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerApplyEnterMpNotify) },
		func() ProtoMessage { return new(PlayerApplyEnterMpReq) },
		func() ProtoMessage { return new(PlayerApplyEnterMpRsp) },
		func() ProtoMessage { return new(PlayerApplyEnterMpResultNotify) },
		func() ProtoMessage { return new(PlayerApplyEnterMpResultReq) },
		func() ProtoMessage { return new(PlayerApplyEnterMpResultRsp) },
		func() ProtoMessage { return new(PlayerQuitFromMpNotify) },
		func() ProtoMessage { return new(PlayerPreEnterMpNotify) },
		func() ProtoMessage { return new(GetPlayerMpModeAvailabilityReq) },
		func() ProtoMessage { return new(GetPlayerMpModeAvailabilityRsp) },
		func() ProtoMessage { return new(PlayerSetOnlyMPWithPSPlayerReq) },
		func() ProtoMessage { return new(PlayerSetOnlyMPWithPSPlayerRsp) },
		func() ProtoMessage { return new(PSPlayerApplyEnterMpReq) },
		func() ProtoMessage { return new(PSPlayerApplyEnterMpRsp) },
		func() ProtoMessage { return new(MpPlayOwnerCheckReq) },
		func() ProtoMessage { return new(MpPlayOwnerCheckRsp) },
		func() ProtoMessage { return new(MpPlayOwnerStartInviteReq) },
		func() ProtoMessage { return new(MpPlayOwnerStartInviteRsp) },
		func() ProtoMessage { return new(MpPlayOwnerInviteNotify) },
		func() ProtoMessage { return new(MpPlayGuestReplyInviteReq) },
		func() ProtoMessage { return new(MpPlayGuestReplyInviteRsp) },
		func() ProtoMessage { return new(MpPlayGuestReplyNotify) },
		func() ProtoMessage { return new(MpPlayPrepareNotify) },
		func() ProtoMessage { return new(MpPlayInviteResultNotify) },
		func() ProtoMessage { return new(MpPlayPrepareInterruptNotify) },
		func() ProtoMessage { return new(MpBlockNotify) },
	)
}

const (
	ProtoCommandPlayerApplyEnterMpNotify       ProtoCommand = 1826
	ProtoCommandPlayerApplyEnterMpReq          ProtoCommand = 1818
	ProtoCommandPlayerApplyEnterMpRsp          ProtoCommand = 1825
	ProtoCommandPlayerApplyEnterMpResultNotify ProtoCommand = 1807
	ProtoCommandPlayerApplyEnterMpResultReq    ProtoCommand = 1802
	ProtoCommandPlayerApplyEnterMpResultRsp    ProtoCommand = 1831
	ProtoCommandPlayerQuitFromMpNotify         ProtoCommand = 1829
	ProtoCommandPlayerPreEnterMpNotify         ProtoCommand = 1822
	ProtoCommandGetPlayerMpModeAvailabilityReq ProtoCommand = 1844
	ProtoCommandGetPlayerMpModeAvailabilityRsp ProtoCommand = 1849
	ProtoCommandPlayerSetOnlyMPWithPSPlayerReq ProtoCommand = 1820
	ProtoCommandPlayerSetOnlyMPWithPSPlayerRsp ProtoCommand = 1845
	ProtoCommandPSPlayerApplyEnterMpReq        ProtoCommand = 1841
	ProtoCommandPSPlayerApplyEnterMpRsp        ProtoCommand = 1842
	ProtoCommandMpPlayOwnerCheckReq            ProtoCommand = 1814
	ProtoCommandMpPlayOwnerCheckRsp            ProtoCommand = 1847
	ProtoCommandMpPlayOwnerStartInviteReq      ProtoCommand = 1837
	ProtoCommandMpPlayOwnerStartInviteRsp      ProtoCommand = 1823
	ProtoCommandMpPlayOwnerInviteNotify        ProtoCommand = 1835
	ProtoCommandMpPlayGuestReplyInviteReq      ProtoCommand = 1848
	ProtoCommandMpPlayGuestReplyInviteRsp      ProtoCommand = 1850
	ProtoCommandMpPlayGuestReplyNotify         ProtoCommand = 1812
	ProtoCommandMpPlayPrepareNotify            ProtoCommand = 1833
	ProtoCommandMpPlayInviteResultNotify       ProtoCommand = 1815
	ProtoCommandMpPlayPrepareInterruptNotify   ProtoCommand = 1813
	ProtoCommandMpBlockNotify                  ProtoCommand = 1801
)

func (*PlayerApplyEnterMpNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterMpNotify
}
func (*PlayerApplyEnterMpNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterMpNotify"
}

func (*PlayerApplyEnterMpReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerApplyEnterMpReq }
func (*PlayerApplyEnterMpReq) ProtoMessageType() ProtoMessageType { return "PlayerApplyEnterMpReq" }

func (*PlayerApplyEnterMpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerApplyEnterMpRsp }
func (*PlayerApplyEnterMpRsp) ProtoMessageType() ProtoMessageType { return "PlayerApplyEnterMpRsp" }

func (*PlayerApplyEnterMpResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterMpResultNotify
}
func (*PlayerApplyEnterMpResultNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterMpResultNotify"
}

func (*PlayerApplyEnterMpResultReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterMpResultReq
}
func (*PlayerApplyEnterMpResultReq) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterMpResultReq"
}

func (*PlayerApplyEnterMpResultRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterMpResultRsp
}
func (*PlayerApplyEnterMpResultRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterMpResultRsp"
}

func (*PlayerQuitFromMpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerQuitFromMpNotify }
func (*PlayerQuitFromMpNotify) ProtoMessageType() ProtoMessageType { return "PlayerQuitFromMpNotify" }

func (*PlayerPreEnterMpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerPreEnterMpNotify }
func (*PlayerPreEnterMpNotify) ProtoMessageType() ProtoMessageType { return "PlayerPreEnterMpNotify" }

func (*GetPlayerMpModeAvailabilityReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpModeAvailabilityReq
}
func (*GetPlayerMpModeAvailabilityReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpModeAvailabilityReq"
}

func (*GetPlayerMpModeAvailabilityRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpModeAvailabilityRsp
}
func (*GetPlayerMpModeAvailabilityRsp) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpModeAvailabilityRsp"
}

func (*PlayerSetOnlyMPWithPSPlayerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerSetOnlyMPWithPSPlayerReq
}
func (*PlayerSetOnlyMPWithPSPlayerReq) ProtoMessageType() ProtoMessageType {
	return "PlayerSetOnlyMPWithPSPlayerReq"
}

func (*PlayerSetOnlyMPWithPSPlayerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerSetOnlyMPWithPSPlayerRsp
}
func (*PlayerSetOnlyMPWithPSPlayerRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerSetOnlyMPWithPSPlayerRsp"
}

func (*PSPlayerApplyEnterMpReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPSPlayerApplyEnterMpReq
}
func (*PSPlayerApplyEnterMpReq) ProtoMessageType() ProtoMessageType { return "PSPlayerApplyEnterMpReq" }

func (*PSPlayerApplyEnterMpRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPSPlayerApplyEnterMpRsp
}
func (*PSPlayerApplyEnterMpRsp) ProtoMessageType() ProtoMessageType { return "PSPlayerApplyEnterMpRsp" }

func (*MpPlayOwnerCheckReq) ProtoCommand() ProtoCommand         { return ProtoCommandMpPlayOwnerCheckReq }
func (*MpPlayOwnerCheckReq) ProtoMessageType() ProtoMessageType { return "MpPlayOwnerCheckReq" }

func (*MpPlayOwnerCheckRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMpPlayOwnerCheckRsp }
func (*MpPlayOwnerCheckRsp) ProtoMessageType() ProtoMessageType { return "MpPlayOwnerCheckRsp" }

func (*MpPlayOwnerStartInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayOwnerStartInviteReq
}
func (*MpPlayOwnerStartInviteReq) ProtoMessageType() ProtoMessageType {
	return "MpPlayOwnerStartInviteReq"
}

func (*MpPlayOwnerStartInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayOwnerStartInviteRsp
}
func (*MpPlayOwnerStartInviteRsp) ProtoMessageType() ProtoMessageType {
	return "MpPlayOwnerStartInviteRsp"
}

func (*MpPlayOwnerInviteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayOwnerInviteNotify
}
func (*MpPlayOwnerInviteNotify) ProtoMessageType() ProtoMessageType { return "MpPlayOwnerInviteNotify" }

func (*MpPlayGuestReplyInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayGuestReplyInviteReq
}
func (*MpPlayGuestReplyInviteReq) ProtoMessageType() ProtoMessageType {
	return "MpPlayGuestReplyInviteReq"
}

func (*MpPlayGuestReplyInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayGuestReplyInviteRsp
}
func (*MpPlayGuestReplyInviteRsp) ProtoMessageType() ProtoMessageType {
	return "MpPlayGuestReplyInviteRsp"
}

func (*MpPlayGuestReplyNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMpPlayGuestReplyNotify }
func (*MpPlayGuestReplyNotify) ProtoMessageType() ProtoMessageType { return "MpPlayGuestReplyNotify" }

func (*MpPlayPrepareNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMpPlayPrepareNotify }
func (*MpPlayPrepareNotify) ProtoMessageType() ProtoMessageType { return "MpPlayPrepareNotify" }

func (*MpPlayInviteResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayInviteResultNotify
}
func (*MpPlayInviteResultNotify) ProtoMessageType() ProtoMessageType {
	return "MpPlayInviteResultNotify"
}

func (*MpPlayPrepareInterruptNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMpPlayPrepareInterruptNotify
}
func (*MpPlayPrepareInterruptNotify) ProtoMessageType() ProtoMessageType {
	return "MpPlayPrepareInterruptNotify"
}

func (*MpBlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMpBlockNotify }
func (*MpBlockNotify) ProtoMessageType() ProtoMessageType { return "MpBlockNotify" }