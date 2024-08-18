
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(DungeonEntryInfoReq) },
		func() ProtoMessage { return new(DungeonEntryInfoRsp) },
		func() ProtoMessage { return new(PlayerEnterDungeonReq) },
		func() ProtoMessage { return new(PlayerEnterDungeonRsp) },
		func() ProtoMessage { return new(PlayerQuitDungeonReq) },
		func() ProtoMessage { return new(PlayerQuitDungeonRsp) },
		func() ProtoMessage { return new(DungeonWayPointNotify) },
		func() ProtoMessage { return new(DungeonWayPointActivateReq) },
		func() ProtoMessage { return new(DungeonWayPointActivateRsp) },
		func() ProtoMessage { return new(DungeonSettleNotify) },
		func() ProtoMessage { return new(DungeonPlayerDieNotify) },
		func() ProtoMessage { return new(DungeonDieOptionReq) },
		func() ProtoMessage { return new(DungeonDieOptionRsp) },
		func() ProtoMessage { return new(DungeonShowReminderNotify) },
		func() ProtoMessage { return new(DungeonPlayerDieReq) },
		func() ProtoMessage { return new(DungeonPlayerDieRsp) },
		func() ProtoMessage { return new(DungeonDataNotify) },
		func() ProtoMessage { return new(DungeonChallengeBeginNotify) },
		func() ProtoMessage { return new(DungeonChallengeFinishNotify) },
		func() ProtoMessage { return new(ChallengeDataNotify) },
		func() ProtoMessage { return new(DungeonFollowNotify) },
		func() ProtoMessage { return new(DungeonGetStatueDropReq) },
		func() ProtoMessage { return new(DungeonGetStatueDropRsp) },
		func() ProtoMessage { return new(ChallengeRecordNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamInfoNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamInviteNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamRefuseNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamPlayerLeaveNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamDismissNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamCreateReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamCreateRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamInviteReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamInviteRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamKickReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamKickRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamLeaveReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamLeaveRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamReplyInviteReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamReplyInviteRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamSetReadyReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamSetReadyRsp) },
		func() ProtoMessage { return new(DungeonCandidateTeamChangeAvatarReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamChangeAvatarRsp) },
		func() ProtoMessage { return new(GetDailyDungeonEntryInfoReq) },
		func() ProtoMessage { return new(GetDailyDungeonEntryInfoRsp) },
		func() ProtoMessage { return new(DungeonSlipRevivePointActivateReq) },
		func() ProtoMessage { return new(DungeonSlipRevivePointActivateRsp) },
		func() ProtoMessage { return new(DungeonInterruptChallengeReq) },
		func() ProtoMessage { return new(DungeonInterruptChallengeRsp) },
		func() ProtoMessage { return new(InteractDailyDungeonInfoNotify) },
		func() ProtoMessage { return new(DungeonRestartReq) },
		func() ProtoMessage { return new(DungeonRestartRsp) },
		func() ProtoMessage { return new(DungeonRestartInviteNotify) },
		func() ProtoMessage { return new(DungeonRestartInviteReplyReq) },
		func() ProtoMessage { return new(DungeonRestartInviteReplyRsp) },
		func() ProtoMessage { return new(DungeonRestartInviteReplyNotify) },
		func() ProtoMessage { return new(DungeonRestartResultNotify) },
		func() ProtoMessage { return new(DungeonCandidateTeamSetChangingAvatarReq) },
		func() ProtoMessage { return new(DungeonCandidateTeamSetChangingAvatarRsp) },
		func() ProtoMessage { return new(MistTrialFloorLevelNotify) },
		func() ProtoMessage { return new(DungeonReviseLevelNotify) },
	)
}

const (
	ProtoCommandDungeonEntryInfoReq                      ProtoCommand = 972
	ProtoCommandDungeonEntryInfoRsp                      ProtoCommand = 998
	ProtoCommandPlayerEnterDungeonReq                    ProtoCommand = 912
	ProtoCommandPlayerEnterDungeonRsp                    ProtoCommand = 935
	ProtoCommandPlayerQuitDungeonReq                     ProtoCommand = 907
	ProtoCommandPlayerQuitDungeonRsp                     ProtoCommand = 921
	ProtoCommandDungeonWayPointNotify                    ProtoCommand = 903
	ProtoCommandDungeonWayPointActivateReq               ProtoCommand = 990
	ProtoCommandDungeonWayPointActivateRsp               ProtoCommand = 973
	ProtoCommandDungeonSettleNotify                      ProtoCommand = 999
	ProtoCommandDungeonPlayerDieNotify                   ProtoCommand = 931
	ProtoCommandDungeonDieOptionReq                      ProtoCommand = 975
	ProtoCommandDungeonDieOptionRsp                      ProtoCommand = 948
	ProtoCommandDungeonShowReminderNotify                ProtoCommand = 997
	ProtoCommandDungeonPlayerDieReq                      ProtoCommand = 981
	ProtoCommandDungeonPlayerDieRsp                      ProtoCommand = 905
	ProtoCommandDungeonDataNotify                        ProtoCommand = 982
	ProtoCommandDungeonChallengeBeginNotify              ProtoCommand = 947
	ProtoCommandDungeonChallengeFinishNotify             ProtoCommand = 939
	ProtoCommandChallengeDataNotify                      ProtoCommand = 953
	ProtoCommandDungeonFollowNotify                      ProtoCommand = 922
	ProtoCommandDungeonGetStatueDropReq                  ProtoCommand = 965
	ProtoCommandDungeonGetStatueDropRsp                  ProtoCommand = 904
	ProtoCommandChallengeRecordNotify                    ProtoCommand = 993
	ProtoCommandDungeonCandidateTeamInfoNotify           ProtoCommand = 927
	ProtoCommandDungeonCandidateTeamInviteNotify         ProtoCommand = 994
	ProtoCommandDungeonCandidateTeamRefuseNotify         ProtoCommand = 988
	ProtoCommandDungeonCandidateTeamPlayerLeaveNotify    ProtoCommand = 926
	ProtoCommandDungeonCandidateTeamDismissNotify        ProtoCommand = 963
	ProtoCommandDungeonCandidateTeamCreateReq            ProtoCommand = 995
	ProtoCommandDungeonCandidateTeamCreateRsp            ProtoCommand = 906
	ProtoCommandDungeonCandidateTeamInviteReq            ProtoCommand = 934
	ProtoCommandDungeonCandidateTeamInviteRsp            ProtoCommand = 950
	ProtoCommandDungeonCandidateTeamKickReq              ProtoCommand = 943
	ProtoCommandDungeonCandidateTeamKickRsp              ProtoCommand = 974
	ProtoCommandDungeonCandidateTeamLeaveReq             ProtoCommand = 976
	ProtoCommandDungeonCandidateTeamLeaveRsp             ProtoCommand = 946
	ProtoCommandDungeonCandidateTeamReplyInviteReq       ProtoCommand = 941
	ProtoCommandDungeonCandidateTeamReplyInviteRsp       ProtoCommand = 949
	ProtoCommandDungeonCandidateTeamSetReadyReq          ProtoCommand = 991
	ProtoCommandDungeonCandidateTeamSetReadyRsp          ProtoCommand = 924
	ProtoCommandDungeonCandidateTeamChangeAvatarReq      ProtoCommand = 956
	ProtoCommandDungeonCandidateTeamChangeAvatarRsp      ProtoCommand = 942
	ProtoCommandGetDailyDungeonEntryInfoReq              ProtoCommand = 930
	ProtoCommandGetDailyDungeonEntryInfoRsp              ProtoCommand = 967
	ProtoCommandDungeonSlipRevivePointActivateReq        ProtoCommand = 958
	ProtoCommandDungeonSlipRevivePointActivateRsp        ProtoCommand = 970
	ProtoCommandDungeonInterruptChallengeReq             ProtoCommand = 917
	ProtoCommandDungeonInterruptChallengeRsp             ProtoCommand = 902
	ProtoCommandInteractDailyDungeonInfoNotify           ProtoCommand = 919
	ProtoCommandDungeonRestartReq                        ProtoCommand = 961
	ProtoCommandDungeonRestartRsp                        ProtoCommand = 929
	ProtoCommandDungeonRestartInviteNotify               ProtoCommand = 957
	ProtoCommandDungeonRestartInviteReplyReq             ProtoCommand = 1000
	ProtoCommandDungeonRestartInviteReplyRsp             ProtoCommand = 916
	ProtoCommandDungeonRestartInviteReplyNotify          ProtoCommand = 987
	ProtoCommandDungeonRestartResultNotify               ProtoCommand = 940
	ProtoCommandDungeonCandidateTeamSetChangingAvatarReq ProtoCommand = 918
	ProtoCommandDungeonCandidateTeamSetChangingAvatarRsp ProtoCommand = 966
	ProtoCommandMistTrialFloorLevelNotify                ProtoCommand = 968
	ProtoCommandDungeonReviseLevelNotify                 ProtoCommand = 933
)

func (*DungeonEntryInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonEntryInfoReq }
func (*DungeonEntryInfoReq) ProtoMessageType() ProtoMessageType { return "DungeonEntryInfoReq" }

func (*DungeonEntryInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonEntryInfoRsp }
func (*DungeonEntryInfoRsp) ProtoMessageType() ProtoMessageType { return "DungeonEntryInfoRsp" }

func (*PlayerEnterDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerEnterDungeonReq }
func (*PlayerEnterDungeonReq) ProtoMessageType() ProtoMessageType { return "PlayerEnterDungeonReq" }

func (*PlayerEnterDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerEnterDungeonRsp }
func (*PlayerEnterDungeonRsp) ProtoMessageType() ProtoMessageType { return "PlayerEnterDungeonRsp" }

func (*PlayerQuitDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerQuitDungeonReq }
func (*PlayerQuitDungeonReq) ProtoMessageType() ProtoMessageType { return "PlayerQuitDungeonReq" }

func (*PlayerQuitDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerQuitDungeonRsp }
func (*PlayerQuitDungeonRsp) ProtoMessageType() ProtoMessageType { return "PlayerQuitDungeonRsp" }

func (*DungeonWayPointNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonWayPointNotify }
func (*DungeonWayPointNotify) ProtoMessageType() ProtoMessageType { return "DungeonWayPointNotify" }

func (*DungeonWayPointActivateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonWayPointActivateReq
}
func (*DungeonWayPointActivateReq) ProtoMessageType() ProtoMessageType {
	return "DungeonWayPointActivateReq"
}

func (*DungeonWayPointActivateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonWayPointActivateRsp
}
func (*DungeonWayPointActivateRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonWayPointActivateRsp"
}

func (*DungeonSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonSettleNotify }
func (*DungeonSettleNotify) ProtoMessageType() ProtoMessageType { return "DungeonSettleNotify" }

func (*DungeonPlayerDieNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonPlayerDieNotify }
func (*DungeonPlayerDieNotify) ProtoMessageType() ProtoMessageType { return "DungeonPlayerDieNotify" }

func (*DungeonDieOptionReq) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonDieOptionReq }
func (*DungeonDieOptionReq) ProtoMessageType() ProtoMessageType { return "DungeonDieOptionReq" }

func (*DungeonDieOptionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonDieOptionRsp }
func (*DungeonDieOptionRsp) ProtoMessageType() ProtoMessageType { return "DungeonDieOptionRsp" }

func (*DungeonShowReminderNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonShowReminderNotify
}
func (*DungeonShowReminderNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonShowReminderNotify"
}

func (*DungeonPlayerDieReq) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonPlayerDieReq }
func (*DungeonPlayerDieReq) ProtoMessageType() ProtoMessageType { return "DungeonPlayerDieReq" }

func (*DungeonPlayerDieRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonPlayerDieRsp }
func (*DungeonPlayerDieRsp) ProtoMessageType() ProtoMessageType { return "DungeonPlayerDieRsp" }

func (*DungeonDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonDataNotify }
func (*DungeonDataNotify) ProtoMessageType() ProtoMessageType { return "DungeonDataNotify" }

func (*DungeonChallengeBeginNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonChallengeBeginNotify
}
func (*DungeonChallengeBeginNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonChallengeBeginNotify"
}

func (*DungeonChallengeFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonChallengeFinishNotify
}
func (*DungeonChallengeFinishNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonChallengeFinishNotify"
}

func (*ChallengeDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChallengeDataNotify }
func (*ChallengeDataNotify) ProtoMessageType() ProtoMessageType { return "ChallengeDataNotify" }

func (*DungeonFollowNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonFollowNotify }
func (*DungeonFollowNotify) ProtoMessageType() ProtoMessageType { return "DungeonFollowNotify" }

func (*DungeonGetStatueDropReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonGetStatueDropReq
}
func (*DungeonGetStatueDropReq) ProtoMessageType() ProtoMessageType { return "DungeonGetStatueDropReq" }

func (*DungeonGetStatueDropRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonGetStatueDropRsp
}
func (*DungeonGetStatueDropRsp) ProtoMessageType() ProtoMessageType { return "DungeonGetStatueDropRsp" }

func (*ChallengeRecordNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChallengeRecordNotify }
func (*ChallengeRecordNotify) ProtoMessageType() ProtoMessageType { return "ChallengeRecordNotify" }

func (*DungeonCandidateTeamInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamInfoNotify
}
func (*DungeonCandidateTeamInfoNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamInfoNotify"
}

func (*DungeonCandidateTeamInviteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamInviteNotify
}
func (*DungeonCandidateTeamInviteNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamInviteNotify"
}

func (*DungeonCandidateTeamRefuseNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamRefuseNotify
}
func (*DungeonCandidateTeamRefuseNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamRefuseNotify"
}

func (*DungeonCandidateTeamPlayerLeaveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamPlayerLeaveNotify
}
func (*DungeonCandidateTeamPlayerLeaveNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamPlayerLeaveNotify"
}

func (*DungeonCandidateTeamDismissNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamDismissNotify
}
func (*DungeonCandidateTeamDismissNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamDismissNotify"
}

func (*DungeonCandidateTeamCreateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamCreateReq
}
func (*DungeonCandidateTeamCreateReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamCreateReq"
}

func (*DungeonCandidateTeamCreateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamCreateRsp
}
func (*DungeonCandidateTeamCreateRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamCreateRsp"
}

func (*DungeonCandidateTeamInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamInviteReq
}
func (*DungeonCandidateTeamInviteReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamInviteReq"
}

func (*DungeonCandidateTeamInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamInviteRsp
}
func (*DungeonCandidateTeamInviteRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamInviteRsp"
}

func (*DungeonCandidateTeamKickReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamKickReq
}
func (*DungeonCandidateTeamKickReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamKickReq"
}

func (*DungeonCandidateTeamKickRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamKickRsp
}
func (*DungeonCandidateTeamKickRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamKickRsp"
}

func (*DungeonCandidateTeamLeaveReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamLeaveReq
}
func (*DungeonCandidateTeamLeaveReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamLeaveReq"
}

func (*DungeonCandidateTeamLeaveRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamLeaveRsp
}
func (*DungeonCandidateTeamLeaveRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamLeaveRsp"
}

func (*DungeonCandidateTeamReplyInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamReplyInviteReq
}
func (*DungeonCandidateTeamReplyInviteReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamReplyInviteReq"
}

func (*DungeonCandidateTeamReplyInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamReplyInviteRsp
}
func (*DungeonCandidateTeamReplyInviteRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamReplyInviteRsp"
}

func (*DungeonCandidateTeamSetReadyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamSetReadyReq
}
func (*DungeonCandidateTeamSetReadyReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamSetReadyReq"
}

func (*DungeonCandidateTeamSetReadyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamSetReadyRsp
}
func (*DungeonCandidateTeamSetReadyRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamSetReadyRsp"
}

func (*DungeonCandidateTeamChangeAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamChangeAvatarReq
}
func (*DungeonCandidateTeamChangeAvatarReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamChangeAvatarReq"
}

func (*DungeonCandidateTeamChangeAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamChangeAvatarRsp
}
func (*DungeonCandidateTeamChangeAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamChangeAvatarRsp"
}

func (*GetDailyDungeonEntryInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetDailyDungeonEntryInfoReq
}
func (*GetDailyDungeonEntryInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetDailyDungeonEntryInfoReq"
}

func (*GetDailyDungeonEntryInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetDailyDungeonEntryInfoRsp
}
func (*GetDailyDungeonEntryInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetDailyDungeonEntryInfoRsp"
}

func (*DungeonSlipRevivePointActivateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonSlipRevivePointActivateReq
}
func (*DungeonSlipRevivePointActivateReq) ProtoMessageType() ProtoMessageType {
	return "DungeonSlipRevivePointActivateReq"
}

func (*DungeonSlipRevivePointActivateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonSlipRevivePointActivateRsp
}
func (*DungeonSlipRevivePointActivateRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonSlipRevivePointActivateRsp"
}

func (*DungeonInterruptChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonInterruptChallengeReq
}
func (*DungeonInterruptChallengeReq) ProtoMessageType() ProtoMessageType {
	return "DungeonInterruptChallengeReq"
}

func (*DungeonInterruptChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonInterruptChallengeRsp
}
func (*DungeonInterruptChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonInterruptChallengeRsp"
}

func (*InteractDailyDungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInteractDailyDungeonInfoNotify
}
func (*InteractDailyDungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "InteractDailyDungeonInfoNotify"
}

func (*DungeonRestartReq) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonRestartReq }
func (*DungeonRestartReq) ProtoMessageType() ProtoMessageType { return "DungeonRestartReq" }

func (*DungeonRestartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDungeonRestartRsp }
func (*DungeonRestartRsp) ProtoMessageType() ProtoMessageType { return "DungeonRestartRsp" }

func (*DungeonRestartInviteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonRestartInviteNotify
}
func (*DungeonRestartInviteNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonRestartInviteNotify"
}

func (*DungeonRestartInviteReplyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonRestartInviteReplyReq
}
func (*DungeonRestartInviteReplyReq) ProtoMessageType() ProtoMessageType {
	return "DungeonRestartInviteReplyReq"
}

func (*DungeonRestartInviteReplyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonRestartInviteReplyRsp
}
func (*DungeonRestartInviteReplyRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonRestartInviteReplyRsp"
}

func (*DungeonRestartInviteReplyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonRestartInviteReplyNotify
}
func (*DungeonRestartInviteReplyNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonRestartInviteReplyNotify"
}

func (*DungeonRestartResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonRestartResultNotify
}
func (*DungeonRestartResultNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonRestartResultNotify"
}

func (*DungeonCandidateTeamSetChangingAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamSetChangingAvatarReq
}
func (*DungeonCandidateTeamSetChangingAvatarReq) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamSetChangingAvatarReq"
}

func (*DungeonCandidateTeamSetChangingAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonCandidateTeamSetChangingAvatarRsp
}
func (*DungeonCandidateTeamSetChangingAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "DungeonCandidateTeamSetChangingAvatarRsp"
}

func (*MistTrialFloorLevelNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialFloorLevelNotify
}
func (*MistTrialFloorLevelNotify) ProtoMessageType() ProtoMessageType {
	return "MistTrialFloorLevelNotify"
}

func (*DungeonReviseLevelNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonReviseLevelNotify
}
func (*DungeonReviseLevelNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonReviseLevelNotify"
}