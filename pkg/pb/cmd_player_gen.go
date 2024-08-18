
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetPlayerTokenReq) },
		func() ProtoMessage { return new(GetPlayerTokenRsp) },
		func() ProtoMessage { return new(PlayerLoginReq) },
		func() ProtoMessage { return new(PlayerLoginRsp) },
		func() ProtoMessage { return new(PlayerLogoutReq) },
		func() ProtoMessage { return new(PlayerLogoutRsp) },
		func() ProtoMessage { return new(PlayerLogoutNotify) },
		func() ProtoMessage { return new(PlayerDataNotify) },
		func() ProtoMessage { return new(ChangeGameTimeReq) },
		func() ProtoMessage { return new(ChangeGameTimeRsp) },
		func() ProtoMessage { return new(PlayerGameTimeNotify) },
		func() ProtoMessage { return new(PlayerPropNotify) },
		func() ProtoMessage { return new(ClientTriggerEventNotify) },
		func() ProtoMessage { return new(SetPlayerPropReq) },
		func() ProtoMessage { return new(SetPlayerPropRsp) },
		func() ProtoMessage { return new(SetPlayerBornDataReq) },
		func() ProtoMessage { return new(SetPlayerBornDataRsp) },
		func() ProtoMessage { return new(DoSetPlayerBornDataNotify) },
		func() ProtoMessage { return new(PlayerPropChangeNotify) },
		func() ProtoMessage { return new(SetPlayerNameReq) },
		func() ProtoMessage { return new(SetPlayerNameRsp) },
		func() ProtoMessage { return new(SetOpenStateReq) },
		func() ProtoMessage { return new(SetOpenStateRsp) },
		func() ProtoMessage { return new(OpenStateUpdateNotify) },
		func() ProtoMessage { return new(OpenStateChangeNotify) },
		func() ProtoMessage { return new(PlayerCookReq) },
		func() ProtoMessage { return new(PlayerCookRsp) },
		func() ProtoMessage { return new(PlayerRandomCookReq) },
		func() ProtoMessage { return new(PlayerRandomCookRsp) },
		func() ProtoMessage { return new(CookDataNotify) },
		func() ProtoMessage { return new(CookRecipeDataNotify) },
		func() ProtoMessage { return new(CookGradeDataNotify) },
		func() ProtoMessage { return new(PlayerCompoundMaterialReq) },
		func() ProtoMessage { return new(PlayerCompoundMaterialRsp) },
		func() ProtoMessage { return new(TakeCompoundOutputReq) },
		func() ProtoMessage { return new(TakeCompoundOutputRsp) },
		func() ProtoMessage { return new(CompoundDataNotify) },
		func() ProtoMessage { return new(GetCompoundDataReq) },
		func() ProtoMessage { return new(GetCompoundDataRsp) },
		func() ProtoMessage { return new(PlayerTimeNotify) },
		func() ProtoMessage { return new(PlayerSetPauseReq) },
		func() ProtoMessage { return new(PlayerSetPauseRsp) },
		func() ProtoMessage { return new(PlayerSetLanguageReq) },
		func() ProtoMessage { return new(PlayerSetLanguageRsp) },
		func() ProtoMessage { return new(DataResVersionNotify) },
		func() ProtoMessage { return new(DailyTaskDataNotify) },
		func() ProtoMessage { return new(DailyTaskProgressNotify) },
		func() ProtoMessage { return new(DailyTaskScoreRewardNotify) },
		func() ProtoMessage { return new(WorldOwnerDailyTaskNotify) },
		func() ProtoMessage { return new(AddRandTaskInfoNotify) },
		func() ProtoMessage { return new(RemoveRandTaskInfoNotify) },
		func() ProtoMessage { return new(TakePlayerLevelRewardReq) },
		func() ProtoMessage { return new(TakePlayerLevelRewardRsp) },
		func() ProtoMessage { return new(PlayerLevelRewardUpdateNotify) },
		func() ProtoMessage { return new(GivingRecordNotify) },
		func() ProtoMessage { return new(GivingRecordChangeNotify) },
		func() ProtoMessage { return new(ItemGivingReq) },
		func() ProtoMessage { return new(ItemGivingRsp) },
		func() ProtoMessage { return new(PlayerCookArgsReq) },
		func() ProtoMessage { return new(PlayerCookArgsRsp) },
		func() ProtoMessage { return new(PlayerLuaShellNotify) },
		func() ProtoMessage { return new(ServerDisconnectClientNotify) },
		func() ProtoMessage { return new(AntiAddictNotify) },
		func() ProtoMessage { return new(PlayerForceExitReq) },
		func() ProtoMessage { return new(PlayerForceExitRsp) },
		func() ProtoMessage { return new(PlayerInjectFixNotify) },
		func() ProtoMessage { return new(TaskVarNotify) },
		func() ProtoMessage { return new(ClientLockGameTimeNotify) },
		func() ProtoMessage { return new(GetNextResourceInfoReq) },
		func() ProtoMessage { return new(GetNextResourceInfoRsp) },
		func() ProtoMessage { return new(AdjustWorldLevelReq) },
		func() ProtoMessage { return new(AdjustWorldLevelRsp) },
		func() ProtoMessage { return new(DailyTaskFilterCityReq) },
		func() ProtoMessage { return new(DailyTaskFilterCityRsp) },
		func() ProtoMessage { return new(DailyTaskUnlockedCitiesNotify) },
		func() ProtoMessage { return new(ExclusiveRuleNotify) },
		func() ProtoMessage { return new(CompoundUnlockNotify) },
		func() ProtoMessage { return new(GetGameplayRecommendationReq) },
		func() ProtoMessage { return new(GetGameplayRecommendationRsp) },
		func() ProtoMessage { return new(TakeBackGivingItemReq) },
		func() ProtoMessage { return new(TakeBackGivingItemRsp) },
		func() ProtoMessage { return new(PlayerNicknameAuditDataNotify) },
		func() ProtoMessage { return new(PlayerNicknameNotify) },
		func() ProtoMessage { return new(NicknameAuditConfigNotify) },
		func() ProtoMessage { return new(ReadNicknameAuditReq) },
		func() ProtoMessage { return new(ReadNicknameAuditRsp) },
		func() ProtoMessage { return new(PlayerCompoundMaterialBoostReq) },
		func() ProtoMessage { return new(PlayerCompoundMaterialBoostRsp) },
		func() ProtoMessage { return new(PlayerGameTimeByLuaNotify) },
		func() ProtoMessage { return new(PlayerIpRegionNotify) },
	)
}

const (
	ProtoCommandGetPlayerTokenReq              ProtoCommand = 172
	ProtoCommandGetPlayerTokenRsp              ProtoCommand = 198
	ProtoCommandPlayerLoginReq                 ProtoCommand = 112
	ProtoCommandPlayerLoginRsp                 ProtoCommand = 135
	ProtoCommandPlayerLogoutReq                ProtoCommand = 107
	ProtoCommandPlayerLogoutRsp                ProtoCommand = 121
	ProtoCommandPlayerLogoutNotify             ProtoCommand = 103
	ProtoCommandPlayerDataNotify               ProtoCommand = 190
	ProtoCommandChangeGameTimeReq              ProtoCommand = 173
	ProtoCommandChangeGameTimeRsp              ProtoCommand = 199
	ProtoCommandPlayerGameTimeNotify           ProtoCommand = 131
	ProtoCommandPlayerPropNotify               ProtoCommand = 175
	ProtoCommandClientTriggerEventNotify       ProtoCommand = 148
	ProtoCommandSetPlayerPropReq               ProtoCommand = 197
	ProtoCommandSetPlayerPropRsp               ProtoCommand = 181
	ProtoCommandSetPlayerBornDataReq           ProtoCommand = 105
	ProtoCommandSetPlayerBornDataRsp           ProtoCommand = 182
	ProtoCommandDoSetPlayerBornDataNotify      ProtoCommand = 147
	ProtoCommandPlayerPropChangeNotify         ProtoCommand = 139
	ProtoCommandSetPlayerNameReq               ProtoCommand = 153
	ProtoCommandSetPlayerNameRsp               ProtoCommand = 122
	ProtoCommandSetOpenStateReq                ProtoCommand = 165
	ProtoCommandSetOpenStateRsp                ProtoCommand = 104
	ProtoCommandOpenStateUpdateNotify          ProtoCommand = 193
	ProtoCommandOpenStateChangeNotify          ProtoCommand = 127
	ProtoCommandPlayerCookReq                  ProtoCommand = 194
	ProtoCommandPlayerCookRsp                  ProtoCommand = 188
	ProtoCommandPlayerRandomCookReq            ProtoCommand = 126
	ProtoCommandPlayerRandomCookRsp            ProtoCommand = 163
	ProtoCommandCookDataNotify                 ProtoCommand = 195
	ProtoCommandCookRecipeDataNotify           ProtoCommand = 106
	ProtoCommandCookGradeDataNotify            ProtoCommand = 134
	ProtoCommandPlayerCompoundMaterialReq      ProtoCommand = 150
	ProtoCommandPlayerCompoundMaterialRsp      ProtoCommand = 143
	ProtoCommandTakeCompoundOutputReq          ProtoCommand = 174
	ProtoCommandTakeCompoundOutputRsp          ProtoCommand = 176
	ProtoCommandCompoundDataNotify             ProtoCommand = 146
	ProtoCommandGetCompoundDataReq             ProtoCommand = 141
	ProtoCommandGetCompoundDataRsp             ProtoCommand = 149
	ProtoCommandPlayerTimeNotify               ProtoCommand = 191
	ProtoCommandPlayerSetPauseReq              ProtoCommand = 124
	ProtoCommandPlayerSetPauseRsp              ProtoCommand = 156
	ProtoCommandPlayerSetLanguageReq           ProtoCommand = 142
	ProtoCommandPlayerSetLanguageRsp           ProtoCommand = 130
	ProtoCommandDataResVersionNotify           ProtoCommand = 167
	ProtoCommandDailyTaskDataNotify            ProtoCommand = 158
	ProtoCommandDailyTaskProgressNotify        ProtoCommand = 170
	ProtoCommandDailyTaskScoreRewardNotify     ProtoCommand = 117
	ProtoCommandWorldOwnerDailyTaskNotify      ProtoCommand = 102
	ProtoCommandAddRandTaskInfoNotify          ProtoCommand = 119
	ProtoCommandRemoveRandTaskInfoNotify       ProtoCommand = 161
	ProtoCommandTakePlayerLevelRewardReq       ProtoCommand = 129
	ProtoCommandTakePlayerLevelRewardRsp       ProtoCommand = 157
	ProtoCommandPlayerLevelRewardUpdateNotify  ProtoCommand = 200
	ProtoCommandGivingRecordNotify             ProtoCommand = 116
	ProtoCommandGivingRecordChangeNotify       ProtoCommand = 187
	ProtoCommandItemGivingReq                  ProtoCommand = 140
	ProtoCommandItemGivingRsp                  ProtoCommand = 118
	ProtoCommandPlayerCookArgsReq              ProtoCommand = 166
	ProtoCommandPlayerCookArgsRsp              ProtoCommand = 168
	ProtoCommandPlayerLuaShellNotify           ProtoCommand = 133
	ProtoCommandServerDisconnectClientNotify   ProtoCommand = 184
	ProtoCommandAntiAddictNotify               ProtoCommand = 180
	ProtoCommandPlayerForceExitReq             ProtoCommand = 189
	ProtoCommandPlayerForceExitRsp             ProtoCommand = 159
	ProtoCommandPlayerInjectFixNotify          ProtoCommand = 132
	ProtoCommandTaskVarNotify                  ProtoCommand = 160
	ProtoCommandClientLockGameTimeNotify       ProtoCommand = 114
	ProtoCommandGetNextResourceInfoReq         ProtoCommand = 192
	ProtoCommandGetNextResourceInfoRsp         ProtoCommand = 120
	ProtoCommandAdjustWorldLevelReq            ProtoCommand = 164
	ProtoCommandAdjustWorldLevelRsp            ProtoCommand = 138
	ProtoCommandDailyTaskFilterCityReq         ProtoCommand = 111
	ProtoCommandDailyTaskFilterCityRsp         ProtoCommand = 144
	ProtoCommandDailyTaskUnlockedCitiesNotify  ProtoCommand = 186
	ProtoCommandExclusiveRuleNotify            ProtoCommand = 101
	ProtoCommandCompoundUnlockNotify           ProtoCommand = 128
	ProtoCommandGetGameplayRecommendationReq   ProtoCommand = 151
	ProtoCommandGetGameplayRecommendationRsp   ProtoCommand = 123
	ProtoCommandTakeBackGivingItemReq          ProtoCommand = 171
	ProtoCommandTakeBackGivingItemRsp          ProtoCommand = 145
	ProtoCommandPlayerNicknameAuditDataNotify  ProtoCommand = 108
	ProtoCommandPlayerNicknameNotify           ProtoCommand = 109
	ProtoCommandNicknameAuditConfigNotify      ProtoCommand = 152
	ProtoCommandReadNicknameAuditReq           ProtoCommand = 177
	ProtoCommandReadNicknameAuditRsp           ProtoCommand = 137
	ProtoCommandPlayerCompoundMaterialBoostReq ProtoCommand = 185
	ProtoCommandPlayerCompoundMaterialBoostRsp ProtoCommand = 125
	ProtoCommandPlayerGameTimeByLuaNotify      ProtoCommand = 178
	ProtoCommandPlayerIpRegionNotify           ProtoCommand = 136
)

func (*GetPlayerTokenReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerTokenReq }
func (*GetPlayerTokenReq) ProtoMessageType() ProtoMessageType { return "GetPlayerTokenReq" }

func (*GetPlayerTokenRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerTokenRsp }
func (*GetPlayerTokenRsp) ProtoMessageType() ProtoMessageType { return "GetPlayerTokenRsp" }

func (*PlayerLoginReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLoginReq }
func (*PlayerLoginReq) ProtoMessageType() ProtoMessageType { return "PlayerLoginReq" }

func (*PlayerLoginRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLoginRsp }
func (*PlayerLoginRsp) ProtoMessageType() ProtoMessageType { return "PlayerLoginRsp" }

func (*PlayerLogoutReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLogoutReq }
func (*PlayerLogoutReq) ProtoMessageType() ProtoMessageType { return "PlayerLogoutReq" }

func (*PlayerLogoutRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLogoutRsp }
func (*PlayerLogoutRsp) ProtoMessageType() ProtoMessageType { return "PlayerLogoutRsp" }

func (*PlayerLogoutNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLogoutNotify }
func (*PlayerLogoutNotify) ProtoMessageType() ProtoMessageType { return "PlayerLogoutNotify" }

func (*PlayerDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerDataNotify }
func (*PlayerDataNotify) ProtoMessageType() ProtoMessageType { return "PlayerDataNotify" }

func (*ChangeGameTimeReq) ProtoCommand() ProtoCommand         { return ProtoCommandChangeGameTimeReq }
func (*ChangeGameTimeReq) ProtoMessageType() ProtoMessageType { return "ChangeGameTimeReq" }

func (*ChangeGameTimeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChangeGameTimeRsp }
func (*ChangeGameTimeRsp) ProtoMessageType() ProtoMessageType { return "ChangeGameTimeRsp" }

func (*PlayerGameTimeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerGameTimeNotify }
func (*PlayerGameTimeNotify) ProtoMessageType() ProtoMessageType { return "PlayerGameTimeNotify" }

func (*PlayerPropNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerPropNotify }
func (*PlayerPropNotify) ProtoMessageType() ProtoMessageType { return "PlayerPropNotify" }

func (*ClientTriggerEventNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientTriggerEventNotify
}
func (*ClientTriggerEventNotify) ProtoMessageType() ProtoMessageType {
	return "ClientTriggerEventNotify"
}

func (*SetPlayerPropReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerPropReq }
func (*SetPlayerPropReq) ProtoMessageType() ProtoMessageType { return "SetPlayerPropReq" }

func (*SetPlayerPropRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerPropRsp }
func (*SetPlayerPropRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerPropRsp" }

func (*SetPlayerBornDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerBornDataReq }
func (*SetPlayerBornDataReq) ProtoMessageType() ProtoMessageType { return "SetPlayerBornDataReq" }

func (*SetPlayerBornDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerBornDataRsp }
func (*SetPlayerBornDataRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerBornDataRsp" }

func (*DoSetPlayerBornDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDoSetPlayerBornDataNotify
}
func (*DoSetPlayerBornDataNotify) ProtoMessageType() ProtoMessageType {
	return "DoSetPlayerBornDataNotify"
}

func (*PlayerPropChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerPropChangeNotify }
func (*PlayerPropChangeNotify) ProtoMessageType() ProtoMessageType { return "PlayerPropChangeNotify" }

func (*SetPlayerNameReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerNameReq }
func (*SetPlayerNameReq) ProtoMessageType() ProtoMessageType { return "SetPlayerNameReq" }

func (*SetPlayerNameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerNameRsp }
func (*SetPlayerNameRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerNameRsp" }

func (*SetOpenStateReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetOpenStateReq }
func (*SetOpenStateReq) ProtoMessageType() ProtoMessageType { return "SetOpenStateReq" }

func (*SetOpenStateRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetOpenStateRsp }
func (*SetOpenStateRsp) ProtoMessageType() ProtoMessageType { return "SetOpenStateRsp" }

func (*OpenStateUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOpenStateUpdateNotify }
func (*OpenStateUpdateNotify) ProtoMessageType() ProtoMessageType { return "OpenStateUpdateNotify" }

func (*OpenStateChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOpenStateChangeNotify }
func (*OpenStateChangeNotify) ProtoMessageType() ProtoMessageType { return "OpenStateChangeNotify" }

func (*PlayerCookReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCookReq }
func (*PlayerCookReq) ProtoMessageType() ProtoMessageType { return "PlayerCookReq" }

func (*PlayerCookRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCookRsp }
func (*PlayerCookRsp) ProtoMessageType() ProtoMessageType { return "PlayerCookRsp" }

func (*PlayerRandomCookReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerRandomCookReq }
func (*PlayerRandomCookReq) ProtoMessageType() ProtoMessageType { return "PlayerRandomCookReq" }

func (*PlayerRandomCookRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerRandomCookRsp }
func (*PlayerRandomCookRsp) ProtoMessageType() ProtoMessageType { return "PlayerRandomCookRsp" }

func (*CookDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCookDataNotify }
func (*CookDataNotify) ProtoMessageType() ProtoMessageType { return "CookDataNotify" }

func (*CookRecipeDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCookRecipeDataNotify }
func (*CookRecipeDataNotify) ProtoMessageType() ProtoMessageType { return "CookRecipeDataNotify" }

func (*CookGradeDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCookGradeDataNotify }
func (*CookGradeDataNotify) ProtoMessageType() ProtoMessageType { return "CookGradeDataNotify" }

func (*PlayerCompoundMaterialReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerCompoundMaterialReq
}
func (*PlayerCompoundMaterialReq) ProtoMessageType() ProtoMessageType {
	return "PlayerCompoundMaterialReq"
}

func (*PlayerCompoundMaterialRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerCompoundMaterialRsp
}
func (*PlayerCompoundMaterialRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerCompoundMaterialRsp"
}

func (*TakeCompoundOutputReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeCompoundOutputReq }
func (*TakeCompoundOutputReq) ProtoMessageType() ProtoMessageType { return "TakeCompoundOutputReq" }

func (*TakeCompoundOutputRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeCompoundOutputRsp }
func (*TakeCompoundOutputRsp) ProtoMessageType() ProtoMessageType { return "TakeCompoundOutputRsp" }

func (*CompoundDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCompoundDataNotify }
func (*CompoundDataNotify) ProtoMessageType() ProtoMessageType { return "CompoundDataNotify" }

func (*GetCompoundDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetCompoundDataReq }
func (*GetCompoundDataReq) ProtoMessageType() ProtoMessageType { return "GetCompoundDataReq" }

func (*GetCompoundDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetCompoundDataRsp }
func (*GetCompoundDataRsp) ProtoMessageType() ProtoMessageType { return "GetCompoundDataRsp" }

func (*PlayerTimeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerTimeNotify }
func (*PlayerTimeNotify) ProtoMessageType() ProtoMessageType { return "PlayerTimeNotify" }

func (*PlayerSetPauseReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerSetPauseReq }
func (*PlayerSetPauseReq) ProtoMessageType() ProtoMessageType { return "PlayerSetPauseReq" }

func (*PlayerSetPauseRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerSetPauseRsp }
func (*PlayerSetPauseRsp) ProtoMessageType() ProtoMessageType { return "PlayerSetPauseRsp" }

func (*PlayerSetLanguageReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerSetLanguageReq }
func (*PlayerSetLanguageReq) ProtoMessageType() ProtoMessageType { return "PlayerSetLanguageReq" }

func (*PlayerSetLanguageRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerSetLanguageRsp }
func (*PlayerSetLanguageRsp) ProtoMessageType() ProtoMessageType { return "PlayerSetLanguageRsp" }

func (*DataResVersionNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDataResVersionNotify }
func (*DataResVersionNotify) ProtoMessageType() ProtoMessageType { return "DataResVersionNotify" }

func (*DailyTaskDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDailyTaskDataNotify }
func (*DailyTaskDataNotify) ProtoMessageType() ProtoMessageType { return "DailyTaskDataNotify" }

func (*DailyTaskProgressNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDailyTaskProgressNotify
}
func (*DailyTaskProgressNotify) ProtoMessageType() ProtoMessageType { return "DailyTaskProgressNotify" }

func (*DailyTaskScoreRewardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDailyTaskScoreRewardNotify
}
func (*DailyTaskScoreRewardNotify) ProtoMessageType() ProtoMessageType {
	return "DailyTaskScoreRewardNotify"
}

func (*WorldOwnerDailyTaskNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldOwnerDailyTaskNotify
}
func (*WorldOwnerDailyTaskNotify) ProtoMessageType() ProtoMessageType {
	return "WorldOwnerDailyTaskNotify"
}

func (*AddRandTaskInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAddRandTaskInfoNotify }
func (*AddRandTaskInfoNotify) ProtoMessageType() ProtoMessageType { return "AddRandTaskInfoNotify" }

func (*RemoveRandTaskInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRemoveRandTaskInfoNotify
}
func (*RemoveRandTaskInfoNotify) ProtoMessageType() ProtoMessageType {
	return "RemoveRandTaskInfoNotify"
}

func (*TakePlayerLevelRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakePlayerLevelRewardReq
}
func (*TakePlayerLevelRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakePlayerLevelRewardReq"
}

func (*TakePlayerLevelRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakePlayerLevelRewardRsp
}
func (*TakePlayerLevelRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakePlayerLevelRewardRsp"
}

func (*PlayerLevelRewardUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerLevelRewardUpdateNotify
}
func (*PlayerLevelRewardUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerLevelRewardUpdateNotify"
}

func (*GivingRecordNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGivingRecordNotify }
func (*GivingRecordNotify) ProtoMessageType() ProtoMessageType { return "GivingRecordNotify" }

func (*GivingRecordChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGivingRecordChangeNotify
}
func (*GivingRecordChangeNotify) ProtoMessageType() ProtoMessageType {
	return "GivingRecordChangeNotify"
}

func (*ItemGivingReq) ProtoCommand() ProtoCommand         { return ProtoCommandItemGivingReq }
func (*ItemGivingReq) ProtoMessageType() ProtoMessageType { return "ItemGivingReq" }

func (*ItemGivingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandItemGivingRsp }
func (*ItemGivingRsp) ProtoMessageType() ProtoMessageType { return "ItemGivingRsp" }

func (*PlayerCookArgsReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCookArgsReq }
func (*PlayerCookArgsReq) ProtoMessageType() ProtoMessageType { return "PlayerCookArgsReq" }

func (*PlayerCookArgsRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCookArgsRsp }
func (*PlayerCookArgsRsp) ProtoMessageType() ProtoMessageType { return "PlayerCookArgsRsp" }

func (*PlayerLuaShellNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerLuaShellNotify }
func (*PlayerLuaShellNotify) ProtoMessageType() ProtoMessageType { return "PlayerLuaShellNotify" }

func (*ServerDisconnectClientNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerDisconnectClientNotify
}
func (*ServerDisconnectClientNotify) ProtoMessageType() ProtoMessageType {
	return "ServerDisconnectClientNotify"
}

func (*AntiAddictNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAntiAddictNotify }
func (*AntiAddictNotify) ProtoMessageType() ProtoMessageType { return "AntiAddictNotify" }

func (*PlayerForceExitReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerForceExitReq }
func (*PlayerForceExitReq) ProtoMessageType() ProtoMessageType { return "PlayerForceExitReq" }

func (*PlayerForceExitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerForceExitRsp }
func (*PlayerForceExitRsp) ProtoMessageType() ProtoMessageType { return "PlayerForceExitRsp" }

func (*PlayerInjectFixNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerInjectFixNotify }
func (*PlayerInjectFixNotify) ProtoMessageType() ProtoMessageType { return "PlayerInjectFixNotify" }

func (*TaskVarNotify) ProtoCommand() ProtoCommand         { return ProtoCommandTaskVarNotify }
func (*TaskVarNotify) ProtoMessageType() ProtoMessageType { return "TaskVarNotify" }

func (*ClientLockGameTimeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientLockGameTimeNotify
}
func (*ClientLockGameTimeNotify) ProtoMessageType() ProtoMessageType {
	return "ClientLockGameTimeNotify"
}

func (*GetNextResourceInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetNextResourceInfoReq }
func (*GetNextResourceInfoReq) ProtoMessageType() ProtoMessageType { return "GetNextResourceInfoReq" }

func (*GetNextResourceInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetNextResourceInfoRsp }
func (*GetNextResourceInfoRsp) ProtoMessageType() ProtoMessageType { return "GetNextResourceInfoRsp" }

func (*AdjustWorldLevelReq) ProtoCommand() ProtoCommand         { return ProtoCommandAdjustWorldLevelReq }
func (*AdjustWorldLevelReq) ProtoMessageType() ProtoMessageType { return "AdjustWorldLevelReq" }

func (*AdjustWorldLevelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAdjustWorldLevelRsp }
func (*AdjustWorldLevelRsp) ProtoMessageType() ProtoMessageType { return "AdjustWorldLevelRsp" }

func (*DailyTaskFilterCityReq) ProtoCommand() ProtoCommand         { return ProtoCommandDailyTaskFilterCityReq }
func (*DailyTaskFilterCityReq) ProtoMessageType() ProtoMessageType { return "DailyTaskFilterCityReq" }

func (*DailyTaskFilterCityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDailyTaskFilterCityRsp }
func (*DailyTaskFilterCityRsp) ProtoMessageType() ProtoMessageType { return "DailyTaskFilterCityRsp" }

func (*DailyTaskUnlockedCitiesNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDailyTaskUnlockedCitiesNotify
}
func (*DailyTaskUnlockedCitiesNotify) ProtoMessageType() ProtoMessageType {
	return "DailyTaskUnlockedCitiesNotify"
}

func (*ExclusiveRuleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandExclusiveRuleNotify }
func (*ExclusiveRuleNotify) ProtoMessageType() ProtoMessageType { return "ExclusiveRuleNotify" }

func (*CompoundUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCompoundUnlockNotify }
func (*CompoundUnlockNotify) ProtoMessageType() ProtoMessageType { return "CompoundUnlockNotify" }

func (*GetGameplayRecommendationReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetGameplayRecommendationReq
}
func (*GetGameplayRecommendationReq) ProtoMessageType() ProtoMessageType {
	return "GetGameplayRecommendationReq"
}

func (*GetGameplayRecommendationRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetGameplayRecommendationRsp
}
func (*GetGameplayRecommendationRsp) ProtoMessageType() ProtoMessageType {
	return "GetGameplayRecommendationRsp"
}

func (*TakeBackGivingItemReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeBackGivingItemReq }
func (*TakeBackGivingItemReq) ProtoMessageType() ProtoMessageType { return "TakeBackGivingItemReq" }

func (*TakeBackGivingItemRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeBackGivingItemRsp }
func (*TakeBackGivingItemRsp) ProtoMessageType() ProtoMessageType { return "TakeBackGivingItemRsp" }

func (*PlayerNicknameAuditDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerNicknameAuditDataNotify
}
func (*PlayerNicknameAuditDataNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerNicknameAuditDataNotify"
}

func (*PlayerNicknameNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerNicknameNotify }
func (*PlayerNicknameNotify) ProtoMessageType() ProtoMessageType { return "PlayerNicknameNotify" }

func (*NicknameAuditConfigNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandNicknameAuditConfigNotify
}
func (*NicknameAuditConfigNotify) ProtoMessageType() ProtoMessageType {
	return "NicknameAuditConfigNotify"
}

func (*ReadNicknameAuditReq) ProtoCommand() ProtoCommand         { return ProtoCommandReadNicknameAuditReq }
func (*ReadNicknameAuditReq) ProtoMessageType() ProtoMessageType { return "ReadNicknameAuditReq" }

func (*ReadNicknameAuditRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReadNicknameAuditRsp }
func (*ReadNicknameAuditRsp) ProtoMessageType() ProtoMessageType { return "ReadNicknameAuditRsp" }

func (*PlayerCompoundMaterialBoostReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerCompoundMaterialBoostReq
}
func (*PlayerCompoundMaterialBoostReq) ProtoMessageType() ProtoMessageType {
	return "PlayerCompoundMaterialBoostReq"
}

func (*PlayerCompoundMaterialBoostRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerCompoundMaterialBoostRsp
}
func (*PlayerCompoundMaterialBoostRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerCompoundMaterialBoostRsp"
}

func (*PlayerGameTimeByLuaNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGameTimeByLuaNotify
}
func (*PlayerGameTimeByLuaNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerGameTimeByLuaNotify"
}

func (*PlayerIpRegionNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerIpRegionNotify }
func (*PlayerIpRegionNotify) ProtoMessageType() ProtoMessageType { return "PlayerIpRegionNotify" }