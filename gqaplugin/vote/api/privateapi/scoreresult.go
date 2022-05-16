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

func VoteResultVoted(c *gin.Context) {
	var getVoteScoreVoted model.RequestScoreResultVoted
	if err := gqaModel.RequestShouldBindJSON(c, &getVoteScoreVoted); err != nil {
		return
	}
	if err, voter := privateservice.VoteResultVoted(getVoteScoreVoted); err != nil {
		gqaGlobal.GqaLogger.Error("获取已经投票人列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取已经投票人列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records: voter,
		}, c)
	}
}

func VoteResultList(c *gin.Context) {
	var getVoteScoreList model.RequestScoreResultList
	if err := gqaModel.RequestShouldBindJSON(c, &getVoteScoreList); err != nil {
		return
	}
	if err, voter, total := privateservice.VoteResultList(getVoteScoreList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取投票结果列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取投票结果列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  voter,
			Page:     getVoteScoreList.Page,
			PageSize: getVoteScoreList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoteResultChart(c *gin.Context) {
	var getVoteScoreList model.RequestScoreResultList
	if err := gqaModel.RequestShouldBindJSON(c, &getVoteScoreList); err != nil {
		return
	}
	if err, voter, total := privateservice.VoteResultChart(getVoteScoreList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取投票结果图形失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取投票结果图形失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  voter,
			Page:     getVoteScoreList.Page,
			PageSize: getVoteScoreList.PageSize,
			Total:    total,
		}, c)
	}
}
