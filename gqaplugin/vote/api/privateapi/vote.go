package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandle(c *gin.Context) {
	var toAddScore model.RequestAddScore
	if err := gqaModel.RequestShouldBindJSON(c, &toAddScore); err != nil {
		return
	}
	if err := privateservice.VoteHandle(&toAddScore, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("投票失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("投票失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("投票成功！", c)
	}
}

func CanVoteDy(c *gin.Context) {
	if err, canVote := privateservice.CanVoteDy(gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("检查投票结果失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("检查投票结果失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": canVote}, "检查投票结果成功！", c)
	}
}

func CanVoteGl(c *gin.Context) {
	if err, canVote := privateservice.CanVoteGl(gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("检查投票结果失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("检查投票结果失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": canVote}, "检查投票结果成功！", c)
	}
}
