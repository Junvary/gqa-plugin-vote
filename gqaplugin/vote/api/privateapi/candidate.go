package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CandidateList(c *gin.Context) {
	var getCandidateList model.RequestCandidateList
	if err := gqaModel.RequestShouldBindJSON(c, &getCandidateList); err != nil {
		return
	}
	if err, candidate, total := privateservice.CandidateList(getCandidateList, utils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取候选人列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取候选人列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  candidate,
			Page:     getCandidateList.Page,
			PageSize: getCandidateList.PageSize,
			Total:    total,
		}, c)
	}
}

func CandidateAdd(c *gin.Context) {
	var toAddCandidate model.RequestAddCandidate
	if err := gqaModel.RequestShouldBindJSON(c, &toAddCandidate); err != nil {
		return
	}
	if err := privateservice.CandidateAdd(&toAddCandidate, utils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加候选人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加候选人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加候选人成功！", c)
	}
}

func CandidateDelete(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.CandidateDelete(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除候选人失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除候选人失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除候选人成功！", c)
	}
}
