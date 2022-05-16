package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginVoteSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-vote").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中vote插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_api 表中vote插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 701, Memo: "插件vote：随机和固定投票人列表", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-list", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 702, Memo: "插件vote：获取随机投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-get", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 703, Memo: "插件vote：新增随机和固定投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-add", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 704, Memo: "插件vote：删除随机和固定投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-delete", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 705, Memo: "插件vote：投票", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-handle", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 706, Memo: "插件vote：检查是否可以投票-党员", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-dy", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 707, Memo: "插件vote：检查是否可以投票-管理", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-gl", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 708, Memo: "插件vote：投票结果查看已经投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-voted", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 709, Memo: "插件vote：投票结果列表", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-list", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 710, Memo: "插件vote：投票结果图形", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-chart", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 711, Memo: "插件vote：固定投票人列表", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-list", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 712, Memo: "插件vote：新增固定投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-add", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 713, Memo: "插件vote：删除固定投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-delete", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 714, Memo: "插件vote：候选人列表", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-list", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 715, Memo: "插件vote：新增候选人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-add", ApiMethod: "POST"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "yes", Sort: 716, Memo: "插件vote：删除候选人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-delete", ApiMethod: "POST"},
}
