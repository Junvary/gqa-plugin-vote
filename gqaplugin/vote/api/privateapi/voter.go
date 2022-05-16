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

func VoterList(c *gin.Context) {
	var getVoterList model.RequestVoterList
	if err := gqaModel.RequestShouldBindJSON(c, &getVoterList); err != nil {
		return
	}
	if err, voter, total := privateservice.VoterList(getVoterList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取投票人列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取投票人列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  voter,
			Page:     getVoterList.Page,
			PageSize: getVoterList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoterAdd(c *gin.Context) {
	var toAddVoter model.RequestAddVoter
	if err := gqaModel.RequestShouldBindJSON(c, &toAddVoter); err != nil {
		return
	}
	if err := privateservice.VoterAdd(&toAddVoter, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加投票人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加投票人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加投票人成功！", c)
	}
}

func VoterDelete(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.VoterDelete(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除投票人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除投票人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除投票人成功！", c)
	}
}
