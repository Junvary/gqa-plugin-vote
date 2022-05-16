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

func VoterRandomList(c *gin.Context) {
	var getVoterRandomList model.RequestVoterRandomList
	if err := gqaModel.RequestShouldBindJSON(c, &getVoterRandomList); err != nil {
		return
	}
	if err, voter, total := privateservice.VoterRandomList(getVoterRandomList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取随机投票人列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取随机投票人列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  voter,
			Page:     getVoterRandomList.Page,
			PageSize: getVoterRandomList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoterRandomGet(c *gin.Context) {
	var getVoterRandomNumber model.RequestGetVoterRandom
	if err := gqaModel.RequestShouldBindJSON(c, &getVoterRandomNumber); err != nil {
		return
	}
	if err, voter, total := privateservice.VoterRandomGet(getVoterRandomNumber, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("随机选择投票人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("随机选择投票人失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gin.H{"records": voter, "randomNumber": total}, c)
	}
}

func VoterRandomAdd(c *gin.Context) {
	var toAddVoterRandom model.RequestAddVoterRandom
	if err := gqaModel.RequestShouldBindJSON(c, &toAddVoterRandom); err != nil {
		return
	}
	if err := privateservice.VoterRandomAdd(&toAddVoterRandom, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加投票人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加投票人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加投票人成功！", c)
	}
}

func VoterRandomDelete(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.VoterRandomDelete(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除投票人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除投票人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除投票人成功！", c)
	}
}
