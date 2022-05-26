package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginVoteSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginVote").Or("name = ?", "GqaPluginVote").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中vote插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_menu 表中vote插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 701, Memo: "投票系统", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "GqaPluginVote", Title: "投票系统", Icon: "how_to_vote", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 1, Memo: "民主评议党员投票", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-vote-dy", ParentCode: "GqaPluginVote", Title: "民主评议党员投票", Icon: "how_to_vote",
		Path: "/plugin-vote/vote/vote-dy", Component: "plugins/Vote/VoteDy/index"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 2, Memo: "民主评议业务骨干及管理人员投票", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-vote-gl", ParentCode: "GqaPluginVote", Title: "民主评议业务骨干及管理人员投票", Icon: "how_to_vote",
		Path: "/plugin-vote/vote/vote-gl", Component: "plugins/Vote/VoteGl/index"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 3, Memo: "随机抽取投票人", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-voter-random", ParentCode: "GqaPluginVote", Title: "随机抽取投票人", Icon: "refresh",
		Path: "/plugin-vote/vote/voter-random", Component: "plugins/Vote/VoterRandom/index"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 4, Memo: "固定投票人配置", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-voter", ParentCode: "GqaPluginVote", Title: "固定投票人配置", Icon: "emoji_people",
		Path: "/plugin-vote/vote/voter", Component: "plugins/Vote/Voter/index"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 5, Memo: "候选人配置", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-candidate", ParentCode: "GqaPluginVote", Title: "候选人配置", Icon: "people",
		Path: "/plugin-vote/vote/candidate", Component: "plugins/Vote/Candidate/index"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Status: "on", Sort: 6, Memo: "投票结果", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, IsPlugin: "yes", Name: "plugin-vote-score-result", ParentCode: "GqaPluginVote", Title: "投票结果分析", Icon: "dvr",
		Path: "/plugin-vote/vote/score-result", Component: "plugins/Vote/ScoreResult/index"},
}
