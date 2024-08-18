
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(TryEnterHomeReq) },
		func() ProtoMessage { return new(TryEnterHomeRsp) },
		func() ProtoMessage { return new(JoinHomeWorldFailNotify) },
		func() ProtoMessage { return new(HomeBlockNotify) },
		func() ProtoMessage { return new(HomeGetBasicInfoReq) },
		func() ProtoMessage { return new(HomeBasicInfoNotify) },
		func() ProtoMessage { return new(HomeChangeEditModeReq) },
		func() ProtoMessage { return new(HomeChangeEditModeRsp) },
		func() ProtoMessage { return new(HomeChangeModuleReq) },
		func() ProtoMessage { return new(HomeChangeModuleRsp) },
		func() ProtoMessage { return new(HomeGetArrangementInfoReq) },
		func() ProtoMessage { return new(HomeGetArrangementInfoRsp) },
		func() ProtoMessage { return new(HomeUpdateArrangementInfoReq) },
		func() ProtoMessage { return new(HomeUpdateArrangementInfoRsp) },
		func() ProtoMessage { return new(GetPlayerHomeCompInfoReq) },
		func() ProtoMessage { return new(PlayerHomeCompInfoNotify) },
		func() ProtoMessage { return new(SetFriendEnterHomeOptionReq) },
		func() ProtoMessage { return new(SetFriendEnterHomeOptionRsp) },
		func() ProtoMessage { return new(PlayerApplyEnterHomeNotify) },
		func() ProtoMessage { return new(PlayerApplyEnterHomeResultReq) },
		func() ProtoMessage { return new(PlayerApplyEnterHomeResultRsp) },
		func() ProtoMessage { return new(PlayerApplyEnterHomeResultNotify) },
		func() ProtoMessage { return new(HomeSceneJumpReq) },
		func() ProtoMessage { return new(HomeSceneJumpRsp) },
		func() ProtoMessage { return new(HomeChooseModuleReq) },
		func() ProtoMessage { return new(HomeChooseModuleRsp) },
		func() ProtoMessage { return new(HomeModuleUnlockNotify) },
		func() ProtoMessage { return new(HomeGetOnlineStatusReq) },
		func() ProtoMessage { return new(HomeGetOnlineStatusRsp) },
		func() ProtoMessage { return new(HomeKickPlayerReq) },
		func() ProtoMessage { return new(HomeKickPlayerRsp) },
		func() ProtoMessage { return new(HomeModuleSeenReq) },
		func() ProtoMessage { return new(HomeModuleSeenRsp) },
		func() ProtoMessage { return new(UnlockedFurnitureFormulaDataNotify) },
		func() ProtoMessage { return new(UnlockedFurnitureSuiteDataNotify) },
		func() ProtoMessage { return new(GetHomeLevelUpRewardReq) },
		func() ProtoMessage { return new(GetHomeLevelUpRewardRsp) },
		func() ProtoMessage { return new(GetFurnitureCurModuleArrangeCountReq) },
		func() ProtoMessage { return new(FurnitureCurModuleArrangeCountNotify) },
		func() ProtoMessage { return new(HomeComfortInfoNotify) },
		func() ProtoMessage { return new(PlayerQuitFromHomeNotify) },
		func() ProtoMessage { return new(OtherPlayerEnterHomeNotify) },
		func() ProtoMessage { return new(HomePriorCheckNotify) },
		func() ProtoMessage { return new(HomeMarkPointNotify) },
		func() ProtoMessage { return new(HomeAllUnlockedBgmIdListNotify) },
		func() ProtoMessage { return new(HomeNewUnlockedBgmIdListNotify) },
		func() ProtoMessage { return new(HomeChangeBgmReq) },
		func() ProtoMessage { return new(HomeChangeBgmRsp) },
		func() ProtoMessage { return new(HomeChangeBgmNotify) },
		func() ProtoMessage { return new(HomePreChangeEditModeNotify) },
		func() ProtoMessage { return new(HomeEnterEditModeFinishReq) },
		func() ProtoMessage { return new(HomeEnterEditModeFinishRsp) },
		func() ProtoMessage { return new(FurnitureMakeReq) },
		func() ProtoMessage { return new(FurnitureMakeRsp) },
		func() ProtoMessage { return new(TakeFurnitureMakeReq) },
		func() ProtoMessage { return new(TakeFurnitureMakeRsp) },
		func() ProtoMessage { return new(FurnitureMakeFinishNotify) },
		func() ProtoMessage { return new(FurnitureMakeStartReq) },
		func() ProtoMessage { return new(FurnitureMakeStartRsp) },
		func() ProtoMessage { return new(FurnitureMakeCancelReq) },
		func() ProtoMessage { return new(FurnitureMakeCancelRsp) },
		func() ProtoMessage { return new(FurnitureMakeBeHelpedNotify) },
		func() ProtoMessage { return new(FurnitureMakeHelpReq) },
		func() ProtoMessage { return new(FurnitureMakeHelpRsp) },
		func() ProtoMessage { return new(FunitureMakeMakeInfoChangeNotify) },
		func() ProtoMessage { return new(HomeLimitedShopInfoReq) },
		func() ProtoMessage { return new(HomeLimitedShopInfoRsp) },
		func() ProtoMessage { return new(HomeLimitedShopInfoNotify) },
		func() ProtoMessage { return new(HomeLimitedShopGoodsListReq) },
		func() ProtoMessage { return new(HomeLimitedShopGoodsListRsp) },
		func() ProtoMessage { return new(HomeLimitedShopBuyGoodsReq) },
		func() ProtoMessage { return new(HomeLimitedShopBuyGoodsRsp) },
		func() ProtoMessage { return new(HomeLimitedShopInfoChangeNotify) },
		func() ProtoMessage { return new(HomeResourceNotify) },
		func() ProtoMessage { return new(HomeResourceTakeHomeCoinReq) },
		func() ProtoMessage { return new(HomeResourceTakeHomeCoinRsp) },
		func() ProtoMessage { return new(HomeResourceTakeFetterExpReq) },
		func() ProtoMessage { return new(HomeResourceTakeFetterExpRsp) },
		func() ProtoMessage { return new(HomeAvatarTalkFinishInfoNotify) },
		func() ProtoMessage { return new(HomeAvatarTalkReq) },
		func() ProtoMessage { return new(HomeAvatarTalkRsp) },
		func() ProtoMessage { return new(HomeAvatarRewardEventNotify) },
		func() ProtoMessage { return new(HomeAvatarRewardEventGetReq) },
		func() ProtoMessage { return new(HomeAvatarRewardEventGetRsp) },
		func() ProtoMessage { return new(HomeAvatarSummonAllEventNotify) },
		func() ProtoMessage { return new(HomeAvatarSummonEventReq) },
		func() ProtoMessage { return new(HomeAvatarSummonEventRsp) },
		func() ProtoMessage { return new(HomeAvatarCostumeChangeNotify) },
		func() ProtoMessage { return new(HomeAvatarSummonFinishReq) },
		func() ProtoMessage { return new(HomeAvatarSummonFinishRsp) },
		func() ProtoMessage { return new(HomeAvtarAllFinishRewardNotify) },
		func() ProtoMessage { return new(HomeAvatarAllFinishRewardNotify) },
		func() ProtoMessage { return new(HomeSceneInitFinishReq) },
		func() ProtoMessage { return new(HomeSceneInitFinishRsp) },
		func() ProtoMessage { return new(HomePlantSeedReq) },
		func() ProtoMessage { return new(HomePlantSeedRsp) },
		func() ProtoMessage { return new(HomePlantWeedReq) },
		func() ProtoMessage { return new(HomePlantWeedRsp) },
		func() ProtoMessage { return new(HomePlantInfoNotify) },
		func() ProtoMessage { return new(HomePlantFieldNotify) },
		func() ProtoMessage { return new(HomePlantInfoReq) },
		func() ProtoMessage { return new(HomePlantInfoRsp) },
		func() ProtoMessage { return new(HomeTransferReq) },
		func() ProtoMessage { return new(HomeTransferRsp) },
		func() ProtoMessage { return new(HomeGetFishFarmingInfoReq) },
		func() ProtoMessage { return new(HomeGetFishFarmingInfoRsp) },
		func() ProtoMessage { return new(HomeFishFarmingInfoNotify) },
		func() ProtoMessage { return new(HomeUpdateFishFarmingInfoReq) },
		func() ProtoMessage { return new(HomeUpdateFishFarmingInfoRsp) },
		func() ProtoMessage { return new(HomeUpdateScenePointFishFarmingInfoReq) },
		func() ProtoMessage { return new(HomeUpdateScenePointFishFarmingInfoRsp) },
		func() ProtoMessage { return new(HomeScenePointFishFarmingInfoNotify) },
		func() ProtoMessage { return new(HomeCustomFurnitureInfoNotify) },
		func() ProtoMessage { return new(HomeEditCustomFurnitureReq) },
		func() ProtoMessage { return new(HomeEditCustomFurnitureRsp) },
		func() ProtoMessage { return new(HomePictureFrameInfoNotify) },
		func() ProtoMessage { return new(HomeUpdatePictureFrameInfoReq) },
		func() ProtoMessage { return new(HomeUpdatePictureFrameInfoRsp) },
		func() ProtoMessage { return new(HomeRacingGallerySettleNotify) },
		func() ProtoMessage { return new(HomeGetGroupRecordReq) },
		func() ProtoMessage { return new(HomeGetGroupRecordRsp) },
		func() ProtoMessage { return new(HomeClearGroupRecordReq) },
		func() ProtoMessage { return new(HomeClearGroupRecordRsp) },
		func() ProtoMessage { return new(HomeBalloonGallerySettleNotify) },
		func() ProtoMessage { return new(HomeBalloonGalleryScoreNotify) },
		func() ProtoMessage { return new(HomeSeekFurnitureGalleryScoreNotify) },
		func() ProtoMessage { return new(GetHomeExchangeWoodInfoReq) },
		func() ProtoMessage { return new(GetHomeExchangeWoodInfoRsp) },
		func() ProtoMessage { return new(HomeExchangeWoodReq) },
		func() ProtoMessage { return new(HomeExchangeWoodRsp) },
		func() ProtoMessage { return new(HomeGetBlueprintSlotInfoReq) },
		func() ProtoMessage { return new(HomeGetBlueprintSlotInfoRsp) },
		func() ProtoMessage { return new(HomeSetBlueprintSlotOptionReq) },
		func() ProtoMessage { return new(HomeSetBlueprintSlotOptionRsp) },
		func() ProtoMessage { return new(HomeSetBlueprintFriendOptionReq) },
		func() ProtoMessage { return new(HomeSetBlueprintFriendOptionRsp) },
		func() ProtoMessage { return new(HomeBlueprintInfoNotify) },
		func() ProtoMessage { return new(HomePreviewBlueprintReq) },
		func() ProtoMessage { return new(HomePreviewBlueprintRsp) },
		func() ProtoMessage { return new(HomeCreateBlueprintReq) },
		func() ProtoMessage { return new(HomeCreateBlueprintRsp) },
		func() ProtoMessage { return new(HomeDeleteBlueprintReq) },
		func() ProtoMessage { return new(HomeDeleteBlueprintRsp) },
		func() ProtoMessage { return new(HomeSearchBlueprintReq) },
		func() ProtoMessage { return new(HomeSearchBlueprintRsp) },
		func() ProtoMessage { return new(HomeSaveArrangementNoChangeReq) },
		func() ProtoMessage { return new(HomeSaveArrangementNoChangeRsp) },
	)
}

const (
	ProtoCommandTryEnterHomeReq                        ProtoCommand = 4482
	ProtoCommandTryEnterHomeRsp                        ProtoCommand = 4653
	ProtoCommandJoinHomeWorldFailNotify                ProtoCommand = 4530
	ProtoCommandHomeBlockNotify                        ProtoCommand = 4543
	ProtoCommandHomeGetBasicInfoReq                    ProtoCommand = 4655
	ProtoCommandHomeBasicInfoNotify                    ProtoCommand = 4885
	ProtoCommandHomeChangeEditModeReq                  ProtoCommand = 4564
	ProtoCommandHomeChangeEditModeRsp                  ProtoCommand = 4559
	ProtoCommandHomeChangeModuleReq                    ProtoCommand = 4809
	ProtoCommandHomeChangeModuleRsp                    ProtoCommand = 4596
	ProtoCommandHomeGetArrangementInfoReq              ProtoCommand = 4848
	ProtoCommandHomeGetArrangementInfoRsp              ProtoCommand = 4844
	ProtoCommandHomeUpdateArrangementInfoReq           ProtoCommand = 4510
	ProtoCommandHomeUpdateArrangementInfoRsp           ProtoCommand = 4757
	ProtoCommandGetPlayerHomeCompInfoReq               ProtoCommand = 4597
	ProtoCommandPlayerHomeCompInfoNotify               ProtoCommand = 4880
	ProtoCommandSetFriendEnterHomeOptionReq            ProtoCommand = 4494
	ProtoCommandSetFriendEnterHomeOptionRsp            ProtoCommand = 4743
	ProtoCommandPlayerApplyEnterHomeNotify             ProtoCommand = 4533
	ProtoCommandPlayerApplyEnterHomeResultReq          ProtoCommand = 4693
	ProtoCommandPlayerApplyEnterHomeResultRsp          ProtoCommand = 4706
	ProtoCommandPlayerApplyEnterHomeResultNotify       ProtoCommand = 4468
	ProtoCommandHomeSceneJumpReq                       ProtoCommand = 4528
	ProtoCommandHomeSceneJumpRsp                       ProtoCommand = 4698
	ProtoCommandHomeChooseModuleReq                    ProtoCommand = 4524
	ProtoCommandHomeChooseModuleRsp                    ProtoCommand = 4648
	ProtoCommandHomeModuleUnlockNotify                 ProtoCommand = 4560
	ProtoCommandHomeGetOnlineStatusReq                 ProtoCommand = 4820
	ProtoCommandHomeGetOnlineStatusRsp                 ProtoCommand = 4705
	ProtoCommandHomeKickPlayerReq                      ProtoCommand = 4870
	ProtoCommandHomeKickPlayerRsp                      ProtoCommand = 4691
	ProtoCommandHomeModuleSeenReq                      ProtoCommand = 4499
	ProtoCommandHomeModuleSeenRsp                      ProtoCommand = 4821
	ProtoCommandUnlockedFurnitureFormulaDataNotify     ProtoCommand = 4846
	ProtoCommandUnlockedFurnitureSuiteDataNotify       ProtoCommand = 4454
	ProtoCommandGetHomeLevelUpRewardReq                ProtoCommand = 4557
	ProtoCommandGetHomeLevelUpRewardRsp                ProtoCommand = 4603
	ProtoCommandGetFurnitureCurModuleArrangeCountReq   ProtoCommand = 4711
	ProtoCommandFurnitureCurModuleArrangeCountNotify   ProtoCommand = 4498
	ProtoCommandHomeComfortInfoNotify                  ProtoCommand = 4699
	ProtoCommandPlayerQuitFromHomeNotify               ProtoCommand = 4656
	ProtoCommandOtherPlayerEnterHomeNotify             ProtoCommand = 4628
	ProtoCommandHomePriorCheckNotify                   ProtoCommand = 4599
	ProtoCommandHomeMarkPointNotify                    ProtoCommand = 4474
	ProtoCommandHomeAllUnlockedBgmIdListNotify         ProtoCommand = 4608
	ProtoCommandHomeNewUnlockedBgmIdListNotify         ProtoCommand = 4847
	ProtoCommandHomeChangeBgmReq                       ProtoCommand = 4558
	ProtoCommandHomeChangeBgmRsp                       ProtoCommand = 4488
	ProtoCommandHomeChangeBgmNotify                    ProtoCommand = 4872
	ProtoCommandHomePreChangeEditModeNotify            ProtoCommand = 4639
	ProtoCommandHomeEnterEditModeFinishReq             ProtoCommand = 4537
	ProtoCommandHomeEnterEditModeFinishRsp             ProtoCommand = 4615
	ProtoCommandFurnitureMakeReq                       ProtoCommand = 4477
	ProtoCommandFurnitureMakeRsp                       ProtoCommand = 4782
	ProtoCommandTakeFurnitureMakeReq                   ProtoCommand = 4772
	ProtoCommandTakeFurnitureMakeRsp                   ProtoCommand = 4769
	ProtoCommandFurnitureMakeFinishNotify              ProtoCommand = 4841
	ProtoCommandFurnitureMakeStartReq                  ProtoCommand = 4633
	ProtoCommandFurnitureMakeStartRsp                  ProtoCommand = 4729
	ProtoCommandFurnitureMakeCancelReq                 ProtoCommand = 4555
	ProtoCommandFurnitureMakeCancelRsp                 ProtoCommand = 4683
	ProtoCommandFurnitureMakeBeHelpedNotify            ProtoCommand = 4578
	ProtoCommandFurnitureMakeHelpReq                   ProtoCommand = 4865
	ProtoCommandFurnitureMakeHelpRsp                   ProtoCommand = 4756
	ProtoCommandFunitureMakeMakeInfoChangeNotify       ProtoCommand = 4898
	ProtoCommandHomeLimitedShopInfoReq                 ProtoCommand = 4825
	ProtoCommandHomeLimitedShopInfoRsp                 ProtoCommand = 4796
	ProtoCommandHomeLimitedShopInfoNotify              ProtoCommand = 4887
	ProtoCommandHomeLimitedShopGoodsListReq            ProtoCommand = 4552
	ProtoCommandHomeLimitedShopGoodsListRsp            ProtoCommand = 4546
	ProtoCommandHomeLimitedShopBuyGoodsReq             ProtoCommand = 4760
	ProtoCommandHomeLimitedShopBuyGoodsRsp             ProtoCommand = 4750
	ProtoCommandHomeLimitedShopInfoChangeNotify        ProtoCommand = 4790
	ProtoCommandHomeResourceNotify                     ProtoCommand = 4892
	ProtoCommandHomeResourceTakeHomeCoinReq            ProtoCommand = 4479
	ProtoCommandHomeResourceTakeHomeCoinRsp            ProtoCommand = 4541
	ProtoCommandHomeResourceTakeFetterExpReq           ProtoCommand = 4768
	ProtoCommandHomeResourceTakeFetterExpRsp           ProtoCommand = 4645
	ProtoCommandHomeAvatarTalkFinishInfoNotify         ProtoCommand = 4896
	ProtoCommandHomeAvatarTalkReq                      ProtoCommand = 4688
	ProtoCommandHomeAvatarTalkRsp                      ProtoCommand = 4464
	ProtoCommandHomeAvatarRewardEventNotify            ProtoCommand = 4852
	ProtoCommandHomeAvatarRewardEventGetReq            ProtoCommand = 4551
	ProtoCommandHomeAvatarRewardEventGetRsp            ProtoCommand = 4833
	ProtoCommandHomeAvatarSummonAllEventNotify         ProtoCommand = 4808
	ProtoCommandHomeAvatarSummonEventReq               ProtoCommand = 4806
	ProtoCommandHomeAvatarSummonEventRsp               ProtoCommand = 4817
	ProtoCommandHomeAvatarCostumeChangeNotify          ProtoCommand = 4748
	ProtoCommandHomeAvatarSummonFinishReq              ProtoCommand = 4629
	ProtoCommandHomeAvatarSummonFinishRsp              ProtoCommand = 4696
	ProtoCommandHomeAvtarAllFinishRewardNotify         ProtoCommand = 4453
	ProtoCommandHomeAvatarAllFinishRewardNotify        ProtoCommand = 4741
	ProtoCommandHomeSceneInitFinishReq                 ProtoCommand = 4674
	ProtoCommandHomeSceneInitFinishRsp                 ProtoCommand = 4505
	ProtoCommandHomePlantSeedReq                       ProtoCommand = 4804
	ProtoCommandHomePlantSeedRsp                       ProtoCommand = 4556
	ProtoCommandHomePlantWeedReq                       ProtoCommand = 4640
	ProtoCommandHomePlantWeedRsp                       ProtoCommand = 4527
	ProtoCommandHomePlantInfoNotify                    ProtoCommand = 4587
	ProtoCommandHomePlantFieldNotify                   ProtoCommand = 4549
	ProtoCommandHomePlantInfoReq                       ProtoCommand = 4647
	ProtoCommandHomePlantInfoRsp                       ProtoCommand = 4701
	ProtoCommandHomeTransferReq                        ProtoCommand = 4726
	ProtoCommandHomeTransferRsp                        ProtoCommand = 4616
	ProtoCommandHomeGetFishFarmingInfoReq              ProtoCommand = 4476
	ProtoCommandHomeGetFishFarmingInfoRsp              ProtoCommand = 4678
	ProtoCommandHomeFishFarmingInfoNotify              ProtoCommand = 4677
	ProtoCommandHomeUpdateFishFarmingInfoReq           ProtoCommand = 4544
	ProtoCommandHomeUpdateFishFarmingInfoRsp           ProtoCommand = 4857
	ProtoCommandHomeUpdateScenePointFishFarmingInfoReq ProtoCommand = 4511
	ProtoCommandHomeUpdateScenePointFishFarmingInfoRsp ProtoCommand = 4540
	ProtoCommandHomeScenePointFishFarmingInfoNotify    ProtoCommand = 4547
	ProtoCommandHomeCustomFurnitureInfoNotify          ProtoCommand = 4712
	ProtoCommandHomeEditCustomFurnitureReq             ProtoCommand = 4724
	ProtoCommandHomeEditCustomFurnitureRsp             ProtoCommand = 4496
	ProtoCommandHomePictureFrameInfoNotify             ProtoCommand = 4878
	ProtoCommandHomeUpdatePictureFrameInfoReq          ProtoCommand = 4486
	ProtoCommandHomeUpdatePictureFrameInfoRsp          ProtoCommand = 4641
	ProtoCommandHomeRacingGallerySettleNotify          ProtoCommand = 4805
	ProtoCommandHomeGetGroupRecordReq                  ProtoCommand = 4523
	ProtoCommandHomeGetGroupRecordRsp                  ProtoCommand = 4538
	ProtoCommandHomeClearGroupRecordReq                ProtoCommand = 4759
	ProtoCommandHomeClearGroupRecordRsp                ProtoCommand = 4605
	ProtoCommandHomeBalloonGallerySettleNotify         ProtoCommand = 4811
	ProtoCommandHomeBalloonGalleryScoreNotify          ProtoCommand = 4654
	ProtoCommandHomeSeekFurnitureGalleryScoreNotify    ProtoCommand = 4583
	ProtoCommandGetHomeExchangeWoodInfoReq             ProtoCommand = 4473
	ProtoCommandGetHomeExchangeWoodInfoRsp             ProtoCommand = 4659
	ProtoCommandHomeExchangeWoodReq                    ProtoCommand = 4576
	ProtoCommandHomeExchangeWoodRsp                    ProtoCommand = 4622
	ProtoCommandHomeGetBlueprintSlotInfoReq            ProtoCommand = 4584
	ProtoCommandHomeGetBlueprintSlotInfoRsp            ProtoCommand = 4662
	ProtoCommandHomeSetBlueprintSlotOptionReq          ProtoCommand = 4798
	ProtoCommandHomeSetBlueprintSlotOptionRsp          ProtoCommand = 4786
	ProtoCommandHomeSetBlueprintFriendOptionReq        ProtoCommand = 4554
	ProtoCommandHomeSetBlueprintFriendOptionRsp        ProtoCommand = 4604
	ProtoCommandHomeBlueprintInfoNotify                ProtoCommand = 4765
	ProtoCommandHomePreviewBlueprintReq                ProtoCommand = 4478
	ProtoCommandHomePreviewBlueprintRsp                ProtoCommand = 4738
	ProtoCommandHomeCreateBlueprintReq                 ProtoCommand = 4619
	ProtoCommandHomeCreateBlueprintRsp                 ProtoCommand = 4606
	ProtoCommandHomeDeleteBlueprintReq                 ProtoCommand = 4502
	ProtoCommandHomeDeleteBlueprintRsp                 ProtoCommand = 4586
	ProtoCommandHomeSearchBlueprintReq                 ProtoCommand = 4889
	ProtoCommandHomeSearchBlueprintRsp                 ProtoCommand = 4593
	ProtoCommandHomeSaveArrangementNoChangeReq         ProtoCommand = 4704
	ProtoCommandHomeSaveArrangementNoChangeRsp         ProtoCommand = 4668
)

func (*TryEnterHomeReq) ProtoCommand() ProtoCommand         { return ProtoCommandTryEnterHomeReq }
func (*TryEnterHomeReq) ProtoMessageType() ProtoMessageType { return "TryEnterHomeReq" }

func (*TryEnterHomeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTryEnterHomeRsp }
func (*TryEnterHomeRsp) ProtoMessageType() ProtoMessageType { return "TryEnterHomeRsp" }

func (*JoinHomeWorldFailNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandJoinHomeWorldFailNotify
}
func (*JoinHomeWorldFailNotify) ProtoMessageType() ProtoMessageType { return "JoinHomeWorldFailNotify" }

func (*HomeBlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeBlockNotify }
func (*HomeBlockNotify) ProtoMessageType() ProtoMessageType { return "HomeBlockNotify" }

func (*HomeGetBasicInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeGetBasicInfoReq }
func (*HomeGetBasicInfoReq) ProtoMessageType() ProtoMessageType { return "HomeGetBasicInfoReq" }

func (*HomeBasicInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeBasicInfoNotify }
func (*HomeBasicInfoNotify) ProtoMessageType() ProtoMessageType { return "HomeBasicInfoNotify" }

func (*HomeChangeEditModeReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeEditModeReq }
func (*HomeChangeEditModeReq) ProtoMessageType() ProtoMessageType { return "HomeChangeEditModeReq" }

func (*HomeChangeEditModeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeEditModeRsp }
func (*HomeChangeEditModeRsp) ProtoMessageType() ProtoMessageType { return "HomeChangeEditModeRsp" }

func (*HomeChangeModuleReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeModuleReq }
func (*HomeChangeModuleReq) ProtoMessageType() ProtoMessageType { return "HomeChangeModuleReq" }

func (*HomeChangeModuleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeModuleRsp }
func (*HomeChangeModuleRsp) ProtoMessageType() ProtoMessageType { return "HomeChangeModuleRsp" }

func (*HomeGetArrangementInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetArrangementInfoReq
}
func (*HomeGetArrangementInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeGetArrangementInfoReq"
}

func (*HomeGetArrangementInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetArrangementInfoRsp
}
func (*HomeGetArrangementInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeGetArrangementInfoRsp"
}

func (*HomeUpdateArrangementInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateArrangementInfoReq
}
func (*HomeUpdateArrangementInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateArrangementInfoReq"
}

func (*HomeUpdateArrangementInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateArrangementInfoRsp
}
func (*HomeUpdateArrangementInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateArrangementInfoRsp"
}

func (*GetPlayerHomeCompInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerHomeCompInfoReq
}
func (*GetPlayerHomeCompInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerHomeCompInfoReq"
}

func (*PlayerHomeCompInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerHomeCompInfoNotify
}
func (*PlayerHomeCompInfoNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerHomeCompInfoNotify"
}

func (*SetFriendEnterHomeOptionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetFriendEnterHomeOptionReq
}
func (*SetFriendEnterHomeOptionReq) ProtoMessageType() ProtoMessageType {
	return "SetFriendEnterHomeOptionReq"
}

func (*SetFriendEnterHomeOptionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetFriendEnterHomeOptionRsp
}
func (*SetFriendEnterHomeOptionRsp) ProtoMessageType() ProtoMessageType {
	return "SetFriendEnterHomeOptionRsp"
}

func (*PlayerApplyEnterHomeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterHomeNotify
}
func (*PlayerApplyEnterHomeNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterHomeNotify"
}

func (*PlayerApplyEnterHomeResultReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterHomeResultReq
}
func (*PlayerApplyEnterHomeResultReq) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterHomeResultReq"
}

func (*PlayerApplyEnterHomeResultRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterHomeResultRsp
}
func (*PlayerApplyEnterHomeResultRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterHomeResultRsp"
}

func (*PlayerApplyEnterHomeResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterHomeResultNotify
}
func (*PlayerApplyEnterHomeResultNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterHomeResultNotify"
}

func (*HomeSceneJumpReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSceneJumpReq }
func (*HomeSceneJumpReq) ProtoMessageType() ProtoMessageType { return "HomeSceneJumpReq" }

func (*HomeSceneJumpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSceneJumpRsp }
func (*HomeSceneJumpRsp) ProtoMessageType() ProtoMessageType { return "HomeSceneJumpRsp" }

func (*HomeChooseModuleReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChooseModuleReq }
func (*HomeChooseModuleReq) ProtoMessageType() ProtoMessageType { return "HomeChooseModuleReq" }

func (*HomeChooseModuleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChooseModuleRsp }
func (*HomeChooseModuleRsp) ProtoMessageType() ProtoMessageType { return "HomeChooseModuleRsp" }

func (*HomeModuleUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeModuleUnlockNotify }
func (*HomeModuleUnlockNotify) ProtoMessageType() ProtoMessageType { return "HomeModuleUnlockNotify" }

func (*HomeGetOnlineStatusReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeGetOnlineStatusReq }
func (*HomeGetOnlineStatusReq) ProtoMessageType() ProtoMessageType { return "HomeGetOnlineStatusReq" }

func (*HomeGetOnlineStatusRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeGetOnlineStatusRsp }
func (*HomeGetOnlineStatusRsp) ProtoMessageType() ProtoMessageType { return "HomeGetOnlineStatusRsp" }

func (*HomeKickPlayerReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeKickPlayerReq }
func (*HomeKickPlayerReq) ProtoMessageType() ProtoMessageType { return "HomeKickPlayerReq" }

func (*HomeKickPlayerRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeKickPlayerRsp }
func (*HomeKickPlayerRsp) ProtoMessageType() ProtoMessageType { return "HomeKickPlayerRsp" }

func (*HomeModuleSeenReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeModuleSeenReq }
func (*HomeModuleSeenReq) ProtoMessageType() ProtoMessageType { return "HomeModuleSeenReq" }

func (*HomeModuleSeenRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeModuleSeenRsp }
func (*HomeModuleSeenRsp) ProtoMessageType() ProtoMessageType { return "HomeModuleSeenRsp" }

func (*UnlockedFurnitureFormulaDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUnlockedFurnitureFormulaDataNotify
}
func (*UnlockedFurnitureFormulaDataNotify) ProtoMessageType() ProtoMessageType {
	return "UnlockedFurnitureFormulaDataNotify"
}

func (*UnlockedFurnitureSuiteDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUnlockedFurnitureSuiteDataNotify
}
func (*UnlockedFurnitureSuiteDataNotify) ProtoMessageType() ProtoMessageType {
	return "UnlockedFurnitureSuiteDataNotify"
}

func (*GetHomeLevelUpRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHomeLevelUpRewardReq
}
func (*GetHomeLevelUpRewardReq) ProtoMessageType() ProtoMessageType { return "GetHomeLevelUpRewardReq" }

func (*GetHomeLevelUpRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHomeLevelUpRewardRsp
}
func (*GetHomeLevelUpRewardRsp) ProtoMessageType() ProtoMessageType { return "GetHomeLevelUpRewardRsp" }

func (*GetFurnitureCurModuleArrangeCountReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetFurnitureCurModuleArrangeCountReq
}
func (*GetFurnitureCurModuleArrangeCountReq) ProtoMessageType() ProtoMessageType {
	return "GetFurnitureCurModuleArrangeCountReq"
}

func (*FurnitureCurModuleArrangeCountNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFurnitureCurModuleArrangeCountNotify
}
func (*FurnitureCurModuleArrangeCountNotify) ProtoMessageType() ProtoMessageType {
	return "FurnitureCurModuleArrangeCountNotify"
}

func (*HomeComfortInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeComfortInfoNotify }
func (*HomeComfortInfoNotify) ProtoMessageType() ProtoMessageType { return "HomeComfortInfoNotify" }

func (*PlayerQuitFromHomeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerQuitFromHomeNotify
}
func (*PlayerQuitFromHomeNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerQuitFromHomeNotify"
}

func (*OtherPlayerEnterHomeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandOtherPlayerEnterHomeNotify
}
func (*OtherPlayerEnterHomeNotify) ProtoMessageType() ProtoMessageType {
	return "OtherPlayerEnterHomeNotify"
}

func (*HomePriorCheckNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomePriorCheckNotify }
func (*HomePriorCheckNotify) ProtoMessageType() ProtoMessageType { return "HomePriorCheckNotify" }

func (*HomeMarkPointNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeMarkPointNotify }
func (*HomeMarkPointNotify) ProtoMessageType() ProtoMessageType { return "HomeMarkPointNotify" }

func (*HomeAllUnlockedBgmIdListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAllUnlockedBgmIdListNotify
}
func (*HomeAllUnlockedBgmIdListNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAllUnlockedBgmIdListNotify"
}

func (*HomeNewUnlockedBgmIdListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeNewUnlockedBgmIdListNotify
}
func (*HomeNewUnlockedBgmIdListNotify) ProtoMessageType() ProtoMessageType {
	return "HomeNewUnlockedBgmIdListNotify"
}

func (*HomeChangeBgmReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeBgmReq }
func (*HomeChangeBgmReq) ProtoMessageType() ProtoMessageType { return "HomeChangeBgmReq" }

func (*HomeChangeBgmRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeBgmRsp }
func (*HomeChangeBgmRsp) ProtoMessageType() ProtoMessageType { return "HomeChangeBgmRsp" }

func (*HomeChangeBgmNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeChangeBgmNotify }
func (*HomeChangeBgmNotify) ProtoMessageType() ProtoMessageType { return "HomeChangeBgmNotify" }

func (*HomePreChangeEditModeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomePreChangeEditModeNotify
}
func (*HomePreChangeEditModeNotify) ProtoMessageType() ProtoMessageType {
	return "HomePreChangeEditModeNotify"
}

func (*HomeEnterEditModeFinishReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeEnterEditModeFinishReq
}
func (*HomeEnterEditModeFinishReq) ProtoMessageType() ProtoMessageType {
	return "HomeEnterEditModeFinishReq"
}

func (*HomeEnterEditModeFinishRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeEnterEditModeFinishRsp
}
func (*HomeEnterEditModeFinishRsp) ProtoMessageType() ProtoMessageType {
	return "HomeEnterEditModeFinishRsp"
}

func (*FurnitureMakeReq) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeReq }
func (*FurnitureMakeReq) ProtoMessageType() ProtoMessageType { return "FurnitureMakeReq" }

func (*FurnitureMakeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeRsp }
func (*FurnitureMakeRsp) ProtoMessageType() ProtoMessageType { return "FurnitureMakeRsp" }

func (*TakeFurnitureMakeReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeFurnitureMakeReq }
func (*TakeFurnitureMakeReq) ProtoMessageType() ProtoMessageType { return "TakeFurnitureMakeReq" }

func (*TakeFurnitureMakeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeFurnitureMakeRsp }
func (*TakeFurnitureMakeRsp) ProtoMessageType() ProtoMessageType { return "TakeFurnitureMakeRsp" }

func (*FurnitureMakeFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFurnitureMakeFinishNotify
}
func (*FurnitureMakeFinishNotify) ProtoMessageType() ProtoMessageType {
	return "FurnitureMakeFinishNotify"
}

func (*FurnitureMakeStartReq) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeStartReq }
func (*FurnitureMakeStartReq) ProtoMessageType() ProtoMessageType { return "FurnitureMakeStartReq" }

func (*FurnitureMakeStartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeStartRsp }
func (*FurnitureMakeStartRsp) ProtoMessageType() ProtoMessageType { return "FurnitureMakeStartRsp" }

func (*FurnitureMakeCancelReq) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeCancelReq }
func (*FurnitureMakeCancelReq) ProtoMessageType() ProtoMessageType { return "FurnitureMakeCancelReq" }

func (*FurnitureMakeCancelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeCancelRsp }
func (*FurnitureMakeCancelRsp) ProtoMessageType() ProtoMessageType { return "FurnitureMakeCancelRsp" }

func (*FurnitureMakeBeHelpedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFurnitureMakeBeHelpedNotify
}
func (*FurnitureMakeBeHelpedNotify) ProtoMessageType() ProtoMessageType {
	return "FurnitureMakeBeHelpedNotify"
}

func (*FurnitureMakeHelpReq) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeHelpReq }
func (*FurnitureMakeHelpReq) ProtoMessageType() ProtoMessageType { return "FurnitureMakeHelpReq" }

func (*FurnitureMakeHelpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFurnitureMakeHelpRsp }
func (*FurnitureMakeHelpRsp) ProtoMessageType() ProtoMessageType { return "FurnitureMakeHelpRsp" }

func (*FunitureMakeMakeInfoChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFunitureMakeMakeInfoChangeNotify
}
func (*FunitureMakeMakeInfoChangeNotify) ProtoMessageType() ProtoMessageType {
	return "FunitureMakeMakeInfoChangeNotify"
}

func (*HomeLimitedShopInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeLimitedShopInfoReq }
func (*HomeLimitedShopInfoReq) ProtoMessageType() ProtoMessageType { return "HomeLimitedShopInfoReq" }

func (*HomeLimitedShopInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeLimitedShopInfoRsp }
func (*HomeLimitedShopInfoRsp) ProtoMessageType() ProtoMessageType { return "HomeLimitedShopInfoRsp" }

func (*HomeLimitedShopInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopInfoNotify
}
func (*HomeLimitedShopInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopInfoNotify"
}

func (*HomeLimitedShopGoodsListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopGoodsListReq
}
func (*HomeLimitedShopGoodsListReq) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopGoodsListReq"
}

func (*HomeLimitedShopGoodsListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopGoodsListRsp
}
func (*HomeLimitedShopGoodsListRsp) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopGoodsListRsp"
}

func (*HomeLimitedShopBuyGoodsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopBuyGoodsReq
}
func (*HomeLimitedShopBuyGoodsReq) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopBuyGoodsReq"
}

func (*HomeLimitedShopBuyGoodsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopBuyGoodsRsp
}
func (*HomeLimitedShopBuyGoodsRsp) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopBuyGoodsRsp"
}

func (*HomeLimitedShopInfoChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeLimitedShopInfoChangeNotify
}
func (*HomeLimitedShopInfoChangeNotify) ProtoMessageType() ProtoMessageType {
	return "HomeLimitedShopInfoChangeNotify"
}

func (*HomeResourceNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomeResourceNotify }
func (*HomeResourceNotify) ProtoMessageType() ProtoMessageType { return "HomeResourceNotify" }

func (*HomeResourceTakeHomeCoinReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeResourceTakeHomeCoinReq
}
func (*HomeResourceTakeHomeCoinReq) ProtoMessageType() ProtoMessageType {
	return "HomeResourceTakeHomeCoinReq"
}

func (*HomeResourceTakeHomeCoinRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeResourceTakeHomeCoinRsp
}
func (*HomeResourceTakeHomeCoinRsp) ProtoMessageType() ProtoMessageType {
	return "HomeResourceTakeHomeCoinRsp"
}

func (*HomeResourceTakeFetterExpReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeResourceTakeFetterExpReq
}
func (*HomeResourceTakeFetterExpReq) ProtoMessageType() ProtoMessageType {
	return "HomeResourceTakeFetterExpReq"
}

func (*HomeResourceTakeFetterExpRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeResourceTakeFetterExpRsp
}
func (*HomeResourceTakeFetterExpRsp) ProtoMessageType() ProtoMessageType {
	return "HomeResourceTakeFetterExpRsp"
}

func (*HomeAvatarTalkFinishInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarTalkFinishInfoNotify
}
func (*HomeAvatarTalkFinishInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarTalkFinishInfoNotify"
}

func (*HomeAvatarTalkReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeAvatarTalkReq }
func (*HomeAvatarTalkReq) ProtoMessageType() ProtoMessageType { return "HomeAvatarTalkReq" }

func (*HomeAvatarTalkRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeAvatarTalkRsp }
func (*HomeAvatarTalkRsp) ProtoMessageType() ProtoMessageType { return "HomeAvatarTalkRsp" }

func (*HomeAvatarRewardEventNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarRewardEventNotify
}
func (*HomeAvatarRewardEventNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarRewardEventNotify"
}

func (*HomeAvatarRewardEventGetReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarRewardEventGetReq
}
func (*HomeAvatarRewardEventGetReq) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarRewardEventGetReq"
}

func (*HomeAvatarRewardEventGetRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarRewardEventGetRsp
}
func (*HomeAvatarRewardEventGetRsp) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarRewardEventGetRsp"
}

func (*HomeAvatarSummonAllEventNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarSummonAllEventNotify
}
func (*HomeAvatarSummonAllEventNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarSummonAllEventNotify"
}

func (*HomeAvatarSummonEventReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarSummonEventReq
}
func (*HomeAvatarSummonEventReq) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarSummonEventReq"
}

func (*HomeAvatarSummonEventRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarSummonEventRsp
}
func (*HomeAvatarSummonEventRsp) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarSummonEventRsp"
}

func (*HomeAvatarCostumeChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarCostumeChangeNotify
}
func (*HomeAvatarCostumeChangeNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarCostumeChangeNotify"
}

func (*HomeAvatarSummonFinishReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarSummonFinishReq
}
func (*HomeAvatarSummonFinishReq) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarSummonFinishReq"
}

func (*HomeAvatarSummonFinishRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarSummonFinishRsp
}
func (*HomeAvatarSummonFinishRsp) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarSummonFinishRsp"
}

func (*HomeAvtarAllFinishRewardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvtarAllFinishRewardNotify
}
func (*HomeAvtarAllFinishRewardNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvtarAllFinishRewardNotify"
}

func (*HomeAvatarAllFinishRewardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeAvatarAllFinishRewardNotify
}
func (*HomeAvatarAllFinishRewardNotify) ProtoMessageType() ProtoMessageType {
	return "HomeAvatarAllFinishRewardNotify"
}

func (*HomeSceneInitFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSceneInitFinishReq }
func (*HomeSceneInitFinishReq) ProtoMessageType() ProtoMessageType { return "HomeSceneInitFinishReq" }

func (*HomeSceneInitFinishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSceneInitFinishRsp }
func (*HomeSceneInitFinishRsp) ProtoMessageType() ProtoMessageType { return "HomeSceneInitFinishRsp" }

func (*HomePlantSeedReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantSeedReq }
func (*HomePlantSeedReq) ProtoMessageType() ProtoMessageType { return "HomePlantSeedReq" }

func (*HomePlantSeedRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantSeedRsp }
func (*HomePlantSeedRsp) ProtoMessageType() ProtoMessageType { return "HomePlantSeedRsp" }

func (*HomePlantWeedReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantWeedReq }
func (*HomePlantWeedReq) ProtoMessageType() ProtoMessageType { return "HomePlantWeedReq" }

func (*HomePlantWeedRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantWeedRsp }
func (*HomePlantWeedRsp) ProtoMessageType() ProtoMessageType { return "HomePlantWeedRsp" }

func (*HomePlantInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantInfoNotify }
func (*HomePlantInfoNotify) ProtoMessageType() ProtoMessageType { return "HomePlantInfoNotify" }

func (*HomePlantFieldNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantFieldNotify }
func (*HomePlantFieldNotify) ProtoMessageType() ProtoMessageType { return "HomePlantFieldNotify" }

func (*HomePlantInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantInfoReq }
func (*HomePlantInfoReq) ProtoMessageType() ProtoMessageType { return "HomePlantInfoReq" }

func (*HomePlantInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomePlantInfoRsp }
func (*HomePlantInfoRsp) ProtoMessageType() ProtoMessageType { return "HomePlantInfoRsp" }

func (*HomeTransferReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeTransferReq }
func (*HomeTransferReq) ProtoMessageType() ProtoMessageType { return "HomeTransferReq" }

func (*HomeTransferRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeTransferRsp }
func (*HomeTransferRsp) ProtoMessageType() ProtoMessageType { return "HomeTransferRsp" }

func (*HomeGetFishFarmingInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetFishFarmingInfoReq
}
func (*HomeGetFishFarmingInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeGetFishFarmingInfoReq"
}

func (*HomeGetFishFarmingInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetFishFarmingInfoRsp
}
func (*HomeGetFishFarmingInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeGetFishFarmingInfoRsp"
}

func (*HomeFishFarmingInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeFishFarmingInfoNotify
}
func (*HomeFishFarmingInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomeFishFarmingInfoNotify"
}

func (*HomeUpdateFishFarmingInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateFishFarmingInfoReq
}
func (*HomeUpdateFishFarmingInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateFishFarmingInfoReq"
}

func (*HomeUpdateFishFarmingInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateFishFarmingInfoRsp
}
func (*HomeUpdateFishFarmingInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateFishFarmingInfoRsp"
}

func (*HomeUpdateScenePointFishFarmingInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateScenePointFishFarmingInfoReq
}
func (*HomeUpdateScenePointFishFarmingInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateScenePointFishFarmingInfoReq"
}

func (*HomeUpdateScenePointFishFarmingInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdateScenePointFishFarmingInfoRsp
}
func (*HomeUpdateScenePointFishFarmingInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeUpdateScenePointFishFarmingInfoRsp"
}

func (*HomeScenePointFishFarmingInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeScenePointFishFarmingInfoNotify
}
func (*HomeScenePointFishFarmingInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomeScenePointFishFarmingInfoNotify"
}

func (*HomeCustomFurnitureInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeCustomFurnitureInfoNotify
}
func (*HomeCustomFurnitureInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomeCustomFurnitureInfoNotify"
}

func (*HomeEditCustomFurnitureReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeEditCustomFurnitureReq
}
func (*HomeEditCustomFurnitureReq) ProtoMessageType() ProtoMessageType {
	return "HomeEditCustomFurnitureReq"
}

func (*HomeEditCustomFurnitureRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeEditCustomFurnitureRsp
}
func (*HomeEditCustomFurnitureRsp) ProtoMessageType() ProtoMessageType {
	return "HomeEditCustomFurnitureRsp"
}

func (*HomePictureFrameInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomePictureFrameInfoNotify
}
func (*HomePictureFrameInfoNotify) ProtoMessageType() ProtoMessageType {
	return "HomePictureFrameInfoNotify"
}

func (*HomeUpdatePictureFrameInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdatePictureFrameInfoReq
}
func (*HomeUpdatePictureFrameInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeUpdatePictureFrameInfoReq"
}

func (*HomeUpdatePictureFrameInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeUpdatePictureFrameInfoRsp
}
func (*HomeUpdatePictureFrameInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeUpdatePictureFrameInfoRsp"
}

func (*HomeRacingGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeRacingGallerySettleNotify
}
func (*HomeRacingGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "HomeRacingGallerySettleNotify"
}

func (*HomeGetGroupRecordReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeGetGroupRecordReq }
func (*HomeGetGroupRecordReq) ProtoMessageType() ProtoMessageType { return "HomeGetGroupRecordReq" }

func (*HomeGetGroupRecordRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeGetGroupRecordRsp }
func (*HomeGetGroupRecordRsp) ProtoMessageType() ProtoMessageType { return "HomeGetGroupRecordRsp" }

func (*HomeClearGroupRecordReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeClearGroupRecordReq
}
func (*HomeClearGroupRecordReq) ProtoMessageType() ProtoMessageType { return "HomeClearGroupRecordReq" }

func (*HomeClearGroupRecordRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeClearGroupRecordRsp
}
func (*HomeClearGroupRecordRsp) ProtoMessageType() ProtoMessageType { return "HomeClearGroupRecordRsp" }

func (*HomeBalloonGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeBalloonGallerySettleNotify
}
func (*HomeBalloonGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "HomeBalloonGallerySettleNotify"
}

func (*HomeBalloonGalleryScoreNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeBalloonGalleryScoreNotify
}
func (*HomeBalloonGalleryScoreNotify) ProtoMessageType() ProtoMessageType {
	return "HomeBalloonGalleryScoreNotify"
}

func (*HomeSeekFurnitureGalleryScoreNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSeekFurnitureGalleryScoreNotify
}
func (*HomeSeekFurnitureGalleryScoreNotify) ProtoMessageType() ProtoMessageType {
	return "HomeSeekFurnitureGalleryScoreNotify"
}

func (*GetHomeExchangeWoodInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHomeExchangeWoodInfoReq
}
func (*GetHomeExchangeWoodInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetHomeExchangeWoodInfoReq"
}

func (*GetHomeExchangeWoodInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHomeExchangeWoodInfoRsp
}
func (*GetHomeExchangeWoodInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetHomeExchangeWoodInfoRsp"
}

func (*HomeExchangeWoodReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeExchangeWoodReq }
func (*HomeExchangeWoodReq) ProtoMessageType() ProtoMessageType { return "HomeExchangeWoodReq" }

func (*HomeExchangeWoodRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeExchangeWoodRsp }
func (*HomeExchangeWoodRsp) ProtoMessageType() ProtoMessageType { return "HomeExchangeWoodRsp" }

func (*HomeGetBlueprintSlotInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetBlueprintSlotInfoReq
}
func (*HomeGetBlueprintSlotInfoReq) ProtoMessageType() ProtoMessageType {
	return "HomeGetBlueprintSlotInfoReq"
}

func (*HomeGetBlueprintSlotInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGetBlueprintSlotInfoRsp
}
func (*HomeGetBlueprintSlotInfoRsp) ProtoMessageType() ProtoMessageType {
	return "HomeGetBlueprintSlotInfoRsp"
}

func (*HomeSetBlueprintSlotOptionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSetBlueprintSlotOptionReq
}
func (*HomeSetBlueprintSlotOptionReq) ProtoMessageType() ProtoMessageType {
	return "HomeSetBlueprintSlotOptionReq"
}

func (*HomeSetBlueprintSlotOptionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSetBlueprintSlotOptionRsp
}
func (*HomeSetBlueprintSlotOptionRsp) ProtoMessageType() ProtoMessageType {
	return "HomeSetBlueprintSlotOptionRsp"
}

func (*HomeSetBlueprintFriendOptionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSetBlueprintFriendOptionReq
}
func (*HomeSetBlueprintFriendOptionReq) ProtoMessageType() ProtoMessageType {
	return "HomeSetBlueprintFriendOptionReq"
}

func (*HomeSetBlueprintFriendOptionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSetBlueprintFriendOptionRsp
}
func (*HomeSetBlueprintFriendOptionRsp) ProtoMessageType() ProtoMessageType {
	return "HomeSetBlueprintFriendOptionRsp"
}

func (*HomeBlueprintInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeBlueprintInfoNotify
}
func (*HomeBlueprintInfoNotify) ProtoMessageType() ProtoMessageType { return "HomeBlueprintInfoNotify" }

func (*HomePreviewBlueprintReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomePreviewBlueprintReq
}
func (*HomePreviewBlueprintReq) ProtoMessageType() ProtoMessageType { return "HomePreviewBlueprintReq" }

func (*HomePreviewBlueprintRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomePreviewBlueprintRsp
}
func (*HomePreviewBlueprintRsp) ProtoMessageType() ProtoMessageType { return "HomePreviewBlueprintRsp" }

func (*HomeCreateBlueprintReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeCreateBlueprintReq }
func (*HomeCreateBlueprintReq) ProtoMessageType() ProtoMessageType { return "HomeCreateBlueprintReq" }

func (*HomeCreateBlueprintRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeCreateBlueprintRsp }
func (*HomeCreateBlueprintRsp) ProtoMessageType() ProtoMessageType { return "HomeCreateBlueprintRsp" }

func (*HomeDeleteBlueprintReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeDeleteBlueprintReq }
func (*HomeDeleteBlueprintReq) ProtoMessageType() ProtoMessageType { return "HomeDeleteBlueprintReq" }

func (*HomeDeleteBlueprintRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeDeleteBlueprintRsp }
func (*HomeDeleteBlueprintRsp) ProtoMessageType() ProtoMessageType { return "HomeDeleteBlueprintRsp" }

func (*HomeSearchBlueprintReq) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSearchBlueprintReq }
func (*HomeSearchBlueprintReq) ProtoMessageType() ProtoMessageType { return "HomeSearchBlueprintReq" }

func (*HomeSearchBlueprintRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHomeSearchBlueprintRsp }
func (*HomeSearchBlueprintRsp) ProtoMessageType() ProtoMessageType { return "HomeSearchBlueprintRsp" }

func (*HomeSaveArrangementNoChangeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSaveArrangementNoChangeReq
}
func (*HomeSaveArrangementNoChangeReq) ProtoMessageType() ProtoMessageType {
	return "HomeSaveArrangementNoChangeReq"
}

func (*HomeSaveArrangementNoChangeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeSaveArrangementNoChangeRsp
}
func (*HomeSaveArrangementNoChangeRsp) ProtoMessageType() ProtoMessageType {
	return "HomeSaveArrangementNoChangeRsp"
}